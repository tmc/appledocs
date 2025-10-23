// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
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
	DropOffLocation() corelocation.Placemark
	SetDropOffLocation(value corelocation.IPlacemark)
	PickupLocation() corelocation.Placemark
	SetPickupLocation(value corelocation.IPlacemark)
	RentalCar() unsafe.Pointer
	SetRentalCar(value unsafe.Pointer)
	RentalDuration() INDateComponentsRange
	SetRentalDuration(value INDateComponentsRange)
}

// The information that describes a rental car reservation.


// The information that describes a rental car reservation.
//
// [Full Topic]
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



// The name and location where the user can drop off the car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/dropofflocation
func (i_ INRentalCarReservation) DropOffLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("dropOffLocation"))
	return rv
}


// The name and location where the user can drop off the car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/dropofflocation
func (i_ INRentalCarReservation) SetDropOffLocation(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDropOffLocation:"), value)
}


// The name and location where the user can pick up the car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/pickuplocation
func (i_ INRentalCarReservation) PickupLocation() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("pickupLocation"))
	return rv
}


// The name and location where the user can pick up the car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/pickuplocation
func (i_ INRentalCarReservation) SetPickupLocation(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPickupLocation:"), value)
}


// An object containing detailed information about the rental car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/rentalcar
func (i_ INRentalCarReservation) RentalCar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("rentalCar"))
	return rv
}


// An object containing detailed information about the rental car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/rentalcar
func (i_ INRentalCarReservation) SetRentalCar(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRentalCar:"), value)
}


// The date and time range that indicates the pickup and drop off times for the rental.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/rentalduration
func (i_ INRentalCarReservation) RentalDuration() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("rentalDuration"))
	return rv
}


// The date and time range that indicates the pickup and drop off times for the rental.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inrentalcarreservation/rentalduration
func (i_ INRentalCarReservation) SetRentalDuration(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRentalDuration:"), value)
}



