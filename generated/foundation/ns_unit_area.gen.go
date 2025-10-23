// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitArea] class.
var (
	UnitAreaClass     _UnitAreaClass
	UnitAreaClassOnce sync.Once
)

func getUnitAreaClass() _UnitAreaClass {
	UnitAreaClassOnce.Do(func() {
		UnitAreaClass = _UnitAreaClass{objc.GetClass("NSUnitArea")}
	})
	return UnitAreaClass
}

type _UnitAreaClass struct {
	class objc.Class
}

// An interface definition for the [UnitArea] class.
type IUnitArea interface {
	IDimension
}

// A unit of measure for area.
//
// You typically use instances of to represent specific quantities of area using the class.


// A unit of measure for area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea
type UnitArea struct {
	Dimension
}

// UnitAreaFrom constructs a [UnitArea] from an unsafe.Pointer.
//
// A unit of measure for area.
func UnitAreaFrom(ptr unsafe.Pointer) UnitArea {
	return UnitArea{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitAreaClass) Alloc() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitAreaClass) New() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitArea) Init() UnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitArea) Autorelease() UnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitArea creates a new UnitArea instance.
func NewUnitArea() UnitArea {
	return getUnitAreaClass().New()
}



// The acres unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/acres
func (uc _UnitAreaClass) Acres() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("acres"))
	return rv
}

// The ares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/ares
func (uc _UnitAreaClass) Ares() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("ares"))
	return rv
}

// The hectares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/hectares
func (uc _UnitAreaClass) Hectares() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("hectares"))
	return rv
}

// The square centimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareCentimeters
func (uc _UnitAreaClass) SquareCentimeters() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareCentimeters"))
	return rv
}

// The square feet unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareFeet
func (uc _UnitAreaClass) SquareFeet() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareFeet"))
	return rv
}

// The square inches unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareInches
func (uc _UnitAreaClass) SquareInches() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareInches"))
	return rv
}

// The square kilometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareKilometers
func (uc _UnitAreaClass) SquareKilometers() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareKilometers"))
	return rv
}

// The square megameters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMegameters
func (uc _UnitAreaClass) SquareMegameters() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareMegameters"))
	return rv
}

// The square meters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMeters
func (uc _UnitAreaClass) SquareMeters() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareMeters"))
	return rv
}

// The square micrometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMicrometers
func (uc _UnitAreaClass) SquareMicrometers() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareMicrometers"))
	return rv
}

// The square miles unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMiles
func (uc _UnitAreaClass) SquareMiles() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareMiles"))
	return rv
}

// The square millimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMillimeters
func (uc _UnitAreaClass) SquareMillimeters() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareMillimeters"))
	return rv
}

// The square nanometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareNanometers
func (uc _UnitAreaClass) SquareNanometers() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareNanometers"))
	return rv
}

// The square yards unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareYards
func (uc _UnitAreaClass) SquareYards() UnitArea {
	rv := objc.Send[NSUnitArea](objc.ID(uc.class), objc.Sel("squareYards"))
	return rv
}

// The acres unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/acres
func (u_ UnitArea) Acres() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("acres"))
	return rv
}


// The ares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/ares
func (u_ UnitArea) Ares() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("ares"))
	return rv
}


// The hectares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/hectares
func (u_ UnitArea) Hectares() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("hectares"))
	return rv
}


// The square centimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareCentimeters
func (u_ UnitArea) SquareCentimeters() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareCentimeters"))
	return rv
}


// The square feet unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareFeet
func (u_ UnitArea) SquareFeet() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareFeet"))
	return rv
}


// The square inches unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareInches
func (u_ UnitArea) SquareInches() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareInches"))
	return rv
}


// The square kilometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareKilometers
func (u_ UnitArea) SquareKilometers() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareKilometers"))
	return rv
}


// The square megameters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMegameters
func (u_ UnitArea) SquareMegameters() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareMegameters"))
	return rv
}


// The square meters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMeters
func (u_ UnitArea) SquareMeters() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareMeters"))
	return rv
}


// The square micrometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMicrometers
func (u_ UnitArea) SquareMicrometers() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareMicrometers"))
	return rv
}


// The square miles unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMiles
func (u_ UnitArea) SquareMiles() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareMiles"))
	return rv
}


// The square millimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMillimeters
func (u_ UnitArea) SquareMillimeters() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareMillimeters"))
	return rv
}


// The square nanometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareNanometers
func (u_ UnitArea) SquareNanometers() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareNanometers"))
	return rv
}


// The square yards unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareYards
func (u_ UnitArea) SquareYards() NSUnitArea {
	rv := objc.Send[NSUnitArea](u_.ID, objc.Sel("squareYards"))
	return rv
}



