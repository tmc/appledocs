// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterDeviceEnergyManagementMode */


/* debug [class_header]: Header for MTRBaseClusterDeviceEnergyManagementMode */
// The class instance for the [MTRBaseClusterDeviceEnergyManagementMode] class.
var (
	MTRBaseClusterDeviceEnergyManagementModeClass     _MTRBaseClusterDeviceEnergyManagementModeClass
	MTRBaseClusterDeviceEnergyManagementModeClassOnce sync.Once
)

func getMTRBaseClusterDeviceEnergyManagementModeClass() _MTRBaseClusterDeviceEnergyManagementModeClass {
	MTRBaseClusterDeviceEnergyManagementModeClassOnce.Do(func() {
		MTRBaseClusterDeviceEnergyManagementModeClass = _MTRBaseClusterDeviceEnergyManagementModeClass{objc.GetClass("MTRBaseClusterDeviceEnergyManagementMode")}
	})
	return MTRBaseClusterDeviceEnergyManagementModeClass
}

type _MTRBaseClusterDeviceEnergyManagementModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterDeviceEnergyManagementMode */
// An interface definition for the [MTRBaseClusterDeviceEnergyManagementMode] class.
type IMTRBaseClusterDeviceEnergyManagementMode interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterDeviceEnergyManagementMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterDeviceEnergyManagementMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterDeviceEnergyManagementMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) Alloc() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) New() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDeviceEnergyManagementMode) Init() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDeviceEnergyManagementMode) Autorelease() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDeviceEnergyManagementMode creates a new MTRBaseClusterDeviceEnergyManagementMode instance.
func NewMTRBaseClusterDeviceEnergyManagementMode() MTRBaseClusterDeviceEnergyManagementMode {
	return getMTRBaseClusterDeviceEnergyManagementModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterDeviceEnergyManagementMode */
// Cluster Device Energy Management Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Device Energy Management Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode
type MTRBaseClusterDeviceEnergyManagementMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDeviceEnergyManagementModeFrom constructs a [MTRBaseClusterDeviceEnergyManagementMode] from an unsafe.Pointer.
//
// Cluster Device Energy Management Mode
func MTRBaseClusterDeviceEnergyManagementModeFrom(ptr unsafe.Pointer) MTRBaseClusterDeviceEnergyManagementMode {
	return MTRBaseClusterDeviceEnergyManagementMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterDeviceEnergyManagementMode */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/init(device:endpointID:queue:)
func NewMTRBaseClusterDeviceEnergyManagementModeWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterDeviceEnergyManagementMode {
	instance := getMTRBaseClusterDeviceEnergyManagementModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterDeviceEnergyManagementModeWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterDeviceEnergyManagementMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterDeviceEnergyManagementMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterDeviceEnergyManagementMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterDeviceEnergyManagementMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterDeviceEnergyManagementMode */


