// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PedometerEvent] class.
var (
	PedometerEventClass     _PedometerEventClass
	PedometerEventClassOnce sync.Once
)

func getPedometerEventClass() _PedometerEventClass {
	PedometerEventClassOnce.Do(func() {
		PedometerEventClass = _PedometerEventClass{objc.GetClass("CMPedometerEvent")}
	})
	return PedometerEventClass
}

type _PedometerEventClass struct {
	class objc.Class
}

// An interface definition for the [PedometerEvent] class.
type IPedometerEvent interface {
	objectivec.IObject
	Date() foundation.NSDate
	Type() PedometerEventType
}

// A change in the user’s pedestrian activity.


// A change in the user’s pedestrian activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEvent
type PedometerEvent struct {
	objectivec.Object
}

// PedometerEventFrom constructs a [PedometerEvent] from an unsafe.Pointer.
//
// A change in the user’s pedestrian activity.
func PedometerEventFrom(ptr unsafe.Pointer) PedometerEvent {
	return PedometerEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PedometerEventClass) Alloc() PedometerEvent {
	rv := objc.Send[PedometerEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PedometerEventClass) New() PedometerEvent {
	rv := objc.Send[PedometerEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PedometerEvent) Init() PedometerEvent {
	rv := objc.Send[PedometerEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PedometerEvent) Autorelease() PedometerEvent {
	rv := objc.Send[PedometerEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPedometerEvent creates a new PedometerEvent instance.
func NewPedometerEvent() PedometerEvent {
	return getPedometerEventClass().New()
}



// The date on which the pedometer event was recorded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEvent/date
func (p_ PedometerEvent) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("date"))
	return rv
}


// The type of change that occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEvent/type
func (p_ PedometerEvent) Type() PedometerEventType {
	rv := objc.Send[PedometerEventType](p_.ID, objc.Sel("type"))
	return rv
}



