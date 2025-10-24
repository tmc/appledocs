// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSeekCommandEvent */


/* debug [class_header]: Header for MPSeekCommandEvent */
// The class instance for the [SeekCommandEvent] class.
var (
	SeekCommandEventClass     _SeekCommandEventClass
	SeekCommandEventClassOnce sync.Once
)

func getSeekCommandEventClass() _SeekCommandEventClass {
	SeekCommandEventClassOnce.Do(func() {
		SeekCommandEventClass = _SeekCommandEventClass{objc.GetClass("MPSeekCommandEvent")}
	})
	return SeekCommandEventClass
}

type _SeekCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SeekCommandEvent */
// An interface definition for the [SeekCommandEvent] class.
type ISeekCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for SeekCommandEvent */
	// properties:
	Type() SeekCommandEventType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SeekCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SeekCommandEvent */
// Alloc allocates a new instance without initialization.
func (sc _SeekCommandEventClass) Alloc() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SeekCommandEventClass) New() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SeekCommandEvent) Init() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SeekCommandEvent) Autorelease() SeekCommandEvent {
	rv := objc.Send[SeekCommandEvent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSeekCommandEvent creates a new SeekCommandEvent instance.
func NewSeekCommandEvent() SeekCommandEvent {
	return getSeekCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SeekCommandEvent */
// An event requesting that the player seek to a new position.


// An event requesting that the player seek to a new position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSeekCommandEvent
type SeekCommandEvent struct {
	RemoteCommandEvent
}

// SeekCommandEventFrom constructs a [SeekCommandEvent] from an unsafe.Pointer.
//
// An event requesting that the player seek to a new position.
func SeekCommandEventFrom(ptr unsafe.Pointer) SeekCommandEvent {
	return SeekCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SeekCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SeekCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SeekCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SeekCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SeekCommandEvent */

// The type of seek command event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSeekCommandEvent/type
func (s_ SeekCommandEvent) Type() SeekCommandEventType {
	rv := objc.Send[SeekCommandEventType](s_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSeekCommandEvent */



