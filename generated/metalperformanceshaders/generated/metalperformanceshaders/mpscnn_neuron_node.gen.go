// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronNode */


/* debug [class_header]: Header for MPSCNNNeuronNode */
// The class instance for the [CNNNeuronNode] class.
var (
	CNNNeuronNodeClass     _CNNNeuronNodeClass
	CNNNeuronNodeClassOnce sync.Once
)

func getCNNNeuronNodeClass() _CNNNeuronNodeClass {
	CNNNeuronNodeClassOnce.Do(func() {
		CNNNeuronNodeClass = _CNNNeuronNodeClass{objc.GetClass("MPSCNNNeuronNode")}
	})
	return CNNNeuronNodeClass
}

type _CNNNeuronNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronNode */
// An interface definition for the [CNNNeuronNode] class.
type ICNNNeuronNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNNeuronNode */
	// properties:
	A() objectivec.IObject
	SetA(value objectivec.IObject)
	B() objectivec.IObject
	SetB(value objectivec.IObject)
	C() objectivec.IObject
	SetC(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronNodeClass) Alloc() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronNodeClass) New() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronNode) Init() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronNode) Autorelease() CNNNeuronNode {
	rv := objc.Send[CNNNeuronNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronNode creates a new CNNNeuronNode instance.
func NewCNNNeuronNode() CNNNeuronNode {
	return getCNNNeuronNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronNode */
// The virtual base class for MPS CNN neuron nodes.


// The virtual base class for MPS CNN neuron nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronNode
type CNNNeuronNode struct {
	FilterNode
}

// CNNNeuronNodeFrom constructs a [CNNNeuronNode] from an unsafe.Pointer.
//
// The virtual base class for MPS CNN neuron nodes.
func CNNNeuronNodeFrom(ptr unsafe.Pointer) CNNNeuronNode {
	return CNNNeuronNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource
func (cc _CNNNeuronNodeClass) NodeWithSourceDescriptor(sourceNode IImageNode, descriptor INeuronDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2921459-a
func (c_ CNNNeuronNode) A() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2921459-a
func (c_ CNNNeuronNode) SetA(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2921461-b
func (c_ CNNNeuronNode) B() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2921461-b
func (c_ CNNNeuronNode) SetB(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setB:"), value)
}/* debug [instance_properties/setter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2935553-c
func (c_ CNNNeuronNode) C() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2935553-c
func (c_ CNNNeuronNode) SetC(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setC:"), value)
}/* debug [instance_properties/setter]: c */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronNode */



