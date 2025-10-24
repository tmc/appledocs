// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterWiFiNetworkDiagnostics] class.
var (
	MTRClusterWiFiNetworkDiagnosticsClass     _MTRClusterWiFiNetworkDiagnosticsClass
	MTRClusterWiFiNetworkDiagnosticsClassOnce sync.Once
)

func getMTRClusterWiFiNetworkDiagnosticsClass() _MTRClusterWiFiNetworkDiagnosticsClass {
	MTRClusterWiFiNetworkDiagnosticsClassOnce.Do(func() {
		MTRClusterWiFiNetworkDiagnosticsClass = _MTRClusterWiFiNetworkDiagnosticsClass{objc.GetClass("MTRClusterWiFiNetworkDiagnostics")}
	})
	return MTRClusterWiFiNetworkDiagnosticsClass
}

type _MTRClusterWiFiNetworkDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterWiFiNetworkDiagnostics] class.
type IMTRClusterWiFiNetworkDiagnostics interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkDiagnostics
type MTRClusterWiFiNetworkDiagnostics struct {
	MTRGenericCluster
}

// MTRClusterWiFiNetworkDiagnosticsFrom constructs a [MTRClusterWiFiNetworkDiagnostics] from an unsafe.Pointer.
func MTRClusterWiFiNetworkDiagnosticsFrom(ptr unsafe.Pointer) MTRClusterWiFiNetworkDiagnostics {
	return MTRClusterWiFiNetworkDiagnostics{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWiFiNetworkDiagnosticsClass) Alloc() MTRClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRClusterWiFiNetworkDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterWiFiNetworkDiagnosticsClass) New() MTRClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRClusterWiFiNetworkDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWiFiNetworkDiagnostics) Init() MTRClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRClusterWiFiNetworkDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWiFiNetworkDiagnostics) Autorelease() MTRClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRClusterWiFiNetworkDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWiFiNetworkDiagnostics creates a new MTRClusterWiFiNetworkDiagnostics instance.
func NewMTRClusterWiFiNetworkDiagnostics() MTRClusterWiFiNetworkDiagnostics {
	return getMTRClusterWiFiNetworkDiagnosticsClass().New()
}
