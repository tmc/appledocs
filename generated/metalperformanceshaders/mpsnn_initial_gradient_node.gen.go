// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InitialGradientNode] class.
var (
	InitialGradientNodeClass     _InitialGradientNodeClass
	InitialGradientNodeClassOnce sync.Once
)

func getInitialGradientNodeClass() _InitialGradientNodeClass {
	InitialGradientNodeClassOnce.Do(func() {
		InitialGradientNodeClass = _InitialGradientNodeClass{objc.GetClass("MPSNNInitialGradientNode")}
	})
	return InitialGradientNodeClass
}

type _InitialGradientNodeClass struct {
	class objc.Class
}

// An interface definition for the [InitialGradientNode] class.
type IInitialGradientNode interface {
	IFilterNode
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode
type InitialGradientNode struct {
	FilterNode
}

// InitialGradientNodeFrom constructs a [InitialGradientNode] from an unsafe.Pointer.
func InitialGradientNodeFrom(ptr unsafe.Pointer) InitialGradientNode {
	return InitialGradientNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InitialGradientNodeClass) Alloc() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InitialGradientNodeClass) New() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InitialGradientNode) Init() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InitialGradientNode) Autorelease() InitialGradientNode {
	rv := objc.Send[InitialGradientNode](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInitialGradientNode creates a new InitialGradientNode instance.
func NewInitialGradientNode() InitialGradientNode {
	return getInitialGradientNodeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode/init(source:)
func NewInitialGradientNodeWithSource(source IMPSNNImageNode) InitialGradientNode {
	instance := getInitialGradientNodeClass().Alloc()
	rv := objc.Send[InitialGradientNode](instance.ID, objc.Sel("initWithSource:"), source)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNInitialGradientNode/nodeWithSource:
func (ic _InitialGradientNodeClass) NodeWithSource(source IMPSNNImageNode) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("nodeWithSource:"), source)
	return rv
}


