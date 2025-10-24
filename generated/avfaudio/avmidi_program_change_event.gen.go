// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDIProgramChangeEvent] class.
var (
	MIDIProgramChangeEventClass     _MIDIProgramChangeEventClass
	MIDIProgramChangeEventClassOnce sync.Once
)

func getMIDIProgramChangeEventClass() _MIDIProgramChangeEventClass {
	MIDIProgramChangeEventClassOnce.Do(func() {
		MIDIProgramChangeEventClass = _MIDIProgramChangeEventClass{objc.GetClass("AVMIDIProgramChangeEvent")}
	})
	return MIDIProgramChangeEventClass
}

type _MIDIProgramChangeEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDIProgramChangeEvent] class.
type IMIDIProgramChangeEvent interface {
	IMIDIChannelEvent
	

	// properties:
	ProgramNumber() objectivec.IObject
	SetProgramNumber(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDIProgramChangeEventClass) Alloc() MIDIProgramChangeEvent {
	rv := objc.Send[MIDIProgramChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIProgramChangeEventClass) New() MIDIProgramChangeEvent {
	rv := objc.Send[MIDIProgramChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIProgramChangeEvent) Init() MIDIProgramChangeEvent {
	rv := objc.Send[MIDIProgramChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIProgramChangeEvent) Autorelease() MIDIProgramChangeEvent {
	rv := objc.Send[MIDIProgramChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIProgramChangeEvent creates a new MIDIProgramChangeEvent instance.
func NewMIDIProgramChangeEvent() MIDIProgramChangeEvent {
	return getMIDIProgramChangeEventClass().New()
}





// An object that represents a MIDI program or patch change message.
//
// The effect of this message depends on the destination audio unit.


// An object that represents a MIDI program or patch change message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent
type MIDIProgramChangeEvent struct {
	MIDIChannelEvent
}

// MIDIProgramChangeEventFrom constructs a [MIDIProgramChangeEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI program or patch change message.
func MIDIProgramChangeEventFrom(ptr unsafe.Pointer) MIDIProgramChangeEvent {
	return MIDIProgramChangeEvent{
		MIDIChannelEvent: MIDIChannelEventFrom(ptr),
	}
}






// Creates a program change event with a channel and program number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/init(channel:programNumber:)
func NewMIDIProgramChangeEventWithChannelProgramNumber(channel objectivec.IObject, programNumber objectivec.IObject) MIDIProgramChangeEvent {
	instance := getMIDIProgramChangeEventClass().Alloc()
	rv := objc.Send[MIDIProgramChangeEvent](instance.ID, objc.Sel("initWithChannel:programNumber:"), channel, programNumber)
	rv.Autorelease()
	return rv
}






















// The MIDI program number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/programNumber
func (m_ MIDIProgramChangeEvent) ProgramNumber() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("programNumber"))
	return rv
}


// The MIDI program number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/programNumber
func (m_ MIDIProgramChangeEvent) SetProgramNumber(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramNumber:"), value)
}







