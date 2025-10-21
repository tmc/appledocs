// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXDiskWriteExceptionDiagnostic] class.
var (
	MXDiskWriteExceptionDiagnosticClass     _MXDiskWriteExceptionDiagnosticClass
	MXDiskWriteExceptionDiagnosticClassOnce sync.Once
)

func getMXDiskWriteExceptionDiagnosticClass() _MXDiskWriteExceptionDiagnosticClass {
	MXDiskWriteExceptionDiagnosticClassOnce.Do(func() {
		MXDiskWriteExceptionDiagnosticClass = _MXDiskWriteExceptionDiagnosticClass{objc.GetClass("MXDiskWriteExceptionDiagnostic")}
	})
	return MXDiskWriteExceptionDiagnosticClass
}

type _MXDiskWriteExceptionDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [MXDiskWriteExceptionDiagnostic] class.
type IMXDiskWriteExceptionDiagnostic interface {
	IMXDiagnostic
}

// An object representing a diagnostic report for a disk write exception.
//
// A disk write exception occurs when the app writes an excessive amount of data to the disk.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskWriteExceptionDiagnostic
type MXDiskWriteExceptionDiagnostic struct {
	MXDiagnostic
}

// MXDiskWriteExceptionDiagnosticFrom constructs a [MXDiskWriteExceptionDiagnostic] from an unsafe.Pointer.
//
// An object representing a diagnostic report for a disk write exception.
func MXDiskWriteExceptionDiagnosticFrom(ptr unsafe.Pointer) MXDiskWriteExceptionDiagnostic {
	return MXDiskWriteExceptionDiagnostic{
		MXDiagnostic: MXDiagnosticFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXDiskWriteExceptionDiagnosticClass) Alloc() MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[MXDiskWriteExceptionDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXDiskWriteExceptionDiagnosticClass) New() MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[MXDiskWriteExceptionDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiskWriteExceptionDiagnostic) Init() MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[MXDiskWriteExceptionDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiskWriteExceptionDiagnostic) Autorelease() MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[MXDiskWriteExceptionDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiskWriteExceptionDiagnostic creates a new MXDiskWriteExceptionDiagnostic instance.
func NewMXDiskWriteExceptionDiagnostic() MXDiskWriteExceptionDiagnostic {
	return getMXDiskWriteExceptionDiagnosticClass().New()
}


// The call stack for the disk write exception.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskWriteExceptionDiagnostic/callStackTree
func (m_ MXDiskWriteExceptionDiagnostic) CallStackTree() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}

// The total amount of data written to disk or other long-term storage during the disk write exception.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskWriteExceptionDiagnostic/totalWritesCaused
func (m_ MXDiskWriteExceptionDiagnostic) TotalWritesCaused() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalWritesCaused"))
	return rv
}



