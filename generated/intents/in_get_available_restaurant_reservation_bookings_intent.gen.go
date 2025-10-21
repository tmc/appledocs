// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INGetAvailableRestaurantReservationBookingsIntent] class.
var (
	INGetAvailableRestaurantReservationBookingsIntentClass     _INGetAvailableRestaurantReservationBookingsIntentClass
	INGetAvailableRestaurantReservationBookingsIntentClassOnce sync.Once
)

func getINGetAvailableRestaurantReservationBookingsIntentClass() _INGetAvailableRestaurantReservationBookingsIntentClass {
	INGetAvailableRestaurantReservationBookingsIntentClassOnce.Do(func() {
		INGetAvailableRestaurantReservationBookingsIntentClass = _INGetAvailableRestaurantReservationBookingsIntentClass{objc.GetClass("INGetAvailableRestaurantReservationBookingsIntent")}
	})
	return INGetAvailableRestaurantReservationBookingsIntentClass
}

type _INGetAvailableRestaurantReservationBookingsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetAvailableRestaurantReservationBookingsIntent] class.
type IINGetAvailableRestaurantReservationBookingsIntent interface {
	IINIntent
}

// A request for the time slots available for making a reservation.
//
// An object asks you to generate details regarding the available time slots offered by a restaurant for a given party size. Maps sends this intent to your Intents extension when the user begins the booking process. You use this intent to obtain the initial details about the reservation request, including the number of people and the preferred date for the reservation. You use those details to identify potential time slots that can accommodate the party and return those time slots in your response object. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should resolve and confirm any parameters and create an object with the list of potential time slots.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetAvailableRestaurantReservationBookingsIntent
type INGetAvailableRestaurantReservationBookingsIntent struct {
	INIntent
}

// INGetAvailableRestaurantReservationBookingsIntentFrom constructs a [INGetAvailableRestaurantReservationBookingsIntent] from an unsafe.Pointer.
//
// A request for the time slots available for making a reservation.
func INGetAvailableRestaurantReservationBookingsIntentFrom(ptr unsafe.Pointer) INGetAvailableRestaurantReservationBookingsIntent {
	return INGetAvailableRestaurantReservationBookingsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetAvailableRestaurantReservationBookingsIntentClass) Alloc() INGetAvailableRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetAvailableRestaurantReservationBookingsIntentClass) New() INGetAvailableRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetAvailableRestaurantReservationBookingsIntent) Init() INGetAvailableRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetAvailableRestaurantReservationBookingsIntent) Autorelease() INGetAvailableRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetAvailableRestaurantReservationBookingsIntent creates a new INGetAvailableRestaurantReservationBookingsIntent instance.
func NewINGetAvailableRestaurantReservationBookingsIntent() INGetAvailableRestaurantReservationBookingsIntent {
	return getINGetAvailableRestaurantReservationBookingsIntentClass().New()
}


// The earliest date for which to return results.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/earliestbookingdateforresults
func (i_ INGetAvailableRestaurantReservationBookingsIntent) EarliestBookingDateForResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("earliestBookingDateForResults"))
	return rv
}


// SetEarliestBookingDateForResults sets the value of the earliestBookingDateForResults property.
// The earliest date for which to return results.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/earliestbookingdateforresults
func (i_ INGetAvailableRestaurantReservationBookingsIntent) SetEarliestBookingDateForResults(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEarliestBookingDateForResults:"), value)
}

// The latest date for which to return results.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/latestbookingdateforresults
func (i_ INGetAvailableRestaurantReservationBookingsIntent) LatestBookingDateForResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("latestBookingDateForResults"))
	return rv
}


// SetLatestBookingDateForResults sets the value of the latestBookingDateForResults property.
// The latest date for which to return results.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/latestbookingdateforresults
func (i_ INGetAvailableRestaurantReservationBookingsIntent) SetLatestBookingDateForResults(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLatestBookingDateForResults:"), value)
}

// The maximum number of reservation results to return.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/maximumnumberofresults
func (i_ INGetAvailableRestaurantReservationBookingsIntent) MaximumNumberOfResults() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("maximumNumberOfResults"))
	return rv
}


// SetMaximumNumberOfResults sets the value of the maximumNumberOfResults property.
// The maximum number of reservation results to return.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/maximumnumberofresults
func (i_ INGetAvailableRestaurantReservationBookingsIntent) SetMaximumNumberOfResults(value foundation.Number) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaximumNumberOfResults:"), value)
}

// The number of people in the guest’s party.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/partysize
func (i_ INGetAvailableRestaurantReservationBookingsIntent) PartySize() int {
	rv := objc.Send[int](i_.ID, objc.Sel("partySize"))
	return rv
}


// SetPartySize sets the value of the partySize property.
// The number of people in the guest’s party.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/partysize
func (i_ INGetAvailableRestaurantReservationBookingsIntent) SetPartySize(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPartySize:"), value)
}

// The date and time preferred by the user for the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/preferredbookingdatecomponents
func (i_ INGetAvailableRestaurantReservationBookingsIntent) PreferredBookingDateComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("preferredBookingDateComponents"))
	return rv
}


// SetPreferredBookingDateComponents sets the value of the preferredBookingDateComponents property.
// The date and time preferred by the user for the reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/preferredbookingdatecomponents
func (i_ INGetAvailableRestaurantReservationBookingsIntent) SetPreferredBookingDateComponents(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredBookingDateComponents:"), value)
}

// The restaurant associated with the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/restaurant
func (i_ INGetAvailableRestaurantReservationBookingsIntent) Restaurant() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("restaurant"))
	return rv
}


// SetRestaurant sets the value of the restaurant property.
// The restaurant associated with the reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetavailablerestaurantreservationbookingsintent/restaurant
func (i_ INGetAvailableRestaurantReservationBookingsIntent) SetRestaurant(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRestaurant:"), value)
}



