// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitDispersion */


/* debug [class_header]: Header for NSUnitDispersion */
// The class instance for the [UnitDispersion] class.
var (
	UnitDispersionClass     _UnitDispersionClass
	UnitDispersionClassOnce sync.Once
)

func getUnitDispersionClass() _UnitDispersionClass {
	UnitDispersionClassOnce.Do(func() {
		UnitDispersionClass = _UnitDispersionClass{objc.GetClass("NSUnitDispersion")}
	})
	return UnitDispersionClass
}

type _UnitDispersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitDispersion */
// An interface definition for the [UnitDispersion] class.
type IUnitDispersion interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitDispersion */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitDispersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitDispersion */
// Alloc allocates a new instance without initialization.
func (uc _UnitDispersionClass) Alloc() UnitDispersion {
	rv := objc.Send[UnitDispersion](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitDispersionClass) New() UnitDispersion {
	rv := objc.Send[UnitDispersion](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitDispersion) Init() UnitDispersion {
	rv := objc.Send[UnitDispersion](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitDispersion) Autorelease() UnitDispersion {
	rv := objc.Send[UnitDispersion](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitDispersion creates a new UnitDispersion instance.
func NewUnitDispersion() UnitDispersion {
	return getUnitDispersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitDispersion */
// A unit of measure for specific quantities of dispersion.
//
// You typically use instances of to represent specific quantities of dispersion using the class.


// A unit of measure for specific quantities of dispersion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion
type UnitDispersion struct {
	Dimension
}

// UnitDispersionFrom constructs a [UnitDispersion] from an unsafe.Pointer.
//
// A unit of measure for specific quantities of dispersion.
func UnitDispersionFrom(ptr unsafe.Pointer) UnitDispersion {
	return UnitDispersion{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitDispersion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitDispersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitDispersion */

// The parts per million unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion/partsPerMillion
func (uc _UnitDispersionClass) PartsPerMillion() UnitDispersion {
	rv := objc.Send[UnitDispersion](objc.ID(uc.class), objc.Sel("partsPerMillion"))
	return rv
}/* debug [class_properties_class/property]: partsPerMillion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitDispersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitDispersion */

// The parts per million unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion/partsPerMillion
func (u_ UnitDispersion) PartsPerMillion() IUnitDispersion {
	rv := objc.Send[UnitDispersion](u_.ID, objc.Sel("partsPerMillion"))
	return rv
}/* debug [instance_properties/getter]: partsPerMillion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitDispersion */



