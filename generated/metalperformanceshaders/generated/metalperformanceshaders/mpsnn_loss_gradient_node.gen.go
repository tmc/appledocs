// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNLossGradientNode */


/* debug [class_header]: Header for MPSNNLossGradientNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LossGradientNode */
// An interface definition for the [LossGradientNode] class.
type ILossGradientNode interface {
	IGradientFilterNode
	
/* debug [class_interface_properties]: Properties for LossGradientNode */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LossGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LossGradientNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LossGradientNode */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LossGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131855-initwithsourcegradient
func NewLossGradientNodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) LossGradientNode {
	instance := getLossGradientNodeClass().Alloc()
	rv := objc.Send[LossGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLossGradientNodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131856-initwithsourcegradient
func NewLossGradientNodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, weights IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) LossGradientNode {
	instance := getLossGradientNodeClass().Alloc()
	rv := objc.Send[LossGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLossGradientNodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131857-initwithsources
func NewLossGradientNodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes unsafe.Pointer, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) LossGradientNode {
	instance := getLossGradientNodeClass().Alloc()
	rv := objc.Send[LossGradientNode](instance.ID, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLossGradientNodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LossGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131861-nodewithsourcegradient
func (lc _LossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131862-nodewithsourcegradient
func (lc _LossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient IImageNode, sourceImage IImageNode, labels IImageNode, weights IImageNode, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("nodeWithSourceGradient:sourceImage:labels:weights:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceGradient, sourceImage, labels, weights, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceGradientSourceImageLabelsWeightsGradientStateLossDescriptorIsLabelsGradientFilter) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131863-nodewithsources
func (lc _LossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes unsafe.Pointer, gradientState IGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LossGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LossGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LossGradientNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131853-delta
func (l_ LossGradientNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131853-delta
func (l_ LossGradientNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131854-epsilon
func (l_ LossGradientNode) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131854-epsilon
func (l_ LossGradientNode) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131858-islabelsgradientfilter
func (l_ LossGradientNode) IsLabelsGradientFilter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("isLabelsGradientFilter"))
	return rv
}/* debug [instance_properties/getter]: isLabelsGradientFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131858-islabelsgradientfilter
func (l_ LossGradientNode) SetIsLabelsGradientFilter(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsLabelsGradientFilter:"), value)
}/* debug [instance_properties/setter]: isLabelsGradientFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131859-labelsmoothing
func (l_ LossGradientNode) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131859-labelsmoothing
func (l_ LossGradientNode) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}/* debug [instance_properties/setter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131860-losstype
func (l_ LossGradientNode) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131860-losstype
func (l_ LossGradientNode) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}/* debug [instance_properties/setter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131864-numberofclasses
func (l_ LossGradientNode) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}/* debug [instance_properties/getter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131864-numberofclasses
func (l_ LossGradientNode) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}/* debug [instance_properties/setter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback
func (l_ LossGradientNode) PropertyCallBack() LossCallback get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("propertyCallBack"))
	return rv
}/* debug [instance_properties/getter]: propertyCallBack */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback
func (l_ LossGradientNode) SetPropertyCallBack(value LossCallback get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPropertyCallBack:"), value)
}/* debug [instance_properties/setter]: propertyCallBack */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131866-reductiontype
func (l_ LossGradientNode) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131866-reductiontype
func (l_ LossGradientNode) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131867-weight
func (l_ LossGradientNode) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131867-weight
func (l_ LossGradientNode) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3547988-reduceacrossbatch
func (l_ LossGradientNode) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3547988-reduceacrossbatch
func (l_ LossGradientNode) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNLossGradientNode */


