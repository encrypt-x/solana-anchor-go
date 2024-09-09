// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UserInitializeRewardAccounts is the `user_initialize_reward_accounts` instruction.
type UserInitializeRewardAccounts struct {

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [] system_program
	//
	// [2] = [] receipt_token_mint
	//
	// [3] = [WRITE] user_reward_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUserInitializeRewardAccountsInstructionBuilder creates a new `UserInitializeRewardAccounts` instruction builder.
func NewUserInitializeRewardAccountsInstructionBuilder() *UserInitializeRewardAccounts {
	nd := &UserInitializeRewardAccounts{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetUserAccount sets the "user" account.
func (inst *UserInitializeRewardAccounts) SetUserAccount(user ag_solanago.PublicKey) *UserInitializeRewardAccounts {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UserInitializeRewardAccounts) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UserInitializeRewardAccounts) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UserInitializeRewardAccounts {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UserInitializeRewardAccounts) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *UserInitializeRewardAccounts) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *UserInitializeRewardAccounts {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *UserInitializeRewardAccounts) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserRewardAccountAccount sets the "user_reward_account" account.
func (inst *UserInitializeRewardAccounts) SetUserRewardAccountAccount(userRewardAccount ag_solanago.PublicKey) *UserInitializeRewardAccounts {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userRewardAccount).WRITE()
	return inst
}

func (inst *UserInitializeRewardAccounts) findFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_reward
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserRewardAccountAddressWithBumpSeed calculates UserRewardAccount account address with given seeds and a known bump seed.
func (inst *UserInitializeRewardAccounts) FindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserInitializeRewardAccounts) MustFindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserRewardAccountAddress finds UserRewardAccount account address with given seeds.
func (inst *UserInitializeRewardAccounts) FindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserInitializeRewardAccounts) MustFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserRewardAccountAccount gets the "user_reward_account" account.
func (inst *UserInitializeRewardAccounts) GetUserRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst UserInitializeRewardAccounts) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UserInitializeRewardAccounts,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UserInitializeRewardAccounts) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UserInitializeRewardAccounts) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserRewardAccount is not set")
		}
	}
	return nil
}

func (inst *UserInitializeRewardAccounts) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UserInitializeRewardAccounts")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      user_reward_", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj UserInitializeRewardAccounts) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UserInitializeRewardAccounts) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUserInitializeRewardAccountsInstruction declares a new UserInitializeRewardAccounts instruction with the provided parameters and accounts.
func NewUserInitializeRewardAccountsInstruction(
	// Accounts:
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	userRewardAccount ag_solanago.PublicKey) *UserInitializeRewardAccounts {
	return NewUserInitializeRewardAccountsInstructionBuilder().
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetUserRewardAccountAccount(userRewardAccount)
}