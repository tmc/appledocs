// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionNode */


/* debug [class_header]: Header for MPSCNNConvolutionNode */
// The class instance for the [CNNConvolutionNode] class.
var (
	CNNConvolutionNodeClass     _CNNConvolutionNodeClass
	CNNConvolutionNodeClassOnce sync.Once
)

func getCNNConvolutionNodeClass() _CNNConvolutionNodeClass {
	CNNConvolutionNodeClassOnce.Do(func() {
		CNNConvolutionNodeClass = _CNNConvolutionNodeClass{objc.GetClass("MPSCNNConvolutionNode")}
	})
	return CNNConvolutionNodeClass
}

type _CNNConvolutionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionNode */
// An interface definition for the [CNNConvolutionNode] class.
type ICNNConvolutionNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNConvolutionNode */
	// properties:
	ConvolutionGradientState() IMPSCNNConvolutionGradientStateNode
	SetConvolutionGradientState(value IMPSCNNConvolutionGradientStateNode)
	AccumulatorPrecision() ConvolutionAccumulatorPrecisionOption get set /* not a class type */
	SetAccumulatorPrecision(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */)
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionNodeClass) Alloc() CNNConvolutionNode {
	rv := objc.Send[CNNConvolutionNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionNodeClass) New() CNNConvolutionNode {
	rv := objc.Send[CNNConvolutionNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionNode) Init() CNNConvolutionNode {
	rv := objc.Send[CNNConvolutionNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionNode) Autorelease() CNNConvolutionNode {
	rv := objc.Send[CNNConvolutionNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionNode creates a new CNNConvolutionNode instance.
func NewCNNConvolutionNode() CNNConvolutionNode {
	return getCNNConvolutionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionNode */
// A representation of a convolution kernel.


// A representation of a convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionNode
type CNNConvolutionNode struct {
	FilterNode
}

// CNNConvolutionNodeFrom constructs a [CNNConvolutionNode] from an unsafe.Pointer.
//
// A representation of a convolution kernel.
func CNNConvolutionNodeFrom(ptr unsafe.Pointer) CNNConvolutionNode {
	return CNNConvolutionNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866470-initwithsource
func NewCNNConvolutionNodeWithSourceWeights(sourceNode IImageNode, weights unsafe.Pointer) CNNConvolutionNode {
	instance := getCNNConvolutionNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionNode](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionNodeWithSourceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866436-nodewithsource
func (cc _CNNConvolutionNodeClass) NodeWithSourceWeights(sourceNode IImageNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:"), sourceNode, weights)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceWeights) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2942634-convolutiongradientstate
func (c_ CNNConvolutionNode) ConvolutionGradientState() IMPSCNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](c_.ID, objc.Sel("convolutionGradientState"))
	return rv
}/* debug [instance_properties/getter]: convolutionGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2942634-convolutiongradientstate
func (c_ CNNConvolutionNode) SetConvolutionGradientState(value IMPSCNNConvolutionGradientStateNode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConvolutionGradientState:"), value)
}/* debug [instance_properties/setter]: convolutionGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2980757-accumulatorprecision
func (c_ CNNConvolutionNode) AccumulatorPrecision() ConvolutionAccumulatorPrecisionOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("accumulatorPrecision"))
	return rv
}/* debug [instance_properties/getter]: accumulatorPrecision */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2980757-accumulatorprecision
func (c_ CNNConvolutionNode) SetAccumulatorPrecision(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccumulatorPrecision:"), value)
}/* debug [instance_properties/setter]: accumulatorPrecision */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/3197822-trainingstyle
func (c_ CNNConvolutionNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}/* debug [instance_properties/getter]: trainingStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/3197822-trainingstyle
func (c_ CNNConvolutionNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}/* debug [instance_properties/setter]: trainingStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionNode */


