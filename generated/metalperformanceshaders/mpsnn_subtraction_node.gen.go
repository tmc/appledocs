// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [SubtractionNode] class.
var (
	SubtractionNodeClass     _SubtractionNodeClass
	SubtractionNodeClassOnce sync.Once
)

func getSubtractionNodeClass() _SubtractionNodeClass {
	SubtractionNodeClassOnce.Do(func() {
		SubtractionNodeClass = _SubtractionNodeClass{objc.GetClass("MPSNNSubtractionNode")}
	})
	return SubtractionNodeClass
}

type _SubtractionNodeClass struct {
	class objc.Class
}





// An interface definition for the [SubtractionNode] class.
type ISubtractionNode interface {
	IBinaryArithmeticNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SubtractionNodeClass) Alloc() SubtractionNode {
	rv := objc.Send[SubtractionNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubtractionNodeClass) New() SubtractionNode {
	rv := objc.Send[SubtractionNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SubtractionNode) Init() SubtractionNode {
	rv := objc.Send[SubtractionNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SubtractionNode) Autorelease() SubtractionNode {
	rv := objc.Send[SubtractionNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubtractionNode creates a new SubtractionNode instance.
func NewSubtractionNode() SubtractionNode {
	return getSubtractionNodeClass().New()
}





// A representation of an subtraction operator.


// A representation of an subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionNode
type SubtractionNode struct {
	BinaryArithmeticNode
}

// SubtractionNodeFrom constructs a [SubtractionNode] from an unsafe.Pointer.
//
// A representation of an subtraction operator.
func SubtractionNodeFrom(ptr unsafe.Pointer) SubtractionNode {
	return SubtractionNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}































