// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNBinaryGradientStateNode */


/* debug [class_header]: Header for MPSNNBinaryGradientStateNode */
// The class instance for the [BinaryGradientStateNode] class.
var (
	BinaryGradientStateNodeClass     _BinaryGradientStateNodeClass
	BinaryGradientStateNodeClassOnce sync.Once
)

func getBinaryGradientStateNodeClass() _BinaryGradientStateNodeClass {
	BinaryGradientStateNodeClassOnce.Do(func() {
		BinaryGradientStateNodeClass = _BinaryGradientStateNodeClass{objc.GetClass("MPSNNBinaryGradientStateNode")}
	})
	return BinaryGradientStateNodeClass
}

type _BinaryGradientStateNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BinaryGradientStateNode */
// An interface definition for the [BinaryGradientStateNode] class.
type IBinaryGradientStateNode interface {
	IStateNode
	
/* debug [class_interface_properties]: Properties for BinaryGradientStateNode */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	PaddingPolicy() Padding /* not a class type */
	SetPaddingPolicy(value Padding /* not a class type */)
	ResultImage() IMPSNNImageNode
	SetResultImage(value IMPSNNImageNode)
	ResultState() IMPSNNStateNode
	SetResultState(value IMPSNNStateNode)
	ResultStates() IMPSNNStateNode
	SetResultStates(value IMPSNNStateNode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BinaryGradientStateNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BinaryGradientStateNode */
// Alloc allocates a new instance without initialization.
func (bc _BinaryGradientStateNodeClass) Alloc() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BinaryGradientStateNodeClass) New() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryGradientStateNode) Init() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryGradientStateNode) Autorelease() BinaryGradientStateNode {
	rv := objc.Send[BinaryGradientStateNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryGradientStateNode creates a new BinaryGradientStateNode instance.
func NewBinaryGradientStateNode() BinaryGradientStateNode {
	return getBinaryGradientStateNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BinaryGradientStateNode */
// A representation of the state created to record the properties of a binary gradient kernel.


// A representation of the state created to record the properties of a binary gradient kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientStateNode
type BinaryGradientStateNode struct {
	StateNode
}

// BinaryGradientStateNodeFrom constructs a [BinaryGradientStateNode] from an unsafe.Pointer.
//
// A representation of the state created to record the properties of a binary gradient kernel.
func BinaryGradientStateNodeFrom(ptr unsafe.Pointer) BinaryGradientStateNode {
	return BinaryGradientStateNode{
		StateNode: StateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BinaryGradientStateNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BinaryGradientStateNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BinaryGradientStateNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BinaryGradientStateNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BinaryGradientStateNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (b_ BinaryGradientStateNode) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (b_ BinaryGradientStateNode) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (b_ BinaryGradientStateNode) PaddingPolicy() Padding /* not a class type */ {
	rv := objc.Send[Padding](b_.ID, objc.Sel("paddingPolicy"))
	return rv
}/* debug [instance_properties/getter]: paddingPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (b_ BinaryGradientStateNode) SetPaddingPolicy(value Padding /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPaddingPolicy:"), value)
}/* debug [instance_properties/setter]: paddingPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (b_ BinaryGradientStateNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[ImageNode](b_.ID, objc.Sel("resultImage"))
	return rv
}/* debug [instance_properties/getter]: resultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (b_ BinaryGradientStateNode) SetResultImage(value IMPSNNImageNode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultImage:"), value)
}/* debug [instance_properties/setter]: resultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (b_ BinaryGradientStateNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[StateNode](b_.ID, objc.Sel("resultState"))
	return rv
}/* debug [instance_properties/getter]: resultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (b_ BinaryGradientStateNode) SetResultState(value IMPSNNStateNode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultState:"), value)
}/* debug [instance_properties/setter]: resultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (b_ BinaryGradientStateNode) ResultStates() IMPSNNStateNode {
	rv := objc.Send[StateNode](b_.ID, objc.Sel("resultStates"))
	return rv
}/* debug [instance_properties/getter]: resultStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (b_ BinaryGradientStateNode) SetResultStates(value IMPSNNStateNode) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResultStates:"), value)
}/* debug [instance_properties/setter]: resultStates */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNBinaryGradientStateNode */



