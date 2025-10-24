// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMIDINoteEvent */


/* debug [class_header]: Header for AVMIDINoteEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDINoteEvent */
// An interface definition for the [MIDINoteEvent] class.
type IMIDINoteEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for MIDINoteEvent */
	// properties:
	Channel() objectivec.IObject
	SetChannel(value objectivec.IObject)
	Duration() MusicTimeStamp /* typedef */
	SetDuration(value MusicTimeStamp /* typedef */)
	Key() objectivec.IObject
	SetKey(value objectivec.IObject)
	Velocity() objectivec.IObject
	SetVelocity(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDINoteEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDINoteEvent */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDINoteEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDINoteEvent */

// Creates an event with a MIDI channel, key number, velocity, and duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/init(channel:key:velocity:duration:)
func NewMIDINoteEventWithChannelKeyVelocityDuration(channel objectivec.IObject, keyNum objectivec.IObject, velocity objectivec.IObject, duration MusicTimeStamp /* typedef */) MIDINoteEvent {
	instance := getMIDINoteEventClass().Alloc()
	rv := objc.Send[MIDINoteEvent](instance.ID, objc.Sel("initWithChannel:key:velocity:duration:"), channel, keyNum, velocity, duration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDINoteEventWithChannelKeyVelocityDuration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDINoteEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDINoteEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDINoteEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDINoteEvent */

// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/channel
func (m_ MIDINoteEvent) Channel() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("channel"))
	return rv
}/* debug [instance_properties/getter]: channel */


// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/channel
func (m_ MIDINoteEvent) SetChannel(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}/* debug [instance_properties/setter]: channel */


// The duration for the note, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/duration
func (m_ MIDINoteEvent) Duration() MusicTimeStamp /* typedef */ {
	rv := objc.Send[float64](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration for the note, in beats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/duration
func (m_ MIDINoteEvent) SetDuration(value MusicTimeStamp /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The MIDI key number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/key
func (m_ MIDINoteEvent) Key() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// The MIDI key number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/key
func (m_ MIDINoteEvent) SetKey(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}/* debug [instance_properties/setter]: key */


// The MIDI velocity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/velocity
func (m_ MIDINoteEvent) Velocity() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("velocity"))
	return rv
}/* debug [instance_properties/getter]: velocity */


// The MIDI velocity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/velocity
func (m_ MIDINoteEvent) SetVelocity(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVelocity:"), value)
}/* debug [instance_properties/setter]: velocity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDINoteEvent */


