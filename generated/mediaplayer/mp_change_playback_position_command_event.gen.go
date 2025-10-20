// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangePlaybackPositionCommandEvent] class.
var (
	ChangePlaybackPositionCommandEventClass     _ChangePlaybackPositionCommandEventClass
	ChangePlaybackPositionCommandEventClassOnce sync.Once
)

func getChangePlaybackPositionCommandEventClass() _ChangePlaybackPositionCommandEventClass {
	ChangePlaybackPositionCommandEventClassOnce.Do(func() {
		ChangePlaybackPositionCommandEventClass = _ChangePlaybackPositionCommandEventClass{objc.GetClass("MPChangePlaybackPositionCommandEvent")}
	})
	return ChangePlaybackPositionCommandEventClass
}

type _ChangePlaybackPositionCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [ChangePlaybackPositionCommandEvent] class.
type IChangePlaybackPositionCommandEvent interface {
	IRemoteCommandEvent
}

// An event requesting a change in the playback position.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackPositionCommandEvent
type ChangePlaybackPositionCommandEvent struct {
	RemoteCommandEvent
}

// ChangePlaybackPositionCommandEventFrom constructs a [ChangePlaybackPositionCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the playback position.
func ChangePlaybackPositionCommandEventFrom(ptr unsafe.Pointer) ChangePlaybackPositionCommandEvent {
	return ChangePlaybackPositionCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackPositionCommandEventClass) Alloc() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangePlaybackPositionCommandEventClass) New() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackPositionCommandEvent) Init() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackPositionCommandEvent) Autorelease() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackPositionCommandEvent creates a new ChangePlaybackPositionCommandEvent instance.
func NewChangePlaybackPositionCommandEvent() ChangePlaybackPositionCommandEvent {
	return getChangePlaybackPositionCommandEventClass().New()
}


// The playback position used when setting the current time of the player.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackPositionCommandEvent/positionTime
func (c_ ChangePlaybackPositionCommandEvent) PositionTime() TimeInterval {
	rv := objc.Send[TimeInterval](c_.ID, objc.Sel("positionTime"))
	return rv
}



