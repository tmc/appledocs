// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ForwardLoss] class.
var (
	ForwardLossClass     _ForwardLossClass
	ForwardLossClassOnce sync.Once
)

func getForwardLossClass() _ForwardLossClass {
	ForwardLossClassOnce.Do(func() {
		ForwardLossClass = _ForwardLossClass{objc.GetClass("MPSNNForwardLoss")}
	})
	return ForwardLossClass
}

type _ForwardLossClass struct {
	class objc.Class
}





// An interface definition for the [ForwardLoss] class.
type IForwardLoss interface {
	ICNNKernel
	

	// properties:
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	LabelSmoothing() objectivec.IObject
	SetLabelSmoothing(value objectivec.IObject)
	LossType() CNNLossType get /* not a class type */
	SetLossType(value CNNLossType get /* not a class type */)
	NumberOfClasses() objectivec.IObject
	SetNumberOfClasses(value objectivec.IObject)
	ReductionType() CNNReductionType get /* not a class type */
	SetReductionType(value CNNReductionType get /* not a class type */)
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)


	

	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */


}





// Alloc allocates a new instance without initialization.
func (fc _ForwardLossClass) Alloc() ForwardLoss {
	rv := objc.Send[ForwardLoss](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _ForwardLossClass) New() ForwardLoss {
	rv := objc.Send[ForwardLoss](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ ForwardLoss) Init() ForwardLoss {
	rv := objc.Send[ForwardLoss](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ ForwardLoss) Autorelease() ForwardLoss {
	rv := objc.Send[ForwardLoss](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewForwardLoss creates a new ForwardLoss instance.
func NewForwardLoss() ForwardLoss {
	return getForwardLossClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss
type ForwardLoss struct {
	CNNKernel
}

// ForwardLossFrom constructs a [ForwardLoss] from an unsafe.Pointer.
func ForwardLossFrom(ptr unsafe.Pointer) ForwardLoss {
	return ForwardLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131801-initwithcoder
func NewForwardLossWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ForwardLoss {
	instance := getForwardLossClass().Alloc()
	rv := objc.Send[ForwardLoss](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131802-initwithdevice
func NewForwardLossWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNLossDescriptor) ForwardLoss {
	instance := getForwardLossClass().Alloc()
	rv := objc.Send[ForwardLoss](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131798-encodebatch
func (f_ ForwardLoss) EncodeBatch() {
	objc.Send[objc.ID](f_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131798-encodebatchtocommandbuffer
func (f_ ForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:"), commandBuffer, sourceImages, labels, weights, destinationStates, destinationImages)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131799-encodebatchtocommandbuffer
func (f_ ForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](f_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, labels, weights, outStates, isTemporary)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131797-delta
func (f_ ForwardLoss) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131797-delta
func (f_ ForwardLoss) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131800-epsilon
func (f_ ForwardLoss) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131800-epsilon
func (f_ ForwardLoss) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131803-labelsmoothing
func (f_ ForwardLoss) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131803-labelsmoothing
func (f_ ForwardLoss) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131804-losstype
func (f_ ForwardLoss) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131804-losstype
func (f_ ForwardLoss) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131805-numberofclasses
func (f_ ForwardLoss) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131805-numberofclasses
func (f_ ForwardLoss) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131806-reductiontype
func (f_ ForwardLoss) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131806-reductiontype
func (f_ ForwardLoss) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131807-weight
func (f_ ForwardLoss) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131807-weight
func (f_ ForwardLoss) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3547985-reduceacrossbatch
func (f_ ForwardLoss) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3547985-reduceacrossbatch
func (f_ ForwardLoss) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}







