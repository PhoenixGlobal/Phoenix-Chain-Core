//go:generate go-bindata -nometadata -o assets.go -pkg tracers -ignore tracers.go -ignore assets.go ./...
//go:generate gofmt -s -w assets.go

// Package tracers contains the actual JavaScript tracer assets.
package tracers
