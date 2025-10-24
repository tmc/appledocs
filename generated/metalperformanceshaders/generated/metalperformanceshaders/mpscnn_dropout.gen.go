// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDropout */


/* debug [class_header]: Header for MPSCNNDropout */
// The class instance for the [CNNDropout] class.
var (
	CNNDropoutClass     _CNNDropoutClass
	CNNDropoutClassOnce sync.Once
)

func getCNNDropoutClass() _CNNDropoutClass {
	CNNDropoutClassOnce.Do(func() {
		CNNDropoutClass = _CNNDropoutClass{objc.GetClass("MPSCNNDropout")}
	})
	return CNNDropoutClass
}

type _CNNDropoutClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDropout */
// An interface definition for the [CNNDropout] class.
type ICNNDropout interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNDropout */
	// properties:
	Seed() objectivec.IObject
	SetSeed(value objectivec.IObject)
	MaskStrideInPixels() Size get /* not a class type */
	SetMaskStrideInPixels(value Size get /* not a class type */)
	KeepProbability() objectivec.IObject
	SetKeepProbability(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDropout */
	// methods:
	ResultStateBatch()
	ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) ICNNDropoutGradientState
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNDropoutGradientState
	TemporaryResultStateBatch()
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNDropoutGradientStateBatch /* not a class type */
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNDropoutGradientState
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDropout */
// Alloc allocates a new instance without initialization.
func (cc _CNNDropoutClass) Alloc() CNNDropout {
	rv := objc.Send[CNNDropout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDropoutClass) New() CNNDropout {
	rv := objc.Send[CNNDropout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDropout) Init() CNNDropout {
	rv := objc.Send[CNNDropout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDropout) Autorelease() CNNDropout {
	rv := objc.Send[CNNDropout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDropout creates a new CNNDropout instance.
func NewCNNDropout() CNNDropout {
	return getCNNDropoutClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDropout */
// A dropout filter.


// A dropout filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropout
type CNNDropout struct {
	CNNKernel
}

// CNNDropoutFrom constructs a [CNNDropout] from an unsafe.Pointer.
//
// A dropout filter.
func CNNDropoutFrom(ptr unsafe.Pointer) CNNDropout {
	return CNNDropout{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDropout */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942514-initwithcoder
func NewCNNDropoutWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNDropout {
	instance := getCNNDropoutClass().Alloc()
	rv := objc.Send[CNNDropout](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942522-initwithdevice
func NewCNNDropoutWithDeviceKeepProbabilitySeedMaskStrideInPixels(device unsafe.Pointer, keepProbability float32, seed uint, maskStrideInPixels Size /* not a class type */) CNNDropout {
	instance := getCNNDropoutClass().Alloc()
	rv := objc.Send[CNNDropout](instance.ID, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), device, keepProbability, seed, maskStrideInPixels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDropoutWithDeviceKeepProbabilitySeedMaskStrideInPixels */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDropout */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDropout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDropout */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131792-resultstatebatch
func (c_ CNNDropout) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}/* debug [instance_methods/method]: ResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131792-resultstatebatchforsourceimage
func (c_ CNNDropout) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) ICNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateBatchForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131793-resultstate
func (c_ CNNDropout) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131793-resultstateforsourceimage
func (c_ CNNDropout) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131794-temporaryresultstatebatch
func (c_ CNNDropout) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}/* debug [instance_methods/method]: TemporaryResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131794-temporaryresultstatebatchforcomm
func (c_ CNNDropout) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNDropoutGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNDropoutGradientStateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131795-temporaryresultstate
func (c_ CNNDropout) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/3131795-temporaryresultstateforcommandbu
func (c_ CNNDropout) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDropout */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942517-seed
func (c_ CNNDropout) Seed() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942517-seed
func (c_ CNNDropout) SetSeed(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}/* debug [instance_properties/setter]: seed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942519-maskstrideinpixels
func (c_ CNNDropout) MaskStrideInPixels() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("maskStrideInPixels"))
	return rv
}/* debug [instance_properties/getter]: maskStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942519-maskstrideinpixels
func (c_ CNNDropout) SetMaskStrideInPixels(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaskStrideInPixels:"), value)
}/* debug [instance_properties/setter]: maskStrideInPixels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942524-keepprobability
func (c_ CNNDropout) KeepProbability() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("keepProbability"))
	return rv
}/* debug [instance_properties/getter]: keepProbability */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942524-keepprobability
func (c_ CNNDropout) SetKeepProbability(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeepProbability:"), value)
}/* debug [instance_properties/setter]: keepProbability */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDropout */


