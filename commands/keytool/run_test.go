package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/internal/cmdtest"
	"github.com/docker/docker/pkg/reexec"
)

type testKeytool struct {
	*cmdtest.TestCmd
}

// spawns phoenixkey with the given command line args.
func runKeytool(t *testing.T, args ...string) *testKeytool {
	tt := new(testKeytool)
	tt.TestCmd = cmdtest.NewTestCmd(t, tt)
	tt.Run("phoenixkey-test", args...)
	return tt
}

func TestMain(m *testing.M) {
	// Run the app if we've been exec'd as "phoenixkey-test" in runKeytool.
	reexec.Register("phoenixkey-test", func() {
		if err := app.Run(os.Args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	})
	// check if we have been reexec'd
	if reexec.Init() {
		return
	}
	os.Exit(m.Run())
}
