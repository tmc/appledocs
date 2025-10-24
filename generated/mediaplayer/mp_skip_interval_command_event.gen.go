// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSkipIntervalCommandEvent */


/* debug [class_header]: Header for MPSkipIntervalCommandEvent */
// The class instance for the [SkipIntervalCommandEvent] class.
var (
	SkipIntervalCommandEventClass     _SkipIntervalCommandEventClass
	SkipIntervalCommandEventClassOnce sync.Once
)

func getSkipIntervalCommandEventClass() _SkipIntervalCommandEventClass {
	SkipIntervalCommandEventClassOnce.Do(func() {
		SkipIntervalCommandEventClass = _SkipIntervalCommandEventClass{objc.GetClass("MPSkipIntervalCommandEvent")}
	})
	return SkipIntervalCommandEventClass
}

type _SkipIntervalCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SkipIntervalCommandEvent */
// An interface definition for the [SkipIntervalCommandEvent] class.
type ISkipIntervalCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for SkipIntervalCommandEvent */
	// properties:
	Interval() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SkipIntervalCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SkipIntervalCommandEvent */
// Alloc allocates a new instance without initialization.
func (sc _SkipIntervalCommandEventClass) Alloc() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SkipIntervalCommandEventClass) New() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SkipIntervalCommandEvent) Init() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SkipIntervalCommandEvent) Autorelease() SkipIntervalCommandEvent {
	rv := objc.Send[SkipIntervalCommandEvent](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkipIntervalCommandEvent creates a new SkipIntervalCommandEvent instance.
func NewSkipIntervalCommandEvent() SkipIntervalCommandEvent {
	return getSkipIntervalCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SkipIntervalCommandEvent */
// An event requesting a change in the current skip interval.


// An event requesting a change in the current skip interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommandEvent
type SkipIntervalCommandEvent struct {
	RemoteCommandEvent
}

// SkipIntervalCommandEventFrom constructs a [SkipIntervalCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the current skip interval.
func SkipIntervalCommandEventFrom(ptr unsafe.Pointer) SkipIntervalCommandEvent {
	return SkipIntervalCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SkipIntervalCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SkipIntervalCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SkipIntervalCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SkipIntervalCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SkipIntervalCommandEvent */

// The chosen interval, in seconds, for the skip command event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommandEvent/interval
func (s_ SkipIntervalCommandEvent) Interval() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("interval"))
	return rv
}/* debug [instance_properties/getter]: interval */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSkipIntervalCommandEvent */



