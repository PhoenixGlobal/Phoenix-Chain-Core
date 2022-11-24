package slashing

import "github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"

var (
	ErrDuplicateSignVerify = common.NewBizError(303000, "Double-signning verification failed")
	ErrSlashingExist       = common.NewBizError(303001, "Punishment has been executed already")
	ErrBlockNumberTooHigh  = common.NewBizError(303002, "BlockNumber for the reported double-spending attack is higher than the current value")
	ErrIntervalTooLong     = common.NewBizError(303003, "Reported evidence expired")
	ErrGetCandidate        = common.NewBizError(303004, "Failed to retrieve the reported validator information")
	ErrAddrMismatch        = common.NewBizError(303005, "The evidence address is inconsistent with the validator address")
	ErrNodeIdMismatch      = common.NewBizError(303006, "NodeId does not match")
	ErrBlsPubKeyMismatch   = common.NewBizError(303007, "BlsPubKey does not match")
	ErrSlashingFail        = common.NewBizError(303008, "Slashing node failed")
	ErrNotValidator        = common.NewBizError(303009, "This node is not a validator")
	ErrSameAddr            = common.NewBizError(303010, "Can't report yourself")
)
