// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LossGradientNode] class.
var (
	LossGradientNodeClass     _LossGradientNodeClass
	LossGradientNodeClassOnce sync.Once
)

func getLossGradientNodeClass() _LossGradientNodeClass {
	LossGradientNodeClassOnce.Do(func() {
		LossGradientNodeClass = _LossGradientNodeClass{objc.GetClass("MPSNNLossGradientNode")}
	})
	return LossGradientNodeClass
}

type _LossGradientNodeClass struct {
	class objc.Class
}

// An interface definition for the [LossGradientNode] class.
type ILossGradientNode interface {
	IGradientFilterNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode
type LossGradientNode struct {
	GradientFilterNode
}

// LossGradientNodeFrom constructs a [LossGradientNode] from an unsafe.Pointer.
func LossGradientNodeFrom(ptr unsafe.Pointer) LossGradientNode {
	return LossGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LossGradientNodeClass) Alloc() LossGradientNode {
	rv := objc.Send[LossGradientNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LossGradientNodeClass) New() LossGradientNode {
	rv := objc.Send[LossGradientNode](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LossGradientNode) Init() LossGradientNode {
	rv := objc.Send[LossGradientNode](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LossGradientNode) Autorelease() LossGradientNode {
	rv := objc.Send[LossGradientNode](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLossGradientNode creates a new LossGradientNode instance.
func NewLossGradientNode() LossGradientNode {
	return getLossGradientNodeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func (lc _LossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes unsafe.Pointer, gradientState unsafe.Pointer, descriptor unsafe.Pointer, isLabelsGradientFilter bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/delta
func (l_ LossGradientNode) Delta() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("delta"))
	return rv
}


// SetDelta sets the value of the delta property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/delta
func (l_ LossGradientNode) SetDelta(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/numberofclasses
func (l_ LossGradientNode) NumberOfClasses() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// SetNumberOfClasses sets the value of the numberOfClasses property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/numberofclasses
func (l_ LossGradientNode) SetNumberOfClasses(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/reduceacrossbatch
func (l_ LossGradientNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// SetReduceAcrossBatch sets the value of the reduceAcrossBatch property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/reduceacrossbatch
func (l_ LossGradientNode) SetReduceAcrossBatch(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/labelsmoothing
func (l_ LossGradientNode) LabelSmoothing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// SetLabelSmoothing sets the value of the labelSmoothing property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/labelsmoothing
func (l_ LossGradientNode) SetLabelSmoothing(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/islabelsgradientfilter
func (l_ LossGradientNode) IsLabelsGradientFilter() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isLabelsGradientFilter"))
	return rv
}


// SetIsLabelsGradientFilter sets the value of the isLabelsGradientFilter property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/islabelsgradientfilter
func (l_ LossGradientNode) SetIsLabelsGradientFilter(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsLabelsGradientFilter:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/propertycallback
func (l_ LossGradientNode) PropertyCallBack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// SetPropertyCallBack sets the value of the propertyCallBack property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/propertycallback
func (l_ LossGradientNode) SetPropertyCallBack(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPropertyCallBack:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/weight
func (l_ LossGradientNode) Weight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("weight"))
	return rv
}


// SetWeight sets the value of the weight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/weight
func (l_ LossGradientNode) SetWeight(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/losstype
func (l_ LossGradientNode) LossType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("lossType"))
	return rv
}


// SetLossType sets the value of the lossType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/losstype
func (l_ LossGradientNode) SetLossType(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/epsilon
func (l_ LossGradientNode) Epsilon() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("epsilon"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/reductionType
func (l_ LossGradientNode) ReductionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("reductionType"))
	return rv
}



