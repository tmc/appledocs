// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitTemperature] class.
var unitTemperatureClass = _UnitTemperatureClass{objc.GetClass("NSUnitTemperature")}

type _UnitTemperatureClass struct {
	class objc.Class
}

// An interface definition for the [UnitTemperature] class.
type IUnitTemperature interface {
	IDimension
}

// A unit of measure for temperature. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return unitTemperatureClass.New()
}




