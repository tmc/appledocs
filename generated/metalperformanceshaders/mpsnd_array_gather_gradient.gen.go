// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNDArrayGatherGradient */


/* debug [class_header]: Header for MPSNDArrayGatherGradient */
// The class instance for the [NDArrayGatherGradient] class.
var (
	NDArrayGatherGradientClass     _NDArrayGatherGradientClass
	NDArrayGatherGradientClassOnce sync.Once
)

func getNDArrayGatherGradientClass() _NDArrayGatherGradientClass {
	NDArrayGatherGradientClassOnce.Do(func() {
		NDArrayGatherGradientClass = _NDArrayGatherGradientClass{objc.GetClass("MPSNDArrayGatherGradient")}
	})
	return NDArrayGatherGradientClass
}

type _NDArrayGatherGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayGatherGradient */
// An interface definition for the [NDArrayGatherGradient] class.
type INDArrayGatherGradient interface {
	INDArrayBinaryPrimaryGradientKernel
	
/* debug [class_interface_properties]: Properties for NDArrayGatherGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayGatherGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayGatherGradient */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayGatherGradientClass) Alloc() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayGatherGradientClass) New() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGatherGradient) Init() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGatherGradient) Autorelease() NDArrayGatherGradient {
	rv := objc.Send[NDArrayGatherGradient](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGatherGradient creates a new NDArrayGatherGradient instance.
func NewNDArrayGatherGradient() NDArrayGatherGradient {
	return getNDArrayGatherGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayGatherGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradient
type NDArrayGatherGradient struct {
	NDArrayBinaryPrimaryGradientKernel
}

// NDArrayGatherGradientFrom constructs a [NDArrayGatherGradient] from an unsafe.Pointer.
func NDArrayGatherGradientFrom(ptr unsafe.Pointer) NDArrayGatherGradient {
	return NDArrayGatherGradient{
		NDArrayBinaryPrimaryGradientKernel: NDArrayBinaryPrimaryGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayGatherGradient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayGatherGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayGatherGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayGatherGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayGatherGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayGatherGradient */



