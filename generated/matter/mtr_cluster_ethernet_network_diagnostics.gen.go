// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterEthernetNetworkDiagnostics] class.
var (
	MTRClusterEthernetNetworkDiagnosticsClass     _MTRClusterEthernetNetworkDiagnosticsClass
	MTRClusterEthernetNetworkDiagnosticsClassOnce sync.Once
)

func getMTRClusterEthernetNetworkDiagnosticsClass() _MTRClusterEthernetNetworkDiagnosticsClass {
	MTRClusterEthernetNetworkDiagnosticsClassOnce.Do(func() {
		MTRClusterEthernetNetworkDiagnosticsClass = _MTRClusterEthernetNetworkDiagnosticsClass{objc.GetClass("MTRClusterEthernetNetworkDiagnostics")}
	})
	return MTRClusterEthernetNetworkDiagnosticsClass
}

type _MTRClusterEthernetNetworkDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterEthernetNetworkDiagnostics] class.
type IMTRClusterEthernetNetworkDiagnostics interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEthernetNetworkDiagnostics
type MTRClusterEthernetNetworkDiagnostics struct {
	MTRGenericCluster
}

// MTRClusterEthernetNetworkDiagnosticsFrom constructs a [MTRClusterEthernetNetworkDiagnostics] from an unsafe.Pointer.
func MTRClusterEthernetNetworkDiagnosticsFrom(ptr unsafe.Pointer) MTRClusterEthernetNetworkDiagnostics {
	return MTRClusterEthernetNetworkDiagnostics{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterEthernetNetworkDiagnosticsClass) Alloc() MTRClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRClusterEthernetNetworkDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterEthernetNetworkDiagnosticsClass) New() MTRClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRClusterEthernetNetworkDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterEthernetNetworkDiagnostics) Init() MTRClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRClusterEthernetNetworkDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterEthernetNetworkDiagnostics) Autorelease() MTRClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRClusterEthernetNetworkDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterEthernetNetworkDiagnostics creates a new MTRClusterEthernetNetworkDiagnostics instance.
func NewMTRClusterEthernetNetworkDiagnostics() MTRClusterEthernetNetworkDiagnostics {
	return getMTRClusterEthernetNetworkDiagnosticsClass().New()
}
