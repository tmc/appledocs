// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
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
	RestaurantLocation() corelocation.Placemark
	PartySize() int
	SetPartySize(value int)
	ReservationDuration() INDateComponentsRange
	SetReservationDuration(value INDateComponentsRange)
}

// The information that describes a restaurant reservation.


// The information that describes a restaurant reservation.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRestaurantReservation/initWithItemReference:reservationNumber:bookingTime:reservationStatus:reservationHolderName:actions:reservationDuration:partySize:restaurantLocation:
func NewINRestaurantReservationWithItemReferenceReservationNumberBookingTimeReservationStatusReservationHolderNameActionsReservationDurationPartySizeRestaurantLocation(itemReference INSpeakableString, reservationNumber string, bookingTime foundation.IDate, reservationStatus INReservationStatus, reservationHolderName string, actions []INReservationAction, reservationDuration INDateComponentsRange, partySize foundation.INumber, restaurantLocation corelocation.IPlacemark) INRestaurantReservation {
	instance := getINRestaurantReservationClass().Alloc()
	rv := objc.Send[INRestaurantReservation](instance.ID, objc.Sel("initWithItemReference:reservationNumber:bookingTime:reservationStatus:reservationHolderName:actions:reservationDuration:partySize:restaurantLocation:"), itemReference, objc.String(reservationNumber), bookingTime, reservationStatus, objc.String(reservationHolderName), actions, reservationDuration, partySize, restaurantLocation)
	rv.Autorelease()
	return rv
}



// The name and location of the restaurant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRestaurantReservation/restaurantLocation
func (i_ INRestaurantReservation) RestaurantLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("restaurantLocation"))
	return rv
}


// The number of people in the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrestaurantreservation/partysize-9ux0p
func (i_ INRestaurantReservation) PartySize() int {
	rv := objc.Send[int](i_.ID, objc.Sel("partySize"))
	return rv
}


// The number of people in the party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrestaurantreservation/partysize-9ux0p
func (i_ INRestaurantReservation) SetPartySize(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPartySize:"), value)
}


// The date and time range that defines beginning and end of the restaurant reservation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrestaurantreservation/reservationduration
func (i_ INRestaurantReservation) ReservationDuration() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("reservationDuration"))
	return rv
}


// The date and time range that defines beginning and end of the restaurant reservation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrestaurantreservation/reservationduration
func (i_ INRestaurantReservation) SetReservationDuration(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservationDuration:"), value)
}


