// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SaberAddDecimalsDeposit is the `saberAddDecimalsDeposit` instruction.
type SaberAddDecimalsDeposit struct {
	InAmount         *uint64 `bin:"optional"`
	MinimumOutAmount *uint64
	PlatformFeeBps   *uint8

	// [0] = [] addDecimalsProgram
	//
	// [1] = [] wrapper
	//
	// [2] = [WRITE] wrapperMint
	//
	// [3] = [WRITE] wrapperUnderlyingTokens
	//
	// [4] = [SIGNER] owner
	//
	// [5] = [WRITE] userUnderlyingTokens
	//
	// [6] = [WRITE] userWrappedTokens
	//
	// [7] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSaberAddDecimalsDepositInstructionBuilder creates a new `SaberAddDecimalsDeposit` instruction builder.
func NewSaberAddDecimalsDepositInstructionBuilder() *SaberAddDecimalsDeposit {
	nd := &SaberAddDecimalsDeposit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetInAmount sets the "inAmount" parameter.
func (inst *SaberAddDecimalsDeposit) SetInAmount(inAmount uint64) *SaberAddDecimalsDeposit {
	inst.InAmount = &inAmount
	return inst
}

// SetMinimumOutAmount sets the "minimumOutAmount" parameter.
func (inst *SaberAddDecimalsDeposit) SetMinimumOutAmount(minimumOutAmount uint64) *SaberAddDecimalsDeposit {
	inst.MinimumOutAmount = &minimumOutAmount
	return inst
}

// SetPlatformFeeBps sets the "platformFeeBps" parameter.
func (inst *SaberAddDecimalsDeposit) SetPlatformFeeBps(platformFeeBps uint8) *SaberAddDecimalsDeposit {
	inst.PlatformFeeBps = &platformFeeBps
	return inst
}

// SetAddDecimalsProgramAccount sets the "addDecimalsProgram" account.
func (inst *SaberAddDecimalsDeposit) SetAddDecimalsProgramAccount(addDecimalsProgram ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(addDecimalsProgram)
	return inst
}

// GetAddDecimalsProgramAccount gets the "addDecimalsProgram" account.
func (inst *SaberAddDecimalsDeposit) GetAddDecimalsProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetWrapperAccount sets the "wrapper" account.
func (inst *SaberAddDecimalsDeposit) SetWrapperAccount(wrapper ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(wrapper)
	return inst
}

// GetWrapperAccount gets the "wrapper" account.
func (inst *SaberAddDecimalsDeposit) GetWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetWrapperMintAccount sets the "wrapperMint" account.
func (inst *SaberAddDecimalsDeposit) SetWrapperMintAccount(wrapperMint ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(wrapperMint).WRITE()
	return inst
}

// GetWrapperMintAccount gets the "wrapperMint" account.
func (inst *SaberAddDecimalsDeposit) GetWrapperMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetWrapperUnderlyingTokensAccount sets the "wrapperUnderlyingTokens" account.
func (inst *SaberAddDecimalsDeposit) SetWrapperUnderlyingTokensAccount(wrapperUnderlyingTokens ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(wrapperUnderlyingTokens).WRITE()
	return inst
}

// GetWrapperUnderlyingTokensAccount gets the "wrapperUnderlyingTokens" account.
func (inst *SaberAddDecimalsDeposit) GetWrapperUnderlyingTokensAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOwnerAccount sets the "owner" account.
func (inst *SaberAddDecimalsDeposit) SetOwnerAccount(owner ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *SaberAddDecimalsDeposit) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUserUnderlyingTokensAccount sets the "userUnderlyingTokens" account.
func (inst *SaberAddDecimalsDeposit) SetUserUnderlyingTokensAccount(userUnderlyingTokens ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(userUnderlyingTokens).WRITE()
	return inst
}

// GetUserUnderlyingTokensAccount gets the "userUnderlyingTokens" account.
func (inst *SaberAddDecimalsDeposit) GetUserUnderlyingTokensAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUserWrappedTokensAccount sets the "userWrappedTokens" account.
func (inst *SaberAddDecimalsDeposit) SetUserWrappedTokensAccount(userWrappedTokens ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(userWrappedTokens).WRITE()
	return inst
}

// GetUserWrappedTokensAccount gets the "userWrappedTokens" account.
func (inst *SaberAddDecimalsDeposit) GetUserWrappedTokensAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *SaberAddDecimalsDeposit) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *SaberAddDecimalsDeposit) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst SaberAddDecimalsDeposit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SaberAddDecimalsDeposit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SaberAddDecimalsDeposit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SaberAddDecimalsDeposit) Validate() error {
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
			return errors.New("accounts.AddDecimalsProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Wrapper is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.WrapperMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.WrapperUnderlyingTokens is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.UserUnderlyingTokens is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.UserWrappedTokens is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *SaberAddDecimalsDeposit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SaberAddDecimalsDeposit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("        InAmount (OPT)", inst.InAmount))
						paramsBranch.Child(ag_format.Param("MinimumOutAmount", *inst.MinimumOutAmount))
						paramsBranch.Child(ag_format.Param("  PlatformFeeBps", *inst.PlatformFeeBps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     addDecimalsProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                wrapper", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            wrapperMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("wrapperUnderlyingTokens", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                  owner", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   userUnderlyingTokens", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      userWrappedTokens", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           tokenProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj SaberAddDecimalsDeposit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *SaberAddDecimalsDeposit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewSaberAddDecimalsDepositInstruction declares a new SaberAddDecimalsDeposit instruction with the provided parameters and accounts.
func NewSaberAddDecimalsDepositInstruction(
	// Parameters:
	inAmount uint64,
	minimumOutAmount uint64,
	platformFeeBps uint8,
	// Accounts:
	addDecimalsProgram ag_solanago.PublicKey,
	wrapper ag_solanago.PublicKey,
	wrapperMint ag_solanago.PublicKey,
	wrapperUnderlyingTokens ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	userUnderlyingTokens ag_solanago.PublicKey,
	userWrappedTokens ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *SaberAddDecimalsDeposit {
	return NewSaberAddDecimalsDepositInstructionBuilder().
		SetInAmount(inAmount).
		SetMinimumOutAmount(minimumOutAmount).
		SetPlatformFeeBps(platformFeeBps).
		SetAddDecimalsProgramAccount(addDecimalsProgram).
		SetWrapperAccount(wrapper).
		SetWrapperMintAccount(wrapperMint).
		SetWrapperUnderlyingTokensAccount(wrapperUnderlyingTokens).
		SetOwnerAccount(owner).
		SetUserUnderlyingTokensAccount(userUnderlyingTokens).
		SetUserWrappedTokensAccount(userWrappedTokens).
		SetTokenProgramAccount(tokenProgram)
}
