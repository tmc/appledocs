// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMIDISysexEvent */


/* debug [class_header]: Header for AVMIDISysexEvent */
// The class instance for the [MIDISysexEvent] class.
var (
	MIDISysexEventClass     _MIDISysexEventClass
	MIDISysexEventClassOnce sync.Once
)

func getMIDISysexEventClass() _MIDISysexEventClass {
	MIDISysexEventClassOnce.Do(func() {
		MIDISysexEventClass = _MIDISysexEventClass{objc.GetClass("AVMIDISysexEvent")}
	})
	return MIDISysexEventClass
}

type _MIDISysexEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDISysexEvent */
// An interface definition for the [MIDISysexEvent] class.
type IMIDISysexEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for MIDISysexEvent */
	// properties:
	SizeInBytes() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDISysexEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDISysexEvent */
// Alloc allocates a new instance without initialization.
func (mc _MIDISysexEventClass) Alloc() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDISysexEventClass) New() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDISysexEvent) Init() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDISysexEvent) Autorelease() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDISysexEvent creates a new MIDISysexEvent instance.
func NewMIDISysexEvent() MIDISysexEvent {
	return getMIDISysexEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDISysexEvent */
// An object that represents a MIDI system exclusive message.
//
// You can’t modify the size and contents of this event once you create it.


// An object that represents a MIDI system exclusive message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent
type MIDISysexEvent struct {
	MusicEvent
}

// MIDISysexEventFrom constructs a [MIDISysexEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI system exclusive message.
func MIDISysexEventFrom(ptr unsafe.Pointer) MIDISysexEvent {
	return MIDISysexEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDISysexEvent */

// Creates a system event with the data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/init(data:)
func NewMIDISysexEventWithData(data objc.IObject /* cross-framework: NSData */) MIDISysexEvent {
	instance := getMIDISysexEventClass().Alloc()
	rv := objc.Send[MIDISysexEvent](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDISysexEventWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDISysexEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDISysexEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDISysexEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDISysexEvent */

// The size of the data that this event contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/sizeInBytes
func (m_ MIDISysexEvent) SizeInBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sizeInBytes"))
	return rv
}/* debug [instance_properties/getter]: sizeInBytes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMIDISysexEvent */


