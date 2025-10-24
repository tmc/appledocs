// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDIControlChangeEvent] class.
var (
	MIDIControlChangeEventClass     _MIDIControlChangeEventClass
	MIDIControlChangeEventClassOnce sync.Once
)

func getMIDIControlChangeEventClass() _MIDIControlChangeEventClass {
	MIDIControlChangeEventClassOnce.Do(func() {
		MIDIControlChangeEventClass = _MIDIControlChangeEventClass{objc.GetClass("AVMIDIControlChangeEvent")}
	})
	return MIDIControlChangeEventClass
}

type _MIDIControlChangeEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDIControlChangeEvent] class.
type IMIDIControlChangeEvent interface {
	IMIDIChannelEvent
	

	// properties:
	MessageType() MIDIControlChangeMessageType
	Value() objectivec.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDIControlChangeEventClass) Alloc() MIDIControlChangeEvent {
	rv := objc.Send[MIDIControlChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIControlChangeEventClass) New() MIDIControlChangeEvent {
	rv := objc.Send[MIDIControlChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIControlChangeEvent) Init() MIDIControlChangeEvent {
	rv := objc.Send[MIDIControlChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIControlChangeEvent) Autorelease() MIDIControlChangeEvent {
	rv := objc.Send[MIDIControlChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIControlChangeEvent creates a new MIDIControlChangeEvent instance.
func NewMIDIControlChangeEvent() MIDIControlChangeEvent {
	return getMIDIControlChangeEventClass().New()
}





// An object that represents a MIDI control change message.


// An object that represents a MIDI control change message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent
type MIDIControlChangeEvent struct {
	MIDIChannelEvent
}

// MIDIControlChangeEventFrom constructs a [MIDIControlChangeEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI control change message.
func MIDIControlChangeEventFrom(ptr unsafe.Pointer) MIDIControlChangeEvent {
	return MIDIControlChangeEvent{
		MIDIChannelEvent: MIDIChannelEventFrom(ptr),
	}
}






// Creates an event with a channel, control change type, and a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/init(channel:messageType:value:)
func NewMIDIControlChangeEventWithChannelMessageTypeValue(channel objectivec.IObject, messageType MIDIControlChangeMessageType, value objectivec.IObject) MIDIControlChangeEvent {
	instance := getMIDIControlChangeEventClass().Alloc()
	rv := objc.Send[MIDIControlChangeEvent](instance.ID, objc.Sel("initWithChannel:messageType:value:"), channel, messageType, value)
	rv.Autorelease()
	return rv
}






















// The type of control change message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/messageType-swift.property
func (m_ MIDIControlChangeEvent) MessageType() MIDIControlChangeMessageType {
	rv := objc.Send[MIDIControlChangeMessageType](m_.ID, objc.Sel("messageType"))
	return rv
}


// The value of the control change event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/value
func (m_ MIDIControlChangeEvent) Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("value"))
	return rv
}







