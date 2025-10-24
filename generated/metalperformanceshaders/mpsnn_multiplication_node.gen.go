// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNMultiplicationNode */


/* debug [class_header]: Header for MPSNNMultiplicationNode */
// The class instance for the [MultiplicationNode] class.
var (
	MultiplicationNodeClass     _MultiplicationNodeClass
	MultiplicationNodeClassOnce sync.Once
)

func getMultiplicationNodeClass() _MultiplicationNodeClass {
	MultiplicationNodeClassOnce.Do(func() {
		MultiplicationNodeClass = _MultiplicationNodeClass{objc.GetClass("MPSNNMultiplicationNode")}
	})
	return MultiplicationNodeClass
}

type _MultiplicationNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiplicationNode */
// An interface definition for the [MultiplicationNode] class.
type IMultiplicationNode interface {
	IBinaryArithmeticNode
	
/* debug [class_interface_properties]: Properties for MultiplicationNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MultiplicationNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiplicationNode */
// Alloc allocates a new instance without initialization.
func (mc _MultiplicationNodeClass) Alloc() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiplicationNodeClass) New() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiplicationNode) Init() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiplicationNode) Autorelease() MultiplicationNode {
	rv := objc.Send[MultiplicationNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiplicationNode creates a new MultiplicationNode instance.
func NewMultiplicationNode() MultiplicationNode {
	return getMultiplicationNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiplicationNode */
// A representation of a multiplication operator.


// A representation of a multiplication operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiplicationNode
type MultiplicationNode struct {
	BinaryArithmeticNode
}

// MultiplicationNodeFrom constructs a [MultiplicationNode] from an unsafe.Pointer.
//
// A representation of a multiplication operator.
func MultiplicationNodeFrom(ptr unsafe.Pointer) MultiplicationNode {
	return MultiplicationNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiplicationNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiplicationNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiplicationNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiplicationNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiplicationNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNMultiplicationNode */



