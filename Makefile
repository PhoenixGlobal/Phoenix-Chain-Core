# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: android ios all test clean

GOBIN = $(shell pwd)/build/bin
GO ?= latest
GPATH = $(shell go env GOPATH)
GORUN = env GO111MODULE=on GOPATH=$(GPATH) go run

phoenixchain:
	build/build_deps.sh
	$(GORUN) build/ci.go install ./commands/phoenixchain
	@echo "Done building."
	@echo "Run \"$(GOBIN)/phoenixchain\" to launch phoenixchain."

phoenixchain-with-mpc:
	build/build_deps.sh
	$(GORUN) build/ci.go install -mpc on ./commands/phoenixchain
	@echo "Done building phoenixchain with mpc."
	@echo "Run \"$(GOBIN)/phoenixchain\" to launch phoenixchain."

all:
	build/build_deps.sh
	$(GORUN) build/ci.go install
	@mv $(GOBIN)/keytool $(GOBIN)/phoenixkey

all-debug:
	build/build_deps.sh
	$(GORUN) build/ci.go install -gcflags on

all-with-mpc:
	build/build_deps.sh
	$(GORUN) build/ci.go install -mpc on

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/geth.aar\" to use the library."

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	./build/clean_deps.sh
	./build/clean_go_build_cache.sh
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./commands/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'