// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitMass */


/* debug [class_header]: Header for NSUnitMass */
// The class instance for the [UnitMass] class.
var (
	UnitMassClass     _UnitMassClass
	UnitMassClassOnce sync.Once
)

func getUnitMassClass() _UnitMassClass {
	UnitMassClassOnce.Do(func() {
		UnitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}
	})
	return UnitMassClass
}

type _UnitMassClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitMass */
// An interface definition for the [UnitMass] class.
type IUnitMass interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitMass */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitMass */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitMass */
// Alloc allocates a new instance without initialization.
func (uc _UnitMassClass) Alloc() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitMassClass) New() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitMass) Init() UnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitMass) Autorelease() UnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitMass creates a new UnitMass instance.
func NewUnitMass() UnitMass {
	return getUnitMassClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitMass */
// A unit of measure for mass.
//
// You typically use instances of to represent specific quantities of mass using the class.


// A unit of measure for mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass
type UnitMass struct {
	Dimension
}

// UnitMassFrom constructs a [UnitMass] from an unsafe.Pointer.
//
// A unit of measure for mass.
func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitMass *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitMass */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitMass */

// The kilograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/kilograms
func (uc _UnitMassClass) Kilograms() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("kilograms"))
	return rv
}/* debug [class_properties_class/property]: kilograms */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitMass */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitMass */

// The kilograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/kilograms
func (u_ UnitMass) Kilograms() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("kilograms"))
	return rv
}/* debug [instance_properties/getter]: kilograms */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitMass */



