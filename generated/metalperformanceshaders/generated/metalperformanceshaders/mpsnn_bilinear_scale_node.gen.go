// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNBilinearScaleNode */


/* debug [class_header]: Header for MPSNNBilinearScaleNode */
// The class instance for the [BilinearScaleNode] class.
var (
	BilinearScaleNodeClass     _BilinearScaleNodeClass
	BilinearScaleNodeClassOnce sync.Once
)

func getBilinearScaleNodeClass() _BilinearScaleNodeClass {
	BilinearScaleNodeClassOnce.Do(func() {
		BilinearScaleNodeClass = _BilinearScaleNodeClass{objc.GetClass("MPSNNBilinearScaleNode")}
	})
	return BilinearScaleNodeClass
}

type _BilinearScaleNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BilinearScaleNode */
// An interface definition for the [BilinearScaleNode] class.
type IBilinearScaleNode interface {
	IScaleNode
	
/* debug [class_interface_properties]: Properties for BilinearScaleNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BilinearScaleNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BilinearScaleNode */
// Alloc allocates a new instance without initialization.
func (bc _BilinearScaleNodeClass) Alloc() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BilinearScaleNodeClass) New() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BilinearScaleNode) Init() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BilinearScaleNode) Autorelease() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBilinearScaleNode creates a new BilinearScaleNode instance.
func NewBilinearScaleNode() BilinearScaleNode {
	return getBilinearScaleNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BilinearScaleNode */
// A representation of a bilinear resampling filter.


// A representation of a bilinear resampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBilinearScaleNode
type BilinearScaleNode struct {
	ScaleNode
}

// BilinearScaleNodeFrom constructs a [BilinearScaleNode] from an unsafe.Pointer.
//
// A representation of a bilinear resampling filter.
func BilinearScaleNodeFrom(ptr unsafe.Pointer) BilinearScaleNode {
	return BilinearScaleNode{
		ScaleNode: ScaleNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BilinearScaleNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BilinearScaleNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BilinearScaleNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BilinearScaleNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BilinearScaleNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNBilinearScaleNode */



