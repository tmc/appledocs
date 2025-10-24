// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterDishwasherAlarm */


/* debug [class_header]: Header for MTRBaseClusterDishwasherAlarm */
// The class instance for the [MTRBaseClusterDishwasherAlarm] class.
var (
	MTRBaseClusterDishwasherAlarmClass     _MTRBaseClusterDishwasherAlarmClass
	MTRBaseClusterDishwasherAlarmClassOnce sync.Once
)

func getMTRBaseClusterDishwasherAlarmClass() _MTRBaseClusterDishwasherAlarmClass {
	MTRBaseClusterDishwasherAlarmClassOnce.Do(func() {
		MTRBaseClusterDishwasherAlarmClass = _MTRBaseClusterDishwasherAlarmClass{objc.GetClass("MTRBaseClusterDishwasherAlarm")}
	})
	return MTRBaseClusterDishwasherAlarmClass
}

type _MTRBaseClusterDishwasherAlarmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterDishwasherAlarm */
// An interface definition for the [MTRBaseClusterDishwasherAlarm] class.
type IMTRBaseClusterDishwasherAlarm interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterDishwasherAlarm */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterDishwasherAlarm */
	// methods:
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterDishwasherAlarm */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDishwasherAlarmClass) Alloc() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterDishwasherAlarmClass) New() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDishwasherAlarm) Init() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDishwasherAlarm) Autorelease() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDishwasherAlarm creates a new MTRBaseClusterDishwasherAlarm instance.
func NewMTRBaseClusterDishwasherAlarm() MTRBaseClusterDishwasherAlarm {
	return getMTRBaseClusterDishwasherAlarmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterDishwasherAlarm */
// Cluster Dishwasher Alarm
//
// Attributes and commands for configuring the Dishwasher alarm.


// Cluster Dishwasher Alarm
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm
type MTRBaseClusterDishwasherAlarm struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDishwasherAlarmFrom constructs a [MTRBaseClusterDishwasherAlarm] from an unsafe.Pointer.
//
// Cluster Dishwasher Alarm
func MTRBaseClusterDishwasherAlarmFrom(ptr unsafe.Pointer) MTRBaseClusterDishwasherAlarm {
	return MTRBaseClusterDishwasherAlarm{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterDishwasherAlarm *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterDishwasherAlarm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterDishwasherAlarm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterDishwasherAlarm */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterDishwasherAlarm */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterDishwasherAlarm */



