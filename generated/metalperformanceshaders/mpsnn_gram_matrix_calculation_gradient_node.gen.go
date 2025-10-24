// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [GramMatrixCalculationGradientNode] class.
var (
	GramMatrixCalculationGradientNodeClass     _GramMatrixCalculationGradientNodeClass
	GramMatrixCalculationGradientNodeClassOnce sync.Once
)

func getGramMatrixCalculationGradientNodeClass() _GramMatrixCalculationGradientNodeClass {
	GramMatrixCalculationGradientNodeClassOnce.Do(func() {
		GramMatrixCalculationGradientNodeClass = _GramMatrixCalculationGradientNodeClass{objc.GetClass("MPSNNGramMatrixCalculationGradientNode")}
	})
	return GramMatrixCalculationGradientNodeClass
}

type _GramMatrixCalculationGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [GramMatrixCalculationGradientNode] class.
type IGramMatrixCalculationGradientNode interface {
	IGradientFilterNode
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationGradientNodeClass) Alloc() GramMatrixCalculationGradientNode {
	rv := objc.Send[GramMatrixCalculationGradientNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GramMatrixCalculationGradientNodeClass) New() GramMatrixCalculationGradientNode {
	rv := objc.Send[GramMatrixCalculationGradientNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GramMatrixCalculationGradientNode) Init() GramMatrixCalculationGradientNode {
	rv := objc.Send[GramMatrixCalculationGradientNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GramMatrixCalculationGradientNode) Autorelease() GramMatrixCalculationGradientNode {
	rv := objc.Send[GramMatrixCalculationGradientNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGramMatrixCalculationGradientNode creates a new GramMatrixCalculationGradientNode instance.
func NewGramMatrixCalculationGradientNode() GramMatrixCalculationGradientNode {
	return getGramMatrixCalculationGradientNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode
type GramMatrixCalculationGradientNode struct {
	GradientFilterNode
}

// GramMatrixCalculationGradientNodeFrom constructs a [GramMatrixCalculationGradientNode] from an unsafe.Pointer.
func GramMatrixCalculationGradientNodeFrom(ptr unsafe.Pointer) GramMatrixCalculationGradientNode {
	return GramMatrixCalculationGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114089-initwithsourcegradient
func NewGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) GramMatrixCalculationGradientNode {
	instance := getGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[GramMatrixCalculationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114090-initwithsourcegradient
func NewGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientStateAlpha(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, alpha float32) GramMatrixCalculationGradientNode {
	instance := getGramMatrixCalculationGradientNodeClass().Alloc()
	rv := objc.Send[GramMatrixCalculationGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:alpha:"), sourceGradient, sourceImage, gradientState, alpha)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114091-nodewithsourcegradient
func (gc _GramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), sourceGradient, sourceImage, gradientState)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114092-nodewithsourcegradient
func (gc _GramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateAlpha(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, alpha float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:alpha:"), sourceGradient, sourceImage, gradientState, alpha)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114088-alpha
func (g_ GramMatrixCalculationGradientNode) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114088-alpha
func (g_ GramMatrixCalculationGradientNode) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAlpha:"), value)
}







