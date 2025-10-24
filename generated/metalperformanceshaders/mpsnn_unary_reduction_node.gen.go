// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNUnaryReductionNode */


/* debug [class_header]: Header for MPSNNUnaryReductionNode */
// The class instance for the [UnaryReductionNode] class.
var (
	UnaryReductionNodeClass     _UnaryReductionNodeClass
	UnaryReductionNodeClassOnce sync.Once
)

func getUnaryReductionNodeClass() _UnaryReductionNodeClass {
	UnaryReductionNodeClassOnce.Do(func() {
		UnaryReductionNodeClass = _UnaryReductionNodeClass{objc.GetClass("MPSNNUnaryReductionNode")}
	})
	return UnaryReductionNodeClass
}

type _UnaryReductionNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnaryReductionNode */
// An interface definition for the [UnaryReductionNode] class.
type IUnaryReductionNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for UnaryReductionNode */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnaryReductionNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnaryReductionNode */
// Alloc allocates a new instance without initialization.
func (uc _UnaryReductionNodeClass) Alloc() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnaryReductionNodeClass) New() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnaryReductionNode) Init() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnaryReductionNode) Autorelease() UnaryReductionNode {
	rv := objc.Send[UnaryReductionNode](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnaryReductionNode creates a new UnaryReductionNode instance.
func NewUnaryReductionNode() UnaryReductionNode {
	return getUnaryReductionNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnaryReductionNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNUnaryReductionNode
type UnaryReductionNode struct {
	FilterNode
}

// UnaryReductionNodeFrom constructs a [UnaryReductionNode] from an unsafe.Pointer.
func UnaryReductionNodeFrom(ptr unsafe.Pointer) UnaryReductionNode {
	return UnaryReductionNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnaryReductionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource
func NewUnaryReductionNodeWithSource(sourceNode IImageNode) UnaryReductionNode {
	instance := getUnaryReductionNodeClass().Alloc()
	rv := objc.Send[UnaryReductionNode](instance.ID, objc.Sel("initWithSource:"), sourceNode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUnaryReductionNodeWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnaryReductionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource
func (uc _UnaryReductionNodeClass) NodeWithSource(sourceNode IImageNode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnaryReductionNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnaryReductionNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnaryReductionNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037423-cliprectsource
func (u_ UnaryReductionNode) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037423-cliprectsource
func (u_ UnaryReductionNode) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNUnaryReductionNode */


