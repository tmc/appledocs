// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterDeviceEnergyManagement */


/* debug [class_header]: Header for MTRBaseClusterDeviceEnergyManagement */
// The class instance for the [MTRBaseClusterDeviceEnergyManagement] class.
var (
	MTRBaseClusterDeviceEnergyManagementClass     _MTRBaseClusterDeviceEnergyManagementClass
	MTRBaseClusterDeviceEnergyManagementClassOnce sync.Once
)

func getMTRBaseClusterDeviceEnergyManagementClass() _MTRBaseClusterDeviceEnergyManagementClass {
	MTRBaseClusterDeviceEnergyManagementClassOnce.Do(func() {
		MTRBaseClusterDeviceEnergyManagementClass = _MTRBaseClusterDeviceEnergyManagementClass{objc.GetClass("MTRBaseClusterDeviceEnergyManagement")}
	})
	return MTRBaseClusterDeviceEnergyManagementClass
}

type _MTRBaseClusterDeviceEnergyManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterDeviceEnergyManagement */
// An interface definition for the [MTRBaseClusterDeviceEnergyManagement] class.
type IMTRBaseClusterDeviceEnergyManagement interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterDeviceEnergyManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterDeviceEnergyManagement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterDeviceEnergyManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDeviceEnergyManagementClass) Alloc() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterDeviceEnergyManagementClass) New() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDeviceEnergyManagement) Init() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDeviceEnergyManagement) Autorelease() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDeviceEnergyManagement creates a new MTRBaseClusterDeviceEnergyManagement instance.
func NewMTRBaseClusterDeviceEnergyManagement() MTRBaseClusterDeviceEnergyManagement {
	return getMTRBaseClusterDeviceEnergyManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterDeviceEnergyManagement */
// Cluster Device Energy Management
//
// This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).


// Cluster Device Energy Management
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement
type MTRBaseClusterDeviceEnergyManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDeviceEnergyManagementFrom constructs a [MTRBaseClusterDeviceEnergyManagement] from an unsafe.Pointer.
//
// Cluster Device Energy Management
func MTRBaseClusterDeviceEnergyManagementFrom(ptr unsafe.Pointer) MTRBaseClusterDeviceEnergyManagement {
	return MTRBaseClusterDeviceEnergyManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterDeviceEnergyManagement */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterDeviceEnergyManagementWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterDeviceEnergyManagement {
	instance := getMTRBaseClusterDeviceEnergyManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterDeviceEnergyManagementWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterDeviceEnergyManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterDeviceEnergyManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterDeviceEnergyManagement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterDeviceEnergyManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterDeviceEnergyManagement */


