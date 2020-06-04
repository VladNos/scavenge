package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Scavenge is the Scavenge struct
type Scavenge struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`           // address of the scavenger creator
	Description  string         `json:"description" yaml:"description"`   // description of the scavenge
	SolutionHash string         `json:"solutionHash" yaml:"solutionHash"` // solution hash of the scavenge
	Reward       sdk.Coins      `json:"reward" yaml:"reward"`             // reward of the scavenger
	Solution     string         `json:"solution" yaml:"solution"`         // the solution to the scagenve
	Scavenger    sdk.AccAddress `json:"scavenger" yaml:"scavenger"`       // the scavenger who found the solution
}

// implement fmt.Stringer
func (s Scavenge) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Description: %s
	SolutionHash: %s
	Reward: %s
	Solution: %s
	Scavenger: %s`,
		s.Creator,
		s.Description,
		s.SolutionHash,
		s.Reward,
		s.Solution,
		s.Scavenger,
	))
}

// Commit is the commit struct
type Commit struct {
	Scavenger             sdk.AccAddress `json:"scavenger" yaml:"scavenger"`                         // address of the scavenger scavenger
	SolutionHash          string         `json:"solutionHash" yaml:"solutionHash"`                   // SolutionHash of the scavenge
	SolutionScavengerHash string         `json:"solutionScavengerHash" yaml:"solutionScavengerHash"` // solution hash of the scavenge
}

// implement fmt.Stringer
func (c Commit) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Scavenger: %s
	SolutionHash: %s
	SolutionScavengerHash: %s`,
		c.Scavenger,
		c.SolutionHash,
		c.SolutionScavengerHash,
	))
}

//Question type 

type Question struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`           // address of the question creator
	Description  string         `json:"description" yaml:"description"`   // description of the Question
	QuestionID 	 string         `json:"questionID" yaml:"questionID"`     // questionID
	Reward       sdk.Coins      `json:"reward" yaml:"reward"`             // reward of the scavenger
	Company 	 string         `json:"company" yaml:"company"`             // answer hash of the scavenge
	Phase 	 string         `json:"phase" yaml:"phase"`             // answer hash of the scavenge
}

// implement fmt.Stringer
func (q Question) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Description: %s
	QuestionID: %s
	Reward: %s
	Company: %s
	Phase: %s`,
		q.Creator,
		q.Description,
		q.QuestionID,
		q.Reward,
		q.Company,
		q.Phase,
	))
}


//Answer type
type Answer struct {
	Worker    sdk.AccAddress `json:"worker" yaml:"worker"`       // address of the workers wallet
	Description  string         `json:"description" yaml:"description"`   // description of the Question
	QuestionID 	 string         `json:"questionID" yaml:"questionID"`     // questionID
	Answer     string         `json:"answer" yaml:"answer"`         // Answer of the scavenge
}

// implement fmt.Stringer
func (a Answer) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Worker: %s
	Description: %s
	QuestionID: %s
	Answer: %s`,
		a.Worker,
		a.Description,
		a.QuestionID,
		a.Answer,
	))
}