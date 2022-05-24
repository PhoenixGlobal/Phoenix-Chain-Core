package types


//-----------------------------------------------------------------------------
// RoundStepType enum type

// RoundStepType enumerates the state of the consensus state machine
type RoundStepType uint32 // These must be numeric, ordered.

// RoundStepType
const (
	RoundStepPrepareBlock       = RoundStepType(0x01) // PrepareBlock Step
	RoundStepPrepareVote       = RoundStepType(0x02) // PrepareVote Step
	RoundStepPreCommit     = RoundStepType(0x03) // PreCommit Step

	// NOTE: Update IsValid method if you change this!
)

// IsValid returns true if the step is valid, false if unknown/undefined.
func (rs RoundStepType) IsValid() bool {
	return uint8(rs) >= 0x01 && uint8(rs) <= 0x08
}

// String returns a string
func (rs RoundStepType) String() string {
	switch rs {
	case RoundStepPrepareBlock:
		return "RoundStepPrepareBlock"
	case RoundStepPrepareVote:
		return "RoundStepPrepareVote"
	case RoundStepPreCommit:
		return "RoundStepPreCommit"
	default:
		return "RoundStepUnknown" // Cannot panic.
	}
}

