// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayGather */


/* debug [class_header]: Header for MPSNDArrayGather */
// The class instance for the [NDArrayGather] class.
var (
	NDArrayGatherClass     _NDArrayGatherClass
	NDArrayGatherClassOnce sync.Once
)

func getNDArrayGatherClass() _NDArrayGatherClass {
	NDArrayGatherClassOnce.Do(func() {
		NDArrayGatherClass = _NDArrayGatherClass{objc.GetClass("MPSNDArrayGather")}
	})
	return NDArrayGatherClass
}

type _NDArrayGatherClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayGather */
// An interface definition for the [NDArrayGather] class.
type INDArrayGather interface {
	INDArrayBinaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayGather */
	// properties:
	Axis() objectivec.IObject
	SetAxis(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayGather */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayGather */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayGatherClass) Alloc() NDArrayGather {
	rv := objc.Send[NDArrayGather](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayGatherClass) New() NDArrayGather {
	rv := objc.Send[NDArrayGather](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGather) Init() NDArrayGather {
	rv := objc.Send[NDArrayGather](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGather) Autorelease() NDArrayGather {
	rv := objc.Send[NDArrayGather](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGather creates a new NDArrayGather instance.
func NewNDArrayGather() NDArrayGather {
	return getNDArrayGatherClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayGather */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGather
type NDArrayGather struct {
	NDArrayBinaryKernel
}

// NDArrayGatherFrom constructs a [NDArrayGather] from an unsafe.Pointer.
func NDArrayGatherFrom(ptr unsafe.Pointer) NDArrayGather {
	return NDArrayGather{
		NDArrayBinaryKernel: NDArrayBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayGather *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayGather */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayGather */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayGather */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayGather */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather/3152529-axis
func (n_ NDArrayGather) Axis() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("axis"))
	return rv
}/* debug [instance_properties/getter]: axis */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather/3152529-axis
func (n_ NDArrayGather) SetAxis(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAxis:"), value)
}/* debug [instance_properties/setter]: axis */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayGather */



