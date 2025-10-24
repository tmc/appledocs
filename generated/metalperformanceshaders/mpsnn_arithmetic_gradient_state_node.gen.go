// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNArithmeticGradientStateNode */


/* debug [class_header]: Header for MPSNNArithmeticGradientStateNode */
// The class instance for the [ArithmeticGradientStateNode] class.
var (
	ArithmeticGradientStateNodeClass     _ArithmeticGradientStateNodeClass
	ArithmeticGradientStateNodeClassOnce sync.Once
)

func getArithmeticGradientStateNodeClass() _ArithmeticGradientStateNodeClass {
	ArithmeticGradientStateNodeClassOnce.Do(func() {
		ArithmeticGradientStateNodeClass = _ArithmeticGradientStateNodeClass{objc.GetClass("MPSNNArithmeticGradientStateNode")}
	})
	return ArithmeticGradientStateNodeClass
}

type _ArithmeticGradientStateNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ArithmeticGradientStateNode */
// An interface definition for the [ArithmeticGradientStateNode] class.
type IArithmeticGradientStateNode interface {
	IBinaryGradientStateNode
	
/* debug [class_interface_properties]: Properties for ArithmeticGradientStateNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ArithmeticGradientStateNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ArithmeticGradientStateNode */
// Alloc allocates a new instance without initialization.
func (ac _ArithmeticGradientStateNodeClass) Alloc() ArithmeticGradientStateNode {
	rv := objc.Send[ArithmeticGradientStateNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArithmeticGradientStateNodeClass) New() ArithmeticGradientStateNode {
	rv := objc.Send[ArithmeticGradientStateNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArithmeticGradientStateNode) Init() ArithmeticGradientStateNode {
	rv := objc.Send[ArithmeticGradientStateNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArithmeticGradientStateNode) Autorelease() ArithmeticGradientStateNode {
	rv := objc.Send[ArithmeticGradientStateNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArithmeticGradientStateNode creates a new ArithmeticGradientStateNode instance.
func NewArithmeticGradientStateNode() ArithmeticGradientStateNode {
	return getArithmeticGradientStateNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ArithmeticGradientStateNode */
// A representation of the clamp mask used by gradient arithmetic operators.


// A representation of the clamp mask used by gradient arithmetic operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNArithmeticGradientStateNode
type ArithmeticGradientStateNode struct {
	BinaryGradientStateNode
}

// ArithmeticGradientStateNodeFrom constructs a [ArithmeticGradientStateNode] from an unsafe.Pointer.
//
// A representation of the clamp mask used by gradient arithmetic operators.
func ArithmeticGradientStateNodeFrom(ptr unsafe.Pointer) ArithmeticGradientStateNode {
	return ArithmeticGradientStateNode{
		BinaryGradientStateNode: BinaryGradientStateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ArithmeticGradientStateNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ArithmeticGradientStateNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ArithmeticGradientStateNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ArithmeticGradientStateNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ArithmeticGradientStateNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNArithmeticGradientStateNode */



