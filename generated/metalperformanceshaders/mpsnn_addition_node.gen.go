// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNAdditionNode */


/* debug [class_header]: Header for MPSNNAdditionNode */
// The class instance for the [AdditionNode] class.
var (
	AdditionNodeClass     _AdditionNodeClass
	AdditionNodeClassOnce sync.Once
)

func getAdditionNodeClass() _AdditionNodeClass {
	AdditionNodeClassOnce.Do(func() {
		AdditionNodeClass = _AdditionNodeClass{objc.GetClass("MPSNNAdditionNode")}
	})
	return AdditionNodeClass
}

type _AdditionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AdditionNode */
// An interface definition for the [AdditionNode] class.
type IAdditionNode interface {
	IBinaryArithmeticNode
	
/* debug [class_interface_properties]: Properties for AdditionNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AdditionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AdditionNode */
// Alloc allocates a new instance without initialization.
func (ac _AdditionNodeClass) Alloc() AdditionNode {
	rv := objc.Send[AdditionNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdditionNodeClass) New() AdditionNode {
	rv := objc.Send[AdditionNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdditionNode) Init() AdditionNode {
	rv := objc.Send[AdditionNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdditionNode) Autorelease() AdditionNode {
	rv := objc.Send[AdditionNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdditionNode creates a new AdditionNode instance.
func NewAdditionNode() AdditionNode {
	return getAdditionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AdditionNode */
// A representation of an addition operator.


// A representation of an addition operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNAdditionNode
type AdditionNode struct {
	BinaryArithmeticNode
}

// AdditionNodeFrom constructs a [AdditionNode] from an unsafe.Pointer.
//
// A representation of an addition operator.
func AdditionNodeFrom(ptr unsafe.Pointer) AdditionNode {
	return AdditionNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AdditionNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AdditionNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AdditionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AdditionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AdditionNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNAdditionNode */



