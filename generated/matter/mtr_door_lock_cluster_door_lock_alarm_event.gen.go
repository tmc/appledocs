// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterDoorLockAlarmEvent */


/* debug [class_header]: Header for MTRDoorLockClusterDoorLockAlarmEvent */
// The class instance for the [MTRDoorLockClusterDoorLockAlarmEvent] class.
var (
	MTRDoorLockClusterDoorLockAlarmEventClass     _MTRDoorLockClusterDoorLockAlarmEventClass
	MTRDoorLockClusterDoorLockAlarmEventClassOnce sync.Once
)

func getMTRDoorLockClusterDoorLockAlarmEventClass() _MTRDoorLockClusterDoorLockAlarmEventClass {
	MTRDoorLockClusterDoorLockAlarmEventClassOnce.Do(func() {
		MTRDoorLockClusterDoorLockAlarmEventClass = _MTRDoorLockClusterDoorLockAlarmEventClass{objc.GetClass("MTRDoorLockClusterDoorLockAlarmEvent")}
	})
	return MTRDoorLockClusterDoorLockAlarmEventClass
}

type _MTRDoorLockClusterDoorLockAlarmEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterDoorLockAlarmEvent */
// An interface definition for the [MTRDoorLockClusterDoorLockAlarmEvent] class.
type IMTRDoorLockClusterDoorLockAlarmEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterDoorLockAlarmEvent */
	// properties:
	AlarmCode() objc.IObject /* cross-framework: NSNumber */
	SetAlarmCode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterDoorLockAlarmEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterDoorLockAlarmEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterDoorLockAlarmEventClass) Alloc() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterDoorLockAlarmEventClass) New() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) Init() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) Autorelease() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterDoorLockAlarmEvent creates a new MTRDoorLockClusterDoorLockAlarmEvent instance.
func NewMTRDoorLockClusterDoorLockAlarmEvent() MTRDoorLockClusterDoorLockAlarmEvent {
	return getMTRDoorLockClusterDoorLockAlarmEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterDoorLockAlarmEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorLockAlarmEvent
type MTRDoorLockClusterDoorLockAlarmEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterDoorLockAlarmEventFrom constructs a [MTRDoorLockClusterDoorLockAlarmEvent] from an unsafe.Pointer.
func MTRDoorLockClusterDoorLockAlarmEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterDoorLockAlarmEvent {
	return MTRDoorLockClusterDoorLockAlarmEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterDoorLockAlarmEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterDoorLockAlarmEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterDoorLockAlarmEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterDoorLockAlarmEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterDoorLockAlarmEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorLockAlarmEvent/alarmCode
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) AlarmCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarmCode"))
	return rv
}/* debug [instance_properties/getter]: alarmCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorLockAlarmEvent/alarmCode
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) SetAlarmCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmCode:"), value)
}/* debug [instance_properties/setter]: alarmCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterDoorLockAlarmEvent */



