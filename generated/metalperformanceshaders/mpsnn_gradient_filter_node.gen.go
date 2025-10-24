// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GradientFilterNode] class.
var (
	GradientFilterNodeClass     _GradientFilterNodeClass
	GradientFilterNodeClassOnce sync.Once
)

func getGradientFilterNodeClass() _GradientFilterNodeClass {
	GradientFilterNodeClassOnce.Do(func() {
		GradientFilterNodeClass = _GradientFilterNodeClass{objc.GetClass("MPSNNGradientFilterNode")}
	})
	return GradientFilterNodeClass
}

type _GradientFilterNodeClass struct {
	class objc.Class
}

// An interface definition for the [GradientFilterNode] class.
type IGradientFilterNode interface {
	IFilterNode
	// properties:
	// methods:
}

// A representation of a gradient filter.


// A representation of a gradient filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientFilterNode
type GradientFilterNode struct {
	FilterNode
}

// GradientFilterNodeFrom constructs a [GradientFilterNode] from an unsafe.Pointer.
//
// A representation of a gradient filter.
func GradientFilterNodeFrom(ptr unsafe.Pointer) GradientFilterNode {
	return GradientFilterNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GradientFilterNodeClass) Alloc() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GradientFilterNodeClass) New() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GradientFilterNode) Init() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GradientFilterNode) Autorelease() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradientFilterNode creates a new GradientFilterNode instance.
func NewGradientFilterNode() GradientFilterNode {
	return getGradientFilterNodeClass().New()
}




