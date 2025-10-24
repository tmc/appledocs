// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [FeedbackCommandEvent] class.
type IFeedbackCommandEvent interface {
	IRemoteCommandEvent
	// properties:
	IsNegative() bool
	SetIsNegative(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (fc _FeedbackCommandEventClass) Alloc() FeedbackCommandEvent {
	rv := objc.Send[FeedbackCommandEvent](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether an app should perform a negative command appropriate to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommandevent/isnegative
func (f_ FeedbackCommandEvent) IsNegative() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isNegative"))
	return rv
}


// A Boolean value that indicates whether an app should perform a negative command appropriate to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpfeedbackcommandevent/isnegative
func (f_ FeedbackCommandEvent) SetIsNegative(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsNegative:"), value)
}



