// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionGradientStateNode */


/* debug [class_header]: Header for MPSCNNConvolutionGradientStateNode */
// The class instance for the [CNNConvolutionGradientStateNode] class.
var (
	CNNConvolutionGradientStateNodeClass     _CNNConvolutionGradientStateNodeClass
	CNNConvolutionGradientStateNodeClassOnce sync.Once
)

func getCNNConvolutionGradientStateNodeClass() _CNNConvolutionGradientStateNodeClass {
	CNNConvolutionGradientStateNodeClassOnce.Do(func() {
		CNNConvolutionGradientStateNodeClass = _CNNConvolutionGradientStateNodeClass{objc.GetClass("MPSCNNConvolutionGradientStateNode")}
	})
	return CNNConvolutionGradientStateNodeClass
}

type _CNNConvolutionGradientStateNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionGradientStateNode */
// An interface definition for the [CNNConvolutionGradientStateNode] class.
type ICNNConvolutionGradientStateNode interface {
	IGradientStateNode
	
/* debug [class_interface_properties]: Properties for CNNConvolutionGradientStateNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionGradientStateNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionGradientStateNode */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientStateNodeClass) Alloc() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionGradientStateNodeClass) New() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientStateNode) Init() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientStateNode) Autorelease() CNNConvolutionGradientStateNode {
	rv := objc.Send[CNNConvolutionGradientStateNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientStateNode creates a new CNNConvolutionGradientStateNode instance.
func NewCNNConvolutionGradientStateNode() CNNConvolutionGradientStateNode {
	return getCNNConvolutionGradientStateNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionGradientStateNode */
// A representation of a gradient convolution state.


// A representation of a gradient convolution state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientStateNode
type CNNConvolutionGradientStateNode struct {
	GradientStateNode
}

// CNNConvolutionGradientStateNodeFrom constructs a [CNNConvolutionGradientStateNode] from an unsafe.Pointer.
//
// A representation of a gradient convolution state.
func CNNConvolutionGradientStateNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientStateNode {
	return CNNConvolutionGradientStateNode{
		GradientStateNode: GradientStateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionGradientStateNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionGradientStateNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionGradientStateNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionGradientStateNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionGradientStateNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionGradientStateNode */



