// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INBookRestaurantReservationIntent] class.
var (
	INBookRestaurantReservationIntentClass     _INBookRestaurantReservationIntentClass
	INBookRestaurantReservationIntentClassOnce sync.Once
)

func getINBookRestaurantReservationIntentClass() _INBookRestaurantReservationIntentClass {
	INBookRestaurantReservationIntentClassOnce.Do(func() {
		INBookRestaurantReservationIntentClass = _INBookRestaurantReservationIntentClass{objc.GetClass("INBookRestaurantReservationIntent")}
	})
	return INBookRestaurantReservationIntentClass
}

type _INBookRestaurantReservationIntentClass struct {
	class objc.Class
}

// An interface definition for the [INBookRestaurantReservationIntent] class.
type IINBookRestaurantReservationIntent interface {
	IINIntent
}

// A request to create a reservation at the specified restaurant.
//
// An object asks you to book the reservation time selected by the user. Maps sends this intent to your Intents extension when the user selects a reservation time from the available options. Use this intent object to obtain the details of the reservation, including the time slot and any selected special offers. Use those details to confirm the reservation with the restaurant and store the results in your system. Booking a reservation is the last step in the reservation creation process. By the time the system delivers this intent to your Intents extension, the user has already had an opportunity to view a list of potential reservation times and configure the details of the reservation request. This intent object contains all of the final choices made by the user. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should resolve and confirm any parameters and create an object with the status of the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INBookRestaurantReservationIntent
type INBookRestaurantReservationIntent struct {
	INIntent
}

// INBookRestaurantReservationIntentFrom constructs a [INBookRestaurantReservationIntent] from an unsafe.Pointer.
//
// A request to create a reservation at the specified restaurant.
func INBookRestaurantReservationIntentFrom(ptr unsafe.Pointer) INBookRestaurantReservationIntent {
	return INBookRestaurantReservationIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INBookRestaurantReservationIntentClass) Alloc() INBookRestaurantReservationIntent {
	rv := objc.Send[INBookRestaurantReservationIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INBookRestaurantReservationIntentClass) New() INBookRestaurantReservationIntent {
	rv := objc.Send[INBookRestaurantReservationIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INBookRestaurantReservationIntent) Init() INBookRestaurantReservationIntent {
	rv := objc.Send[INBookRestaurantReservationIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INBookRestaurantReservationIntent) Autorelease() INBookRestaurantReservationIntent {
	rv := objc.Send[INBookRestaurantReservationIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINBookRestaurantReservationIntent creates a new INBookRestaurantReservationIntent instance.
func NewINBookRestaurantReservationIntent() INBookRestaurantReservationIntent {
	return getINBookRestaurantReservationIntentClass().New()
}




