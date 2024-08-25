// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundCancelWithdrawalRequest is the `fund_cancel_withdrawal_request` instruction.
type FundCancelWithdrawalRequest struct {
	RequestId *uint64

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [WRITE] user_receipt
	//
	// [2] = [WRITE] fund
	//
	// [3] = [WRITE] fund_token_authority
	//
	// [4] = [WRITE] receipt_token_mint
	//
	// [5] = [WRITE] receipt_token_account
	//
	// [6] = [WRITE] receipt_token_lock_account
	//
	// [7] = [] token_program
	//
	// [8] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundCancelWithdrawalRequestInstructionBuilder creates a new `FundCancelWithdrawalRequest` instruction builder.
func NewFundCancelWithdrawalRequestInstructionBuilder() *FundCancelWithdrawalRequest {
	nd := &FundCancelWithdrawalRequest{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("FRAGsJAbW4cHk2DYhtAWohV6MUMauJHCFtT1vGvRwnXN")).WRITE()
	nd.AccountMetaSlice[7] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"))
	nd.AccountMetaSlice[8] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"))
	return nd
}

// SetRequestId sets the "request_id" parameter.
func (inst *FundCancelWithdrawalRequest) SetRequestId(request_id uint64) *FundCancelWithdrawalRequest {
	inst.RequestId = &request_id
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *FundCancelWithdrawalRequest) SetUserAccount(user ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *FundCancelWithdrawalRequest) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserReceiptAccount sets the "user_receipt" account.
func (inst *FundCancelWithdrawalRequest) SetUserReceiptAccount(userReceipt ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(userReceipt).WRITE()
	return inst
}

func (inst *FundCancelWithdrawalRequest) findFindUserReceiptAddress(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_receipt_seed
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x73), byte(0x65), byte(0x65), byte(0x64)})
	// path: user
	seeds = append(seeds, user.Bytes())
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

// FindUserReceiptAddressWithBumpSeed calculates UserReceipt account address with given seeds and a known bump seed.
func (inst *FundCancelWithdrawalRequest) FindUserReceiptAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserReceiptAddress(user, receiptTokenMint, bumpSeed)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindUserReceiptAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptAddress(user, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserReceiptAddress finds UserReceipt account address with given seeds.
func (inst *FundCancelWithdrawalRequest) FindUserReceiptAddress(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserReceiptAddress(user, receiptTokenMint, 0)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindUserReceiptAddress(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptAddress(user, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserReceiptAccount gets the "user_receipt" account.
func (inst *FundCancelWithdrawalRequest) GetUserReceiptAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccount sets the "fund" account.
func (inst *FundCancelWithdrawalRequest) SetFundAccount(fund ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fund).WRITE()
	return inst
}

func (inst *FundCancelWithdrawalRequest) findFindFundAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundCancelWithdrawalRequest) FindFundAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindFundAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAddress finds Fund account address with given seeds.
func (inst *FundCancelWithdrawalRequest) FindFundAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAddress(receiptTokenMint, 0)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindFundAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccount gets the "fund" account.
func (inst *FundCancelWithdrawalRequest) GetFundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFundTokenAuthorityAccount sets the "fund_token_authority" account.
func (inst *FundCancelWithdrawalRequest) SetFundTokenAuthorityAccount(fundTokenAuthority ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(fundTokenAuthority).WRITE()
	return inst
}

func (inst *FundCancelWithdrawalRequest) findFindFundTokenAuthorityAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundCancelWithdrawalRequest) FindFundTokenAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundTokenAuthorityAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindFundTokenAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundTokenAuthorityAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundTokenAuthorityAddress finds FundTokenAuthority account address with given seeds.
func (inst *FundCancelWithdrawalRequest) FindFundTokenAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundTokenAuthorityAddress(receiptTokenMint, 0)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindFundTokenAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundTokenAuthorityAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundTokenAuthorityAccount gets the "fund_token_authority" account.
func (inst *FundCancelWithdrawalRequest) GetFundTokenAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundCancelWithdrawalRequest) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundCancelWithdrawalRequest) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReceiptTokenAccountAccount sets the "receipt_token_account" account.
func (inst *FundCancelWithdrawalRequest) SetReceiptTokenAccountAccount(receiptTokenAccount ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(receiptTokenAccount).WRITE()
	return inst
}

func (inst *FundCancelWithdrawalRequest) findFindReceiptTokenAccountAddress(user ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: user
	seeds = append(seeds, user.Bytes())
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

// FindReceiptTokenAccountAddressWithBumpSeed calculates ReceiptTokenAccount account address with given seeds and a known bump seed.
func (inst *FundCancelWithdrawalRequest) FindReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenAccountAddress(user, tokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenAccountAddress(user, tokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenAccountAddress finds ReceiptTokenAccount account address with given seeds.
func (inst *FundCancelWithdrawalRequest) FindReceiptTokenAccountAddress(user ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenAccountAddress(user, tokenProgram, receiptTokenMint, 0)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindReceiptTokenAccountAddress(user ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenAccountAddress(user, tokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenAccountAccount gets the "receipt_token_account" account.
func (inst *FundCancelWithdrawalRequest) GetReceiptTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReceiptTokenLockAccountAccount sets the "receipt_token_lock_account" account.
func (inst *FundCancelWithdrawalRequest) SetReceiptTokenLockAccountAccount(receiptTokenLockAccount ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(receiptTokenLockAccount).WRITE()
	return inst
}

func (inst *FundCancelWithdrawalRequest) findFindReceiptTokenLockAccountAddress(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundCancelWithdrawalRequest) FindReceiptTokenLockAccountAddressWithBumpSeed(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindReceiptTokenLockAccountAddressWithBumpSeed(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenLockAccountAddress finds ReceiptTokenLockAccount account address with given seeds.
func (inst *FundCancelWithdrawalRequest) FindReceiptTokenLockAccountAddress(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, 0)
	return
}

func (inst *FundCancelWithdrawalRequest) MustFindReceiptTokenLockAccountAddress(fundTokenAuthority ag_solanago.PublicKey, tokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(fundTokenAuthority, tokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenLockAccountAccount gets the "receipt_token_lock_account" account.
func (inst *FundCancelWithdrawalRequest) GetReceiptTokenLockAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *FundCancelWithdrawalRequest) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *FundCancelWithdrawalRequest) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *FundCancelWithdrawalRequest) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *FundCancelWithdrawalRequest) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst FundCancelWithdrawalRequest) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundCancelWithdrawalRequest,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundCancelWithdrawalRequest) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundCancelWithdrawalRequest) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RequestId == nil {
			return errors.New("RequestId parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UserReceipt is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Fund is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FundTokenAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReceiptTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReceiptTokenLockAccount is not set")
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

func (inst *FundCancelWithdrawalRequest) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundCancelWithdrawalRequest")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" RequestId", *inst.RequestId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            user_receipt", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    fund", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    fund_token_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      receipt_token_mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          receipt_token_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     receipt_token_lock_", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj FundCancelWithdrawalRequest) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundCancelWithdrawalRequest) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}

// NewFundCancelWithdrawalRequestInstruction declares a new FundCancelWithdrawalRequest instruction with the provided parameters and accounts.
func NewFundCancelWithdrawalRequestInstruction(
	// Parameters:
	request_id uint64,
	// Accounts:
	user ag_solanago.PublicKey,
	userReceipt ag_solanago.PublicKey,
	fund ag_solanago.PublicKey,
	fundTokenAuthority ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	receiptTokenAccount ag_solanago.PublicKey,
	receiptTokenLockAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *FundCancelWithdrawalRequest {
	return NewFundCancelWithdrawalRequestInstructionBuilder().
		SetRequestId(request_id).
		SetUserAccount(user).
		SetUserReceiptAccount(userReceipt).
		SetFundAccount(fund).
		SetFundTokenAuthorityAccount(fundTokenAuthority).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetReceiptTokenAccountAccount(receiptTokenAccount).
		SetReceiptTokenLockAccountAccount(receiptTokenLockAccount).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}