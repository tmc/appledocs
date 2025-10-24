// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPFeedbackCommandEvent */


/* debug [class_header]: Header for MPFeedbackCommandEvent */
// The class instance for the [FeedbackCommandEvent] class.
var (
	FeedbackCommandEventClass     _FeedbackCommandEventClass
	FeedbackCommandEventClassOnce sync.Once
)

func getFeedbackCommandEventClass() _FeedbackCommandEventClass {
	FeedbackCommandEventClassOnce.Do(func() {
		FeedbackCommandEventClass = _FeedbackCommandEventClass{objc.GetClass("MPFeedbackCommandEvent")}
	})
	return FeedbackCommandEventClass
}

type _FeedbackCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FeedbackCommandEvent */
// An interface definition for the [FeedbackCommandEvent] class.
type IFeedbackCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for FeedbackCommandEvent */
	// properties:
	Negative() bool
	IsNegative() bool
	SetIsNegative(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FeedbackCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FeedbackCommandEvent */
// Alloc allocates a new instance without initialization.
func (fc _FeedbackCommandEventClass) Alloc() FeedbackCommandEvent {
	rv := objc.Send[FeedbackCommandEvent](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FeedbackCommandEventClass) New() FeedbackCommandEvent {
	rv := objc.Send[FeedbackCommandEvent](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeedbackCommandEvent) Init() FeedbackCommandEvent {
	rv := objc.Send[FeedbackCommandEvent](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeedbackCommandEvent) Autorelease() FeedbackCommandEvent {
	rv := objc.Send[FeedbackCommandEvent](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeedbackCommandEvent creates a new FeedbackCommandEvent instance.
func NewFeedbackCommandEvent() FeedbackCommandEvent {
	return getFeedbackCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FeedbackCommandEvent */
// An event requesting a change in the feedback setting.


// An event requesting a change in the feedback setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommandEvent
type FeedbackCommandEvent struct {
	RemoteCommandEvent
}

// FeedbackCommandEventFrom constructs a [FeedbackCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the feedback setting.
func FeedbackCommandEventFrom(ptr unsafe.Pointer) FeedbackCommandEvent {
	return FeedbackCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FeedbackCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FeedbackCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FeedbackCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FeedbackCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FeedbackCommandEvent */

// A Boolean value that indicates whether an app should perform a negative command appropriate to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPFeedbackCommandEvent/isNegative
func (f_ FeedbackCommandEvent) Negative() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("negative"))
	return rv
}/* debug [instance_properties/getter]: negative */


// A Boolean value that indicates whether an app should perform a negative command appropriate to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommandevent/isnegative
func (f_ FeedbackCommandEvent) IsNegative() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isNegative"))
	return rv
}/* debug [instance_properties/getter]: isNegative */


// A Boolean value that indicates whether an app should perform a negative command appropriate to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommandevent/isnegative
func (f_ FeedbackCommandEvent) SetIsNegative(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsNegative:"), value)
}/* debug [instance_properties/setter]: isNegative */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPFeedbackCommandEvent */



