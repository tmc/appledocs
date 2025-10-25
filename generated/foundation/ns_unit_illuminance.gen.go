// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitIlluminance */


/* debug [class_header]: Header for NSUnitIlluminance */
// The class instance for the [UnitIlluminance] class.
var (
	UnitIlluminanceClass     _UnitIlluminanceClass
	UnitIlluminanceClassOnce sync.Once
)

func getUnitIlluminanceClass() _UnitIlluminanceClass {
	UnitIlluminanceClassOnce.Do(func() {
		UnitIlluminanceClass = _UnitIlluminanceClass{objc.GetClass("NSUnitIlluminance")}
	})
	return UnitIlluminanceClass
}

type _UnitIlluminanceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitIlluminance */
// An interface definition for the [UnitIlluminance] class.
type IUnitIlluminance interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitIlluminance */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitIlluminance */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitIlluminance */
// Alloc allocates a new instance without initialization.
func (uc _UnitIlluminanceClass) Alloc() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitIlluminanceClass) New() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitIlluminance) Init() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitIlluminance) Autorelease() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitIlluminance creates a new UnitIlluminance instance.
func NewUnitIlluminance() UnitIlluminance {
	return getUnitIlluminanceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitIlluminance */
// A unit of measure for illuminance.
//
// You typically use instances of to represent specific quantities of illuminance using the class.


// A unit of measure for illuminance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance
type UnitIlluminance struct {
	Dimension
}

// UnitIlluminanceFrom constructs a [UnitIlluminance] from an unsafe.Pointer.
//
// A unit of measure for illuminance.
func UnitIlluminanceFrom(ptr unsafe.Pointer) UnitIlluminance {
	return UnitIlluminance{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitIlluminance *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitIlluminance */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitIlluminance */

// The lux unit of illuminance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance/lux
func (uc _UnitIlluminanceClass) Lux() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](objc.ID(uc.class), objc.Sel("lux"))
	return rv
}/* debug [class_properties_class/property]: lux */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitIlluminance */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitIlluminance */

// The lux unit of illuminance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance/lux
func (u_ UnitIlluminance) Lux() IUnitIlluminance {
	rv := objc.Send[UnitIlluminance](u_.ID, objc.Sel("lux"))
	return rv
}/* debug [instance_properties/getter]: lux */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitIlluminance */



