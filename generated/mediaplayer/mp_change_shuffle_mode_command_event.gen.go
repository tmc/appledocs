// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangeShuffleModeCommandEvent */


/* debug [class_header]: Header for MPChangeShuffleModeCommandEvent */
// The class instance for the [ChangeShuffleModeCommandEvent] class.
var (
	ChangeShuffleModeCommandEventClass     _ChangeShuffleModeCommandEventClass
	ChangeShuffleModeCommandEventClassOnce sync.Once
)

func getChangeShuffleModeCommandEventClass() _ChangeShuffleModeCommandEventClass {
	ChangeShuffleModeCommandEventClassOnce.Do(func() {
		ChangeShuffleModeCommandEventClass = _ChangeShuffleModeCommandEventClass{objc.GetClass("MPChangeShuffleModeCommandEvent")}
	})
	return ChangeShuffleModeCommandEventClass
}

type _ChangeShuffleModeCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangeShuffleModeCommandEvent */
// An interface definition for the [ChangeShuffleModeCommandEvent] class.
type IChangeShuffleModeCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for ChangeShuffleModeCommandEvent */
	// properties:
	PreservesShuffleMode() bool
	ShuffleType() ShuffleType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangeShuffleModeCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangeShuffleModeCommandEvent */
// Alloc allocates a new instance without initialization.
func (cc _ChangeShuffleModeCommandEventClass) Alloc() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangeShuffleModeCommandEventClass) New() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeShuffleModeCommandEvent) Init() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeShuffleModeCommandEvent) Autorelease() ChangeShuffleModeCommandEvent {
	rv := objc.Send[ChangeShuffleModeCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeShuffleModeCommandEvent creates a new ChangeShuffleModeCommandEvent instance.
func NewChangeShuffleModeCommandEvent() ChangeShuffleModeCommandEvent {
	return getChangeShuffleModeCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangeShuffleModeCommandEvent */
// An event requesting a change in the shuffle mode.


// An event requesting a change in the shuffle mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommandEvent
type ChangeShuffleModeCommandEvent struct {
	RemoteCommandEvent
}

// ChangeShuffleModeCommandEventFrom constructs a [ChangeShuffleModeCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the shuffle mode.
func ChangeShuffleModeCommandEventFrom(ptr unsafe.Pointer) ChangeShuffleModeCommandEvent {
	return ChangeShuffleModeCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangeShuffleModeCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangeShuffleModeCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangeShuffleModeCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangeShuffleModeCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangeShuffleModeCommandEvent */

// A Boolean value that indicates whether the shuffle mode is preserved between playback sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommandEvent/preservesShuffleMode
func (c_ ChangeShuffleModeCommandEvent) PreservesShuffleMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesShuffleMode"))
	return rv
}/* debug [instance_properties/getter]: preservesShuffleMode */


// The shuffle type used when fulfilling the event request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommandEvent/shuffleType
func (c_ ChangeShuffleModeCommandEvent) ShuffleType() ShuffleType {
	rv := objc.Send[ShuffleType](c_.ID, objc.Sel("shuffleType"))
	return rv
}/* debug [instance_properties/getter]: shuffleType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangeShuffleModeCommandEvent */



