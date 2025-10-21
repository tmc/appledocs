// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterTimeZoneStatusEvent] class.
var (
	MTRTimeSynchronizationClusterTimeZoneStatusEventClass     _MTRTimeSynchronizationClusterTimeZoneStatusEventClass
	MTRTimeSynchronizationClusterTimeZoneStatusEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneStatusEventClass() _MTRTimeSynchronizationClusterTimeZoneStatusEventClass {
	MTRTimeSynchronizationClusterTimeZoneStatusEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneStatusEventClass = _MTRTimeSynchronizationClusterTimeZoneStatusEventClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneStatusEvent")}
	})
	return MTRTimeSynchronizationClusterTimeZoneStatusEventClass
}

type _MTRTimeSynchronizationClusterTimeZoneStatusEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneStatusEvent] class.
type IMTRTimeSynchronizationClusterTimeZoneStatusEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent
type MTRTimeSynchronizationClusterTimeZoneStatusEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTimeZoneStatusEventFrom constructs a [MTRTimeSynchronizationClusterTimeZoneStatusEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneStatusEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	return MTRTimeSynchronizationClusterTimeZoneStatusEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneStatusEventClass) Alloc() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterTimeZoneStatusEventClass) New() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Init() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Autorelease() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneStatusEvent creates a new MTRTimeSynchronizationClusterTimeZoneStatusEvent instance.
func NewMTRTimeSynchronizationClusterTimeZoneStatusEvent() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	return getMTRTimeSynchronizationClusterTimeZoneStatusEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Offset() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offset"))
	return rv
}


// SetOffset sets the value of the offset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) SetOffset(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}



