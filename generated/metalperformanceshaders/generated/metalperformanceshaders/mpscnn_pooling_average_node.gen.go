// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingAverageNode */


/* debug [class_header]: Header for MPSCNNPoolingAverageNode */
// The class instance for the [CNNPoolingAverageNode] class.
var (
	CNNPoolingAverageNodeClass     _CNNPoolingAverageNodeClass
	CNNPoolingAverageNodeClassOnce sync.Once
)

func getCNNPoolingAverageNodeClass() _CNNPoolingAverageNodeClass {
	CNNPoolingAverageNodeClassOnce.Do(func() {
		CNNPoolingAverageNodeClass = _CNNPoolingAverageNodeClass{objc.GetClass("MPSCNNPoolingAverageNode")}
	})
	return CNNPoolingAverageNodeClass
}

type _CNNPoolingAverageNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingAverageNode */
// An interface definition for the [CNNPoolingAverageNode] class.
type ICNNPoolingAverageNode interface {
	ICNNPoolingNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingAverageNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingAverageNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingAverageNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageNodeClass) Alloc() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageNodeClass) New() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverageNode) Init() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverageNode) Autorelease() CNNPoolingAverageNode {
	rv := objc.Send[CNNPoolingAverageNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverageNode creates a new CNNPoolingAverageNode instance.
func NewCNNPoolingAverageNode() CNNPoolingAverageNode {
	return getCNNPoolingAverageNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingAverageNode */
// A representation of an average pooling filter.


// A representation of an average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageNode
type CNNPoolingAverageNode struct {
	CNNPoolingNode
}

// CNNPoolingAverageNodeFrom constructs a [CNNPoolingAverageNode] from an unsafe.Pointer.
//
// A representation of an average pooling filter.
func CNNPoolingAverageNodeFrom(ptr unsafe.Pointer) CNNPoolingAverageNode {
	return CNNPoolingAverageNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingAverageNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingAverageNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingAverageNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingAverageNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingAverageNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingAverageNode */



