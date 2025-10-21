// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INRentalCarReservation] class.
var (
	INRentalCarReservationClass     _INRentalCarReservationClass
	INRentalCarReservationClassOnce sync.Once
)

func getINRentalCarReservationClass() _INRentalCarReservationClass {
	INRentalCarReservationClassOnce.Do(func() {
		INRentalCarReservationClass = _INRentalCarReservationClass{objc.GetClass("INRentalCarReservation")}
	})
	return INRentalCarReservationClass
}

type _INRentalCarReservationClass struct {
	class objc.Class
}

// An interface definition for the [INRentalCarReservation] class.
type IINRentalCarReservation interface {
	IINReservation
}

// The information that describes a rental car reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRentalCarReservation
type INRentalCarReservation struct {
	INReservation
}

// INRentalCarReservationFrom constructs a [INRentalCarReservation] from an unsafe.Pointer.
//
// The information that describes a rental car reservation.
func INRentalCarReservationFrom(ptr unsafe.Pointer) INRentalCarReservation {
	return INRentalCarReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INRentalCarReservationClass) Alloc() INRentalCarReservation {
	rv := objc.Send[INRentalCarReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRentalCarReservationClass) New() INRentalCarReservation {
	rv := objc.Send[INRentalCarReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRentalCarReservation) Init() INRentalCarReservation {
	rv := objc.Send[INRentalCarReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRentalCarReservation) Autorelease() INRentalCarReservation {
	rv := objc.Send[INRentalCarReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRentalCarReservation creates a new INRentalCarReservation instance.
func NewINRentalCarReservation() INRentalCarReservation {
	return getINRentalCarReservationClass().New()
}




