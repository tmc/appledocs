// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetReservationDetailsIntentResponse] class.
var (
	INGetReservationDetailsIntentResponseClass     _INGetReservationDetailsIntentResponseClass
	INGetReservationDetailsIntentResponseClassOnce sync.Once
)

func getINGetReservationDetailsIntentResponseClass() _INGetReservationDetailsIntentResponseClass {
	INGetReservationDetailsIntentResponseClassOnce.Do(func() {
		INGetReservationDetailsIntentResponseClass = _INGetReservationDetailsIntentResponseClass{objc.GetClass("INGetReservationDetailsIntentResponse")}
	})
	return INGetReservationDetailsIntentResponseClass
}

type _INGetReservationDetailsIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INGetReservationDetailsIntentResponse] class.
type IINGetReservationDetailsIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a request for reservation details.
//
// Use an object to specify the results of a user requesting reservation details in your app. After getting the reservation details action using the criteria specified in the object, create an instance of this class with the results of the action. Siri can then use this information for system integrations, such as populating the calendar with an event.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetReservationDetailsIntentResponse
type INGetReservationDetailsIntentResponse struct {
	INIntentResponse
}

// INGetReservationDetailsIntentResponseFrom constructs a [INGetReservationDetailsIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a request for reservation details.
func INGetReservationDetailsIntentResponseFrom(ptr unsafe.Pointer) INGetReservationDetailsIntentResponse {
	return INGetReservationDetailsIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetReservationDetailsIntentResponseClass) Alloc() INGetReservationDetailsIntentResponse {
	rv := objc.Send[INGetReservationDetailsIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetReservationDetailsIntentResponseClass) New() INGetReservationDetailsIntentResponse {
	rv := objc.Send[INGetReservationDetailsIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetReservationDetailsIntentResponse) Init() INGetReservationDetailsIntentResponse {
	rv := objc.Send[INGetReservationDetailsIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetReservationDetailsIntentResponse) Autorelease() INGetReservationDetailsIntentResponse {
	rv := objc.Send[INGetReservationDetailsIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetReservationDetailsIntentResponse creates a new INGetReservationDetailsIntentResponse instance.
func NewINGetReservationDetailsIntentResponse() INGetReservationDetailsIntentResponse {
	return getINGetReservationDetailsIntentResponseClass().New()
}


// The code that indicates whether your app successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetreservationdetailsintentresponse/code
func (i_ INGetReservationDetailsIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}


// SetCode sets the value of the code property.
// The code that indicates whether your app successfully handled the intent.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetreservationdetailsintentresponse/code
func (i_ INGetReservationDetailsIntentResponse) SetCode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCode:"), value)
}

// An array containing reservations reqeusted by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetreservationdetailsintentresponse/reservations
func (i_ INGetReservationDetailsIntentResponse) Reservations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reservations"))
	return rv
}


// SetReservations sets the value of the reservations property.
// An array containing reservations reqeusted by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetreservationdetailsintentresponse/reservations
func (i_ INGetReservationDetailsIntentResponse) SetReservations(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservations:"), value)
}



