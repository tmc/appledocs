// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


// The date and time of the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/bookingdatecomponents
func (i_ INBookRestaurantReservationIntent) BookingDateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](i_.ID, objc.Sel("bookingDateComponents"))
	return rv
}


// SetBookingDateComponents sets the value of the bookingDateComponents property.
// The date and time of the reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/bookingdatecomponents
func (i_ INBookRestaurantReservationIntent) SetBookingDateComponents(value foundation.IDateComponents) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBookingDateComponents:"), value)
}

// The unique identifier associated with the initial reservation data.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/bookingidentifier
func (i_ INBookRestaurantReservationIntent) BookingIdentifier() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("bookingIdentifier"))
	return rv
}


// SetBookingIdentifier sets the value of the bookingIdentifier property.
// The unique identifier associated with the initial reservation data.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/bookingidentifier
func (i_ INBookRestaurantReservationIntent) SetBookingIdentifier(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBookingIdentifier:"), value)
}

// The identity of the guest associated with the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/guest
func (i_ INBookRestaurantReservationIntent) Guest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("guest"))
	return rv
}


// SetGuest sets the value of the guest property.
// The identity of the guest associated with the reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/guest
func (i_ INBookRestaurantReservationIntent) SetGuest(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGuest:"), value)
}

// Information about any special requests made by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/guestprovidedspecialrequesttext
func (i_ INBookRestaurantReservationIntent) GuestProvidedSpecialRequestText() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("guestProvidedSpecialRequestText"))
	return rv
}


// SetGuestProvidedSpecialRequestText sets the value of the guestProvidedSpecialRequestText property.
// Information about any special requests made by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/guestprovidedspecialrequesttext
func (i_ INBookRestaurantReservationIntent) SetGuestProvidedSpecialRequestText(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGuestProvidedSpecialRequestText:"), value)
}

// The total number of people in the user’s party.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/partysize
func (i_ INBookRestaurantReservationIntent) PartySize() int {
	rv := objc.Send[int](i_.ID, objc.Sel("partySize"))
	return rv
}


// SetPartySize sets the value of the partySize property.
// The total number of people in the user’s party.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/partysize
func (i_ INBookRestaurantReservationIntent) SetPartySize(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPartySize:"), value)
}

// The restaurant to contact regarding the booking.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/restaurant
func (i_ INBookRestaurantReservationIntent) Restaurant() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("restaurant"))
	return rv
}


// SetRestaurant sets the value of the restaurant property.
// The restaurant to contact regarding the booking.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/restaurant
func (i_ INBookRestaurantReservationIntent) SetRestaurant(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRestaurant:"), value)
}

// The special offer, if any, selected by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/selectedoffer
func (i_ INBookRestaurantReservationIntent) SelectedOffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("selectedOffer"))
	return rv
}


// SetSelectedOffer sets the value of the selectedOffer property.
// The special offer, if any, selected by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbookrestaurantreservationintent/selectedoffer
func (i_ INBookRestaurantReservationIntent) SetSelectedOffer(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectedOffer:"), value)
}



