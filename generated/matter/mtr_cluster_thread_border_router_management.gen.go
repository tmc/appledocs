// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterThreadBorderRouterManagement */


/* debug [class_header]: Header for MTRClusterThreadBorderRouterManagement */
// The class instance for the [MTRClusterThreadBorderRouterManagement] class.
var (
	MTRClusterThreadBorderRouterManagementClass     _MTRClusterThreadBorderRouterManagementClass
	MTRClusterThreadBorderRouterManagementClassOnce sync.Once
)

func getMTRClusterThreadBorderRouterManagementClass() _MTRClusterThreadBorderRouterManagementClass {
	MTRClusterThreadBorderRouterManagementClassOnce.Do(func() {
		MTRClusterThreadBorderRouterManagementClass = _MTRClusterThreadBorderRouterManagementClass{objc.GetClass("MTRClusterThreadBorderRouterManagement")}
	})
	return MTRClusterThreadBorderRouterManagementClass
}

type _MTRClusterThreadBorderRouterManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterThreadBorderRouterManagement */
// An interface definition for the [MTRClusterThreadBorderRouterManagement] class.
type IMTRClusterThreadBorderRouterManagement interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterThreadBorderRouterManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterThreadBorderRouterManagement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterThreadBorderRouterManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterThreadBorderRouterManagementClass) Alloc() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterThreadBorderRouterManagementClass) New() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterThreadBorderRouterManagement) Init() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterThreadBorderRouterManagement) Autorelease() MTRClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterThreadBorderRouterManagement creates a new MTRClusterThreadBorderRouterManagement instance.
func NewMTRClusterThreadBorderRouterManagement() MTRClusterThreadBorderRouterManagement {
	return getMTRClusterThreadBorderRouterManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterThreadBorderRouterManagement */
// Cluster Thread Border Router Management Manage the Thread network of Thread Border Router


// Cluster Thread Border Router Management Manage the Thread network of Thread Border Router
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement
type MTRClusterThreadBorderRouterManagement struct {
	MTRGenericCluster
}

// MTRClusterThreadBorderRouterManagementFrom constructs a [MTRClusterThreadBorderRouterManagement] from an unsafe.Pointer.
//
// Cluster Thread Border Router Management Manage the Thread network of Thread Border Router
func MTRClusterThreadBorderRouterManagementFrom(ptr unsafe.Pointer) MTRClusterThreadBorderRouterManagement {
	return MTRClusterThreadBorderRouterManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterThreadBorderRouterManagement */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterThreadBorderRouterManagement/init(device:endpointID:queue:)
func NewMTRClusterThreadBorderRouterManagementWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterThreadBorderRouterManagement {
	instance := getMTRClusterThreadBorderRouterManagementClass().Alloc()
	rv := objc.Send[MTRClusterThreadBorderRouterManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterThreadBorderRouterManagementWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterThreadBorderRouterManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterThreadBorderRouterManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterThreadBorderRouterManagement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterThreadBorderRouterManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterThreadBorderRouterManagement */


