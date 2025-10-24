// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterEnergyEVSEMode */


/* debug [class_header]: Header for MTRBaseClusterEnergyEVSEMode */
// The class instance for the [MTRBaseClusterEnergyEVSEMode] class.
var (
	MTRBaseClusterEnergyEVSEModeClass     _MTRBaseClusterEnergyEVSEModeClass
	MTRBaseClusterEnergyEVSEModeClassOnce sync.Once
)

func getMTRBaseClusterEnergyEVSEModeClass() _MTRBaseClusterEnergyEVSEModeClass {
	MTRBaseClusterEnergyEVSEModeClassOnce.Do(func() {
		MTRBaseClusterEnergyEVSEModeClass = _MTRBaseClusterEnergyEVSEModeClass{objc.GetClass("MTRBaseClusterEnergyEVSEMode")}
	})
	return MTRBaseClusterEnergyEVSEModeClass
}

type _MTRBaseClusterEnergyEVSEModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterEnergyEVSEMode */
// An interface definition for the [MTRBaseClusterEnergyEVSEMode] class.
type IMTRBaseClusterEnergyEVSEMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterEnergyEVSEMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterEnergyEVSEMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterEnergyEVSEMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterEnergyEVSEModeClass) Alloc() MTRBaseClusterEnergyEVSEMode {
	rv := objc.Send[MTRBaseClusterEnergyEVSEMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterEnergyEVSEModeClass) New() MTRBaseClusterEnergyEVSEMode {
	rv := objc.Send[MTRBaseClusterEnergyEVSEMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterEnergyEVSEMode) Init() MTRBaseClusterEnergyEVSEMode {
	rv := objc.Send[MTRBaseClusterEnergyEVSEMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterEnergyEVSEMode) Autorelease() MTRBaseClusterEnergyEVSEMode {
	rv := objc.Send[MTRBaseClusterEnergyEVSEMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterEnergyEVSEMode creates a new MTRBaseClusterEnergyEVSEMode instance.
func NewMTRBaseClusterEnergyEVSEMode() MTRBaseClusterEnergyEVSEMode {
	return getMTRBaseClusterEnergyEVSEModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterEnergyEVSEMode */
// Cluster Energy EVSE Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Energy EVSE Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSEMode
type MTRBaseClusterEnergyEVSEMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterEnergyEVSEModeFrom constructs a [MTRBaseClusterEnergyEVSEMode] from an unsafe.Pointer.
//
// Cluster Energy EVSE Mode
func MTRBaseClusterEnergyEVSEModeFrom(ptr unsafe.Pointer) MTRBaseClusterEnergyEVSEMode {
	return MTRBaseClusterEnergyEVSEMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterEnergyEVSEMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterEnergyEVSEMode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSEMode/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEModeClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterEnergyEVSEMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterEnergyEVSEMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterEnergyEVSEMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterEnergyEVSEMode */



