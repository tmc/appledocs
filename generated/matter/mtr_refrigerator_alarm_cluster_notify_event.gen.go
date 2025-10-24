// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRRefrigeratorAlarmClusterNotifyEvent */


/* debug [class_header]: Header for MTRRefrigeratorAlarmClusterNotifyEvent */
// The class instance for the [MTRRefrigeratorAlarmClusterNotifyEvent] class.
var (
	MTRRefrigeratorAlarmClusterNotifyEventClass     _MTRRefrigeratorAlarmClusterNotifyEventClass
	MTRRefrigeratorAlarmClusterNotifyEventClassOnce sync.Once
)

func getMTRRefrigeratorAlarmClusterNotifyEventClass() _MTRRefrigeratorAlarmClusterNotifyEventClass {
	MTRRefrigeratorAlarmClusterNotifyEventClassOnce.Do(func() {
		MTRRefrigeratorAlarmClusterNotifyEventClass = _MTRRefrigeratorAlarmClusterNotifyEventClass{objc.GetClass("MTRRefrigeratorAlarmClusterNotifyEvent")}
	})
	return MTRRefrigeratorAlarmClusterNotifyEventClass
}

type _MTRRefrigeratorAlarmClusterNotifyEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRRefrigeratorAlarmClusterNotifyEvent */
// An interface definition for the [MTRRefrigeratorAlarmClusterNotifyEvent] class.
type IMTRRefrigeratorAlarmClusterNotifyEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRRefrigeratorAlarmClusterNotifyEvent */
	// properties:
	Inactive() objc.IObject /* cross-framework: NSNumber */
	SetInactive(value objc.IObject /* cross-framework: NSNumber */)
	Mask() objc.IObject /* cross-framework: NSNumber */
	SetMask(value objc.IObject /* cross-framework: NSNumber */)
	Active() objc.IObject /* cross-framework: NSNumber */
	SetActive(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRRefrigeratorAlarmClusterNotifyEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRRefrigeratorAlarmClusterNotifyEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAlarmClusterNotifyEventClass) Alloc() MTRRefrigeratorAlarmClusterNotifyEvent {
	rv := objc.Send[MTRRefrigeratorAlarmClusterNotifyEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRRefrigeratorAlarmClusterNotifyEventClass) New() MTRRefrigeratorAlarmClusterNotifyEvent {
	rv := objc.Send[MTRRefrigeratorAlarmClusterNotifyEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Init() MTRRefrigeratorAlarmClusterNotifyEvent {
	rv := objc.Send[MTRRefrigeratorAlarmClusterNotifyEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Autorelease() MTRRefrigeratorAlarmClusterNotifyEvent {
	rv := objc.Send[MTRRefrigeratorAlarmClusterNotifyEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAlarmClusterNotifyEvent creates a new MTRRefrigeratorAlarmClusterNotifyEvent instance.
func NewMTRRefrigeratorAlarmClusterNotifyEvent() MTRRefrigeratorAlarmClusterNotifyEvent {
	return getMTRRefrigeratorAlarmClusterNotifyEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRRefrigeratorAlarmClusterNotifyEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent
type MTRRefrigeratorAlarmClusterNotifyEvent struct {
	objectivec.Object
}

// MTRRefrigeratorAlarmClusterNotifyEventFrom constructs a [MTRRefrigeratorAlarmClusterNotifyEvent] from an unsafe.Pointer.
func MTRRefrigeratorAlarmClusterNotifyEventFrom(ptr unsafe.Pointer) MTRRefrigeratorAlarmClusterNotifyEvent {
	return MTRRefrigeratorAlarmClusterNotifyEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRRefrigeratorAlarmClusterNotifyEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRRefrigeratorAlarmClusterNotifyEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRRefrigeratorAlarmClusterNotifyEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRRefrigeratorAlarmClusterNotifyEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRRefrigeratorAlarmClusterNotifyEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/inactive
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Inactive() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("inactive"))
	return rv
}/* debug [instance_properties/getter]: inactive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/inactive
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetInactive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInactive:"), value)
}/* debug [instance_properties/setter]: inactive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/mask
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Mask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mask"))
	return rv
}/* debug [instance_properties/getter]: mask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/mask
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}/* debug [instance_properties/setter]: mask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratoralarmclusternotifyevent/active
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Active() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratoralarmclusternotifyevent/active
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetActive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActive:"), value)
}/* debug [instance_properties/setter]: active */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratoralarmclusternotifyevent/state
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratoralarmclusternotifyevent/state
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRRefrigeratorAlarmClusterNotifyEvent */



