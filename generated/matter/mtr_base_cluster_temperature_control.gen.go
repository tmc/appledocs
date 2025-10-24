// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterTemperatureControl */


/* debug [class_header]: Header for MTRBaseClusterTemperatureControl */
// The class instance for the [MTRBaseClusterTemperatureControl] class.
var (
	MTRBaseClusterTemperatureControlClass     _MTRBaseClusterTemperatureControlClass
	MTRBaseClusterTemperatureControlClassOnce sync.Once
)

func getMTRBaseClusterTemperatureControlClass() _MTRBaseClusterTemperatureControlClass {
	MTRBaseClusterTemperatureControlClassOnce.Do(func() {
		MTRBaseClusterTemperatureControlClass = _MTRBaseClusterTemperatureControlClass{objc.GetClass("MTRBaseClusterTemperatureControl")}
	})
	return MTRBaseClusterTemperatureControlClass
}

type _MTRBaseClusterTemperatureControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterTemperatureControl */
// An interface definition for the [MTRBaseClusterTemperatureControl] class.
type IMTRBaseClusterTemperatureControl interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterTemperatureControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterTemperatureControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterTemperatureControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTemperatureControlClass) Alloc() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterTemperatureControlClass) New() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTemperatureControl) Init() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTemperatureControl) Autorelease() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTemperatureControl creates a new MTRBaseClusterTemperatureControl instance.
func NewMTRBaseClusterTemperatureControl() MTRBaseClusterTemperatureControl {
	return getMTRBaseClusterTemperatureControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterTemperatureControl */
// Cluster Temperature Control
//
// Attributes and commands for configuring the temperature control, and reporting temperature.


// Cluster Temperature Control
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl
type MTRBaseClusterTemperatureControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTemperatureControlFrom constructs a [MTRBaseClusterTemperatureControl] from an unsafe.Pointer.
//
// Cluster Temperature Control
func MTRBaseClusterTemperatureControlFrom(ptr unsafe.Pointer) MTRBaseClusterTemperatureControl {
	return MTRBaseClusterTemperatureControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterTemperatureControl */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/init(device:endpointID:queue:)
func NewMTRBaseClusterTemperatureControlWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterTemperatureControl {
	instance := getMTRBaseClusterTemperatureControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterTemperatureControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterTemperatureControlWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterTemperatureControl */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterTemperatureControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterTemperatureControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterTemperatureControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterTemperatureControl */


