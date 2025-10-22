// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
)

// The class instance for the [UnaryReductionNode] class.
var (
	UnaryReductionNodeClass     _UnaryReductionNodeClass
	UnaryReductionNodeClassOnce sync.Once
)

func getUnaryReductionNodeClass() _UnaryReductionNodeClass {
	UnaryReductionNodeClassOnce.Do(func() {
		UnaryReductionNodeClass = _UnaryReductionNodeClass{objc.GetClass("MPSNNUnaryReductionNode")}
	})
	return UnaryReductionNodeClass
}

type _UnaryReductionNodeClass struct {
	class objc.Class
}

// An interface definition for the [UnaryReductionNode] class.
type IUnaryReductionNode interface {
	IFilterNode
	ClipRectSource() corelocation.Region
	SetClipRectSource(value corelocation.IRegion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode
type UnaryReductionNode struct {
	FilterNode
}

// UnaryReductionNodeFrom constructs a [UnaryReductionNode] from an unsafe.Pointer.
func UnaryReductionNodeFrom(ptr unsafe.Pointer) UnaryReductionNode {
	return UnaryReductionNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnaryReductionNodeClass) Alloc() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnaryReductionNodeClass) New() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnaryReductionNode) Init() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnaryReductionNode) Autorelease() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnaryReductionNode creates a new UnaryReductionNode instance.
func NewUnaryReductionNode() UnaryReductionNode {
	return getUnaryReductionNodeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode/init(source:)
func NewUnaryReductionNodeWithSource(sourceNode IMPSNNImageNode) UnaryReductionNode {
	instance := getUnaryReductionNodeClass().Alloc()
	rv := objc.Send[UnaryReductionNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/cliprectsource
func (u_ UnaryReductionNode) ClipRectSource() corelocation.Region {
	rv := objc.Send[corelocation.Region](u_.ID, objc.Sel("clipRectSource"))
	return rv
}


// SetClipRectSource sets the value of the clipRectSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/cliprectsource
func (u_ UnaryReductionNode) SetClipRectSource(value corelocation.IRegion) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRectSource:"), value)
}


