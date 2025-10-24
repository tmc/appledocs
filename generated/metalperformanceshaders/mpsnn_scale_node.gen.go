// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNScaleNode */


/* debug [class_header]: Header for MPSNNScaleNode */
// The class instance for the [ScaleNode] class.
var (
	ScaleNodeClass     _ScaleNodeClass
	ScaleNodeClassOnce sync.Once
)

func getScaleNodeClass() _ScaleNodeClass {
	ScaleNodeClassOnce.Do(func() {
		ScaleNodeClass = _ScaleNodeClass{objc.GetClass("MPSNNScaleNode")}
	})
	return ScaleNodeClass
}

type _ScaleNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScaleNode */
// An interface definition for the [ScaleNode] class.
type IScaleNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for ScaleNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScaleNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScaleNode */
// Alloc allocates a new instance without initialization.
func (sc _ScaleNodeClass) Alloc() ScaleNode {
	rv := objc.Send[ScaleNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScaleNodeClass) New() ScaleNode {
	rv := objc.Send[ScaleNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScaleNode) Init() ScaleNode {
	rv := objc.Send[ScaleNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScaleNode) Autorelease() ScaleNode {
	rv := objc.Send[ScaleNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScaleNode creates a new ScaleNode instance.
func NewScaleNode() ScaleNode {
	return getScaleNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScaleNode */
// Abstract node representing an image resampling filter.


// Abstract node representing an image resampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode
type ScaleNode struct {
	FilterNode
}

// ScaleNodeFrom constructs a [ScaleNode] from an unsafe.Pointer.
//
// Abstract node representing an image resampling filter.
func ScaleNodeFrom(ptr unsafe.Pointer) ScaleNode {
	return ScaleNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScaleNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915285-initwithsource
func NewScaleNodeWithSourceOutputSize(sourceNode IImageNode, size objc.IObject /* cross-framework: MTLSize */) ScaleNode {
	instance := getScaleNodeClass().Alloc()
	rv := objc.Send[ScaleNode](instance.ID, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScaleNodeWithSourceOutputSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915278-initwithsource
func NewScaleNodeWithSourceTransformProviderOutputSize(sourceNode IImageNode, transformProvider unsafe.Pointer, size objc.IObject /* cross-framework: MTLSize */) ScaleNode {
	instance := getScaleNodeClass().Alloc()
	rv := objc.Send[ScaleNode](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScaleNodeWithSourceTransformProviderOutputSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScaleNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915280-nodewithsource
func (sc _ScaleNodeClass) NodeWithSourceOutputSize(sourceNode IImageNode, size objc.IObject /* cross-framework: MTLSize */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("nodeWithSource:outputSize:"), sourceNode, size)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceOutputSize) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915286-nodewithsource
func (sc _ScaleNodeClass) NodeWithSourceTransformProviderOutputSize(sourceNode IImageNode, transformProvider unsafe.Pointer, size objc.IObject /* cross-framework: MTLSize */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("nodeWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceTransformProviderOutputSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScaleNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScaleNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScaleNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNScaleNode */


