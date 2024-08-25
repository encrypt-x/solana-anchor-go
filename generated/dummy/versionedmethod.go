// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dummy

import (
	"errors"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// VersionedMethod is the `versioned_method` instruction.
type VersionedMethod struct {
	Data VersionedState

	// [0] = [WRITE] user_token_amount
	//
	// [1] = [SIGNER] user
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewVersionedMethodInstructionBuilder creates a new `VersionedMethod` instruction builder.
func NewVersionedMethodInstructionBuilder() *VersionedMethod {
	nd := &VersionedMethod{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *VersionedMethod) SetData(data VersionedState) *VersionedMethod {
	inst.Data = data
	return inst
}

// SetUserTokenAmountAccount sets the "user_token_amount" account.
func (inst *VersionedMethod) SetUserTokenAmountAccount(userTokenAmount ag_solanago.PublicKey) *VersionedMethod {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(userTokenAmount).WRITE()
	return inst
}

// GetUserTokenAmountAccount gets the "user_token_amount" account.
func (inst *VersionedMethod) GetUserTokenAmountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *VersionedMethod) SetUserAccount(user ag_solanago.PublicKey) *VersionedMethod {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *VersionedMethod) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst VersionedMethod) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_VersionedMethod,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst VersionedMethod) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *VersionedMethod) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UserTokenAmount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.User is not set")
		}
	}
	return nil
}

func (inst *VersionedMethod) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("VersionedMethod")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("user_token_amount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             user", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj VersionedMethod) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	{
		tmp := versionedStateContainer{}
		switch realvalue := obj.Data.(type) {
		case *VersionedStateV1Tuple:
			tmp.Enum = 0
			tmp.V1 = *realvalue
		case *VersionedStateV2Tuple:
			tmp.Enum = 1
			tmp.V2 = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	return nil
}
func (obj *VersionedMethod) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	{
		tmp := new(versionedStateContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.Data = &tmp.V1
		case 1:
			obj.Data = &tmp.V2
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	return nil
}

// NewVersionedMethodInstruction declares a new VersionedMethod instruction with the provided parameters and accounts.
func NewVersionedMethodInstruction(
	// Parameters:
	data VersionedState,
	// Accounts:
	userTokenAmount ag_solanago.PublicKey,
	user ag_solanago.PublicKey) *VersionedMethod {
	return NewVersionedMethodInstructionBuilder().
		SetData(data).
		SetUserTokenAmountAccount(userTokenAmount).
		SetUserAccount(user)
}