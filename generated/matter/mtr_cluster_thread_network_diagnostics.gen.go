// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterThreadNetworkDiagnostics] class.
var (
	MTRClusterThreadNetworkDiagnosticsClass     _MTRClusterThreadNetworkDiagnosticsClass
	MTRClusterThreadNetworkDiagnosticsClassOnce sync.Once
)

func getMTRClusterThreadNetworkDiagnosticsClass() _MTRClusterThreadNetworkDiagnosticsClass {
	MTRClusterThreadNetworkDiagnosticsClassOnce.Do(func() {
		MTRClusterThreadNetworkDiagnosticsClass = _MTRClusterThreadNetworkDiagnosticsClass{objc.GetClass("MTRClusterThreadNetworkDiagnostics")}
	})
	return MTRClusterThreadNetworkDiagnosticsClass
}

type _MTRClusterThreadNetworkDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterThreadNetworkDiagnostics] class.
type IMTRClusterThreadNetworkDiagnostics interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadNetworkDiagnostics
type MTRClusterThreadNetworkDiagnostics struct {
	MTRGenericCluster
}

// MTRClusterThreadNetworkDiagnosticsFrom constructs a [MTRClusterThreadNetworkDiagnostics] from an unsafe.Pointer.
func MTRClusterThreadNetworkDiagnosticsFrom(ptr unsafe.Pointer) MTRClusterThreadNetworkDiagnostics {
	return MTRClusterThreadNetworkDiagnostics{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThreadNetworkDiagnosticsClass) Alloc() MTRClusterThreadNetworkDiagnostics {
	rv := objc.Send[MTRClusterThreadNetworkDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterThreadNetworkDiagnosticsClass) New() MTRClusterThreadNetworkDiagnostics {
	rv := objc.Send[MTRClusterThreadNetworkDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThreadNetworkDiagnostics) Init() MTRClusterThreadNetworkDiagnostics {
	rv := objc.Send[MTRClusterThreadNetworkDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThreadNetworkDiagnostics) Autorelease() MTRClusterThreadNetworkDiagnostics {
	rv := objc.Send[MTRClusterThreadNetworkDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThreadNetworkDiagnostics creates a new MTRClusterThreadNetworkDiagnostics instance.
func NewMTRClusterThreadNetworkDiagnostics() MTRClusterThreadNetworkDiagnostics {
	return getMTRClusterThreadNetworkDiagnosticsClass().New()
}




