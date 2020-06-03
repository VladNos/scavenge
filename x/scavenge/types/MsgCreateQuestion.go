package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateQuestion
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreateQuestion{}

// MsgCreateQuestion - struct for unjailing jailed validator
type MsgCreateQuestion struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`           // address of the question creator
	Description  string         `json:"description" yaml:"description"`   // description of the Question
	QuestionID 	 string         `json:"questionID" yaml:"questionID"`     // questionID
	Reward       sdk.Coins      `json:"reward" yaml:"reward"`             // reward of the scavenger
	Company 	 string         `json:"company" yaml:"company"`             // answer hash of the scavenge
	Phase 	 string         `json:"phase" yaml:"phase"`             // answer hash of the scavenge
}

// NewMsgCreateQuestion creates a new MsgCreateQuestion instance
func NewMsgCreateQuestion(creator sdk.AccAddress, description, questionID string, reward sdk.Coins, company string, phase string) MsgCreateQuestion {
	return MsgCreateQuestion{
		Creator:      creator,
		Description:  description,
		QuestionID:		questionID,
		Reward:       reward,
		Company:      company,
		Phase:		  phase, 
	}
}

// CreateQuestionConst is CreateQuestion Constant
const CreateQuestionConst = "CreateQuestion"

// nolint
func (msg MsgCreateQuestion) Route() string { return RouterKey }
func (msg MsgCreateQuestion) Type() string  { return CreateQuestionConst }
func (msg MsgCreateQuestion) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateQuestion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateQuestion) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.AnswerHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "answer can't be empty")
	}
	return nil
}
