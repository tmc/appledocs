// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	IsLabelsGradientFilter() objectivec.IObject
	SetIsLabelsGradientFilter(value objectivec.IObject)
	LabelSmoothing() objectivec.IObject
	SetLabelSmoothing(value objectivec.IObject)
	LossType() CNNLossType get /* not a class type */
	SetLossType(value CNNLossType get /* not a class type */)
	NumberOfClasses() objectivec.IObject
	SetNumberOfClasses(value objectivec.IObject)
	PropertyCallBack() LossCallback get set /* not a class type */
	SetPropertyCallBack(value LossCallback get set /* not a class type */)
	ReductionType() CNNReductionType get /* not a class type */
	SetReductionType(value CNNReductionType get /* not a class type */)
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LossGradientNodeClass) Alloc() LossGradientNode {
	rv := objc.Send[LossGradientNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131855-initwithsourcegradient
func NewLossGradientNodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) LossGradientNode {
	instance := getLossGradientNodeClass().Alloc()
	rv := objc.Send[LossGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131856-initwithsourcegradient
func NewLossGradientNodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, weights IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) LossGradientNode {
	instance := getLossGradientNodeClass().Alloc()
	rv := objc.Send[LossGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131857-initwithsources
func NewLossGradientNodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes unsafe.Pointer, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) LossGradientNode {
	instance := getLossGradientNodeClass().Alloc()
	rv := objc.Send[LossGradientNode](instance.ID, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131861-nodewithsourcegradient
func (lc _LossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131862-nodewithsourcegradient
func (lc _LossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, weights IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131863-nodewithsources
func (lc _LossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes unsafe.Pointer, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131853-delta
func (l_ LossGradientNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131853-delta
func (l_ LossGradientNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131854-epsilon
func (l_ LossGradientNode) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131854-epsilon
func (l_ LossGradientNode) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131858-islabelsgradientfilter
func (l_ LossGradientNode) IsLabelsGradientFilter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("isLabelsGradientFilter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131858-islabelsgradientfilter
func (l_ LossGradientNode) SetIsLabelsGradientFilter(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsLabelsGradientFilter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131859-labelsmoothing
func (l_ LossGradientNode) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131859-labelsmoothing
func (l_ LossGradientNode) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131860-losstype
func (l_ LossGradientNode) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131860-losstype
func (l_ LossGradientNode) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131864-numberofclasses
func (l_ LossGradientNode) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131864-numberofclasses
func (l_ LossGradientNode) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback
func (l_ LossGradientNode) PropertyCallBack() LossCallback get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback
func (l_ LossGradientNode) SetPropertyCallBack(value LossCallback get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPropertyCallBack:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131866-reductiontype
func (l_ LossGradientNode) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131866-reductiontype
func (l_ LossGradientNode) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131867-weight
func (l_ LossGradientNode) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131867-weight
func (l_ LossGradientNode) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3547988-reduceacrossbatch
func (l_ LossGradientNode) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3547988-reduceacrossbatch
func (l_ LossGradientNode) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}







