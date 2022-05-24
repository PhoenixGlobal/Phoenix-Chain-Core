package main

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cespare/cp"
)

// These tests are 'smoke tests' for the account related
// subcommands and flags.
//
// For most tests, the test files from package accounts
// are copied into a temporary keystore directory.

func tmpDatadirWithKeystore(t *testing.T) string {
	datadir := tmpdir(t)
	keystore := filepath.Join(datadir, "keystore")
	source := filepath.Join("..", "..", "accounts", "keystore", "testdata", "keystore")
	if err := cp.CopyAll(keystore, source); err != nil {
		t.Fatal(err)
	}
	return datadir
}

func TestAccountListEmpty(t *testing.T) {
	phoenixchain := runPhoenixChain(t, "account", "list")
	phoenixchain.ExpectExit()
}

func TestAccountList(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t, "account", "list", "--datadir", datadir)
	defer phoenixchain.ExpectExit()
	if runtime.GOOS == "windows" {
		phoenixchain.Expect(`
Account #0: {0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32} keystore://{{.Datadir}}\keystore\UTC--2016-03-22T12-57-55.920751759Z--7ef5a6135f1fd6a02593eedc869c6d41d934aef8
Account #1: {0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6} keystore://{{.Datadir}}\keystore\aaa
Account #2: {0x9zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg} keystore://{{.Datadir}}\keystore\zzz
`)
	} else {
		phoenixchain.Expect(`
Account #0: {0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32} keystore://{{.Datadir}}/keystore/UTC--2016-03-22T12-57-55.920751759Z--7ef5a6135f1fd6a02593eedc869c6d41d934aef8
Account #1: {0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6} keystore://{{.Datadir}}/keystore/aaa
Account #2: {0x9zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg} keystore://{{.Datadir}}/keystore/zzz
`)
	}
}

func TestAccountNew(t *testing.T) {
	phoenixchain := runPhoenixChain(t, "account", "new", "--lightkdf")
	defer phoenixchain.ExpectExit()
	phoenixchain.Expect(`
Your new account is locked with a password. Please give a password. Do not forget this password.
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "foobar"}}
Repeat passphrase: {{.InputLine "foobar"}}

Your new key was generated
`)

	phoenixchain.ExpectRegexp(`
Public address of the key:   0x[0-9a-z]{38}
Path of the secret key file: .*UTC--.+--[0-9a-f]{40}

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
`)
}

func TestAccountNewBadRepeat(t *testing.T) {
	phoenixchain := runPhoenixChain(t, "account", "new", "--lightkdf")
	defer phoenixchain.ExpectExit()
	phoenixchain.Expect(`
Your new account is locked with a password. Please give a password. Do not forget this password.
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "something"}}
Repeat passphrase: {{.InputLine "something else"}}
Fatal: Passphrases do not match
`)
}

func TestAccountUpdate(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t, "account", "update",
		"--datadir", datadir, "--lightkdf",
		"0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6")
	defer phoenixchain.ExpectExit()
	phoenixchain.Expect(`
Unlocking account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "foobar"}}
Please give a new password. Do not forget this password.
Passphrase: {{.InputLine "foobar2"}}
Repeat passphrase: {{.InputLine "foobar2"}}
`)
}

func TestUnlockFlag(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t,
		"--datadir", datadir, "--ipcdisable", "--testnet", "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0",
		"--unlock", "0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
		"js", "testdata/empty.js")
	phoenixchain.Expect(`
Unlocking account 0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "foobar"}}
`)
	phoenixchain.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
	}
	for _, m := range wantMessages {
		if !strings.Contains(phoenixchain.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagWrongPassword(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6")
	defer phoenixchain.ExpectExit()
	phoenixchain.Expect(`
Unlocking account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "wrong1"}}
Unlocking account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 2/3
Passphrase: {{.InputLine "wrong2"}}
Unlocking account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 3/3
Passphrase: {{.InputLine "wrong3"}}
Fatal: Failed to unlock account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 (could not decrypt key with given passphrase)
`)
}

// https://github.com/ethereum/go-ethereum/issues/1785
func TestUnlockFlagMultiIndex(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "0,2",
		"js", "testdata/empty.js")
	phoenixchain.Expect(`
Unlocking account 0 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "foobar"}}
Unlocking account 2 | Attempt 1/3
Passphrase: {{.InputLine "foobar"}}
`)
	phoenixchain.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
		"=0x9zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg",
	}
	for _, m := range wantMessages {
		if !strings.Contains(phoenixchain.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagPasswordFile(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0",
		"--password", "testdata/passwords.txt", "--unlock", "0,2", "--ipcdisable", "--testnet",
		"js", "testdata/empty.js")
	phoenixchain.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=0x0m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
		"=0x9zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg",
	}
	for _, m := range wantMessages {
		if !strings.Contains(phoenixchain.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagPasswordFileWrongPassword(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	phoenixchain := runPhoenixChain(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--password", "testdata/wrong-passwords.txt", "--unlock", "0,2")
	defer phoenixchain.ExpectExit()
	phoenixchain.Expect(`
Fatal: Failed to unlock account 0 (could not decrypt key with given passphrase)
`)
}

func TestUnlockFlagAmbiguous(t *testing.T) {
	store := filepath.Join("..", "..", "accounts", "keystore", "testdata", "dupes")
	phoenixchain := runPhoenixChain(t,
		"--keystore", store, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6",
		"js", "testdata/empty.js")
	defer phoenixchain.ExpectExit()

	// Helper for the expect template, returns absolute keystore path.
	phoenixchain.SetTemplateFunc("keypath", func(file string) string {
		abs, _ := filepath.Abs(filepath.Join(store, file))
		return abs
	})
	phoenixchain.Expect(`
Unlocking account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "foobar"}}
Multiple key files exist for address 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6:
   keystore://{{keypath "1"}}
   keystore://{{keypath "2"}}
Testing your passphrase against all of them...
Your passphrase unlocked keystore://{{keypath "1"}}
In order to avoid this warning, you need to remove the following duplicate key files:
   keystore://{{keypath "2"}}
`)
	phoenixchain.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6",
	}
	for _, m := range wantMessages {
		if !strings.Contains(phoenixchain.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagAmbiguousWrongPassword(t *testing.T) {
	store := filepath.Join("..", "..", "accounts", "keystore", "testdata", "dupes")
	phoenixchain := runPhoenixChain(t,
		"--keystore", store, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6")
	defer phoenixchain.ExpectExit()

	// Helper for the expect template, returns absolute keystore path.
	phoenixchain.SetTemplateFunc("keypath", func(file string) string {
		abs, _ := filepath.Abs(filepath.Join(store, file))
		return abs
	})
	phoenixchain.Expect(`
Unlocking account 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Passphrase: {{.InputLine "wrong"}}
Multiple key files exist for address 0x73ngt84dryedws7kyt9hflq93zpwsey2m0wqp6:
   keystore://{{keypath "1"}}
   keystore://{{keypath "2"}}
Testing your passphrase against all of them...
Fatal: None of the listed files could be unlocked.
`)
	phoenixchain.ExpectExit()
}
