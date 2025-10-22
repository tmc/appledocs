// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WaterSubmersionEvent] class.
var (
	WaterSubmersionEventClass     _WaterSubmersionEventClass
	WaterSubmersionEventClassOnce sync.Once
)

func getWaterSubmersionEventClass() _WaterSubmersionEventClass {
	WaterSubmersionEventClassOnce.Do(func() {
		WaterSubmersionEventClass = _WaterSubmersionEventClass{objc.GetClass("CMWaterSubmersionEvent")}
	})
	return WaterSubmersionEventClass
}

type _WaterSubmersionEventClass struct {
	class objc.Class
}

// An interface definition for the [WaterSubmersionEvent] class.
type IWaterSubmersionEvent interface {
	objectivec.IObject
	Date() foundation.NSDate
	State() WaterSubmersionState
}

// An event indicating that the device’s submersion state has changed.


// An event indicating that the device’s submersion state has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent

type WaterSubmersionEvent struct {
	objectivec.Object
}

// WaterSubmersionEventFrom constructs a [WaterSubmersionEvent] from an unsafe.Pointer.
//
// An event indicating that the device’s submersion state has changed.
func WaterSubmersionEventFrom(ptr unsafe.Pointer) WaterSubmersionEvent {
	return WaterSubmersionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WaterSubmersionEventClass) Alloc() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WaterSubmersionEventClass) New() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterSubmersionEvent) Init() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterSubmersionEvent) Autorelease() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterSubmersionEvent creates a new WaterSubmersionEvent instance.
func NewWaterSubmersionEvent() WaterSubmersionEvent {
	return getWaterSubmersionEventClass().New()
}



// The time and date of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/date

func (w_ WaterSubmersionEvent) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](w_.ID, objc.Sel("date"))
	return rv
}


// The new submersion state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/state-swift.property

func (w_ WaterSubmersionEvent) State() WaterSubmersionState {
	rv := objc.Send[WaterSubmersionState](w_.ID, objc.Sel("state"))
	return rv
}



