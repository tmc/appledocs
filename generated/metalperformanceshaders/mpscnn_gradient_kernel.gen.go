// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNGradientKernel */


/* debug [class_header]: Header for MPSCNNGradientKernel */
// The class instance for the [CNNGradientKernel] class.
var (
	CNNGradientKernelClass     _CNNGradientKernelClass
	CNNGradientKernelClassOnce sync.Once
)

func getCNNGradientKernelClass() _CNNGradientKernelClass {
	CNNGradientKernelClassOnce.Do(func() {
		CNNGradientKernelClass = _CNNGradientKernelClass{objc.GetClass("MPSCNNGradientKernel")}
	})
	return CNNGradientKernelClass
}

type _CNNGradientKernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNGradientKernel */
// An interface definition for the [CNNGradientKernel] class.
type ICNNGradientKernel interface {
	ICNNBinaryKernel
	
/* debug [class_interface_properties]: Properties for CNNGradientKernel */
	// properties:
	KernelOffsetY() objectivec.IObject
	SetKernelOffsetY(value objectivec.IObject)
	KernelOffsetX() objectivec.IObject
	SetKernelOffsetX(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNGradientKernel */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, gradientStates StateBatch /* not a class type */, destinationGradients ImageBatch /* not a class type */)
	Encode()
	EncodeToCommandBufferSourceGradientSourceImageGradientState(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, gradientState IState) IImage
	EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, gradientStates StateBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, gradientState IState, destinationGradient IImage)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNGradientKernel */
// Alloc allocates a new instance without initialization.
func (cc _CNNGradientKernelClass) Alloc() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNGradientKernelClass) New() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGradientKernel) Init() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGradientKernel) Autorelease() CNNGradientKernel {
	rv := objc.Send[CNNGradientKernel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGradientKernel creates a new CNNGradientKernel instance.
func NewCNNGradientKernel() CNNGradientKernel {
	return getCNNGradientKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNGradientKernel */
// The base class for gradient layers.


// The base class for gradient layers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGradientKernel
type CNNGradientKernel struct {
	CNNBinaryKernel
}

// CNNGradientKernelFrom constructs a [CNNGradientKernel] from an unsafe.Pointer.
//
// The base class for gradient layers.
func CNNGradientKernelFrom(ptr unsafe.Pointer) CNNGradientKernel {
	return CNNGradientKernel{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942647-initwithcoder
func NewCNNGradientKernelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNGradientKernel {
	instance := getCNNGradientKernelClass().Alloc()
	rv := objc.Send[CNNGradientKernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNGradientKernelWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice
func NewCNNGradientKernelWithDevice(device unsafe.Pointer) CNNGradientKernel {
	instance := getCNNGradientKernelClass().Alloc()
	rv := objc.Send[CNNGradientKernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNGradientKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNGradientKernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNGradientKernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942653-encodebatch
func (c_ CNNGradientKernel) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942653-encodebatchtocommandbuffer
func (c_ CNNGradientKernel) EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, gradientStates StateBatch /* not a class type */, destinationGradients ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:gradientStates:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, gradientStates, destinationGradients)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStatesDestinationGradients */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942663-encode
func (c_ CNNGradientKernel) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942663-encodetocommandbuffer
func (c_ CNNGradientKernel) EncodeToCommandBufferSourceGradientSourceImageGradientState(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, gradientState IState) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:gradientState:"), commandBuffer, sourceGradient, sourceImage, gradientState)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceGradientSourceImageGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942668-encodebatchtocommandbuffer
func (c_ CNNGradientKernel) EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, gradientStates StateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:gradientStates:"), commandBuffer, sourceGradients, sourceImages, gradientStates)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942675-encodetocommandbuffer
func (c_ CNNGradientKernel) EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, gradientState IState, destinationGradient IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:gradientState:destinationGradient:"), commandBuffer, sourceGradient, sourceImage, gradientState, destinationGradient)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceGradientSourceImageGradientStateDestinationGradient */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNGradientKernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942644-kerneloffsety
func (c_ CNNGradientKernel) KernelOffsetY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelOffsetY"))
	return rv
}/* debug [instance_properties/getter]: kernelOffsetY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942644-kerneloffsety
func (c_ CNNGradientKernel) SetKernelOffsetY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelOffsetY:"), value)
}/* debug [instance_properties/setter]: kernelOffsetY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942676-kerneloffsetx
func (c_ CNNGradientKernel) KernelOffsetX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelOffsetX"))
	return rv
}/* debug [instance_properties/getter]: kernelOffsetX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942676-kerneloffsetx
func (c_ CNNGradientKernel) SetKernelOffsetX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelOffsetX:"), value)
}/* debug [instance_properties/setter]: kernelOffsetX */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNGradientKernel */


