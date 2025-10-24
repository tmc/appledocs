// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MXCrashDiagnostic] class.
var (
	MXCrashDiagnosticClass     _MXCrashDiagnosticClass
	MXCrashDiagnosticClassOnce sync.Once
)

func getMXCrashDiagnosticClass() _MXCrashDiagnosticClass {
	MXCrashDiagnosticClassOnce.Do(func() {
		MXCrashDiagnosticClass = _MXCrashDiagnosticClass{objc.GetClass("MXCrashDiagnostic")}
	})
	return MXCrashDiagnosticClass
}

type _MXCrashDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [MXCrashDiagnostic] class.
type IMXCrashDiagnostic interface {
	IMXDiagnostic
	// properties:
	CallStackTree() IMXCallStackTree
	ExceptionCode() objc.IObject /* cross-framework: NSNumber */
	ExceptionReason() IMXCrashDiagnosticObjectiveCExceptionReason
	ExceptionType() objc.IObject /* cross-framework: NSNumber */
	Signal() objc.IObject /* cross-framework: NSNumber */
	TerminationReason() objc.IObject /* cross-framework: NSString */
	VirtualMemoryRegionInfo() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An object representing a diagnostic report for an app crash.


// An object representing a diagnostic report for an app crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic
type MXCrashDiagnostic struct {
	MXDiagnostic
}

// MXCrashDiagnosticFrom constructs a [MXCrashDiagnostic] from an unsafe.Pointer.
//
// An object representing a diagnostic report for an app crash.
func MXCrashDiagnosticFrom(ptr unsafe.Pointer) MXCrashDiagnostic {
	return MXCrashDiagnostic{
		MXDiagnostic: MXDiagnosticFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXCrashDiagnosticClass) Alloc() MXCrashDiagnostic {
	rv := objc.Send[MXCrashDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXCrashDiagnosticClass) New() MXCrashDiagnostic {
	rv := objc.Send[MXCrashDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCrashDiagnostic) Init() MXCrashDiagnostic {
	rv := objc.Send[MXCrashDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCrashDiagnostic) Autorelease() MXCrashDiagnostic {
	rv := objc.Send[MXCrashDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCrashDiagnostic creates a new MXCrashDiagnostic instance.
func NewMXCrashDiagnostic() MXCrashDiagnostic {
	return getMXCrashDiagnosticClass().New()
}



// The call stack for the crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/callStackTree
func (m_ MXCrashDiagnostic) CallStackTree() IMXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}


// The encoded processor-specific information for the crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/exceptionCode
func (m_ MXCrashDiagnostic) ExceptionCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("exceptionCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/exceptionReason
func (m_ MXCrashDiagnostic) ExceptionReason() IMXCrashDiagnosticObjectiveCExceptionReason {
	rv := objc.Send[MXCrashDiagnosticObjectiveCExceptionReason](m_.ID, objc.Sel("exceptionReason"))
	return rv
}


// The Mach exception type of the crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/exceptionType
func (m_ MXCrashDiagnostic) ExceptionType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("exceptionType"))
	return rv
}


// The signal associated with the crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/signal
func (m_ MXCrashDiagnostic) Signal() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("signal"))
	return rv
}


// The reason the app was terminated as a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/terminationReason
func (m_ MXCrashDiagnostic) TerminationReason() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("terminationReason"))
	return rv
}


// Information about the region of memory an app accessed incorrectly, resulting in a bad-access crash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCrashDiagnostic/virtualMemoryRegionInfo
func (m_ MXCrashDiagnostic) VirtualMemoryRegionInfo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("virtualMemoryRegionInfo"))
	return rv
}



