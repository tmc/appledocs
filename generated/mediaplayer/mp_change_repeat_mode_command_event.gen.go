// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangeRepeatModeCommandEvent */


/* debug [class_header]: Header for MPChangeRepeatModeCommandEvent */
// The class instance for the [ChangeRepeatModeCommandEvent] class.
var (
	ChangeRepeatModeCommandEventClass     _ChangeRepeatModeCommandEventClass
	ChangeRepeatModeCommandEventClassOnce sync.Once
)

func getChangeRepeatModeCommandEventClass() _ChangeRepeatModeCommandEventClass {
	ChangeRepeatModeCommandEventClassOnce.Do(func() {
		ChangeRepeatModeCommandEventClass = _ChangeRepeatModeCommandEventClass{objc.GetClass("MPChangeRepeatModeCommandEvent")}
	})
	return ChangeRepeatModeCommandEventClass
}

type _ChangeRepeatModeCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangeRepeatModeCommandEvent */
// An interface definition for the [ChangeRepeatModeCommandEvent] class.
type IChangeRepeatModeCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for ChangeRepeatModeCommandEvent */
	// properties:
	PreservesRepeatMode() bool
	RepeatType() RepeatType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangeRepeatModeCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangeRepeatModeCommandEvent */
// Alloc allocates a new instance without initialization.
func (cc _ChangeRepeatModeCommandEventClass) Alloc() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangeRepeatModeCommandEventClass) New() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeRepeatModeCommandEvent) Init() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeRepeatModeCommandEvent) Autorelease() ChangeRepeatModeCommandEvent {
	rv := objc.Send[ChangeRepeatModeCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeRepeatModeCommandEvent creates a new ChangeRepeatModeCommandEvent instance.
func NewChangeRepeatModeCommandEvent() ChangeRepeatModeCommandEvent {
	return getChangeRepeatModeCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangeRepeatModeCommandEvent */
// An event requesting a change in the repeat mode.


// An event requesting a change in the repeat mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommandEvent
type ChangeRepeatModeCommandEvent struct {
	RemoteCommandEvent
}

// ChangeRepeatModeCommandEventFrom constructs a [ChangeRepeatModeCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the repeat mode.
func ChangeRepeatModeCommandEventFrom(ptr unsafe.Pointer) ChangeRepeatModeCommandEvent {
	return ChangeRepeatModeCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangeRepeatModeCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangeRepeatModeCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangeRepeatModeCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangeRepeatModeCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangeRepeatModeCommandEvent */

// A Boolean value that indicates whether the chosen repeat mode is preserved between playback sessions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommandEvent/preservesRepeatMode
func (c_ ChangeRepeatModeCommandEvent) PreservesRepeatMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesRepeatMode"))
	return rv
}/* debug [instance_properties/getter]: preservesRepeatMode */


// The repeat type used when fulfilling the event request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommandEvent/repeatType
func (c_ ChangeRepeatModeCommandEvent) RepeatType() RepeatType {
	rv := objc.Send[RepeatType](c_.ID, objc.Sel("repeatType"))
	return rv
}/* debug [instance_properties/getter]: repeatType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangeRepeatModeCommandEvent */



