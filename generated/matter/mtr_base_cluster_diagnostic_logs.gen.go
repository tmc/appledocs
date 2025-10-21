// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterDiagnosticLogs] class.
var (
	MTRBaseClusterDiagnosticLogsClass     _MTRBaseClusterDiagnosticLogsClass
	MTRBaseClusterDiagnosticLogsClassOnce sync.Once
)

func getMTRBaseClusterDiagnosticLogsClass() _MTRBaseClusterDiagnosticLogsClass {
	MTRBaseClusterDiagnosticLogsClassOnce.Do(func() {
		MTRBaseClusterDiagnosticLogsClass = _MTRBaseClusterDiagnosticLogsClass{objc.GetClass("MTRBaseClusterDiagnosticLogs")}
	})
	return MTRBaseClusterDiagnosticLogsClass
}

type _MTRBaseClusterDiagnosticLogsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterDiagnosticLogs] class.
type IMTRBaseClusterDiagnosticLogs interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDiagnosticLogs
type MTRBaseClusterDiagnosticLogs struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDiagnosticLogsFrom constructs a [MTRBaseClusterDiagnosticLogs] from an unsafe.Pointer.
func MTRBaseClusterDiagnosticLogsFrom(ptr unsafe.Pointer) MTRBaseClusterDiagnosticLogs {
	return MTRBaseClusterDiagnosticLogs{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDiagnosticLogsClass) Alloc() MTRBaseClusterDiagnosticLogs {
	rv := objc.Send[MTRBaseClusterDiagnosticLogs](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterDiagnosticLogsClass) New() MTRBaseClusterDiagnosticLogs {
	rv := objc.Send[MTRBaseClusterDiagnosticLogs](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDiagnosticLogs) Init() MTRBaseClusterDiagnosticLogs {
	rv := objc.Send[MTRBaseClusterDiagnosticLogs](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDiagnosticLogs) Autorelease() MTRBaseClusterDiagnosticLogs {
	rv := objc.Send[MTRBaseClusterDiagnosticLogs](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDiagnosticLogs creates a new MTRBaseClusterDiagnosticLogs instance.
func NewMTRBaseClusterDiagnosticLogs() MTRBaseClusterDiagnosticLogs {
	return getMTRBaseClusterDiagnosticLogsClass().New()
}




