// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNLabelsNode */


/* debug [class_header]: Header for MPSNNLabelsNode */
// The class instance for the [LabelsNode] class.
var (
	LabelsNodeClass     _LabelsNodeClass
	LabelsNodeClassOnce sync.Once
)

func getLabelsNodeClass() _LabelsNodeClass {
	LabelsNodeClassOnce.Do(func() {
		LabelsNodeClass = _LabelsNodeClass{objc.GetClass("MPSNNLabelsNode")}
	})
	return LabelsNodeClass
}

type _LabelsNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LabelsNode */
// An interface definition for the [LabelsNode] class.
type ILabelsNode interface {
	IStateNode
	
/* debug [class_interface_properties]: Properties for LabelsNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LabelsNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LabelsNode */
// Alloc allocates a new instance without initialization.
func (lc _LabelsNodeClass) Alloc() LabelsNode {
	rv := objc.Send[LabelsNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LabelsNodeClass) New() LabelsNode {
	rv := objc.Send[LabelsNode](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LabelsNode) Init() LabelsNode {
	rv := objc.Send[LabelsNode](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LabelsNode) Autorelease() LabelsNode {
	rv := objc.Send[LabelsNode](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLabelsNode creates a new LabelsNode instance.
func NewLabelsNode() LabelsNode {
	return getLabelsNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LabelsNode */
// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels.


// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLabelsNode
type LabelsNode struct {
	StateNode
}

// LabelsNodeFrom constructs a [LabelsNode] from an unsafe.Pointer.
//
// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels.
func LabelsNodeFrom(ptr unsafe.Pointer) LabelsNode {
	return LabelsNode{
		StateNode: StateNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LabelsNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LabelsNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LabelsNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LabelsNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LabelsNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNLabelsNode */



