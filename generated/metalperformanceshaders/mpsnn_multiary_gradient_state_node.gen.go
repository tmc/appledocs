// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MultiaryGradientStateNode] class.
var (
	MultiaryGradientStateNodeClass     _MultiaryGradientStateNodeClass
	MultiaryGradientStateNodeClassOnce sync.Once
)

func getMultiaryGradientStateNodeClass() _MultiaryGradientStateNodeClass {
	MultiaryGradientStateNodeClassOnce.Do(func() {
		MultiaryGradientStateNodeClass = _MultiaryGradientStateNodeClass{objc.GetClass("MPSNNMultiaryGradientStateNode")}
	})
	return MultiaryGradientStateNodeClass
}

type _MultiaryGradientStateNodeClass struct {
	class objc.Class
}





// An interface definition for the [MultiaryGradientStateNode] class.
type IMultiaryGradientStateNode interface {
	IStateNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MultiaryGradientStateNodeClass) Alloc() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiaryGradientStateNodeClass) New() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiaryGradientStateNode) Init() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiaryGradientStateNode) Autorelease() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiaryGradientStateNode creates a new MultiaryGradientStateNode instance.
func NewMultiaryGradientStateNode() MultiaryGradientStateNode {
	return getMultiaryGradientStateNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientStateNode
type MultiaryGradientStateNode struct {
	StateNode
}

// MultiaryGradientStateNodeFrom constructs a [MultiaryGradientStateNode] from an unsafe.Pointer.
func MultiaryGradientStateNodeFrom(ptr unsafe.Pointer) MultiaryGradientStateNode {
	return MultiaryGradientStateNode{
		StateNode: StateNodeFrom(ptr),
	}
}































