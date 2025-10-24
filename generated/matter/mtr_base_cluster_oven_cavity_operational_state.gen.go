// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterOvenCavityOperationalState */


/* debug [class_header]: Header for MTRBaseClusterOvenCavityOperationalState */
// The class instance for the [MTRBaseClusterOvenCavityOperationalState] class.
var (
	MTRBaseClusterOvenCavityOperationalStateClass     _MTRBaseClusterOvenCavityOperationalStateClass
	MTRBaseClusterOvenCavityOperationalStateClassOnce sync.Once
)

func getMTRBaseClusterOvenCavityOperationalStateClass() _MTRBaseClusterOvenCavityOperationalStateClass {
	MTRBaseClusterOvenCavityOperationalStateClassOnce.Do(func() {
		MTRBaseClusterOvenCavityOperationalStateClass = _MTRBaseClusterOvenCavityOperationalStateClass{objc.GetClass("MTRBaseClusterOvenCavityOperationalState")}
	})
	return MTRBaseClusterOvenCavityOperationalStateClass
}

type _MTRBaseClusterOvenCavityOperationalStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterOvenCavityOperationalState */
// An interface definition for the [MTRBaseClusterOvenCavityOperationalState] class.
type IMTRBaseClusterOvenCavityOperationalState interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterOvenCavityOperationalState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterOvenCavityOperationalState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterOvenCavityOperationalState */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) Alloc() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) New() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOvenCavityOperationalState) Init() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOvenCavityOperationalState) Autorelease() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOvenCavityOperationalState creates a new MTRBaseClusterOvenCavityOperationalState instance.
func NewMTRBaseClusterOvenCavityOperationalState() MTRBaseClusterOvenCavityOperationalState {
	return getMTRBaseClusterOvenCavityOperationalStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterOvenCavityOperationalState */
// Cluster Oven Cavity Operational State
//
// This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.


// Cluster Oven Cavity Operational State
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState
type MTRBaseClusterOvenCavityOperationalState struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOvenCavityOperationalStateFrom constructs a [MTRBaseClusterOvenCavityOperationalState] from an unsafe.Pointer.
//
// Cluster Oven Cavity Operational State
func MTRBaseClusterOvenCavityOperationalStateFrom(ptr unsafe.Pointer) MTRBaseClusterOvenCavityOperationalState {
	return MTRBaseClusterOvenCavityOperationalState{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterOvenCavityOperationalState */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/init(device:endpointID:queue:)
func NewMTRBaseClusterOvenCavityOperationalStateWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterOvenCavityOperationalState {
	instance := getMTRBaseClusterOvenCavityOperationalStateClass().Alloc()
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterOvenCavityOperationalStateWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterOvenCavityOperationalState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterOvenCavityOperationalState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterOvenCavityOperationalState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterOvenCavityOperationalState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterOvenCavityOperationalState */


