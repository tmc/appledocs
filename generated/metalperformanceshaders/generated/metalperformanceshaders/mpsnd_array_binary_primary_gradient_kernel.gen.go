// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayBinaryPrimaryGradientKernel */


/* debug [class_header]: Header for MPSNDArrayBinaryPrimaryGradientKernel */
// The class instance for the [NDArrayBinaryPrimaryGradientKernel] class.
var (
	NDArrayBinaryPrimaryGradientKernelClass     _NDArrayBinaryPrimaryGradientKernelClass
	NDArrayBinaryPrimaryGradientKernelClassOnce sync.Once
)

func getNDArrayBinaryPrimaryGradientKernelClass() _NDArrayBinaryPrimaryGradientKernelClass {
	NDArrayBinaryPrimaryGradientKernelClassOnce.Do(func() {
		NDArrayBinaryPrimaryGradientKernelClass = _NDArrayBinaryPrimaryGradientKernelClass{objc.GetClass("MPSNDArrayBinaryPrimaryGradientKernel")}
	})
	return NDArrayBinaryPrimaryGradientKernelClass
}

type _NDArrayBinaryPrimaryGradientKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayBinaryPrimaryGradientKernel */
// An interface definition for the [NDArrayBinaryPrimaryGradientKernel] class.
type INDArrayBinaryPrimaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	
/* debug [class_interface_properties]: Properties for NDArrayBinaryPrimaryGradientKernel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayBinaryPrimaryGradientKernel */
	// methods:
	Encode()
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) INDArray
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState, destination INDArray)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayBinaryPrimaryGradientKernel */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayBinaryPrimaryGradientKernelClass) Alloc() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayBinaryPrimaryGradientKernelClass) New() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayBinaryPrimaryGradientKernel) Init() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayBinaryPrimaryGradientKernel) Autorelease() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayBinaryPrimaryGradientKernel creates a new NDArrayBinaryPrimaryGradientKernel instance.
func NewNDArrayBinaryPrimaryGradientKernel() NDArrayBinaryPrimaryGradientKernel {
	return getNDArrayBinaryPrimaryGradientKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayBinaryPrimaryGradientKernel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayBinaryPrimaryGradientKernel
type NDArrayBinaryPrimaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

// NDArrayBinaryPrimaryGradientKernelFrom constructs a [NDArrayBinaryPrimaryGradientKernel] from an unsafe.Pointer.
func NDArrayBinaryPrimaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayBinaryPrimaryGradientKernel {
	return NDArrayBinaryPrimaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayBinaryPrimaryGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3175006-initwithcoder
func NewNDArrayBinaryPrimaryGradientKernelWithCoderDevice(coder Coder /* not a class type */, device unsafe.Pointer) NDArrayBinaryPrimaryGradientKernel {
	instance := getNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayBinaryPrimaryGradientKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143515-initwithdevice
func NewNDArrayBinaryPrimaryGradientKernelWithDevice(device unsafe.Pointer) NDArrayBinaryPrimaryGradientKernel {
	instance := getNDArrayBinaryPrimaryGradientKernelClass().Alloc()
	rv := objc.Send[NDArrayBinaryPrimaryGradientKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayBinaryPrimaryGradientKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayBinaryPrimaryGradientKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayBinaryPrimaryGradientKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayBinaryPrimaryGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143513-encode
func (n_ NDArrayBinaryPrimaryGradientKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143513-encodetocommandbuffer
func (n_ NDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143514-encodetocommandbuffer
func (n_ NDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray(cmdBuf unsafe.Pointer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:destinationArray:"), cmdBuf, primarySourceArray, secondarySourceArray, gradient, state, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientStateDestinationArray */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayBinaryPrimaryGradientKernel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayBinaryPrimaryGradientKernel */


