package gov

import (
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/ethereum/p2p/discover"
)

type TallyResult struct {
	ProposalID    common.Hash    `json:"proposalID"`
	Yeas          uint64         `json:"yeas"`
	Nays          uint64         `json:"nays"`
	Abstentions   uint64         `json:"abstentions"`
	AccuVerifiers uint64         `json:"accuVerifiers"`
	Status        ProposalStatus `json:"status"`
	CanceledBy    common.Hash    `json:"canceledBy"`
}

type VoteInfo struct {
	ProposalID common.Hash     `json:"proposalID"`
	VoteNodeID discover.NodeID `json:"voteNodeID"`
	VoteOption VoteOption      `json:"voteOption"`
}

type VoteValue struct {
	VoteNodeID discover.NodeID `json:"voteNodeID"`
	VoteOption VoteOption      `json:"voteOption"`
}

type ActiveVersionValue struct {
	ActiveVersion uint32 `json:"ActiveVersion"`
	ActiveBlock   uint64 `json:"ActiveBlock"`
}
type ParamItem struct {
	Module string `json:"Module"`
	Name   string `json:"Name"`
	Desc   string `json:"Desc"`
}

type ParamValue struct {
	StaleValue  string `json:"StaleValue"`
	Value       string `json:"Value"`
	ActiveBlock uint64 `json:"ActiveBlock"`
}

type GovernParam struct {
	ParamItem     *ParamItem
	ParamValue    *ParamValue
	ParamVerifier ParamVerifier `json:"-"`
}
