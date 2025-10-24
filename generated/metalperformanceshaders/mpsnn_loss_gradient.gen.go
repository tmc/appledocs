// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [LossGradient] class.
var (
	LossGradientClass     _LossGradientClass
	LossGradientClassOnce sync.Once
)

func getLossGradientClass() _LossGradientClass {
	LossGradientClassOnce.Do(func() {
		LossGradientClass = _LossGradientClass{objc.GetClass("MPSNNLossGradient")}
	})
	return LossGradientClass
}

type _LossGradientClass struct {
	class objc.Class
}





// An interface definition for the [LossGradient] class.
type ILossGradient interface {
	ICNNBinaryKernel
	

	// properties:
	ComputeLabelGradients() objectivec.IObject
	SetComputeLabelGradients(value objectivec.IObject)
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
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationGradients ImageBatch /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (lc _LossGradientClass) Alloc() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LossGradientClass) New() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LossGradient) Init() LossGradient {
	rv := objc.Send[LossGradient](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LossGradient) Autorelease() LossGradient {
	rv := objc.Send[LossGradient](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLossGradient creates a new LossGradient instance.
func NewLossGradient() LossGradient {
	return getLossGradientClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient
type LossGradient struct {
	CNNBinaryKernel
}

// LossGradientFrom constructs a [LossGradient] from an unsafe.Pointer.
func LossGradientFrom(ptr unsafe.Pointer) LossGradient {
	return LossGradient{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131816-initwithcoder
func NewLossGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) LossGradient {
	instance := getLossGradientClass().Alloc()
	rv := objc.Send[LossGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131817-initwithdevice
func NewLossGradientWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNLossDescriptor) LossGradient {
	instance := getLossGradientClass().Alloc()
	rv := objc.Send[LossGradient](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131813-encodebatch
func (l_ LossGradient) EncodeBatch() {
	objc.Send[objc.ID](l_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131813-encodebatchtocommandbuffer
func (l_ LossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](l_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131814-encodebatchtocommandbuffer
func (l_ LossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationGradients ImageBatch /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates, destinationGradients)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131811-computelabelgradients
func (l_ LossGradient) ComputeLabelGradients() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("computeLabelGradients"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131811-computelabelgradients
func (l_ LossGradient) SetComputeLabelGradients(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setComputeLabelGradients:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131812-delta
func (l_ LossGradient) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131812-delta
func (l_ LossGradient) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131815-epsilon
func (l_ LossGradient) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131815-epsilon
func (l_ LossGradient) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131818-labelsmoothing
func (l_ LossGradient) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131818-labelsmoothing
func (l_ LossGradient) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131819-losstype
func (l_ LossGradient) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131819-losstype
func (l_ LossGradient) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131820-numberofclasses
func (l_ LossGradient) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131820-numberofclasses
func (l_ LossGradient) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131821-reductiontype
func (l_ LossGradient) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131821-reductiontype
func (l_ LossGradient) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131822-weight
func (l_ LossGradient) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131822-weight
func (l_ LossGradient) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3547986-reduceacrossbatch
func (l_ LossGradient) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3547986-reduceacrossbatch
func (l_ LossGradient) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}







