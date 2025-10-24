// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterRefrigeratorAlarm */


/* debug [class_header]: Header for MTRClusterRefrigeratorAlarm */
// The class instance for the [MTRClusterRefrigeratorAlarm] class.
var (
	MTRClusterRefrigeratorAlarmClass     _MTRClusterRefrigeratorAlarmClass
	MTRClusterRefrigeratorAlarmClassOnce sync.Once
)

func getMTRClusterRefrigeratorAlarmClass() _MTRClusterRefrigeratorAlarmClass {
	MTRClusterRefrigeratorAlarmClassOnce.Do(func() {
		MTRClusterRefrigeratorAlarmClass = _MTRClusterRefrigeratorAlarmClass{objc.GetClass("MTRClusterRefrigeratorAlarm")}
	})
	return MTRClusterRefrigeratorAlarmClass
}

type _MTRClusterRefrigeratorAlarmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterRefrigeratorAlarm */
// An interface definition for the [MTRClusterRefrigeratorAlarm] class.
type IMTRClusterRefrigeratorAlarm interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterRefrigeratorAlarm */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterRefrigeratorAlarm */
	// methods:
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterRefrigeratorAlarm */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRefrigeratorAlarmClass) Alloc() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterRefrigeratorAlarmClass) New() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRefrigeratorAlarm) Init() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRefrigeratorAlarm) Autorelease() MTRClusterRefrigeratorAlarm {
	rv := objc.Send[MTRClusterRefrigeratorAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRefrigeratorAlarm creates a new MTRClusterRefrigeratorAlarm instance.
func NewMTRClusterRefrigeratorAlarm() MTRClusterRefrigeratorAlarm {
	return getMTRClusterRefrigeratorAlarmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterRefrigeratorAlarm */
// Cluster Refrigerator Alarm Attributes and commands for configuring the Refrigerator alarm.


// Cluster Refrigerator Alarm Attributes and commands for configuring the Refrigerator alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm
type MTRClusterRefrigeratorAlarm struct {
	MTRGenericCluster
}

// MTRClusterRefrigeratorAlarmFrom constructs a [MTRClusterRefrigeratorAlarm] from an unsafe.Pointer.
//
// Cluster Refrigerator Alarm Attributes and commands for configuring the Refrigerator alarm.
func MTRClusterRefrigeratorAlarmFrom(ptr unsafe.Pointer) MTRClusterRefrigeratorAlarm {
	return MTRClusterRefrigeratorAlarm{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterRefrigeratorAlarm *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterRefrigeratorAlarm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterRefrigeratorAlarm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterRefrigeratorAlarm */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAlarm/readAttributeFeatureMap(with:)
func (m_ MTRClusterRefrigeratorAlarm) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterRefrigeratorAlarm */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterRefrigeratorAlarm */



