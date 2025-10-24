// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterDeviceEnergyManagement */


/* debug [class_header]: Header for MTRClusterDeviceEnergyManagement */
// The class instance for the [MTRClusterDeviceEnergyManagement] class.
var (
	MTRClusterDeviceEnergyManagementClass     _MTRClusterDeviceEnergyManagementClass
	MTRClusterDeviceEnergyManagementClassOnce sync.Once
)

func getMTRClusterDeviceEnergyManagementClass() _MTRClusterDeviceEnergyManagementClass {
	MTRClusterDeviceEnergyManagementClassOnce.Do(func() {
		MTRClusterDeviceEnergyManagementClass = _MTRClusterDeviceEnergyManagementClass{objc.GetClass("MTRClusterDeviceEnergyManagement")}
	})
	return MTRClusterDeviceEnergyManagementClass
}

type _MTRClusterDeviceEnergyManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterDeviceEnergyManagement */
// An interface definition for the [MTRClusterDeviceEnergyManagement] class.
type IMTRClusterDeviceEnergyManagement interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterDeviceEnergyManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterDeviceEnergyManagement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterDeviceEnergyManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDeviceEnergyManagementClass) Alloc() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterDeviceEnergyManagementClass) New() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDeviceEnergyManagement) Init() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDeviceEnergyManagement) Autorelease() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDeviceEnergyManagement creates a new MTRClusterDeviceEnergyManagement instance.
func NewMTRClusterDeviceEnergyManagement() MTRClusterDeviceEnergyManagement {
	return getMTRClusterDeviceEnergyManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterDeviceEnergyManagement */
// Cluster Device Energy Management This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).


// Cluster Device Energy Management This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement
type MTRClusterDeviceEnergyManagement struct {
	MTRGenericCluster
}

// MTRClusterDeviceEnergyManagementFrom constructs a [MTRClusterDeviceEnergyManagement] from an unsafe.Pointer.
//
// Cluster Device Energy Management This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).
func MTRClusterDeviceEnergyManagementFrom(ptr unsafe.Pointer) MTRClusterDeviceEnergyManagement {
	return MTRClusterDeviceEnergyManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterDeviceEnergyManagement */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/init(device:endpointID:queue:)
func NewMTRClusterDeviceEnergyManagementWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDeviceEnergyManagement {
	instance := getMTRClusterDeviceEnergyManagementClass().Alloc()
	rv := objc.Send[MTRClusterDeviceEnergyManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterDeviceEnergyManagementWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterDeviceEnergyManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterDeviceEnergyManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterDeviceEnergyManagement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterDeviceEnergyManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterDeviceEnergyManagement */


