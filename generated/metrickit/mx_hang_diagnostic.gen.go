// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXHangDiagnostic] class.
var (
	MXHangDiagnosticClass     _MXHangDiagnosticClass
	MXHangDiagnosticClassOnce sync.Once
)

func getMXHangDiagnosticClass() _MXHangDiagnosticClass {
	MXHangDiagnosticClassOnce.Do(func() {
		MXHangDiagnosticClass = _MXHangDiagnosticClass{objc.GetClass("MXHangDiagnostic")}
	})
	return MXHangDiagnosticClass
}

type _MXHangDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [MXHangDiagnostic] class.
type IMXHangDiagnostic interface {
	IMXDiagnostic
	CallStackTree() MXCallStackTree
	HangDuration() unsafe.Pointer
}

// An object representing a diagnostic report for an app that is too busy to handle user input responsively.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHangDiagnostic
type MXHangDiagnostic struct {
	MXDiagnostic
}

// MXHangDiagnosticFrom constructs a [MXHangDiagnostic] from an unsafe.Pointer.
//
// An object representing a diagnostic report for an app that is too busy to handle user input responsively.
func MXHangDiagnosticFrom(ptr unsafe.Pointer) MXHangDiagnostic {
	return MXHangDiagnostic{
		MXDiagnostic: MXDiagnosticFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXHangDiagnosticClass) Alloc() MXHangDiagnostic {
	rv := objc.Send[MXHangDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXHangDiagnosticClass) New() MXHangDiagnostic {
	rv := objc.Send[MXHangDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXHangDiagnostic) Init() MXHangDiagnostic {
	rv := objc.Send[MXHangDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXHangDiagnostic) Autorelease() MXHangDiagnostic {
	rv := objc.Send[MXHangDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXHangDiagnostic creates a new MXHangDiagnostic instance.
func NewMXHangDiagnostic() MXHangDiagnostic {
	return getMXHangDiagnosticClass().New()
}


// The call stack for the app hang report.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHangDiagnostic/callStackTree
func (m_ MXHangDiagnostic) CallStackTree() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}

// The amount of time the app is busy and unable to respond to user interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHangDiagnostic/hangDuration
func (m_ MXHangDiagnostic) HangDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("hangDuration"))
	return rv
}



