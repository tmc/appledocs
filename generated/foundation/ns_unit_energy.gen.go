// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [UnitEnergy] class.
var (
	UnitEnergyClass     _UnitEnergyClass
	UnitEnergyClassOnce sync.Once
)

func getUnitEnergyClass() _UnitEnergyClass {
	UnitEnergyClassOnce.Do(func() {
		UnitEnergyClass = _UnitEnergyClass{objc.GetClass("NSUnitEnergy")}
	})
	return UnitEnergyClass
}

type _UnitEnergyClass struct {
	class objc.Class
}





// An interface definition for the [UnitEnergy] class.
type IUnitEnergy interface {
	IDimension
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (uc _UnitEnergyClass) Alloc() UnitEnergy {
	rv := objc.Send[UnitEnergy](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitEnergyClass) New() UnitEnergy {
	rv := objc.Send[UnitEnergy](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitEnergy) Init() UnitEnergy {
	rv := objc.Send[UnitEnergy](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitEnergy) Autorelease() UnitEnergy {
	rv := objc.Send[UnitEnergy](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitEnergy creates a new UnitEnergy instance.
func NewUnitEnergy() UnitEnergy {
	return getUnitEnergyClass().New()
}





// A unit of measure for energy.
//
// You typically use instances of to represent specific quantities of energy using the class.


// A unit of measure for energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitEnergy
type UnitEnergy struct {
	Dimension
}

// UnitEnergyFrom constructs a [UnitEnergy] from an unsafe.Pointer.
//
// A unit of measure for energy.
func UnitEnergyFrom(ptr unsafe.Pointer) UnitEnergy {
	return UnitEnergy{
		Dimension: DimensionFrom(ptr),
	}
}















// The joules unit of energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitEnergy/joules
func (uc _UnitEnergyClass) Joules() UnitEnergy {
	rv := objc.Send[UnitEnergy](objc.ID(uc.class), objc.Sel("joules"))
	return rv
}











// The joules unit of energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitEnergy/joules
func (u_ UnitEnergy) Joules() IUnitEnergy {
	rv := objc.Send[UnitEnergy](u_.ID, objc.Sel("joules"))
	return rv
}








