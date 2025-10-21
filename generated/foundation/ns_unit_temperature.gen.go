// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitTemperature] class.
var (
	UnitTemperatureClass     _UnitTemperatureClass
	UnitTemperatureClassOnce sync.Once
)

func getUnitTemperatureClass() _UnitTemperatureClass {
	UnitTemperatureClassOnce.Do(func() {
		UnitTemperatureClass = _UnitTemperatureClass{objc.GetClass("NSUnitTemperature")}
	})
	return UnitTemperatureClass
}

type _UnitTemperatureClass struct {
	class objc.Class
}

// An interface definition for the [UnitTemperature] class.
type IUnitTemperature interface {
	IDimension
}

// A unit of measure for temperature.
//
// You typically use instances of to represent specific quantities of temperature using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature
type UnitTemperature struct {
	Dimension
}

// UnitTemperatureFrom constructs a [UnitTemperature] from an unsafe.Pointer.
//
// A unit of measure for temperature.
func UnitTemperatureFrom(ptr unsafe.Pointer) UnitTemperature {
	return UnitTemperature{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitTemperatureClass) Alloc() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitTemperatureClass) New() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitTemperature) Init() UnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitTemperature) Autorelease() UnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitTemperature creates a new UnitTemperature instance.
func NewUnitTemperature() UnitTemperature {
	return getUnitTemperatureClass().New()
}


// The degree Celsius unit of temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/celsius
func (uc _UnitTemperatureClass) Celsius() UnitTemperature {
	rv := objc.Send[NSUnitTemperature](objc.ID(uc.class), objc.Sel("celsius"))
	return rv
}
// The degree Fahrenheit unit of temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/fahrenheit
func (uc _UnitTemperatureClass) Fahrenheit() UnitTemperature {
	rv := objc.Send[NSUnitTemperature](objc.ID(uc.class), objc.Sel("fahrenheit"))
	return rv
}
// The kelvin unit of temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/kelvin
func (uc _UnitTemperatureClass) Kelvin() UnitTemperature {
	rv := objc.Send[NSUnitTemperature](objc.ID(uc.class), objc.Sel("kelvin"))
	return rv
}
// The degree Celsius unit of temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/celsius
func (u_ UnitTemperature) Celsius() NSUnitTemperature {
	rv := objc.Send[NSUnitTemperature](u_.ID, objc.Sel("celsius"))
	return rv
}

// The degree Fahrenheit unit of temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/fahrenheit
func (u_ UnitTemperature) Fahrenheit() NSUnitTemperature {
	rv := objc.Send[NSUnitTemperature](u_.ID, objc.Sel("fahrenheit"))
	return rv
}

// The kelvin unit of temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/kelvin
func (u_ UnitTemperature) Kelvin() NSUnitTemperature {
	rv := objc.Send[NSUnitTemperature](u_.ID, objc.Sel("kelvin"))
	return rv
}



