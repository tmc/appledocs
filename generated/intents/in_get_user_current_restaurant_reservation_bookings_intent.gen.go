// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INGetUserCurrentRestaurantReservationBookingsIntent] class.
var (
	INGetUserCurrentRestaurantReservationBookingsIntentClass     _INGetUserCurrentRestaurantReservationBookingsIntentClass
	INGetUserCurrentRestaurantReservationBookingsIntentClassOnce sync.Once
)

func getINGetUserCurrentRestaurantReservationBookingsIntentClass() _INGetUserCurrentRestaurantReservationBookingsIntentClass {
	INGetUserCurrentRestaurantReservationBookingsIntentClassOnce.Do(func() {
		INGetUserCurrentRestaurantReservationBookingsIntentClass = _INGetUserCurrentRestaurantReservationBookingsIntentClass{objc.GetClass("INGetUserCurrentRestaurantReservationBookingsIntent")}
	})
	return INGetUserCurrentRestaurantReservationBookingsIntentClass
}

type _INGetUserCurrentRestaurantReservationBookingsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetUserCurrentRestaurantReservationBookingsIntent] class.
type IINGetUserCurrentRestaurantReservationBookingsIntent interface {
	IINIntent
}

// A request for the list of the user’s current reservations.
//
// An object asks you to retrieve the current restaurant reservations associated with the user. Maps sends this intent to your Intents extension when it needs information about all of the user’s current reservations, or when it needs information about one or more specific reservations. Use the properties of this object to determine which reservations to return. When searching for reservations, use the properties to fetch only the specified reservation, or use the property to fetch pending reservations only at the specified restaurant. If both of those properties are , retrieve all of the user’s currently pending reservations. After fetching the appropriate set of reservations, use the and properties to limit the set of results you return as part of your response. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should resolve and confirm any parameters and create an object using the found results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetUserCurrentRestaurantReservationBookingsIntent
type INGetUserCurrentRestaurantReservationBookingsIntent struct {
	INIntent
}

// INGetUserCurrentRestaurantReservationBookingsIntentFrom constructs a [INGetUserCurrentRestaurantReservationBookingsIntent] from an unsafe.Pointer.
//
// A request for the list of the user’s current reservations.
func INGetUserCurrentRestaurantReservationBookingsIntentFrom(ptr unsafe.Pointer) INGetUserCurrentRestaurantReservationBookingsIntent {
	return INGetUserCurrentRestaurantReservationBookingsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetUserCurrentRestaurantReservationBookingsIntentClass) Alloc() INGetUserCurrentRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetUserCurrentRestaurantReservationBookingsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetUserCurrentRestaurantReservationBookingsIntentClass) New() INGetUserCurrentRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetUserCurrentRestaurantReservationBookingsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) Init() INGetUserCurrentRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetUserCurrentRestaurantReservationBookingsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) Autorelease() INGetUserCurrentRestaurantReservationBookingsIntent {
	rv := objc.Send[INGetUserCurrentRestaurantReservationBookingsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetUserCurrentRestaurantReservationBookingsIntent creates a new INGetUserCurrentRestaurantReservationBookingsIntent instance.
func NewINGetUserCurrentRestaurantReservationBookingsIntent() INGetUserCurrentRestaurantReservationBookingsIntent {
	return getINGetUserCurrentRestaurantReservationBookingsIntentClass().New()
}


// An identifier to use when searching for the user’s reservations.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetUserCurrentRestaurantReservationBookingsIntent/reservationIdentifier
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) ReservationIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("reservationIdentifier"))
	return rv
}


// SetReservationIdentifier sets the value of the reservationIdentifier property.
// An identifier to use when searching for the user’s reservations.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetUserCurrentRestaurantReservationBookingsIntent/reservationIdentifier
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) SetReservationIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservationIdentifier:"), objc.String(value))
}

// A restaurant to use as a filter when searching for reservations.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetUserCurrentRestaurantReservationBookingsIntent/restaurant
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) Restaurant() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("restaurant"))
	return rv
}


// SetRestaurant sets the value of the restaurant property.
// A restaurant to use as a filter when searching for reservations.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetUserCurrentRestaurantReservationBookingsIntent/restaurant
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) SetRestaurant(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRestaurant:"), value)
}

// The earliest date to associate with any reservations.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetusercurrentrestaurantreservationbookingsintent/earliestbookingdateforresults
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) EarliestBookingDateForResults() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("earliestBookingDateForResults"))
	return rv
}


// SetEarliestBookingDateForResults sets the value of the earliestBookingDateForResults property.
// The earliest date to associate with any reservations.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetusercurrentrestaurantreservationbookingsintent/earliestbookingdateforresults
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) SetEarliestBookingDateForResults(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEarliestBookingDateForResults:"), value)
}

// The maximum number of reservations to include in your response object.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetusercurrentrestaurantreservationbookingsintent/maximumnumberofresults
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) MaximumNumberOfResults() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("maximumNumberOfResults"))
	return rv
}


// SetMaximumNumberOfResults sets the value of the maximumNumberOfResults property.
// The maximum number of reservations to include in your response object.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ingetusercurrentrestaurantreservationbookingsintent/maximumnumberofresults
func (i_ INGetUserCurrentRestaurantReservationBookingsIntent) SetMaximumNumberOfResults(value foundation.Number) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaximumNumberOfResults:"), value)
}



