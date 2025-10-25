// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitAcceleration */


/* debug [class_header]: Header for NSUnitAcceleration */
// The class instance for the [UnitAcceleration] class.
var (
	UnitAccelerationClass     _UnitAccelerationClass
	UnitAccelerationClassOnce sync.Once
)

func getUnitAccelerationClass() _UnitAccelerationClass {
	UnitAccelerationClassOnce.Do(func() {
		UnitAccelerationClass = _UnitAccelerationClass{objc.GetClass("NSUnitAcceleration")}
	})
	return UnitAccelerationClass
}

type _UnitAccelerationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitAcceleration */
// An interface definition for the [UnitAcceleration] class.
type IUnitAcceleration interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitAcceleration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitAcceleration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitAcceleration */
// Alloc allocates a new instance without initialization.
func (uc _UnitAccelerationClass) Alloc() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitAccelerationClass) New() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitAcceleration) Init() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitAcceleration) Autorelease() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitAcceleration creates a new UnitAcceleration instance.
func NewUnitAcceleration() UnitAcceleration {
	return getUnitAccelerationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitAcceleration */
// A unit of measure for acceleration.
//
// You typically use instances of to represent specific quantities of acceleration using the class.


// A unit of measure for acceleration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration
type UnitAcceleration struct {
	Dimension
}

// UnitAccelerationFrom constructs a [UnitAcceleration] from an unsafe.Pointer.
//
// A unit of measure for acceleration.
func UnitAccelerationFrom(ptr unsafe.Pointer) UnitAcceleration {
	return UnitAcceleration{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitAcceleration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitAcceleration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitAcceleration */

// Returns the gravity unit of acceleration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration/gravity
func (uc _UnitAccelerationClass) Gravity() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](objc.ID(uc.class), objc.Sel("gravity"))
	return rv
}/* debug [class_properties_class/property]: gravity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitAcceleration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitAcceleration */

// Returns the gravity unit of acceleration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration/gravity
func (u_ UnitAcceleration) Gravity() IUnitAcceleration {
	rv := objc.Send[UnitAcceleration](u_.ID, objc.Sel("gravity"))
	return rv
}/* debug [instance_properties/getter]: gravity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitAcceleration */



