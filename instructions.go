// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Jupiter"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_MercurialExchange = ag_binary.TypeID([8]byte{31, 248, 60, 226, 215, 168, 55, 199})

	Instruction_SaberExchange = ag_binary.TypeID([8]byte{145, 158, 184, 212, 3, 74, 156, 118})

	Instruction_SaberSwap = ag_binary.TypeID([8]byte{64, 62, 98, 226, 52, 74, 37, 178})

	Instruction_SaberAddDecimalsDeposit = ag_binary.TypeID([8]byte{30, 130, 91, 146, 128, 80, 145, 232})

	Instruction_SaberAddDecimalsWithdraw = ag_binary.TypeID([8]byte{228, 191, 206, 12, 242, 90, 137, 34})

	Instruction_SenchaExchange = ag_binary.TypeID([8]byte{94, 31, 159, 252, 144, 67, 4, 100})

	Instruction_SerumSwap = ag_binary.TypeID([8]byte{88, 183, 70, 249, 214, 118, 82, 210})

	Instruction_TokenSwap = ag_binary.TypeID([8]byte{187, 192, 118, 212, 62, 109, 28, 213})

	Instruction_StepTokenSwap = ag_binary.TypeID([8]byte{55, 100, 17, 243, 242, 181, 43, 165})

	Instruction_CropperTokenSwap = ag_binary.TypeID([8]byte{167, 38, 59, 37, 132, 60, 95, 68})

	Instruction_RaydiumSwap = ag_binary.TypeID([8]byte{177, 173, 42, 240, 184, 4, 124, 81})

	Instruction_RaydiumSwapV2 = ag_binary.TypeID([8]byte{69, 227, 98, 93, 237, 202, 223, 140})

	Instruction_AldrinSwap = ag_binary.TypeID([8]byte{251, 232, 119, 166, 225, 185, 169, 161})

	Instruction_AldrinV2Swap = ag_binary.TypeID([8]byte{190, 166, 89, 139, 33, 152, 16, 10})

	Instruction_CremaTokenSwap = ag_binary.TypeID([8]byte{235, 160, 175, 122, 61, 177, 2, 247})

	Instruction_LifinityTokenSwap = ag_binary.TypeID([8]byte{0, 49, 246, 1, 36, 153, 11, 93})

	Instruction_CykuraSwap = ag_binary.TypeID([8]byte{38, 241, 21, 107, 120, 59, 184, 249})

	Instruction_WhirlpoolSwap = ag_binary.TypeID([8]byte{123, 229, 184, 63, 12, 0, 92, 145})

	Instruction_RiskCheckAndFee = ag_binary.TypeID([8]byte{81, 42, 179, 152, 221, 1, 181, 120})

	Instruction_InitializeTokenLedger = ag_binary.TypeID([8]byte{244, 63, 250, 192, 50, 44, 172, 250})

	Instruction_SetTokenLedger = ag_binary.TypeID([8]byte{228, 85, 185, 112, 78, 79, 77, 2})

	Instruction_CreateOpenOrders = ag_binary.TypeID([8]byte{229, 194, 212, 172, 8, 10, 134, 147})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_MercurialExchange:
		return "MercurialExchange"
	case Instruction_SaberExchange:
		return "SaberExchange"
	case Instruction_SaberSwap:
		return "SaberSwap"
	case Instruction_SaberAddDecimalsDeposit:
		return "SaberAddDecimalsDeposit"
	case Instruction_SaberAddDecimalsWithdraw:
		return "SaberAddDecimalsWithdraw"
	case Instruction_SenchaExchange:
		return "SenchaExchange"
	case Instruction_SerumSwap:
		return "SerumSwap"
	case Instruction_TokenSwap:
		return "TokenSwap"
	case Instruction_StepTokenSwap:
		return "StepTokenSwap"
	case Instruction_CropperTokenSwap:
		return "CropperTokenSwap"
	case Instruction_RaydiumSwap:
		return "RaydiumSwap"
	case Instruction_RaydiumSwapV2:
		return "RaydiumSwapV2"
	case Instruction_AldrinSwap:
		return "AldrinSwap"
	case Instruction_AldrinV2Swap:
		return "AldrinV2Swap"
	case Instruction_CremaTokenSwap:
		return "CremaTokenSwap"
	case Instruction_LifinityTokenSwap:
		return "LifinityTokenSwap"
	case Instruction_CykuraSwap:
		return "CykuraSwap"
	case Instruction_WhirlpoolSwap:
		return "WhirlpoolSwap"
	case Instruction_RiskCheckAndFee:
		return "RiskCheckAndFee"
	case Instruction_InitializeTokenLedger:
		return "InitializeTokenLedger"
	case Instruction_SetTokenLedger:
		return "SetTokenLedger"
	case Instruction_CreateOpenOrders:
		return "CreateOpenOrders"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"mercurial_exchange", (*MercurialExchange)(nil),
		},
		{
			"saber_exchange", (*SaberExchange)(nil),
		},
		{
			"saber_swap", (*SaberSwap)(nil),
		},
		{
			"saber_add_decimals_deposit", (*SaberAddDecimalsDeposit)(nil),
		},
		{
			"saber_add_decimals_withdraw", (*SaberAddDecimalsWithdraw)(nil),
		},
		{
			"sencha_exchange", (*SenchaExchange)(nil),
		},
		{
			"serum_swap", (*SerumSwap)(nil),
		},
		{
			"token_swap", (*TokenSwap)(nil),
		},
		{
			"step_token_swap", (*StepTokenSwap)(nil),
		},
		{
			"cropper_token_swap", (*CropperTokenSwap)(nil),
		},
		{
			"raydium_swap", (*RaydiumSwap)(nil),
		},
		{
			"raydium_swap_v2", (*RaydiumSwapV2)(nil),
		},
		{
			"aldrin_swap", (*AldrinSwap)(nil),
		},
		{
			"aldrin_v2_swap", (*AldrinV2Swap)(nil),
		},
		{
			"crema_token_swap", (*CremaTokenSwap)(nil),
		},
		{
			"lifinity_token_swap", (*LifinityTokenSwap)(nil),
		},
		{
			"cykura_swap", (*CykuraSwap)(nil),
		},
		{
			"whirlpool_swap", (*WhirlpoolSwap)(nil),
		},
		{
			"risk_check_and_fee", (*RiskCheckAndFee)(nil),
		},
		{
			"initialize_token_ledger", (*InitializeTokenLedger)(nil),
		},
		{
			"set_token_ledger", (*SetTokenLedger)(nil),
		},
		{
			"create_open_orders", (*CreateOpenOrders)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}