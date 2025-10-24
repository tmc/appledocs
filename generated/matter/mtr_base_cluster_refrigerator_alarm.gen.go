// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterRefrigeratorAlarm */


/* debug [class_header]: Header for MTRBaseClusterRefrigeratorAlarm */
// The class instance for the [MTRBaseClusterRefrigeratorAlarm] class.
var (
	MTRBaseClusterRefrigeratorAlarmClass     _MTRBaseClusterRefrigeratorAlarmClass
	MTRBaseClusterRefrigeratorAlarmClassOnce sync.Once
)

func getMTRBaseClusterRefrigeratorAlarmClass() _MTRBaseClusterRefrigeratorAlarmClass {
	MTRBaseClusterRefrigeratorAlarmClassOnce.Do(func() {
		MTRBaseClusterRefrigeratorAlarmClass = _MTRBaseClusterRefrigeratorAlarmClass{objc.GetClass("MTRBaseClusterRefrigeratorAlarm")}
	})
	return MTRBaseClusterRefrigeratorAlarmClass
}

type _MTRBaseClusterRefrigeratorAlarmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterRefrigeratorAlarm */
// An interface definition for the [MTRBaseClusterRefrigeratorAlarm] class.
type IMTRBaseClusterRefrigeratorAlarm interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterRefrigeratorAlarm */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterRefrigeratorAlarm */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterRefrigeratorAlarm */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRefrigeratorAlarmClass) Alloc() MTRBaseClusterRefrigeratorAlarm {
	rv := objc.Send[MTRBaseClusterRefrigeratorAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterRefrigeratorAlarmClass) New() MTRBaseClusterRefrigeratorAlarm {
	rv := objc.Send[MTRBaseClusterRefrigeratorAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRefrigeratorAlarm) Init() MTRBaseClusterRefrigeratorAlarm {
	rv := objc.Send[MTRBaseClusterRefrigeratorAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRefrigeratorAlarm) Autorelease() MTRBaseClusterRefrigeratorAlarm {
	rv := objc.Send[MTRBaseClusterRefrigeratorAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRefrigeratorAlarm creates a new MTRBaseClusterRefrigeratorAlarm instance.
func NewMTRBaseClusterRefrigeratorAlarm() MTRBaseClusterRefrigeratorAlarm {
	return getMTRBaseClusterRefrigeratorAlarmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterRefrigeratorAlarm */
// Cluster Refrigerator Alarm
//
// Attributes and commands for configuring the Refrigerator alarm.


// Cluster Refrigerator Alarm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAlarm
type MTRBaseClusterRefrigeratorAlarm struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRefrigeratorAlarmFrom constructs a [MTRBaseClusterRefrigeratorAlarm] from an unsafe.Pointer.
//
// Cluster Refrigerator Alarm
func MTRBaseClusterRefrigeratorAlarmFrom(ptr unsafe.Pointer) MTRBaseClusterRefrigeratorAlarm {
	return MTRBaseClusterRefrigeratorAlarm{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterRefrigeratorAlarm */

// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAlarm/init(device:endpointID:queue:)
func NewMTRBaseClusterRefrigeratorAlarmWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterRefrigeratorAlarm {
	instance := getMTRBaseClusterRefrigeratorAlarmClass().Alloc()
	rv := objc.Send[MTRBaseClusterRefrigeratorAlarm](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterRefrigeratorAlarmWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterRefrigeratorAlarm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterRefrigeratorAlarm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterRefrigeratorAlarm */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterRefrigeratorAlarm */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterRefrigeratorAlarm */


