// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [DivisionNode] class.
var (
	DivisionNodeClass     _DivisionNodeClass
	DivisionNodeClassOnce sync.Once
)

func getDivisionNodeClass() _DivisionNodeClass {
	DivisionNodeClassOnce.Do(func() {
		DivisionNodeClass = _DivisionNodeClass{objc.GetClass("MPSNNDivisionNode")}
	})
	return DivisionNodeClass
}

type _DivisionNodeClass struct {
	class objc.Class
}





// An interface definition for the [DivisionNode] class.
type IDivisionNode interface {
	IBinaryArithmeticNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DivisionNodeClass) Alloc() DivisionNode {
	rv := objc.Send[DivisionNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DivisionNodeClass) New() DivisionNode {
	rv := objc.Send[DivisionNode](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DivisionNode) Init() DivisionNode {
	rv := objc.Send[DivisionNode](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DivisionNode) Autorelease() DivisionNode {
	rv := objc.Send[DivisionNode](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDivisionNode creates a new DivisionNode instance.
func NewDivisionNode() DivisionNode {
	return getDivisionNodeClass().New()
}





// A representation of a division operator.


// A representation of a division operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDivisionNode
type DivisionNode struct {
	BinaryArithmeticNode
}

// DivisionNodeFrom constructs a [DivisionNode] from an unsafe.Pointer.
//
// A representation of a division operator.
func DivisionNodeFrom(ptr unsafe.Pointer) DivisionNode {
	return DivisionNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}































