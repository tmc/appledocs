// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitConcentrationMass] class.
var unitConcentrationMassClass = _UnitConcentrationMassClass{objc.GetClass("NSUnitConcentrationMass")}

type _UnitConcentrationMassClass struct {
	class objc.Class
}

// An interface definition for the [UnitConcentrationMass] class.
type IUnitConcentrationMass interface {
	IDimension
}

// A unit of measure for concentration of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass

type UnitConcentrationMass struct {
	Dimension
}

// UnitConcentrationMassFrom constructs a [UnitConcentrationMass] from an unsafe.Pointer.
//
// A unit of measure for concentration of mass.
func UnitConcentrationMassFrom(ptr unsafe.Pointer) UnitConcentrationMass {
	return UnitConcentrationMass{
		Dimension: DimensionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (uc _UnitConcentrationMassClass) Alloc() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _UnitConcentrationMassClass) New() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitConcentrationMass) Init() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitConcentrationMass) Autorelease() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConcentrationMass creates a new UnitConcentrationMass instance.
func NewUnitConcentrationMass() UnitConcentrationMass {
	return unitConcentrationMassClass.New()
}




