package network

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/consensus/pbft/protocols"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/metrics"
)

var (

	// Network delay record
	propPeerLatencyMeter = metrics.NewRegisteredMeter("pbft/prop/pong/latency", nil)

	// PrepareBlockMsg
	propPrepareBlockInPacketsMeter  = metrics.NewRegisteredMeter("pbft/prop/prepare_blocks/in/packets", nil)
	propPrepareBlockInTrafficMeter  = metrics.NewRegisteredMeter("pbft/prop/prepare_blocks/in/traffic", nil)
	propPrepareBlockOutPacketsMeter = metrics.NewRegisteredMeter("pbft/prop/prepare_blocks/out/packets", nil)
	propPrepareBlockOutTrafficMeter = metrics.NewRegisteredMeter("pbft/prop/prepare_blocks/out/traffic", nil)

	// PrepareVoteMsg
	propPrepareVoteInPacketsMeter  = metrics.NewRegisteredMeter("pbft/prop/prepare_vote/in/packets", nil)
	propPrepareVoteInTrafficMeter  = metrics.NewRegisteredMeter("pbft/prop/prepare_vote/in/traffic", nil)
	propPrepareVoteOutPacketsMeter = metrics.NewRegisteredMeter("pbft/prop/prepare_vote/out/packets", nil)
	propPrepareVoteOutTrafficMeter = metrics.NewRegisteredMeter("pbft/prop/prepare_vote/out/traffic", nil)

	// ViewChangeMsg
	propViewChangeInPacketsMeter  = metrics.NewRegisteredMeter("pbft/prop/view_change/in/packets", nil)
	propViewChangeInTrafficMeter  = metrics.NewRegisteredMeter("pbft/prop/view_change/in/traffic", nil)
	propViewChangeOutPacketsMeter = metrics.NewRegisteredMeter("pbft/prop/view_change/out/packets", nil)
	propViewChangeOutTrafficMeter = metrics.NewRegisteredMeter("pbft/prop/view_change/out/traffic", nil)

	// PrepareBlockHashMsg
	propPrepareBlockHashInPacketsMeter  = metrics.NewRegisteredMeter("pbft/prop/prepare_block_hash/in/packets", nil)
	propPrepareBlockHashInTrafficMeter  = metrics.NewRegisteredMeter("pbft/prop/prepare_block_hash/in/traffic", nil)
	propPrepareBlockHashOutPacketsMeter = metrics.NewRegisteredMeter("pbft/prop/prepare_block_hash/out/packets", nil)
	propPrepareBlockHashOutTrafficMeter = metrics.NewRegisteredMeter("pbft/prop/prepare_block_hash/out/traffic", nil)

	// GetPrepareBlockMsg
	reqGetPrepareBlockInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/get_prepare_block/in/packets", nil)
	reqGetPrepareBlockInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/get_prepare_block/in/traffic", nil)
	reqGetPrepareBlockOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/get_prepare_block/out/packets", nil)
	reqGetPrepareBlockOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/get_prepare_block/out/traffic", nil)

	// GetQuorumCertMsg
	reqGetQuorumCertInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/get_quorum_cert/in/packets", nil)
	reqGetQuorumCertInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/get_quorum_cert/in/traffic", nil)
	reqGetQuorumCertOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/get_quorum_cert/out/packets", nil)
	reqGetQuorumCertOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/get_quorum_cert/out/traffic", nil)

	// BlockQuorumCertMsg
	reqBlockQuorumCertInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/block_quorum_cert/in/packets", nil)
	reqBlockQuorumCertInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/block_quorum_cert/in/traffic", nil)
	reqBlockQuorumCertOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/block_quorum_cert/out/packets", nil)
	reqBlockQuorumCertOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/block_quorum_cert/out/traffic", nil)

	// GetPrepareVoteMsg
	reqGetPrepareVoteInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/get_prepare_vote/in/packets", nil)
	reqGetPrepareVoteInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/get_prepare_vote/in/traffic", nil)
	reqGetPrepareVoteOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/get_prepare_vote/out/packets", nil)
	reqGetPrepareVoteOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/get_prepare_vote/out/traffic", nil)

	// PrepareVotesMsg
	reqPrepareVotesInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/prepare_votes/in/packets", nil)
	reqPrepareVotesInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/prepare_votes/in/traffic", nil)
	reqPrepareVotesOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/prepare_votes/out/packets", nil)
	reqPrepareVotesOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/prepare_votes/out/traffic", nil)

	// GetQCBlockListMsg
	reqGetQCBlockListInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/get_qc_block_list/in/packets", nil)
	reqGetQCBlockListInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/get_qc_block_list/in/traffic", nil)
	reqGetQCBlockListOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/get_qc_block_list/out/packets", nil)
	reqGetQCBlockListOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/get_qc_block_list/out/traffic", nil)

	// QCBlockListMsg
	reqQCBlockListInPacketsMeter  = metrics.NewRegisteredMeter("pbft/req/qc_block_list/in/packets", nil)
	reqQCBlockListInTrafficMeter  = metrics.NewRegisteredMeter("pbft/req/qc_block_list/in/traffic", nil)
	reqQCBlockListOutPacketsMeter = metrics.NewRegisteredMeter("pbft/req/qc_block_list/out/packets", nil)
	reqQCBlockListOutTrafficMeter = metrics.NewRegisteredMeter("pbft/req/qc_block_list/out/traffic", nil)

	// Unmatched message type
	miscInPacketsMeter  = metrics.NewRegisteredMeter("pbft/misc/in/packets", nil)
	miscInTrafficMeter  = metrics.NewRegisteredMeter("pbft/misc/in/traffic", nil)
	miscOutPacketsMeter = metrics.NewRegisteredMeter("pbft/misc/out/packets", nil)
	miscOutTrafficMeter = metrics.NewRegisteredMeter("pbft/misc/out/traffic", nil)

	messageGossipMeter = metrics.NewRegisteredMeter("pbft/meter/message/gossip", nil)
	messageRepeatMeter = metrics.NewRegisteredMeter("pbft/meter/message/repeat", nil)

	neighborPeerGauage = metrics.NewRegisteredGauge("pbft/gauage/peer/value", nil)
)

// meteredMsgReadWriter is a wrapper around a p2p.MsgReadWriter, capable of
// accumulating the above defined metrics based on the data stream contents.
type meteredMsgReadWriter struct {
	p2p.MsgReadWriter     // Wrapped message stream to meter
	version           int // Protocol version to select correct meters
}

// newMeteredMsgWriter wraps a p2p MsgReadWriter with metering support. If the
// metrics system is disabled, this function returns the original object.
func newMeteredMsgWriter(rw p2p.MsgReadWriter) p2p.MsgReadWriter {
	if !metrics.Enabled {
		return rw
	}
	return &meteredMsgReadWriter{MsgReadWriter: rw}
}

// Init sets the protocol version used by the stream to know which meters to
// increment in case of overlapping message ids between protocol versions.
func (rw *meteredMsgReadWriter) Init(version int) {
	rw.version = version
}

func (rw *meteredMsgReadWriter) ReadMsg() (p2p.Msg, error) {
	// Read the message and short circuit in case of an error
	msg, err := rw.MsgReadWriter.ReadMsg()
	if err != nil {
		return msg, err
	}
	packets, traffic := miscInPacketsMeter, miscInTrafficMeter
	switch {
	case msg.Code == protocols.PrepareBlockMsg:
		packets, traffic = propPrepareBlockInPacketsMeter, propPrepareBlockInTrafficMeter
	case msg.Code == protocols.PrepareVoteMsg:
		packets, traffic = propPrepareVoteInPacketsMeter, propPrepareVoteInTrafficMeter
	case msg.Code == protocols.ViewChangeMsg:
		packets, traffic = propViewChangeInPacketsMeter, propViewChangeInTrafficMeter
	case msg.Code == protocols.GetPrepareBlockMsg:
		packets, traffic = reqGetPrepareBlockInPacketsMeter, reqGetPrepareBlockInTrafficMeter
	case msg.Code == protocols.GetBlockQuorumCertMsg:
		packets, traffic = reqGetQuorumCertInPacketsMeter, reqGetQuorumCertInTrafficMeter
	case msg.Code == protocols.BlockQuorumCertMsg:
		packets, traffic = reqBlockQuorumCertInPacketsMeter, reqBlockQuorumCertInTrafficMeter
	case msg.Code == protocols.GetPrepareVoteMsg:
		packets, traffic = reqGetPrepareVoteInPacketsMeter, reqGetPrepareVoteInTrafficMeter
	case msg.Code == protocols.PrepareVotesMsg:
		packets, traffic = reqPrepareVotesInPacketsMeter, reqPrepareVotesInTrafficMeter
	case msg.Code == protocols.GetQCBlockListMsg:
		packets, traffic = reqGetQCBlockListInPacketsMeter, reqGetQCBlockListInTrafficMeter
	case msg.Code == protocols.QCBlockListMsg:
		packets, traffic = reqQCBlockListInPacketsMeter, reqQCBlockListInTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	return msg, err
}

func (rw *meteredMsgReadWriter) WriteMsg(msg p2p.Msg) error {
	// Account for the data traffic
	packets, traffic := miscOutPacketsMeter, miscOutTrafficMeter
	switch {
	case msg.Code == protocols.PrepareBlockMsg:
		packets, traffic = propPrepareBlockOutPacketsMeter, propPrepareBlockOutTrafficMeter
		common.PrepareBlockPBFTEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.PrepareVoteMsg:
		packets, traffic = propPrepareVoteOutPacketsMeter, propPrepareVoteOutTrafficMeter
		common.PrepareVoteEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.ViewChangeMsg:
		packets, traffic = propViewChangeOutPacketsMeter, propViewChangeOutTrafficMeter
		common.ViewChangeEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.GetPrepareBlockMsg:
		//packets, traffic = reqGetPrepareBlockOutPacketsMeter, reqGetPrepareBlockOutTrafficMeter
		common.GetPrepareBlockEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.PrepareBlockHashMsg:
		//packets, traffic = propPrepareBlockHashOutPacketsMeter, propPrepareBlockHashOutTrafficMeter
		common.PrepareBlockHashEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.GetPrepareVoteMsg:
		//packets, traffic = reqGetPrepareBlockOutPacketsMeter, reqGetPrepareVoteOutTrafficMeter
		common.GetPrepareVoteEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.GetBlockQuorumCertMsg:
		//packets, traffic = reqGetQuorumCertOutPacketsMeter, reqGetQuorumCertOutTrafficMeter
		common.GetBlockQuorumCertEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.BlockQuorumCertMsg:
		//packets, traffic = reqBlockQuorumCertOutPacketsMeter, reqBlockQuorumCertOutTrafficMeter
		common.BlockQuorumCertEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.PrepareVotesMsg:
		//packets, traffic = reqPrepareVotesOutPacketsMeter, reqPrepareVotesOutTrafficMeter
		common.PrepareVotesEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.GetQCBlockListMsg:
		//packets, traffic = reqGetQCBlockListOutPacketsMeter, reqGetQCBlockListOutTrafficMeter
		common.GetQCBlockListEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.QCBlockListMsg:
		//packets, traffic = reqQCBlockListOutPacketsMeter, reqQCBlockListOutTrafficMeter
		common.QCBlockListEgressTrafficMeter.Mark(int64(msg.Size))

	case msg.Code == protocols.PBFTStatusMsg:
		common.PBFTStatusEgressTrafficMeter.Mark(int64(msg.Size))
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	// Send the packet to the p2p layer
	return rw.MsgReadWriter.WriteMsg(msg)
}
