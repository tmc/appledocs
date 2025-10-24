// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNSubtractionNode */


/* debug [class_header]: Header for MPSNNSubtractionNode */
// The class instance for the [SubtractionNode] class.
var (
	SubtractionNodeClass     _SubtractionNodeClass
	SubtractionNodeClassOnce sync.Once
)

func getSubtractionNodeClass() _SubtractionNodeClass {
	SubtractionNodeClassOnce.Do(func() {
		SubtractionNodeClass = _SubtractionNodeClass{objc.GetClass("MPSNNSubtractionNode")}
	})
	return SubtractionNodeClass
}

type _SubtractionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SubtractionNode */
// An interface definition for the [SubtractionNode] class.
type ISubtractionNode interface {
	IBinaryArithmeticNode
	
/* debug [class_interface_properties]: Properties for SubtractionNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SubtractionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SubtractionNode */
// Alloc allocates a new instance without initialization.
func (sc _SubtractionNodeClass) Alloc() SubtractionNode {
	rv := objc.Send[SubtractionNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubtractionNodeClass) New() SubtractionNode {
	rv := objc.Send[SubtractionNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SubtractionNode) Init() SubtractionNode {
	rv := objc.Send[SubtractionNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SubtractionNode) Autorelease() SubtractionNode {
	rv := objc.Send[SubtractionNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubtractionNode creates a new SubtractionNode instance.
func NewSubtractionNode() SubtractionNode {
	return getSubtractionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SubtractionNode */
// A representation of an subtraction operator.


// A representation of an subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNSubtractionNode
type SubtractionNode struct {
	BinaryArithmeticNode
}

// SubtractionNodeFrom constructs a [SubtractionNode] from an unsafe.Pointer.
//
// A representation of an subtraction operator.
func SubtractionNodeFrom(ptr unsafe.Pointer) SubtractionNode {
	return SubtractionNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SubtractionNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SubtractionNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SubtractionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SubtractionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SubtractionNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNSubtractionNode */



