// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXAppLaunchDiagnostic */


/* debug [class_header]: Header for MXAppLaunchDiagnostic */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXAppLaunchDiagnostic */
// An interface definition for the [MXAppLaunchDiagnostic] class.
type IMXAppLaunchDiagnostic interface {
	IMXDiagnostic
	
/* debug [class_interface_properties]: Properties for MXAppLaunchDiagnostic */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXAppLaunchDiagnostic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXAppLaunchDiagnostic */
// Alloc allocates a new instance without initialization.
func (mc _MXAppLaunchDiagnosticClass) Alloc() MXAppLaunchDiagnostic {
	rv := objc.Send[MXAppLaunchDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXAppLaunchDiagnostic */
// A diagnostic subclass that encapsulates app launch diagnostic reports.


// A diagnostic subclass that encapsulates app launch diagnostic reports.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXAppLaunchDiagnostic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXAppLaunchDiagnostic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXAppLaunchDiagnostic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXAppLaunchDiagnostic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXAppLaunchDiagnostic */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXAppLaunchDiagnostic */


