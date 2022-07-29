// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AldrinV2Swap is the `aldrinV2Swap` instruction.
type AldrinV2Swap struct {
	InAmount         *uint64 `bin:"optional"`
	MinimumOutAmount *uint64
	Side             *Side
	PlatformFeeBps   *uint8

	// [0] = [] swapProgram
	//
	// [1] = [] pool
	//
	// [2] = [] poolSigner
	//
	// [3] = [WRITE] poolMint
	//
	// [4] = [WRITE] baseTokenVault
	//
	// [5] = [WRITE] quoteTokenVault
	//
	// [6] = [WRITE] feePoolTokenAccount
	//
	// [7] = [SIGNER] walletAuthority
	//
	// [8] = [WRITE] userBaseTokenAccount
	//
	// [9] = [WRITE] userQuoteTokenAccount
	//
	// [10] = [] curve
	//
	// [11] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAldrinV2SwapInstructionBuilder creates a new `AldrinV2Swap` instruction builder.
func NewAldrinV2SwapInstructionBuilder() *AldrinV2Swap {
	nd := &AldrinV2Swap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetInAmount sets the "inAmount" parameter.
func (inst *AldrinV2Swap) SetInAmount(inAmount uint64) *AldrinV2Swap {
	inst.InAmount = &inAmount
	return inst
}

// SetMinimumOutAmount sets the "minimumOutAmount" parameter.
func (inst *AldrinV2Swap) SetMinimumOutAmount(minimumOutAmount uint64) *AldrinV2Swap {
	inst.MinimumOutAmount = &minimumOutAmount
	return inst
}

// SetSide sets the "side" parameter.
func (inst *AldrinV2Swap) SetSide(side Side) *AldrinV2Swap {
	inst.Side = &side
	return inst
}

// SetPlatformFeeBps sets the "platformFeeBps" parameter.
func (inst *AldrinV2Swap) SetPlatformFeeBps(platformFeeBps uint8) *AldrinV2Swap {
	inst.PlatformFeeBps = &platformFeeBps
	return inst
}

// SetSwapProgramAccount sets the "swapProgram" account.
func (inst *AldrinV2Swap) SetSwapProgramAccount(swapProgram ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swapProgram)
	return inst
}

// GetSwapProgramAccount gets the "swapProgram" account.
func (inst *AldrinV2Swap) GetSwapProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPoolAccount sets the "pool" account.
func (inst *AldrinV2Swap) SetPoolAccount(pool ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pool)
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *AldrinV2Swap) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolSignerAccount sets the "poolSigner" account.
func (inst *AldrinV2Swap) SetPoolSignerAccount(poolSigner ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(poolSigner)
	return inst
}

// GetPoolSignerAccount gets the "poolSigner" account.
func (inst *AldrinV2Swap) GetPoolSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolMintAccount sets the "poolMint" account.
func (inst *AldrinV2Swap) SetPoolMintAccount(poolMint ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

// GetPoolMintAccount gets the "poolMint" account.
func (inst *AldrinV2Swap) GetPoolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBaseTokenVaultAccount sets the "baseTokenVault" account.
func (inst *AldrinV2Swap) SetBaseTokenVaultAccount(baseTokenVault ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(baseTokenVault).WRITE()
	return inst
}

// GetBaseTokenVaultAccount gets the "baseTokenVault" account.
func (inst *AldrinV2Swap) GetBaseTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetQuoteTokenVaultAccount sets the "quoteTokenVault" account.
func (inst *AldrinV2Swap) SetQuoteTokenVaultAccount(quoteTokenVault ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(quoteTokenVault).WRITE()
	return inst
}

// GetQuoteTokenVaultAccount gets the "quoteTokenVault" account.
func (inst *AldrinV2Swap) GetQuoteTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetFeePoolTokenAccount sets the "feePoolTokenAccount" account.
func (inst *AldrinV2Swap) SetFeePoolTokenAccount(feePoolTokenAccount ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(feePoolTokenAccount).WRITE()
	return inst
}

// GetFeePoolTokenAccount gets the "feePoolTokenAccount" account.
func (inst *AldrinV2Swap) GetFeePoolTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetWalletAuthorityAccount sets the "walletAuthority" account.
func (inst *AldrinV2Swap) SetWalletAuthorityAccount(walletAuthority ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(walletAuthority).SIGNER()
	return inst
}

// GetWalletAuthorityAccount gets the "walletAuthority" account.
func (inst *AldrinV2Swap) GetWalletAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetUserBaseTokenAccount sets the "userBaseTokenAccount" account.
func (inst *AldrinV2Swap) SetUserBaseTokenAccount(userBaseTokenAccount ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(userBaseTokenAccount).WRITE()
	return inst
}

// GetUserBaseTokenAccount gets the "userBaseTokenAccount" account.
func (inst *AldrinV2Swap) GetUserBaseTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetUserQuoteTokenAccount sets the "userQuoteTokenAccount" account.
func (inst *AldrinV2Swap) SetUserQuoteTokenAccount(userQuoteTokenAccount ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(userQuoteTokenAccount).WRITE()
	return inst
}

// GetUserQuoteTokenAccount gets the "userQuoteTokenAccount" account.
func (inst *AldrinV2Swap) GetUserQuoteTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetCurveAccount sets the "curve" account.
func (inst *AldrinV2Swap) SetCurveAccount(curve ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(curve)
	return inst
}

// GetCurveAccount gets the "curve" account.
func (inst *AldrinV2Swap) GetCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *AldrinV2Swap) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AldrinV2Swap {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *AldrinV2Swap) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst AldrinV2Swap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AldrinV2Swap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AldrinV2Swap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AldrinV2Swap) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MinimumOutAmount == nil {
			return errors.New("MinimumOutAmount parameter is not set")
		}
		if inst.Side == nil {
			return errors.New("Side parameter is not set")
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
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PoolSigner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.BaseTokenVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.QuoteTokenVault is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FeePoolTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.WalletAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.UserBaseTokenAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.UserQuoteTokenAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Curve is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *AldrinV2Swap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AldrinV2Swap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("        InAmount (OPT)", inst.InAmount))
						paramsBranch.Child(ag_format.Param("MinimumOutAmount", *inst.MinimumOutAmount))
						paramsBranch.Child(ag_format.Param("            Side", *inst.Side))
						paramsBranch.Child(ag_format.Param("  PlatformFeeBps", *inst.PlatformFeeBps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    swapProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           pool", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     poolSigner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       poolMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" baseTokenVault", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("quoteTokenVault", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   feePoolToken", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("walletAuthority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("  userBaseToken", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta(" userQuoteToken", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("          curve", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj AldrinV2Swap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `Side` param:
	err = encoder.Encode(obj.Side)
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
func (obj *AldrinV2Swap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `Side`:
	err = decoder.Decode(&obj.Side)
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

// NewAldrinV2SwapInstruction declares a new AldrinV2Swap instruction with the provided parameters and accounts.
func NewAldrinV2SwapInstruction(
	// Parameters:
	inAmount uint64,
	minimumOutAmount uint64,
	side Side,
	platformFeeBps uint8,
	// Accounts:
	swapProgram ag_solanago.PublicKey,
	pool ag_solanago.PublicKey,
	poolSigner ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	baseTokenVault ag_solanago.PublicKey,
	quoteTokenVault ag_solanago.PublicKey,
	feePoolTokenAccount ag_solanago.PublicKey,
	walletAuthority ag_solanago.PublicKey,
	userBaseTokenAccount ag_solanago.PublicKey,
	userQuoteTokenAccount ag_solanago.PublicKey,
	curve ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *AldrinV2Swap {
	return NewAldrinV2SwapInstructionBuilder().
		SetInAmount(inAmount).
		SetMinimumOutAmount(minimumOutAmount).
		SetSide(side).
		SetPlatformFeeBps(platformFeeBps).
		SetSwapProgramAccount(swapProgram).
		SetPoolAccount(pool).
		SetPoolSignerAccount(poolSigner).
		SetPoolMintAccount(poolMint).
		SetBaseTokenVaultAccount(baseTokenVault).
		SetQuoteTokenVaultAccount(quoteTokenVault).
		SetFeePoolTokenAccount(feePoolTokenAccount).
		SetWalletAuthorityAccount(walletAuthority).
		SetUserBaseTokenAccount(userBaseTokenAccount).
		SetUserQuoteTokenAccount(userQuoteTokenAccount).
		SetCurveAccount(curve).
		SetTokenProgramAccount(tokenProgram)
}