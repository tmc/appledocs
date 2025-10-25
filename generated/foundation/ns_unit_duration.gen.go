// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitDuration */


/* debug [class_header]: Header for NSUnitDuration */
// The class instance for the [UnitDuration] class.
var (
	UnitDurationClass     _UnitDurationClass
	UnitDurationClassOnce sync.Once
)

func getUnitDurationClass() _UnitDurationClass {
	UnitDurationClassOnce.Do(func() {
		UnitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}
	})
	return UnitDurationClass
}

type _UnitDurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitDuration */
// An interface definition for the [UnitDuration] class.
type IUnitDuration interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitDuration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitDuration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitDuration */
// Alloc allocates a new instance without initialization.
func (uc _UnitDurationClass) Alloc() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitDurationClass) New() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitDuration) Init() UnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitDuration) Autorelease() UnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitDuration creates a new UnitDuration instance.
func NewUnitDuration() UnitDuration {
	return getUnitDurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitDuration */
// A unit of measure for a duration of time.
//
// You typically use instances of to represent specific quantities of planar angle using the class.


// A unit of measure for a duration of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration
type UnitDuration struct {
	Dimension
}

// UnitDurationFrom constructs a [UnitDuration] from an unsafe.Pointer.
//
// A unit of measure for a duration of time.
func UnitDurationFrom(ptr unsafe.Pointer) UnitDuration {
	return UnitDuration{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitDuration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitDuration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitDuration */

// The second unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/seconds
func (uc _UnitDurationClass) Seconds() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("seconds"))
	return rv
}/* debug [class_properties_class/property]: seconds */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitDuration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitDuration */

// The second unit of duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration/seconds
func (u_ UnitDuration) Seconds() IUnitDuration {
	rv := objc.Send[UnitDuration](u_.ID, objc.Sel("seconds"))
	return rv
}/* debug [instance_properties/getter]: seconds */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitDuration */



