// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXDiagnostic */


/* debug [class_header]: Header for MXDiagnostic */
// The class instance for the [MXDiagnostic] class.
var (
	MXDiagnosticClass     _MXDiagnosticClass
	MXDiagnosticClassOnce sync.Once
)

func getMXDiagnosticClass() _MXDiagnosticClass {
	MXDiagnosticClassOnce.Do(func() {
		MXDiagnosticClass = _MXDiagnosticClass{objc.GetClass("MXDiagnostic")}
	})
	return MXDiagnosticClass
}

type _MXDiagnosticClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXDiagnostic */
// An interface definition for the [MXDiagnostic] class.
type IMXDiagnostic interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXDiagnostic */
	// properties:
	ApplicationVersion() objc.IObject /* cross-framework: NSString */
	MetaData() IMXMetaData
	SignpostData() []MXSignpostRecord
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXDiagnostic */
	// methods:
	DictionaryRepresentation() foundation.Dictionary
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXDiagnostic */
// Alloc allocates a new instance without initialization.
func (mc _MXDiagnosticClass) Alloc() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXDiagnosticClass) New() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiagnostic) Init() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiagnostic) Autorelease() MXDiagnostic {
	rv := objc.Send[MXDiagnostic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiagnostic creates a new MXDiagnostic instance.
func NewMXDiagnostic() MXDiagnostic {
	return getMXDiagnosticClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXDiagnostic */
// An abstract data class for a diagnostic.


// An abstract data class for a diagnostic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic
type MXDiagnostic struct {
	objectivec.Object
}

// MXDiagnosticFrom constructs a [MXDiagnostic] from an unsafe.Pointer.
//
// An abstract data class for a diagnostic.
func MXDiagnosticFrom(ptr unsafe.Pointer) MXDiagnostic {
	return MXDiagnostic{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXDiagnostic *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXDiagnostic */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXDiagnostic */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXDiagnostic */

// Returns the contents of a diagnostic as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/dictionaryRepresentation()
func (m_ MXDiagnostic) DictionaryRepresentation() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// Returns the contents of the diagnostic in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/jsonRepresentation()
func (m_ MXDiagnostic) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXDiagnostic */

// The value of the bundle version key, short form, in the app’s property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/applicationVersion
func (m_ MXDiagnostic) ApplicationVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationVersion"))
	return rv
}/* debug [instance_properties/getter]: applicationVersion */


// A set of system-level information for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/metaData
func (m_ MXDiagnostic) MetaData() IMXMetaData {
	rv := objc.Send[MXMetaData](m_.ID, objc.Sel("metaData"))
	return rv
}/* debug [instance_properties/getter]: metaData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnostic/signpostData
func (m_ MXDiagnostic) SignpostData() []MXSignpostRecord {
	rv := objc.Send[[]MXSignpostRecord](m_.ID, objc.Sel("signpostData"))
	return rv
}/* debug [instance_properties/getter]: signpostData */


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXDiagnostic) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXDiagnostic */



