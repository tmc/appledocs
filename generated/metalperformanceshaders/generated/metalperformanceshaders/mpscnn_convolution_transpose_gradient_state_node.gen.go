// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionTransposeGradientStateNode */


/* debug [class_header]: Header for MPSCNNConvolutionTransposeGradientStateNode */
// The class instance for the [CNNConvolutionTransposeGradientStateNode] class.
var (
	CNNConvolutionTransposeGradientStateNodeClass     _CNNConvolutionTransposeGradientStateNodeClass
	CNNConvolutionTransposeGradientStateNodeClassOnce sync.Once
)

func getCNNConvolutionTransposeGradientStateNodeClass() _CNNConvolutionTransposeGradientStateNodeClass {
	CNNConvolutionTransposeGradientStateNodeClassOnce.Do(func() {
		CNNConvolutionTransposeGradientStateNodeClass = _CNNConvolutionTransposeGradientStateNodeClass{objc.GetClass("MPSCNNConvolutionTransposeGradientStateNode")}
	})
	return CNNConvolutionTransposeGradientStateNodeClass
}

type _CNNConvolutionTransposeGradientStateNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionTransposeGradientStateNode */
// An interface definition for the [CNNConvolutionTransposeGradientStateNode] class.
type ICNNConvolutionTransposeGradientStateNode interface {
	ICNNConvolutionGradientStateNode
	
/* debug [class_interface_properties]: Properties for CNNConvolutionTransposeGradientStateNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionTransposeGradientStateNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionTransposeGradientStateNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeGradientStateNodeClass) Alloc() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeGradientStateNodeClass) New() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTransposeGradientStateNode) Init() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTransposeGradientStateNode) Autorelease() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Send[CNNConvolutionTransposeGradientStateNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTransposeGradientStateNode creates a new CNNConvolutionTransposeGradientStateNode instance.
func NewCNNConvolutionTransposeGradientStateNode() CNNConvolutionTransposeGradientStateNode {
	return getCNNConvolutionTransposeGradientStateNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionTransposeGradientStateNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTransposeGradientStateNode
type CNNConvolutionTransposeGradientStateNode struct {
	CNNConvolutionGradientStateNode
}

// CNNConvolutionTransposeGradientStateNodeFrom constructs a [CNNConvolutionTransposeGradientStateNode] from an unsafe.Pointer.
func CNNConvolutionTransposeGradientStateNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientStateNode {
	return CNNConvolutionTransposeGradientStateNode{
		CNNConvolutionGradientStateNode: CNNConvolutionGradientStateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionTransposeGradientStateNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionTransposeGradientStateNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionTransposeGradientStateNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionTransposeGradientStateNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionTransposeGradientStateNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionTransposeGradientStateNode */



