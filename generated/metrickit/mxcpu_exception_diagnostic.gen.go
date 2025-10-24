// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXCPUExceptionDiagnostic */


/* debug [class_header]: Header for MXCPUExceptionDiagnostic */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXCPUExceptionDiagnostic */
// An interface definition for the [MXCPUExceptionDiagnostic] class.
type IMXCPUExceptionDiagnostic interface {
	IMXDiagnostic
	
/* debug [class_interface_properties]: Properties for MXCPUExceptionDiagnostic */
	// properties:
	CallStackTree() IMXCallStackTree
	TotalCPUTime() unsafe.Pointer
	TotalSampledTime() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXCPUExceptionDiagnostic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXCPUExceptionDiagnostic */
// Alloc allocates a new instance without initialization.
func (mc _MXCPUExceptionDiagnosticClass) Alloc() MXCPUExceptionDiagnostic {
	rv := objc.Send[MXCPUExceptionDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXCPUExceptionDiagnostic */
// An object representing a diagnostic report for a fatal or nonfatal CPU exception.
//
// A CPU exception occurs when your app uses an excessive amount of CPU time over a short period.


// An object representing a diagnostic report for a fatal or nonfatal CPU exception.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXCPUExceptionDiagnostic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXCPUExceptionDiagnostic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXCPUExceptionDiagnostic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXCPUExceptionDiagnostic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXCPUExceptionDiagnostic */

// The app call stack associated with the CPU exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic/callStackTree
func (m_ MXCPUExceptionDiagnostic) CallStackTree() IMXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}/* debug [instance_properties/getter]: callStackTree */


// The total CPU time used during the exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic/totalCPUTime
func (m_ MXCPUExceptionDiagnostic) TotalCPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalCPUTime"))
	return rv
}/* debug [instance_properties/getter]: totalCPUTime */


// The total time the app was sampled during the exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCPUExceptionDiagnostic/totalSampledTime
func (m_ MXCPUExceptionDiagnostic) TotalSampledTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalSampledTime"))
	return rv
}/* debug [instance_properties/getter]: totalSampledTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXCPUExceptionDiagnostic */



