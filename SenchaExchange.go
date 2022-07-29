// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SenchaExchange is the `senchaExchange` instruction.
type SenchaExchange struct {
	InAmount         *uint64 `bin:"optional"`
	MinimumOutAmount *uint64
	PlatformFeeBps   *uint8

	// [0] = [] swapProgram
	//
	// [1] = [] tokenProgram
	//
	// [2] = [WRITE] swap
	//
	// [3] = [] userAuthority
	//
	// [4] = [WRITE] inputUserAccount
	//
	// [5] = [WRITE] inputTokenAccount
	//
	// [6] = [WRITE] inputFeesAccount
	//
	// [7] = [WRITE] outputUserAccount
	//
	// [8] = [WRITE] outputTokenAccount
	//
	// [9] = [WRITE] outputFeesAccount
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSenchaExchangeInstructionBuilder creates a new `SenchaExchange` instruction builder.
func NewSenchaExchangeInstructionBuilder() *SenchaExchange {
	nd := &SenchaExchange{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetInAmount sets the "inAmount" parameter.
func (inst *SenchaExchange) SetInAmount(inAmount uint64) *SenchaExchange {
	inst.InAmount = &inAmount
	return inst
}

// SetMinimumOutAmount sets the "minimumOutAmount" parameter.
func (inst *SenchaExchange) SetMinimumOutAmount(minimumOutAmount uint64) *SenchaExchange {
	inst.MinimumOutAmount = &minimumOutAmount
	return inst
}

// SetPlatformFeeBps sets the "platformFeeBps" parameter.
func (inst *SenchaExchange) SetPlatformFeeBps(platformFeeBps uint8) *SenchaExchange {
	inst.PlatformFeeBps = &platformFeeBps
	return inst
}

// SetSwapProgramAccount sets the "swapProgram" account.
func (inst *SenchaExchange) SetSwapProgramAccount(swapProgram ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swapProgram)
	return inst
}

// GetSwapProgramAccount gets the "swapProgram" account.
func (inst *SenchaExchange) GetSwapProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *SenchaExchange) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *SenchaExchange) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSwapAccount sets the "swap" account.
func (inst *SenchaExchange) SetSwapAccount(swap ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(swap).WRITE()
	return inst
}

// GetSwapAccount gets the "swap" account.
func (inst *SenchaExchange) GetSwapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserAuthorityAccount sets the "userAuthority" account.
func (inst *SenchaExchange) SetUserAuthorityAccount(userAuthority ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userAuthority)
	return inst
}

// GetUserAuthorityAccount gets the "userAuthority" account.
func (inst *SenchaExchange) GetUserAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetInputUserAccount sets the "inputUserAccount" account.
func (inst *SenchaExchange) SetInputUserAccount(inputUserAccount ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(inputUserAccount).WRITE()
	return inst
}

// GetInputUserAccount gets the "inputUserAccount" account.
func (inst *SenchaExchange) GetInputUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetInputTokenAccount sets the "inputTokenAccount" account.
func (inst *SenchaExchange) SetInputTokenAccount(inputTokenAccount ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(inputTokenAccount).WRITE()
	return inst
}

// GetInputTokenAccount gets the "inputTokenAccount" account.
func (inst *SenchaExchange) GetInputTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetInputFeesAccount sets the "inputFeesAccount" account.
func (inst *SenchaExchange) SetInputFeesAccount(inputFeesAccount ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(inputFeesAccount).WRITE()
	return inst
}

// GetInputFeesAccount gets the "inputFeesAccount" account.
func (inst *SenchaExchange) GetInputFeesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetOutputUserAccount sets the "outputUserAccount" account.
func (inst *SenchaExchange) SetOutputUserAccount(outputUserAccount ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(outputUserAccount).WRITE()
	return inst
}

// GetOutputUserAccount gets the "outputUserAccount" account.
func (inst *SenchaExchange) GetOutputUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetOutputTokenAccount sets the "outputTokenAccount" account.
func (inst *SenchaExchange) SetOutputTokenAccount(outputTokenAccount ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(outputTokenAccount).WRITE()
	return inst
}

// GetOutputTokenAccount gets the "outputTokenAccount" account.
func (inst *SenchaExchange) GetOutputTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetOutputFeesAccount sets the "outputFeesAccount" account.
func (inst *SenchaExchange) SetOutputFeesAccount(outputFeesAccount ag_solanago.PublicKey) *SenchaExchange {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(outputFeesAccount).WRITE()
	return inst
}

// GetOutputFeesAccount gets the "outputFeesAccount" account.
func (inst *SenchaExchange) GetOutputFeesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst SenchaExchange) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SenchaExchange,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SenchaExchange) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SenchaExchange) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MinimumOutAmount == nil {
			return errors.New("MinimumOutAmount parameter is not set")
		}
		if inst.PlatformFeeBps == nil {
			return errors.New("PlatformFeeBps parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SwapProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Swap is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.InputUserAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.InputTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.InputFeesAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.OutputUserAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.OutputTokenAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.OutputFeesAccount is not set")
		}
	}
	return nil
}

func (inst *SenchaExchange) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SenchaExchange")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("        InAmount (OPT)", inst.InAmount))
						paramsBranch.Child(ag_format.Param("MinimumOutAmount", *inst.MinimumOutAmount))
						paramsBranch.Child(ag_format.Param("  PlatformFeeBps", *inst.PlatformFeeBps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  swapProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         swap", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("userAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    inputUser", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   inputToken", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    inputFees", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("   outputUser", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("  outputToken", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   outputFees", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj SenchaExchange) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `InAmount` param (optional):
	{
		if obj.InAmount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.InAmount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MinimumOutAmount` param:
	err = encoder.Encode(obj.MinimumOutAmount)
	if err != nil {
		return err
	}
	// Serialize `PlatformFeeBps` param:
	err = encoder.Encode(obj.PlatformFeeBps)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SenchaExchange) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `InAmount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.InAmount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MinimumOutAmount`:
	err = decoder.Decode(&obj.MinimumOutAmount)
	if err != nil {
		return err
	}
	// Deserialize `PlatformFeeBps`:
	err = decoder.Decode(&obj.PlatformFeeBps)
	if err != nil {
		return err
	}
	return nil
}

// NewSenchaExchangeInstruction declares a new SenchaExchange instruction with the provided parameters and accounts.
func NewSenchaExchangeInstruction(
	// Parameters:
	inAmount uint64,
	minimumOutAmount uint64,
	platformFeeBps uint8,
	// Accounts:
	swapProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	swap ag_solanago.PublicKey,
	userAuthority ag_solanago.PublicKey,
	inputUserAccount ag_solanago.PublicKey,
	inputTokenAccount ag_solanago.PublicKey,
	inputFeesAccount ag_solanago.PublicKey,
	outputUserAccount ag_solanago.PublicKey,
	outputTokenAccount ag_solanago.PublicKey,
	outputFeesAccount ag_solanago.PublicKey) *SenchaExchange {
	return NewSenchaExchangeInstructionBuilder().
		SetInAmount(inAmount).
		SetMinimumOutAmount(minimumOutAmount).
		SetPlatformFeeBps(platformFeeBps).
		SetSwapProgramAccount(swapProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSwapAccount(swap).
		SetUserAuthorityAccount(userAuthority).
		SetInputUserAccount(inputUserAccount).
		SetInputTokenAccount(inputTokenAccount).
		SetInputFeesAccount(inputFeesAccount).
		SetOutputUserAccount(outputUserAccount).
		SetOutputTokenAccount(outputTokenAccount).
		SetOutputFeesAccount(outputFeesAccount)
}