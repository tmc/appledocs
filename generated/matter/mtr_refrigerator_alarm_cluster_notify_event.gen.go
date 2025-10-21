// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MTRRefrigeratorAlarmClusterNotifyEvent] class.
type IMTRRefrigeratorAlarmClusterNotifyEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent
type MTRRefrigeratorAlarmClusterNotifyEvent struct {
	objectivec.Object
}

// MTRRefrigeratorAlarmClusterNotifyEventFrom constructs a [MTRRefrigeratorAlarmClusterNotifyEvent] from an unsafe.Pointer.
func MTRRefrigeratorAlarmClusterNotifyEventFrom(ptr unsafe.Pointer) MTRRefrigeratorAlarmClusterNotifyEvent {
	return MTRRefrigeratorAlarmClusterNotifyEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAlarmClusterNotifyEventClass) Alloc() MTRRefrigeratorAlarmClusterNotifyEvent {
	rv := objc.Send[MTRRefrigeratorAlarmClusterNotifyEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/active
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Active() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("active"))
	return rv
}


// SetActive sets the value of the active property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/active
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetActive(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActive:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/inactive
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Inactive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("inactive"))
	return rv
}


// SetInactive sets the value of the inactive property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/inactive
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetInactive(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInactive:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/mask
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) Mask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mask"))
	return rv
}


// SetMask sets the value of the mask property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/mask
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetMask(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/state
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAlarmClusterNotifyEvent/state
func (m_ MTRRefrigeratorAlarmClusterNotifyEvent) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


