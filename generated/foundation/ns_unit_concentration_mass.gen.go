// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitConcentrationMass */


/* debug [class_header]: Header for NSUnitConcentrationMass */
// The class instance for the [UnitConcentrationMass] class.
var (
	UnitConcentrationMassClass     _UnitConcentrationMassClass
	UnitConcentrationMassClassOnce sync.Once
)

func getUnitConcentrationMassClass() _UnitConcentrationMassClass {
	UnitConcentrationMassClassOnce.Do(func() {
		UnitConcentrationMassClass = _UnitConcentrationMassClass{objc.GetClass("NSUnitConcentrationMass")}
	})
	return UnitConcentrationMassClass
}

type _UnitConcentrationMassClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitConcentrationMass */
// An interface definition for the [UnitConcentrationMass] class.
type IUnitConcentrationMass interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitConcentrationMass */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitConcentrationMass */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitConcentrationMass */
// Alloc allocates a new instance without initialization.
func (uc _UnitConcentrationMassClass) Alloc() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getUnitConcentrationMassClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitConcentrationMass */
// A unit of measure for concentration of mass.
//
// You typically use instances of to represent specific quantities of concentration using the class.


// A unit of measure for concentration of mass.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitConcentrationMass *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitConcentrationMass */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitConcentrationMass */

// The grams per liter unit of concentration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass/gramsPerLiter
func (uc _UnitConcentrationMassClass) GramsPerLiter() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](objc.ID(uc.class), objc.Sel("gramsPerLiter"))
	return rv
}/* debug [class_properties_class/property]: gramsPerLiter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitConcentrationMass */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitConcentrationMass */

// The grams per liter unit of concentration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass/gramsPerLiter
func (u_ UnitConcentrationMass) GramsPerLiter() IUnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](u_.ID, objc.Sel("gramsPerLiter"))
	return rv
}/* debug [instance_properties/getter]: gramsPerLiter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitConcentrationMass */



