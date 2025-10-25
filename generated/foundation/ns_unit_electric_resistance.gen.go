// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitElectricResistance */


/* debug [class_header]: Header for NSUnitElectricResistance */
// The class instance for the [UnitElectricResistance] class.
var (
	UnitElectricResistanceClass     _UnitElectricResistanceClass
	UnitElectricResistanceClassOnce sync.Once
)

func getUnitElectricResistanceClass() _UnitElectricResistanceClass {
	UnitElectricResistanceClassOnce.Do(func() {
		UnitElectricResistanceClass = _UnitElectricResistanceClass{objc.GetClass("NSUnitElectricResistance")}
	})
	return UnitElectricResistanceClass
}

type _UnitElectricResistanceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitElectricResistance */
// An interface definition for the [UnitElectricResistance] class.
type IUnitElectricResistance interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitElectricResistance */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitElectricResistance */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitElectricResistance */
// Alloc allocates a new instance without initialization.
func (uc _UnitElectricResistanceClass) Alloc() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitElectricResistanceClass) New() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricResistance) Init() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricResistance) Autorelease() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricResistance creates a new UnitElectricResistance instance.
func NewUnitElectricResistance() UnitElectricResistance {
	return getUnitElectricResistanceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitElectricResistance */
// A unit of measure for electric resistance.
//
// You typically use instances of to represent specific quantities of electric resistance using the class.


// A unit of measure for electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance
type UnitElectricResistance struct {
	Dimension
}

// UnitElectricResistanceFrom constructs a [UnitElectricResistance] from an unsafe.Pointer.
//
// A unit of measure for electric resistance.
func UnitElectricResistanceFrom(ptr unsafe.Pointer) UnitElectricResistance {
	return UnitElectricResistance{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitElectricResistance *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitElectricResistance */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitElectricResistance */

// The ohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/ohms
func (uc _UnitElectricResistanceClass) Ohms() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("ohms"))
	return rv
}/* debug [class_properties_class/property]: ohms */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitElectricResistance */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitElectricResistance */

// The ohms unit of electric resistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/ohms
func (u_ UnitElectricResistance) Ohms() IUnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u_.ID, objc.Sel("ohms"))
	return rv
}/* debug [instance_properties/getter]: ohms */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitElectricResistance */



