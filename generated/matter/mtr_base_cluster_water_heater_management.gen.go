// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterWaterHeaterManagement */


/* debug [class_header]: Header for MTRBaseClusterWaterHeaterManagement */
// The class instance for the [MTRBaseClusterWaterHeaterManagement] class.
var (
	MTRBaseClusterWaterHeaterManagementClass     _MTRBaseClusterWaterHeaterManagementClass
	MTRBaseClusterWaterHeaterManagementClassOnce sync.Once
)

func getMTRBaseClusterWaterHeaterManagementClass() _MTRBaseClusterWaterHeaterManagementClass {
	MTRBaseClusterWaterHeaterManagementClassOnce.Do(func() {
		MTRBaseClusterWaterHeaterManagementClass = _MTRBaseClusterWaterHeaterManagementClass{objc.GetClass("MTRBaseClusterWaterHeaterManagement")}
	})
	return MTRBaseClusterWaterHeaterManagementClass
}

type _MTRBaseClusterWaterHeaterManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterWaterHeaterManagement */
// An interface definition for the [MTRBaseClusterWaterHeaterManagement] class.
type IMTRBaseClusterWaterHeaterManagement interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterWaterHeaterManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterWaterHeaterManagement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterWaterHeaterManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWaterHeaterManagementClass) Alloc() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterWaterHeaterManagementClass) New() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWaterHeaterManagement) Init() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWaterHeaterManagement) Autorelease() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWaterHeaterManagement creates a new MTRBaseClusterWaterHeaterManagement instance.
func NewMTRBaseClusterWaterHeaterManagement() MTRBaseClusterWaterHeaterManagement {
	return getMTRBaseClusterWaterHeaterManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterWaterHeaterManagement */
// Cluster Water Heater Management
//
// This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.


// Cluster Water Heater Management
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement
type MTRBaseClusterWaterHeaterManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWaterHeaterManagementFrom constructs a [MTRBaseClusterWaterHeaterManagement] from an unsafe.Pointer.
//
// Cluster Water Heater Management
func MTRBaseClusterWaterHeaterManagementFrom(ptr unsafe.Pointer) MTRBaseClusterWaterHeaterManagement {
	return MTRBaseClusterWaterHeaterManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterWaterHeaterManagement */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterWaterHeaterManagementWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterWaterHeaterManagement {
	instance := getMTRBaseClusterWaterHeaterManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterWaterHeaterManagementWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterWaterHeaterManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterWaterHeaterManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterWaterHeaterManagement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterWaterHeaterManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterWaterHeaterManagement */


