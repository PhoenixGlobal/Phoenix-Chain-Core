package node

import (
	"crypto/ecdsa"
	"path/filepath"
	"reflect"

	"Phoenix-Chain-Core/ethereum/core/db/rawdb"

	"Phoenix-Chain-Core/ethereum/accounts"
	"Phoenix-Chain-Core/libs/ethdb"
	"Phoenix-Chain-Core/libs/event"
	"Phoenix-Chain-Core/ethereum/p2p"
	"Phoenix-Chain-Core/libs/rpc"
)

// ServiceContext is a collection of service independent options inherited from
// the protocol stack, that is passed to all constructors to be optionally used;
// as well as utility methods to operate on the service environment.
type ServiceContext struct {
	config         *Config
	services       map[reflect.Type]Service // Index of the already constructed services
	EventMux       *event.TypeMux           // Event multiplexer used for decoupled notifications
	AccountManager *accounts.Manager        // Account manager created by the node.
	serverConfig   p2p.Config
}

func NewServiceContext(config *Config, services map[reflect.Type]Service, EventMux *event.TypeMux, AccountManager *accounts.Manager) *ServiceContext {
	return &ServiceContext{
		config:         config,
		services:       services,
		EventMux:       EventMux,
		AccountManager: AccountManager,
	}
}

// OpenDatabase opens an existing database with the given name (or creates one
// if no previous can be found) from within the node's data directory. If the
// node is an ephemeral one, a memory database is returned.
func (ctx *ServiceContext) OpenDatabase(name string, cache int, handles int, namespace string) (ethdb.Database, error) {
	if ctx.config.DataDir == "" {
		return rawdb.NewMemoryDatabase(), nil
	}
	return rawdb.NewLevelDBDatabase(ctx.config.ResolvePath(name), cache, handles, namespace)
}

// OpenDatabaseWithFreezer opens an existing database with the given name (or
// creates one if no previous can be found) from within the node's data directory,
// also attaching a chain freezer to it that moves ancient chain data from the
// database to immutable append-only files. If the node is an ephemeral one, a
// memory database is returned.
func (ctx *ServiceContext) OpenDatabaseWithFreezer(name string, cache int, handles int, freezer string, namespace string) (ethdb.Database, error) {
	if ctx.config.DataDir == "" {
		return rawdb.NewMemoryDatabase(), nil
	}
	root := ctx.config.ResolvePath(name)

	switch {
	case freezer == "":
		freezer = filepath.Join(root, "ancient")
	case !filepath.IsAbs(freezer):
		freezer = ctx.config.ResolvePath(freezer)
	}
	return rawdb.NewLevelDBDatabaseWithFreezer(root, cache, handles, freezer, namespace)
}

func (ctx *ServiceContext) ResolveFreezerPath(name string, freezer string) string {
	root := ctx.config.ResolvePath(name)
	switch {
	case freezer == "":
		freezer = filepath.Join(root, "ancient")
	case !filepath.IsAbs(freezer):
		freezer = ctx.config.ResolvePath(freezer)
	}
	return freezer
}

// ResolvePath resolves a user path into the data directory if that was relative
// and if the user actually uses persistent storage. It will return an empty string
// for emphemeral storage and the user's own input for absolute paths.
func (ctx *ServiceContext) ResolvePath(path string) string {
	return ctx.config.ResolvePath(path)
}

func (ctx *ServiceContext) GenesisPath() string {
	return ctx.config.GenesisPath()
}

// Service retrieves a currently running service registered of a specific type.
func (ctx *ServiceContext) Service(service interface{}) error {
	element := reflect.ValueOf(service).Elem()
	if running, ok := ctx.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return ErrServiceUnknown
}

func (ctx *ServiceContext) NodePriKey() *ecdsa.PrivateKey {
	return ctx.serverConfig.PrivateKey
}

// ExtRPCEnabled returns the indicator whether node enables the external
// RPC(http, ws or graphql).
func (ctx *ServiceContext) ExtRPCEnabled() bool {
	return ctx.config.ExtRPCEnabled()
}

// ServiceConstructor is the function signature of the constructors needed to be
// registered for service instantiation.
type ServiceConstructor func(ctx *ServiceContext) (Service, error)

// Service is an individual protocol that can be registered into a node.
//
// Notes:
//
// • Service life-cycle management is delegated to the node. The service is allowed to
// initialize itself upon creation, but no goroutines should be spun up outside of the
// Start method.
//
// • Restart logic is not required as the node will create a fresh instance
// every time a service is started.
type Service interface {

	// Protocols retrieves the P2P protocols the service wishes to start.
	Protocols() []p2p.Protocol

	// APIs retrieves the list of RPC descriptors the service provides
	APIs() []rpc.API

	// Start is called after all services have been constructed and the networking
	// layer was also initialized to spawn any goroutines required by the service.
	Start(server *p2p.Server) error

	// Stop terminates all goroutines belonging to the service, blocking until they
	// are all terminated.
	Stop() error
}
