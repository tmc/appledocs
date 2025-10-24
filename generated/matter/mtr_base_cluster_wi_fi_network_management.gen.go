// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterWiFiNetworkManagement */


/* debug [class_header]: Header for MTRBaseClusterWiFiNetworkManagement */
// The class instance for the [MTRBaseClusterWiFiNetworkManagement] class.
var (
	MTRBaseClusterWiFiNetworkManagementClass     _MTRBaseClusterWiFiNetworkManagementClass
	MTRBaseClusterWiFiNetworkManagementClassOnce sync.Once
)

func getMTRBaseClusterWiFiNetworkManagementClass() _MTRBaseClusterWiFiNetworkManagementClass {
	MTRBaseClusterWiFiNetworkManagementClassOnce.Do(func() {
		MTRBaseClusterWiFiNetworkManagementClass = _MTRBaseClusterWiFiNetworkManagementClass{objc.GetClass("MTRBaseClusterWiFiNetworkManagement")}
	})
	return MTRBaseClusterWiFiNetworkManagementClass
}

type _MTRBaseClusterWiFiNetworkManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterWiFiNetworkManagement */
// An interface definition for the [MTRBaseClusterWiFiNetworkManagement] class.
type IMTRBaseClusterWiFiNetworkManagement interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterWiFiNetworkManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterWiFiNetworkManagement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterWiFiNetworkManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWiFiNetworkManagementClass) Alloc() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterWiFiNetworkManagementClass) New() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWiFiNetworkManagement) Init() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWiFiNetworkManagement) Autorelease() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWiFiNetworkManagement creates a new MTRBaseClusterWiFiNetworkManagement instance.
func NewMTRBaseClusterWiFiNetworkManagement() MTRBaseClusterWiFiNetworkManagement {
	return getMTRBaseClusterWiFiNetworkManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterWiFiNetworkManagement */
// Cluster Wi-Fi Network Management
//
// Functionality to retrieve operational information about a managed Wi-Fi network.


// Cluster Wi-Fi Network Management
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement
type MTRBaseClusterWiFiNetworkManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWiFiNetworkManagementFrom constructs a [MTRBaseClusterWiFiNetworkManagement] from an unsafe.Pointer.
//
// Cluster Wi-Fi Network Management
func MTRBaseClusterWiFiNetworkManagementFrom(ptr unsafe.Pointer) MTRBaseClusterWiFiNetworkManagement {
	return MTRBaseClusterWiFiNetworkManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterWiFiNetworkManagement */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterWiFiNetworkManagementWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterWiFiNetworkManagement {
	instance := getMTRBaseClusterWiFiNetworkManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterWiFiNetworkManagementWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterWiFiNetworkManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterWiFiNetworkManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterWiFiNetworkManagement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterWiFiNetworkManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterWiFiNetworkManagement */


