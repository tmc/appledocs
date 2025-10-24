// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayVectorLUTDequantize */


/* debug [class_header]: Header for MPSNDArrayVectorLUTDequantize */
// The class instance for the [NDArrayVectorLUTDequantize] class.
var (
	NDArrayVectorLUTDequantizeClass     _NDArrayVectorLUTDequantizeClass
	NDArrayVectorLUTDequantizeClassOnce sync.Once
)

func getNDArrayVectorLUTDequantizeClass() _NDArrayVectorLUTDequantizeClass {
	NDArrayVectorLUTDequantizeClassOnce.Do(func() {
		NDArrayVectorLUTDequantizeClass = _NDArrayVectorLUTDequantizeClass{objc.GetClass("MPSNDArrayVectorLUTDequantize")}
	})
	return NDArrayVectorLUTDequantizeClass
}

type _NDArrayVectorLUTDequantizeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayVectorLUTDequantize */
// An interface definition for the [NDArrayVectorLUTDequantize] class.
type INDArrayVectorLUTDequantize interface {
	INDArrayMultiaryKernel
	
/* debug [class_interface_properties]: Properties for NDArrayVectorLUTDequantize */
	// properties:
	VectorAxis() objectivec.IObject
	SetVectorAxis(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayVectorLUTDequantize */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayVectorLUTDequantize */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayVectorLUTDequantizeClass) Alloc() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayVectorLUTDequantizeClass) New() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayVectorLUTDequantize) Init() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayVectorLUTDequantize) Autorelease() NDArrayVectorLUTDequantize {
	rv := objc.Send[NDArrayVectorLUTDequantize](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayVectorLUTDequantize creates a new NDArrayVectorLUTDequantize instance.
func NewNDArrayVectorLUTDequantize() NDArrayVectorLUTDequantize {
	return getNDArrayVectorLUTDequantizeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayVectorLUTDequantize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayVectorLUTDequantize
type NDArrayVectorLUTDequantize struct {
	NDArrayMultiaryKernel
}

// NDArrayVectorLUTDequantizeFrom constructs a [NDArrayVectorLUTDequantize] from an unsafe.Pointer.
func NDArrayVectorLUTDequantizeFrom(ptr unsafe.Pointer) NDArrayVectorLUTDequantize {
	return NDArrayVectorLUTDequantize{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayVectorLUTDequantize */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayvectorlutdequantize/4446155-initwithdevice
func NewNDArrayVectorLUTDequantizeWithDeviceAxis(device unsafe.Pointer, axis uint) NDArrayVectorLUTDequantize {
	instance := getNDArrayVectorLUTDequantizeClass().Alloc()
	rv := objc.Send[NDArrayVectorLUTDequantize](instance.ID, objc.Sel("initWithDevice:axis:"), device, axis)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayVectorLUTDequantizeWithDeviceAxis */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayVectorLUTDequantize */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayVectorLUTDequantize */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayVectorLUTDequantize */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayVectorLUTDequantize */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayvectorlutdequantize/4446156-vectoraxis
func (n_ NDArrayVectorLUTDequantize) VectorAxis() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("vectorAxis"))
	return rv
}/* debug [instance_properties/getter]: vectorAxis */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayvectorlutdequantize/4446156-vectoraxis
func (n_ NDArrayVectorLUTDequantize) SetVectorAxis(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setVectorAxis:"), value)
}/* debug [instance_properties/setter]: vectorAxis */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayVectorLUTDequantize */


