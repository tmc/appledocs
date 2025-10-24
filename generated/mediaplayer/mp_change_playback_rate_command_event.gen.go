// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangePlaybackRateCommandEvent */


/* debug [class_header]: Header for MPChangePlaybackRateCommandEvent */
// The class instance for the [ChangePlaybackRateCommandEvent] class.
var (
	ChangePlaybackRateCommandEventClass     _ChangePlaybackRateCommandEventClass
	ChangePlaybackRateCommandEventClassOnce sync.Once
)

func getChangePlaybackRateCommandEventClass() _ChangePlaybackRateCommandEventClass {
	ChangePlaybackRateCommandEventClassOnce.Do(func() {
		ChangePlaybackRateCommandEventClass = _ChangePlaybackRateCommandEventClass{objc.GetClass("MPChangePlaybackRateCommandEvent")}
	})
	return ChangePlaybackRateCommandEventClass
}

type _ChangePlaybackRateCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangePlaybackRateCommandEvent */
// An interface definition for the [ChangePlaybackRateCommandEvent] class.
type IChangePlaybackRateCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for ChangePlaybackRateCommandEvent */
	// properties:
	PlaybackRate() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangePlaybackRateCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangePlaybackRateCommandEvent */
// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackRateCommandEventClass) Alloc() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangePlaybackRateCommandEventClass) New() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackRateCommandEvent) Init() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackRateCommandEvent) Autorelease() ChangePlaybackRateCommandEvent {
	rv := objc.Send[ChangePlaybackRateCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackRateCommandEvent creates a new ChangePlaybackRateCommandEvent instance.
func NewChangePlaybackRateCommandEvent() ChangePlaybackRateCommandEvent {
	return getChangePlaybackRateCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangePlaybackRateCommandEvent */
// An event requesting a change in the playback rate.


// An event requesting a change in the playback rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommandEvent
type ChangePlaybackRateCommandEvent struct {
	RemoteCommandEvent
}

// ChangePlaybackRateCommandEventFrom constructs a [ChangePlaybackRateCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the playback rate.
func ChangePlaybackRateCommandEventFrom(ptr unsafe.Pointer) ChangePlaybackRateCommandEvent {
	return ChangePlaybackRateCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangePlaybackRateCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangePlaybackRateCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangePlaybackRateCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangePlaybackRateCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangePlaybackRateCommandEvent */

// The chosen playback rate for the command event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommandEvent/playbackRate
func (c_ ChangePlaybackRateCommandEvent) PlaybackRate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("playbackRate"))
	return rv
}/* debug [instance_properties/getter]: playbackRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangePlaybackRateCommandEvent */



