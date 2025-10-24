// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterTemperatureControl */


/* debug [class_header]: Header for MTRClusterTemperatureControl */
// The class instance for the [MTRClusterTemperatureControl] class.
var (
	MTRClusterTemperatureControlClass     _MTRClusterTemperatureControlClass
	MTRClusterTemperatureControlClassOnce sync.Once
)

func getMTRClusterTemperatureControlClass() _MTRClusterTemperatureControlClass {
	MTRClusterTemperatureControlClassOnce.Do(func() {
		MTRClusterTemperatureControlClass = _MTRClusterTemperatureControlClass{objc.GetClass("MTRClusterTemperatureControl")}
	})
	return MTRClusterTemperatureControlClass
}

type _MTRClusterTemperatureControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterTemperatureControl */
// An interface definition for the [MTRClusterTemperatureControl] class.
type IMTRClusterTemperatureControl interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterTemperatureControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterTemperatureControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterTemperatureControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTemperatureControlClass) Alloc() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterTemperatureControlClass) New() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTemperatureControl) Init() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTemperatureControl) Autorelease() MTRClusterTemperatureControl {
	rv := objc.Send[MTRClusterTemperatureControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTemperatureControl creates a new MTRClusterTemperatureControl instance.
func NewMTRClusterTemperatureControl() MTRClusterTemperatureControl {
	return getMTRClusterTemperatureControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterTemperatureControl */
// Cluster Temperature Control Attributes and commands for configuring the temperature control, and reporting temperature.


// Cluster Temperature Control Attributes and commands for configuring the temperature control, and reporting temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl
type MTRClusterTemperatureControl struct {
	MTRGenericCluster
}

// MTRClusterTemperatureControlFrom constructs a [MTRClusterTemperatureControl] from an unsafe.Pointer.
//
// Cluster Temperature Control Attributes and commands for configuring the temperature control, and reporting temperature.
func MTRClusterTemperatureControlFrom(ptr unsafe.Pointer) MTRClusterTemperatureControl {
	return MTRClusterTemperatureControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterTemperatureControl */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureControl/init(device:endpointID:queue:)
func NewMTRClusterTemperatureControlWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterTemperatureControl {
	instance := getMTRClusterTemperatureControlClass().Alloc()
	rv := objc.Send[MTRClusterTemperatureControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterTemperatureControlWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterTemperatureControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterTemperatureControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterTemperatureControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterTemperatureControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterTemperatureControl */


