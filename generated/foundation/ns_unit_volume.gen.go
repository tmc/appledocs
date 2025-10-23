// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitVolume] class.
var (
	UnitVolumeClass     _UnitVolumeClass
	UnitVolumeClassOnce sync.Once
)

func getUnitVolumeClass() _UnitVolumeClass {
	UnitVolumeClassOnce.Do(func() {
		UnitVolumeClass = _UnitVolumeClass{objc.GetClass("NSUnitVolume")}
	})
	return UnitVolumeClass
}

type _UnitVolumeClass struct {
	class objc.Class
}

// An interface definition for the [UnitVolume] class.
type IUnitVolume interface {
	IDimension
}

// A unit of measure for volume.
//
// You typically use instances of to represent specific quantities of volume using the class.


// A unit of measure for volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume
type UnitVolume struct {
	Dimension
}

// UnitVolumeFrom constructs a [UnitVolume] from an unsafe.Pointer.
//
// A unit of measure for volume.
func UnitVolumeFrom(ptr unsafe.Pointer) UnitVolume {
	return UnitVolume{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitVolumeClass) Alloc() UnitVolume {
	rv := objc.Send[UnitVolume](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitVolumeClass) New() UnitVolume {
	rv := objc.Send[UnitVolume](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitVolume) Init() UnitVolume {
	rv := objc.Send[UnitVolume](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitVolume) Autorelease() UnitVolume {
	rv := objc.Send[UnitVolume](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitVolume creates a new UnitVolume instance.
func NewUnitVolume() UnitVolume {
	return getUnitVolumeClass().New()
}



// The acre feet unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/acreFeet
func (uc _UnitVolumeClass) AcreFeet() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("acreFeet"))
	return rv
}

// The bushels unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/bushels
func (uc _UnitVolumeClass) Bushels() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("bushels"))
	return rv
}

// The centiliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/centiliters
func (uc _UnitVolumeClass) Centiliters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("centiliters"))
	return rv
}

// The cubic centimeters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicCentimeters
func (uc _UnitVolumeClass) CubicCentimeters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicCentimeters"))
	return rv
}

// The cubic decimeters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicDecimeters
func (uc _UnitVolumeClass) CubicDecimeters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicDecimeters"))
	return rv
}

// The cubic feet unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicFeet
func (uc _UnitVolumeClass) CubicFeet() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicFeet"))
	return rv
}

// The cubic inches unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicInches
func (uc _UnitVolumeClass) CubicInches() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicInches"))
	return rv
}

// The cubic kilometers unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicKilometers
func (uc _UnitVolumeClass) CubicKilometers() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicKilometers"))
	return rv
}

// The cubic meters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMeters
func (uc _UnitVolumeClass) CubicMeters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicMeters"))
	return rv
}

// The cubic miles unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMiles
func (uc _UnitVolumeClass) CubicMiles() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicMiles"))
	return rv
}

// The cubic millimeters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMillimeters
func (uc _UnitVolumeClass) CubicMillimeters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicMillimeters"))
	return rv
}

// The cubic yards unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicYards
func (uc _UnitVolumeClass) CubicYards() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cubicYards"))
	return rv
}

// The cups unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cups
func (uc _UnitVolumeClass) Cups() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("cups"))
	return rv
}

// The deciliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/deciliters
func (uc _UnitVolumeClass) Deciliters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("deciliters"))
	return rv
}

// The fluid ounces unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/fluidOunces
func (uc _UnitVolumeClass) FluidOunces() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("fluidOunces"))
	return rv
}

// The gallons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/gallons
func (uc _UnitVolumeClass) Gallons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("gallons"))
	return rv
}

// The imperial fluid ounces unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialFluidOunces
func (uc _UnitVolumeClass) ImperialFluidOunces() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialFluidOunces"))
	return rv
}

// The imperial gallons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialGallons
func (uc _UnitVolumeClass) ImperialGallons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialGallons"))
	return rv
}

// The imperial pints unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialPints
func (uc _UnitVolumeClass) ImperialPints() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialPints"))
	return rv
}

// The imperial quarts unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialQuarts
func (uc _UnitVolumeClass) ImperialQuarts() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialQuarts"))
	return rv
}

// The imperial tablespoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialTablespoons
func (uc _UnitVolumeClass) ImperialTablespoons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialTablespoons"))
	return rv
}

// The imperial teaspoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialTeaspoons
func (uc _UnitVolumeClass) ImperialTeaspoons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("imperialTeaspoons"))
	return rv
}

// The kiloliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/kiloliters
func (uc _UnitVolumeClass) Kiloliters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("kiloliters"))
	return rv
}

// The liters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/liters
func (uc _UnitVolumeClass) Liters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("liters"))
	return rv
}

// The megaliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/megaliters
func (uc _UnitVolumeClass) Megaliters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("megaliters"))
	return rv
}

// The metric cups unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/metricCups
func (uc _UnitVolumeClass) MetricCups() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("metricCups"))
	return rv
}

// The milliliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/milliliters
func (uc _UnitVolumeClass) Milliliters() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("milliliters"))
	return rv
}

// The pints unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/pints
func (uc _UnitVolumeClass) Pints() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("pints"))
	return rv
}

// The quarts unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/quarts
func (uc _UnitVolumeClass) Quarts() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("quarts"))
	return rv
}

// The tablespoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/tablespoons
func (uc _UnitVolumeClass) Tablespoons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("tablespoons"))
	return rv
}

// The teaspoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/teaspoons
func (uc _UnitVolumeClass) Teaspoons() UnitVolume {
	rv := objc.Send[NSUnitVolume](objc.ID(uc.class), objc.Sel("teaspoons"))
	return rv
}

// The acre feet unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/acreFeet
func (u_ UnitVolume) AcreFeet() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("acreFeet"))
	return rv
}


// The bushels unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/bushels
func (u_ UnitVolume) Bushels() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("bushels"))
	return rv
}


// The centiliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/centiliters
func (u_ UnitVolume) Centiliters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("centiliters"))
	return rv
}


// The cubic centimeters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicCentimeters
func (u_ UnitVolume) CubicCentimeters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicCentimeters"))
	return rv
}


// The cubic decimeters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicDecimeters
func (u_ UnitVolume) CubicDecimeters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicDecimeters"))
	return rv
}


// The cubic feet unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicFeet
func (u_ UnitVolume) CubicFeet() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicFeet"))
	return rv
}


// The cubic inches unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicInches
func (u_ UnitVolume) CubicInches() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicInches"))
	return rv
}


// The cubic kilometers unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicKilometers
func (u_ UnitVolume) CubicKilometers() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicKilometers"))
	return rv
}


// The cubic meters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMeters
func (u_ UnitVolume) CubicMeters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicMeters"))
	return rv
}


// The cubic miles unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMiles
func (u_ UnitVolume) CubicMiles() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicMiles"))
	return rv
}


// The cubic millimeters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMillimeters
func (u_ UnitVolume) CubicMillimeters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicMillimeters"))
	return rv
}


// The cubic yards unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicYards
func (u_ UnitVolume) CubicYards() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cubicYards"))
	return rv
}


// The cups unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/cups
func (u_ UnitVolume) Cups() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("cups"))
	return rv
}


// The deciliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/deciliters
func (u_ UnitVolume) Deciliters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("deciliters"))
	return rv
}


// The fluid ounces unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/fluidOunces
func (u_ UnitVolume) FluidOunces() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("fluidOunces"))
	return rv
}


// The gallons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/gallons
func (u_ UnitVolume) Gallons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("gallons"))
	return rv
}


// The imperial fluid ounces unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialFluidOunces
func (u_ UnitVolume) ImperialFluidOunces() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialFluidOunces"))
	return rv
}


// The imperial gallons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialGallons
func (u_ UnitVolume) ImperialGallons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialGallons"))
	return rv
}


// The imperial pints unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialPints
func (u_ UnitVolume) ImperialPints() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialPints"))
	return rv
}


// The imperial quarts unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialQuarts
func (u_ UnitVolume) ImperialQuarts() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialQuarts"))
	return rv
}


// The imperial tablespoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialTablespoons
func (u_ UnitVolume) ImperialTablespoons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialTablespoons"))
	return rv
}


// The imperial teaspoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialTeaspoons
func (u_ UnitVolume) ImperialTeaspoons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("imperialTeaspoons"))
	return rv
}


// The kiloliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/kiloliters
func (u_ UnitVolume) Kiloliters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("kiloliters"))
	return rv
}


// The liters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/liters
func (u_ UnitVolume) Liters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("liters"))
	return rv
}


// The megaliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/megaliters
func (u_ UnitVolume) Megaliters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("megaliters"))
	return rv
}


// The metric cups unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/metricCups
func (u_ UnitVolume) MetricCups() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("metricCups"))
	return rv
}


// The milliliters unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/milliliters
func (u_ UnitVolume) Milliliters() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("milliliters"))
	return rv
}


// The pints unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/pints
func (u_ UnitVolume) Pints() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("pints"))
	return rv
}


// The quarts unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/quarts
func (u_ UnitVolume) Quarts() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("quarts"))
	return rv
}


// The tablespoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/tablespoons
func (u_ UnitVolume) Tablespoons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("tablespoons"))
	return rv
}


// The teaspoons unit of volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume/teaspoons
func (u_ UnitVolume) Teaspoons() NSUnitVolume {
	rv := objc.Send[NSUnitVolume](u_.ID, objc.Sel("teaspoons"))
	return rv
}



