// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundWithdraw is the `fund_withdraw` instruction.
type FundWithdraw struct {
	RequestId *uint64

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [WRITE] user_receipt
	//
	// [2] = [WRITE] fund
	//
	// [3] = [] receipt_token_mint
	//
	// [4] = [] token_pricing_source_0
	//
	// [5] = [] token_pricing_source_1
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundWithdrawInstructionBuilder creates a new `FundWithdraw` instruction builder.
func NewFundWithdrawInstructionBuilder() *FundWithdraw {
	nd := &FundWithdraw{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[3] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("FRAGsJAbW4cHk2DYhtAWohV6MUMauJHCFtT1vGvRwnXN"))
	nd.AccountMetaSlice[4] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("azFVdHtAJN8BX3sbGAYkXvtdjdrT5U6rj9rovvUFos9"))
	nd.AccountMetaSlice[5] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("8szGkuLTAux9XMgZ2vtY39jVSowEcpBfFfD8hXSEqdGC"))
	return nd
}

// SetRequestId sets the "request_id" parameter.
func (inst *FundWithdraw) SetRequestId(request_id uint64) *FundWithdraw {
	inst.RequestId = &request_id
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *FundWithdraw) SetUserAccount(user ag_solanago.PublicKey) *FundWithdraw {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *FundWithdraw) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserReceiptAccount sets the "user_receipt" account.
func (inst *FundWithdraw) SetUserReceiptAccount(userReceipt ag_solanago.PublicKey) *FundWithdraw {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(userReceipt).WRITE()
	return inst
}

func (inst *FundWithdraw) findFindUserReceiptAddress(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundWithdraw) FindUserReceiptAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserReceiptAddress(user, receiptTokenMint, bumpSeed)
	return
}

func (inst *FundWithdraw) MustFindUserReceiptAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptAddress(user, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserReceiptAddress finds UserReceipt account address with given seeds.
func (inst *FundWithdraw) FindUserReceiptAddress(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserReceiptAddress(user, receiptTokenMint, 0)
	return
}

func (inst *FundWithdraw) MustFindUserReceiptAddress(user ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptAddress(user, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserReceiptAccount gets the "user_receipt" account.
func (inst *FundWithdraw) GetUserReceiptAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccount sets the "fund" account.
func (inst *FundWithdraw) SetFundAccount(fund ag_solanago.PublicKey) *FundWithdraw {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fund).WRITE()
	return inst
}

func (inst *FundWithdraw) findFindFundAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundWithdraw) FindFundAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundWithdraw) MustFindFundAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAddress finds Fund account address with given seeds.
func (inst *FundWithdraw) FindFundAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAddress(receiptTokenMint, 0)
	return
}

func (inst *FundWithdraw) MustFindFundAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccount gets the "fund" account.
func (inst *FundWithdraw) GetFundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundWithdraw) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundWithdraw {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundWithdraw) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenPricingSource0Account sets the "token_pricing_source_0" account.
func (inst *FundWithdraw) SetTokenPricingSource0Account(tokenPricingSource0 ag_solanago.PublicKey) *FundWithdraw {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenPricingSource0)
	return inst
}

// GetTokenPricingSource0Account gets the "token_pricing_source_0" account.
func (inst *FundWithdraw) GetTokenPricingSource0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenPricingSource1Account sets the "token_pricing_source_1" account.
func (inst *FundWithdraw) SetTokenPricingSource1Account(tokenPricingSource1 ag_solanago.PublicKey) *FundWithdraw {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenPricingSource1)
	return inst
}

// GetTokenPricingSource1Account gets the "token_pricing_source_1" account.
func (inst *FundWithdraw) GetTokenPricingSource1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst FundWithdraw) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundWithdraw,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundWithdraw) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundWithdraw) Validate() error {
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
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenPricingSource0 is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenPricingSource1 is not set")
		}
	}
	return nil
}

func (inst *FundWithdraw) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundWithdraw")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" RequestId", *inst.RequestId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          user_receipt", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                  fund", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("token_pricing_source_0", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("token_pricing_source_1", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj FundWithdraw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundWithdraw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}

// NewFundWithdrawInstruction declares a new FundWithdraw instruction with the provided parameters and accounts.
func NewFundWithdrawInstruction(
	// Parameters:
	request_id uint64,
	// Accounts:
	user ag_solanago.PublicKey,
	userReceipt ag_solanago.PublicKey,
	fund ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	tokenPricingSource0 ag_solanago.PublicKey,
	tokenPricingSource1 ag_solanago.PublicKey) *FundWithdraw {
	return NewFundWithdrawInstructionBuilder().
		SetRequestId(request_id).
		SetUserAccount(user).
		SetUserReceiptAccount(userReceipt).
		SetFundAccount(fund).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetTokenPricingSource0Account(tokenPricingSource0).
		SetTokenPricingSource1Account(tokenPricingSource1)
}