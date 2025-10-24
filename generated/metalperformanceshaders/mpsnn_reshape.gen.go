// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReshape */


/* debug [class_header]: Header for MPSNNReshape */
// The class instance for the [Reshape] class.
var (
	ReshapeClass     _ReshapeClass
	ReshapeClassOnce sync.Once
)

func getReshapeClass() _ReshapeClass {
	ReshapeClassOnce.Do(func() {
		ReshapeClass = _ReshapeClass{objc.GetClass("MPSNNReshape")}
	})
	return ReshapeClass
}

type _ReshapeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Reshape */
// An interface definition for the [Reshape] class.
type IReshape interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for Reshape */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Reshape */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */
	Encode()
	EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, outState objectivec.IObject, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage
	EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Reshape */
// Alloc allocates a new instance without initialization.
func (rc _ReshapeClass) Alloc() Reshape {
	rv := objc.Send[Reshape](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReshapeClass) New() Reshape {
	rv := objc.Send[Reshape](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Reshape) Init() Reshape {
	rv := objc.Send[Reshape](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Reshape) Autorelease() Reshape {
	rv := objc.Send[Reshape](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshape creates a new Reshape instance.
func NewReshape() Reshape {
	return getReshapeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Reshape */
// The base class for reshape operations.


// The base class for reshape operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshape
type Reshape struct {
	CNNKernel
}

// ReshapeFrom constructs a [Reshape] from an unsafe.Pointer.
//
// The base class for reshape operations.
func ReshapeFrom(ptr unsafe.Pointer) Reshape {
	return Reshape{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Reshape */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/2951930-initwithcoder
func NewReshapeWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) Reshape {
	instance := getReshapeClass().Alloc()
	rv := objc.Send[Reshape](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReshapeWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/2951929-initwithdevice
func NewReshapeWithDevice(device unsafe.Pointer) Reshape {
	instance := getReshapeClass().Alloc()
	rv := objc.Send[Reshape](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReshapeWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Reshape */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Reshape */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Reshape */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547989-encodebatch
func (r_ Reshape) EncodeBatch() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547989-encodebatchtocommandbuffer
func (r_ Reshape) EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](r_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationStates:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImages, outStates, isTemporary, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesDestinationStatesDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547990-encodebatchtocommandbuffer
func (r_ Reshape) EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](r_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImages, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesReshapedWidthReshapedHeightReshapedFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547991-encode
func (r_ Reshape) Encode() {
	objc.Send[objc.ID](r_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547991-encodetocommandbuffer
func (r_ Reshape) EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, outState objectivec.IObject, isTemporary bool, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:destinationState:destinationStateIsTemporary:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImage, outState, isTemporary, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageDestinationStateDestinationStateIsTemporaryReshapedWidthReshapedHeightReshapedFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshape/3547992-encodetocommandbuffer
func (r_ Reshape) EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels(commandBuffer unsafe.Pointer, sourceImage IImage, reshapedWidth uint, reshapedHeight uint, reshapedFeatureChannels uint) IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:reshapedWidth:reshapedHeight:reshapedFeatureChannels:"), commandBuffer, sourceImage, reshapedWidth, reshapedHeight, reshapedFeatureChannels)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageReshapedWidthReshapedHeightReshapedFeatureChannels */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Reshape */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReshape */


