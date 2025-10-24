// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMIDIProgramChangeEvent */


/* debug [class_header]: Header for AVMIDIProgramChangeEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIProgramChangeEvent */
// An interface definition for the [MIDIProgramChangeEvent] class.
type IMIDIProgramChangeEvent interface {
	IMIDIChannelEvent
	
/* debug [class_interface_properties]: Properties for MIDIProgramChangeEvent */
	// properties:
	ProgramNumber() objectivec.IObject
	SetProgramNumber(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIProgramChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIProgramChangeEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIProgramChangeEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIProgramChangeEvent */

// Creates a program change event with a channel and program number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/init(channel:programNumber:)
func NewMIDIProgramChangeEventWithChannelProgramNumber(channel objectivec.IObject, programNumber objectivec.IObject) MIDIProgramChangeEvent {
	instance := getMIDIProgramChangeEventClass().Alloc()
	rv := objc.Send[MIDIProgramChangeEvent](instance.ID, objc.Sel("initWithChannel:programNumber:"), channel, programNumber)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDIProgramChangeEventWithChannelProgramNumber */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIProgramChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIProgramChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIProgramChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIProgramChangeEvent */

// The MIDI program number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/programNumber
func (m_ MIDIProgramChangeEvent) ProgramNumber() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("programNumber"))
	return rv
}/* debug [instance_properties/getter]: programNumber */


// The MIDI program number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/programNumber
func (m_ MIDIProgramChangeEvent) SetProgramNumber(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramNumber:"), value)
}/* debug [instance_properties/setter]: programNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDIProgramChangeEvent */


