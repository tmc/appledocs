// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitFrequency */


/* debug [class_header]: Header for NSUnitFrequency */
// The class instance for the [UnitFrequency] class.
var (
	UnitFrequencyClass     _UnitFrequencyClass
	UnitFrequencyClassOnce sync.Once
)

func getUnitFrequencyClass() _UnitFrequencyClass {
	UnitFrequencyClassOnce.Do(func() {
		UnitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}
	})
	return UnitFrequencyClass
}

type _UnitFrequencyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitFrequency */
// An interface definition for the [UnitFrequency] class.
type IUnitFrequency interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitFrequency */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitFrequency */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitFrequency */
// Alloc allocates a new instance without initialization.
func (uc _UnitFrequencyClass) Alloc() UnitFrequency {
	rv := objc.Send[UnitFrequency](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitFrequencyClass) New() UnitFrequency {
	rv := objc.Send[UnitFrequency](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitFrequency) Init() UnitFrequency {
	rv := objc.Send[UnitFrequency](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitFrequency) Autorelease() UnitFrequency {
	rv := objc.Send[UnitFrequency](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitFrequency creates a new UnitFrequency instance.
func NewUnitFrequency() UnitFrequency {
	return getUnitFrequencyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitFrequency */
// A unit of measure for frequency.
//
// You typically use instances of to represent specific quantities of frequency using the class.


// A unit of measure for frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency
type UnitFrequency struct {
	Dimension
}

// UnitFrequencyFrom constructs a [UnitFrequency] from an unsafe.Pointer.
//
// A unit of measure for frequency.
func UnitFrequencyFrom(ptr unsafe.Pointer) UnitFrequency {
	return UnitFrequency{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitFrequency *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitFrequency */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitFrequency */

// The gigahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/gigahertz
func (uc _UnitFrequencyClass) Gigahertz() UnitFrequency {
	rv := objc.Send[UnitFrequency](objc.ID(uc.class), objc.Sel("gigahertz"))
	return rv
}/* debug [class_properties_class/property]: gigahertz */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitFrequency */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitFrequency */

// The gigahertz unit of frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency/gigahertz
func (u_ UnitFrequency) Gigahertz() IUnitFrequency {
	rv := objc.Send[UnitFrequency](u_.ID, objc.Sel("gigahertz"))
	return rv
}/* debug [instance_properties/getter]: gigahertz */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitFrequency */



