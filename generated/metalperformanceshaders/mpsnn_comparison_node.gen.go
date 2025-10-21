// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComparisonNode] class.
var (
	ComparisonNodeClass     _ComparisonNodeClass
	ComparisonNodeClassOnce sync.Once
)

func getComparisonNodeClass() _ComparisonNodeClass {
	ComparisonNodeClassOnce.Do(func() {
		ComparisonNodeClass = _ComparisonNodeClass{objc.GetClass("MPSNNComparisonNode")}
	})
	return ComparisonNodeClass
}

type _ComparisonNodeClass struct {
	class objc.Class
}

// An interface definition for the [ComparisonNode] class.
type IComparisonNode interface {
	IBinaryArithmeticNode
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode
type ComparisonNode struct {
	BinaryArithmeticNode
}

// ComparisonNodeFrom constructs a [ComparisonNode] from an unsafe.Pointer.
func ComparisonNodeFrom(ptr unsafe.Pointer) ComparisonNode {
	return ComparisonNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ComparisonNodeClass) Alloc() ComparisonNode {
	rv := objc.Send[ComparisonNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComparisonNodeClass) New() ComparisonNode {
	rv := objc.Send[ComparisonNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComparisonNode) Init() ComparisonNode {
	rv := objc.Send[ComparisonNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComparisonNode) Autorelease() ComparisonNode {
	rv := objc.Send[ComparisonNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComparisonNode creates a new ComparisonNode instance.
func NewComparisonNode() ComparisonNode {
	return getComparisonNodeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode/comparisonType
func (c_ ComparisonNode) ComparisonType() ComparisonType {
	rv := objc.Send[ComparisonType](c_.ID, objc.Sel("comparisonType"))
	return rv
}


// SetComparisonType sets the value of the comparisonType property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode/comparisonType
func (c_ ComparisonNode) SetComparisonType(value ComparisonType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComparisonType:"), value)
}



