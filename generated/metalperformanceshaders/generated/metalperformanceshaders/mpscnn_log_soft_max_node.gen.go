// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLogSoftMaxNode */


/* debug [class_header]: Header for MPSCNNLogSoftMaxNode */
// The class instance for the [CNNLogSoftMaxNode] class.
var (
	CNNLogSoftMaxNodeClass     _CNNLogSoftMaxNodeClass
	CNNLogSoftMaxNodeClassOnce sync.Once
)

func getCNNLogSoftMaxNodeClass() _CNNLogSoftMaxNodeClass {
	CNNLogSoftMaxNodeClassOnce.Do(func() {
		CNNLogSoftMaxNodeClass = _CNNLogSoftMaxNodeClass{objc.GetClass("MPSCNNLogSoftMaxNode")}
	})
	return CNNLogSoftMaxNodeClass
}

type _CNNLogSoftMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLogSoftMaxNode */
// An interface definition for the [CNNLogSoftMaxNode] class.
type ICNNLogSoftMaxNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for CNNLogSoftMaxNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLogSoftMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLogSoftMaxNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNLogSoftMaxNodeClass) Alloc() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLogSoftMaxNodeClass) New() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLogSoftMaxNode) Init() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLogSoftMaxNode) Autorelease() CNNLogSoftMaxNode {
	rv := objc.Send[CNNLogSoftMaxNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLogSoftMaxNode creates a new CNNLogSoftMaxNode instance.
func NewCNNLogSoftMaxNode() CNNLogSoftMaxNode {
	return getCNNLogSoftMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLogSoftMaxNode */
// A representation of a logarithmic softmax filter kernel.


// A representation of a logarithmic softmax filter kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxNode
type CNNLogSoftMaxNode struct {
	FilterNode
}

// CNNLogSoftMaxNodeFrom constructs a [CNNLogSoftMaxNode] from an unsafe.Pointer.
//
// A representation of a logarithmic softmax filter kernel.
func CNNLogSoftMaxNodeFrom(ptr unsafe.Pointer) CNNLogSoftMaxNode {
	return CNNLogSoftMaxNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLogSoftMaxNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode/2866457-initwithsource
func NewCNNLogSoftMaxNodeWithSource(sourceNode IImageNode) CNNLogSoftMaxNode {
	instance := getCNNLogSoftMaxNodeClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLogSoftMaxNodeWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLogSoftMaxNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode/2866434-nodewithsource
func (cc _CNNLogSoftMaxNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLogSoftMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLogSoftMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLogSoftMaxNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLogSoftMaxNode */


