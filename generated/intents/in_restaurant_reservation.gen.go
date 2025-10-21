// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INRestaurantReservation] class.
var (
	INRestaurantReservationClass     _INRestaurantReservationClass
	INRestaurantReservationClassOnce sync.Once
)

func getINRestaurantReservationClass() _INRestaurantReservationClass {
	INRestaurantReservationClassOnce.Do(func() {
		INRestaurantReservationClass = _INRestaurantReservationClass{objc.GetClass("INRestaurantReservation")}
	})
	return INRestaurantReservationClass
}

type _INRestaurantReservationClass struct {
	class objc.Class
}

// An interface definition for the [INRestaurantReservation] class.
type IINRestaurantReservation interface {
	IINReservation
}

// The information that describes a restaurant reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRestaurantReservation
type INRestaurantReservation struct {
	INReservation
}

// INRestaurantReservationFrom constructs a [INRestaurantReservation] from an unsafe.Pointer.
//
// The information that describes a restaurant reservation.
func INRestaurantReservationFrom(ptr unsafe.Pointer) INRestaurantReservation {
	return INRestaurantReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INRestaurantReservationClass) Alloc() INRestaurantReservation {
	rv := objc.Send[INRestaurantReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRestaurantReservationClass) New() INRestaurantReservation {
	rv := objc.Send[INRestaurantReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRestaurantReservation) Init() INRestaurantReservation {
	rv := objc.Send[INRestaurantReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRestaurantReservation) Autorelease() INRestaurantReservation {
	rv := objc.Send[INRestaurantReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRestaurantReservation creates a new INRestaurantReservation instance.
func NewINRestaurantReservation() INRestaurantReservation {
	return getINRestaurantReservationClass().New()
}




// Creates a new restaurant reservation with the provided information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRestaurantReservation/initWithItemReference:reservationNumber:bookingTime:reservationStatus:reservationHolderName:actions:reservationDuration:partySize:restaurantLocation:
func NewINRestaurantReservationWithItemReferenceReservationNumberBookingTimeReservationStatusReservationHolderNameActionsReservationDurationPartySizeRestaurantLocation(itemReference unsafe.Pointer, reservationNumber string, bookingTime unsafe.Pointer, reservationStatus unsafe.Pointer, reservationHolderName string, actions unsafe.Pointer, reservationDuration unsafe.Pointer, partySize unsafe.Pointer, restaurantLocation unsafe.Pointer) INRestaurantReservation {
	instance := getINRestaurantReservationClass().Alloc()
	rv := objc.Send[INRestaurantReservation](instance.ID, objc.Sel("initWithItemReference:reservationNumber:bookingTime:reservationStatus:reservationHolderName:actions:reservationDuration:partySize:restaurantLocation:"), itemReference, objc.String(reservationNumber), bookingTime, reservationStatus, objc.String(reservationHolderName), actions, reservationDuration, partySize, restaurantLocation)
	rv.Autorelease()
	return rv
}


// The name and location of the restaurant.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRestaurantReservation/restaurantLocation
func (i_ INRestaurantReservation) RestaurantLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("restaurantLocation"))
	return rv
}


