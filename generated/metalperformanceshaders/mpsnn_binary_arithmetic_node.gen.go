// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BinaryArithmeticNode] class.
var (
	BinaryArithmeticNodeClass     _BinaryArithmeticNodeClass
	BinaryArithmeticNodeClassOnce sync.Once
)

func getBinaryArithmeticNodeClass() _BinaryArithmeticNodeClass {
	BinaryArithmeticNodeClassOnce.Do(func() {
		BinaryArithmeticNodeClass = _BinaryArithmeticNodeClass{objc.GetClass("MPSNNBinaryArithmeticNode")}
	})
	return BinaryArithmeticNodeClass
}

type _BinaryArithmeticNodeClass struct {
	class objc.Class
}

// An interface definition for the [BinaryArithmeticNode] class.
type IBinaryArithmeticNode interface {
	IFilterNode
}

// Virtual base class for basic arithmetic nodes.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryArithmeticNode
type BinaryArithmeticNode struct {
	FilterNode
}

// BinaryArithmeticNodeFrom constructs a [BinaryArithmeticNode] from an unsafe.Pointer.
//
// Virtual base class for basic arithmetic nodes.
func BinaryArithmeticNodeFrom(ptr unsafe.Pointer) BinaryArithmeticNode {
	return BinaryArithmeticNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BinaryArithmeticNodeClass) Alloc() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BinaryArithmeticNodeClass) New() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryArithmeticNode) Init() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryArithmeticNode) Autorelease() BinaryArithmeticNode {
	rv := objc.Send[BinaryArithmeticNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryArithmeticNode creates a new BinaryArithmeticNode instance.
func NewBinaryArithmeticNode() BinaryArithmeticNode {
	return getBinaryArithmeticNodeClass().New()
}
