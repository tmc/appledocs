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
	// properties:
	LodgingBusinessLocation() objc.IObject /* cross-framework: Placemark */
	SetLodgingBusinessLocation(value objc.IObject /* cross-framework: Placemark */)
	NumberOfAdults() int /* primitive/slice/pointer. */
	SetNumberOfAdults(value int /* primitive/slice/pointer. */)
	NumberOfChildren() int /* primitive/slice/pointer. */
	SetNumberOfChildren(value int /* primitive/slice/pointer. */)
	ReservationDuration() INDateComponentsRange /* already interface */
	SetReservationDuration(value INDateComponentsRange /* already interface */)
	// methods:
}

// The information that describes a lodging reservation.


// The information that describes a lodging reservation.
//
// [Full Topic]
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



// The name and location of the lodging establishment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/lodgingbusinesslocation
func (i_ INLodgingReservation) LodgingBusinessLocation() objc.IObject /* cross-framework: Placemark */ {
	rv := objc.Send[Placemark](i_.ID, objc.Sel("lodgingBusinessLocation"))
	return rv
}


// The name and location of the lodging establishment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/lodgingbusinesslocation
func (i_ INLodgingReservation) SetLodgingBusinessLocation(value objc.IObject /* cross-framework: Placemark */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLodgingBusinessLocation:"), value)
}


// The number of adults staying at the lodging location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/numberofadults-6fsnq
func (i_ INLodgingReservation) NumberOfAdults() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("numberOfAdults"))
	return rv
}


// The number of adults staying at the lodging location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/numberofadults-6fsnq
func (i_ INLodgingReservation) SetNumberOfAdults(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumberOfAdults:"), value)
}


// The number of children staying at the lodging location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/numberofchildren-1dm3g
func (i_ INLodgingReservation) NumberOfChildren() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("numberOfChildren"))
	return rv
}


// The number of children staying at the lodging location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/numberofchildren-1dm3g
func (i_ INLodgingReservation) SetNumberOfChildren(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumberOfChildren:"), value)
}


// The date and time range that indicates the beginning and end of the reservation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/reservationduration
func (i_ INLodgingReservation) ReservationDuration() INDateComponentsRange /* already interface */ {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("reservationDuration"))
	return rv
}


// The date and time range that indicates the beginning and end of the reservation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlodgingreservation/reservationduration
func (i_ INLodgingReservation) SetReservationDuration(value INDateComponentsRange /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReservationDuration:"), value)
}



