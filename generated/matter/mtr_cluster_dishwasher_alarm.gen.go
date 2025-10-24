// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterDishwasherAlarm */


/* debug [class_header]: Header for MTRClusterDishwasherAlarm */
// The class instance for the [MTRClusterDishwasherAlarm] class.
var (
	MTRClusterDishwasherAlarmClass     _MTRClusterDishwasherAlarmClass
	MTRClusterDishwasherAlarmClassOnce sync.Once
)

func getMTRClusterDishwasherAlarmClass() _MTRClusterDishwasherAlarmClass {
	MTRClusterDishwasherAlarmClassOnce.Do(func() {
		MTRClusterDishwasherAlarmClass = _MTRClusterDishwasherAlarmClass{objc.GetClass("MTRClusterDishwasherAlarm")}
	})
	return MTRClusterDishwasherAlarmClass
}

type _MTRClusterDishwasherAlarmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterDishwasherAlarm */
// An interface definition for the [MTRClusterDishwasherAlarm] class.
type IMTRClusterDishwasherAlarm interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterDishwasherAlarm */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterDishwasherAlarm */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterDishwasherAlarm */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDishwasherAlarmClass) Alloc() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterDishwasherAlarmClass) New() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDishwasherAlarm) Init() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDishwasherAlarm) Autorelease() MTRClusterDishwasherAlarm {
	rv := objc.Send[MTRClusterDishwasherAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDishwasherAlarm creates a new MTRClusterDishwasherAlarm instance.
func NewMTRClusterDishwasherAlarm() MTRClusterDishwasherAlarm {
	return getMTRClusterDishwasherAlarmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterDishwasherAlarm */
// Cluster Dishwasher Alarm Attributes and commands for configuring the Dishwasher alarm.


// Cluster Dishwasher Alarm Attributes and commands for configuring the Dishwasher alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm
type MTRClusterDishwasherAlarm struct {
	MTRGenericCluster
}

// MTRClusterDishwasherAlarmFrom constructs a [MTRClusterDishwasherAlarm] from an unsafe.Pointer.
//
// Cluster Dishwasher Alarm Attributes and commands for configuring the Dishwasher alarm.
func MTRClusterDishwasherAlarmFrom(ptr unsafe.Pointer) MTRClusterDishwasherAlarm {
	return MTRClusterDishwasherAlarm{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterDishwasherAlarm */

// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDishwasherAlarm/init(device:endpointID:queue:)
func NewMTRClusterDishwasherAlarmWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDishwasherAlarm {
	instance := getMTRClusterDishwasherAlarmClass().Alloc()
	rv := objc.Send[MTRClusterDishwasherAlarm](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRClusterDishwasherAlarmWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterDishwasherAlarm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterDishwasherAlarm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterDishwasherAlarm */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterDishwasherAlarm */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterDishwasherAlarm */


