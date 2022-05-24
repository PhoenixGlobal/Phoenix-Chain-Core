package rules

import (
	"Phoenix-Chain-Core/consensus/pbft/protocols"
	"Phoenix-Chain-Core/consensus/pbft/state"
)

type VoteError interface {
	error
	Discard() bool //Is the error need discard
}

type voteError struct {
	s       string
	discard bool
}

func (v *voteError) Error() string {
	return v.s
}

func (v *voteError) Discard() bool {
	return v.discard
}

func newVoteError(s string, discard bool) VoteError {
	return &voteError{
		s:       s,
		discard: discard,
	}
}

type VoteRules interface {
	// Determine if the resulting vote is allowed to be sent
	AllowVote(vote *protocols.PrepareVote) VoteError
	AllowPreCommit(vote *protocols.PreCommit) VoteError
}

type baseVoteRules struct {
	viewState *state.ViewState
}

// Determine if voting is possible
// viewNumber should be equal
// have you voted
func (v *baseVoteRules) AllowVote(vote *protocols.PrepareVote) VoteError {
	if v.viewState.BlockNumber() != vote.BlockNumber||v.viewState.ViewNumber() != vote.ViewNumber {
		return newVoteError("v.viewState.BlockNumber() != vote.BlockNumber||v.viewState.ViewNumber() != vote.ViewNumber", true)
	}

	if v.viewState.HadSendPreVote(vote) {
		return newVoteError("v.viewState.HadSendPreVote(vote.BlockNumber)", true)
	}
	return nil
}

// Determine if voting is possible
// viewNumber should be equal
// have you voted
func (v *baseVoteRules) AllowPreCommit(vote *protocols.PreCommit) VoteError {
	if v.viewState.BlockNumber() != vote.BlockNumber||v.viewState.ViewNumber() != vote.ViewNumber {
		return newVoteError("v.viewState.BlockNumber() != vote.BlockNumber||v.viewState.ViewNumber() != vote.ViewNumber", true)
	}

	if v.viewState.HadSendPreCommit(vote) {
		return newVoteError("v.viewState.HadSendPreCommit(vote.BlockNumber)", true)
	}
	return nil
}

func NewVoteRules(viewState *state.ViewState) VoteRules {
	return &baseVoteRules{
		viewState: viewState,
	}

}
