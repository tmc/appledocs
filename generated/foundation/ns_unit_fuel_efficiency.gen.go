// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitFuelEfficiency */


/* debug [class_header]: Header for NSUnitFuelEfficiency */
// The class instance for the [UnitFuelEfficiency] class.
var (
	UnitFuelEfficiencyClass     _UnitFuelEfficiencyClass
	UnitFuelEfficiencyClassOnce sync.Once
)

func getUnitFuelEfficiencyClass() _UnitFuelEfficiencyClass {
	UnitFuelEfficiencyClassOnce.Do(func() {
		UnitFuelEfficiencyClass = _UnitFuelEfficiencyClass{objc.GetClass("NSUnitFuelEfficiency")}
	})
	return UnitFuelEfficiencyClass
}

type _UnitFuelEfficiencyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitFuelEfficiency */
// An interface definition for the [UnitFuelEfficiency] class.
type IUnitFuelEfficiency interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitFuelEfficiency */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitFuelEfficiency */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitFuelEfficiency */
// Alloc allocates a new instance without initialization.
func (uc _UnitFuelEfficiencyClass) Alloc() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitFuelEfficiencyClass) New() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitFuelEfficiency) Init() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitFuelEfficiency) Autorelease() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitFuelEfficiency creates a new UnitFuelEfficiency instance.
func NewUnitFuelEfficiency() UnitFuelEfficiency {
	return getUnitFuelEfficiencyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitFuelEfficiency */
// A unit of measure for fuel efficiency.
//
// You typically use instances of to represent specific quantities of fuel efficiency using the class.


// A unit of measure for fuel efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency
type UnitFuelEfficiency struct {
	Dimension
}

// UnitFuelEfficiencyFrom constructs a [UnitFuelEfficiency] from an unsafe.Pointer.
//
// A unit of measure for fuel efficiency.
func UnitFuelEfficiencyFrom(ptr unsafe.Pointer) UnitFuelEfficiency {
	return UnitFuelEfficiency{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitFuelEfficiency *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitFuelEfficiency */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitFuelEfficiency */

// The miles per imperial gallon unit of fuel efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerImperialGallon
func (uc _UnitFuelEfficiencyClass) MilesPerImperialGallon() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](objc.ID(uc.class), objc.Sel("milesPerImperialGallon"))
	return rv
}/* debug [class_properties_class/property]: milesPerImperialGallon */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitFuelEfficiency */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitFuelEfficiency */

// The miles per imperial gallon unit of fuel efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerImperialGallon
func (u_ UnitFuelEfficiency) MilesPerImperialGallon() IUnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u_.ID, objc.Sel("milesPerImperialGallon"))
	return rv
}/* debug [instance_properties/getter]: milesPerImperialGallon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitFuelEfficiency */



