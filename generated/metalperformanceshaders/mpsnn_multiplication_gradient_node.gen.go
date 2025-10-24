// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MultiplicationGradientNode] class.
var (
	MultiplicationGradientNodeClass     _MultiplicationGradientNodeClass
	MultiplicationGradientNodeClassOnce sync.Once
)

func getMultiplicationGradientNodeClass() _MultiplicationGradientNodeClass {
	MultiplicationGradientNodeClassOnce.Do(func() {
		MultiplicationGradientNodeClass = _MultiplicationGradientNodeClass{objc.GetClass("MPSNNMultiplicationGradientNode")}
	})
	return MultiplicationGradientNodeClass
}

type _MultiplicationGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [MultiplicationGradientNode] class.
type IMultiplicationGradientNode interface {
	IArithmeticGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MultiplicationGradientNodeClass) Alloc() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiplicationGradientNodeClass) New() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiplicationGradientNode) Init() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiplicationGradientNode) Autorelease() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiplicationGradientNode creates a new MultiplicationGradientNode instance.
func NewMultiplicationGradientNode() MultiplicationGradientNode {
	return getMultiplicationGradientNodeClass().New()
}





// A representation of a gradient multiplication operator.


// A representation of a gradient multiplication operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiplicationGradientNode
type MultiplicationGradientNode struct {
	ArithmeticGradientNode
}

// MultiplicationGradientNodeFrom constructs a [MultiplicationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient multiplication operator.
func MultiplicationGradientNodeFrom(ptr unsafe.Pointer) MultiplicationGradientNode {
	return MultiplicationGradientNode{
		ArithmeticGradientNode: ArithmeticGradientNodeFrom(ptr),
	}
}































