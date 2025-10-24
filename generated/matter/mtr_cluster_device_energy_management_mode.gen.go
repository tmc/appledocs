// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRClusterDeviceEnergyManagementMode */


/* debug [class_header]: Header for MTRClusterDeviceEnergyManagementMode */
// The class instance for the [MTRClusterDeviceEnergyManagementMode] class.
var (
	MTRClusterDeviceEnergyManagementModeClass     _MTRClusterDeviceEnergyManagementModeClass
	MTRClusterDeviceEnergyManagementModeClassOnce sync.Once
)

func getMTRClusterDeviceEnergyManagementModeClass() _MTRClusterDeviceEnergyManagementModeClass {
	MTRClusterDeviceEnergyManagementModeClassOnce.Do(func() {
		MTRClusterDeviceEnergyManagementModeClass = _MTRClusterDeviceEnergyManagementModeClass{objc.GetClass("MTRClusterDeviceEnergyManagementMode")}
	})
	return MTRClusterDeviceEnergyManagementModeClass
}

type _MTRClusterDeviceEnergyManagementModeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterDeviceEnergyManagementMode */
// An interface definition for the [MTRClusterDeviceEnergyManagementMode] class.
type IMTRClusterDeviceEnergyManagementMode interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterDeviceEnergyManagementMode */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterDeviceEnergyManagementMode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterDeviceEnergyManagementMode */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDeviceEnergyManagementModeClass) Alloc() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterDeviceEnergyManagementModeClass) New() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDeviceEnergyManagementMode) Init() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDeviceEnergyManagementMode) Autorelease() MTRClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDeviceEnergyManagementMode creates a new MTRClusterDeviceEnergyManagementMode instance.
func NewMTRClusterDeviceEnergyManagementMode() MTRClusterDeviceEnergyManagementMode {
	return getMTRClusterDeviceEnergyManagementModeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterDeviceEnergyManagementMode */
// Cluster Device Energy Management Mode Attributes and commands for selecting a mode from a list of supported options.


// Cluster Device Energy Management Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagementMode
type MTRClusterDeviceEnergyManagementMode struct {
	MTRGenericCluster
}

// MTRClusterDeviceEnergyManagementModeFrom constructs a [MTRClusterDeviceEnergyManagementMode] from an unsafe.Pointer.
//
// Cluster Device Energy Management Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterDeviceEnergyManagementModeFrom(ptr unsafe.Pointer) MTRClusterDeviceEnergyManagementMode {
	return MTRClusterDeviceEnergyManagementMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterDeviceEnergyManagementMode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterDeviceEnergyManagementMode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterDeviceEnergyManagementMode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterDeviceEnergyManagementMode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterDeviceEnergyManagementMode */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterDeviceEnergyManagementMode */



