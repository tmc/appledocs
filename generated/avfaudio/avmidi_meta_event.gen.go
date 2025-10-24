// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDIMetaEvent] class.
var (
	MIDIMetaEventClass     _MIDIMetaEventClass
	MIDIMetaEventClassOnce sync.Once
)

func getMIDIMetaEventClass() _MIDIMetaEventClass {
	MIDIMetaEventClassOnce.Do(func() {
		MIDIMetaEventClass = _MIDIMetaEventClass{objc.GetClass("AVMIDIMetaEvent")}
	})
	return MIDIMetaEventClass
}

type _MIDIMetaEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDIMetaEvent] class.
type IMIDIMetaEvent interface {
	IMusicEvent
	

	// properties:
	Type() MIDIMetaEventType


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDIMetaEventClass) Alloc() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIMetaEventClass) New() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIMetaEvent) Init() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIMetaEvent) Autorelease() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIMetaEvent creates a new MIDIMetaEvent instance.
func NewMIDIMetaEvent() MIDIMetaEvent {
	return getMIDIMetaEventClass().New()
}





// An object that represents MIDI meta event messages.
//
// You can’t modify the size and contents of this event once you create it. This doesn’t verify that the content matches the MIDI specification. You can only add , , or to a sequence’s tempo track.


// An object that represents MIDI meta event messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent
type MIDIMetaEvent struct {
	MusicEvent
}

// MIDIMetaEventFrom constructs a [MIDIMetaEvent] from an unsafe.Pointer.
//
// An object that represents MIDI meta event messages.
func MIDIMetaEventFrom(ptr unsafe.Pointer) MIDIMetaEvent {
	return MIDIMetaEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}






// Creates an event with a MIDI meta event type and data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/init(type:data:)
func NewMIDIMetaEventWithTypeData(type_ MIDIMetaEventType, data objc.IObject /* cross-framework: NSData */) MIDIMetaEvent {
	instance := getMIDIMetaEventClass().Alloc()
	rv := objc.Send[MIDIMetaEvent](instance.ID, objc.Sel("initWithType:data:"), type_, data)
	rv.Autorelease()
	return rv
}






















// The type of meta event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/type
func (m_ MIDIMetaEvent) Type() MIDIMetaEventType {
	rv := objc.Send[MIDIMetaEventType](m_.ID, objc.Sel("type"))
	return rv
}







