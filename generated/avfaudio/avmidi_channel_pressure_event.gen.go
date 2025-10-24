// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDIChannelPressureEvent] class.
var (
	MIDIChannelPressureEventClass     _MIDIChannelPressureEventClass
	MIDIChannelPressureEventClassOnce sync.Once
)

func getMIDIChannelPressureEventClass() _MIDIChannelPressureEventClass {
	MIDIChannelPressureEventClassOnce.Do(func() {
		MIDIChannelPressureEventClass = _MIDIChannelPressureEventClass{objc.GetClass("AVMIDIChannelPressureEvent")}
	})
	return MIDIChannelPressureEventClass
}

type _MIDIChannelPressureEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDIChannelPressureEvent] class.
type IMIDIChannelPressureEvent interface {
	IMIDIChannelEvent
	

	// properties:
	Pressure() objectivec.IObject
	SetPressure(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDIChannelPressureEventClass) Alloc() MIDIChannelPressureEvent {
	rv := objc.Send[MIDIChannelPressureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIChannelPressureEventClass) New() MIDIChannelPressureEvent {
	rv := objc.Send[MIDIChannelPressureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIChannelPressureEvent) Init() MIDIChannelPressureEvent {
	rv := objc.Send[MIDIChannelPressureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIChannelPressureEvent) Autorelease() MIDIChannelPressureEvent {
	rv := objc.Send[MIDIChannelPressureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIChannelPressureEvent creates a new MIDIChannelPressureEvent instance.
func NewMIDIChannelPressureEvent() MIDIChannelPressureEvent {
	return getMIDIChannelPressureEventClass().New()
}





// An object that represents a MIDI channel pressure message.
//
// The effect of this message depends on the destination audio unit, and the capabilities of the destination’s loaded instrument.


// An object that represents a MIDI channel pressure message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent
type MIDIChannelPressureEvent struct {
	MIDIChannelEvent
}

// MIDIChannelPressureEventFrom constructs a [MIDIChannelPressureEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI channel pressure message.
func MIDIChannelPressureEventFrom(ptr unsafe.Pointer) MIDIChannelPressureEvent {
	return MIDIChannelPressureEvent{
		MIDIChannelEvent: MIDIChannelEventFrom(ptr),
	}
}






// Creates a pressure event with a channel and pressure value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent/init(channel:pressure:)
func NewMIDIChannelPressureEventWithChannelPressure(channel objectivec.IObject, pressure objectivec.IObject) MIDIChannelPressureEvent {
	instance := getMIDIChannelPressureEventClass().Alloc()
	rv := objc.Send[MIDIChannelPressureEvent](instance.ID, objc.Sel("initWithChannel:pressure:"), channel, pressure)
	rv.Autorelease()
	return rv
}






















// The MIDI channel pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent/pressure
func (m_ MIDIChannelPressureEvent) Pressure() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("pressure"))
	return rv
}


// The MIDI channel pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent/pressure
func (m_ MIDIChannelPressureEvent) SetPressure(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPressure:"), value)
}







