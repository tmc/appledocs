// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CaptureEvent] class.
var (
	CaptureEventClass     _CaptureEventClass
	CaptureEventClassOnce sync.Once
)

func getCaptureEventClass() _CaptureEventClass {
	CaptureEventClassOnce.Do(func() {
		CaptureEventClass = _CaptureEventClass{objc.GetClass("AVCaptureEvent")}
	})
	return CaptureEventClass
}

type _CaptureEventClass struct {
	class objc.Class
}

// An interface definition for the [CaptureEvent] class.
type ICaptureEvent interface {
	objectivec.IObject
	PlaySound(sound unsafe.Pointer) bool
}

// An object that describes a user interaction with a system hardware button.
//
// Inspect a capture event’s to determine whether the event begins, ends, or is in a canceled state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent
type CaptureEvent struct {
	objectivec.Object
}

// CaptureEventFrom constructs a [CaptureEvent] from an unsafe.Pointer.
//
// An object that describes a user interaction with a system hardware button.
func CaptureEventFrom(ptr unsafe.Pointer) CaptureEvent {
	return CaptureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureEventClass) Alloc() CaptureEvent {
	rv := objc.Send[CaptureEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureEventClass) New() CaptureEvent {
	rv := objc.Send[CaptureEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureEvent) Init() CaptureEvent {
	rv := objc.Send[CaptureEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureEvent) Autorelease() CaptureEvent {
	rv := objc.Send[CaptureEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureEvent creates a new CaptureEvent instance.
func NewCaptureEvent() CaptureEvent {
	return getCaptureEventClass().New()
}


// Plays the specified capture sound through AirPods.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent/play(_:)
func (c_ CaptureEvent) PlaySound(sound unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("playSound:"), sound)
	return rv
}

// The current phase of a capture event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent/phase
func (c_ CaptureEvent) Phase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("phase"))
	return rv
}

// A Boolean value that indicates whether you must play a sound manually.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEvent/shouldPlaySound
func (c_ CaptureEvent) ShouldPlaySound() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldPlaySound"))
	return rv
}



