// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRClusterStateCacheContainer */


/* debug [class_header]: Header for MTRClusterStateCacheContainer */
// The class instance for the [MTRClusterStateCacheContainer] class.
var (
	MTRClusterStateCacheContainerClass     _MTRClusterStateCacheContainerClass
	MTRClusterStateCacheContainerClassOnce sync.Once
)

func getMTRClusterStateCacheContainerClass() _MTRClusterStateCacheContainerClass {
	MTRClusterStateCacheContainerClassOnce.Do(func() {
		MTRClusterStateCacheContainerClass = _MTRClusterStateCacheContainerClass{objc.GetClass("MTRClusterStateCacheContainer")}
	})
	return MTRClusterStateCacheContainerClass
}

type _MTRClusterStateCacheContainerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterStateCacheContainer */
// An interface definition for the [MTRClusterStateCacheContainer] class.
type IMTRClusterStateCacheContainer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRClusterStateCacheContainer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterStateCacheContainer */
	// methods:
	ReadAttributesWithEndpointIDClusterIDAttributeIDQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterStateCacheContainer */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterStateCacheContainerClass) Alloc() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterStateCacheContainerClass) New() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterStateCacheContainer) Init() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterStateCacheContainer) Autorelease() MTRClusterStateCacheContainer {
	rv := objc.Send[MTRClusterStateCacheContainer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterStateCacheContainer creates a new MTRClusterStateCacheContainer instance.
func NewMTRClusterStateCacheContainer() MTRClusterStateCacheContainer {
	return getMTRClusterStateCacheContainerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterStateCacheContainer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterStateCacheContainer
type MTRClusterStateCacheContainer struct {
	objectivec.Object
}

// MTRClusterStateCacheContainerFrom constructs a [MTRClusterStateCacheContainer] from an unsafe.Pointer.
func MTRClusterStateCacheContainerFrom(ptr unsafe.Pointer) MTRClusterStateCacheContainer {
	return MTRClusterStateCacheContainer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterStateCacheContainer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterStateCacheContainer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterStateCacheContainer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterStateCacheContainer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterStateCacheContainer/readAttributes(withEndpointID:clusterID:attributeID:queue:completion:)
func (m_ MTRClusterStateCacheContainer) ReadAttributesWithEndpointIDClusterIDAttributeIDQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributesWithEndpointID:clusterID:attributeID:queue:completion:"), endpointID, clusterID, attributeID, queue, completion)
}/* debug [instance_methods/method]: ReadAttributesWithEndpointIDClusterIDAttributeIDQueueCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterStateCacheContainer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterStateCacheContainer */



