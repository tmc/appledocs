// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMusicUserEvent */


/* debug [class_header]: Header for AVMusicUserEvent */
// The class instance for the [MusicUserEvent] class.
var (
	MusicUserEventClass     _MusicUserEventClass
	MusicUserEventClassOnce sync.Once
)

func getMusicUserEventClass() _MusicUserEventClass {
	MusicUserEventClassOnce.Do(func() {
		MusicUserEventClass = _MusicUserEventClass{objc.GetClass("AVMusicUserEvent")}
	})
	return MusicUserEventClass
}

type _MusicUserEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicUserEvent */
// An interface definition for the [MusicUserEvent] class.
type IMusicUserEvent interface {
	IMusicEvent
	
/* debug [class_interface_properties]: Properties for MusicUserEvent */
	// properties:
	SizeInBytes() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicUserEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicUserEvent */
// Alloc allocates a new instance without initialization.
func (mc _MusicUserEventClass) Alloc() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicUserEventClass) New() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicUserEvent) Init() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicUserEvent) Autorelease() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicUserEvent creates a new MusicUserEvent instance.
func NewMusicUserEvent() MusicUserEvent {
	return getMusicUserEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicUserEvent */
// An object that represents a custom user message.
//
// When playback of an reaches this event, the system calls the track’s callback. You can’t modify the size and contents of an once you create it.


// An object that represents a custom user message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent
type MusicUserEvent struct {
	MusicEvent
}

// MusicUserEventFrom constructs a [MusicUserEvent] from an unsafe.Pointer.
//
// An object that represents a custom user message.
func MusicUserEventFrom(ptr unsafe.Pointer) MusicUserEvent {
	return MusicUserEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicUserEvent */

// Creates a user event with the data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/init(data:)
func NewMusicUserEventWithData(data objc.IObject /* cross-framework: NSData */) MusicUserEvent {
	instance := getMusicUserEventClass().Alloc()
	rv := objc.Send[MusicUserEvent](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMusicUserEventWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicUserEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicUserEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicUserEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicUserEvent */

// The size of the data that the user event represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/sizeInBytes
func (m_ MusicUserEvent) SizeInBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sizeInBytes"))
	return rv
}/* debug [instance_properties/getter]: sizeInBytes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMusicUserEvent */


