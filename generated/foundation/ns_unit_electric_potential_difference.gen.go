// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricPotentialDifference] class.
var (
	UnitElectricPotentialDifferenceClass     _UnitElectricPotentialDifferenceClass
	UnitElectricPotentialDifferenceClassOnce sync.Once
)

func getUnitElectricPotentialDifferenceClass() _UnitElectricPotentialDifferenceClass {
	UnitElectricPotentialDifferenceClassOnce.Do(func() {
		UnitElectricPotentialDifferenceClass = _UnitElectricPotentialDifferenceClass{objc.GetClass("NSUnitElectricPotentialDifference")}
	})
	return UnitElectricPotentialDifferenceClass
}

type _UnitElectricPotentialDifferenceClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricPotentialDifference] class.
type IUnitElectricPotentialDifference interface {
	IDimension
}

// A unit of measure for electric potential difference.
//
// You typically use instances of to represent specific quantities of electric potential difference using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference
type UnitElectricPotentialDifference struct {
	Dimension
}

// UnitElectricPotentialDifferenceFrom constructs a [UnitElectricPotentialDifference] from an unsafe.Pointer.
//
// A unit of measure for electric potential difference.
func UnitElectricPotentialDifferenceFrom(ptr unsafe.Pointer) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifference{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitElectricPotentialDifferenceClass) Alloc() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitElectricPotentialDifferenceClass) New() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricPotentialDifference) Init() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricPotentialDifference) Autorelease() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricPotentialDifference creates a new UnitElectricPotentialDifference instance.
func NewUnitElectricPotentialDifference() UnitElectricPotentialDifference {
	return getUnitElectricPotentialDifferenceClass().New()
}




