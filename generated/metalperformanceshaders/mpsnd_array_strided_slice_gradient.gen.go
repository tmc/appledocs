// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNDArrayStridedSliceGradient */


/* debug [class_header]: Header for MPSNDArrayStridedSliceGradient */
// The class instance for the [NDArrayStridedSliceGradient] class.
var (
	NDArrayStridedSliceGradientClass     _NDArrayStridedSliceGradientClass
	NDArrayStridedSliceGradientClassOnce sync.Once
)

func getNDArrayStridedSliceGradientClass() _NDArrayStridedSliceGradientClass {
	NDArrayStridedSliceGradientClassOnce.Do(func() {
		NDArrayStridedSliceGradientClass = _NDArrayStridedSliceGradientClass{objc.GetClass("MPSNDArrayStridedSliceGradient")}
	})
	return NDArrayStridedSliceGradientClass
}

type _NDArrayStridedSliceGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayStridedSliceGradient */
// An interface definition for the [NDArrayStridedSliceGradient] class.
type INDArrayStridedSliceGradient interface {
	INDArrayUnaryGradientKernel
	
/* debug [class_interface_properties]: Properties for NDArrayStridedSliceGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayStridedSliceGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayStridedSliceGradient */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayStridedSliceGradientClass) Alloc() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayStridedSliceGradientClass) New() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayStridedSliceGradient) Init() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayStridedSliceGradient) Autorelease() NDArrayStridedSliceGradient {
	rv := objc.Send[NDArrayStridedSliceGradient](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayStridedSliceGradient creates a new NDArrayStridedSliceGradient instance.
func NewNDArrayStridedSliceGradient() NDArrayStridedSliceGradient {
	return getNDArrayStridedSliceGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayStridedSliceGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayStridedSliceGradient
type NDArrayStridedSliceGradient struct {
	NDArrayUnaryGradientKernel
}

// NDArrayStridedSliceGradientFrom constructs a [NDArrayStridedSliceGradient] from an unsafe.Pointer.
func NDArrayStridedSliceGradientFrom(ptr unsafe.Pointer) NDArrayStridedSliceGradient {
	return NDArrayStridedSliceGradient{
		NDArrayUnaryGradientKernel: NDArrayUnaryGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayStridedSliceGradient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayStridedSliceGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayStridedSliceGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayStridedSliceGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayStridedSliceGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayStridedSliceGradient */



