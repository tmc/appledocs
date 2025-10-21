// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFuelEfficiency] class.
var (
	UnitFuelEfficiencyClass     _UnitFuelEfficiencyClass
	UnitFuelEfficiencyClassOnce sync.Once
)

func getUnitFuelEfficiencyClass() _UnitFuelEfficiencyClass {
	UnitFuelEfficiencyClassOnce.Do(func() {
		UnitFuelEfficiencyClass = _UnitFuelEfficiencyClass{objc.GetClass("NSUnitFuelEfficiency")}
	})
	return UnitFuelEfficiencyClass
}

type _UnitFuelEfficiencyClass struct {
	class objc.Class
}

// An interface definition for the [UnitFuelEfficiency] class.
type IUnitFuelEfficiency interface {
	IDimension
}

// A unit of measure for fuel efficiency.
//
// You typically use instances of to represent specific quantities of fuel efficiency using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency
type UnitFuelEfficiency struct {
	Dimension
}

// UnitFuelEfficiencyFrom constructs a [UnitFuelEfficiency] from an unsafe.Pointer.
//
// A unit of measure for fuel efficiency.
func UnitFuelEfficiencyFrom(ptr unsafe.Pointer) UnitFuelEfficiency {
	return UnitFuelEfficiency{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitFuelEfficiencyClass) Alloc() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitFuelEfficiencyClass) New() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitFuelEfficiency) Init() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitFuelEfficiency) Autorelease() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitFuelEfficiency creates a new UnitFuelEfficiency instance.
func NewUnitFuelEfficiency() UnitFuelEfficiency {
	return getUnitFuelEfficiencyClass().New()
}


// The miles per gallon unit of fuel efficiency.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerGallon
func (uc _UnitFuelEfficiencyClass) MilesPerGallon() UnitFuelEfficiency {
	rv := objc.Send[NSUnitFuelEfficiency](objc.ID(uc.class), objc.Sel("milesPerGallon"))
	return rv
}
// The miles per imperial gallon unit of fuel efficiency.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerImperialGallon
func (uc _UnitFuelEfficiencyClass) MilesPerImperialGallon() UnitFuelEfficiency {
	rv := objc.Send[NSUnitFuelEfficiency](objc.ID(uc.class), objc.Sel("milesPerImperialGallon"))
	return rv
}
// The miles per gallon unit of fuel efficiency.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerGallon
func (u_ UnitFuelEfficiency) MilesPerGallon() NSUnitFuelEfficiency {
	rv := objc.Send[NSUnitFuelEfficiency](u_.ID, objc.Sel("milesPerGallon"))
	return rv
}

// The miles per imperial gallon unit of fuel efficiency.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerImperialGallon
func (u_ UnitFuelEfficiency) MilesPerImperialGallon() NSUnitFuelEfficiency {
	rv := objc.Send[NSUnitFuelEfficiency](u_.ID, objc.Sel("milesPerImperialGallon"))
	return rv
}



