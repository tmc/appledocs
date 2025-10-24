// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMIDIMetaEvent */


/* debug [class_header]: Header for AVMIDIMetaEvent */
// The class instance for the [MIDIMetaEvent] class.
var (
	MIDIMetaEventClass     _MIDIMetaEventClass
	MIDIMetaEventClassOnce sync.Once
)

func getMIDIMetaEventClass() _MIDIMetaEventClass {
	MIDIMetaEventClassOnce.Do(func() {
		MIDIMetaEventClass = _MIDIMetaEventClass{objc.GetClass("AVMIDIMetaEvent")}
	})
	return MIDIMetaEventClass
}

type _MIDIMetaEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIMetaEvent */
// An interface definition for the [MIDIMetaEvent] class.
type IMIDIMetaEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for MIDIMetaEvent */
	// properties:
	Type() MIDIMetaEventType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIMetaEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIMetaEvent */
// Alloc allocates a new instance without initialization.
func (mc _MIDIMetaEventClass) Alloc() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIMetaEventClass) New() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIMetaEvent) Init() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIMetaEvent) Autorelease() MIDIMetaEvent {
	rv := objc.Send[MIDIMetaEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIMetaEvent creates a new MIDIMetaEvent instance.
func NewMIDIMetaEvent() MIDIMetaEvent {
	return getMIDIMetaEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIMetaEvent */
// An object that represents MIDI meta event messages.
//
// You can’t modify the size and contents of this event once you create it. This doesn’t verify that the content matches the MIDI specification. You can only add , , or to a sequence’s tempo track.


// An object that represents MIDI meta event messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent
type MIDIMetaEvent struct {
	MusicEvent
}

// MIDIMetaEventFrom constructs a [MIDIMetaEvent] from an unsafe.Pointer.
//
// An object that represents MIDI meta event messages.
func MIDIMetaEventFrom(ptr unsafe.Pointer) MIDIMetaEvent {
	return MIDIMetaEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIMetaEvent */

// Creates an event with a MIDI meta event type and data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/init(type:data:)
func NewMIDIMetaEventWithTypeData(type_ MIDIMetaEventType, data objc.IObject /* cross-framework: NSData */) MIDIMetaEvent {
	instance := getMIDIMetaEventClass().Alloc()
	rv := objc.Send[MIDIMetaEvent](instance.ID, objc.Sel("initWithType:data:"), type_, data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDIMetaEventWithTypeData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIMetaEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIMetaEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIMetaEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIMetaEvent */

// The type of meta event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/type
func (m_ MIDIMetaEvent) Type() MIDIMetaEventType {
	rv := objc.Send[MIDIMetaEventType](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDIMetaEvent */


