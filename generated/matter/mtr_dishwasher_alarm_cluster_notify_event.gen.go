// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherAlarmClusterNotifyEvent */


/* debug [class_header]: Header for MTRDishwasherAlarmClusterNotifyEvent */
// The class instance for the [MTRDishwasherAlarmClusterNotifyEvent] class.
var (
	MTRDishwasherAlarmClusterNotifyEventClass     _MTRDishwasherAlarmClusterNotifyEventClass
	MTRDishwasherAlarmClusterNotifyEventClassOnce sync.Once
)

func getMTRDishwasherAlarmClusterNotifyEventClass() _MTRDishwasherAlarmClusterNotifyEventClass {
	MTRDishwasherAlarmClusterNotifyEventClassOnce.Do(func() {
		MTRDishwasherAlarmClusterNotifyEventClass = _MTRDishwasherAlarmClusterNotifyEventClass{objc.GetClass("MTRDishwasherAlarmClusterNotifyEvent")}
	})
	return MTRDishwasherAlarmClusterNotifyEventClass
}

type _MTRDishwasherAlarmClusterNotifyEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherAlarmClusterNotifyEvent */
// An interface definition for the [MTRDishwasherAlarmClusterNotifyEvent] class.
type IMTRDishwasherAlarmClusterNotifyEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherAlarmClusterNotifyEvent */
	// properties:
	Active() objc.IObject /* cross-framework: NSNumber */
	SetActive(value objc.IObject /* cross-framework: NSNumber */)
	Inactive() objc.IObject /* cross-framework: NSNumber */
	SetInactive(value objc.IObject /* cross-framework: NSNumber */)
	Mask() objc.IObject /* cross-framework: NSNumber */
	SetMask(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherAlarmClusterNotifyEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherAlarmClusterNotifyEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherAlarmClusterNotifyEventClass) Alloc() MTRDishwasherAlarmClusterNotifyEvent {
	rv := objc.Send[MTRDishwasherAlarmClusterNotifyEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherAlarmClusterNotifyEventClass) New() MTRDishwasherAlarmClusterNotifyEvent {
	rv := objc.Send[MTRDishwasherAlarmClusterNotifyEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Init() MTRDishwasherAlarmClusterNotifyEvent {
	rv := objc.Send[MTRDishwasherAlarmClusterNotifyEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Autorelease() MTRDishwasherAlarmClusterNotifyEvent {
	rv := objc.Send[MTRDishwasherAlarmClusterNotifyEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherAlarmClusterNotifyEvent creates a new MTRDishwasherAlarmClusterNotifyEvent instance.
func NewMTRDishwasherAlarmClusterNotifyEvent() MTRDishwasherAlarmClusterNotifyEvent {
	return getMTRDishwasherAlarmClusterNotifyEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherAlarmClusterNotifyEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent
type MTRDishwasherAlarmClusterNotifyEvent struct {
	objectivec.Object
}

// MTRDishwasherAlarmClusterNotifyEventFrom constructs a [MTRDishwasherAlarmClusterNotifyEvent] from an unsafe.Pointer.
func MTRDishwasherAlarmClusterNotifyEventFrom(ptr unsafe.Pointer) MTRDishwasherAlarmClusterNotifyEvent {
	return MTRDishwasherAlarmClusterNotifyEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherAlarmClusterNotifyEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherAlarmClusterNotifyEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherAlarmClusterNotifyEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherAlarmClusterNotifyEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherAlarmClusterNotifyEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/active
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Active() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/active
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetActive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActive:"), value)
}/* debug [instance_properties/setter]: active */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusternotifyevent/inactive
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Inactive() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("inactive"))
	return rv
}/* debug [instance_properties/getter]: inactive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusternotifyevent/inactive
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetInactive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInactive:"), value)
}/* debug [instance_properties/setter]: inactive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusternotifyevent/mask
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Mask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mask"))
	return rv
}/* debug [instance_properties/getter]: mask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusternotifyevent/mask
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}/* debug [instance_properties/setter]: mask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusternotifyevent/state
func (m_ MTRDishwasherAlarmClusterNotifyEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwasheralarmclusternotifyevent/state
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherAlarmClusterNotifyEvent */



