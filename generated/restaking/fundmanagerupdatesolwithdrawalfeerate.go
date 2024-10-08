// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerUpdateSolWithdrawalFeeRate is the `fund_manager_update_sol_withdrawal_fee_rate` instruction.
type FundManagerUpdateSolWithdrawalFeeRate struct {
	SolWithdrawalFeeRate *uint16

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] receipt_token_mint
	//
	// [2] = [WRITE] fund_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerUpdateSolWithdrawalFeeRateInstructionBuilder creates a new `FundManagerUpdateSolWithdrawalFeeRate` instruction builder.
func NewFundManagerUpdateSolWithdrawalFeeRateInstructionBuilder() *FundManagerUpdateSolWithdrawalFeeRate {
	nd := &FundManagerUpdateSolWithdrawalFeeRate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetSolWithdrawalFeeRate sets the "sol_withdrawal_fee_rate" parameter.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) SetSolWithdrawalFeeRate(sol_withdrawal_fee_rate uint16) *FundManagerUpdateSolWithdrawalFeeRate {
	inst.SolWithdrawalFeeRate = &sol_withdrawal_fee_rate
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerUpdateSolWithdrawalFeeRate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerUpdateSolWithdrawalFeeRate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerUpdateSolWithdrawalFeeRate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerUpdateSolWithdrawalFeeRate) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerUpdateSolWithdrawalFeeRate) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerUpdateSolWithdrawalFeeRate) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerUpdateSolWithdrawalFeeRate) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst FundManagerUpdateSolWithdrawalFeeRate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerUpdateSolWithdrawalFeeRate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerUpdateSolWithdrawalFeeRate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerUpdateSolWithdrawalFeeRate) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SolWithdrawalFeeRate == nil {
			return errors.New("SolWithdrawalFeeRate parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
	}
	return nil
}

func (inst *FundManagerUpdateSolWithdrawalFeeRate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerUpdateSolWithdrawalFeeRate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   SolWithdrawalFeeRate", *inst.SolWithdrawalFeeRate))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj FundManagerUpdateSolWithdrawalFeeRate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SolWithdrawalFeeRate` param:
	err = encoder.Encode(obj.SolWithdrawalFeeRate)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerUpdateSolWithdrawalFeeRate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SolWithdrawalFeeRate`:
	err = decoder.Decode(&obj.SolWithdrawalFeeRate)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerUpdateSolWithdrawalFeeRateInstruction declares a new FundManagerUpdateSolWithdrawalFeeRate instruction with the provided parameters and accounts.
func NewFundManagerUpdateSolWithdrawalFeeRateInstruction(
	// Parameters:
	sol_withdrawal_fee_rate uint16,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey) *FundManagerUpdateSolWithdrawalFeeRate {
	return NewFundManagerUpdateSolWithdrawalFeeRateInstructionBuilder().
		SetSolWithdrawalFeeRate(sol_withdrawal_fee_rate).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount)
}
