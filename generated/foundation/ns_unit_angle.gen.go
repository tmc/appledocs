// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitAngle */


/* debug [class_header]: Header for NSUnitAngle */
// The class instance for the [UnitAngle] class.
var (
	UnitAngleClass     _UnitAngleClass
	UnitAngleClassOnce sync.Once
)

func getUnitAngleClass() _UnitAngleClass {
	UnitAngleClassOnce.Do(func() {
		UnitAngleClass = _UnitAngleClass{objc.GetClass("NSUnitAngle")}
	})
	return UnitAngleClass
}

type _UnitAngleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitAngle */
// An interface definition for the [UnitAngle] class.
type IUnitAngle interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitAngle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitAngle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitAngle */
// Alloc allocates a new instance without initialization.
func (uc _UnitAngleClass) Alloc() UnitAngle {
	rv := objc.Send[UnitAngle](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitAngleClass) New() UnitAngle {
	rv := objc.Send[UnitAngle](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitAngle) Init() UnitAngle {
	rv := objc.Send[UnitAngle](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitAngle) Autorelease() UnitAngle {
	rv := objc.Send[UnitAngle](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitAngle creates a new UnitAngle instance.
func NewUnitAngle() UnitAngle {
	return getUnitAngleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitAngle */
// A unit of measure for planar angle and rotation.
//
// You typically use instances of to represent specific quantities of planar angle using the class.


// A unit of measure for planar angle and rotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle
type UnitAngle struct {
	Dimension
}

// UnitAngleFrom constructs a [UnitAngle] from an unsafe.Pointer.
//
// A unit of measure for planar angle and rotation.
func UnitAngleFrom(ptr unsafe.Pointer) UnitAngle {
	return UnitAngle{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitAngle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitAngle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitAngle */

// The degrees unit of angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle/degrees
func (uc _UnitAngleClass) Degrees() UnitAngle {
	rv := objc.Send[UnitAngle](objc.ID(uc.class), objc.Sel("degrees"))
	return rv
}/* debug [class_properties_class/property]: degrees */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitAngle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitAngle */

// The degrees unit of angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle/degrees
func (u_ UnitAngle) Degrees() IUnitAngle {
	rv := objc.Send[UnitAngle](u_.ID, objc.Sel("degrees"))
	return rv
}/* debug [instance_properties/getter]: degrees */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitAngle */



