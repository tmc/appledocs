// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitEnergy */


/* debug [class_header]: Header for NSUnitEnergy */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitEnergy */
// An interface definition for the [UnitEnergy] class.
type IUnitEnergy interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitEnergy */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitEnergy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitEnergy */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitEnergy */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitEnergy *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitEnergy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitEnergy */

// The joules unit of energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitEnergy/joules
func (uc _UnitEnergyClass) Joules() UnitEnergy {
	rv := objc.Send[UnitEnergy](objc.ID(uc.class), objc.Sel("joules"))
	return rv
}/* debug [class_properties_class/property]: joules */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitEnergy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitEnergy */

// The joules unit of energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitEnergy/joules
func (u_ UnitEnergy) Joules() IUnitEnergy {
	rv := objc.Send[UnitEnergy](u_.ID, objc.Sel("joules"))
	return rv
}/* debug [instance_properties/getter]: joules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitEnergy */



