// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayAffineInt4Dequantize */


/* debug [class_header]: Header for MPSNDArrayAffineInt4Dequantize */
// The class instance for the [NDArrayAffineInt4Dequantize] class.
var (
	NDArrayAffineInt4DequantizeClass     _NDArrayAffineInt4DequantizeClass
	NDArrayAffineInt4DequantizeClassOnce sync.Once
)

func getNDArrayAffineInt4DequantizeClass() _NDArrayAffineInt4DequantizeClass {
	NDArrayAffineInt4DequantizeClassOnce.Do(func() {
		NDArrayAffineInt4DequantizeClass = _NDArrayAffineInt4DequantizeClass{objc.GetClass("MPSNDArrayAffineInt4Dequantize")}
	})
	return NDArrayAffineInt4DequantizeClass
}

type _NDArrayAffineInt4DequantizeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayAffineInt4Dequantize */
// An interface definition for the [NDArrayAffineInt4Dequantize] class.
type INDArrayAffineInt4Dequantize interface {
	INDArrayMultiaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayAffineInt4Dequantize */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayAffineInt4Dequantize */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayAffineInt4Dequantize */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayAffineInt4DequantizeClass) Alloc() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayAffineInt4DequantizeClass) New() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayAffineInt4Dequantize) Init() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayAffineInt4Dequantize) Autorelease() NDArrayAffineInt4Dequantize {
	rv := objc.Send[NDArrayAffineInt4Dequantize](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayAffineInt4Dequantize creates a new NDArrayAffineInt4Dequantize instance.
func NewNDArrayAffineInt4Dequantize() NDArrayAffineInt4Dequantize {
	return getNDArrayAffineInt4DequantizeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayAffineInt4Dequantize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAffineInt4Dequantize
type NDArrayAffineInt4Dequantize struct {
	NDArrayMultiaryKernel
}

// NDArrayAffineInt4DequantizeFrom constructs a [NDArrayAffineInt4Dequantize] from an unsafe.Pointer.
func NDArrayAffineInt4DequantizeFrom(ptr unsafe.Pointer) NDArrayAffineInt4Dequantize {
	return NDArrayAffineInt4Dequantize{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayAffineInt4Dequantize */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayaffineint4dequantize/4446149-initwithdevice
func NewNDArrayAffineInt4DequantizeWithDeviceQuantizationDescriptor(device unsafe.Pointer, quantizationDescriptor INDArrayAffineQuantizationDescriptor) NDArrayAffineInt4Dequantize {
	instance := getNDArrayAffineInt4DequantizeClass().Alloc()
	rv := objc.Send[NDArrayAffineInt4Dequantize](instance.ID, objc.Sel("initWithDevice:quantizationDescriptor:"), device, quantizationDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayAffineInt4DequantizeWithDeviceQuantizationDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayAffineInt4Dequantize */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayAffineInt4Dequantize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayAffineInt4Dequantize */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayAffineInt4Dequantize */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayAffineInt4Dequantize */


