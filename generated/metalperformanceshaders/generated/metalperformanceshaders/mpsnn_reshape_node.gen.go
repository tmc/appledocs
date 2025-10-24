// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNReshapeNode */


/* debug [class_header]: Header for MPSNNReshapeNode */
// The class instance for the [ReshapeNode] class.
var (
	ReshapeNodeClass     _ReshapeNodeClass
	ReshapeNodeClassOnce sync.Once
)

func getReshapeNodeClass() _ReshapeNodeClass {
	ReshapeNodeClassOnce.Do(func() {
		ReshapeNodeClass = _ReshapeNodeClass{objc.GetClass("MPSNNReshapeNode")}
	})
	return ReshapeNodeClass
}

type _ReshapeNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReshapeNode */
// An interface definition for the [ReshapeNode] class.
type IReshapeNode interface {
	IFilterNode
	
/* debug [class_interface_properties]: Properties for ReshapeNode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReshapeNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReshapeNode */
// Alloc allocates a new instance without initialization.
func (rc _ReshapeNodeClass) Alloc() ReshapeNode {
	rv := objc.Send[ReshapeNode](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReshapeNodeClass) New() ReshapeNode {
	rv := objc.Send[ReshapeNode](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReshapeNode) Init() ReshapeNode {
	rv := objc.Send[ReshapeNode](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReshapeNode) Autorelease() ReshapeNode {
	rv := objc.Send[ReshapeNode](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReshapeNode creates a new ReshapeNode instance.
func NewReshapeNode() ReshapeNode {
	return getReshapeNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReshapeNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReshapeNode
type ReshapeNode struct {
	FilterNode
}

// ReshapeNodeFrom constructs a [ReshapeNode] from an unsafe.Pointer.
func ReshapeNodeFrom(ptr unsafe.Pointer) ReshapeNode {
	return ReshapeNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReshapeNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode/3037420-initwithsource
func NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(source IImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) ReshapeNode {
	instance := getReshapeNodeClass().Alloc()
	rv := objc.Send[ReshapeNode](instance.ID, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReshapeNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode/3037421-nodewithsource
func (rc _ReshapeNodeClass) NodeWithSourceResultWidthResultHeightResultFeatureChannels(source IImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:"), source, resultWidth, resultHeight, resultFeatureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NodeWithSourceResultWidthResultHeightResultFeatureChannels) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReshapeNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReshapeNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReshapeNode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReshapeNode */


