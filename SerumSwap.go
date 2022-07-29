// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SerumSwap is the `serumSwap` instruction.
type SerumSwap struct {
	Side             *Side
	InAmount         *uint64 `bin:"optional"`
	MinimumOutAmount *uint64
	PlatformFeeBps   *uint8

	// ····· market: [0] = [WRITE] market
	//
	// ············· [1] = [WRITE] openOrders
	//
	// ············· [2] = [WRITE] requestQueue
	//
	// ············· [3] = [WRITE] eventQueue
	//
	// ············· [4] = [WRITE] bids
	//
	// ············· [5] = [WRITE] asks
	//
	// ············· [6] = [WRITE] coinVault
	//
	// ············· [7] = [WRITE] pcVault
	//
	// ············· [8] = [] vaultSigner
	//
	// [9] = [SIGNER] authority
	//
	// [10] = [WRITE] orderPayerTokenAccount
	//
	// [11] = [WRITE] coinWallet
	//
	// [12] = [WRITE] pcWallet
	//
	// [13] = [] dexProgram
	//
	// [14] = [] tokenProgram
	//
	// [15] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSerumSwapInstructionBuilder creates a new `SerumSwap` instruction builder.
func NewSerumSwapInstructionBuilder() *SerumSwap {
	nd := &SerumSwap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetSide sets the "side" parameter.
func (inst *SerumSwap) SetSide(side Side) *SerumSwap {
	inst.Side = &side
	return inst
}

// SetInAmount sets the "inAmount" parameter.
func (inst *SerumSwap) SetInAmount(inAmount uint64) *SerumSwap {
	inst.InAmount = &inAmount
	return inst
}

// SetMinimumOutAmount sets the "minimumOutAmount" parameter.
func (inst *SerumSwap) SetMinimumOutAmount(minimumOutAmount uint64) *SerumSwap {
	inst.MinimumOutAmount = &minimumOutAmount
	return inst
}

// SetPlatformFeeBps sets the "platformFeeBps" parameter.
func (inst *SerumSwap) SetPlatformFeeBps(platformFeeBps uint8) *SerumSwap {
	inst.PlatformFeeBps = &platformFeeBps
	return inst
}

type SerumSwapMarketAccountsBuilder struct {
	ag_solanago.AccountMetaSlice `bin:"-"`
}

func NewSerumSwapMarketAccountsBuilder() *SerumSwapMarketAccountsBuilder {
	return &SerumSwapMarketAccountsBuilder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
}

func (inst *SerumSwap) SetMarketAccountsFromBuilder(serumSwapMarketAccountsBuilder *SerumSwapMarketAccountsBuilder) *SerumSwap {
	inst.AccountMetaSlice[0] = serumSwapMarketAccountsBuilder.GetMarketAccount()
	inst.AccountMetaSlice[1] = serumSwapMarketAccountsBuilder.GetOpenOrdersAccount()
	inst.AccountMetaSlice[2] = serumSwapMarketAccountsBuilder.GetRequestQueueAccount()
	inst.AccountMetaSlice[3] = serumSwapMarketAccountsBuilder.GetEventQueueAccount()
	inst.AccountMetaSlice[4] = serumSwapMarketAccountsBuilder.GetBidsAccount()
	inst.AccountMetaSlice[5] = serumSwapMarketAccountsBuilder.GetAsksAccount()
	inst.AccountMetaSlice[6] = serumSwapMarketAccountsBuilder.GetCoinVaultAccount()
	inst.AccountMetaSlice[7] = serumSwapMarketAccountsBuilder.GetPcVaultAccount()
	inst.AccountMetaSlice[8] = serumSwapMarketAccountsBuilder.GetVaultSignerAccount()
	return inst
}

// SetMarketAccount sets the "market" account.
func (inst *SerumSwapMarketAccountsBuilder) SetMarketAccount(market ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *SerumSwapMarketAccountsBuilder) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccount sets the "openOrders" account.
func (inst *SerumSwapMarketAccountsBuilder) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "openOrders" account.
func (inst *SerumSwapMarketAccountsBuilder) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRequestQueueAccount sets the "requestQueue" account.
func (inst *SerumSwapMarketAccountsBuilder) SetRequestQueueAccount(requestQueue ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(requestQueue).WRITE()
	return inst
}

// GetRequestQueueAccount gets the "requestQueue" account.
func (inst *SerumSwapMarketAccountsBuilder) GetRequestQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEventQueueAccount sets the "eventQueue" account.
func (inst *SerumSwapMarketAccountsBuilder) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
func (inst *SerumSwapMarketAccountsBuilder) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBidsAccount sets the "bids" account.
func (inst *SerumSwapMarketAccountsBuilder) SetBidsAccount(bids ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *SerumSwapMarketAccountsBuilder) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAsksAccount sets the "asks" account.
func (inst *SerumSwapMarketAccountsBuilder) SetAsksAccount(asks ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *SerumSwapMarketAccountsBuilder) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCoinVaultAccount sets the "coinVault" account.
func (inst *SerumSwapMarketAccountsBuilder) SetCoinVaultAccount(coinVault ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(coinVault).WRITE()
	return inst
}

// GetCoinVaultAccount gets the "coinVault" account.
func (inst *SerumSwapMarketAccountsBuilder) GetCoinVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPcVaultAccount sets the "pcVault" account.
func (inst *SerumSwapMarketAccountsBuilder) SetPcVaultAccount(pcVault ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(pcVault).WRITE()
	return inst
}

// GetPcVaultAccount gets the "pcVault" account.
func (inst *SerumSwapMarketAccountsBuilder) GetPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetVaultSignerAccount sets the "vaultSigner" account.
func (inst *SerumSwapMarketAccountsBuilder) SetVaultSignerAccount(vaultSigner ag_solanago.PublicKey) *SerumSwapMarketAccountsBuilder {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(vaultSigner)
	return inst
}

// GetVaultSignerAccount gets the "vaultSigner" account.
func (inst *SerumSwapMarketAccountsBuilder) GetVaultSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SerumSwap) SetAuthorityAccount(authority ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SerumSwap) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetOrderPayerTokenAccount sets the "orderPayerTokenAccount" account.
func (inst *SerumSwap) SetOrderPayerTokenAccount(orderPayerTokenAccount ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(orderPayerTokenAccount).WRITE()
	return inst
}

// GetOrderPayerTokenAccount gets the "orderPayerTokenAccount" account.
func (inst *SerumSwap) GetOrderPayerTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetCoinWalletAccount sets the "coinWallet" account.
func (inst *SerumSwap) SetCoinWalletAccount(coinWallet ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(coinWallet).WRITE()
	return inst
}

// GetCoinWalletAccount gets the "coinWallet" account.
func (inst *SerumSwap) GetCoinWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetPcWalletAccount sets the "pcWallet" account.
func (inst *SerumSwap) SetPcWalletAccount(pcWallet ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(pcWallet).WRITE()
	return inst
}

// GetPcWalletAccount gets the "pcWallet" account.
func (inst *SerumSwap) GetPcWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetDexProgramAccount sets the "dexProgram" account.
func (inst *SerumSwap) SetDexProgramAccount(dexProgram ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(dexProgram)
	return inst
}

// GetDexProgramAccount gets the "dexProgram" account.
func (inst *SerumSwap) GetDexProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *SerumSwap) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *SerumSwap) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetRentAccount sets the "rent" account.
func (inst *SerumSwap) SetRentAccount(rent ag_solanago.PublicKey) *SerumSwap {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *SerumSwap) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst SerumSwap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SerumSwap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SerumSwap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SerumSwap) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Side == nil {
			return errors.New("Side parameter is not set")
		}
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
			return errors.New("accounts.MarketMarket is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarketOpenOrders is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MarketRequestQueue is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MarketEventQueue is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MarketBids is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MarketAsks is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MarketCoinVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MarketPcVault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MarketVaultSigner is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.OrderPayerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.CoinWallet is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.PcWallet is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.DexProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *SerumSwap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SerumSwap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("            Side", *inst.Side))
						paramsBranch.Child(ag_format.Param("        InAmount (OPT)", inst.InAmount))
						paramsBranch.Child(ag_format.Param("MinimumOutAmount", *inst.MinimumOutAmount))
						paramsBranch.Child(ag_format.Param("  PlatformFeeBps", *inst.PlatformFeeBps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      market/market", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  market/openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("market/requestQueue", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  market/eventQueue", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        market/bids", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        market/asks", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   market/coinVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     market/pcVault", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" market/vaultSigner", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          authority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("    orderPayerToken", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         coinWallet", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("           pcWallet", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("         dexProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("       tokenProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("               rent", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj SerumSwap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.Encode(obj.Side)
	if err != nil {
		return err
	}
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
func (obj *SerumSwap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	err = decoder.Decode(&obj.Side)
	if err != nil {
		return err
	}
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

// NewSerumSwapInstruction declares a new SerumSwap instruction with the provided parameters and accounts.
func NewSerumSwapInstruction(
	// Parameters:
	side Side,
	inAmount uint64,
	minimumOutAmount uint64,
	platformFeeBps uint8,
	// Accounts:
	marketMarket ag_solanago.PublicKey,
	marketOpenOrders ag_solanago.PublicKey,
	marketRequestQueue ag_solanago.PublicKey,
	marketEventQueue ag_solanago.PublicKey,
	marketBids ag_solanago.PublicKey,
	marketAsks ag_solanago.PublicKey,
	marketCoinVault ag_solanago.PublicKey,
	marketPcVault ag_solanago.PublicKey,
	marketVaultSigner ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	orderPayerTokenAccount ag_solanago.PublicKey,
	coinWallet ag_solanago.PublicKey,
	pcWallet ag_solanago.PublicKey,
	dexProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *SerumSwap {
	return NewSerumSwapInstructionBuilder().
		SetSide(side).
		SetInAmount(inAmount).
		SetMinimumOutAmount(minimumOutAmount).
		SetPlatformFeeBps(platformFeeBps).
		SetMarketAccountsFromBuilder(
			NewSerumSwapMarketAccountsBuilder().
				SetMarketAccount(marketMarket).
				SetOpenOrdersAccount(marketOpenOrders).
				SetRequestQueueAccount(marketRequestQueue).
				SetEventQueueAccount(marketEventQueue).
				SetBidsAccount(marketBids).
				SetAsksAccount(marketAsks).
				SetCoinVaultAccount(marketCoinVault).
				SetPcVaultAccount(marketPcVault).
				SetVaultSignerAccount(marketVaultSigner),
		).
		SetAuthorityAccount(authority).
		SetOrderPayerTokenAccount(orderPayerTokenAccount).
		SetCoinWalletAccount(coinWallet).
		SetPcWalletAccount(pcWallet).
		SetDexProgramAccount(dexProgram).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent)
}