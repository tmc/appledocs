// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MultiplicationNode] class.
var (
	MultiplicationNodeClass     _MultiplicationNodeClass
	MultiplicationNodeClassOnce sync.Once
)

func getMultiplicationNodeClass() _MultiplicationNodeClass {
	MultiplicationNodeClassOnce.Do(func() {
		MultiplicationNodeClass = _MultiplicationNodeClass{objc.GetClass("MPSNNMultiplicationNode")}
	})
	return MultiplicationNodeClass
}

type _MultiplicationNodeClass struct {
	class objc.Class
}





// An interface definition for the [MultiplicationNode] class.
type IMultiplicationNode interface {
	IBinaryArithmeticNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MultiplicationNodeClass) Alloc() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiplicationNodeClass) New() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiplicationNode) Init() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiplicationNode) Autorelease() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiplicationNode creates a new MultiplicationNode instance.
func NewMultiplicationNode() MultiplicationNode {
	return getMultiplicationNodeClass().New()
}





// A representation of a multiplication operator.


// A representation of a multiplication operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiplicationNode
type MultiplicationNode struct {
	BinaryArithmeticNode
}

// MultiplicationNodeFrom constructs a [MultiplicationNode] from an unsafe.Pointer.
//
// A representation of a multiplication operator.
func MultiplicationNodeFrom(ptr unsafe.Pointer) MultiplicationNode {
	return MultiplicationNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}































