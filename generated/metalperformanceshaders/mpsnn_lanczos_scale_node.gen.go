// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNLanczosScaleNode */


/* debug [class_header]: Header for MPSNNLanczosScaleNode */
// The class instance for the [LanczosScaleNode] class.
var (
	LanczosScaleNodeClass     _LanczosScaleNodeClass
	LanczosScaleNodeClassOnce sync.Once
)

func getLanczosScaleNodeClass() _LanczosScaleNodeClass {
	LanczosScaleNodeClassOnce.Do(func() {
		LanczosScaleNodeClass = _LanczosScaleNodeClass{objc.GetClass("MPSNNLanczosScaleNode")}
	})
	return LanczosScaleNodeClass
}

type _LanczosScaleNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LanczosScaleNode */
// An interface definition for the [LanczosScaleNode] class.
type ILanczosScaleNode interface {
	IScaleNode
	
/* debug [class_interface_properties]: Properties for LanczosScaleNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LanczosScaleNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LanczosScaleNode */
// Alloc allocates a new instance without initialization.
func (lc _LanczosScaleNodeClass) Alloc() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LanczosScaleNodeClass) New() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LanczosScaleNode) Init() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LanczosScaleNode) Autorelease() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLanczosScaleNode creates a new LanczosScaleNode instance.
func NewLanczosScaleNode() LanczosScaleNode {
	return getLanczosScaleNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LanczosScaleNode */
// A representation of a Lanczos resampling filter.


// A representation of a Lanczos resampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLanczosScaleNode
type LanczosScaleNode struct {
	ScaleNode
}

// LanczosScaleNodeFrom constructs a [LanczosScaleNode] from an unsafe.Pointer.
//
// A representation of a Lanczos resampling filter.
func LanczosScaleNodeFrom(ptr unsafe.Pointer) LanczosScaleNode {
	return LanczosScaleNode{
		ScaleNode: ScaleNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LanczosScaleNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LanczosScaleNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LanczosScaleNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LanczosScaleNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LanczosScaleNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNLanczosScaleNode */



