// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitElectricCharge */


/* debug [class_header]: Header for NSUnitElectricCharge */
// The class instance for the [UnitElectricCharge] class.
var (
	UnitElectricChargeClass     _UnitElectricChargeClass
	UnitElectricChargeClassOnce sync.Once
)

func getUnitElectricChargeClass() _UnitElectricChargeClass {
	UnitElectricChargeClassOnce.Do(func() {
		UnitElectricChargeClass = _UnitElectricChargeClass{objc.GetClass("NSUnitElectricCharge")}
	})
	return UnitElectricChargeClass
}

type _UnitElectricChargeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitElectricCharge */
// An interface definition for the [UnitElectricCharge] class.
type IUnitElectricCharge interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitElectricCharge */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitElectricCharge */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitElectricCharge */
// Alloc allocates a new instance without initialization.
func (uc _UnitElectricChargeClass) Alloc() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitElectricChargeClass) New() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricCharge) Init() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricCharge) Autorelease() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricCharge creates a new UnitElectricCharge instance.
func NewUnitElectricCharge() UnitElectricCharge {
	return getUnitElectricChargeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitElectricCharge */
// A unit of measure for electric charge.
//
// You typically use instances of to represent specific quantities of electric charge using the class.


// A unit of measure for electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge
type UnitElectricCharge struct {
	Dimension
}

// UnitElectricChargeFrom constructs a [UnitElectricCharge] from an unsafe.Pointer.
//
// A unit of measure for electric charge.
func UnitElectricChargeFrom(ptr unsafe.Pointer) UnitElectricCharge {
	return UnitElectricCharge{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitElectricCharge *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitElectricCharge */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitElectricCharge */

// The ampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/ampereHours
func (uc _UnitElectricChargeClass) AmpereHours() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("ampereHours"))
	return rv
}/* debug [class_properties_class/property]: ampereHours */

// The kiloampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/kiloampereHours
func (uc _UnitElectricChargeClass) KiloampereHours() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("kiloampereHours"))
	return rv
}/* debug [class_properties_class/property]: kiloampereHours */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitElectricCharge */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitElectricCharge */

// The ampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/ampereHours
func (u_ UnitElectricCharge) AmpereHours() IUnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u_.ID, objc.Sel("ampereHours"))
	return rv
}/* debug [instance_properties/getter]: ampereHours */


// The kiloampere hours unit of electric charge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/kiloampereHours
func (u_ UnitElectricCharge) KiloampereHours() IUnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u_.ID, objc.Sel("kiloampereHours"))
	return rv
}/* debug [instance_properties/getter]: kiloampereHours */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitElectricCharge */



