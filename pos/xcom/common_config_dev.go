// +build testpackage xcom

package xcom

import "github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"

func init() {
	log.Info("Init dpos common config", "network name", "DefaultTestNet", "network value", DefaultUnitTestNet)
	GetEc(DefaultUnitTestNet)
}
