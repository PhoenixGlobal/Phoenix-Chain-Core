package phoenixchain

import (
	"os"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/log"
)

// SetVerbosity sets the global verbosity level (between 0 and 6 - see logger/verbosity.go).
func SetVerbosity(level int) {
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(level), log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}
