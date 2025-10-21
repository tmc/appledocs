// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterDiagnosticLogs] class.
var (
	MTRClusterDiagnosticLogsClass     _MTRClusterDiagnosticLogsClass
	MTRClusterDiagnosticLogsClassOnce sync.Once
)

func getMTRClusterDiagnosticLogsClass() _MTRClusterDiagnosticLogsClass {
	MTRClusterDiagnosticLogsClassOnce.Do(func() {
		MTRClusterDiagnosticLogsClass = _MTRClusterDiagnosticLogsClass{objc.GetClass("MTRClusterDiagnosticLogs")}
	})
	return MTRClusterDiagnosticLogsClass
}

type _MTRClusterDiagnosticLogsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDiagnosticLogs] class.
type IMTRClusterDiagnosticLogs interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDiagnosticLogs
type MTRClusterDiagnosticLogs struct {
	MTRGenericCluster
}

// MTRClusterDiagnosticLogsFrom constructs a [MTRClusterDiagnosticLogs] from an unsafe.Pointer.
func MTRClusterDiagnosticLogsFrom(ptr unsafe.Pointer) MTRClusterDiagnosticLogs {
	return MTRClusterDiagnosticLogs{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDiagnosticLogsClass) Alloc() MTRClusterDiagnosticLogs {
	rv := objc.Send[MTRClusterDiagnosticLogs](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDiagnosticLogsClass) New() MTRClusterDiagnosticLogs {
	rv := objc.Send[MTRClusterDiagnosticLogs](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDiagnosticLogs) Init() MTRClusterDiagnosticLogs {
	rv := objc.Send[MTRClusterDiagnosticLogs](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDiagnosticLogs) Autorelease() MTRClusterDiagnosticLogs {
	rv := objc.Send[MTRClusterDiagnosticLogs](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDiagnosticLogs creates a new MTRClusterDiagnosticLogs instance.
func NewMTRClusterDiagnosticLogs() MTRClusterDiagnosticLogs {
	return getMTRClusterDiagnosticLogsClass().New()
}




