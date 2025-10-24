// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterWiFiNetworkManagement */


/* debug [class_header]: Header for MTRClusterWiFiNetworkManagement */
// The class instance for the [MTRClusterWiFiNetworkManagement] class.
var (
	MTRClusterWiFiNetworkManagementClass     _MTRClusterWiFiNetworkManagementClass
	MTRClusterWiFiNetworkManagementClassOnce sync.Once
)

func getMTRClusterWiFiNetworkManagementClass() _MTRClusterWiFiNetworkManagementClass {
	MTRClusterWiFiNetworkManagementClassOnce.Do(func() {
		MTRClusterWiFiNetworkManagementClass = _MTRClusterWiFiNetworkManagementClass{objc.GetClass("MTRClusterWiFiNetworkManagement")}
	})
	return MTRClusterWiFiNetworkManagementClass
}

type _MTRClusterWiFiNetworkManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterWiFiNetworkManagement */
// An interface definition for the [MTRClusterWiFiNetworkManagement] class.
type IMTRClusterWiFiNetworkManagement interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterWiFiNetworkManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterWiFiNetworkManagement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterWiFiNetworkManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWiFiNetworkManagementClass) Alloc() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterWiFiNetworkManagementClass) New() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWiFiNetworkManagement) Init() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWiFiNetworkManagement) Autorelease() MTRClusterWiFiNetworkManagement {
	rv := objc.Send[MTRClusterWiFiNetworkManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWiFiNetworkManagement creates a new MTRClusterWiFiNetworkManagement instance.
func NewMTRClusterWiFiNetworkManagement() MTRClusterWiFiNetworkManagement {
	return getMTRClusterWiFiNetworkManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterWiFiNetworkManagement */
// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.


// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement
type MTRClusterWiFiNetworkManagement struct {
	MTRGenericCluster
}

// MTRClusterWiFiNetworkManagementFrom constructs a [MTRClusterWiFiNetworkManagement] from an unsafe.Pointer.
//
// Cluster Wi-Fi Network Management Functionality to retrieve operational information about a managed Wi-Fi network.
func MTRClusterWiFiNetworkManagementFrom(ptr unsafe.Pointer) MTRClusterWiFiNetworkManagement {
	return MTRClusterWiFiNetworkManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterWiFiNetworkManagement */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWiFiNetworkManagement/init(device:endpointID:queue:)
func NewMTRClusterWiFiNetworkManagementWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterWiFiNetworkManagement {
	instance := getMTRClusterWiFiNetworkManagementClass().Alloc()
	rv := objc.Send[MTRClusterWiFiNetworkManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterWiFiNetworkManagementWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterWiFiNetworkManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterWiFiNetworkManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterWiFiNetworkManagement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterWiFiNetworkManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterWiFiNetworkManagement */


