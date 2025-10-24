// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNPoolingAverageGradientNode */


/* debug [class_header]: Header for MPSCNNPoolingAverageGradientNode */
// The class instance for the [CNNPoolingAverageGradientNode] class.
var (
	CNNPoolingAverageGradientNodeClass     _CNNPoolingAverageGradientNodeClass
	CNNPoolingAverageGradientNodeClassOnce sync.Once
)

func getCNNPoolingAverageGradientNodeClass() _CNNPoolingAverageGradientNodeClass {
	CNNPoolingAverageGradientNodeClassOnce.Do(func() {
		CNNPoolingAverageGradientNodeClass = _CNNPoolingAverageGradientNodeClass{objc.GetClass("MPSCNNPoolingAverageGradientNode")}
	})
	return CNNPoolingAverageGradientNodeClass
}

type _CNNPoolingAverageGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNPoolingAverageGradientNode */
// An interface definition for the [CNNPoolingAverageGradientNode] class.
type ICNNPoolingAverageGradientNode interface {
	ICNNPoolingGradientNode
	
/* debug [class_interface_properties]: Properties for CNNPoolingAverageGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNPoolingAverageGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNPoolingAverageGradientNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageGradientNodeClass) Alloc() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageGradientNodeClass) New() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverageGradientNode) Init() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverageGradientNode) Autorelease() CNNPoolingAverageGradientNode {
	rv := objc.Send[CNNPoolingAverageGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverageGradientNode creates a new CNNPoolingAverageGradientNode instance.
func NewCNNPoolingAverageGradientNode() CNNPoolingAverageGradientNode {
	return getCNNPoolingAverageGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNPoolingAverageGradientNode */
// A representation of a gradient average pooling filter.


// A representation of a gradient average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradientNode
type CNNPoolingAverageGradientNode struct {
	CNNPoolingGradientNode
}

// CNNPoolingAverageGradientNodeFrom constructs a [CNNPoolingAverageGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient average pooling filter.
func CNNPoolingAverageGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingAverageGradientNode {
	return CNNPoolingAverageGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNPoolingAverageGradientNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNPoolingAverageGradientNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNPoolingAverageGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNPoolingAverageGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNPoolingAverageGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNPoolingAverageGradientNode */



