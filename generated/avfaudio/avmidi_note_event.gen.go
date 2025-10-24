// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDINoteEvent] class.
var (
	MIDINoteEventClass     _MIDINoteEventClass
	MIDINoteEventClassOnce sync.Once
)

func getMIDINoteEventClass() _MIDINoteEventClass {
	MIDINoteEventClassOnce.Do(func() {
		MIDINoteEventClass = _MIDINoteEventClass{objc.GetClass("AVMIDINoteEvent")}
	})
	return MIDINoteEventClass
}

type _MIDINoteEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDINoteEvent] class.
type IMIDINoteEvent interface {
	IMusicEvent
	

	// properties:
	Channel() objectivec.IObject
	SetChannel(value objectivec.IObject)
	Duration() MusicTimeStamp /* typedef */
	SetDuration(value MusicTimeStamp /* typedef */)
	Key() objectivec.IObject
	SetKey(value objectivec.IObject)
	Velocity() objectivec.IObject
	SetVelocity(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDINoteEventClass) Alloc() MIDINoteEvent {
	rv := objc.Send[MIDINoteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDINoteEventClass) New() MIDINoteEvent {
	rv := objc.Send[MIDINoteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINoteEvent) Init() MIDINoteEvent {
	rv := objc.Send[MIDINoteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINoteEvent) Autorelease() MIDINoteEvent {
	rv := objc.Send[MIDINoteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINoteEvent creates a new MIDINoteEvent instance.
func NewMIDINoteEvent() MIDINoteEvent {
	return getMIDINoteEventClass().New()
}





// An object that represents MIDI note on or off messages.


// An object that represents MIDI note on or off messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent
type MIDINoteEvent struct {
	MusicEvent
}

// MIDINoteEventFrom constructs a [MIDINoteEvent] from an unsafe.Pointer.
//
// An object that represents MIDI note on or off messages.
func MIDINoteEventFrom(ptr unsafe.Pointer) MIDINoteEvent {
	return MIDINoteEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}






// Creates an event with a MIDI channel, key number, velocity, and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/init(channel:key:velocity:duration:)
func NewMIDINoteEventWithChannelKeyVelocityDuration(channel objectivec.IObject, keyNum objectivec.IObject, velocity objectivec.IObject, duration MusicTimeStamp /* typedef */) MIDINoteEvent {
	instance := getMIDINoteEventClass().Alloc()
	rv := objc.Send[MIDINoteEvent](instance.ID, objc.Sel("initWithChannel:key:velocity:duration:"), channel, keyNum, velocity, duration)
	rv.Autorelease()
	return rv
}






















// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/channel
func (m_ MIDINoteEvent) Channel() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("channel"))
	return rv
}


// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/channel
func (m_ MIDINoteEvent) SetChannel(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}


// The duration for the note, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/duration
func (m_ MIDINoteEvent) Duration() MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("duration"))
	return rv
}


// The duration for the note, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/duration
func (m_ MIDINoteEvent) SetDuration(value MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// The MIDI key number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/key
func (m_ MIDINoteEvent) Key() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("key"))
	return rv
}


// The MIDI key number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/key
func (m_ MIDINoteEvent) SetKey(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}


// The MIDI velocity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/velocity
func (m_ MIDINoteEvent) Velocity() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("velocity"))
	return rv
}


// The MIDI velocity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/velocity
func (m_ MIDINoteEvent) SetVelocity(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVelocity:"), value)
}







