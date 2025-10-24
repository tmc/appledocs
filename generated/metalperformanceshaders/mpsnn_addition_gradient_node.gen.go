// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AdditionGradientNode] class.
var (
	AdditionGradientNodeClass     _AdditionGradientNodeClass
	AdditionGradientNodeClassOnce sync.Once
)

func getAdditionGradientNodeClass() _AdditionGradientNodeClass {
	AdditionGradientNodeClassOnce.Do(func() {
		AdditionGradientNodeClass = _AdditionGradientNodeClass{objc.GetClass("MPSNNAdditionGradientNode")}
	})
	return AdditionGradientNodeClass
}

type _AdditionGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [AdditionGradientNode] class.
type IAdditionGradientNode interface {
	IArithmeticGradientNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AdditionGradientNodeClass) Alloc() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdditionGradientNodeClass) New() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdditionGradientNode) Init() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdditionGradientNode) Autorelease() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdditionGradientNode creates a new AdditionGradientNode instance.
func NewAdditionGradientNode() AdditionGradientNode {
	return getAdditionGradientNodeClass().New()
}





// A representation of a gradient addition operator.


// A representation of a gradient addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionGradientNode
type AdditionGradientNode struct {
	ArithmeticGradientNode
}

// AdditionGradientNodeFrom constructs a [AdditionGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient addition operator.
func AdditionGradientNodeFrom(ptr unsafe.Pointer) AdditionGradientNode {
	return AdditionGradientNode{
		ArithmeticGradientNode: ArithmeticGradientNodeFrom(ptr),
	}
}































