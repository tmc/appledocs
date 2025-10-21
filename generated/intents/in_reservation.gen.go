// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INReservation] class.
var (
	INReservationClass     _INReservationClass
	INReservationClassOnce sync.Once
)

func getINReservationClass() _INReservationClass {
	INReservationClassOnce.Do(func() {
		INReservationClass = _INReservationClass{objc.GetClass("INReservation")}
	})
	return INReservationClass
}

type _INReservationClass struct {
	class objc.Class
}

// An interface definition for the [INReservation] class.
type IINReservation interface {
	objectivec.IObject
}

// An object that describes a reservation.
//
// Don’t create instances of this class directly. Instead, use the subclass associated with the type of reservation created.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation
type INReservation struct {
	objectivec.Object
}

// INReservationFrom constructs a [INReservation] from an unsafe.Pointer.
//
// An object that describes a reservation.
func INReservationFrom(ptr unsafe.Pointer) INReservation {
	return INReservation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INReservationClass) Alloc() INReservation {
	rv := objc.Send[INReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INReservationClass) New() INReservation {
	rv := objc.Send[INReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INReservation) Init() INReservation {
	rv := objc.Send[INReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INReservation) Autorelease() INReservation {
	rv := objc.Send[INReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINReservation creates a new INReservation instance.
func NewINReservation() INReservation {
	return getINReservationClass().New()
}


// An array containing actions the user can perform on the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/actions
func (i_ INReservation) Actions() []INReservationAction {
	rv := objc.Send[[]INReservationAction](i_.ID, objc.Sel("actions"))
	return rv
}

// The date and time the user booked the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/bookingTime
func (i_ INReservation) BookingTime() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](i_.ID, objc.Sel("bookingTime"))
	return rv
}

// A unique reference for the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/itemReference
func (i_ INReservation) ItemReference() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("itemReference"))
	return rv
}

// The reservation holder’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/reservationHolderName
func (i_ INReservation) ReservationHolderName() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("reservationHolderName"))
	return rv
}

// The reservation number.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/reservationNumber
func (i_ INReservation) ReservationNumber() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("reservationNumber"))
	return rv
}

// The current status of the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/reservationStatus
func (i_ INReservation) ReservationStatus() INReservationStatus {
	rv := objc.Send[INReservationStatus](i_.ID, objc.Sel("reservationStatus"))
	return rv
}

// A webpage the user can access to view reservation details.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INReservation/url
func (i_ INReservation) URL() foundation.URL {
	rv := objc.Send[foundation.URL](i_.ID, objc.Sel("URL"))
	return rv
}



