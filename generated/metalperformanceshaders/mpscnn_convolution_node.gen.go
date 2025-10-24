// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNConvolutionNode] class.
type ICNNConvolutionNode interface {
	IFilterNode
	

	// properties:
	ConvolutionGradientState() IMPSCNNConvolutionGradientStateNode
	SetConvolutionGradientState(value IMPSCNNConvolutionGradientStateNode)
	AccumulatorPrecision() ConvolutionAccumulatorPrecisionOption get set /* not a class type */
	SetAccumulatorPrecision(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */)
	TrainingStyle() TrainingStyle get set /* not a class type */
	SetTrainingStyle(value TrainingStyle get set /* not a class type */)


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866470-initwithsource
func NewCNNConvolutionNodeWithSourceWeights(sourceNode IImageNode, weights unsafe.Pointer) CNNConvolutionNode {
	instance := getCNNConvolutionNodeClass().Alloc()
	rv := objc.Send[CNNConvolutionNode](instance.ID, objc.Sel("initWithSource:weights:"), sourceNode, weights)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866436-nodewithsource
func (cc _CNNConvolutionNodeClass) NodeWithSourceWeights(sourceNode IImageNode, weights unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:weights:"), sourceNode, weights)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2942634-convolutiongradientstate
func (c_ CNNConvolutionNode) ConvolutionGradientState() IMPSCNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](c_.ID, objc.Sel("convolutionGradientState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2942634-convolutiongradientstate
func (c_ CNNConvolutionNode) SetConvolutionGradientState(value IMPSCNNConvolutionGradientStateNode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConvolutionGradientState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2980757-accumulatorprecision
func (c_ CNNConvolutionNode) AccumulatorPrecision() ConvolutionAccumulatorPrecisionOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("accumulatorPrecision"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2980757-accumulatorprecision
func (c_ CNNConvolutionNode) SetAccumulatorPrecision(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccumulatorPrecision:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/3197822-trainingstyle
func (c_ CNNConvolutionNode) TrainingStyle() TrainingStyle get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("trainingStyle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/3197822-trainingstyle
func (c_ CNNConvolutionNode) SetTrainingStyle(value TrainingStyle get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTrainingStyle:"), value)
}







