// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterWiFiNetworkDiagnostics] class.
var (
	MTRBaseClusterWiFiNetworkDiagnosticsClass     _MTRBaseClusterWiFiNetworkDiagnosticsClass
	MTRBaseClusterWiFiNetworkDiagnosticsClassOnce sync.Once
)

func getMTRBaseClusterWiFiNetworkDiagnosticsClass() _MTRBaseClusterWiFiNetworkDiagnosticsClass {
	MTRBaseClusterWiFiNetworkDiagnosticsClassOnce.Do(func() {
		MTRBaseClusterWiFiNetworkDiagnosticsClass = _MTRBaseClusterWiFiNetworkDiagnosticsClass{objc.GetClass("MTRBaseClusterWiFiNetworkDiagnostics")}
	})
	return MTRBaseClusterWiFiNetworkDiagnosticsClass
}

type _MTRBaseClusterWiFiNetworkDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterWiFiNetworkDiagnostics] class.
type IMTRBaseClusterWiFiNetworkDiagnostics interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkDiagnostics
type MTRBaseClusterWiFiNetworkDiagnostics struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWiFiNetworkDiagnosticsFrom constructs a [MTRBaseClusterWiFiNetworkDiagnostics] from an unsafe.Pointer.
func MTRBaseClusterWiFiNetworkDiagnosticsFrom(ptr unsafe.Pointer) MTRBaseClusterWiFiNetworkDiagnostics {
	return MTRBaseClusterWiFiNetworkDiagnostics{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWiFiNetworkDiagnosticsClass) Alloc() MTRBaseClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterWiFiNetworkDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterWiFiNetworkDiagnosticsClass) New() MTRBaseClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterWiFiNetworkDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWiFiNetworkDiagnostics) Init() MTRBaseClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterWiFiNetworkDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWiFiNetworkDiagnostics) Autorelease() MTRBaseClusterWiFiNetworkDiagnostics {
	rv := objc.Send[MTRBaseClusterWiFiNetworkDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWiFiNetworkDiagnostics creates a new MTRBaseClusterWiFiNetworkDiagnostics instance.
func NewMTRBaseClusterWiFiNetworkDiagnostics() MTRBaseClusterWiFiNetworkDiagnostics {
	return getMTRBaseClusterWiFiNetworkDiagnosticsClass().New()
}




