// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingL2NormNode */


/* debug [class_header]: Header for MPSCNNPoolingL2NormNode */
// The class instance for the [CNNPoolingL2NormNode] class.
var (
	CNNPoolingL2NormNodeClass     _CNNPoolingL2NormNodeClass
	CNNPoolingL2NormNodeClassOnce sync.Once
)

func getCNNPoolingL2NormNodeClass() _CNNPoolingL2NormNodeClass {
	CNNPoolingL2NormNodeClassOnce.Do(func() {
		CNNPoolingL2NormNodeClass = _CNNPoolingL2NormNodeClass{objc.GetClass("MPSCNNPoolingL2NormNode")}
	})
	return CNNPoolingL2NormNodeClass
}

type _CNNPoolingL2NormNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingL2NormNode */
// An interface definition for the [CNNPoolingL2NormNode] class.
type ICNNPoolingL2NormNode interface {
	ICNNPoolingNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingL2NormNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingL2NormNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingL2NormNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingL2NormNodeClass) Alloc() CNNPoolingL2NormNode {
	rv := objc.Send[CNNPoolingL2NormNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingL2NormNodeClass) New() CNNPoolingL2NormNode {
	rv := objc.Send[CNNPoolingL2NormNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingL2NormNode) Init() CNNPoolingL2NormNode {
	rv := objc.Send[CNNPoolingL2NormNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingL2NormNode) Autorelease() CNNPoolingL2NormNode {
	rv := objc.Send[CNNPoolingL2NormNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingL2NormNode creates a new CNNPoolingL2NormNode instance.
func NewCNNPoolingL2NormNode() CNNPoolingL2NormNode {
	return getCNNPoolingL2NormNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingL2NormNode */
// A representation of a L2-norm pooling filter.


// A representation of a L2-norm pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormNode
type CNNPoolingL2NormNode struct {
	CNNPoolingNode
}

// CNNPoolingL2NormNodeFrom constructs a [CNNPoolingL2NormNode] from an unsafe.Pointer.
//
// A representation of a L2-norm pooling filter.
func CNNPoolingL2NormNodeFrom(ptr unsafe.Pointer) CNNPoolingL2NormNode {
	return CNNPoolingL2NormNode{
		CNNPoolingNode: CNNPoolingNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingL2NormNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingL2NormNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingL2NormNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingL2NormNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingL2NormNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingL2NormNode */



