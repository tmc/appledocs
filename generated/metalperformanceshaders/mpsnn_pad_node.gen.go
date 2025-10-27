// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PadNode] class.
var (
	PadNodeClass     _PadNodeClass
	PadNodeClassOnce sync.Once
)

func getPadNodeClass() _PadNodeClass {
	PadNodeClassOnce.Do(func() {
		PadNodeClass = _PadNodeClass{objc.GetClass("MPSNNPadNode")}
	})
	return PadNodeClass
}

type _PadNodeClass struct {
	class objc.Class
}





// An interface definition for the [PadNode] class.
type IPadNode interface {
	IFilterNode
	

	// properties:
	FillValue() objectivec.IObject
	SetFillValue(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PadNodeClass) Alloc() PadNode {
	rv := objc.Send[PadNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PadNodeClass) New() PadNode {
	rv := objc.Send[PadNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PadNode) Init() PadNode {
	rv := objc.Send[PadNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PadNode) Autorelease() PadNode {
	rv := objc.Send[PadNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPadNode creates a new PadNode instance.
func NewPadNode() PadNode {
	return getPadNodeClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNPadNode
type PadNode struct {
	FilterNode
}

// PadNodeFrom constructs a [PadNode] from an unsafe.Pointer.
func PadNodeFrom(ptr unsafe.Pointer) PadNode {
	return PadNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037395-initwithsource
func NewPadNodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IImageNode, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, edgeMode ImageEdgeMode) PadNode {
	instance := getPadNodeClass().Alloc()
	rv := objc.Send[PadNode](instance.ID, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037396-nodewithsource
func (pc _PadNodeClass) NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IImageNode, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, edgeMode ImageEdgeMode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037394-fillvalue
func (p_ PadNode) FillValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("fillValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037394-fillvalue
func (p_ PadNode) SetFillValue(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFillValue:"), value)
}







