package main

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/configs"
)

const (
	ipcAPIs  = "admin:1.0 debug:1.0 miner:1.0 net:1.0 personal:1.0 phoenixchain:1.0 rpc:1.0 txgen:1.0 txpool:1.0 web3:1.0"
	httpAPIs = "net:1.0 phoenixchain:1.0 rpc:1.0 web3:1.0"
)

// Tests that a node embedded within a console can be started up properly and
// then terminated by closing the input stream.
func TestConsoleWelcome(t *testing.T) {
	datadir := tmpdir(t)
	defer os.RemoveAll(datadir)
	phoenixchain := runPhoenixChain(t,
		"--datadir", datadir, "--port", "0", "--ipcdisable", "--testnet", "--maxpeers", "60", "--nodiscover", "--nat", "none", "console")

	// Gather all the infos the welcome message needs to contain
	phoenixchain.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	phoenixchain.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	phoenixchain.SetTemplateFunc("gover", runtime.Version)
	phoenixchain.SetTemplateFunc("gethver", func() string { return configs.VersionWithCommit("", "") })
	phoenixchain.SetTemplateFunc("niltime", func() string { return time.Unix(0, 0).Format(time.RFC1123) })
	phoenixchain.SetTemplateFunc("apis", func() string { return ipcAPIs })

	// Verify the actual welcome message to the required template
	phoenixchain.Expect(`
Welcome to the PhoenixChain JavaScript console!

instance: PhoenixChainnetwork/v{{gethver}}/{{goos}}-{{goarch}}/{{gover}}
at block: 0 ({{niltime}})
 datadir: {{.Datadir}}
 modules: {{apis}}

> {{.InputLine "exit"}}
`)
	phoenixchain.ExpectExit()
}

// Tests that a console can be attached to a running node via various means.
func TestIPCAttachWelcome(t *testing.T) {
	// Configure the instance for IPC attachement
	var ipc string
	if runtime.GOOS == "windows" {
		ipc = `\\.\pipe\phoenixchain` + strconv.Itoa(trulyRandInt(100000, 999999))
	} else {
		ws := tmpdir(t)
		defer os.RemoveAll(ws)
		ipc = filepath.Join(ws, "phoenixchain.ipc")
	}
	phoenixchain := runPhoenixChain(t,
		"--port", "0", "--testnet", "--maxpeers", "60", "--nodiscover", "--nat", "none", "--ipcpath", ipc)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, phoenixchain, "ipc:"+ipc, ipcAPIs)

	phoenixchain.Interrupt()
	phoenixchain.ExpectExit()
}

func TestHTTPAttachWelcome(t *testing.T) {
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P
	phoenixchain := runPhoenixChain(t,
		"--port", "0", "--ipcdisable", "--testnet", "--maxpeers", "60", "--nodiscover", "--nat", "none",
		"--rpc", "--rpcport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, phoenixchain, "http://localhost:"+port, httpAPIs)

	phoenixchain.Interrupt()
	phoenixchain.ExpectExit()
}

func TestWSAttachWelcome(t *testing.T) {
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P

	phoenixchain := runPhoenixChain(t,
		"--port", "0", "--ipcdisable", "--testnet", "--maxpeers", "60", "--nodiscover", "--nat", "none",
		"--ws", "--wsport", port /*, "--testnet"*/)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, phoenixchain, "ws://localhost:"+port, httpAPIs)

	phoenixchain.Interrupt()
	phoenixchain.ExpectExit()
}

func testAttachWelcome(t *testing.T, phoenixchain *testphoenixchain, endpoint, apis string) {
	// Attach to a running phoenixchain note and terminate immediately
	attach := runPhoenixChain(t, "attach", endpoint)
	defer attach.ExpectExit()
	attach.CloseStdin()

	// Gather all the infos the welcome message needs to contain
	attach.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	attach.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	attach.SetTemplateFunc("gover", runtime.Version)
	attach.SetTemplateFunc("gethver", func() string { return configs.VersionWithCommit("", "") })
	attach.SetTemplateFunc("niltime", func() string { return time.Unix(0, 0).Format(time.RFC1123) })
	attach.SetTemplateFunc("ipc", func() bool { return strings.HasPrefix(endpoint, "ipc") })
	attach.SetTemplateFunc("datadir", func() string { return phoenixchain.Datadir })
	attach.SetTemplateFunc("apis", func() string { return apis })

	// Verify the actual welcome message to the required template
	attach.Expect(`
Welcome to the PhoenixChain JavaScript console!

instance: PhoenixChainnetwork/v{{gethver}}/{{goos}}-{{goarch}}/{{gover}}
at block: 0 ({{niltime}}){{if ipc}}
 datadir: {{datadir}}{{end}}
 modules: {{apis}}

> {{.InputLine "exit" }}
`)
	attach.ExpectExit()
}

// trulyRandInt generates a crypto random integer used by the console tests to
// not clash network ports with other tests running cocurrently.
func trulyRandInt(lo, hi int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(hi-lo)))
	return int(num.Int64()) + lo
}
