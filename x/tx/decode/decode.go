package decode

import (
	"errors"

	"github.com/cosmos/cosmos-proto/anyutil"
	"google.golang.org/protobuf/proto"

	v1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/x/tx/signing"
)

// DecodedTx contains the decoded transaction, its signers, and other flags.
type DecodedTx struct {
	Messages                     []proto.Message
	Tx                           *v1beta1.Tx
	TxRaw                        *v1beta1.TxRaw
	Signers                      [][]byte
	TxBodyHasUnknownNonCriticals bool
}

// Decoder contains the dependencies required for decoding transactions.
type Decoder struct {
	signingCtx *signing.Context
}

// Options are options for creating a Decoder.
type Options struct {
	SigningContext *signing.Context
}

// NewDecoder creates a new Decoder for decoding transactions.
func NewDecoder(options Options) (*Decoder, error) {
	if options.SigningContext == nil {
		return nil, errors.New("signing context is required")
	}

	return &Decoder{
		signingCtx: options.SigningContext,
	}, nil
}

// Decode decodes raw protobuf encoded transaction bytes into a DecodedTx.
func (d *Decoder) Decode(txBytes []byte) (*DecodedTx, error) {
	// Make sure txBytes follow ADR-027.
	err := rejectNonADR027TxRaw(txBytes)
	if err != nil {
		return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
	}

	var raw v1beta1.TxRaw

	// reject all unknown proto fields in the root TxRaw
	fileResolver := d.signingCtx.FileResolver()
	err = RejectUnknownFieldsStrict(txBytes, raw.ProtoReflect().Descriptor(), fileResolver)
	if err != nil {
		return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
	}

	err = proto.Unmarshal(txBytes, &raw)
	if err != nil {
		return nil, err
	}

	var body v1beta1.TxBody

	// allow non-critical unknown fields in TxBody
	txBodyHasUnknownNonCriticals, err := RejectUnknownFields(raw.BodyBytes, body.ProtoReflect().Descriptor(), true, fileResolver)
	if err != nil {
		return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
	}

	err = proto.Unmarshal(raw.BodyBytes, &body)
	if err != nil {
		return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
	}

	var authInfo v1beta1.AuthInfo

	// reject all unknown proto fields in AuthInfo
	err = RejectUnknownFieldsStrict(raw.AuthInfoBytes, authInfo.ProtoReflect().Descriptor(), fileResolver)
	if err != nil {
		return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
	}

	err = proto.Unmarshal(raw.AuthInfoBytes, &authInfo)
	if err != nil {
		return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
	}

	theTx := &v1beta1.Tx{
		Body:       &body,
		AuthInfo:   &authInfo,
		Signatures: raw.Signatures,
	}

	var signers [][]byte
	var msgs []proto.Message
	seenSigners := map[string]struct{}{}
	for _, anyMsg := range body.Messages {
		msg, signerErr := anyutil.Unpack(anyMsg, fileResolver, d.signingCtx.TypeResolver())
		if signerErr != nil {
			return nil, errorsmod.Wrap(ErrTxDecode, signerErr.Error())
		}
		msgs = append(msgs, msg)
		ss, signerErr := d.signingCtx.GetSigners(msg)
		if signerErr != nil {
			return nil, errorsmod.Wrap(ErrTxDecode, signerErr.Error())
		}
		for _, s := range ss {
			_, seen := seenSigners[string(s)]
			if seen {
				continue
			}
			signers = append(signers, s)
			seenSigners[string(s)] = struct{}{}
		}
	}

	// If a fee payer is specified in the AuthInfo, it must be added to the list of signers
	if authInfo.Fee != nil && authInfo.Fee.Payer != "" {
		feeAddr, err := d.signingCtx.AddressCodec().StringToBytes(authInfo.Fee.Payer)
		if err != nil {
			return nil, errorsmod.Wrap(ErrTxDecode, err.Error())
		}

		if _, seen := seenSigners[string(feeAddr)]; !seen {
			signers = append(signers, feeAddr)
		}
	}

	return &DecodedTx{
		Messages:                     msgs,
		Tx:                           theTx,
		TxRaw:                        &raw,
		TxBodyHasUnknownNonCriticals: txBodyHasUnknownNonCriticals,
		Signers:                      signers,
	}, nil
}
