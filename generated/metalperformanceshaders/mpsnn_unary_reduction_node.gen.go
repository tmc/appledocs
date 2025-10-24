// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [UnaryReductionNode] class.
var (
	UnaryReductionNodeClass     _UnaryReductionNodeClass
	UnaryReductionNodeClassOnce sync.Once
)

func getUnaryReductionNodeClass() _UnaryReductionNodeClass {
	UnaryReductionNodeClassOnce.Do(func() {
		UnaryReductionNodeClass = _UnaryReductionNodeClass{objc.GetClass("MPSNNUnaryReductionNode")}
	})
	return UnaryReductionNodeClass
}

type _UnaryReductionNodeClass struct {
	class objc.Class
}





// An interface definition for the [UnaryReductionNode] class.
type IUnaryReductionNode interface {
	IFilterNode
	

	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (uc _UnaryReductionNodeClass) Alloc() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnaryReductionNodeClass) New() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnaryReductionNode) Init() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnaryReductionNode) Autorelease() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnaryReductionNode creates a new UnaryReductionNode instance.
func NewUnaryReductionNode() UnaryReductionNode {
	return getUnaryReductionNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode
type UnaryReductionNode struct {
	FilterNode
}

// UnaryReductionNodeFrom constructs a [UnaryReductionNode] from an unsafe.Pointer.
func UnaryReductionNodeFrom(ptr unsafe.Pointer) UnaryReductionNode {
	return UnaryReductionNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource
func NewUnaryReductionNodeWithSource(sourceNode IImageNode) UnaryReductionNode {
	instance := getUnaryReductionNodeClass().Alloc()
	rv := objc.Send[UnaryReductionNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource
func (uc _UnaryReductionNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037423-cliprectsource
func (u_ UnaryReductionNode) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037423-cliprectsource
func (u_ UnaryReductionNode) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRectSource:"), value)
}







