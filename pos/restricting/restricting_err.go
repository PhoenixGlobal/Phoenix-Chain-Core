package restricting

import (
	"fmt"

	"Phoenix-Chain-Core/libs/common"
)

const (
	RestrictTxPlanSize = 36
)

var (
	ErrParamEpochInvalid                    = common.NewBizError(304001, "The initial epoch for staking cannot be zero")
	ErrCountRestrictPlansInvalid            = common.NewBizError(304002, fmt.Sprintf("The number of the restricting plan cannot be (0, %d]", RestrictTxPlanSize))
	ErrLockedAmountTooLess                  = common.NewBizError(304003, "Total staking amount shall be more than 1 PHC")
	ErrBalanceNotEnough                     = common.NewBizError(304004, "Create plan,the sender balance is not enough in restrict")
	ErrAccountNotFound                      = common.NewBizError(304005, "Account is not found on restricting contract")
	ErrSlashingTooMuch                      = common.NewBizError(304006, "Slashing amount is larger than staking amount")
	ErrStakingAmountEmpty                   = common.NewBizError(304007, "Staking amount cannot be 0")
	ErrAdvanceLockedFundsAmountLessThanZero = common.NewBizError(304008, "Staking lock funds amount cannot be less than or equal to 0")
	ErrReturnLockFundsAmountLessThanZero    = common.NewBizError(304009, "Return lock funds amount cannot be less than or equal to 0")
	ErrSlashingAmountLessThanZero           = common.NewBizError(304010, "Slashing amount cannot be less than 0")
	ErrCreatePlanAmountLessThanZero         = common.NewBizError(304011, "Create plan each amount cannot be less than or equal to 0")
	ErrStakingAmountInvalid                 = common.NewBizError(304012, "The staking amount is less than the return amount")
	ErrRestrictBalanceNotEnough             = common.NewBizError(304013, "The user restricting balance is not enough for staking lock funds")
	ErrCreatePlanAmountLessThanMiniAmount   = common.NewBizError(304014, "Create plan each amount should greater than mini amount")
	ErrRestrictBalanceAndFreeNotEnough      = common.NewBizError(304015, "The user restricting  and free balance is not enough for staking lock funds")
)
