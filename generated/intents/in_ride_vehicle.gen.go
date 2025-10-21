// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INRideVehicle] class.
var (
	INRideVehicleClass     _INRideVehicleClass
	INRideVehicleClassOnce sync.Once
)

func getINRideVehicleClass() _INRideVehicleClass {
	INRideVehicleClassOnce.Do(func() {
		INRideVehicleClass = _INRideVehicleClass{objc.GetClass("INRideVehicle")}
	})
	return INRideVehicleClass
}

type _INRideVehicleClass struct {
	class objc.Class
}

// An interface definition for the [INRideVehicle] class.
type IINRideVehicle interface {
	objectivec.IObject
}

// A specific vehicle used by a ride-booking service.
//
// An object provides information about a vehicle in your fleet. Use a vehicle object to convey details about a specific vehicle to the user. Siri and Maps present information about your vehicle’s location on the map and let the user know when the vehicle arrives. You create an object when the user books a ride or when you provide the current status of a ride. You assign the vehicle object to the property of the object that you include with your response.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle
type INRideVehicle struct {
	objectivec.Object
}

// INRideVehicleFrom constructs a [INRideVehicle] from an unsafe.Pointer.
//
// A specific vehicle used by a ride-booking service.
func INRideVehicleFrom(ptr unsafe.Pointer) INRideVehicle {
	return INRideVehicle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INRideVehicleClass) Alloc() INRideVehicle {
	rv := objc.Send[INRideVehicle](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRideVehicleClass) New() INRideVehicle {
	rv := objc.Send[INRideVehicle](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRideVehicle) Init() INRideVehicle {
	rv := objc.Send[INRideVehicle](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRideVehicle) Autorelease() INRideVehicle {
	rv := objc.Send[INRideVehicle](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRideVehicle creates a new INRideVehicle instance.
func NewINRideVehicle() INRideVehicle {
	return getINRideVehicleClass().New()
}


// The vehicle assigned to pick up the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/vehicle
func (i_ INRideVehicle) Vehicle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("vehicle"))
	return rv
}


// SetVehicle sets the value of the vehicle property.
// The vehicle assigned to pick up the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inridestatus/vehicle
func (i_ INRideVehicle) SetVehicle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVehicle:"), value)
}

// The most recent location of the vehicle.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/location
func (i_ INRideVehicle) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The most recent location of the vehicle.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/location
func (i_ INRideVehicle) SetLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocation:"), value)
}

// The name of the vehicle’s manufacturer.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/manufacturer
func (i_ INRideVehicle) Manufacturer() string {
	rv := objc.Send[string](i_.ID, objc.Sel("manufacturer"))
	return rv
}


// SetManufacturer sets the value of the manufacturer property.
// The name of the vehicle’s manufacturer.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/manufacturer
func (i_ INRideVehicle) SetManufacturer(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setManufacturer:"), objc.String(value))
}

// The image to use for the vehicle when displaying its position on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/mapAnnotationImage
func (i_ INRideVehicle) MapAnnotationImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mapAnnotationImage"))
	return rv
}


// SetMapAnnotationImage sets the value of the mapAnnotationImage property.
// The image to use for the vehicle when displaying its position on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/mapAnnotationImage
func (i_ INRideVehicle) SetMapAnnotationImage(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMapAnnotationImage:"), value)
}

// The model of the vehicle.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/model
func (i_ INRideVehicle) Model() string {
	rv := objc.Send[string](i_.ID, objc.Sel("model"))
	return rv
}


// SetModel sets the value of the model property.
// The model of the vehicle.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/model
func (i_ INRideVehicle) SetModel(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setModel:"), objc.String(value))
}

// The text on the license plate or registration plate of the vehicle.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/registrationPlate
func (i_ INRideVehicle) RegistrationPlate() string {
	rv := objc.Send[string](i_.ID, objc.Sel("registrationPlate"))
	return rv
}


// SetRegistrationPlate sets the value of the registrationPlate property.
// The text on the license plate or registration plate of the vehicle.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRideVehicle/registrationPlate
func (i_ INRideVehicle) SetRegistrationPlate(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegistrationPlate:"), objc.String(value))
}



