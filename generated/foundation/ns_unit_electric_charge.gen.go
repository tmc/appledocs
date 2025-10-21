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
//
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


// The coulombs unit of electric charge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/coulombs
func (uc _UnitElectricChargeClass) Coulombs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("coulombs"))
	return rv
}
// The coulombs unit of electric charge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/coulombs
func (u_ UnitElectricCharge) Coulombs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("coulombs"))
	return rv
}



