// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLossNode */


/* debug [class_header]: Header for MPSCNNLossNode */
// The class instance for the [CNNLossNode] class.
var (
	CNNLossNodeClass     _CNNLossNodeClass
	CNNLossNodeClassOnce sync.Once
)

func getCNNLossNodeClass() _CNNLossNodeClass {
	CNNLossNodeClassOnce.Do(func() {
		CNNLossNodeClass = _CNNLossNodeClass{objc.GetClass("MPSCNNLossNode")}
	})
	return CNNLossNodeClass
}

type _CNNLossNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLossNode */
// An interface definition for the [CNNLossNode] class.
type ICNNLossNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNLossNode */
	// properties:
	InputLabels() IMPSNNLabelsNode
	SetInputLabels(value IMPSNNLabelsNode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLossNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLossNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNLossNodeClass) Alloc() CNNLossNode {
	rv := objc.Send[CNNLossNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLossNodeClass) New() CNNLossNode {
	rv := objc.Send[CNNLossNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLossNode) Init() CNNLossNode {
	rv := objc.Send[CNNLossNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLossNode) Autorelease() CNNLossNode {
	rv := objc.Send[CNNLossNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLossNode creates a new CNNLossNode instance.
func NewCNNLossNode() CNNLossNode {
	return getCNNLossNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLossNode */
// A representation of a loss kernel.


// A representation of a loss kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossNode
type CNNLossNode struct {
	FilterNode
}

// CNNLossNodeFrom constructs a [CNNLossNode] from an unsafe.Pointer.
//
// A representation of a loss kernel.
func CNNLossNodeFrom(ptr unsafe.Pointer) CNNLossNode {
	return CNNLossNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951947-initwithsource
func NewCNNLossNodeWithSourceLossDescriptor(source IImageNode, descriptor ICNNLossDescriptor) CNNLossNode {
	instance := getCNNLossNodeClass().Alloc()
	rv := objc.Send[CNNLossNode](instance.ID, objc.Sel("initWithSource:lossDescriptor:"), source, descriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLossNodeWithSourceLossDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951956-nodewithsource
func (cc _CNNLossNodeClass) NodeWithSourceLossDescriptor(source IImageNode, descriptor ICNNLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:lossDescriptor:"), source, descriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceLossDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLossNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLossNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLossNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951942-inputlabels
func (c_ CNNLossNode) InputLabels() IMPSNNLabelsNode {
	rv := objc.Send[LabelsNode](c_.ID, objc.Sel("inputLabels"))
	return rv
}/* debug [instance_properties/getter]: inputLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951942-inputlabels
func (c_ CNNLossNode) SetInputLabels(value IMPSNNLabelsNode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputLabels:"), value)
}/* debug [instance_properties/setter]: inputLabels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLossNode */


