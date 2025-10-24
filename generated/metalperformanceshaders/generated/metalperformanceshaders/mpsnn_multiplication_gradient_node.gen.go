// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNMultiplicationGradientNode */


/* debug [class_header]: Header for MPSNNMultiplicationGradientNode */
// The class instance for the [MultiplicationGradientNode] class.
var (
	MultiplicationGradientNodeClass     _MultiplicationGradientNodeClass
	MultiplicationGradientNodeClassOnce sync.Once
)

func getMultiplicationGradientNodeClass() _MultiplicationGradientNodeClass {
	MultiplicationGradientNodeClassOnce.Do(func() {
		MultiplicationGradientNodeClass = _MultiplicationGradientNodeClass{objc.GetClass("MPSNNMultiplicationGradientNode")}
	})
	return MultiplicationGradientNodeClass
}

type _MultiplicationGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiplicationGradientNode */
// An interface definition for the [MultiplicationGradientNode] class.
type IMultiplicationGradientNode interface {
	IArithmeticGradientNode
	
/* debug [class_interface_properties]: Properties for MultiplicationGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MultiplicationGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiplicationGradientNode */
// Alloc allocates a new instance without initialization.
func (mc _MultiplicationGradientNodeClass) Alloc() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiplicationGradientNodeClass) New() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiplicationGradientNode) Init() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiplicationGradientNode) Autorelease() MultiplicationGradientNode {
	rv := objc.Send[MultiplicationGradientNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiplicationGradientNode creates a new MultiplicationGradientNode instance.
func NewMultiplicationGradientNode() MultiplicationGradientNode {
	return getMultiplicationGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiplicationGradientNode */
// A representation of a gradient multiplication operator.


// A representation of a gradient multiplication operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiplicationGradientNode
type MultiplicationGradientNode struct {
	ArithmeticGradientNode
}

// MultiplicationGradientNodeFrom constructs a [MultiplicationGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient multiplication operator.
func MultiplicationGradientNodeFrom(ptr unsafe.Pointer) MultiplicationGradientNode {
	return MultiplicationGradientNode{
		ArithmeticGradientNode: ArithmeticGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiplicationGradientNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiplicationGradientNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiplicationGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiplicationGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiplicationGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNMultiplicationGradientNode */



