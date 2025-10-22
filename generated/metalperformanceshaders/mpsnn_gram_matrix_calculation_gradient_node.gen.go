// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Alpha() float32
}

//
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

// Alloc allocates a new instance without initialization.
func (gc _GramMatrixCalculationGradientNodeClass) Alloc() GramMatrixCalculationGradientNode {
	rv := objc.Send[GramMatrixCalculationGradientNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGramMatrixCalculationGradientNode/alpha
func (g_ GramMatrixCalculationGradientNode) Alpha() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("alpha"))
	return rv
}



