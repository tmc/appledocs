// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Active() objc.IObject /* cross-framework: NSNumber */
	SetActive(value objc.IObject /* cross-framework: NSNumber */)
	Inactive() objc.IObject /* cross-framework: NSNumber */
	SetInactive(value objc.IObject /* cross-framework: NSNumber */)
	Mask() objc.IObject /* cross-framework: NSNumber */
	SetMask(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/active
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Active() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("active"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/active
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetActive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActive:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/inactive
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Inactive() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("inactive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/inactive
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetInactive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInactive:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/mask
func (m_ MTRDishwasherAlarmClusterNotifyEvent) Mask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/mask
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/state
func (m_ MTRDishwasherAlarmClusterNotifyEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherAlarmClusterNotifyEvent/state
func (m_ MTRDishwasherAlarmClusterNotifyEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}



