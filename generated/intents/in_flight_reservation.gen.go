// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INFlightReservation] class.
var (
	INFlightReservationClass     _INFlightReservationClass
	INFlightReservationClassOnce sync.Once
)

func getINFlightReservationClass() _INFlightReservationClass {
	INFlightReservationClassOnce.Do(func() {
		INFlightReservationClass = _INFlightReservationClass{objc.GetClass("INFlightReservation")}
	})
	return INFlightReservationClass
}

type _INFlightReservationClass struct {
	class objc.Class
}

// An interface definition for the [INFlightReservation] class.
type IINFlightReservation interface {
	IINReservation
}

// The information that describes a flight reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INFlightReservation
type INFlightReservation struct {
	INReservation
}

// INFlightReservationFrom constructs a [INFlightReservation] from an unsafe.Pointer.
//
// The information that describes a flight reservation.
func INFlightReservationFrom(ptr unsafe.Pointer) INFlightReservation {
	return INFlightReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INFlightReservationClass) Alloc() INFlightReservation {
	rv := objc.Send[INFlightReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INFlightReservationClass) New() INFlightReservation {
	rv := objc.Send[INFlightReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INFlightReservation) Init() INFlightReservation {
	rv := objc.Send[INFlightReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INFlightReservation) Autorelease() INFlightReservation {
	rv := objc.Send[INFlightReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINFlightReservation creates a new INFlightReservation instance.
func NewINFlightReservation() INFlightReservation {
	return getINFlightReservationClass().New()
}


// The flight information associated with the reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inflightreservation/flight
func (i_ INFlightReservation) Flight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("flight"))
	return rv
}


// SetFlight sets the value of the flight property.
// The flight information associated with the reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inflightreservation/flight
func (i_ INFlightReservation) SetFlight(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFlight:"), value)
}

// The user’s seat for the flight.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inflightreservation/reservedseat
func (i_ INFlightReservation) ReservedSeat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reservedSeat"))
	return rv
}


// SetReservedSeat sets the value of the reservedSeat property.
// The user’s seat for the flight.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inflightreservation/reservedseat
func (i_ INFlightReservation) SetReservedSeat(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservedSeat:"), value)
}



