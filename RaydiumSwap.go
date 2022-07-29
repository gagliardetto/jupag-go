// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RaydiumSwap is the `raydiumSwap` instruction.
type RaydiumSwap struct {
	InAmount         *uint64 `bin:"optional"`
	MinimumOutAmount *uint64
	PlatformFeeBps   *uint8

	// [0] = [] swapProgram
	//
	// [1] = [] tokenProgram
	//
	// [2] = [WRITE] ammId
	//
	// [3] = [] ammAuthority
	//
	// [4] = [WRITE] ammOpenOrders
	//
	// [5] = [WRITE] ammTargetOrders
	//
	// [6] = [WRITE] poolCoinTokenAccount
	//
	// [7] = [WRITE] poolPcTokenAccount
	//
	// [8] = [] serumProgramId
	//
	// [9] = [WRITE] serumMarket
	//
	// [10] = [WRITE] serumBids
	//
	// [11] = [WRITE] serumAsks
	//
	// [12] = [WRITE] serumEventQueue
	//
	// [13] = [WRITE] serumCoinVaultAccount
	//
	// [14] = [WRITE] serumPcVaultAccount
	//
	// [15] = [] serumVaultSigner
	//
	// [16] = [WRITE] userSourceTokenAccount
	//
	// [17] = [WRITE] userDestinationTokenAccount
	//
	// [18] = [SIGNER] userSourceOwner
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRaydiumSwapInstructionBuilder creates a new `RaydiumSwap` instruction builder.
func NewRaydiumSwapInstructionBuilder() *RaydiumSwap {
	nd := &RaydiumSwap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 19),
	}
	return nd
}

// SetInAmount sets the "inAmount" parameter.
func (inst *RaydiumSwap) SetInAmount(inAmount uint64) *RaydiumSwap {
	inst.InAmount = &inAmount
	return inst
}

// SetMinimumOutAmount sets the "minimumOutAmount" parameter.
func (inst *RaydiumSwap) SetMinimumOutAmount(minimumOutAmount uint64) *RaydiumSwap {
	inst.MinimumOutAmount = &minimumOutAmount
	return inst
}

// SetPlatformFeeBps sets the "platformFeeBps" parameter.
func (inst *RaydiumSwap) SetPlatformFeeBps(platformFeeBps uint8) *RaydiumSwap {
	inst.PlatformFeeBps = &platformFeeBps
	return inst
}

// SetSwapProgramAccount sets the "swapProgram" account.
func (inst *RaydiumSwap) SetSwapProgramAccount(swapProgram ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swapProgram)
	return inst
}

// GetSwapProgramAccount gets the "swapProgram" account.
func (inst *RaydiumSwap) GetSwapProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *RaydiumSwap) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *RaydiumSwap) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAmmIdAccount sets the "ammId" account.
func (inst *RaydiumSwap) SetAmmIdAccount(ammId ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ammId).WRITE()
	return inst
}

// GetAmmIdAccount gets the "ammId" account.
func (inst *RaydiumSwap) GetAmmIdAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAmmAuthorityAccount sets the "ammAuthority" account.
func (inst *RaydiumSwap) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "ammAuthority" account.
func (inst *RaydiumSwap) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAmmOpenOrdersAccount sets the "ammOpenOrders" account.
func (inst *RaydiumSwap) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ammOpenOrders).WRITE()
	return inst
}

// GetAmmOpenOrdersAccount gets the "ammOpenOrders" account.
func (inst *RaydiumSwap) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAmmTargetOrdersAccount sets the "ammTargetOrders" account.
func (inst *RaydiumSwap) SetAmmTargetOrdersAccount(ammTargetOrders ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(ammTargetOrders).WRITE()
	return inst
}

// GetAmmTargetOrdersAccount gets the "ammTargetOrders" account.
func (inst *RaydiumSwap) GetAmmTargetOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolCoinTokenAccount sets the "poolCoinTokenAccount" account.
func (inst *RaydiumSwap) SetPoolCoinTokenAccount(poolCoinTokenAccount ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolCoinTokenAccount).WRITE()
	return inst
}

// GetPoolCoinTokenAccount gets the "poolCoinTokenAccount" account.
func (inst *RaydiumSwap) GetPoolCoinTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPoolPcTokenAccount sets the "poolPcTokenAccount" account.
func (inst *RaydiumSwap) SetPoolPcTokenAccount(poolPcTokenAccount ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(poolPcTokenAccount).WRITE()
	return inst
}

// GetPoolPcTokenAccount gets the "poolPcTokenAccount" account.
func (inst *RaydiumSwap) GetPoolPcTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSerumProgramIdAccount sets the "serumProgramId" account.
func (inst *RaydiumSwap) SetSerumProgramIdAccount(serumProgramId ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(serumProgramId)
	return inst
}

// GetSerumProgramIdAccount gets the "serumProgramId" account.
func (inst *RaydiumSwap) GetSerumProgramIdAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSerumMarketAccount sets the "serumMarket" account.
func (inst *RaydiumSwap) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(serumMarket).WRITE()
	return inst
}

// GetSerumMarketAccount gets the "serumMarket" account.
func (inst *RaydiumSwap) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSerumBidsAccount sets the "serumBids" account.
func (inst *RaydiumSwap) SetSerumBidsAccount(serumBids ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(serumBids).WRITE()
	return inst
}

// GetSerumBidsAccount gets the "serumBids" account.
func (inst *RaydiumSwap) GetSerumBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSerumAsksAccount sets the "serumAsks" account.
func (inst *RaydiumSwap) SetSerumAsksAccount(serumAsks ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(serumAsks).WRITE()
	return inst
}

// GetSerumAsksAccount gets the "serumAsks" account.
func (inst *RaydiumSwap) GetSerumAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSerumEventQueueAccount sets the "serumEventQueue" account.
func (inst *RaydiumSwap) SetSerumEventQueueAccount(serumEventQueue ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(serumEventQueue).WRITE()
	return inst
}

// GetSerumEventQueueAccount gets the "serumEventQueue" account.
func (inst *RaydiumSwap) GetSerumEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSerumCoinVaultAccount sets the "serumCoinVaultAccount" account.
func (inst *RaydiumSwap) SetSerumCoinVaultAccount(serumCoinVaultAccount ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(serumCoinVaultAccount).WRITE()
	return inst
}

// GetSerumCoinVaultAccount gets the "serumCoinVaultAccount" account.
func (inst *RaydiumSwap) GetSerumCoinVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSerumPcVaultAccount sets the "serumPcVaultAccount" account.
func (inst *RaydiumSwap) SetSerumPcVaultAccount(serumPcVaultAccount ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(serumPcVaultAccount).WRITE()
	return inst
}

// GetSerumPcVaultAccount gets the "serumPcVaultAccount" account.
func (inst *RaydiumSwap) GetSerumPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSerumVaultSignerAccount sets the "serumVaultSigner" account.
func (inst *RaydiumSwap) SetSerumVaultSignerAccount(serumVaultSigner ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(serumVaultSigner)
	return inst
}

// GetSerumVaultSignerAccount gets the "serumVaultSigner" account.
func (inst *RaydiumSwap) GetSerumVaultSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetUserSourceTokenAccount sets the "userSourceTokenAccount" account.
func (inst *RaydiumSwap) SetUserSourceTokenAccount(userSourceTokenAccount ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(userSourceTokenAccount).WRITE()
	return inst
}

// GetUserSourceTokenAccount gets the "userSourceTokenAccount" account.
func (inst *RaydiumSwap) GetUserSourceTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetUserDestinationTokenAccount sets the "userDestinationTokenAccount" account.
func (inst *RaydiumSwap) SetUserDestinationTokenAccount(userDestinationTokenAccount ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(userDestinationTokenAccount).WRITE()
	return inst
}

// GetUserDestinationTokenAccount gets the "userDestinationTokenAccount" account.
func (inst *RaydiumSwap) GetUserDestinationTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetUserSourceOwnerAccount sets the "userSourceOwner" account.
func (inst *RaydiumSwap) SetUserSourceOwnerAccount(userSourceOwner ag_solanago.PublicKey) *RaydiumSwap {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(userSourceOwner).SIGNER()
	return inst
}

// GetUserSourceOwnerAccount gets the "userSourceOwner" account.
func (inst *RaydiumSwap) GetUserSourceOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

func (inst RaydiumSwap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RaydiumSwap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RaydiumSwap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RaydiumSwap) Validate() error {
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
			return errors.New("accounts.AmmId is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AmmAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AmmTargetOrders is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PoolPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SerumProgramId is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SerumMarket is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SerumBids is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SerumAsks is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SerumEventQueue is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SerumCoinVaultAccount is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SerumPcVaultAccount is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SerumVaultSigner is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.UserSourceTokenAccount is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.UserDestinationTokenAccount is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.UserSourceOwner is not set")
		}
	}
	return nil
}

func (inst *RaydiumSwap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RaydiumSwap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("        InAmount (OPT)", inst.InAmount))
						paramsBranch.Child(ag_format.Param("MinimumOutAmount", *inst.MinimumOutAmount))
						paramsBranch.Child(ag_format.Param("  PlatformFeeBps", *inst.PlatformFeeBps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=19]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         swapProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               ammId", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        ammAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       ammOpenOrders", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     ammTargetOrders", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       poolCoinToken", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         poolPcToken", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      serumProgramId", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("         serumMarket", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           serumBids", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("           serumAsks", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("     serumEventQueue", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("      serumCoinVault", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("        serumPcVault", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("    serumVaultSigner", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("     userSourceToken", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("userDestinationToken", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("     userSourceOwner", inst.AccountMetaSlice.Get(18)))
					})
				})
		})
}

func (obj RaydiumSwap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *RaydiumSwap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewRaydiumSwapInstruction declares a new RaydiumSwap instruction with the provided parameters and accounts.
func NewRaydiumSwapInstruction(
	// Parameters:
	inAmount uint64,
	minimumOutAmount uint64,
	platformFeeBps uint8,
	// Accounts:
	swapProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	ammId ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	serumProgramId ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumBids ag_solanago.PublicKey,
	serumAsks ag_solanago.PublicKey,
	serumEventQueue ag_solanago.PublicKey,
	serumCoinVaultAccount ag_solanago.PublicKey,
	serumPcVaultAccount ag_solanago.PublicKey,
	serumVaultSigner ag_solanago.PublicKey,
	userSourceTokenAccount ag_solanago.PublicKey,
	userDestinationTokenAccount ag_solanago.PublicKey,
	userSourceOwner ag_solanago.PublicKey) *RaydiumSwap {
	return NewRaydiumSwapInstructionBuilder().
		SetInAmount(inAmount).
		SetMinimumOutAmount(minimumOutAmount).
		SetPlatformFeeBps(platformFeeBps).
		SetSwapProgramAccount(swapProgram).
		SetTokenProgramAccount(tokenProgram).
		SetAmmIdAccount(ammId).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetPoolCoinTokenAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccount(poolPcTokenAccount).
		SetSerumProgramIdAccount(serumProgramId).
		SetSerumMarketAccount(serumMarket).
		SetSerumBidsAccount(serumBids).
		SetSerumAsksAccount(serumAsks).
		SetSerumEventQueueAccount(serumEventQueue).
		SetSerumCoinVaultAccount(serumCoinVaultAccount).
		SetSerumPcVaultAccount(serumPcVaultAccount).
		SetSerumVaultSignerAccount(serumVaultSigner).
		SetUserSourceTokenAccount(userSourceTokenAccount).
		SetUserDestinationTokenAccount(userDestinationTokenAccount).
		SetUserSourceOwnerAccount(userSourceOwner)
}
