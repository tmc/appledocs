// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingMaxGradientNode */


/* debug [class_header]: Header for MPSCNNPoolingMaxGradientNode */
// The class instance for the [CNNPoolingMaxGradientNode] class.
var (
	CNNPoolingMaxGradientNodeClass     _CNNPoolingMaxGradientNodeClass
	CNNPoolingMaxGradientNodeClassOnce sync.Once
)

func getCNNPoolingMaxGradientNodeClass() _CNNPoolingMaxGradientNodeClass {
	CNNPoolingMaxGradientNodeClassOnce.Do(func() {
		CNNPoolingMaxGradientNodeClass = _CNNPoolingMaxGradientNodeClass{objc.GetClass("MPSCNNPoolingMaxGradientNode")}
	})
	return CNNPoolingMaxGradientNodeClass
}

type _CNNPoolingMaxGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingMaxGradientNode */
// An interface definition for the [CNNPoolingMaxGradientNode] class.
type ICNNPoolingMaxGradientNode interface {
	ICNNPoolingGradientNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingMaxGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingMaxGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingMaxGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxGradientNodeClass) Alloc() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxGradientNodeClass) New() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMaxGradientNode) Init() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMaxGradientNode) Autorelease() CNNPoolingMaxGradientNode {
	rv := objc.Send[CNNPoolingMaxGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMaxGradientNode creates a new CNNPoolingMaxGradientNode instance.
func NewCNNPoolingMaxGradientNode() CNNPoolingMaxGradientNode {
	return getCNNPoolingMaxGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingMaxGradientNode */
// A representation of a gradient max pooling filter.


// A representation of a gradient max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradientNode
type CNNPoolingMaxGradientNode struct {
	CNNPoolingGradientNode
}

// CNNPoolingMaxGradientNodeFrom constructs a [CNNPoolingMaxGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient max pooling filter.
func CNNPoolingMaxGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingMaxGradientNode {
	return CNNPoolingMaxGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingMaxGradientNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingMaxGradientNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingMaxGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingMaxGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingMaxGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingMaxGradientNode */



