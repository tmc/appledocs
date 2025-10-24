// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDIChannelEvent] class.
var (
	MIDIChannelEventClass     _MIDIChannelEventClass
	MIDIChannelEventClassOnce sync.Once
)

func getMIDIChannelEventClass() _MIDIChannelEventClass {
	MIDIChannelEventClassOnce.Do(func() {
		MIDIChannelEventClass = _MIDIChannelEventClass{objc.GetClass("AVMIDIChannelEvent")}
	})
	return MIDIChannelEventClass
}

type _MIDIChannelEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDIChannelEvent] class.
type IMIDIChannelEvent interface {
	IMusicEvent
	

	// properties:
	Channel() objectivec.IObject
	SetChannel(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDIChannelEventClass) Alloc() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIChannelEventClass) New() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIChannelEvent) Init() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIChannelEvent) Autorelease() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIChannelEvent creates a new MIDIChannelEvent instance.
func NewMIDIChannelEvent() MIDIChannelEvent {
	return getMIDIChannelEventClass().New()
}





// A base class for all MIDI messages that operate on a single MIDI channel.


// A base class for all MIDI messages that operate on a single MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent
type MIDIChannelEvent struct {
	MusicEvent
}

// MIDIChannelEventFrom constructs a [MIDIChannelEvent] from an unsafe.Pointer.
//
// A base class for all MIDI messages that operate on a single MIDI channel.
func MIDIChannelEventFrom(ptr unsafe.Pointer) MIDIChannelEvent {
	return MIDIChannelEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}

























// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/channel
func (m_ MIDIChannelEvent) Channel() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("channel"))
	return rv
}


// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/channel
func (m_ MIDIChannelEvent) SetChannel(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}








