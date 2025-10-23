// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricResistance] class.
var (
	UnitElectricResistanceClass     _UnitElectricResistanceClass
	UnitElectricResistanceClassOnce sync.Once
)

func getUnitElectricResistanceClass() _UnitElectricResistanceClass {
	UnitElectricResistanceClassOnce.Do(func() {
		UnitElectricResistanceClass = _UnitElectricResistanceClass{objc.GetClass("NSUnitElectricResistance")}
	})
	return UnitElectricResistanceClass
}

type _UnitElectricResistanceClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricResistance] class.
type IUnitElectricResistance interface {
	IDimension
}

// A unit of measure for electric resistance.
//
// You typically use instances of to represent specific quantities of electric resistance using the class.


// A unit of measure for electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance
type UnitElectricResistance struct {
	Dimension
}

// UnitElectricResistanceFrom constructs a [UnitElectricResistance] from an unsafe.Pointer.
//
// A unit of measure for electric resistance.
func UnitElectricResistanceFrom(ptr unsafe.Pointer) UnitElectricResistance {
	return UnitElectricResistance{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitElectricResistanceClass) Alloc() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitElectricResistanceClass) New() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricResistance) Init() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricResistance) Autorelease() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricResistance creates a new UnitElectricResistance instance.
func NewUnitElectricResistance() UnitElectricResistance {
	return getUnitElectricResistanceClass().New()
}



// The kiloohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/kiloohms
func (uc _UnitElectricResistanceClass) Kiloohms() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("kiloohms"))
	return rv
}

// The megaohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/megaohms
func (uc _UnitElectricResistanceClass) Megaohms() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("megaohms"))
	return rv
}

// The microohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/microohms
func (uc _UnitElectricResistanceClass) Microohms() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("microohms"))
	return rv
}

// The milliohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/milliohms
func (uc _UnitElectricResistanceClass) Milliohms() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("milliohms"))
	return rv
}

// The ohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/ohms
func (uc _UnitElectricResistanceClass) Ohms() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("ohms"))
	return rv
}

// The kiloohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/kiloohms
func (u_ UnitElectricResistance) Kiloohms() IUnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("kiloohms"))
	return rv
}


// The megaohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/megaohms
func (u_ UnitElectricResistance) Megaohms() IUnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("megaohms"))
	return rv
}


// The microohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/microohms
func (u_ UnitElectricResistance) Microohms() IUnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("microohms"))
	return rv
}


// The milliohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/milliohms
func (u_ UnitElectricResistance) Milliohms() IUnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("milliohms"))
	return rv
}


// The ohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/ohms
func (u_ UnitElectricResistance) Ohms() IUnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("ohms"))
	return rv
}



