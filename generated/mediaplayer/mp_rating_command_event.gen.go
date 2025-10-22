// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [RatingCommandEvent] class.
type IRatingCommandEvent interface {
	IRemoteCommandEvent
	Rating() float32
	SetRating(value float32)
}

// An event requesting a change in the rating.
//
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

// Alloc allocates a new instance without initialization.
func (rc _RatingCommandEventClass) Alloc() RatingCommandEvent {
	rv := objc.Send[RatingCommandEvent](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The rating for the command event.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommandevent/rating
func (r_ RatingCommandEvent) Rating() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("rating"))
	return rv
}


// SetRating sets the value of the rating property.
// The rating for the command event.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpratingcommandevent/rating
func (r_ RatingCommandEvent) SetRating(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRating:"), value)
}



