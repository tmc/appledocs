// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




