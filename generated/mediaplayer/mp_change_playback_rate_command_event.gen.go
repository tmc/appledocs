// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangePlaybackRateCommandEvent] class.
var (
	ChangePlaybackRateCommandEventClass     _ChangePlaybackRateCommandEventClass
	ChangePlaybackRateCommandEventClassOnce sync.Once
)

func getChangePlaybackRateCommandEventClass() _ChangePlaybackRateCommandEventClass {
	ChangePlaybackRateCommandEventClassOnce.Do(func() {
		ChangePlaybackRateCommandEventClass = _ChangePlaybackRateCommandEventClass{objc.GetClass("MPChangePlaybackRateCommandEvent")}
	})
	return ChangePlaybackRateCommandEventClass
}

type _ChangePlaybackRateCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [ChangePlaybackRateCommandEvent] class.
type IChangePlaybackRateCommandEvent interface {
	IRemoteCommandEvent
	PlaybackRate() float32
	SetPlaybackRate(value float32)
}

// An event requesting a change in the playback rate.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommandEvent
type ChangePlaybackRateCommandEvent struct {
	RemoteCommandEvent
}

// ChangePlaybackRateCommandEventFrom constructs a [ChangePlaybackRateCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the playback rate.
func ChangePlaybackRateCommandEventFrom(ptr unsafe.Pointer) ChangePlaybackRateCommandEvent {
	return ChangePlaybackRateCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackRateCommandEventClass) Alloc() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangePlaybackRateCommandEventClass) New() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackRateCommandEvent) Init() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackRateCommandEvent) Autorelease() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackRateCommandEvent creates a new ChangePlaybackRateCommandEvent instance.
func NewChangePlaybackRateCommandEvent() ChangePlaybackRateCommandEvent {
	return getChangePlaybackRateCommandEventClass().New()
}


// The chosen playback rate for the command event.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommandevent/playbackrate
func (c_ ChangePlaybackRateCommandEvent) PlaybackRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("playbackRate"))
	return rv
}


// SetPlaybackRate sets the value of the playbackRate property.
// The chosen playback rate for the command event.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeplaybackratecommandevent/playbackrate
func (c_ ChangePlaybackRateCommandEvent) SetPlaybackRate(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlaybackRate:"), value)
}



