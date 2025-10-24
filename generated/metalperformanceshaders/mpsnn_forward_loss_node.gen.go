// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ForwardLossNode] class.
var (
	ForwardLossNodeClass     _ForwardLossNodeClass
	ForwardLossNodeClassOnce sync.Once
)

func getForwardLossNodeClass() _ForwardLossNodeClass {
	ForwardLossNodeClassOnce.Do(func() {
		ForwardLossNodeClass = _ForwardLossNodeClass{objc.GetClass("MPSNNForwardLossNode")}
	})
	return ForwardLossNodeClass
}

type _ForwardLossNodeClass struct {
	class objc.Class
}





// An interface definition for the [ForwardLossNode] class.
type IForwardLossNode interface {
	IFilterNode
	

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
	PropertyCallBack() LossCallback get set /* not a class type */
	SetPropertyCallBack(value LossCallback get set /* not a class type */)
	ReductionType() CNNReductionType get /* not a class type */
	SetReductionType(value CNNReductionType get /* not a class type */)
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)


	

	// methods:
	GradientFilter()
	GradientFilterWithSource(sourceGradient IImageNode) ILossGradientNode
	GradientFilterWithSources(sourceGradient unsafe.Pointer) ILossGradientNode
	GradientFilters()
	GradientFiltersWithSource(sourceGradient IImageNode) unsafe.Pointer
	GradientFiltersWithSources(sourceGradient unsafe.Pointer) unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (fc _ForwardLossNodeClass) Alloc() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _ForwardLossNodeClass) New() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ ForwardLossNode) Init() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ ForwardLossNode) Autorelease() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewForwardLossNode creates a new ForwardLossNode instance.
func NewForwardLossNode() ForwardLossNode {
	return getForwardLossNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode
type ForwardLossNode struct {
	FilterNode
}

// ForwardLossNodeFrom constructs a [ForwardLossNode] from an unsafe.Pointer.
func ForwardLossNodeFrom(ptr unsafe.Pointer) ForwardLossNode {
	return ForwardLossNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131832-initwithsource
func NewForwardLossNodeWithSourceLabelsLossDescriptor(source IImageNode, labels IImageNode, descriptor ICNNLossDescriptor) ForwardLossNode {
	instance := getForwardLossNodeClass().Alloc()
	rv := objc.Send[ForwardLossNode](instance.ID, objc.Sel("initWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131833-initwithsource
func NewForwardLossNodeWithSourceLabelsWeightsLossDescriptor(source IImageNode, labels IImageNode, weights IImageNode, descriptor ICNNLossDescriptor) ForwardLossNode {
	instance := getForwardLossNodeClass().Alloc()
	rv := objc.Send[ForwardLossNode](instance.ID, objc.Sel("initWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131834-initwithsources
func NewForwardLossNodeWithSourcesLossDescriptor(sourceNodes unsafe.Pointer, descriptor ICNNLossDescriptor) ForwardLossNode {
	instance := getForwardLossNodeClass().Alloc()
	rv := objc.Send[ForwardLossNode](instance.ID, objc.Sel("initWithSources:lossDescriptor:"), sourceNodes, descriptor)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131837-nodewithsource
func (fc _ForwardLossNodeClass) NodeWithSourceLabelsLossDescriptor(source IImageNode, labels IImageNode, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("nodeWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131838-nodewithsource
func (fc _ForwardLossNodeClass) NodeWithSourceLabelsWeightsLossDescriptor(source IImageNode, labels IImageNode, weights IImageNode, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("nodeWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131839-nodewithsources
func (fc _ForwardLossNodeClass) NodeWithSourcesLossDescriptor(sourceNodes unsafe.Pointer, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("nodeWithSources:lossDescriptor:"), sourceNodes, descriptor)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131828-gradientfilter
func (f_ ForwardLossNode) GradientFilter() {
	objc.Send[objc.ID](f_.ID, objc.Sel("gradientFilter"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131828-gradientfilterwithsource
func (f_ ForwardLossNode) GradientFilterWithSource(sourceGradient IImageNode) ILossGradientNode {
	rv := objc.Send[LossGradientNode](f_.ID, objc.Sel("gradientFilterWithSource:"), sourceGradient)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131829-gradientfilterwithsources
func (f_ ForwardLossNode) GradientFilterWithSources(sourceGradient unsafe.Pointer) ILossGradientNode {
	rv := objc.Send[LossGradientNode](f_.ID, objc.Sel("gradientFilterWithSources:"), sourceGradient)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131830-gradientfilters
func (f_ ForwardLossNode) GradientFilters() {
	objc.Send[objc.ID](f_.ID, objc.Sel("gradientFilters"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131830-gradientfilterswithsource
func (f_ ForwardLossNode) GradientFiltersWithSource(sourceGradient IImageNode) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gradientFiltersWithSource:"), sourceGradient)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131831-gradientfilterswithsources
func (f_ ForwardLossNode) GradientFiltersWithSources(sourceGradient unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gradientFiltersWithSources:"), sourceGradient)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131826-delta
func (f_ ForwardLossNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("delta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131826-delta
func (f_ ForwardLossNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131827-epsilon
func (f_ ForwardLossNode) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131827-epsilon
func (f_ ForwardLossNode) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131835-labelsmoothing
func (f_ ForwardLossNode) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("labelSmoothing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131835-labelsmoothing
func (f_ ForwardLossNode) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabelSmoothing:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131836-losstype
func (f_ ForwardLossNode) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("lossType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131836-losstype
func (f_ ForwardLossNode) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLossType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131840-numberofclasses
func (f_ ForwardLossNode) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("numberOfClasses"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131840-numberofclasses
func (f_ ForwardLossNode) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNumberOfClasses:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback
func (f_ ForwardLossNode) PropertyCallBack() LossCallback get set /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback
func (f_ ForwardLossNode) SetPropertyCallBack(value LossCallback get set /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPropertyCallBack:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131842-reductiontype
func (f_ ForwardLossNode) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("reductionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131842-reductiontype
func (f_ ForwardLossNode) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReductionType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131843-weight
func (f_ ForwardLossNode) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("weight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131843-weight
func (f_ ForwardLossNode) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3547987-reduceacrossbatch
func (f_ ForwardLossNode) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3547987-reduceacrossbatch
func (f_ ForwardLossNode) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}







