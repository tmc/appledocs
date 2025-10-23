// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitPower] class.
var (
	UnitPowerClass     _UnitPowerClass
	UnitPowerClassOnce sync.Once
)

func getUnitPowerClass() _UnitPowerClass {
	UnitPowerClassOnce.Do(func() {
		UnitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}
	})
	return UnitPowerClass
}

type _UnitPowerClass struct {
	class objc.Class
}

// An interface definition for the [UnitPower] class.
type IUnitPower interface {
	IDimension
}

// A unit of measure for power.
//
// You typically use instances of to represent specific quantities of power using the class.


// A unit of measure for power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower
type UnitPower struct {
	Dimension
}

// UnitPowerFrom constructs a [UnitPower] from an unsafe.Pointer.
//
// A unit of measure for power.
func UnitPowerFrom(ptr unsafe.Pointer) UnitPower {
	return UnitPower{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitPowerClass) Alloc() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitPowerClass) New() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitPower) Init() UnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitPower) Autorelease() UnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitPower creates a new UnitPower instance.
func NewUnitPower() UnitPower {
	return getUnitPowerClass().New()
}



// The femtowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/femtowatts
func (uc _UnitPowerClass) Femtowatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("femtowatts"))
	return rv
}

// The gigawatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/gigawatts
func (uc _UnitPowerClass) Gigawatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("gigawatts"))
	return rv
}

// The horsepower unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/horsepower
func (uc _UnitPowerClass) Horsepower() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("horsepower"))
	return rv
}

// The kilowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/kilowatts
func (uc _UnitPowerClass) Kilowatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("kilowatts"))
	return rv
}

// The megawatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/megawatts
func (uc _UnitPowerClass) Megawatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("megawatts"))
	return rv
}

// The microwatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/microwatts
func (uc _UnitPowerClass) Microwatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("microwatts"))
	return rv
}

// The milliwatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/milliwatts
func (uc _UnitPowerClass) Milliwatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("milliwatts"))
	return rv
}

// The nanowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/nanowatts
func (uc _UnitPowerClass) Nanowatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("nanowatts"))
	return rv
}

// The picowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/picowatts
func (uc _UnitPowerClass) Picowatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("picowatts"))
	return rv
}

// The terawatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/terawatts
func (uc _UnitPowerClass) Terawatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("terawatts"))
	return rv
}

// The watts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/watts
func (uc _UnitPowerClass) Watts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("watts"))
	return rv
}

// The femtowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/femtowatts
func (u_ UnitPower) Femtowatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("femtowatts"))
	return rv
}


// The gigawatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/gigawatts
func (u_ UnitPower) Gigawatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("gigawatts"))
	return rv
}


// The horsepower unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/horsepower
func (u_ UnitPower) Horsepower() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("horsepower"))
	return rv
}


// The kilowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/kilowatts
func (u_ UnitPower) Kilowatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("kilowatts"))
	return rv
}


// The megawatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/megawatts
func (u_ UnitPower) Megawatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("megawatts"))
	return rv
}


// The microwatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/microwatts
func (u_ UnitPower) Microwatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("microwatts"))
	return rv
}


// The milliwatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/milliwatts
func (u_ UnitPower) Milliwatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("milliwatts"))
	return rv
}


// The nanowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/nanowatts
func (u_ UnitPower) Nanowatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("nanowatts"))
	return rv
}


// The picowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/picowatts
func (u_ UnitPower) Picowatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("picowatts"))
	return rv
}


// The terawatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/terawatts
func (u_ UnitPower) Terawatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("terawatts"))
	return rv
}


// The watts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/watts
func (u_ UnitPower) Watts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("watts"))
	return rv
}



