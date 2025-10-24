// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNDivisionNode */


/* debug [class_header]: Header for MPSNNDivisionNode */
// The class instance for the [DivisionNode] class.
var (
	DivisionNodeClass     _DivisionNodeClass
	DivisionNodeClassOnce sync.Once
)

func getDivisionNodeClass() _DivisionNodeClass {
	DivisionNodeClassOnce.Do(func() {
		DivisionNodeClass = _DivisionNodeClass{objc.GetClass("MPSNNDivisionNode")}
	})
	return DivisionNodeClass
}

type _DivisionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DivisionNode */
// An interface definition for the [DivisionNode] class.
type IDivisionNode interface {
	IBinaryArithmeticNode
	
/* debug [class_interface_properties]: Properties for DivisionNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DivisionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DivisionNode */
// Alloc allocates a new instance without initialization.
func (dc _DivisionNodeClass) Alloc() DivisionNode {
	rv := objc.Send[DivisionNode](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DivisionNodeClass) New() DivisionNode {
	rv := objc.Send[DivisionNode](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DivisionNode) Init() DivisionNode {
	rv := objc.Send[DivisionNode](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DivisionNode) Autorelease() DivisionNode {
	rv := objc.Send[DivisionNode](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDivisionNode creates a new DivisionNode instance.
func NewDivisionNode() DivisionNode {
	return getDivisionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DivisionNode */
// A representation of a division operator.


// A representation of a division operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDivisionNode
type DivisionNode struct {
	BinaryArithmeticNode
}

// DivisionNodeFrom constructs a [DivisionNode] from an unsafe.Pointer.
//
// A representation of a division operator.
func DivisionNodeFrom(ptr unsafe.Pointer) DivisionNode {
	return DivisionNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DivisionNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DivisionNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DivisionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DivisionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DivisionNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNDivisionNode */



