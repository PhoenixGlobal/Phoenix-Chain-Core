package wal

import (
	"fmt"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/protocols"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/rlp"
)

// struct SendPrepareBlock for rlp decode
type MessageSendPrepareBlock struct {
	Timestamp uint64
	Data      *protocols.SendPrepareBlock
}

// struct SendPrepareVote for rlp decode
type MessageSendPrepareVote struct {
	Timestamp uint64
	Data      *protocols.SendPrepareVote
}

// struct SendPrepareVote for rlp decode
type MessageSendPreCommit struct {
	Timestamp uint64
	Data      *protocols.SendPreCommit
}

// struct SendViewChange for rlp decode
type MessageSendViewChange struct {
	Timestamp uint64
	Data      *protocols.SendViewChange
}

// struct ConfirmedViewChange for rlp decode
type MessageConfirmedViewChange struct {
	Timestamp uint64
	Data      *protocols.ConfirmedViewChange
}

func WALDecode(pack []byte, msgType uint16) (interface{}, error) {
	switch msgType {
	case protocols.ConfirmedViewChangeMsg:
		var j MessageConfirmedViewChange
		if err := rlp.DecodeBytes(pack, &j); err != nil {
			return nil, err

		}
		return j.Data, nil

	case protocols.SendViewChangeMsg:
		var j MessageSendViewChange
		if err := rlp.DecodeBytes(pack, &j); err != nil {
			return nil, err

		}
		return j.Data, nil

	case protocols.SendPrepareBlockMsg:
		var j MessageSendPrepareBlock
		if err := rlp.DecodeBytes(pack, &j); err != nil {
			return nil, err
		}
		return j.Data, nil

	case protocols.SendPrepareVoteMsg:
		var j MessageSendPrepareVote
		if err := rlp.DecodeBytes(pack, &j); err != nil {
			return nil, err

		}
		return j.Data, nil

	case protocols.SendPreCommitMsg:
		var j MessageSendPreCommit
		if err := rlp.DecodeBytes(pack, &j); err != nil {
			return nil, err

		}
		return j.Data, nil
	}
	panic(fmt.Sprintf("invalid msg type %d", msgType))
}
