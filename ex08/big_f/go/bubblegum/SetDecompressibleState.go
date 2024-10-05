// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetDecompressibleState is the `setDecompressibleState` instruction.
type SetDecompressibleState struct {
	DecompressableState *DecompressibleState

	// [0] = [WRITE] treeAuthority
	//
	// [1] = [SIGNER] treeCreator
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetDecompressibleStateInstructionBuilder creates a new `SetDecompressibleState` instruction builder.
func NewSetDecompressibleStateInstructionBuilder() *SetDecompressibleState {
	nd := &SetDecompressibleState{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetDecompressableState sets the "decompressableState" parameter.
func (inst *SetDecompressibleState) SetDecompressableState(decompressableState DecompressibleState) *SetDecompressibleState {
	inst.DecompressableState = &decompressableState
	return inst
}

// SetTreeAuthorityAccount sets the "treeAuthority" account.
func (inst *SetDecompressibleState) SetTreeAuthorityAccount(treeAuthority ag_solanago.PublicKey) *SetDecompressibleState {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(treeAuthority).WRITE()
	return inst
}

// GetTreeAuthorityAccount gets the "treeAuthority" account.
func (inst *SetDecompressibleState) GetTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTreeCreatorAccount sets the "treeCreator" account.
func (inst *SetDecompressibleState) SetTreeCreatorAccount(treeCreator ag_solanago.PublicKey) *SetDecompressibleState {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(treeCreator).SIGNER()
	return inst
}

// GetTreeCreatorAccount gets the "treeCreator" account.
func (inst *SetDecompressibleState) GetTreeCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst SetDecompressibleState) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetDecompressibleState,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetDecompressibleState) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetDecompressibleState) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DecompressableState == nil {
			return errors.New("DecompressableState parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TreeAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TreeCreator is not set")
		}
	}
	return nil
}

func (inst *SetDecompressibleState) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetDecompressibleState")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DecompressableState", *inst.DecompressableState))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("treeAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  treeCreator", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj SetDecompressibleState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DecompressableState` param:
	err = encoder.Encode(obj.DecompressableState)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetDecompressibleState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DecompressableState`:
	err = decoder.Decode(&obj.DecompressableState)
	if err != nil {
		return err
	}
	return nil
}

// NewSetDecompressibleStateInstruction declares a new SetDecompressibleState instruction with the provided parameters and accounts.
func NewSetDecompressibleStateInstruction(
	// Parameters:
	decompressableState DecompressibleState,
	// Accounts:
	treeAuthority ag_solanago.PublicKey,
	treeCreator ag_solanago.PublicKey) *SetDecompressibleState {
	return NewSetDecompressibleStateInstructionBuilder().
		SetDecompressableState(decompressableState).
		SetTreeAuthorityAccount(treeAuthority).
		SetTreeCreatorAccount(treeCreator)
}