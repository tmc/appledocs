// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricCharge] class.
var unitElectricChargeClass = _UnitElectricChargeClass{objc.GetClass("NSUnitElectricCharge")}

type _UnitElectricChargeClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricCharge] class.
type IUnitElectricCharge interface {
	IDimension
}

// A unit of measure for electric charge. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return unitElectricChargeClass.New()
}




