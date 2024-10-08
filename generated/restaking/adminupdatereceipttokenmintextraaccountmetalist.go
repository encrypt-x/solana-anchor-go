// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminUpdateReceiptTokenMintExtraAccountMetaList is the `admin_update_receipt_token_mint_extra_account_meta_list` instruction.
type AdminUpdateReceiptTokenMintExtraAccountMetaList struct {

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] system_program
	//
	// [3] = [] receipt_token_program
	//
	// [4] = [] receipt_token_mint
	//
	// [5] = [] receipt_token_mint_authority
	//
	// [6] = [WRITE] extra_account_meta_list
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminUpdateReceiptTokenMintExtraAccountMetaListInstructionBuilder creates a new `AdminUpdateReceiptTokenMintExtraAccountMetaList` instruction builder.
func NewAdminUpdateReceiptTokenMintExtraAccountMetaListInstructionBuilder() *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	nd := &AdminUpdateReceiptTokenMintExtraAccountMetaList{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["fragkamrANLvuZYQPcmPsCATQAabkqNGH6gxqqPG3aP"]).SIGNER()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"])
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetPayerAccount sets the "payer" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetPayerAccount(payer ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetAdminAccount(admin ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenProgramAccount sets the "receipt_token_program" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetReceiptTokenProgramAccount(receiptTokenProgram ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenProgram)
	return inst
}

// GetReceiptTokenProgramAccount gets the "receipt_token_program" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetReceiptTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReceiptTokenMintAuthorityAccount sets the "receipt_token_mint_authority" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetReceiptTokenMintAuthorityAccount(receiptTokenMintAuthority ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(receiptTokenMintAuthority)
	return inst
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) findFindReceiptTokenMintAuthorityAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: receipt_token_mint_authority
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x6d), byte(0x69), byte(0x6e), byte(0x74), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})
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

// FindReceiptTokenMintAuthorityAddressWithBumpSeed calculates ReceiptTokenMintAuthority account address with given seeds and a known bump seed.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) FindReceiptTokenMintAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) MustFindReceiptTokenMintAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenMintAuthorityAddress finds ReceiptTokenMintAuthority account address with given seeds.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) FindReceiptTokenMintAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) MustFindReceiptTokenMintAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenMintAuthorityAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenMintAuthorityAccount gets the "receipt_token_mint_authority" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetReceiptTokenMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetExtraAccountMetaListAccount sets the "extra_account_meta_list" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) SetExtraAccountMetaListAccount(extraAccountMetaList ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(extraAccountMetaList).WRITE()
	return inst
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) findFindExtraAccountMetaListAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: extra-account-metas
	seeds = append(seeds, []byte{byte(0x65), byte(0x78), byte(0x74), byte(0x72), byte(0x61), byte(0x2d), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74), byte(0x2d), byte(0x6d), byte(0x65), byte(0x74), byte(0x61), byte(0x73)})
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

// FindExtraAccountMetaListAddressWithBumpSeed calculates ExtraAccountMetaList account address with given seeds and a known bump seed.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) FindExtraAccountMetaListAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindExtraAccountMetaListAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) MustFindExtraAccountMetaListAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindExtraAccountMetaListAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindExtraAccountMetaListAddress finds ExtraAccountMetaList account address with given seeds.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) FindExtraAccountMetaListAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindExtraAccountMetaListAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) MustFindExtraAccountMetaListAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindExtraAccountMetaListAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetExtraAccountMetaListAccount gets the "extra_account_meta_list" account.
func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) GetExtraAccountMetaListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst AdminUpdateReceiptTokenMintExtraAccountMetaList) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminUpdateReceiptTokenMintExtraAccountMetaList,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminUpdateReceiptTokenMintExtraAccountMetaList) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReceiptTokenMintAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ExtraAccountMetaList is not set")
		}
	}
	return nil
}

func (inst *AdminUpdateReceiptTokenMintExtraAccountMetaList) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminUpdateReceiptTokenMintExtraAccountMetaList")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                       payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                       admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              system_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       receipt_token_program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          receipt_token_mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint_authority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     extra_account_meta_list", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj AdminUpdateReceiptTokenMintExtraAccountMetaList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AdminUpdateReceiptTokenMintExtraAccountMetaList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAdminUpdateReceiptTokenMintExtraAccountMetaListInstruction declares a new AdminUpdateReceiptTokenMintExtraAccountMetaList instruction with the provided parameters and accounts.
func NewAdminUpdateReceiptTokenMintExtraAccountMetaListInstruction(
	// Accounts:
	payer ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	receiptTokenMintAuthority ag_solanago.PublicKey,
	extraAccountMetaList ag_solanago.PublicKey) *AdminUpdateReceiptTokenMintExtraAccountMetaList {
	return NewAdminUpdateReceiptTokenMintExtraAccountMetaListInstructionBuilder().
		SetPayerAccount(payer).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenProgramAccount(receiptTokenProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetReceiptTokenMintAuthorityAccount(receiptTokenMintAuthority).
		SetExtraAccountMetaListAccount(extraAccountMetaList)
}
