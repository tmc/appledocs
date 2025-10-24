// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVExtendedNoteOnEvent */


/* debug [class_header]: Header for AVExtendedNoteOnEvent */
// The class instance for the [ExtendedNoteOnEvent] class.
var (
	ExtendedNoteOnEventClass     _ExtendedNoteOnEventClass
	ExtendedNoteOnEventClassOnce sync.Once
)

func getExtendedNoteOnEventClass() _ExtendedNoteOnEventClass {
	ExtendedNoteOnEventClassOnce.Do(func() {
		ExtendedNoteOnEventClass = _ExtendedNoteOnEventClass{objc.GetClass("AVExtendedNoteOnEvent")}
	})
	return ExtendedNoteOnEventClass
}

type _ExtendedNoteOnEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExtendedNoteOnEvent */
// An interface definition for the [ExtendedNoteOnEvent] class.
type IExtendedNoteOnEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for ExtendedNoteOnEvent */
	// properties:
	Duration() MusicTimeStamp /* typedef */
	SetDuration(value MusicTimeStamp /* typedef */)
	GroupID() objectivec.IObject
	SetGroupID(value objectivec.IObject)
	InstrumentID() objectivec.IObject
	SetInstrumentID(value objectivec.IObject)
	MidiNote() float32
	SetMidiNote(value float32)
	Velocity() float32
	SetVelocity(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExtendedNoteOnEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExtendedNoteOnEvent */
// Alloc allocates a new instance without initialization.
func (ec _ExtendedNoteOnEventClass) Alloc() ExtendedNoteOnEvent {
	rv := objc.Send[ExtendedNoteOnEvent](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExtendedNoteOnEventClass) New() ExtendedNoteOnEvent {
	rv := objc.Send[ExtendedNoteOnEvent](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtendedNoteOnEvent) Init() ExtendedNoteOnEvent {
	rv := objc.Send[ExtendedNoteOnEvent](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtendedNoteOnEvent) Autorelease() ExtendedNoteOnEvent {
	rv := objc.Send[ExtendedNoteOnEvent](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtendedNoteOnEvent creates a new ExtendedNoteOnEvent instance.
func NewExtendedNoteOnEvent() ExtendedNoteOnEvent {
	return getExtendedNoteOnEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExtendedNoteOnEvent */
// An object that represents a custom extension of a MIDI note on event.
//
// Use this to allow an app to trigger a custom note on event on one of several Apple audio units that support it. The floating point note and velocity numbers allow for optional fractional control of the note’s runtime properties that the system modulates by those inputs. This event supports the possibility of an audio unit with more than the standard 16 MIDI channels.


// An object that represents a custom extension of a MIDI note on event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent
type ExtendedNoteOnEvent struct {
	MusicEvent
}

// ExtendedNoteOnEventFrom constructs a [ExtendedNoteOnEvent] from an unsafe.Pointer.
//
// An object that represents a custom extension of a MIDI note on event.
func ExtendedNoteOnEventFrom(ptr unsafe.Pointer) ExtendedNoteOnEvent {
	return ExtendedNoteOnEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExtendedNoteOnEvent */

// Creates an event with a MIDI note, velocity, group identifier, and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/init(midiNote:velocity:groupID:duration:)
func NewExtendedNoteOnEventWithMIDINoteVelocityGroupIDDuration(midiNote float32, velocity float32, groupID objectivec.IObject, duration MusicTimeStamp /* typedef */) ExtendedNoteOnEvent {
	instance := getExtendedNoteOnEventClass().Alloc()
	rv := objc.Send[ExtendedNoteOnEvent](instance.ID, objc.Sel("initWithMIDINote:velocity:groupID:duration:"), midiNote, velocity, groupID, duration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewExtendedNoteOnEventWithMIDINoteVelocityGroupIDDuration */


// Creates a note on event with the default instrument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/init(midiNote:velocity:instrumentID:groupID:duration:)
func NewExtendedNoteOnEventWithMIDINoteVelocityInstrumentIDGroupIDDuration(midiNote float32, velocity float32, instrumentID objectivec.IObject, groupID objectivec.IObject, duration MusicTimeStamp /* typedef */) ExtendedNoteOnEvent {
	instance := getExtendedNoteOnEventClass().Alloc()
	rv := objc.Send[ExtendedNoteOnEvent](instance.ID, objc.Sel("initWithMIDINote:velocity:instrumentID:groupID:duration:"), midiNote, velocity, instrumentID, groupID, duration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewExtendedNoteOnEventWithMIDINoteVelocityInstrumentIDGroupIDDuration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExtendedNoteOnEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExtendedNoteOnEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExtendedNoteOnEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExtendedNoteOnEvent */

// The duration of the event, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/duration
func (e_ ExtendedNoteOnEvent) Duration() MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration of the event, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/duration
func (e_ ExtendedNoteOnEvent) SetDuration(value MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The audio unit channel that handles the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/groupID
func (e_ ExtendedNoteOnEvent) GroupID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("groupID"))
	return rv
}/* debug [instance_properties/getter]: groupID */


// The audio unit channel that handles the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/groupID
func (e_ ExtendedNoteOnEvent) SetGroupID(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setGroupID:"), value)
}/* debug [instance_properties/setter]: groupID */


// The instrument identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/instrumentID
func (e_ ExtendedNoteOnEvent) InstrumentID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("instrumentID"))
	return rv
}/* debug [instance_properties/getter]: instrumentID */


// The instrument identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/instrumentID
func (e_ ExtendedNoteOnEvent) SetInstrumentID(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setInstrumentID:"), value)
}/* debug [instance_properties/setter]: instrumentID */


// The MIDI note number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/midiNote
func (e_ ExtendedNoteOnEvent) MidiNote() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("midiNote"))
	return rv
}/* debug [instance_properties/getter]: midiNote */


// The MIDI note number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/midiNote
func (e_ ExtendedNoteOnEvent) SetMidiNote(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMidiNote:"), value)
}/* debug [instance_properties/setter]: midiNote */


// The MDI velocity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/velocity
func (e_ ExtendedNoteOnEvent) Velocity() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("velocity"))
	return rv
}/* debug [instance_properties/getter]: velocity */


// The MDI velocity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/velocity
func (e_ ExtendedNoteOnEvent) SetVelocity(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setVelocity:"), value)
}/* debug [instance_properties/setter]: velocity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVExtendedNoteOnEvent */


