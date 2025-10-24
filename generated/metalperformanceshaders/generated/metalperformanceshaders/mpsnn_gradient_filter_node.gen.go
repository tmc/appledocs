// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNGradientFilterNode */


/* debug [class_header]: Header for MPSNNGradientFilterNode */
// The class instance for the [GradientFilterNode] class.
var (
	GradientFilterNodeClass     _GradientFilterNodeClass
	GradientFilterNodeClassOnce sync.Once
)

func getGradientFilterNodeClass() _GradientFilterNodeClass {
	GradientFilterNodeClassOnce.Do(func() {
		GradientFilterNodeClass = _GradientFilterNodeClass{objc.GetClass("MPSNNGradientFilterNode")}
	})
	return GradientFilterNodeClass
}

type _GradientFilterNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GradientFilterNode */
// An interface definition for the [GradientFilterNode] class.
type IGradientFilterNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for GradientFilterNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GradientFilterNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GradientFilterNode */
// Alloc allocates a new instance without initialization.
func (gc _GradientFilterNodeClass) Alloc() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GradientFilterNodeClass) New() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GradientFilterNode) Init() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GradientFilterNode) Autorelease() GradientFilterNode {
	rv := objc.Send[GradientFilterNode](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradientFilterNode creates a new GradientFilterNode instance.
func NewGradientFilterNode() GradientFilterNode {
	return getGradientFilterNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GradientFilterNode */
// A representation of a gradient filter.


// A representation of a gradient filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGradientFilterNode
type GradientFilterNode struct {
	FilterNode
}

// GradientFilterNodeFrom constructs a [GradientFilterNode] from an unsafe.Pointer.
//
// A representation of a gradient filter.
func GradientFilterNodeFrom(ptr unsafe.Pointer) GradientFilterNode {
	return GradientFilterNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GradientFilterNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GradientFilterNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GradientFilterNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GradientFilterNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GradientFilterNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNGradientFilterNode */



