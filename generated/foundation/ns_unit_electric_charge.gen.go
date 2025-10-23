// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricCharge] class.
var (
	UnitElectricChargeClass     _UnitElectricChargeClass
	UnitElectricChargeClassOnce sync.Once
)

func getUnitElectricChargeClass() _UnitElectricChargeClass {
	UnitElectricChargeClassOnce.Do(func() {
		UnitElectricChargeClass = _UnitElectricChargeClass{objc.GetClass("NSUnitElectricCharge")}
	})
	return UnitElectricChargeClass
}

type _UnitElectricChargeClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricCharge] class.
type IUnitElectricCharge interface {
	IDimension
}

// A unit of measure for electric charge.
//
// You typically use instances of to represent specific quantities of electric charge using the class.


// A unit of measure for electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge
type UnitElectricCharge struct {
	Dimension
}

// UnitElectricChargeFrom constructs a [UnitElectricCharge] from an unsafe.Pointer.
//
// A unit of measure for electric charge.
func UnitElectricChargeFrom(ptr unsafe.Pointer) UnitElectricCharge {
	return UnitElectricCharge{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitElectricChargeClass) Alloc() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitElectricChargeClass) New() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricCharge) Init() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricCharge) Autorelease() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricCharge creates a new UnitElectricCharge instance.
func NewUnitElectricCharge() UnitElectricCharge {
	return getUnitElectricChargeClass().New()
}



// The ampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/ampereHours
func (uc _UnitElectricChargeClass) AmpereHours() UnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](objc.ID(uc.class), objc.Sel("ampereHours"))
	return rv
}

// The coulombs unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/coulombs
func (uc _UnitElectricChargeClass) Coulombs() UnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](objc.ID(uc.class), objc.Sel("coulombs"))
	return rv
}

// The kiloampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/kiloampereHours
func (uc _UnitElectricChargeClass) KiloampereHours() UnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](objc.ID(uc.class), objc.Sel("kiloampereHours"))
	return rv
}

// The megaampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/megaampereHours
func (uc _UnitElectricChargeClass) MegaampereHours() UnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](objc.ID(uc.class), objc.Sel("megaampereHours"))
	return rv
}

// The microampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/microampereHours
func (uc _UnitElectricChargeClass) MicroampereHours() UnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](objc.ID(uc.class), objc.Sel("microampereHours"))
	return rv
}

// The milliampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/milliampereHours
func (uc _UnitElectricChargeClass) MilliampereHours() UnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](objc.ID(uc.class), objc.Sel("milliampereHours"))
	return rv
}

// The ampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/ampereHours
func (u_ UnitElectricCharge) AmpereHours() IUnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](u_.ID, objc.Sel("ampereHours"))
	return rv
}


// The coulombs unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/coulombs
func (u_ UnitElectricCharge) Coulombs() IUnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](u_.ID, objc.Sel("coulombs"))
	return rv
}


// The kiloampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/kiloampereHours
func (u_ UnitElectricCharge) KiloampereHours() IUnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](u_.ID, objc.Sel("kiloampereHours"))
	return rv
}


// The megaampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/megaampereHours
func (u_ UnitElectricCharge) MegaampereHours() IUnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](u_.ID, objc.Sel("megaampereHours"))
	return rv
}


// The microampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/microampereHours
func (u_ UnitElectricCharge) MicroampereHours() IUnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](u_.ID, objc.Sel("microampereHours"))
	return rv
}


// The milliampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/milliampereHours
func (u_ UnitElectricCharge) MilliampereHours() IUnitElectricCharge {
	rv := objc.Send[NSUnitElectricCharge](u_.ID, objc.Sel("milliampereHours"))
	return rv
}



