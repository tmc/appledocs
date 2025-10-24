// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBatchNormalizationGradient */


/* debug [class_header]: Header for MPSCNNBatchNormalizationGradient */
// The class instance for the [CNNBatchNormalizationGradient] class.
var (
	CNNBatchNormalizationGradientClass     _CNNBatchNormalizationGradientClass
	CNNBatchNormalizationGradientClassOnce sync.Once
)

func getCNNBatchNormalizationGradientClass() _CNNBatchNormalizationGradientClass {
	CNNBatchNormalizationGradientClassOnce.Do(func() {
		CNNBatchNormalizationGradientClass = _CNNBatchNormalizationGradientClass{objc.GetClass("MPSCNNBatchNormalizationGradient")}
	})
	return CNNBatchNormalizationGradientClass
}

type _CNNBatchNormalizationGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBatchNormalizationGradient */
// An interface definition for the [CNNBatchNormalizationGradient] class.
type ICNNBatchNormalizationGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNBatchNormalizationGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBatchNormalizationGradient */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState, destinationGradients ImageBatch /* not a class type */)
	Encode()
	EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState) IImage
	EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationGradient IImage)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBatchNormalizationGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationGradientClass) Alloc() CNNBatchNormalizationGradient {
	rv := objc.Send[CNNBatchNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationGradientClass) New() CNNBatchNormalizationGradient {
	rv := objc.Send[CNNBatchNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalizationGradient) Init() CNNBatchNormalizationGradient {
	rv := objc.Send[CNNBatchNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalizationGradient) Autorelease() CNNBatchNormalizationGradient {
	rv := objc.Send[CNNBatchNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalizationGradient creates a new CNNBatchNormalizationGradient instance.
func NewCNNBatchNormalizationGradient() CNNBatchNormalizationGradient {
	return getCNNBatchNormalizationGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBatchNormalizationGradient */
// A gradient batch normalization kernel.


// A gradient batch normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationGradient
type CNNBatchNormalizationGradient struct {
	CNNGradientKernel
}

// CNNBatchNormalizationGradientFrom constructs a [CNNBatchNormalizationGradient] from an unsafe.Pointer.
//
// A gradient batch normalization kernel.
func CNNBatchNormalizationGradientFrom(ptr unsafe.Pointer) CNNBatchNormalizationGradient {
	return CNNBatchNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBatchNormalizationGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/3019331-initwithcoder
func NewCNNBatchNormalizationGradientWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNBatchNormalizationGradient {
	instance := getCNNBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBatchNormalizationGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/3019332-initwithdevice
func NewCNNBatchNormalizationGradientWithDeviceFusedNeuronDescriptor(device unsafe.Pointer, fusedNeuronDescriptor INeuronDescriptor) CNNBatchNormalizationGradient {
	instance := getCNNBatchNormalizationGradientClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationGradient](instance.ID, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), device, fusedNeuronDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBatchNormalizationGradientWithDeviceFusedNeuronDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBatchNormalizationGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBatchNormalizationGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBatchNormalizationGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2942590-encodebatch
func (c_ CNNBatchNormalizationGradient) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2942590-encodebatchtocommandbuffer
func (c_ CNNBatchNormalizationGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), commandBuffer, sourceGradients, sourceImages, batchNormalizationState)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2942608-encodebatchtocommandbuffer
func (c_ CNNBatchNormalizationGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState, destinationGradients ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, batchNormalizationState, destinationGradients)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationStateDestinationGradients */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2951885-encode
func (c_ CNNBatchNormalizationGradient) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2951885-encodetocommandbuffer
func (c_ CNNBatchNormalizationGradient) EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:batchNormalizationState:"), commandBuffer, sourceGradient, sourceImage, batchNormalizationState)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2951895-encodetocommandbuffer
func (c_ CNNBatchNormalizationGradient) EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient(commandBuffer unsafe.Pointer, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationGradient IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:batchNormalizationState:destinationGradient:"), commandBuffer, sourceGradient, sourceImage, batchNormalizationState, destinationGradient)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationStateDestinationGradient */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBatchNormalizationGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBatchNormalizationGradient */


