// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayMultiaryKernel */


/* debug [class_header]: Header for MPSNDArrayMultiaryKernel */
// The class instance for the [NDArrayMultiaryKernel] class.
var (
	NDArrayMultiaryKernelClass     _NDArrayMultiaryKernelClass
	NDArrayMultiaryKernelClassOnce sync.Once
)

func getNDArrayMultiaryKernelClass() _NDArrayMultiaryKernelClass {
	NDArrayMultiaryKernelClassOnce.Do(func() {
		NDArrayMultiaryKernelClass = _NDArrayMultiaryKernelClass{objc.GetClass("MPSNDArrayMultiaryKernel")}
	})
	return NDArrayMultiaryKernelClass
}

type _NDArrayMultiaryKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayMultiaryKernel */
// An interface definition for the [NDArrayMultiaryKernel] class.
type INDArrayMultiaryKernel interface {
	INDArrayMultiaryBase
	
/* debug [class_interface_properties]: Properties for NDArrayMultiaryKernel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayMultiaryKernel */
	// methods:
	Encode()
	EncodeToCommandBufferSourceArrays(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer) INDArray
	EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray)
	EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState IState, destination INDArray)
	EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray
	EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray(encoder unsafe.Pointer, commandBuffer unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayMultiaryKernel */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryKernelClass) Alloc() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayMultiaryKernelClass) New() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMultiaryKernel) Init() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMultiaryKernel) Autorelease() NDArrayMultiaryKernel {
	rv := objc.Send[NDArrayMultiaryKernel](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMultiaryKernel creates a new NDArrayMultiaryKernel instance.
func NewNDArrayMultiaryKernel() NDArrayMultiaryKernel {
	return getNDArrayMultiaryKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayMultiaryKernel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryKernel
type NDArrayMultiaryKernel struct {
	NDArrayMultiaryBase
}

// NDArrayMultiaryKernelFrom constructs a [NDArrayMultiaryKernel] from an unsafe.Pointer.
func NDArrayMultiaryKernelFrom(ptr unsafe.Pointer) NDArrayMultiaryKernel {
	return NDArrayMultiaryKernel{
		NDArrayMultiaryBase: NDArrayMultiaryBaseFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayMultiaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175009-initwithcoder
func NewNDArrayMultiaryKernelWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayMultiaryKernel {
	instance := getNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryKernel](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayMultiaryKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice
func NewNDArrayMultiaryKernelWithDeviceSourceCount(device unsafe.Pointer, count uint) NDArrayMultiaryKernel {
	instance := getNDArrayMultiaryKernelClass().Alloc()
	rv := objc.Send[NDArrayMultiaryKernel](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayMultiaryKernelWithDeviceSourceCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayMultiaryKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayMultiaryKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayMultiaryKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143525-encode
func (n_ NDArrayMultiaryKernel) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143525-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArrays(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:"), cmdBuf, sourceArrays)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArrays */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143526-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:destinationArray:"), cmdBuf, sourceArrays, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraysDestinationArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143527-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateDestinationArray(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState IState, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:destinationArray:"), cmdBuf, sourceArrays, outGradientState, destination)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraysResultStateDestinationArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3143528-encodetocommandbuffer
func (n_ NDArrayMultiaryKernel) EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary(cmdBuf unsafe.Pointer, sourceArrays unsafe.Pointer, outGradientState objectivec.IObject, outputStateIsTemporary bool) INDArray {
	rv := objc.Send[NDArray](n_.ID, objc.Sel("encodeToCommandBuffer:sourceArrays:resultState:outputStateIsTemporary:"), cmdBuf, sourceArrays, outGradientState, outputStateIsTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceArraysResultStateOutputStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/4462738-encodetocommandencoder
func (n_ NDArrayMultiaryKernel) EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray(encoder unsafe.Pointer, commandBuffer unsafe.Pointer, sourceArrays unsafe.Pointer, destination INDArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeToCommandEncoder:commandBuffer:sourceArrays:destinationArray:"), encoder, commandBuffer, sourceArrays, destination)
}/* debug [instance_methods/method]: EncodeToCommandEncoderCommandBufferSourceArraysDestinationArray */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayMultiaryKernel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayMultiaryKernel */


