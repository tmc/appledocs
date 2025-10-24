// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXHangDiagnostic */


/* debug [class_header]: Header for MXHangDiagnostic */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXHangDiagnostic */
// An interface definition for the [MXHangDiagnostic] class.
type IMXHangDiagnostic interface {
	IMXDiagnostic
	
/* debug [class_interface_properties]: Properties for MXHangDiagnostic */
	// properties:
	CallStackTree() IMXCallStackTree
	HangDuration() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXHangDiagnostic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXHangDiagnostic */
// Alloc allocates a new instance without initialization.
func (mc _MXHangDiagnosticClass) Alloc() MXHangDiagnostic {
	rv := objc.Send[MXHangDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXHangDiagnostic */
// An object representing a diagnostic report for an app that is too busy to handle user input responsively.


// An object representing a diagnostic report for an app that is too busy to handle user input responsively.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXHangDiagnostic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXHangDiagnostic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXHangDiagnostic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXHangDiagnostic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXHangDiagnostic */

// The call stack for the app hang report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHangDiagnostic/callStackTree
func (m_ MXHangDiagnostic) CallStackTree() IMXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}/* debug [instance_properties/getter]: callStackTree */


// The amount of time the app is busy and unable to respond to user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHangDiagnostic/hangDuration
func (m_ MXHangDiagnostic) HangDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("hangDuration"))
	return rv
}/* debug [instance_properties/getter]: hangDuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXHangDiagnostic */



