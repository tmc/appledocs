// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SkipIntervalCommandEvent] class.
var (
	SkipIntervalCommandEventClass     _SkipIntervalCommandEventClass
	SkipIntervalCommandEventClassOnce sync.Once
)

func getSkipIntervalCommandEventClass() _SkipIntervalCommandEventClass {
	SkipIntervalCommandEventClassOnce.Do(func() {
		SkipIntervalCommandEventClass = _SkipIntervalCommandEventClass{objc.GetClass("MPSkipIntervalCommandEvent")}
	})
	return SkipIntervalCommandEventClass
}

type _SkipIntervalCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [SkipIntervalCommandEvent] class.
type ISkipIntervalCommandEvent interface {
	IRemoteCommandEvent
}

// An event requesting a change in the current skip interval.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommandEvent
type SkipIntervalCommandEvent struct {
	RemoteCommandEvent
}

// SkipIntervalCommandEventFrom constructs a [SkipIntervalCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the current skip interval.
func SkipIntervalCommandEventFrom(ptr unsafe.Pointer) SkipIntervalCommandEvent {
	return SkipIntervalCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SkipIntervalCommandEventClass) Alloc() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SkipIntervalCommandEventClass) New() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SkipIntervalCommandEvent) Init() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SkipIntervalCommandEvent) Autorelease() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkipIntervalCommandEvent creates a new SkipIntervalCommandEvent instance.
func NewSkipIntervalCommandEvent() SkipIntervalCommandEvent {
	return getSkipIntervalCommandEventClass().New()
}


// The chosen interval, in seconds, for the skip command event.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommandEvent/interval
func (s_ SkipIntervalCommandEvent) Interval() TimeInterval {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("interval"))
	return rv
}



