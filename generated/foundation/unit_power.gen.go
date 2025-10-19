// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitPower] class.
var (
	unitPowerClass     _UnitPowerClass
	unitPowerClassOnce sync.Once
)

func getUnitPowerClass() _UnitPowerClass {
	unitPowerClassOnce.Do(func() {
		unitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}
	})
	return unitPowerClass
}

type _UnitPowerClass struct {
	class objc.Class
}

// An interface definition for the [UnitPower] class.
type IUnitPower interface {
	IDimension
}

// A unit of measure for power. [Full Topic]
//
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




