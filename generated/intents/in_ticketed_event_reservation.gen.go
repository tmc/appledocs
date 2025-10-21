// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INTicketedEventReservation] class.
var (
	INTicketedEventReservationClass     _INTicketedEventReservationClass
	INTicketedEventReservationClassOnce sync.Once
)

func getINTicketedEventReservationClass() _INTicketedEventReservationClass {
	INTicketedEventReservationClassOnce.Do(func() {
		INTicketedEventReservationClass = _INTicketedEventReservationClass{objc.GetClass("INTicketedEventReservation")}
	})
	return INTicketedEventReservationClass
}

type _INTicketedEventReservationClass struct {
	class objc.Class
}

// An interface definition for the [INTicketedEventReservation] class.
type IINTicketedEventReservation interface {
	IINReservation
}

// The information that describes a ticketed event reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INTicketedEventReservation
type INTicketedEventReservation struct {
	INReservation
}

// INTicketedEventReservationFrom constructs a [INTicketedEventReservation] from an unsafe.Pointer.
//
// The information that describes a ticketed event reservation.
func INTicketedEventReservationFrom(ptr unsafe.Pointer) INTicketedEventReservation {
	return INTicketedEventReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INTicketedEventReservationClass) Alloc() INTicketedEventReservation {
	rv := objc.Send[INTicketedEventReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INTicketedEventReservationClass) New() INTicketedEventReservation {
	rv := objc.Send[INTicketedEventReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INTicketedEventReservation) Init() INTicketedEventReservation {
	rv := objc.Send[INTicketedEventReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INTicketedEventReservation) Autorelease() INTicketedEventReservation {
	rv := objc.Send[INTicketedEventReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINTicketedEventReservation creates a new INTicketedEventReservation instance.
func NewINTicketedEventReservation() INTicketedEventReservation {
	return getINTicketedEventReservationClass().New()
}


// An object containing detailed information about the ticketed event.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INTicketedEventReservation/event
func (i_ INTicketedEventReservation) Event() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("event"))
	return rv
}

// The user’s assigned seat for the ticketed event.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inticketedeventreservation/reservedseat
func (i_ INTicketedEventReservation) ReservedSeat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reservedSeat"))
	return rv
}


// SetReservedSeat sets the value of the reservedSeat property.
// The user’s assigned seat for the ticketed event.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inticketedeventreservation/reservedseat
func (i_ INTicketedEventReservation) SetReservedSeat(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservedSeat:"), value)
}



