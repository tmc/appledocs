// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AdditionNode] class.
var (
	AdditionNodeClass     _AdditionNodeClass
	AdditionNodeClassOnce sync.Once
)

func getAdditionNodeClass() _AdditionNodeClass {
	AdditionNodeClassOnce.Do(func() {
		AdditionNodeClass = _AdditionNodeClass{objc.GetClass("MPSNNAdditionNode")}
	})
	return AdditionNodeClass
}

type _AdditionNodeClass struct {
	class objc.Class
}





// An interface definition for the [AdditionNode] class.
type IAdditionNode interface {
	IBinaryArithmeticNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AdditionNodeClass) Alloc() AdditionNode {
	rv := objc.Send[AdditionNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdditionNodeClass) New() AdditionNode {
	rv := objc.Send[AdditionNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdditionNode) Init() AdditionNode {
	rv := objc.Send[AdditionNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdditionNode) Autorelease() AdditionNode {
	rv := objc.Send[AdditionNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdditionNode creates a new AdditionNode instance.
func NewAdditionNode() AdditionNode {
	return getAdditionNodeClass().New()
}





// A representation of an addition operator.


// A representation of an addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionNode
type AdditionNode struct {
	BinaryArithmeticNode
}

// AdditionNodeFrom constructs a [AdditionNode] from an unsafe.Pointer.
//
// A representation of an addition operator.
func AdditionNodeFrom(ptr unsafe.Pointer) AdditionNode {
	return AdditionNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}































