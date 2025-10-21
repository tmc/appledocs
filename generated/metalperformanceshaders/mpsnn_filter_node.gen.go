// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilterNode] class.
var (
	FilterNodeClass     _FilterNodeClass
	FilterNodeClassOnce sync.Once
)

func getFilterNodeClass() _FilterNodeClass {
	FilterNodeClassOnce.Do(func() {
		FilterNodeClass = _FilterNodeClass{objc.GetClass("MPSNNFilterNode")}
	})
	return FilterNodeClass
}

type _FilterNodeClass struct {
	class objc.Class
}

// An interface definition for the [FilterNode] class.
type IFilterNode interface {
	objectivec.IObject
}

// A placeholder node denoting a neural network filter stage.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNFilterNode
type FilterNode struct {
	objectivec.Object
}

// FilterNodeFrom constructs a [FilterNode] from an unsafe.Pointer.
//
// A placeholder node denoting a neural network filter stage.
func FilterNodeFrom(ptr unsafe.Pointer) FilterNode {
	return FilterNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilterNodeClass) Alloc() FilterNode {
	rv := objc.Send[FilterNode](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilterNodeClass) New() FilterNode {
	rv := objc.Send[FilterNode](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilterNode) Init() FilterNode {
	rv := objc.Send[FilterNode](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilterNode) Autorelease() FilterNode {
	rv := objc.Send[FilterNode](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilterNode creates a new FilterNode instance.
func NewFilterNode() FilterNode {
	return getFilterNodeClass().New()
}




