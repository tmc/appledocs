// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterSoftwareDiagnostics] class.
var (
	MTRClusterSoftwareDiagnosticsClass     _MTRClusterSoftwareDiagnosticsClass
	MTRClusterSoftwareDiagnosticsClassOnce sync.Once
)

func getMTRClusterSoftwareDiagnosticsClass() _MTRClusterSoftwareDiagnosticsClass {
	MTRClusterSoftwareDiagnosticsClassOnce.Do(func() {
		MTRClusterSoftwareDiagnosticsClass = _MTRClusterSoftwareDiagnosticsClass{objc.GetClass("MTRClusterSoftwareDiagnostics")}
	})
	return MTRClusterSoftwareDiagnosticsClass
}

type _MTRClusterSoftwareDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterSoftwareDiagnostics] class.
type IMTRClusterSoftwareDiagnostics interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterSoftwareDiagnostics
type MTRClusterSoftwareDiagnostics struct {
	MTRGenericCluster
}

// MTRClusterSoftwareDiagnosticsFrom constructs a [MTRClusterSoftwareDiagnostics] from an unsafe.Pointer.
func MTRClusterSoftwareDiagnosticsFrom(ptr unsafe.Pointer) MTRClusterSoftwareDiagnostics {
	return MTRClusterSoftwareDiagnostics{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterSoftwareDiagnosticsClass) Alloc() MTRClusterSoftwareDiagnostics {
	rv := objc.Send[MTRClusterSoftwareDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterSoftwareDiagnosticsClass) New() MTRClusterSoftwareDiagnostics {
	rv := objc.Send[MTRClusterSoftwareDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterSoftwareDiagnostics) Init() MTRClusterSoftwareDiagnostics {
	rv := objc.Send[MTRClusterSoftwareDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterSoftwareDiagnostics) Autorelease() MTRClusterSoftwareDiagnostics {
	rv := objc.Send[MTRClusterSoftwareDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterSoftwareDiagnostics creates a new MTRClusterSoftwareDiagnostics instance.
func NewMTRClusterSoftwareDiagnostics() MTRClusterSoftwareDiagnostics {
	return getMTRClusterSoftwareDiagnosticsClass().New()
}




