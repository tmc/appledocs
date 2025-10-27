// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [UnitLength] class.
var (
	UnitLengthClass     _UnitLengthClass
	UnitLengthClassOnce sync.Once
)

func getUnitLengthClass() _UnitLengthClass {
	UnitLengthClassOnce.Do(func() {
		UnitLengthClass = _UnitLengthClass{objc.GetClass("NSUnitLength")}
	})
	return UnitLengthClass
}

type _UnitLengthClass struct {
	class objc.Class
}





// An interface definition for the [UnitLength] class.
type IUnitLength interface {
	IDimension
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (uc _UnitLengthClass) Alloc() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitLengthClass) New() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitLength) Init() UnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitLength) Autorelease() UnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitLength creates a new UnitLength instance.
func NewUnitLength() UnitLength {
	return getUnitLengthClass().New()
}





// A unit of measure for length.
//
// You typically use instances of to represent specific quantities of length using the class.


// A unit of measure for length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength
type UnitLength struct {
	Dimension
}

// UnitLengthFrom constructs a [UnitLength] from an unsafe.Pointer.
//
// A unit of measure for length.
func UnitLengthFrom(ptr unsafe.Pointer) UnitLength {
	return UnitLength{
		Dimension: DimensionFrom(ptr),
	}
}















// The centimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/centimeters
func (uc _UnitLengthClass) Centimeters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("centimeters"))
	return rv
}

// The decameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decameters
func (uc _UnitLengthClass) Decameters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("decameters"))
	return rv
}

// The decimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decimeters
func (uc _UnitLengthClass) Decimeters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("decimeters"))
	return rv
}

// The feet unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/feet
func (uc _UnitLengthClass) Feet() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("feet"))
	return rv
}

// The hectometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (uc _UnitLengthClass) Hectometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("hectometers"))
	return rv
}

// The inches unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/inches
func (uc _UnitLengthClass) Inches() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("inches"))
	return rv
}

// The kilometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/kilometers
func (uc _UnitLengthClass) Kilometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("kilometers"))
	return rv
}

// The megameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/megameters
func (uc _UnitLengthClass) Megameters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("megameters"))
	return rv
}

// The meters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/meters
func (uc _UnitLengthClass) Meters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("meters"))
	return rv
}

// The micrometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/micrometers
func (uc _UnitLengthClass) Micrometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("micrometers"))
	return rv
}

// The miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/miles
func (uc _UnitLengthClass) Miles() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("miles"))
	return rv
}

// The millimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/millimeters
func (uc _UnitLengthClass) Millimeters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("millimeters"))
	return rv
}

// The nanometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/nanometers
func (uc _UnitLengthClass) Nanometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("nanometers"))
	return rv
}

// The picometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/picometers
func (uc _UnitLengthClass) Picometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("picometers"))
	return rv
}

// The Scandinavian miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/scandinavianMiles
func (uc _UnitLengthClass) ScandinavianMiles() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("scandinavianMiles"))
	return rv
}

// The yards unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/yards
func (uc _UnitLengthClass) Yards() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("yards"))
	return rv
}











// The centimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/centimeters
func (u_ UnitLength) Centimeters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("centimeters"))
	return rv
}


// The decameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decameters
func (u_ UnitLength) Decameters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("decameters"))
	return rv
}


// The decimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decimeters
func (u_ UnitLength) Decimeters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("decimeters"))
	return rv
}


// The feet unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/feet
func (u_ UnitLength) Feet() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("feet"))
	return rv
}


// The hectometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (u_ UnitLength) Hectometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("hectometers"))
	return rv
}


// The inches unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/inches
func (u_ UnitLength) Inches() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("inches"))
	return rv
}


// The kilometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/kilometers
func (u_ UnitLength) Kilometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("kilometers"))
	return rv
}


// The megameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/megameters
func (u_ UnitLength) Megameters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("megameters"))
	return rv
}


// The meters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/meters
func (u_ UnitLength) Meters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("meters"))
	return rv
}


// The micrometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/micrometers
func (u_ UnitLength) Micrometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("micrometers"))
	return rv
}


// The miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/miles
func (u_ UnitLength) Miles() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("miles"))
	return rv
}


// The millimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/millimeters
func (u_ UnitLength) Millimeters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("millimeters"))
	return rv
}


// The nanometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/nanometers
func (u_ UnitLength) Nanometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("nanometers"))
	return rv
}


// The picometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/picometers
func (u_ UnitLength) Picometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("picometers"))
	return rv
}


// The Scandinavian miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/scandinavianMiles
func (u_ UnitLength) ScandinavianMiles() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("scandinavianMiles"))
	return rv
}


// The yards unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/yards
func (u_ UnitLength) Yards() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("yards"))
	return rv
}








