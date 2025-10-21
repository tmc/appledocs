// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SeekCommandEvent] class.
var (
	SeekCommandEventClass     _SeekCommandEventClass
	SeekCommandEventClassOnce sync.Once
)

func getSeekCommandEventClass() _SeekCommandEventClass {
	SeekCommandEventClassOnce.Do(func() {
		SeekCommandEventClass = _SeekCommandEventClass{objc.GetClass("MPSeekCommandEvent")}
	})
	return SeekCommandEventClass
}

type _SeekCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [SeekCommandEvent] class.
type ISeekCommandEvent interface {
	IRemoteCommandEvent
}

// An event requesting that the player seek to a new position.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSeekCommandEvent
type SeekCommandEvent struct {
	RemoteCommandEvent
}

// SeekCommandEventFrom constructs a [SeekCommandEvent] from an unsafe.Pointer.
//
// An event requesting that the player seek to a new position.
func SeekCommandEventFrom(ptr unsafe.Pointer) SeekCommandEvent {
	return SeekCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SeekCommandEventClass) Alloc() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SeekCommandEventClass) New() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SeekCommandEvent) Init() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SeekCommandEvent) Autorelease() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSeekCommandEvent creates a new SeekCommandEvent instance.
func NewSeekCommandEvent() SeekCommandEvent {
	return getSeekCommandEventClass().New()
}


// The type of seek command event.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpseekcommandevent/type
func (s_ SeekCommandEvent) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of seek command event.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpseekcommandevent/type
func (s_ SeekCommandEvent) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setType:"), value)
}



