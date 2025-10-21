// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INBusReservation] class.
var (
	INBusReservationClass     _INBusReservationClass
	INBusReservationClassOnce sync.Once
)

func getINBusReservationClass() _INBusReservationClass {
	INBusReservationClassOnce.Do(func() {
		INBusReservationClass = _INBusReservationClass{objc.GetClass("INBusReservation")}
	})
	return INBusReservationClass
}

type _INBusReservationClass struct {
	class objc.Class
}

// An interface definition for the [INBusReservation] class.
type IINBusReservation interface {
	IINReservation
}

// The information that describes a bus reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INBusReservation
type INBusReservation struct {
	INReservation
}

// INBusReservationFrom constructs a [INBusReservation] from an unsafe.Pointer.
//
// The information that describes a bus reservation.
func INBusReservationFrom(ptr unsafe.Pointer) INBusReservation {
	return INBusReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INBusReservationClass) Alloc() INBusReservation {
	rv := objc.Send[INBusReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INBusReservationClass) New() INBusReservation {
	rv := objc.Send[INBusReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INBusReservation) Init() INBusReservation {
	rv := objc.Send[INBusReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INBusReservation) Autorelease() INBusReservation {
	rv := objc.Send[INBusReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINBusReservation creates a new INBusReservation instance.
func NewINBusReservation() INBusReservation {
	return getINBusReservationClass().New()
}


// Information describing a bus trip.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbusreservation/bustrip
func (i_ INBusReservation) BusTrip() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("busTrip"))
	return rv
}


// SetBusTrip sets the value of the busTrip property.
// Information describing a bus trip.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbusreservation/bustrip
func (i_ INBusReservation) SetBusTrip(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBusTrip:"), value)
}

// Seat information for the bus reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbusreservation/reservedseat
func (i_ INBusReservation) ReservedSeat() INSeat {
	rv := objc.Send[INSeat](i_.ID, objc.Sel("reservedSeat"))
	return rv
}


// SetReservedSeat sets the value of the reservedSeat property.
// Seat information for the bus reservation.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inbusreservation/reservedseat
func (i_ INBusReservation) SetReservedSeat(value INSeat) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservedSeat:"), value)
}



