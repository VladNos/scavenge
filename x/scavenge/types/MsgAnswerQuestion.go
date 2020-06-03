package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgAnswerQuestion
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgAnswerQuestion{}

// MsgAnswerQuestion - struct for unjailing jailed validator
type MsgAnswerQuestion struct {
	Worker    sdk.AccAddress `json:"worker" yaml:"worker"`       // address of the workers wallet
	Description  string         `json:"description" yaml:"description"`   // description of the Question
	QuestionID 	 string         `json:"questionID" yaml:"questionID"`     // questionID
	Answer     string         `json:"answer" yaml:"answer"`         // Answer of the scavenge
}

// NewMsgAnswerQuestion creates a new MsgAnswerQuestion instance
func NewMsgAnswerQuestion(scavenger sdk.AccAddress, answer, questionID, Description string) MsgAnswerQuestion {


	return MsgAnswerQuestion{
		Scavenger:    scavenger,
		Answer:     answer,
		QuestionID: questionID,
		Description: description,
	}
}

// AnswerQuestionConst is AnswerQuestion Constant
const AnswerQuestionConst = "AnswerQuestion"

// nolint
func (msg MsgAnswerQuestion) Route() string { return RouterKey }
func (msg MsgAnswerQuestion) Type() string  { return AnswerQuestionConst }
func (msg MsgAnswerQuestion) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Scavenger)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAnswerQuestion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgAnswerQuestion) ValidateBasic() error {
	if msg.Scavenger.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Answer == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "answer can't be empty")
	}
	if msg.QuestionID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "questionID can't be empty")
	}
	

	
	return nil
}
