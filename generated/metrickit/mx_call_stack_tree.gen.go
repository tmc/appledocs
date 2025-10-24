// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXCallStackTree */


/* debug [class_header]: Header for MXCallStackTree */
// The class instance for the [MXCallStackTree] class.
var (
	MXCallStackTreeClass     _MXCallStackTreeClass
	MXCallStackTreeClassOnce sync.Once
)

func getMXCallStackTreeClass() _MXCallStackTreeClass {
	MXCallStackTreeClassOnce.Do(func() {
		MXCallStackTreeClass = _MXCallStackTreeClass{objc.GetClass("MXCallStackTree")}
	})
	return MXCallStackTreeClass
}

type _MXCallStackTreeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXCallStackTree */
// An interface definition for the [MXCallStackTree] class.
type IMXCallStackTree interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXCallStackTree */
	// properties:
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXCallStackTree */
	// methods:
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXCallStackTree */
// Alloc allocates a new instance without initialization.
func (mc _MXCallStackTreeClass) Alloc() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXCallStackTreeClass) New() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCallStackTree) Init() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCallStackTree) Autorelease() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCallStackTree creates a new MXCallStackTree instance.
func NewMXCallStackTree() MXCallStackTree {
	return getMXCallStackTreeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXCallStackTree */
// An object representing the call stack for an exception.


// An object representing the call stack for an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCallStackTree
type MXCallStackTree struct {
	objectivec.Object
}

// MXCallStackTreeFrom constructs a [MXCallStackTree] from an unsafe.Pointer.
//
// An object representing the call stack for an exception.
func MXCallStackTreeFrom(ptr unsafe.Pointer) MXCallStackTree {
	return MXCallStackTree{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXCallStackTree *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXCallStackTree */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXCallStackTree */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXCallStackTree */

// Returns the contents of the stack tree in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCallStackTree/jsonRepresentation()
func (m_ MXCallStackTree) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXCallStackTree */

// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXCallStackTree) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXCallStackTree */



