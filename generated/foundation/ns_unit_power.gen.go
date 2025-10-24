// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitPower */


/* debug [class_header]: Header for NSUnitPower */
// The class instance for the [UnitPower] class.
var (
	UnitPowerClass     _UnitPowerClass
	UnitPowerClassOnce sync.Once
)

func getUnitPowerClass() _UnitPowerClass {
	UnitPowerClassOnce.Do(func() {
		UnitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}
	})
	return UnitPowerClass
}

type _UnitPowerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitPower */
// An interface definition for the [UnitPower] class.
type IUnitPower interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitPower */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitPower */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitPower */
// Alloc allocates a new instance without initialization.
func (uc _UnitPowerClass) Alloc() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitPower */
// A unit of measure for power.
//
// You typically use instances of to represent specific quantities of power using the class.


// A unit of measure for power.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitPower *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitPower */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitPower */

// The picowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/picowatts
func (uc _UnitPowerClass) Picowatts() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("picowatts"))
	return rv
}/* debug [class_properties_class/property]: picowatts */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitPower */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitPower */

// The picowatts unit of power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower/picowatts
func (u_ UnitPower) Picowatts() IUnitPower {
	rv := objc.Send[UnitPower](u_.ID, objc.Sel("picowatts"))
	return rv
}/* debug [instance_properties/getter]: picowatts */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitPower */



