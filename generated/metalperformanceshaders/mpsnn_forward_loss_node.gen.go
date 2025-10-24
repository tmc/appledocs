// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNForwardLossNode */


/* debug [class_header]: Header for MPSNNForwardLossNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ForwardLossNode */
// An interface definition for the [ForwardLossNode] class.
type IForwardLossNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for ForwardLossNode */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ForwardLossNode */
	// methods:
	GradientFilter()
	GradientFilterWithSource(sourceGradient IImageNode) ILossGradientNode
	GradientFilterWithSources(sourceGradient unsafe.Pointer) ILossGradientNode
	GradientFilters()
	GradientFiltersWithSource(sourceGradient IImageNode) unsafe.Pointer
	GradientFiltersWithSources(sourceGradient unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ForwardLossNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ForwardLossNode */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ForwardLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131832-initwithsource
func NewForwardLossNodeWithSourceLabelsLossDescriptor(source IImageNode, labels IImageNode, descriptor ICNNLossDescriptor) ForwardLossNode {
	instance := getForwardLossNodeClass().Alloc()
	rv := objc.Send[ForwardLossNode](instance.ID, objc.Sel("initWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewForwardLossNodeWithSourceLabelsLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131833-initwithsource
func NewForwardLossNodeWithSourceLabelsWeightsLossDescriptor(source IImageNode, labels IImageNode, weights IImageNode, descriptor ICNNLossDescriptor) ForwardLossNode {
	instance := getForwardLossNodeClass().Alloc()
	rv := objc.Send[ForwardLossNode](instance.ID, objc.Sel("initWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewForwardLossNodeWithSourceLabelsWeightsLossDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131834-initwithsources
func NewForwardLossNodeWithSourcesLossDescriptor(sourceNodes unsafe.Pointer, descriptor ICNNLossDescriptor) ForwardLossNode {
	instance := getForwardLossNodeClass().Alloc()
	rv := objc.Send[ForwardLossNode](instance.ID, objc.Sel("initWithSources:lossDescriptor:"), sourceNodes, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewForwardLossNodeWithSourcesLossDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ForwardLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131837-nodewithsource
func (fc _ForwardLossNodeClass) NodeWithSourceLabelsLossDescriptor(source IImageNode, labels IImageNode, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("nodeWithSource:labels:lossDescriptor:"), source, labels, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceLabelsLossDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131838-nodewithsource
func (fc _ForwardLossNodeClass) NodeWithSourceLabelsWeightsLossDescriptor(source IImageNode, labels IImageNode, weights IImageNode, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("nodeWithSource:labels:weights:lossDescriptor:"), source, labels, weights, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceLabelsWeightsLossDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131839-nodewithsources
func (fc _ForwardLossNodeClass) NodeWithSourcesLossDescriptor(sourceNodes unsafe.Pointer, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("nodeWithSources:lossDescriptor:"), sourceNodes, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourcesLossDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ForwardLossNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ForwardLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131828-gradientfilter
func (f_ ForwardLossNode) GradientFilter() {
	objc.Send[objc.ID](f_.ID, objc.Sel("gradientFilter"))
}/* debug [instance_methods/method]: GradientFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131828-gradientfilterwithsource
func (f_ ForwardLossNode) GradientFilterWithSource(sourceGradient IImageNode) ILossGradientNode {
	rv := objc.Send[LossGradientNode](f_.ID, objc.Sel("gradientFilterWithSource:"), sourceGradient)
	return rv
}/* debug [instance_methods/method]: GradientFilterWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131829-gradientfilterwithsources
func (f_ ForwardLossNode) GradientFilterWithSources(sourceGradient unsafe.Pointer) ILossGradientNode {
	rv := objc.Send[LossGradientNode](f_.ID, objc.Sel("gradientFilterWithSources:"), sourceGradient)
	return rv
}/* debug [instance_methods/method]: GradientFilterWithSources */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131830-gradientfilters
func (f_ ForwardLossNode) GradientFilters() {
	objc.Send[objc.ID](f_.ID, objc.Sel("gradientFilters"))
}/* debug [instance_methods/method]: GradientFilters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131830-gradientfilterswithsource
func (f_ ForwardLossNode) GradientFiltersWithSource(sourceGradient IImageNode) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gradientFiltersWithSource:"), sourceGradient)
	return rv
}/* debug [instance_methods/method]: GradientFiltersWithSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131831-gradientfilterswithsources
func (f_ ForwardLossNode) GradientFiltersWithSources(sourceGradient unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("gradientFiltersWithSources:"), sourceGradient)
	return rv
}/* debug [instance_methods/method]: GradientFiltersWithSources */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ForwardLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131826-delta
func (f_ ForwardLossNode) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131826-delta
func (f_ ForwardLossNode) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131827-epsilon
func (f_ ForwardLossNode) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131827-epsilon
func (f_ ForwardLossNode) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131835-labelsmoothing
func (f_ ForwardLossNode) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131835-labelsmoothing
func (f_ ForwardLossNode) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabelSmoothing:"), value)
}/* debug [instance_properties/setter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131836-losstype
func (f_ ForwardLossNode) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131836-losstype
func (f_ ForwardLossNode) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLossType:"), value)
}/* debug [instance_properties/setter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131840-numberofclasses
func (f_ ForwardLossNode) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("numberOfClasses"))
	return rv
}/* debug [instance_properties/getter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131840-numberofclasses
func (f_ ForwardLossNode) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNumberOfClasses:"), value)
}/* debug [instance_properties/setter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback
func (f_ ForwardLossNode) PropertyCallBack() LossCallback get set /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("propertyCallBack"))
	return rv
}/* debug [instance_properties/getter]: propertyCallBack */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131841-propertycallback
func (f_ ForwardLossNode) SetPropertyCallBack(value LossCallback get set /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPropertyCallBack:"), value)
}/* debug [instance_properties/setter]: propertyCallBack */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131842-reductiontype
func (f_ ForwardLossNode) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131842-reductiontype
func (f_ ForwardLossNode) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131843-weight
func (f_ ForwardLossNode) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3131843-weight
func (f_ ForwardLossNode) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3547987-reduceacrossbatch
func (f_ ForwardLossNode) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardlossnode/3547987-reduceacrossbatch
func (f_ ForwardLossNode) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNForwardLossNode */


