// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OperatorRunIfNeeded is the `operator_run_if_needed` instruction.
type OperatorRunIfNeeded struct {

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [WRITE] fund
	//
	// [2] = [WRITE] fund_token_authority
	//
	// [3] = [WRITE] receipt_token_mint
	//
	// [4] = [WRITE] receipt_token_lock_account
	//
	// [5] = [] token_pricing_source_0
	//
	// [6] = [] token_pricing_source_1
	//
	// [7] = [] token_program
	//
	// [8] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOperatorRunIfNeededInstructionBuilder creates a new `OperatorRunIfNeeded` instruction builder.
func NewOperatorRunIfNeededInstructionBuilder() *OperatorRunIfNeeded {
	nd := &OperatorRunIfNeeded{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	nd.AccountMetaSlice[3] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("FRAGsJAbW4cHk2DYhtAWohV6MUMauJHCFtT1vGvRwnXN")).WRITE()
	nd.AccountMetaSlice[5] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("azFVdHtAJN8BX3sbGAYkXvtdjdrT5U6rj9rovvUFos9"))
	nd.AccountMetaSlice[6] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("8szGkuLTAux9XMgZ2vtY39jVSowEcpBfFfD8hXSEqdGC"))
	nd.AccountMetaSlice[7] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"))
	nd.AccountMetaSlice[8] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"))
	return nd
}

// SetPayerAccount sets the "payer" account.
func (inst *OperatorRunIfNeeded) SetPayerAccount(payer ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *OperatorRunIfNeeded) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFundAccount sets the "fund" account.
func (inst *OperatorRunIfNeeded) SetFundAccount(fund ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fund).WRITE()
	return inst
}

func (inst *OperatorRunIfNeeded) findFindFundAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund_seed
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64), byte(0x5f), byte(0x73), byte(0x65), byte(0x65), byte(0x64)})
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

// FindFundAddressWithBumpSeed calculates Fund account address with given seeds and a known bump seed.
func (inst *OperatorRunIfNeeded) FindFundAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorRunIfNeeded) MustFindFundAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAddress finds Fund account address with given seeds.
func (inst *OperatorRunIfNeeded) FindFundAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorRunIfNeeded) MustFindFundAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccount gets the "fund" account.
func (inst *OperatorRunIfNeeded) GetFundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundTokenAuthorityAccount sets the "fund_token_authority" account.
func (inst *OperatorRunIfNeeded) SetFundTokenAuthorityAccount(fundTokenAuthority ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundTokenAuthority).WRITE()
	return inst
}

func (inst *OperatorRunIfNeeded) findFindFundTokenAuthorityAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund_token_authority_seed
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79), byte(0x5f), byte(0x73), byte(0x65), byte(0x65), byte(0x64)})
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

// FindFundTokenAuthorityAddressWithBumpSeed calculates FundTokenAuthority account address with given seeds and a known bump seed.
func (inst *OperatorRunIfNeeded) FindFundTokenAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundTokenAuthorityAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorRunIfNeeded) MustFindFundTokenAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundTokenAuthorityAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundTokenAuthorityAddress finds FundTokenAuthority account address with given seeds.
func (inst *OperatorRunIfNeeded) FindFundTokenAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundTokenAuthorityAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorRunIfNeeded) MustFindFundTokenAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundTokenAuthorityAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundTokenAuthorityAccount gets the "fund_token_authority" account.
func (inst *OperatorRunIfNeeded) GetFundTokenAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *OperatorRunIfNeeded) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *OperatorRunIfNeeded) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenLockAccountAccount sets the "receipt_token_lock_account" account.
func (inst *OperatorRunIfNeeded) SetReceiptTokenLockAccountAccount(receiptTokenLockAccount ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenLockAccount).WRITE()
	return inst
}

func (inst *OperatorRunIfNeeded) findFindReceiptTokenLockAccountAddress(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: fundTokenAuthority
	seeds = append(seeds, fundTokenAuthority.Bytes())
	// path: tokenProgram
	seeds = append(seeds, tokenProgram.Bytes())
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	programID := ag_solanago.PublicKey([]byte{byte(0x8c), byte(0x97), byte(0x25), byte(0x8f), byte(0x4e), byte(0x24), byte(0x89), byte(0xf1), byte(0xbb), byte(0x3d), byte(0x10), byte(0x29), byte(0x14), byte(0x8e), byte(0xd), byte(0x83), byte(0xb), byte(0x5a), byte(0x13), byte(0x99), byte(0xda), byte(0xff), byte(0x10), byte(0x84), byte(0x4), byte(0x8e), byte(0x7b), byte(0xd8), byte(0xdb), byte(0xe9), byte(0xf8), byte(0x59)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindReceiptTokenLockAccountAddressWithBumpSeed calculates ReceiptTokenLockAccount account address with given seeds and a known bump seed.
func (inst *OperatorRunIfNeeded) FindReceiptTokenLockAccountAddressWithBumpSeed(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorRunIfNeeded) MustFindReceiptTokenLockAccountAddressWithBumpSeed(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenLockAccountAddress finds ReceiptTokenLockAccount account address with given seeds.
func (inst *OperatorRunIfNeeded) FindReceiptTokenLockAccountAddress(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, 0)
	return
}

func (inst *OperatorRunIfNeeded) MustFindReceiptTokenLockAccountAddress(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenLockAccountAccount gets the "receipt_token_lock_account" account.
func (inst *OperatorRunIfNeeded) GetReceiptTokenLockAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenPricingSource0Account sets the "token_pricing_source_0" account.
func (inst *OperatorRunIfNeeded) SetTokenPricingSource0Account(tokenPricingSource0 ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenPricingSource0)
	return inst
}

// GetTokenPricingSource0Account gets the "token_pricing_source_0" account.
func (inst *OperatorRunIfNeeded) GetTokenPricingSource0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenPricingSource1Account sets the "token_pricing_source_1" account.
func (inst *OperatorRunIfNeeded) SetTokenPricingSource1Account(tokenPricingSource1 ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenPricingSource1)
	return inst
}

// GetTokenPricingSource1Account gets the "token_pricing_source_1" account.
func (inst *OperatorRunIfNeeded) GetTokenPricingSource1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *OperatorRunIfNeeded) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *OperatorRunIfNeeded) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *OperatorRunIfNeeded) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *OperatorRunIfNeeded {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *OperatorRunIfNeeded) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst OperatorRunIfNeeded) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OperatorRunIfNeeded,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OperatorRunIfNeeded) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OperatorRunIfNeeded) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Fund is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundTokenAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenLockAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenPricingSource0 is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenPricingSource1 is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *OperatorRunIfNeeded) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OperatorRunIfNeeded")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                    fund", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    fund_token_authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     receipt_token_lock_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  token_pricing_source_0", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  token_pricing_source_1", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj OperatorRunIfNeeded) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *OperatorRunIfNeeded) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewOperatorRunIfNeededInstruction declares a new OperatorRunIfNeeded instruction with the provided parameters and accounts.
func NewOperatorRunIfNeededInstruction(
	// Accounts:
	payer ag_solanago.PublicKey,
	fund ag_solanago.PublicKey,
	fundTokenAuthority ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	receiptTokenLockAccount ag_solanago.PublicKey,
	tokenPricingSource0 ag_solanago.PublicKey,
	tokenPricingSource1 ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *OperatorRunIfNeeded {
	return NewOperatorRunIfNeededInstructionBuilder().
		SetPayerAccount(payer).
		SetFundAccount(fund).
		SetFundTokenAuthorityAccount(fundTokenAuthority).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetReceiptTokenLockAccountAccount(receiptTokenLockAccount).
		SetTokenPricingSource0Account(tokenPricingSource0).
		SetTokenPricingSource1Account(tokenPricingSource1).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}