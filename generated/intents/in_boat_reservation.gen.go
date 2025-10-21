// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INBoatReservation] class.
var (
	INBoatReservationClass     _INBoatReservationClass
	INBoatReservationClassOnce sync.Once
)

func getINBoatReservationClass() _INBoatReservationClass {
	INBoatReservationClassOnce.Do(func() {
		INBoatReservationClass = _INBoatReservationClass{objc.GetClass("INBoatReservation")}
	})
	return INBoatReservationClass
}

type _INBoatReservationClass struct {
	class objc.Class
}

// An interface definition for the [INBoatReservation] class.
type IINBoatReservation interface {
	IINReservation
}

// The information that describes a boat reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INBoatReservation
type INBoatReservation struct {
	INReservation
}

// INBoatReservationFrom constructs a [INBoatReservation] from an unsafe.Pointer.
//
// The information that describes a boat reservation.
func INBoatReservationFrom(ptr unsafe.Pointer) INBoatReservation {
	return INBoatReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INBoatReservationClass) Alloc() INBoatReservation {
	rv := objc.Send[INBoatReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INBoatReservationClass) New() INBoatReservation {
	rv := objc.Send[INBoatReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INBoatReservation) Init() INBoatReservation {
	rv := objc.Send[INBoatReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INBoatReservation) Autorelease() INBoatReservation {
	rv := objc.Send[INBoatReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINBoatReservation creates a new INBoatReservation instance.
func NewINBoatReservation() INBoatReservation {
	return getINBoatReservationClass().New()
}




