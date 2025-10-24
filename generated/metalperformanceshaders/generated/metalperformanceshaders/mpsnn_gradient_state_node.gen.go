// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNGradientStateNode */


/* debug [class_header]: Header for MPSNNGradientStateNode */
// The class instance for the [GradientStateNode] class.
var (
	GradientStateNodeClass     _GradientStateNodeClass
	GradientStateNodeClassOnce sync.Once
)

func getGradientStateNodeClass() _GradientStateNodeClass {
	GradientStateNodeClassOnce.Do(func() {
		GradientStateNodeClass = _GradientStateNodeClass{objc.GetClass("MPSNNGradientStateNode")}
	})
	return GradientStateNodeClass
}

type _GradientStateNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GradientStateNode */
// An interface definition for the [GradientStateNode] class.
type IGradientStateNode interface {
	IStateNode
	
/* debug [class_interface_properties]: Properties for GradientStateNode */
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

	
/* debug [class_interface_methods]: Methods for GradientStateNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GradientStateNode */
// Alloc allocates a new instance without initialization.
func (gc _GradientStateNodeClass) Alloc() GradientStateNode {
	rv := objc.Send[GradientStateNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GradientStateNodeClass) New() GradientStateNode {
	rv := objc.Send[GradientStateNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GradientStateNode) Init() GradientStateNode {
	rv := objc.Send[GradientStateNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GradientStateNode) Autorelease() GradientStateNode {
	rv := objc.Send[GradientStateNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradientStateNode creates a new GradientStateNode instance.
func NewGradientStateNode() GradientStateNode {
	return getGradientStateNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GradientStateNode */
// A representation of the state created to record the properties of a gradient kernel at the time it was encoded.


// A representation of the state created to record the properties of a gradient kernel at the time it was encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientStateNode
type GradientStateNode struct {
	StateNode
}

// GradientStateNodeFrom constructs a [GradientStateNode] from an unsafe.Pointer.
//
// A representation of the state created to record the properties of a gradient kernel at the time it was encoded.
func GradientStateNodeFrom(ptr unsafe.Pointer) GradientStateNode {
	return GradientStateNode{
		StateNode: StateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GradientStateNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GradientStateNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GradientStateNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GradientStateNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GradientStateNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (g_ GradientStateNode) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/label
func (g_ GradientStateNode) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (g_ GradientStateNode) PaddingPolicy() Padding /* not a class type */ {
	rv := objc.Send[Padding](g_.ID, objc.Sel("paddingPolicy"))
	return rv
}/* debug [instance_properties/getter]: paddingPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/paddingpolicy
func (g_ GradientStateNode) SetPaddingPolicy(value Padding /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaddingPolicy:"), value)
}/* debug [instance_properties/setter]: paddingPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (g_ GradientStateNode) ResultImage() IMPSNNImageNode {
	rv := objc.Send[ImageNode](g_.ID, objc.Sel("resultImage"))
	return rv
}/* debug [instance_properties/getter]: resultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultimage
func (g_ GradientStateNode) SetResultImage(value IMPSNNImageNode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultImage:"), value)
}/* debug [instance_properties/setter]: resultImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (g_ GradientStateNode) ResultState() IMPSNNStateNode {
	rv := objc.Send[StateNode](g_.ID, objc.Sel("resultState"))
	return rv
}/* debug [instance_properties/getter]: resultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstate
func (g_ GradientStateNode) SetResultState(value IMPSNNStateNode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultState:"), value)
}/* debug [instance_properties/setter]: resultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (g_ GradientStateNode) ResultStates() IMPSNNStateNode {
	rv := objc.Send[StateNode](g_.ID, objc.Sel("resultStates"))
	return rv
}/* debug [instance_properties/getter]: resultStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/resultstates
func (g_ GradientStateNode) SetResultStates(value IMPSNNStateNode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResultStates:"), value)
}/* debug [instance_properties/setter]: resultStates */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNGradientStateNode */



