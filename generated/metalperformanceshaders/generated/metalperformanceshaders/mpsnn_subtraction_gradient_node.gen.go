// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNSubtractionGradientNode */


/* debug [class_header]: Header for MPSNNSubtractionGradientNode */
// The class instance for the [SubtractionGradientNode] class.
var (
	SubtractionGradientNodeClass     _SubtractionGradientNodeClass
	SubtractionGradientNodeClassOnce sync.Once
)

func getSubtractionGradientNodeClass() _SubtractionGradientNodeClass {
	SubtractionGradientNodeClassOnce.Do(func() {
		SubtractionGradientNodeClass = _SubtractionGradientNodeClass{objc.GetClass("MPSNNSubtractionGradientNode")}
	})
	return SubtractionGradientNodeClass
}

type _SubtractionGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SubtractionGradientNode */
// An interface definition for the [SubtractionGradientNode] class.
type ISubtractionGradientNode interface {
	IArithmeticGradientNode
	
/* debug [class_interface_properties]: Properties for SubtractionGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SubtractionGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SubtractionGradientNode */
// Alloc allocates a new instance without initialization.
func (sc _SubtractionGradientNodeClass) Alloc() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubtractionGradientNodeClass) New() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SubtractionGradientNode) Init() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SubtractionGradientNode) Autorelease() SubtractionGradientNode {
	rv := objc.Send[SubtractionGradientNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubtractionGradientNode creates a new SubtractionGradientNode instance.
func NewSubtractionGradientNode() SubtractionGradientNode {
	return getSubtractionGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SubtractionGradientNode */
// A representation of a gradient subtraction operator.


// A representation of a gradient subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionGradientNode
type SubtractionGradientNode struct {
	ArithmeticGradientNode
}

// SubtractionGradientNodeFrom constructs a [SubtractionGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient subtraction operator.
func SubtractionGradientNodeFrom(ptr unsafe.Pointer) SubtractionGradientNode {
	return SubtractionGradientNode{
		ArithmeticGradientNode: ArithmeticGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SubtractionGradientNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SubtractionGradientNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SubtractionGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SubtractionGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SubtractionGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNSubtractionGradientNode */



