package scavenge

import (
	"github.com/vladnos/scavenge/x/scavenge/keeper"
	"github.com/vladnos/scavenge/x/scavenge/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	
	NewMsgCreateScavenge = types.NewMsgCreateScavenge
	NewMsgCommitSolution = types.NewMsgCommitSolution
	NewMsgRevealSolution = types.NewMsgRevealSolution
	NewMsgCreateQuestion = types.NewMsgCreateQuestion
	NewMsgAnswerQuestion = types.NewMsgAnswerQuestion
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgCreateScavenge = types.MsgCreateScavenge
	MsgCommitSolution = types.MsgCommitSolution
	MsgRevealSolution = types.MsgRevealSolution
	MsgCreateQuestion = types.MsgCreateQuestion
	MsgAnswerQuestion = types.MsgAnswerQuestion
)
