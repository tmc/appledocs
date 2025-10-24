// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPRatingCommandEvent */


/* debug [class_header]: Header for MPRatingCommandEvent */
// The class instance for the [RatingCommandEvent] class.
var (
	RatingCommandEventClass     _RatingCommandEventClass
	RatingCommandEventClassOnce sync.Once
)

func getRatingCommandEventClass() _RatingCommandEventClass {
	RatingCommandEventClassOnce.Do(func() {
		RatingCommandEventClass = _RatingCommandEventClass{objc.GetClass("MPRatingCommandEvent")}
	})
	return RatingCommandEventClass
}

type _RatingCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RatingCommandEvent */
// An interface definition for the [RatingCommandEvent] class.
type IRatingCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for RatingCommandEvent */
	// properties:
	Rating() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RatingCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RatingCommandEvent */
// Alloc allocates a new instance without initialization.
func (rc _RatingCommandEventClass) Alloc() RatingCommandEvent {
	rv := objc.Send[RatingCommandEvent](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RatingCommandEventClass) New() RatingCommandEvent {
	rv := objc.Send[RatingCommandEvent](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RatingCommandEvent) Init() RatingCommandEvent {
	rv := objc.Send[RatingCommandEvent](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RatingCommandEvent) Autorelease() RatingCommandEvent {
	rv := objc.Send[RatingCommandEvent](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRatingCommandEvent creates a new RatingCommandEvent instance.
func NewRatingCommandEvent() RatingCommandEvent {
	return getRatingCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RatingCommandEvent */
// An event requesting a change in the rating.


// An event requesting a change in the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommandEvent
type RatingCommandEvent struct {
	RemoteCommandEvent
}

// RatingCommandEventFrom constructs a [RatingCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the rating.
func RatingCommandEventFrom(ptr unsafe.Pointer) RatingCommandEvent {
	return RatingCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RatingCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RatingCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RatingCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RatingCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RatingCommandEvent */

// The rating for the command event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRatingCommandEvent/rating
func (r_ RatingCommandEvent) Rating() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("rating"))
	return rv
}/* debug [instance_properties/getter]: rating */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPRatingCommandEvent */



