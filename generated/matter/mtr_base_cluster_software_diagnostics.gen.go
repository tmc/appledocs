// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterSoftwareDiagnostics] class.
var (
	MTRBaseClusterSoftwareDiagnosticsClass     _MTRBaseClusterSoftwareDiagnosticsClass
	MTRBaseClusterSoftwareDiagnosticsClassOnce sync.Once
)

func getMTRBaseClusterSoftwareDiagnosticsClass() _MTRBaseClusterSoftwareDiagnosticsClass {
	MTRBaseClusterSoftwareDiagnosticsClassOnce.Do(func() {
		MTRBaseClusterSoftwareDiagnosticsClass = _MTRBaseClusterSoftwareDiagnosticsClass{objc.GetClass("MTRBaseClusterSoftwareDiagnostics")}
	})
	return MTRBaseClusterSoftwareDiagnosticsClass
}

type _MTRBaseClusterSoftwareDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterSoftwareDiagnostics] class.
type IMTRBaseClusterSoftwareDiagnostics interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics
type MTRBaseClusterSoftwareDiagnostics struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterSoftwareDiagnosticsFrom constructs a [MTRBaseClusterSoftwareDiagnostics] from an unsafe.Pointer.
func MTRBaseClusterSoftwareDiagnosticsFrom(ptr unsafe.Pointer) MTRBaseClusterSoftwareDiagnostics {
	return MTRBaseClusterSoftwareDiagnostics{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) Alloc() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) New() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterSoftwareDiagnostics) Init() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterSoftwareDiagnostics) Autorelease() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterSoftwareDiagnostics creates a new MTRBaseClusterSoftwareDiagnostics instance.
func NewMTRBaseClusterSoftwareDiagnostics() MTRBaseClusterSoftwareDiagnostics {
	return getMTRBaseClusterSoftwareDiagnosticsClass().New()
}




