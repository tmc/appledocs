// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangeShuffleModeCommandEvent] class.
var (
	ChangeShuffleModeCommandEventClass     _ChangeShuffleModeCommandEventClass
	ChangeShuffleModeCommandEventClassOnce sync.Once
)

func getChangeShuffleModeCommandEventClass() _ChangeShuffleModeCommandEventClass {
	ChangeShuffleModeCommandEventClassOnce.Do(func() {
		ChangeShuffleModeCommandEventClass = _ChangeShuffleModeCommandEventClass{objc.GetClass("MPChangeShuffleModeCommandEvent")}
	})
	return ChangeShuffleModeCommandEventClass
}

type _ChangeShuffleModeCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [ChangeShuffleModeCommandEvent] class.
type IChangeShuffleModeCommandEvent interface {
	IRemoteCommandEvent
}

// An event requesting a change in the shuffle mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommandEvent
type ChangeShuffleModeCommandEvent struct {
	RemoteCommandEvent
}

// ChangeShuffleModeCommandEventFrom constructs a [ChangeShuffleModeCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the shuffle mode.
func ChangeShuffleModeCommandEventFrom(ptr unsafe.Pointer) ChangeShuffleModeCommandEvent {
	return ChangeShuffleModeCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangeShuffleModeCommandEventClass) Alloc() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangeShuffleModeCommandEventClass) New() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeShuffleModeCommandEvent) Init() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeShuffleModeCommandEvent) Autorelease() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeShuffleModeCommandEvent creates a new ChangeShuffleModeCommandEvent instance.
func NewChangeShuffleModeCommandEvent() ChangeShuffleModeCommandEvent {
	return getChangeShuffleModeCommandEventClass().New()
}


// The shuffle type used when fulfilling the event request.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent/shuffletype
func (c_ ChangeShuffleModeCommandEvent) ShuffleType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("shuffleType"))
	return rv
}


// SetShuffleType sets the value of the shuffleType property.
// The shuffle type used when fulfilling the event request.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent/shuffletype
func (c_ ChangeShuffleModeCommandEvent) SetShuffleType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShuffleType:"), value)
}

// A Boolean value that indicates whether the shuffle mode is preserved between playback sessions.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent/preservesshufflemode
func (c_ ChangeShuffleModeCommandEvent) PreservesShuffleMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesShuffleMode"))
	return rv
}


// SetPreservesShuffleMode sets the value of the preservesShuffleMode property.
// A Boolean value that indicates whether the shuffle mode is preserved between playback sessions.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangeshufflemodecommandevent/preservesshufflemode
func (c_ ChangeShuffleModeCommandEvent) SetPreservesShuffleMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesShuffleMode:"), value)
}



