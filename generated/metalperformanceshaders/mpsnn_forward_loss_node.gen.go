// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	GradientFiltersWithSources(sourceGradient unsafe.Pointer) []LossGradientNode
}

//
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

// Alloc allocates a new instance without initialization.
func (fc _ForwardLossNodeClass) Alloc() ForwardLossNode {
	rv := objc.Send[ForwardLossNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLossNode/gradientFilters(withSources:)
func (f_ ForwardLossNode) GradientFiltersWithSources(sourceGradient unsafe.Pointer) []LossGradientNode {
	rv := objc.Send[[]LossGradientNode](f_.ID, objc.Sel("gradientFiltersWithSources:"), sourceGradient)
	return rv
}



