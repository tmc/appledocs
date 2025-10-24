// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioUnitMIDIInstrument] class.
var (
	AudioUnitMIDIInstrumentClass     _AudioUnitMIDIInstrumentClass
	AudioUnitMIDIInstrumentClassOnce sync.Once
)

func getAudioUnitMIDIInstrumentClass() _AudioUnitMIDIInstrumentClass {
	AudioUnitMIDIInstrumentClassOnce.Do(func() {
		AudioUnitMIDIInstrumentClass = _AudioUnitMIDIInstrumentClass{objc.GetClass("AVAudioUnitMIDIInstrument")}
	})
	return AudioUnitMIDIInstrumentClass
}

type _AudioUnitMIDIInstrumentClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnitMIDIInstrument] class.
type IAudioUnitMIDIInstrument interface {
	IAudioUnit
	

	// properties:


	

	// methods:
	SendMIDIEventList(eventList objc.IObject /* cross-framework: MIDIEventList */)
	SendControllerWithValueOnChannel(controller uint8 /* not a class type */, value uint8 /* not a class type */, channel uint8 /* not a class type */)
	SendMIDIEventData1(midiStatus uint8 /* not a class type */, data1 uint8 /* not a class type */)
	SendMIDIEventData1Data2(midiStatus uint8 /* not a class type */, data1 uint8 /* not a class type */, data2 uint8 /* not a class type */)
	SendMIDISysExEvent(midiData objc.IObject /* cross-framework: NSData */)
	SendPitchBendOnChannel(pitchbend uint16 /* not a class type */, channel uint8 /* not a class type */)
	SendPressureOnChannel(pressure uint8 /* not a class type */, channel uint8 /* not a class type */)
	SendPressureForKeyWithValueOnChannel(key uint8 /* not a class type */, value uint8 /* not a class type */, channel uint8 /* not a class type */)
	SendProgramChangeBankMSBBankLSBOnChannel(program uint8 /* not a class type */, bankMSB uint8 /* not a class type */, bankLSB uint8 /* not a class type */, channel uint8 /* not a class type */)
	SendProgramChangeOnChannel(program uint8 /* not a class type */, channel uint8 /* not a class type */)
	StartNoteWithVelocityOnChannel(note uint8 /* not a class type */, velocity uint8 /* not a class type */, channel uint8 /* not a class type */)
	StopNoteOnChannel(note uint8 /* not a class type */, channel uint8 /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitMIDIInstrumentClass) Alloc() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitMIDIInstrumentClass) New() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitMIDIInstrument) Init() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitMIDIInstrument) Autorelease() AudioUnitMIDIInstrument {
	rv := objc.Send[AudioUnitMIDIInstrument](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitMIDIInstrument creates a new AudioUnitMIDIInstrument instance.
func NewAudioUnitMIDIInstrument() AudioUnitMIDIInstrument {
	return getAudioUnitMIDIInstrumentClass().New()
}





// An object that represents music devices or remote instruments.
//
// Use an in a chain that processes real-time (live) input and has the general concept of music events; for example, notes.


// An object that represents music devices or remote instruments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument
type AudioUnitMIDIInstrument struct {
	AudioUnit
}

// AudioUnitMIDIInstrumentFrom constructs a [AudioUnitMIDIInstrument] from an unsafe.Pointer.
//
// An object that represents music devices or remote instruments.
func AudioUnitMIDIInstrumentFrom(ptr unsafe.Pointer) AudioUnitMIDIInstrument {
	return AudioUnitMIDIInstrument{
		AudioUnit: AudioUnitFrom(ptr),
	}
}






// Creates a MIDI instrument audio unit with the component description you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/init(audioComponentDescription:)
func NewAudioUnitMIDIInstrumentWithAudioComponentDescription(description audiotoolbox.AudioComponentDescription) AudioUnitMIDIInstrument {
	instance := getAudioUnitMIDIInstrumentClass().Alloc()
	rv := objc.Send[AudioUnitMIDIInstrument](instance.ID, objc.Sel("initWithAudioComponentDescription:"), description)
	rv.Autorelease()
	return rv
}

















// Sends a MIDI event list to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/send(_:)
func (a_ AudioUnitMIDIInstrument) SendMIDIEventList(eventList objc.IObject /* cross-framework: MIDIEventList */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendMIDIEventList:"), eventList)
}


// Sends a MIDI controller event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendController(_:withValue:onChannel:)
func (a_ AudioUnitMIDIInstrument) SendControllerWithValueOnChannel(controller uint8 /* not a class type */, value uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendController:withValue:onChannel:"), controller, value, channel)
}


// Sends a MIDI event which contains one data byte to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendMIDIEvent(_:data1:)
func (a_ AudioUnitMIDIInstrument) SendMIDIEventData1(midiStatus uint8 /* not a class type */, data1 uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendMIDIEvent:data1:"), midiStatus, data1)
}


// Sends a MIDI event which contains two data bytes to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendMIDIEvent(_:data1:data2:)
func (a_ AudioUnitMIDIInstrument) SendMIDIEventData1Data2(midiStatus uint8 /* not a class type */, data1 uint8 /* not a class type */, data2 uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendMIDIEvent:data1:data2:"), midiStatus, data1, data2)
}


// Sends a MIDI System Exclusive event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendMIDISysExEvent(_:)
func (a_ AudioUnitMIDIInstrument) SendMIDISysExEvent(midiData objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendMIDISysExEvent:"), midiData)
}


// Sends a MIDI Pitch Bend event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendPitchBend(_:onChannel:)
func (a_ AudioUnitMIDIInstrument) SendPitchBendOnChannel(pitchbend uint16 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendPitchBend:onChannel:"), pitchbend, channel)
}


// Sends a MIDI channel pressure event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendPressure(_:onChannel:)
func (a_ AudioUnitMIDIInstrument) SendPressureOnChannel(pressure uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendPressure:onChannel:"), pressure, channel)
}


// Sends a MIDI Polyphonic key pressure event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendPressure(forKey:withValue:onChannel:)
func (a_ AudioUnitMIDIInstrument) SendPressureForKeyWithValueOnChannel(key uint8 /* not a class type */, value uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendPressureForKey:withValue:onChannel:"), key, value, channel)
}


// Sends MIDI Program Change and Bank Select events to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendProgramChange(_:bankMSB:bankLSB:onChannel:)
func (a_ AudioUnitMIDIInstrument) SendProgramChangeBankMSBBankLSBOnChannel(program uint8 /* not a class type */, bankMSB uint8 /* not a class type */, bankLSB uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendProgramChange:bankMSB:bankLSB:onChannel:"), program, bankMSB, bankLSB, channel)
}


// Sends MIDI Program Change and Bank Select events to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/sendProgramChange(_:onChannel:)
func (a_ AudioUnitMIDIInstrument) SendProgramChangeOnChannel(program uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("sendProgramChange:onChannel:"), program, channel)
}


// Sends a MIDI Note On event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/startNote(_:withVelocity:onChannel:)
func (a_ AudioUnitMIDIInstrument) StartNoteWithVelocityOnChannel(note uint8 /* not a class type */, velocity uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startNote:withVelocity:onChannel:"), note, velocity, channel)
}


// Sends a MIDI Note Off event to the instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/stopNote(_:onChannel:)
func (a_ AudioUnitMIDIInstrument) StopNoteOnChannel(note uint8 /* not a class type */, channel uint8 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopNote:onChannel:"), note, channel)
}












