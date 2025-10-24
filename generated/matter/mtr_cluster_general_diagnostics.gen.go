// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterGeneralDiagnostics] class.
var (
	MTRClusterGeneralDiagnosticsClass     _MTRClusterGeneralDiagnosticsClass
	MTRClusterGeneralDiagnosticsClassOnce sync.Once
)

func getMTRClusterGeneralDiagnosticsClass() _MTRClusterGeneralDiagnosticsClass {
	MTRClusterGeneralDiagnosticsClassOnce.Do(func() {
		MTRClusterGeneralDiagnosticsClass = _MTRClusterGeneralDiagnosticsClass{objc.GetClass("MTRClusterGeneralDiagnostics")}
	})
	return MTRClusterGeneralDiagnosticsClass
}

type _MTRClusterGeneralDiagnosticsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterGeneralDiagnostics] class.
type IMTRClusterGeneralDiagnostics interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterGeneralDiagnostics
type MTRClusterGeneralDiagnostics struct {
	MTRGenericCluster
}

// MTRClusterGeneralDiagnosticsFrom constructs a [MTRClusterGeneralDiagnostics] from an unsafe.Pointer.
func MTRClusterGeneralDiagnosticsFrom(ptr unsafe.Pointer) MTRClusterGeneralDiagnostics {
	return MTRClusterGeneralDiagnostics{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterGeneralDiagnosticsClass) Alloc() MTRClusterGeneralDiagnostics {
	rv := objc.Send[MTRClusterGeneralDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterGeneralDiagnosticsClass) New() MTRClusterGeneralDiagnostics {
	rv := objc.Send[MTRClusterGeneralDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterGeneralDiagnostics) Init() MTRClusterGeneralDiagnostics {
	rv := objc.Send[MTRClusterGeneralDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterGeneralDiagnostics) Autorelease() MTRClusterGeneralDiagnostics {
	rv := objc.Send[MTRClusterGeneralDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterGeneralDiagnostics creates a new MTRClusterGeneralDiagnostics instance.
func NewMTRClusterGeneralDiagnostics() MTRClusterGeneralDiagnostics {
	return getMTRClusterGeneralDiagnosticsClass().New()
}




