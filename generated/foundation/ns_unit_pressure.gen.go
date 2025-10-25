// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitPressure */


/* debug [class_header]: Header for NSUnitPressure */
// The class instance for the [UnitPressure] class.
var (
	UnitPressureClass     _UnitPressureClass
	UnitPressureClassOnce sync.Once
)

func getUnitPressureClass() _UnitPressureClass {
	UnitPressureClassOnce.Do(func() {
		UnitPressureClass = _UnitPressureClass{objc.GetClass("NSUnitPressure")}
	})
	return UnitPressureClass
}

type _UnitPressureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitPressure */
// An interface definition for the [UnitPressure] class.
type IUnitPressure interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitPressure */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitPressure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitPressure */
// Alloc allocates a new instance without initialization.
func (uc _UnitPressureClass) Alloc() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitPressureClass) New() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitPressure) Init() UnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitPressure) Autorelease() UnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitPressure creates a new UnitPressure instance.
func NewUnitPressure() UnitPressure {
	return getUnitPressureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitPressure */
// A unit of measure for pressure.
//
// You typically use instances of to represent specific quantities of pressure using the class.


// A unit of measure for pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure
type UnitPressure struct {
	Dimension
}

// UnitPressureFrom constructs a [UnitPressure] from an unsafe.Pointer.
//
// A unit of measure for pressure.
func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitPressure *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitPressure */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitPressure */

// The millimeters of mercury unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/millimetersOfMercury
func (uc _UnitPressureClass) MillimetersOfMercury() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("millimetersOfMercury"))
	return rv
}/* debug [class_properties_class/property]: millimetersOfMercury */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitPressure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitPressure */

// The millimeters of mercury unit of pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure/millimetersOfMercury
func (u_ UnitPressure) MillimetersOfMercury() IUnitPressure {
	rv := objc.Send[UnitPressure](u_.ID, objc.Sel("millimetersOfMercury"))
	return rv
}/* debug [instance_properties/getter]: millimetersOfMercury */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitPressure */



