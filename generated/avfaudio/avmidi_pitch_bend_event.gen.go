// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMIDIPitchBendEvent */


/* debug [class_header]: Header for AVMIDIPitchBendEvent */
// The class instance for the [MIDIPitchBendEvent] class.
var (
	MIDIPitchBendEventClass     _MIDIPitchBendEventClass
	MIDIPitchBendEventClassOnce sync.Once
)

func getMIDIPitchBendEventClass() _MIDIPitchBendEventClass {
	MIDIPitchBendEventClassOnce.Do(func() {
		MIDIPitchBendEventClass = _MIDIPitchBendEventClass{objc.GetClass("AVMIDIPitchBendEvent")}
	})
	return MIDIPitchBendEventClass
}

type _MIDIPitchBendEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIPitchBendEvent */
// An interface definition for the [MIDIPitchBendEvent] class.
type IMIDIPitchBendEvent interface {
	IMIDIChannelEvent
	
/* debug [class_interface_properties]: Properties for MIDIPitchBendEvent */
	// properties:
	Value() objectivec.IObject
	SetValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIPitchBendEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIPitchBendEvent */
// Alloc allocates a new instance without initialization.
func (mc _MIDIPitchBendEventClass) Alloc() MIDIPitchBendEvent {
	rv := objc.Send[MIDIPitchBendEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIPitchBendEventClass) New() MIDIPitchBendEvent {
	rv := objc.Send[MIDIPitchBendEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIPitchBendEvent) Init() MIDIPitchBendEvent {
	rv := objc.Send[MIDIPitchBendEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIPitchBendEvent) Autorelease() MIDIPitchBendEvent {
	rv := objc.Send[MIDIPitchBendEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIPitchBendEvent creates a new MIDIPitchBendEvent instance.
func NewMIDIPitchBendEvent() MIDIPitchBendEvent {
	return getMIDIPitchBendEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIPitchBendEvent */
// An object that represents a MIDI pitch bend message.


// An object that represents a MIDI pitch bend message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent
type MIDIPitchBendEvent struct {
	MIDIChannelEvent
}

// MIDIPitchBendEventFrom constructs a [MIDIPitchBendEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI pitch bend message.
func MIDIPitchBendEventFrom(ptr unsafe.Pointer) MIDIPitchBendEvent {
	return MIDIPitchBendEvent{
		MIDIChannelEvent: MIDIChannelEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIPitchBendEvent */

// Creates an event with a channel and pitch bend value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent/init(channel:value:)
func NewMIDIPitchBendEventWithChannelValue(channel objectivec.IObject, value objectivec.IObject) MIDIPitchBendEvent {
	instance := getMIDIPitchBendEventClass().Alloc()
	rv := objc.Send[MIDIPitchBendEvent](instance.ID, objc.Sel("initWithChannel:value:"), channel, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDIPitchBendEventWithChannelValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIPitchBendEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIPitchBendEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIPitchBendEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIPitchBendEvent */

// The value of the pitch bend event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent/value
func (m_ MIDIPitchBendEvent) Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The value of the pitch bend event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent/value
func (m_ MIDIPitchBendEvent) SetValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDIPitchBendEvent */


