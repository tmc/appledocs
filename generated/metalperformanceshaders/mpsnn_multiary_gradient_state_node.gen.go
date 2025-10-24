// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNMultiaryGradientStateNode */


/* debug [class_header]: Header for MPSNNMultiaryGradientStateNode */
// The class instance for the [MultiaryGradientStateNode] class.
var (
	MultiaryGradientStateNodeClass     _MultiaryGradientStateNodeClass
	MultiaryGradientStateNodeClassOnce sync.Once
)

func getMultiaryGradientStateNodeClass() _MultiaryGradientStateNodeClass {
	MultiaryGradientStateNodeClassOnce.Do(func() {
		MultiaryGradientStateNodeClass = _MultiaryGradientStateNodeClass{objc.GetClass("MPSNNMultiaryGradientStateNode")}
	})
	return MultiaryGradientStateNodeClass
}

type _MultiaryGradientStateNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiaryGradientStateNode */
// An interface definition for the [MultiaryGradientStateNode] class.
type IMultiaryGradientStateNode interface {
	IStateNode
	
/* debug [class_interface_properties]: Properties for MultiaryGradientStateNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MultiaryGradientStateNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiaryGradientStateNode */
// Alloc allocates a new instance without initialization.
func (mc _MultiaryGradientStateNodeClass) Alloc() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiaryGradientStateNodeClass) New() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiaryGradientStateNode) Init() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiaryGradientStateNode) Autorelease() MultiaryGradientStateNode {
	rv := objc.Send[MultiaryGradientStateNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiaryGradientStateNode creates a new MultiaryGradientStateNode instance.
func NewMultiaryGradientStateNode() MultiaryGradientStateNode {
	return getMultiaryGradientStateNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiaryGradientStateNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientStateNode
type MultiaryGradientStateNode struct {
	StateNode
}

// MultiaryGradientStateNodeFrom constructs a [MultiaryGradientStateNode] from an unsafe.Pointer.
func MultiaryGradientStateNodeFrom(ptr unsafe.Pointer) MultiaryGradientStateNode {
	return MultiaryGradientStateNode{
		StateNode: StateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiaryGradientStateNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiaryGradientStateNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiaryGradientStateNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiaryGradientStateNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiaryGradientStateNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNMultiaryGradientStateNode */



