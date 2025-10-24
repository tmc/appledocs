// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLComputePipelineReflection */


/* debug [class_header]: Header for MTLComputePipelineReflection */
// The class instance for the [ComputePipelineReflection] class.
var (
	ComputePipelineReflectionClass     _ComputePipelineReflectionClass
	ComputePipelineReflectionClassOnce sync.Once
)

func getComputePipelineReflectionClass() _ComputePipelineReflectionClass {
	ComputePipelineReflectionClassOnce.Do(func() {
		ComputePipelineReflectionClass = _ComputePipelineReflectionClass{objc.GetClass("MTLComputePipelineReflection")}
	})
	return ComputePipelineReflectionClass
}

type _ComputePipelineReflectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComputePipelineReflection */
// An interface definition for the [ComputePipelineReflection] class.
type IComputePipelineReflection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ComputePipelineReflection */
	// properties:
	Arguments() []Argument
	Bindings() []objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComputePipelineReflection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComputePipelineReflection */
// Alloc allocates a new instance without initialization.
func (cc _ComputePipelineReflectionClass) Alloc() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComputePipelineReflectionClass) New() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePipelineReflection) Init() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePipelineReflection) Autorelease() ComputePipelineReflection {
	rv := objc.Send[ComputePipelineReflection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePipelineReflection creates a new ComputePipelineReflection instance.
func NewComputePipelineReflection() ComputePipelineReflection {
	return getComputePipelineReflectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComputePipelineReflection */
// Information about the arguments of a compute function.
//
// An object provides access to the arguments of the compute function used in an object. An object can be created along with an object. Don’t create an object directly. Instead, call either the or method of to create both an object and an object. objects can use a significant amount of memory; release any strong references to them after you finish creating pipeline objects.


// Information about the arguments of a compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection
type ComputePipelineReflection struct {
	objectivec.Object
}

// ComputePipelineReflectionFrom constructs a [ComputePipelineReflection] from an unsafe.Pointer.
//
// Information about the arguments of a compute function.
func ComputePipelineReflectionFrom(ptr unsafe.Pointer) ComputePipelineReflection {
	return ComputePipelineReflection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComputePipelineReflection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComputePipelineReflection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComputePipelineReflection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComputePipelineReflection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComputePipelineReflection */

// An array of instances that describe the arguments of a compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection/arguments
func (c_ ComputePipelineReflection) Arguments() []Argument {
	rv := objc.Send[[]Argument](c_.ID, objc.Sel("arguments"))
	return rv
}/* debug [instance_properties/getter]: arguments */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePipelineReflection/bindings
func (c_ ComputePipelineReflection) Bindings() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("bindings"))
	return rv
}/* debug [instance_properties/getter]: bindings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLComputePipelineReflection */



