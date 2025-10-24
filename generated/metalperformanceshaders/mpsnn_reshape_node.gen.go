// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ReshapeNode] class.
var (
	ReshapeNodeClass     _ReshapeNodeClass
	ReshapeNodeClassOnce sync.Once
)

func getReshapeNodeClass() _ReshapeNodeClass {
	ReshapeNodeClassOnce.Do(func() {
		ReshapeNodeClass = _ReshapeNodeClass{objc.GetClass("MPSNNReshapeNode")}
	})
	return ReshapeNodeClass
}

type _ReshapeNodeClass struct {
	class objc.Class
}





// An interface definition for the [ReshapeNode] class.
type IReshapeNode interface {
	IFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _ReshapeNodeClass) Alloc() ReshapeNode {
	rv := objc.Send[ReshapeNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReshapeNodeClass) New() ReshapeNode {
	rv := objc.Send[ReshapeNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeNode) Init() ReshapeNode {
	rv := objc.Send[ReshapeNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeNode) Autorelease() ReshapeNode {
	rv := objc.Send[ReshapeNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeNode creates a new ReshapeNode instance.
func NewReshapeNode() ReshapeNode {
	return getReshapeNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode
type ReshapeNode struct {
	FilterNode
}

// ReshapeNodeFrom constructs a [ReshapeNode] from an unsafe.Pointer.
func ReshapeNodeFrom(ptr unsafe.Pointer) ReshapeNode {
	return ReshapeNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode/3037420-initwithsource
func NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(source IImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) ReshapeNode {
	instance := getReshapeNodeClass().Alloc()
	rv := objc.Send[ReshapeNode](instance.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode/3037421-nodewithsource
func (rc _ReshapeNodeClass) NodeWithSourceResultWidthResultHeightResultFeatureChannels(source IImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	return rv
}






















