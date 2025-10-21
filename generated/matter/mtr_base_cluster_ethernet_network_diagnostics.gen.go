// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterEthernetNetworkDiagnostics] class.
var (
	MTRBaseClusterEthernetNetworkDiagnosticsClass     _MTRBaseClusterEthernetNetworkDiagnosticsClass
	MTRBaseClusterEthernetNetworkDiagnosticsClassOnce sync.Once
)

func getMTRBaseClusterEthernetNetworkDiagnosticsClass() _MTRBaseClusterEthernetNetworkDiagnosticsClass {
	MTRBaseClusterEthernetNetworkDiagnosticsClassOnce.Do(func() {
		MTRBaseClusterEthernetNetworkDiagnosticsClass = _MTRBaseClusterEthernetNetworkDiagnosticsClass{objc.GetClass("MTRBaseClusterEthernetNetworkDiagnostics")}
	})
	return MTRBaseClusterEthernetNetworkDiagnosticsClass
}

type _MTRBaseClusterEthernetNetworkDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterEthernetNetworkDiagnostics] class.
type IMTRBaseClusterEthernetNetworkDiagnostics interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEthernetNetworkDiagnostics
type MTRBaseClusterEthernetNetworkDiagnostics struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterEthernetNetworkDiagnosticsFrom constructs a [MTRBaseClusterEthernetNetworkDiagnostics] from an unsafe.Pointer.
func MTRBaseClusterEthernetNetworkDiagnosticsFrom(ptr unsafe.Pointer) MTRBaseClusterEthernetNetworkDiagnostics {
	return MTRBaseClusterEthernetNetworkDiagnostics{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterEthernetNetworkDiagnosticsClass) Alloc() MTRBaseClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterEthernetNetworkDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterEthernetNetworkDiagnosticsClass) New() MTRBaseClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterEthernetNetworkDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterEthernetNetworkDiagnostics) Init() MTRBaseClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterEthernetNetworkDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterEthernetNetworkDiagnostics) Autorelease() MTRBaseClusterEthernetNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterEthernetNetworkDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterEthernetNetworkDiagnostics creates a new MTRBaseClusterEthernetNetworkDiagnostics instance.
func NewMTRBaseClusterEthernetNetworkDiagnostics() MTRBaseClusterEthernetNetworkDiagnostics {
	return getMTRBaseClusterEthernetNetworkDiagnosticsClass().New()
}




