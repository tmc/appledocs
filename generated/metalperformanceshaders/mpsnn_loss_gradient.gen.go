// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer objc.ID, sourceGradients unsafe.Pointer, sourceImages unsafe.Pointer, labels unsafe.Pointer, weights unsafe.Pointer, sourceStates unsafe.Pointer, destinationGradients unsafe.Pointer)
}

//
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

// Alloc allocates a new instance without initialization.
func (lc _LossGradientClass) Alloc() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/init(device:lossDescriptor:)
func NewLossGradientWithDeviceLossDescriptor(device objc.ID, lossDescriptor unsafe.Pointer) LossGradient {
	instance := getLossGradientClass().Alloc()
	rv := objc.Send[LossGradient](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient/encodeBatch(commandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:)
func (l_ LossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer objc.ID, sourceGradients unsafe.Pointer, sourceImages unsafe.Pointer, labels unsafe.Pointer, weights unsafe.Pointer, sourceStates unsafe.Pointer, destinationGradients unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates, destinationGradients)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/computelabelgradients
func (l_ LossGradient) ComputeLabelGradients() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("computeLabelGradients"))
	return rv
}


// SetComputeLabelGradients sets the value of the computeLabelGradients property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/computelabelgradients
func (l_ LossGradient) SetComputeLabelGradients(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setComputeLabelGradients:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/delta
func (l_ LossGradient) Delta() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("delta"))
	return rv
}


// SetDelta sets the value of the delta property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/delta
func (l_ LossGradient) SetDelta(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/epsilon
func (l_ LossGradient) Epsilon() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("epsilon"))
	return rv
}


// SetEpsilon sets the value of the epsilon property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/epsilon
func (l_ LossGradient) SetEpsilon(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEpsilon:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/labelsmoothing
func (l_ LossGradient) LabelSmoothing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// SetLabelSmoothing sets the value of the labelSmoothing property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/labelsmoothing
func (l_ LossGradient) SetLabelSmoothing(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/losstype
func (l_ LossGradient) LossType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("lossType"))
	return rv
}


// SetLossType sets the value of the lossType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/losstype
func (l_ LossGradient) SetLossType(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/numberofclasses
func (l_ LossGradient) NumberOfClasses() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// SetNumberOfClasses sets the value of the numberOfClasses property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/numberofclasses
func (l_ LossGradient) SetNumberOfClasses(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reduceacrossbatch
func (l_ LossGradient) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// SetReduceAcrossBatch sets the value of the reduceAcrossBatch property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reduceacrossbatch
func (l_ LossGradient) SetReduceAcrossBatch(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reductiontype
func (l_ LossGradient) ReductionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("reductionType"))
	return rv
}


// SetReductionType sets the value of the reductionType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/reductiontype
func (l_ LossGradient) SetReductionType(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReductionType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/weight
func (l_ LossGradient) Weight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("weight"))
	return rv
}


// SetWeight sets the value of the weight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/weight
func (l_ LossGradient) SetWeight(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}


