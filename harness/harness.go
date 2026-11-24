// Package harness runs lab implementations against behavioral tests,
// with fault injection (crashes, partitions, delays) to come.
package harness

// Node is the interface every lab implementation satisfies.
type Node interface {
	Start() error
	Stop() error
}
