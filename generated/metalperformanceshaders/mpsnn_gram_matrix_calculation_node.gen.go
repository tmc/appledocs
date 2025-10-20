// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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

// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationNodeClass) Alloc() GramMatrixCalculationNode {
	rv := objc.Send[GramMatrixCalculationNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/init(source:)
func NewGramMatrixCalculationNodeWithSource(sourceNode unsafe.Pointer) GramMatrixCalculationNode {
	instance := getGramMatrixCalculationNodeClass().Alloc()
	rv := objc.Send[GramMatrixCalculationNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/nodeWithSource:alpha:
func (gc _GramMatrixCalculationNodeClass) NodeWithSourceAlpha(sourceNode unsafe.Pointer, alpha unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("nodeWithSource:alpha:"), sourceNode, alpha)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/propertyCallBack
func (g_ GramMatrixCalculationNode) PropertyCallBack() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("propertyCallBack"))
	return rv
}


// SetPropertyCallBack sets the value of the propertyCallBack property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationNode/propertyCallBack
func (g_ GramMatrixCalculationNode) SetPropertyCallBack(value objc.ID) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPropertyCallBack:"), value)
}

