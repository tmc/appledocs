// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INLodgingReservation] class.
var (
	INLodgingReservationClass     _INLodgingReservationClass
	INLodgingReservationClassOnce sync.Once
)

func getINLodgingReservationClass() _INLodgingReservationClass {
	INLodgingReservationClassOnce.Do(func() {
		INLodgingReservationClass = _INLodgingReservationClass{objc.GetClass("INLodgingReservation")}
	})
	return INLodgingReservationClass
}

type _INLodgingReservationClass struct {
	class objc.Class
}

// An interface definition for the [INLodgingReservation] class.
type IINLodgingReservation interface {
	IINReservation
}

// The information that describes a lodging reservation.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INLodgingReservation
type INLodgingReservation struct {
	INReservation
}

// INLodgingReservationFrom constructs a [INLodgingReservation] from an unsafe.Pointer.
//
// The information that describes a lodging reservation.
func INLodgingReservationFrom(ptr unsafe.Pointer) INLodgingReservation {
	return INLodgingReservation{
		INReservation: INReservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INLodgingReservationClass) Alloc() INLodgingReservation {
	rv := objc.Send[INLodgingReservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INLodgingReservationClass) New() INLodgingReservation {
	rv := objc.Send[INLodgingReservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INLodgingReservation) Init() INLodgingReservation {
	rv := objc.Send[INLodgingReservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INLodgingReservation) Autorelease() INLodgingReservation {
	rv := objc.Send[INLodgingReservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINLodgingReservation creates a new INLodgingReservation instance.
func NewINLodgingReservation() INLodgingReservation {
	return getINLodgingReservationClass().New()
}




