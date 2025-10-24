// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingMaxNode */


/* debug [class_header]: Header for MPSCNNPoolingMaxNode */
// The class instance for the [CNNPoolingMaxNode] class.
var (
	CNNPoolingMaxNodeClass     _CNNPoolingMaxNodeClass
	CNNPoolingMaxNodeClassOnce sync.Once
)

func getCNNPoolingMaxNodeClass() _CNNPoolingMaxNodeClass {
	CNNPoolingMaxNodeClassOnce.Do(func() {
		CNNPoolingMaxNodeClass = _CNNPoolingMaxNodeClass{objc.GetClass("MPSCNNPoolingMaxNode")}
	})
	return CNNPoolingMaxNodeClass
}

type _CNNPoolingMaxNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingMaxNode */
// An interface definition for the [CNNPoolingMaxNode] class.
type ICNNPoolingMaxNode interface {
	ICNNPoolingNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingMaxNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingMaxNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingMaxNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxNodeClass) Alloc() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxNodeClass) New() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMaxNode) Init() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMaxNode) Autorelease() CNNPoolingMaxNode {
	rv := objc.Send[CNNPoolingMaxNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMaxNode creates a new CNNPoolingMaxNode instance.
func NewCNNPoolingMaxNode() CNNPoolingMaxNode {
	return getCNNPoolingMaxNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingMaxNode */
// A representation of a max pooling filter.


// A representation of a max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxNode
type CNNPoolingMaxNode struct {
	CNNPoolingNode
}

// CNNPoolingMaxNodeFrom constructs a [CNNPoolingMaxNode] from an unsafe.Pointer.
//
// A representation of a max pooling filter.
func CNNPoolingMaxNodeFrom(ptr unsafe.Pointer) CNNPoolingMaxNode {
	return CNNPoolingMaxNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingMaxNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingMaxNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingMaxNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingMaxNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingMaxNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingMaxNode */



