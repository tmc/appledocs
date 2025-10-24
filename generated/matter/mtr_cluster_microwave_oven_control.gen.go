// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterMicrowaveOvenControl */


/* debug [class_header]: Header for MTRClusterMicrowaveOvenControl */
// The class instance for the [MTRClusterMicrowaveOvenControl] class.
var (
	MTRClusterMicrowaveOvenControlClass     _MTRClusterMicrowaveOvenControlClass
	MTRClusterMicrowaveOvenControlClassOnce sync.Once
)

func getMTRClusterMicrowaveOvenControlClass() _MTRClusterMicrowaveOvenControlClass {
	MTRClusterMicrowaveOvenControlClassOnce.Do(func() {
		MTRClusterMicrowaveOvenControlClass = _MTRClusterMicrowaveOvenControlClass{objc.GetClass("MTRClusterMicrowaveOvenControl")}
	})
	return MTRClusterMicrowaveOvenControlClass
}

type _MTRClusterMicrowaveOvenControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterMicrowaveOvenControl */
// An interface definition for the [MTRClusterMicrowaveOvenControl] class.
type IMTRClusterMicrowaveOvenControl interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterMicrowaveOvenControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterMicrowaveOvenControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterMicrowaveOvenControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMicrowaveOvenControlClass) Alloc() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterMicrowaveOvenControlClass) New() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMicrowaveOvenControl) Init() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMicrowaveOvenControl) Autorelease() MTRClusterMicrowaveOvenControl {
	rv := objc.Send[MTRClusterMicrowaveOvenControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMicrowaveOvenControl creates a new MTRClusterMicrowaveOvenControl instance.
func NewMTRClusterMicrowaveOvenControl() MTRClusterMicrowaveOvenControl {
	return getMTRClusterMicrowaveOvenControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterMicrowaveOvenControl */
// Cluster Microwave Oven Control Attributes and commands for configuring the microwave oven control, and reporting cooking stats.


// Cluster Microwave Oven Control Attributes and commands for configuring the microwave oven control, and reporting cooking stats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl
type MTRClusterMicrowaveOvenControl struct {
	MTRGenericCluster
}

// MTRClusterMicrowaveOvenControlFrom constructs a [MTRClusterMicrowaveOvenControl] from an unsafe.Pointer.
//
// Cluster Microwave Oven Control Attributes and commands for configuring the microwave oven control, and reporting cooking stats.
func MTRClusterMicrowaveOvenControlFrom(ptr unsafe.Pointer) MTRClusterMicrowaveOvenControl {
	return MTRClusterMicrowaveOvenControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterMicrowaveOvenControl */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMicrowaveOvenControl/init(device:endpointID:queue:)
func NewMTRClusterMicrowaveOvenControlWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterMicrowaveOvenControl {
	instance := getMTRClusterMicrowaveOvenControlClass().Alloc()
	rv := objc.Send[MTRClusterMicrowaveOvenControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterMicrowaveOvenControlWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterMicrowaveOvenControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterMicrowaveOvenControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterMicrowaveOvenControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterMicrowaveOvenControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterMicrowaveOvenControl */


