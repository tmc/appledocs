// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNAdditionGradientNode */


/* debug [class_header]: Header for MPSNNAdditionGradientNode */
// The class instance for the [AdditionGradientNode] class.
var (
	AdditionGradientNodeClass     _AdditionGradientNodeClass
	AdditionGradientNodeClassOnce sync.Once
)

func getAdditionGradientNodeClass() _AdditionGradientNodeClass {
	AdditionGradientNodeClassOnce.Do(func() {
		AdditionGradientNodeClass = _AdditionGradientNodeClass{objc.GetClass("MPSNNAdditionGradientNode")}
	})
	return AdditionGradientNodeClass
}

type _AdditionGradientNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AdditionGradientNode */
// An interface definition for the [AdditionGradientNode] class.
type IAdditionGradientNode interface {
	IArithmeticGradientNode
	
/* debug [class_interface_properties]: Properties for AdditionGradientNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AdditionGradientNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AdditionGradientNode */
// Alloc allocates a new instance without initialization.
func (ac _AdditionGradientNodeClass) Alloc() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdditionGradientNodeClass) New() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdditionGradientNode) Init() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdditionGradientNode) Autorelease() AdditionGradientNode {
	rv := objc.Send[AdditionGradientNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdditionGradientNode creates a new AdditionGradientNode instance.
func NewAdditionGradientNode() AdditionGradientNode {
	return getAdditionGradientNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AdditionGradientNode */
// A representation of a gradient addition operator.


// A representation of a gradient addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionGradientNode
type AdditionGradientNode struct {
	ArithmeticGradientNode
}

// AdditionGradientNodeFrom constructs a [AdditionGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient addition operator.
func AdditionGradientNodeFrom(ptr unsafe.Pointer) AdditionGradientNode {
	return AdditionGradientNode{
		ArithmeticGradientNode: ArithmeticGradientNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AdditionGradientNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AdditionGradientNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AdditionGradientNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AdditionGradientNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AdditionGradientNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNAdditionGradientNode */



