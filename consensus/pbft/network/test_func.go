package network

import (
	ctpyes "Phoenix-Chain-Core/consensus/pbft/types"
)

// SetSendQueueHook
func (h *EngineManager) SetSendQueueHook(f func(*ctpyes.MsgPackage)) {
	h.sendQueueHook = f
}
