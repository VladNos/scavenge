package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/vladnos/scavenge/x/scavenge/types"
)

// Keeper of the scavenge store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a scavenge keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey) Kee
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetCommit returns the commit of a solution
func (k Keeper) GetCommit(ctx sdk.Context, solutionScavengerHash string) (types.Commit, error) {
	store := ctx.KVStore(k.storeKey)
	var commit types.Commit
	byteKey := []byte(types.CommitPrefix + solutionScavengerHash)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &commit)
	if err != nil {
		return commit, err
	}
	return commit, nil
}

// GetScavenge returns the scavenge information
func (k Keeper) GetScavenge(ctx sdk.Context, solutionHash string) (types.Scavenge, error) {
	store := ctx.KVStore(k.storeKey)
	var scavenge types.Scavenge
	byteKey := []byte(types.ScavengePrefix + solutionHash)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &scavenge)
	if err != nil {
		return scavenge, err
	}
	return scavenge, nil
}

// SetCommit sets a scavenge
func (k Keeper) SetCommit(ctx sdk.Context, commit types.Commit) {
	solutionScavengerHash := commit.SolutionScavengerHash
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(commit)
	key := []byte(types.CommitPrefix + solutionScavengerHash)
	store.Set(key, bz)
}

// SetScavenge sets a scavenge
func (k Keeper) SetScavenge(ctx sdk.Context, scavenge types.Scavenge) {
	solutionHash := scavenge.SolutionHash
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(scavenge)
	key := []byte(types.ScavengePrefix + solutionHash)
	store.Set(key, bz)
}

// DeleteScavenge deletes a scavenge
func (k Keeper) DeleteScavenge(ctx sdk.Context, solutionHash string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(solutionHash))
}

// GetScavengesIterator gets an iterator over all scavnges in which the keys are the solutionHashes and the values are the scavenges
func (k Keeper) GetScavengesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.ScavengePrefix))
}

// GetCommitsIterator gets an iterator over all commits in which the keys are the prefix and solutionHashes and the values are the scavenges
func (k Keeper) GetCommitsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.CommitPrefix))
}
//Question

// Gets the entire Question metadata struct for an ID
func (k Keeper) GetQuestion(ctx sdk.Context, qID string) types.Question {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(qID))
	var question types.Question
	k.cdc.MustUnmarshalBinaryBare(bz, &Question)
	return question
}

// Sets the entire Question metadata struct for an ID
func (k Keeper) SetQuestion(ctx sdk.Context, question types.Question) {
	qID := question.qID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(question)
	key := []byte(types.questionPrefix + qID)
	store.Set(key, bz)
}

// Deletes the entire Question metadata struct for a name
func (k Keeper) DeleteQuestion(ctx sdk.Context, qID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(qID))
}

// SetCreator - sets the current owner of a question
func (k Keeper) SetCreator(ctx sdk.Context, qID string, creator sdk.AccAddress) {
	question := k.GetQuestion(ctx, qID)
	question.Creator = creator
	k.SetQuestion(ctx, question)
}

// GetCreator - get the current owner of a question
func (k Keeper) GetCreator(ctx sdk.Context, qID string) sdk.AccAddress {
	return k.GetQuestion(ctx, qID).Creator
}

// SetDescription - sets the current description of a question
func (k Keeper) SetDescription(ctx sdk.Context, qID string, description string) {
	question := k.GetQuestion(ctx, qID)
	question.Description = description
	k.SetQuestion(ctx, question)
}

// GetDescription - get the current description of a question
func (k Keeper) GetDescription(ctx sdk.Context, qID string) string {
	return k.GetQuestion(ctx, qID).Description
}

// SetReward - sets the current reward of a question
func (k Keeper) SetReward(ctx sdk.Context, qID string, reward sdk.Coins) {
	question := k.GetQuestion(ctx, qID)
	question.Reward = reward
	k.SetQuestion(ctx, question)
}

// GetReward - get the current reward of a question
func (k Keeper) GetReward(ctx sdk.Context, qID string) sdk.Coins {
	return k.GetQuestion(ctx, qID).Reward
}

// SetCompany - sets the current company of a question
func (k Keeper) SetCompany(ctx sdk.Context, qID string, company string) {
	question := k.GetQuestion(ctx, qID)
	question.Company = company
	k.SetQuestion(ctx, question)
}

// GetCompany - get the current company of a question
func (k Keeper) GetCompany(ctx sdk.Context, qID string) string {
	return k.GetQuestion(ctx, qID).Company
}

// SetPhase - sets the current phase of a question
func (k Keeper) SetPhase(ctx sdk.Context, qID string, phase string) {
	question := k.GetQuestion(ctx, qID)
	question.Phase = phase
	k.SetQuestion(ctx, question)
}

// GetPhase - get the current phase of a question
func (k Keeper) GetPhase(ctx sdk.Context, qID string) string {
	return k.GetQuestion(ctx, qID).Phase
}

// GetQuestionIterator gets an iterator over all commits in which the keys are the prefix and solutionHashes and the values are the scavenges
func (k Keeper) GetQuestionIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.QuestionPrefix))
}


//Answer

// Gets the entire Answer metadata struct for an qID and responder
func (k Keeper) GetAnswer(ctx sdk.Context, qID, responder string) types.Answer {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(qID + responder))
	var answer types.Answer
	k.cdc.MustUnmarshalBinaryBare(bz, &Answer)
	return answer
}

// Sets the entire Answer metadata struct for an qID and responder
func (k Keeper) SetAnswer(ctx sdk.Context, answer types.Answer) {
	name := answer.QuestionID + answer.Responder
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(answer)
	key := []byte(types.AnswerPrefix + name)
	store.Set(key, bz)
}

// Deletes the entire Answer metadata struct for a qID and responder
func (k Keeper) DeleteQuestion(ctx sdk.Context, qID, responder string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(qID + responder))
}

// SetAnswerString - sets the current Y of a answer
func (k Keeper) SetAnswerString(ctx sdk.Context, qID, responder string, answS string) {
	answer := k.GetAnswer(ctx, qID, responder)
	answer.Answer= answS
	k.SetAnswer(ctx,  answer)
}

// GetAnswerString - get the current Y of a answer
func (k Keeper) GetAnswerString(ctx sdk.Context, qID, responder string) string {
	return k.GetAnswer(ctx, qID, responder).Answer
}

// GetAnswerIterator gets an iterator over all commits in which the keys are the prefix and solutionHashes and the values are the scavenges
func (k Keeper) GetAnswerIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.AnswerPrefix))
}

/*
A UpperCase B LowwerCase The object
X     -     Y     -      The Param
C                        the Key

// SetX - sets the current Y of a B
func (k Keeper) SetX(ctx sdk.Context, C string, Y Z) {
	B := k.GetA(ctx, C)
	B.X = Y
	k.SetA(ctx, B)
}

// GetX - get the current Y of a B
func (k Keeper) GetX(ctx sdk.Context, C string) Z {
	return k.GetA(ctx, C).X
}
*/
