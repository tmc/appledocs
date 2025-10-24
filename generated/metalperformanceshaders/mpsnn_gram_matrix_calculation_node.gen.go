// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [GramMatrixCalculationNode] class.
var (
	GramMatrixCalculationNodeClass     _GramMatrixCalculationNodeClass
	GramMatrixCalculationNodeClassOnce sync.Once
)

func getGramMatrixCalculationNodeClass() _GramMatrixCalculationNodeClass {
	GramMatrixCalculationNodeClassOnce.Do(func() {
		GramMatrixCalculationNodeClass = _GramMatrixCalculationNodeClass{objc.GetClass("MPSNNGramMatrixCalculationNode")}
	})
	return GramMatrixCalculationNodeClass
}

type _GramMatrixCalculationNodeClass struct {
	class objc.Class
}





// An interface definition for the [GramMatrixCalculationNode] class.
type IGramMatrixCalculationNode interface {
	IFilterNode
	

	// properties:
	Alpha() objectivec.IObject
	SetAlpha(value objectivec.IObject)
	PropertyCallBack() GramMatrixCallback get set /* not a class type */
	SetPropertyCallBack(value GramMatrixCallback get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationNodeClass) Alloc() GramMatrixCalculationNode {
	rv := objc.Send[GramMatrixCalculationNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GramMatrixCalculationNodeClass) New() GramMatrixCalculationNode {
	rv := objc.Send[GramMatrixCalculationNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GramMatrixCalculationNode) Init() GramMatrixCalculationNode {
	rv := objc.Send[GramMatrixCalculationNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GramMatrixCalculationNode) Autorelease() GramMatrixCalculationNode {
	rv := objc.Send[GramMatrixCalculationNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGramMatrixCalculationNode creates a new GramMatrixCalculationNode instance.
func NewGramMatrixCalculationNode() GramMatrixCalculationNode {
	return getGramMatrixCalculationNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode
type GramMatrixCalculationNode struct {
	FilterNode
}

// GramMatrixCalculationNodeFrom constructs a [GramMatrixCalculationNode] from an unsafe.Pointer.
func GramMatrixCalculationNodeFrom(ptr unsafe.Pointer) GramMatrixCalculationNode {
	return GramMatrixCalculationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114095-initwithsource
func NewGramMatrixCalculationNodeWithSource(sourceNode IImageNode) GramMatrixCalculationNode {
	instance := getGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[GramMatrixCalculationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114096-initwithsource
func NewGramMatrixCalculationNodeWithSourceAlpha(sourceNode IImageNode, alpha float32) GramMatrixCalculationNode {
	instance := getGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[GramMatrixCalculationNode](instance.ID, objc.Sel("initWithSource:alpha:"), sourceNode, alpha)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114097-nodewithsource
func (gc _GramMatrixCalculationNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114098-nodewithsource
func (gc _GramMatrixCalculationNodeClass) NodeWithSourceAlpha(sourceNode IImageNode, alpha float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("nodeWithSource:alpha:"), sourceNode, alpha)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114094-alpha
func (g_ GramMatrixCalculationNode) Alpha() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("alpha"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114094-alpha
func (g_ GramMatrixCalculationNode) SetAlpha(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAlpha:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3131844-propertycallback
func (g_ GramMatrixCalculationNode) PropertyCallBack() GramMatrixCallback get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3131844-propertycallback
func (g_ GramMatrixCalculationNode) SetPropertyCallBack(value GramMatrixCallback get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPropertyCallBack:"), value)
}







