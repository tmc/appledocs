// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXCPUExceptionDiagnostic] class.
var (
	MXCPUExceptionDiagnosticClass     _MXCPUExceptionDiagnosticClass
	MXCPUExceptionDiagnosticClassOnce sync.Once
)

func getMXCPUExceptionDiagnosticClass() _MXCPUExceptionDiagnosticClass {
	MXCPUExceptionDiagnosticClassOnce.Do(func() {
		MXCPUExceptionDiagnosticClass = _MXCPUExceptionDiagnosticClass{objc.GetClass("MXCPUExceptionDiagnostic")}
	})
	return MXCPUExceptionDiagnosticClass
}

type _MXCPUExceptionDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [MXCPUExceptionDiagnostic] class.
type IMXCPUExceptionDiagnostic interface {
	IMXDiagnostic
}

// An object representing a diagnostic report for a fatal or nonfatal CPU exception.
//
// A CPU exception occurs when your app uses an excessive amount of CPU time over a short period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic
type MXCPUExceptionDiagnostic struct {
	MXDiagnostic
}

// MXCPUExceptionDiagnosticFrom constructs a [MXCPUExceptionDiagnostic] from an unsafe.Pointer.
//
// An object representing a diagnostic report for a fatal or nonfatal CPU exception.
func MXCPUExceptionDiagnosticFrom(ptr unsafe.Pointer) MXCPUExceptionDiagnostic {
	return MXCPUExceptionDiagnostic{
		MXDiagnostic: MXDiagnosticFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXCPUExceptionDiagnosticClass) Alloc() MXCPUExceptionDiagnostic {
	rv := objc.Send[MXCPUExceptionDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXCPUExceptionDiagnosticClass) New() MXCPUExceptionDiagnostic {
	rv := objc.Send[MXCPUExceptionDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCPUExceptionDiagnostic) Init() MXCPUExceptionDiagnostic {
	rv := objc.Send[MXCPUExceptionDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCPUExceptionDiagnostic) Autorelease() MXCPUExceptionDiagnostic {
	rv := objc.Send[MXCPUExceptionDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCPUExceptionDiagnostic creates a new MXCPUExceptionDiagnostic instance.
func NewMXCPUExceptionDiagnostic() MXCPUExceptionDiagnostic {
	return getMXCPUExceptionDiagnosticClass().New()
}


// The app call stack associated with the CPU exception.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic/callStackTree
func (m_ MXCPUExceptionDiagnostic) CallStackTree() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}

// The total CPU time used during the exception.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic/totalCPUTime
func (m_ MXCPUExceptionDiagnostic) TotalCPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalCPUTime"))
	return rv
}

// The total time the app was sampled during the exception.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic/totalSampledTime
func (m_ MXCPUExceptionDiagnostic) TotalSampledTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalSampledTime"))
	return rv
}



