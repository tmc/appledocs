// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayLUTDequantize */


/* debug [class_header]: Header for MPSNDArrayLUTDequantize */
// The class instance for the [NDArrayLUTDequantize] class.
var (
	NDArrayLUTDequantizeClass     _NDArrayLUTDequantizeClass
	NDArrayLUTDequantizeClassOnce sync.Once
)

func getNDArrayLUTDequantizeClass() _NDArrayLUTDequantizeClass {
	NDArrayLUTDequantizeClassOnce.Do(func() {
		NDArrayLUTDequantizeClass = _NDArrayLUTDequantizeClass{objc.GetClass("MPSNDArrayLUTDequantize")}
	})
	return NDArrayLUTDequantizeClass
}

type _NDArrayLUTDequantizeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayLUTDequantize */
// An interface definition for the [NDArrayLUTDequantize] class.
type INDArrayLUTDequantize interface {
	INDArrayMultiaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayLUTDequantize */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayLUTDequantize */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayLUTDequantize */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayLUTDequantizeClass) Alloc() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayLUTDequantizeClass) New() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayLUTDequantize) Init() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayLUTDequantize) Autorelease() NDArrayLUTDequantize {
	rv := objc.Send[NDArrayLUTDequantize](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayLUTDequantize creates a new NDArrayLUTDequantize instance.
func NewNDArrayLUTDequantize() NDArrayLUTDequantize {
	return getNDArrayLUTDequantizeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayLUTDequantize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayLUTDequantize
type NDArrayLUTDequantize struct {
	NDArrayMultiaryKernel
}

// NDArrayLUTDequantizeFrom constructs a [NDArrayLUTDequantize] from an unsafe.Pointer.
func NDArrayLUTDequantizeFrom(ptr unsafe.Pointer) NDArrayLUTDequantize {
	return NDArrayLUTDequantize{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayLUTDequantize */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraylutdequantize/4446151-initwithdevice
func NewNDArrayLUTDequantizeWithDevice(device unsafe.Pointer) NDArrayLUTDequantize {
	instance := getNDArrayLUTDequantizeClass().Alloc()
	rv := objc.Send[NDArrayLUTDequantize](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayLUTDequantizeWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayLUTDequantize */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayLUTDequantize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayLUTDequantize */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayLUTDequantize */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayLUTDequantize */


