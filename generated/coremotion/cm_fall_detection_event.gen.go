// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FallDetectionEvent] class.
var (
	FallDetectionEventClass     _FallDetectionEventClass
	FallDetectionEventClassOnce sync.Once
)

func getFallDetectionEventClass() _FallDetectionEventClass {
	FallDetectionEventClassOnce.Do(func() {
		FallDetectionEventClass = _FallDetectionEventClass{objc.GetClass("CMFallDetectionEvent")}
	})
	return FallDetectionEventClass
}

type _FallDetectionEventClass struct {
	class objc.Class
}

// An interface definition for the [FallDetectionEvent] class.
type IFallDetectionEvent interface {
	objectivec.IObject
	Date() foundation.NSDate
	Resolution() FallDetectionEventUserResolution
}

// An object that contains data about a fall detection event.


// An object that contains data about a fall detection event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent
type FallDetectionEvent struct {
	objectivec.Object
}

// FallDetectionEventFrom constructs a [FallDetectionEvent] from an unsafe.Pointer.
//
// An object that contains data about a fall detection event.
func FallDetectionEventFrom(ptr unsafe.Pointer) FallDetectionEvent {
	return FallDetectionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FallDetectionEventClass) Alloc() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FallDetectionEventClass) New() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FallDetectionEvent) Init() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FallDetectionEvent) Autorelease() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFallDetectionEvent creates a new FallDetectionEvent instance.
func NewFallDetectionEvent() FallDetectionEvent {
	return getFallDetectionEventClass().New()
}



// The event’s time and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/date
func (f_ FallDetectionEvent) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](f_.ID, objc.Sel("date"))
	return rv
}


// The event’s resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/resolution
func (f_ FallDetectionEvent) Resolution() FallDetectionEventUserResolution {
	rv := objc.Send[FallDetectionEventUserResolution](f_.ID, objc.Sel("resolution"))
	return rv
}



