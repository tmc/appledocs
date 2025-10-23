// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitPressure] class.
var (
	UnitPressureClass     _UnitPressureClass
	UnitPressureClassOnce sync.Once
)

func getUnitPressureClass() _UnitPressureClass {
	UnitPressureClassOnce.Do(func() {
		UnitPressureClass = _UnitPressureClass{objc.GetClass("NSUnitPressure")}
	})
	return UnitPressureClass
}

type _UnitPressureClass struct {
	class objc.Class
}

// An interface definition for the [UnitPressure] class.
type IUnitPressure interface {
	IDimension
}

// A unit of measure for pressure.
//
// You typically use instances of to represent specific quantities of pressure using the class.


// A unit of measure for pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure
type UnitPressure struct {
	Dimension
}

// UnitPressureFrom constructs a [UnitPressure] from an unsafe.Pointer.
//
// A unit of measure for pressure.
func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitPressureClass) Alloc() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitPressureClass) New() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitPressure) Init() UnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitPressure) Autorelease() UnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitPressure creates a new UnitPressure instance.
func NewUnitPressure() UnitPressure {
	return getUnitPressureClass().New()
}



// The bars unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/bars
func (uc _UnitPressureClass) Bars() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("bars"))
	return rv
}

// The gigapascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/gigapascals
func (uc _UnitPressureClass) Gigapascals() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("gigapascals"))
	return rv
}

// The hectopascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/hectopascals
func (uc _UnitPressureClass) Hectopascals() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("hectopascals"))
	return rv
}

// The inches of mercury unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/inchesOfMercury
func (uc _UnitPressureClass) InchesOfMercury() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("inchesOfMercury"))
	return rv
}

// The kilopascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/kilopascals
func (uc _UnitPressureClass) Kilopascals() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("kilopascals"))
	return rv
}

// The megapascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/megapascals
func (uc _UnitPressureClass) Megapascals() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("megapascals"))
	return rv
}

// The millibars unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/millibars
func (uc _UnitPressureClass) Millibars() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("millibars"))
	return rv
}

// The millimeters of mercury unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/millimetersOfMercury
func (uc _UnitPressureClass) MillimetersOfMercury() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("millimetersOfMercury"))
	return rv
}

// The newtons per square meter unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/newtonsPerMetersSquared
func (uc _UnitPressureClass) NewtonsPerMetersSquared() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("newtonsPerMetersSquared"))
	return rv
}

// The pounds per square inch unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/poundsForcePerSquareInch
func (uc _UnitPressureClass) PoundsForcePerSquareInch() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("poundsForcePerSquareInch"))
	return rv
}

// The bars unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/bars
func (u_ UnitPressure) Bars() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("bars"))
	return rv
}


// The gigapascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/gigapascals
func (u_ UnitPressure) Gigapascals() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("gigapascals"))
	return rv
}


// The hectopascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/hectopascals
func (u_ UnitPressure) Hectopascals() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("hectopascals"))
	return rv
}


// The inches of mercury unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/inchesOfMercury
func (u_ UnitPressure) InchesOfMercury() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("inchesOfMercury"))
	return rv
}


// The kilopascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/kilopascals
func (u_ UnitPressure) Kilopascals() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("kilopascals"))
	return rv
}


// The megapascals unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/megapascals
func (u_ UnitPressure) Megapascals() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("megapascals"))
	return rv
}


// The millibars unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/millibars
func (u_ UnitPressure) Millibars() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("millibars"))
	return rv
}


// The millimeters of mercury unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/millimetersOfMercury
func (u_ UnitPressure) MillimetersOfMercury() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("millimetersOfMercury"))
	return rv
}


// The newtons per square meter unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/newtonsPerMetersSquared
func (u_ UnitPressure) NewtonsPerMetersSquared() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("newtonsPerMetersSquared"))
	return rv
}


// The pounds per square inch unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/poundsForcePerSquareInch
func (u_ UnitPressure) PoundsForcePerSquareInch() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("poundsForcePerSquareInch"))
	return rv
}



