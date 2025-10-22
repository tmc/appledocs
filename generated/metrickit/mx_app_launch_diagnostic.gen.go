// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MXAppLaunchDiagnostic] class.
var (
	MXAppLaunchDiagnosticClass     _MXAppLaunchDiagnosticClass
	MXAppLaunchDiagnosticClassOnce sync.Once
)

func getMXAppLaunchDiagnosticClass() _MXAppLaunchDiagnosticClass {
	MXAppLaunchDiagnosticClassOnce.Do(func() {
		MXAppLaunchDiagnosticClass = _MXAppLaunchDiagnosticClass{objc.GetClass("MXAppLaunchDiagnostic")}
	})
	return MXAppLaunchDiagnosticClass
}

type _MXAppLaunchDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [MXAppLaunchDiagnostic] class.
type IMXAppLaunchDiagnostic interface {
	IMXDiagnostic
	CallStackTree() MXCallStackTree
	LaunchDuration() unsafe.Pointer
}

// A diagnostic subclass that encapsulates app launch diagnostic reports.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchDiagnostic
type MXAppLaunchDiagnostic struct {
	MXDiagnostic
}

// MXAppLaunchDiagnosticFrom constructs a [MXAppLaunchDiagnostic] from an unsafe.Pointer.
//
// A diagnostic subclass that encapsulates app launch diagnostic reports.
func MXAppLaunchDiagnosticFrom(ptr unsafe.Pointer) MXAppLaunchDiagnostic {
	return MXAppLaunchDiagnostic{
		MXDiagnostic: MXDiagnosticFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXAppLaunchDiagnosticClass) Alloc() MXAppLaunchDiagnostic {
	rv := objc.Send[MXAppLaunchDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXAppLaunchDiagnosticClass) New() MXAppLaunchDiagnostic {
	rv := objc.Send[MXAppLaunchDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXAppLaunchDiagnostic) Init() MXAppLaunchDiagnostic {
	rv := objc.Send[MXAppLaunchDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXAppLaunchDiagnostic) Autorelease() MXAppLaunchDiagnostic {
	rv := objc.Send[MXAppLaunchDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXAppLaunchDiagnostic creates a new MXAppLaunchDiagnostic instance.
func NewMXAppLaunchDiagnostic() MXAppLaunchDiagnostic {
	return getMXAppLaunchDiagnosticClass().New()
}


// The call stack tree associated with the app launch.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchDiagnostic/callStackTree
func (m_ MXAppLaunchDiagnostic) CallStackTree() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}

// The total app launch duration.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXAppLaunchDiagnostic/launchDuration
func (m_ MXAppLaunchDiagnostic) LaunchDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("launchDuration"))
	return rv
}



