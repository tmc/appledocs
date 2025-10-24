// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitElectricCurrent */


/* debug [class_header]: Header for NSUnitElectricCurrent */
// The class instance for the [UnitElectricCurrent] class.
var (
	UnitElectricCurrentClass     _UnitElectricCurrentClass
	UnitElectricCurrentClassOnce sync.Once
)

func getUnitElectricCurrentClass() _UnitElectricCurrentClass {
	UnitElectricCurrentClassOnce.Do(func() {
		UnitElectricCurrentClass = _UnitElectricCurrentClass{objc.GetClass("NSUnitElectricCurrent")}
	})
	return UnitElectricCurrentClass
}

type _UnitElectricCurrentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitElectricCurrent */
// An interface definition for the [UnitElectricCurrent] class.
type IUnitElectricCurrent interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitElectricCurrent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitElectricCurrent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitElectricCurrent */
// Alloc allocates a new instance without initialization.
func (uc _UnitElectricCurrentClass) Alloc() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitElectricCurrentClass) New() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitElectricCurrent) Init() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitElectricCurrent) Autorelease() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricCurrent creates a new UnitElectricCurrent instance.
func NewUnitElectricCurrent() UnitElectricCurrent {
	return getUnitElectricCurrentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitElectricCurrent */
// A unit of measure for electric current.
//
// You typically use instances of to represent specific quantities of electric current using the class.


// A unit of measure for electric current.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent
type UnitElectricCurrent struct {
	Dimension
}

// UnitElectricCurrentFrom constructs a [UnitElectricCurrent] from an unsafe.Pointer.
//
// A unit of measure for electric current.
func UnitElectricCurrentFrom(ptr unsafe.Pointer) UnitElectricCurrent {
	return UnitElectricCurrent{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitElectricCurrent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitElectricCurrent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitElectricCurrent */

// The amperes unit of electric current.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/amperes
func (uc _UnitElectricCurrentClass) Amperes() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](objc.ID(uc.class), objc.Sel("amperes"))
	return rv
}/* debug [class_properties_class/property]: amperes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitElectricCurrent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitElectricCurrent */

// The amperes unit of electric current.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/amperes
func (u_ UnitElectricCurrent) Amperes() IUnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u_.ID, objc.Sel("amperes"))
	return rv
}/* debug [instance_properties/getter]: amperes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitElectricCurrent */



