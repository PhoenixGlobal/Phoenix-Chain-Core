package configs

// These are the multipliers for ether denominations.
// Example: To get the von value of an amount in 'gvon', use
//
//    new(big.Int).Mul(value, big.NewInt(configs.GVon))
//
const (
	Von   = 1
	GVon  = 1e9
	PHC   =	1e18
)
