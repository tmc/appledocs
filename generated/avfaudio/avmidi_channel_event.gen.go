// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMIDIChannelEvent */


/* debug [class_header]: Header for AVMIDIChannelEvent */
// The class instance for the [MIDIChannelEvent] class.
var (
	MIDIChannelEventClass     _MIDIChannelEventClass
	MIDIChannelEventClassOnce sync.Once
)

func getMIDIChannelEventClass() _MIDIChannelEventClass {
	MIDIChannelEventClassOnce.Do(func() {
		MIDIChannelEventClass = _MIDIChannelEventClass{objc.GetClass("AVMIDIChannelEvent")}
	})
	return MIDIChannelEventClass
}

type _MIDIChannelEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIChannelEvent */
// An interface definition for the [MIDIChannelEvent] class.
type IMIDIChannelEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for MIDIChannelEvent */
	// properties:
	Channel() objectivec.IObject
	SetChannel(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIChannelEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIChannelEvent */
// Alloc allocates a new instance without initialization.
func (mc _MIDIChannelEventClass) Alloc() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIChannelEventClass) New() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIChannelEvent) Init() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIChannelEvent) Autorelease() MIDIChannelEvent {
	rv := objc.Send[MIDIChannelEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIChannelEvent creates a new MIDIChannelEvent instance.
func NewMIDIChannelEvent() MIDIChannelEvent {
	return getMIDIChannelEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIChannelEvent */
// A base class for all MIDI messages that operate on a single MIDI channel.


// A base class for all MIDI messages that operate on a single MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent
type MIDIChannelEvent struct {
	MusicEvent
}

// MIDIChannelEventFrom constructs a [MIDIChannelEvent] from an unsafe.Pointer.
//
// A base class for all MIDI messages that operate on a single MIDI channel.
func MIDIChannelEventFrom(ptr unsafe.Pointer) MIDIChannelEvent {
	return MIDIChannelEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIChannelEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIChannelEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIChannelEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIChannelEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIChannelEvent */

// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/channel
func (m_ MIDIChannelEvent) Channel() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("channel"))
	return rv
}/* debug [instance_properties/getter]: channel */


// The MIDI channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/channel
func (m_ MIDIChannelEvent) SetChannel(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}/* debug [instance_properties/setter]: channel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDIChannelEvent */



