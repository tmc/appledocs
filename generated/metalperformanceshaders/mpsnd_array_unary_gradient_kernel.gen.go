// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayUnaryGradientKernel */


/* debug [class_header]: Header for MPSNDArrayUnaryGradientKernel */
// The class instance for the [NDArrayUnaryGradientKernel] class.
var (
	NDArrayUnaryGradientKernelClass     _NDArrayUnaryGradientKernelClass
	NDArrayUnaryGradientKernelClassOnce sync.Once
)

func getNDArrayUnaryGradientKernelClass() _NDArrayUnaryGradientKernelClass {
	NDArrayUnaryGradientKernelClassOnce.Do(func() {
		NDArrayUnaryGradientKernelClass = _NDArrayUnaryGradientKernelClass{objc.GetClass("MPSNDArrayUnaryGradientKernel")}
	})
	return NDArrayUnaryGradientKernelClass
}

type _NDArrayUnaryGradientKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayUnaryGradientKernel */
// An interface definition for the [NDArrayUnaryGradientKernel] class.
type INDArrayUnaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	
/* debug [class_interface_properties]: Properties for NDArrayUnaryGradientKernel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayUnaryGradientKernel */
	// methods:
	Encode()
	EncodeToCommandBufferSourceArraySourceGradientGradientState(cmdBuf unsafe.Pointer, sourceArray INDArray, gradient INDArray, state IState) INDArray
	EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, gradient INDArray, state IState, destination INDArray)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayUnaryGradientKernel */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayUnaryGradientKernelClass) Alloc() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayUnaryGradientKernelClass) New() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayUnaryGradientKernel) Init() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayUnaryGradientKernel) Autorelease() NDArrayUnaryGradientKernel {
	rv := objc.Send[NDArrayUnaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayUnaryGradientKernel creates a new NDArrayUnaryGradientKernel instance.
func NewNDArrayUnaryGradientKernel() NDArrayUnaryGradientKernel {
	return getNDArrayUnaryGradientKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayUnaryGradientKernel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayUnaryGradientKernel
type NDArrayUnaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

// NDArrayUnaryGradientKernelFrom constructs a [NDArrayUnaryGradientKernel] from an unsafe.Pointer.
func NDArrayUnaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayUnaryGradientKernel {
	return NDArrayUnaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayUnaryGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3175011-initwithcoder
func NewNDArrayUnaryGradientKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayUnaryGradientKernel {
	instance := getNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayUnaryGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayUnaryGradientKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143532-initwithdevice
func NewNDArrayUnaryGradientKernelWithDevice(device unsafe.Pointer) NDArrayUnaryGradientKernel {
	instance := getNDArrayUnaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayUnaryGradientKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayUnaryGradientKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayUnaryGradientKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayUnaryGradientKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayUnaryGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143530-encode
func (n_ NDArrayUnaryGradientKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143530-encodetocommandbuffer
func (n_ NDArrayUnaryGradientKernel) EncodeToCommandBufferSourceArraySourceGradientGradientState(cmdBuf unsafe.Pointer, sourceArray INDArray, gradient INDArray, state IState) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:"), cmdBuf, sourceArray, gradient, state)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraySourceGradientGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143531-encodetocommandbuffer
func (n_ NDArrayUnaryGradientKernel) EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, sourceArray INDArray, gradient INDArray, state IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, sourceArray, gradient, state, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraySourceGradientGradientStateDestinationArray */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayUnaryGradientKernel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayUnaryGradientKernel */


