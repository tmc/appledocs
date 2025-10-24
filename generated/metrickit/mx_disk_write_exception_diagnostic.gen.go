// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXDiskWriteExceptionDiagnostic */


/* debug [class_header]: Header for MXDiskWriteExceptionDiagnostic */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXDiskWriteExceptionDiagnostic */
// An interface definition for the [MXDiskWriteExceptionDiagnostic] class.
type IMXDiskWriteExceptionDiagnostic interface {
	IMXDiagnostic
	
/* debug [class_interface_properties]: Properties for MXDiskWriteExceptionDiagnostic */
	// properties:
	CallStackTree() IMXCallStackTree
	TotalWritesCaused() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXDiskWriteExceptionDiagnostic */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXDiskWriteExceptionDiagnostic */
// Alloc allocates a new instance without initialization.
func (mc _MXDiskWriteExceptionDiagnosticClass) Alloc() MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[MXDiskWriteExceptionDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXDiskWriteExceptionDiagnostic */
// An object representing a diagnostic report for a disk write exception.
//
// A disk write exception occurs when the app writes an excessive amount of data to the disk.


// An object representing a diagnostic report for a disk write exception.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXDiskWriteExceptionDiagnostic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXDiskWriteExceptionDiagnostic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXDiskWriteExceptionDiagnostic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXDiskWriteExceptionDiagnostic */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXDiskWriteExceptionDiagnostic */

// The call stack for the disk write exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskWriteExceptionDiagnostic/callStackTree
func (m_ MXDiskWriteExceptionDiagnostic) CallStackTree() IMXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("callStackTree"))
	return rv
}/* debug [instance_properties/getter]: callStackTree */


// The total amount of data written to disk or other long-term storage during the disk write exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskWriteExceptionDiagnostic/totalWritesCaused
func (m_ MXDiskWriteExceptionDiagnostic) TotalWritesCaused() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalWritesCaused"))
	return rv
}/* debug [instance_properties/getter]: totalWritesCaused */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXDiskWriteExceptionDiagnostic */



