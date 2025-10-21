// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INSendRideFeedbackIntent] class.
var (
	INSendRideFeedbackIntentClass     _INSendRideFeedbackIntentClass
	INSendRideFeedbackIntentClassOnce sync.Once
)

func getINSendRideFeedbackIntentClass() _INSendRideFeedbackIntentClass {
	INSendRideFeedbackIntentClassOnce.Do(func() {
		INSendRideFeedbackIntentClass = _INSendRideFeedbackIntentClass{objc.GetClass("INSendRideFeedbackIntent")}
	})
	return INSendRideFeedbackIntentClass
}

type _INSendRideFeedbackIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSendRideFeedbackIntent] class.
type IINSendRideFeedbackIntent interface {
	IINIntent
}

// An intent indicating that the user provided feedback for a completed ride.
//
// When the user provides feedback for a completed ride, SiriKit sends an object to your handler. SiriKit populates this intent object with the ride identifier and the feedback, including a possible driver rating and tip. Upon receiving this intent, validate the provided information and forward it along to your service. SiriKit guarantees that it provides at least one piece of feedback. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the status of the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendRideFeedbackIntent
type INSendRideFeedbackIntent struct {
	INIntent
}

// INSendRideFeedbackIntentFrom constructs a [INSendRideFeedbackIntent] from an unsafe.Pointer.
//
// An intent indicating that the user provided feedback for a completed ride.
func INSendRideFeedbackIntentFrom(ptr unsafe.Pointer) INSendRideFeedbackIntent {
	return INSendRideFeedbackIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendRideFeedbackIntentClass) Alloc() INSendRideFeedbackIntent {
	rv := objc.Send[INSendRideFeedbackIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendRideFeedbackIntentClass) New() INSendRideFeedbackIntent {
	rv := objc.Send[INSendRideFeedbackIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendRideFeedbackIntent) Init() INSendRideFeedbackIntent {
	rv := objc.Send[INSendRideFeedbackIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendRideFeedbackIntent) Autorelease() INSendRideFeedbackIntent {
	rv := objc.Send[INSendRideFeedbackIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendRideFeedbackIntent creates a new INSendRideFeedbackIntent instance.
func NewINSendRideFeedbackIntent() INSendRideFeedbackIntent {
	return getINSendRideFeedbackIntentClass().New()
}


// The unique identifier that you assigned to the ride.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendRideFeedbackIntent/rideIdentifier
func (i_ INSendRideFeedbackIntent) RideIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("rideIdentifier"))
	return rv
}

// The user-supplied rating for the driver.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendridefeedbackintent/rating
func (i_ INSendRideFeedbackIntent) Rating() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("rating"))
	return rv
}


// SetRating sets the value of the rating property.
// The user-supplied rating for the driver.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendridefeedbackintent/rating
func (i_ INSendRideFeedbackIntent) SetRating(value foundation.Number) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRating:"), value)
}

// The tip for the driver.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendridefeedbackintent/tip
func (i_ INSendRideFeedbackIntent) Tip() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("tip"))
	return rv
}


// SetTip sets the value of the tip property.
// The tip for the driver.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendridefeedbackintent/tip
func (i_ INSendRideFeedbackIntent) SetTip(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTip:"), value)
}



