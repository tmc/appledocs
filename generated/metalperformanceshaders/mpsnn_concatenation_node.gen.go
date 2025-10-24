// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ConcatenationNode] class.
var (
	ConcatenationNodeClass     _ConcatenationNodeClass
	ConcatenationNodeClassOnce sync.Once
)

func getConcatenationNodeClass() _ConcatenationNodeClass {
	ConcatenationNodeClassOnce.Do(func() {
		ConcatenationNodeClass = _ConcatenationNodeClass{objc.GetClass("MPSNNConcatenationNode")}
	})
	return ConcatenationNodeClass
}

type _ConcatenationNodeClass struct {
	class objc.Class
}





// An interface definition for the [ConcatenationNode] class.
type IConcatenationNode interface {
	IFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ConcatenationNodeClass) Alloc() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ConcatenationNodeClass) New() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConcatenationNode) Init() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConcatenationNode) Autorelease() ConcatenationNode {
	rv := objc.Send[ConcatenationNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConcatenationNode creates a new ConcatenationNode instance.
func NewConcatenationNode() ConcatenationNode {
	return getConcatenationNodeClass().New()
}





// A representation of the results from one or more kernels.


// A representation of the results from one or more kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNConcatenationNode
type ConcatenationNode struct {
	FilterNode
}

// ConcatenationNodeFrom constructs a [ConcatenationNode] from an unsafe.Pointer.
//
// A representation of the results from one or more kernels.
func ConcatenationNodeFrom(ptr unsafe.Pointer) ConcatenationNode {
	return ConcatenationNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode/2866423-initwithsources
func NewConcatenationNodeWithSources(sourceNodes unsafe.Pointer) ConcatenationNode {
	instance := getConcatenationNodeClass().Alloc()
	rv := objc.Send[ConcatenationNode](instance.ID, objc.Sel("initWithSources:"), sourceNodes)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationnode/2866432-nodewithsources
func (cc _ConcatenationNodeClass) NodeWithSources(sourceNodes unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}






















