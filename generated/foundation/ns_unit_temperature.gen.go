// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitTemperature */


/* debug [class_header]: Header for NSUnitTemperature */
// The class instance for the [UnitTemperature] class.
var (
	UnitTemperatureClass     _UnitTemperatureClass
	UnitTemperatureClassOnce sync.Once
)

func getUnitTemperatureClass() _UnitTemperatureClass {
	UnitTemperatureClassOnce.Do(func() {
		UnitTemperatureClass = _UnitTemperatureClass{objc.GetClass("NSUnitTemperature")}
	})
	return UnitTemperatureClass
}

type _UnitTemperatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitTemperature */
// An interface definition for the [UnitTemperature] class.
type IUnitTemperature interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitTemperature */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitTemperature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitTemperature */
// Alloc allocates a new instance without initialization.
func (uc _UnitTemperatureClass) Alloc() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitTemperatureClass) New() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitTemperature) Init() UnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitTemperature) Autorelease() UnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitTemperature creates a new UnitTemperature instance.
func NewUnitTemperature() UnitTemperature {
	return getUnitTemperatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitTemperature */
// A unit of measure for temperature.
//
// You typically use instances of to represent specific quantities of temperature using the class.


// A unit of measure for temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature
type UnitTemperature struct {
	Dimension
}

// UnitTemperatureFrom constructs a [UnitTemperature] from an unsafe.Pointer.
//
// A unit of measure for temperature.
func UnitTemperatureFrom(ptr unsafe.Pointer) UnitTemperature {
	return UnitTemperature{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitTemperature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitTemperature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitTemperature */

// The degree Celsius unit of temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/celsius
func (uc _UnitTemperatureClass) Celsius() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("celsius"))
	return rv
}/* debug [class_properties_class/property]: celsius */

// The degree Fahrenheit unit of temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/fahrenheit
func (uc _UnitTemperatureClass) Fahrenheit() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("fahrenheit"))
	return rv
}/* debug [class_properties_class/property]: fahrenheit */

// The kelvin unit of temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/kelvin
func (uc _UnitTemperatureClass) Kelvin() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("kelvin"))
	return rv
}/* debug [class_properties_class/property]: kelvin */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitTemperature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitTemperature */

// The degree Celsius unit of temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/celsius
func (u_ UnitTemperature) Celsius() IUnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("celsius"))
	return rv
}/* debug [instance_properties/getter]: celsius */


// The degree Fahrenheit unit of temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/fahrenheit
func (u_ UnitTemperature) Fahrenheit() IUnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("fahrenheit"))
	return rv
}/* debug [instance_properties/getter]: fahrenheit */


// The kelvin unit of temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature/kelvin
func (u_ UnitTemperature) Kelvin() IUnitTemperature {
	rv := objc.Send[UnitTemperature](u_.ID, objc.Sel("kelvin"))
	return rv
}/* debug [instance_properties/getter]: kelvin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitTemperature */



