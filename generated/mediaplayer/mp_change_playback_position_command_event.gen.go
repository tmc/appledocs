// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangePlaybackPositionCommandEvent */


/* debug [class_header]: Header for MPChangePlaybackPositionCommandEvent */
// The class instance for the [ChangePlaybackPositionCommandEvent] class.
var (
	ChangePlaybackPositionCommandEventClass     _ChangePlaybackPositionCommandEventClass
	ChangePlaybackPositionCommandEventClassOnce sync.Once
)

func getChangePlaybackPositionCommandEventClass() _ChangePlaybackPositionCommandEventClass {
	ChangePlaybackPositionCommandEventClassOnce.Do(func() {
		ChangePlaybackPositionCommandEventClass = _ChangePlaybackPositionCommandEventClass{objc.GetClass("MPChangePlaybackPositionCommandEvent")}
	})
	return ChangePlaybackPositionCommandEventClass
}

type _ChangePlaybackPositionCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangePlaybackPositionCommandEvent */
// An interface definition for the [ChangePlaybackPositionCommandEvent] class.
type IChangePlaybackPositionCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for ChangePlaybackPositionCommandEvent */
	// properties:
	PositionTime() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangePlaybackPositionCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangePlaybackPositionCommandEvent */
// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackPositionCommandEventClass) Alloc() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangePlaybackPositionCommandEventClass) New() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackPositionCommandEvent) Init() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackPositionCommandEvent) Autorelease() ChangePlaybackPositionCommandEvent {
	rv := objc.Send[ChangePlaybackPositionCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackPositionCommandEvent creates a new ChangePlaybackPositionCommandEvent instance.
func NewChangePlaybackPositionCommandEvent() ChangePlaybackPositionCommandEvent {
	return getChangePlaybackPositionCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangePlaybackPositionCommandEvent */
// An event requesting a change in the playback position.


// An event requesting a change in the playback position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackPositionCommandEvent
type ChangePlaybackPositionCommandEvent struct {
	RemoteCommandEvent
}

// ChangePlaybackPositionCommandEventFrom constructs a [ChangePlaybackPositionCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the playback position.
func ChangePlaybackPositionCommandEventFrom(ptr unsafe.Pointer) ChangePlaybackPositionCommandEvent {
	return ChangePlaybackPositionCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangePlaybackPositionCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangePlaybackPositionCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangePlaybackPositionCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangePlaybackPositionCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangePlaybackPositionCommandEvent */

// The playback position used when setting the current time of the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackPositionCommandEvent/positionTime
func (c_ ChangePlaybackPositionCommandEvent) PositionTime() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("positionTime"))
	return rv
}/* debug [instance_properties/getter]: positionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangePlaybackPositionCommandEvent */



