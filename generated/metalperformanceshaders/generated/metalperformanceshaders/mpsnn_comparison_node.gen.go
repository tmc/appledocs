// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNComparisonNode */


/* debug [class_header]: Header for MPSNNComparisonNode */
// The class instance for the [ComparisonNode] class.
var (
	ComparisonNodeClass     _ComparisonNodeClass
	ComparisonNodeClassOnce sync.Once
)

func getComparisonNodeClass() _ComparisonNodeClass {
	ComparisonNodeClassOnce.Do(func() {
		ComparisonNodeClass = _ComparisonNodeClass{objc.GetClass("MPSNNComparisonNode")}
	})
	return ComparisonNodeClass
}

type _ComparisonNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ComparisonNode */
// An interface definition for the [ComparisonNode] class.
type IComparisonNode interface {
	IBinaryArithmeticNode
	
/* debug [class_interface_properties]: Properties for ComparisonNode */
	// properties:
	ComparisonType() ComparisonType get set /* not a class type */
	SetComparisonType(value ComparisonType get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ComparisonNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ComparisonNode */
// Alloc allocates a new instance without initialization.
func (cc _ComparisonNodeClass) Alloc() ComparisonNode {
	rv := objc.Send[ComparisonNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ComparisonNodeClass) New() ComparisonNode {
	rv := objc.Send[ComparisonNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComparisonNode) Init() ComparisonNode {
	rv := objc.Send[ComparisonNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComparisonNode) Autorelease() ComparisonNode {
	rv := objc.Send[ComparisonNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComparisonNode creates a new ComparisonNode instance.
func NewComparisonNode() ComparisonNode {
	return getComparisonNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ComparisonNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNComparisonNode
type ComparisonNode struct {
	BinaryArithmeticNode
}

// ComparisonNodeFrom constructs a [ComparisonNode] from an unsafe.Pointer.
func ComparisonNodeFrom(ptr unsafe.Pointer) ComparisonNode {
	return ComparisonNode{
		BinaryArithmeticNode: BinaryArithmeticNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ComparisonNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ComparisonNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ComparisonNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ComparisonNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ComparisonNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisonnode/3037389-comparisontype
func (c_ ComparisonNode) ComparisonType() ComparisonType get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("comparisonType"))
	return rv
}/* debug [instance_properties/getter]: comparisonType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisonnode/3037389-comparisontype
func (c_ ComparisonNode) SetComparisonType(value ComparisonType get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComparisonType:"), value)
}/* debug [instance_properties/setter]: comparisonType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNComparisonNode */



