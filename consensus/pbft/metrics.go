package pbft

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/metrics"
)

var (
	blockMinedGauage       = metrics.NewRegisteredGauge("pbft/gauage/block/mined", nil)
	viewChangedTimer       = metrics.NewRegisteredTimer("pbft/timer/view/changed", nil)
	blockQCCollectedGauage = metrics.NewRegisteredGauge("pbft/gauage/block/qc_collected", nil)

	blockProduceMeter          = metrics.NewRegisteredMeter("pbft/meter/block/produce", nil)
	blockCheckFailureMeter     = metrics.NewRegisteredMeter("pbft/meter/block/check_failure", nil)
	signatureCheckFailureMeter = metrics.NewRegisteredMeter("pbft/meter/signature/check_failure", nil)
	blockConfirmedMeter        = metrics.NewRegisteredMeter("pbft/meter/block/confirmed", nil)

	masterCounter    = metrics.NewRegisteredCounter("pbft/counter/view/count", nil)
	consensusCounter = metrics.NewRegisteredCounter("pbft/counter/consensus/count", nil)
	minedCounter     = metrics.NewRegisteredCounter("pbft/counter/mined/count", nil)

	viewNumberGauage          = metrics.NewRegisteredGauge("pbft/gauage/view/number", nil)
	epochNumberGauage         = metrics.NewRegisteredGauge("pbft/gauage/epoch/number", nil)
	proposerIndexGauage       = metrics.NewRegisteredGauge("pbft/gauage/proposer/index", nil)
	validatorCountGauage      = metrics.NewRegisteredGauge("pbft/gauage/validator/count", nil)
	blockNumberGauage         = metrics.NewRegisteredGauge("pbft/gauage/block/number", nil)
	highestQCNumberGauage     = metrics.NewRegisteredGauge("pbft/gauage/block/qc/number", nil)
	highestLockedNumberGauage = metrics.NewRegisteredGauge("pbft/gauage/block/locked/number", nil)
	highestCommitNumberGauage = metrics.NewRegisteredGauge("pbft/gauage/block/commit/number", nil)
)
