// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [SubtractionGradientNode] class.
var (
	SubtractionGradientNodeClass     _SubtractionGradientNodeClass
	SubtractionGradientNodeClassOnce sync.Once
)

func getSubtractionGradientNodeClass() _SubtractionGradientNodeClass {
	SubtractionGradientNodeClassOnce.Do(func() {
		SubtractionGradientNodeClass = _SubtractionGradientNodeClass{objc.GetClass("MPSNNSubtractionGradientNode")}
	})
	return SubtractionGradientNodeClass
}

type _SubtractionGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [SubtractionGradientNode] class.
type ISubtractionGradientNode interface {
	IArithmeticGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SubtractionGradientNodeClass) Alloc() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubtractionGradientNodeClass) New() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SubtractionGradientNode) Init() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SubtractionGradientNode) Autorelease() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubtractionGradientNode creates a new SubtractionGradientNode instance.
func NewSubtractionGradientNode() SubtractionGradientNode {
	return getSubtractionGradientNodeClass().New()
}





// A representation of a gradient subtraction operator.


// A representation of a gradient subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionGradientNode
type SubtractionGradientNode struct {
	ArithmeticGradientNode
}

// SubtractionGradientNodeFrom constructs a [SubtractionGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient subtraction operator.
func SubtractionGradientNodeFrom(ptr unsafe.Pointer) SubtractionGradientNode {
	return SubtractionGradientNode{
		ArithmeticGradientNode: ArithmeticGradientNodeFrom(ptr),
	}
}































