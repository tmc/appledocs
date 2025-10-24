// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayMultiaryGradientKernel */


/* debug [class_header]: Header for MPSNDArrayMultiaryGradientKernel */
// The class instance for the [NDArrayMultiaryGradientKernel] class.
var (
	NDArrayMultiaryGradientKernelClass     _NDArrayMultiaryGradientKernelClass
	NDArrayMultiaryGradientKernelClassOnce sync.Once
)

func getNDArrayMultiaryGradientKernelClass() _NDArrayMultiaryGradientKernelClass {
	NDArrayMultiaryGradientKernelClassOnce.Do(func() {
		NDArrayMultiaryGradientKernelClass = _NDArrayMultiaryGradientKernelClass{objc.GetClass("MPSNDArrayMultiaryGradientKernel")}
	})
	return NDArrayMultiaryGradientKernelClass
}

type _NDArrayMultiaryGradientKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayMultiaryGradientKernel */
// An interface definition for the [NDArrayMultiaryGradientKernel] class.
type INDArrayMultiaryGradientKernel interface {
	INDArrayMultiaryBase
	
/* debug [class_interface_properties]: Properties for NDArrayMultiaryGradientKernel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayMultiaryGradientKernel */
	// methods:
	Encode()
	EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState) INDArray
	EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState, destination INDArray)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayMultiaryGradientKernel */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryGradientKernelClass) Alloc() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayMultiaryGradientKernelClass) New() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMultiaryGradientKernel) Init() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMultiaryGradientKernel) Autorelease() NDArrayMultiaryGradientKernel {
	rv := objc.Send[NDArrayMultiaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMultiaryGradientKernel creates a new NDArrayMultiaryGradientKernel instance.
func NewNDArrayMultiaryGradientKernel() NDArrayMultiaryGradientKernel {
	return getNDArrayMultiaryGradientKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayMultiaryGradientKernel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryGradientKernel
type NDArrayMultiaryGradientKernel struct {
	NDArrayMultiaryBase
}

// NDArrayMultiaryGradientKernelFrom constructs a [NDArrayMultiaryGradientKernel] from an unsafe.Pointer.
func NDArrayMultiaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayMultiaryGradientKernel {
	return NDArrayMultiaryGradientKernel{
		NDArrayMultiaryBase: NDArrayMultiaryBaseFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayMultiaryGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3175008-initwithcoder
func NewNDArrayMultiaryGradientKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayMultiaryGradientKernel {
	instance := getNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayMultiaryGradientKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice
func NewNDArrayMultiaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device unsafe.Pointer, count uint, sourceGradientIndex uint) NDArrayMultiaryGradientKernel {
	instance := getNDArrayMultiaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryGradientKernel](instance.ID, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), device, count, sourceGradientIndex)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayMultiaryGradientKernelWithDeviceSourceCountSourceGradientIndex */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayMultiaryGradientKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayMultiaryGradientKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayMultiaryGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143522-encode
func (n_ NDArrayMultiaryGradientKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143522-encodetocommandbuffer
func (n_ NDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:"), cmdBuf, sources, gradient, state)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraysSourceGradientGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143523-encodetocommandbuffer
func (n_ NDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, sources unsafe.Pointer, gradient INDArray, state IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:destinationArray:"), cmdBuf, sources, gradient, state, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraysSourceGradientGradientStateDestinationArray */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayMultiaryGradientKernel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayMultiaryGradientKernel */


