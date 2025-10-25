// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitElectricPotentialDifference */


/* debug [class_header]: Header for NSUnitElectricPotentialDifference */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitElectricPotentialDifference */
// An interface definition for the [UnitElectricPotentialDifference] class.
type IUnitElectricPotentialDifference interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitElectricPotentialDifference */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitElectricPotentialDifference */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitElectricPotentialDifference */
// Alloc allocates a new instance without initialization.
func (uc _UnitElectricPotentialDifferenceClass) Alloc() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitElectricPotentialDifference */
// A unit of measure for electric potential difference.
//
// You typically use instances of to represent specific quantities of electric potential difference using the class.


// A unit of measure for electric potential difference.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitElectricPotentialDifference *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitElectricPotentialDifference */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitElectricPotentialDifference */

// The volts unit of electric potential difference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/volts
func (uc _UnitElectricPotentialDifferenceClass) Volts() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](objc.ID(uc.class), objc.Sel("volts"))
	return rv
}/* debug [class_properties_class/property]: volts */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitElectricPotentialDifference */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitElectricPotentialDifference */

// The volts unit of electric potential difference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/volts
func (u_ UnitElectricPotentialDifference) Volts() IUnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](u_.ID, objc.Sel("volts"))
	return rv
}/* debug [instance_properties/getter]: volts */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitElectricPotentialDifference */



