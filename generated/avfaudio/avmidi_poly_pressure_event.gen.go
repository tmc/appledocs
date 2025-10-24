// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDIPolyPressureEvent] class.
var (
	MIDIPolyPressureEventClass     _MIDIPolyPressureEventClass
	MIDIPolyPressureEventClassOnce sync.Once
)

func getMIDIPolyPressureEventClass() _MIDIPolyPressureEventClass {
	MIDIPolyPressureEventClassOnce.Do(func() {
		MIDIPolyPressureEventClass = _MIDIPolyPressureEventClass{objc.GetClass("AVMIDIPolyPressureEvent")}
	})
	return MIDIPolyPressureEventClass
}

type _MIDIPolyPressureEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDIPolyPressureEvent] class.
type IMIDIPolyPressureEvent interface {
	IMIDIChannelEvent
	

	// properties:
	Key() objectivec.IObject
	SetKey(value objectivec.IObject)
	Pressure() objectivec.IObject
	SetPressure(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDIPolyPressureEventClass) Alloc() MIDIPolyPressureEvent {
	rv := objc.Send[MIDIPolyPressureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIPolyPressureEventClass) New() MIDIPolyPressureEvent {
	rv := objc.Send[MIDIPolyPressureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIPolyPressureEvent) Init() MIDIPolyPressureEvent {
	rv := objc.Send[MIDIPolyPressureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIPolyPressureEvent) Autorelease() MIDIPolyPressureEvent {
	rv := objc.Send[MIDIPolyPressureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIPolyPressureEvent creates a new MIDIPolyPressureEvent instance.
func NewMIDIPolyPressureEvent() MIDIPolyPressureEvent {
	return getMIDIPolyPressureEventClass().New()
}





// An object that represents a MIDI poly or key pressure event.


// An object that represents a MIDI poly or key pressure event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent
type MIDIPolyPressureEvent struct {
	MIDIChannelEvent
}

// MIDIPolyPressureEventFrom constructs a [MIDIPolyPressureEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI poly or key pressure event.
func MIDIPolyPressureEventFrom(ptr unsafe.Pointer) MIDIPolyPressureEvent {
	return MIDIPolyPressureEvent{
		MIDIChannelEvent: MIDIChannelEventFrom(ptr),
	}
}






// Creates an event with a channel, MIDI key number, and a key pressure value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/init(channel:key:pressure:)
func NewMIDIPolyPressureEventWithChannelKeyPressure(channel objectivec.IObject, key objectivec.IObject, pressure objectivec.IObject) MIDIPolyPressureEvent {
	instance := getMIDIPolyPressureEventClass().Alloc()
	rv := objc.Send[MIDIPolyPressureEvent](instance.ID, objc.Sel("initWithChannel:key:pressure:"), channel, key, pressure)
	rv.Autorelease()
	return rv
}






















// The MIDI key number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/key
func (m_ MIDIPolyPressureEvent) Key() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("key"))
	return rv
}


// The MIDI key number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/key
func (m_ MIDIPolyPressureEvent) SetKey(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}


// The poly pressure value for the requested key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/pressure
func (m_ MIDIPolyPressureEvent) Pressure() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("pressure"))
	return rv
}


// The poly pressure value for the requested key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/pressure
func (m_ MIDIPolyPressureEvent) SetPressure(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPressure:"), value)
}







