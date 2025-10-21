// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangeRepeatModeCommandEvent] class.
var (
	ChangeRepeatModeCommandEventClass     _ChangeRepeatModeCommandEventClass
	ChangeRepeatModeCommandEventClassOnce sync.Once
)

func getChangeRepeatModeCommandEventClass() _ChangeRepeatModeCommandEventClass {
	ChangeRepeatModeCommandEventClassOnce.Do(func() {
		ChangeRepeatModeCommandEventClass = _ChangeRepeatModeCommandEventClass{objc.GetClass("MPChangeRepeatModeCommandEvent")}
	})
	return ChangeRepeatModeCommandEventClass
}

type _ChangeRepeatModeCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [ChangeRepeatModeCommandEvent] class.
type IChangeRepeatModeCommandEvent interface {
	IRemoteCommandEvent
}

// An event requesting a change in the repeat mode.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommandEvent
type ChangeRepeatModeCommandEvent struct {
	RemoteCommandEvent
}

// ChangeRepeatModeCommandEventFrom constructs a [ChangeRepeatModeCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the repeat mode.
func ChangeRepeatModeCommandEventFrom(ptr unsafe.Pointer) ChangeRepeatModeCommandEvent {
	return ChangeRepeatModeCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangeRepeatModeCommandEventClass) Alloc() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangeRepeatModeCommandEventClass) New() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeRepeatModeCommandEvent) Init() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeRepeatModeCommandEvent) Autorelease() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeRepeatModeCommandEvent creates a new ChangeRepeatModeCommandEvent instance.
func NewChangeRepeatModeCommandEvent() ChangeRepeatModeCommandEvent {
	return getChangeRepeatModeCommandEventClass().New()
}


// A Boolean value that indicates whether the chosen repeat mode is preserved between playback sessions.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent/preservesrepeatmode
func (c_ ChangeRepeatModeCommandEvent) PreservesRepeatMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesRepeatMode"))
	return rv
}


// SetPreservesRepeatMode sets the value of the preservesRepeatMode property.
// A Boolean value that indicates whether the chosen repeat mode is preserved between playback sessions.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent/preservesrepeatmode
func (c_ ChangeRepeatModeCommandEvent) SetPreservesRepeatMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesRepeatMode:"), value)
}

// The repeat type used when fulfilling the event request.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent/repeattype
func (c_ ChangeRepeatModeCommandEvent) RepeatType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("repeatType"))
	return rv
}


// SetRepeatType sets the value of the repeatType property.
// The repeat type used when fulfilling the event request.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangerepeatmodecommandevent/repeattype
func (c_ ChangeRepeatModeCommandEvent) SetRepeatType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRepeatType:"), value)
}



