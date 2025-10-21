// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetAvailableRestaurantReservationBookingDefaultsIntent] class.
var (
	INGetAvailableRestaurantReservationBookingDefaultsIntentClass     _INGetAvailableRestaurantReservationBookingDefaultsIntentClass
	INGetAvailableRestaurantReservationBookingDefaultsIntentClassOnce sync.Once
)

func getINGetAvailableRestaurantReservationBookingDefaultsIntentClass() _INGetAvailableRestaurantReservationBookingDefaultsIntentClass {
	INGetAvailableRestaurantReservationBookingDefaultsIntentClassOnce.Do(func() {
		INGetAvailableRestaurantReservationBookingDefaultsIntentClass = _INGetAvailableRestaurantReservationBookingDefaultsIntentClass{objc.GetClass("INGetAvailableRestaurantReservationBookingDefaultsIntent")}
	})
	return INGetAvailableRestaurantReservationBookingDefaultsIntentClass
}

type _INGetAvailableRestaurantReservationBookingDefaultsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetAvailableRestaurantReservationBookingDefaultsIntent] class.
type IINGetAvailableRestaurantReservationBookingDefaultsIntent interface {
	IINIntent
}

// A request for the default values to use when fetching potential reservation options.
//
// An object asks you to provide the set of default options to use when fetching possible reservation times for the specified restaurant. Because restaurants may have different requirements for booking reservations, this intent lets you provide a set of reasonable default values that reflect any restaurant-specific requirements or user tendencies. For example, you use this intent to return the minimum or maximum party size supported by the restaurant. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should resolve and confirm any parameters and create an object with the found results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetAvailableRestaurantReservationBookingDefaultsIntent
type INGetAvailableRestaurantReservationBookingDefaultsIntent struct {
	INIntent
}

// INGetAvailableRestaurantReservationBookingDefaultsIntentFrom constructs a [INGetAvailableRestaurantReservationBookingDefaultsIntent] from an unsafe.Pointer.
//
// A request for the default values to use when fetching potential reservation options.
func INGetAvailableRestaurantReservationBookingDefaultsIntentFrom(ptr unsafe.Pointer) INGetAvailableRestaurantReservationBookingDefaultsIntent {
	return INGetAvailableRestaurantReservationBookingDefaultsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetAvailableRestaurantReservationBookingDefaultsIntentClass) Alloc() INGetAvailableRestaurantReservationBookingDefaultsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingDefaultsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetAvailableRestaurantReservationBookingDefaultsIntentClass) New() INGetAvailableRestaurantReservationBookingDefaultsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingDefaultsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetAvailableRestaurantReservationBookingDefaultsIntent) Init() INGetAvailableRestaurantReservationBookingDefaultsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingDefaultsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetAvailableRestaurantReservationBookingDefaultsIntent) Autorelease() INGetAvailableRestaurantReservationBookingDefaultsIntent {
	rv := objc.Send[INGetAvailableRestaurantReservationBookingDefaultsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetAvailableRestaurantReservationBookingDefaultsIntent creates a new INGetAvailableRestaurantReservationBookingDefaultsIntent instance.
func NewINGetAvailableRestaurantReservationBookingDefaultsIntent() INGetAvailableRestaurantReservationBookingDefaultsIntent {
	return getINGetAvailableRestaurantReservationBookingDefaultsIntentClass().New()
}




