// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ReshapeNode] class.
var (
	ReshapeNodeClass     _ReshapeNodeClass
	ReshapeNodeClassOnce sync.Once
)

func getReshapeNodeClass() _ReshapeNodeClass {
	ReshapeNodeClassOnce.Do(func() {
		ReshapeNodeClass = _ReshapeNodeClass{objc.GetClass("MPSNNReshapeNode")}
	})
	return ReshapeNodeClass
}

type _ReshapeNodeClass struct {
	class objc.Class
}

// An interface definition for the [ReshapeNode] class.
type IReshapeNode interface {
	IFilterNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode
type ReshapeNode struct {
	FilterNode
}

// ReshapeNodeFrom constructs a [ReshapeNode] from an unsafe.Pointer.
func ReshapeNodeFrom(ptr unsafe.Pointer) ReshapeNode {
	return ReshapeNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _ReshapeNodeClass) Alloc() ReshapeNode {
	rv := objc.Send[ReshapeNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReshapeNodeClass) New() ReshapeNode {
	rv := objc.Send[ReshapeNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeNode) Init() ReshapeNode {
	rv := objc.Send[ReshapeNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeNode) Autorelease() ReshapeNode {
	rv := objc.Send[ReshapeNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeNode creates a new ReshapeNode instance.
func NewReshapeNode() ReshapeNode {
	return getReshapeNodeClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode/init(source:resultWidth:resultHeight:resultFeatureChannels:)
func NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(source unsafe.Pointer, resultWidth uint, resultHeight uint, resultFeatureChannels uint) ReshapeNode {
	instance := getReshapeNodeClass().Alloc()
	rv := objc.Send[ReshapeNode](instance.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	rv.Autorelease()
	return rv
}
