// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMusicEvent */


/* debug [class_header]: Header for AVMusicEvent */
// The class instance for the [MusicEvent] class.
var (
	MusicEventClass     _MusicEventClass
	MusicEventClassOnce sync.Once
)

func getMusicEventClass() _MusicEventClass {
	MusicEventClassOnce.Do(func() {
		MusicEventClass = _MusicEventClass{objc.GetClass("AVMusicEvent")}
	})
	return MusicEventClass
}

type _MusicEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicEvent */
// An interface definition for the [MusicEvent] class.
type IMusicEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MusicEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicEvent */
// Alloc allocates a new instance without initialization.
func (mc _MusicEventClass) Alloc() MusicEvent {
	rv := objc.Send[MusicEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicEventClass) New() MusicEvent {
	rv := objc.Send[MusicEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicEvent) Init() MusicEvent {
	rv := objc.Send[MusicEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicEvent) Autorelease() MusicEvent {
	rv := objc.Send[MusicEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicEvent creates a new MusicEvent instance.
func NewMusicEvent() MusicEvent {
	return getMusicEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicEvent */
// A base class for the events you associate with a music track.


// A base class for the events you associate with a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicEvent
type MusicEvent struct {
	objectivec.Object
}

// MusicEventFrom constructs a [MusicEvent] from an unsafe.Pointer.
//
// A base class for the events you associate with a music track.
func MusicEventFrom(ptr unsafe.Pointer) MusicEvent {
	return MusicEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMusicEvent */



