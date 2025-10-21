// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDishwasherAlarmClusterNotifyEvent] class.
type IMTRDishwasherAlarmClusterNotifyEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent
type MTRDishwasherAlarmClusterNotifyEvent struct {
	objectivec.Object
}

// MTRDishwasherAlarmClusterNotifyEventFrom constructs a [MTRDishwasherAlarmClusterNotifyEvent] from an unsafe.Pointer.
func MTRDishwasherAlarmClusterNotifyEventFrom(ptr unsafe.Pointer) MTRDishwasherAlarmClusterNotifyEvent {
	return MTRDishwasherAlarmClusterNotifyEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherAlarmClusterNotifyEventClass) Alloc() MTRDishwasherAlarmClusterNotifyEvent {
	rv := objc.Send[MTRDishwasherAlarmClusterNotifyEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/active
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Active() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("active"))
	return rv
}


// SetActive sets the value of the active property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/active
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetActive(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActive:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/inactive
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Inactive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("inactive"))
	return rv
}


// SetInactive sets the value of the inactive property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/inactive
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetInactive(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInactive:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/mask
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Mask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mask"))
	return rv
}


// SetMask sets the value of the mask property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/mask
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetMask(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/state
func (m_ MTRDishwasherAlarmClusterNotifyEvent) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/state
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


