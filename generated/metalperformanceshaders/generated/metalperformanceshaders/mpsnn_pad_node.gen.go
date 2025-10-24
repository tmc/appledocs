// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNPadNode */


/* debug [class_header]: Header for MPSNNPadNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PadNode */
// An interface definition for the [PadNode] class.
type IPadNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for PadNode */
	// properties:
	FillValue() objectivec.IObject
	SetFillValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PadNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PadNode */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PadNode */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PadNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037395-initwithsource
func NewPadNodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IImageNode, paddingSizeBefore objc.IObject /* cross-framework: MPSImageCoordinate */, paddingSizeAfter objc.IObject /* cross-framework: MPSImageCoordinate */, edgeMode ImageEdgeMode) PadNode {
	instance := getPadNodeClass().Alloc()
	rv := objc.Send[PadNode](instance.ID, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPadNodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PadNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037396-nodewithsource
func (pc _PadNodeClass) NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source IImageNode, paddingSizeBefore objc.IObject /* cross-framework: MPSImageCoordinate */, paddingSizeAfter objc.IObject /* cross-framework: MPSImageCoordinate */, edgeMode ImageEdgeMode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PadNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PadNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PadNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037394-fillvalue
func (p_ PadNode) FillValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("fillValue"))
	return rv
}/* debug [instance_properties/getter]: fillValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037394-fillvalue
func (p_ PadNode) SetFillValue(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFillValue:"), value)
}/* debug [instance_properties/setter]: fillValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNPadNode */


