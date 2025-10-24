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
	// properties:
	Epsilon() float32
	ReductionType() CNNReductionType /* not a class type */
	Delta() float32
	SetDelta(value float32)
	IsLabelsGradientFilter() bool
	SetIsLabelsGradientFilter(value bool)
	LabelSmoothing() float32
	SetLabelSmoothing(value float32)
	LossType() CNNLossType /* not a class type */
	SetLossType(value CNNLossType /* not a class type */)
	NumberOfClasses() int
	SetNumberOfClasses(value int)
	PropertyCallBack() LossCallback /* not a class type */
	SetPropertyCallBack(value LossCallback /* not a class type */)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	Weight() float32
	SetWeight(value float32)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:
func (lc _LossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []IImageNode, gradientState GradientStateNode /* not a class type */, descriptor CNNLossDescriptor /* not a class type */, isLabelsGradientFilter bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/epsilon
func (l_ LossGradientNode) Epsilon() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradientNode/reductionType
func (l_ LossGradientNode) ReductionType() CNNReductionType /* not a class type */ {
	rv := objc.Send[CNNReductionType](l_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/delta
func (l_ LossGradientNode) Delta() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/delta
func (l_ LossGradientNode) SetDelta(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/islabelsgradientfilter
func (l_ LossGradientNode) IsLabelsGradientFilter() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isLabelsGradientFilter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/islabelsgradientfilter
func (l_ LossGradientNode) SetIsLabelsGradientFilter(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsLabelsGradientFilter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/labelsmoothing
func (l_ LossGradientNode) LabelSmoothing() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/labelsmoothing
func (l_ LossGradientNode) SetLabelSmoothing(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/losstype
func (l_ LossGradientNode) LossType() CNNLossType /* not a class type */ {
	rv := objc.Send[CNNLossType](l_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/losstype
func (l_ LossGradientNode) SetLossType(value CNNLossType /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/numberofclasses
func (l_ LossGradientNode) NumberOfClasses() int {
	rv := objc.Send[int](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/numberofclasses
func (l_ LossGradientNode) SetNumberOfClasses(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/propertycallback
func (l_ LossGradientNode) PropertyCallBack() LossCallback /* not a class type */ {
	rv := objc.Send[LossCallback](l_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/propertycallback
func (l_ LossGradientNode) SetPropertyCallBack(value LossCallback /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPropertyCallBack:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/reduceacrossbatch
func (l_ LossGradientNode) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/reduceacrossbatch
func (l_ LossGradientNode) SetReduceAcrossBatch(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/weight
func (l_ LossGradientNode) Weight() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/weight
func (l_ LossGradientNode) SetWeight(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}



