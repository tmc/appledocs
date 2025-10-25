// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitSpeed */


/* debug [class_header]: Header for NSUnitSpeed */
// The class instance for the [UnitSpeed] class.
var (
	UnitSpeedClass     _UnitSpeedClass
	UnitSpeedClassOnce sync.Once
)

func getUnitSpeedClass() _UnitSpeedClass {
	UnitSpeedClassOnce.Do(func() {
		UnitSpeedClass = _UnitSpeedClass{objc.GetClass("NSUnitSpeed")}
	})
	return UnitSpeedClass
}

type _UnitSpeedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitSpeed */
// An interface definition for the [UnitSpeed] class.
type IUnitSpeed interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitSpeed */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitSpeed */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitSpeed */
// Alloc allocates a new instance without initialization.
func (uc _UnitSpeedClass) Alloc() UnitSpeed {
	rv := objc.Send[UnitSpeed](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitSpeedClass) New() UnitSpeed {
	rv := objc.Send[UnitSpeed](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitSpeed) Init() UnitSpeed {
	rv := objc.Send[UnitSpeed](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitSpeed) Autorelease() UnitSpeed {
	rv := objc.Send[UnitSpeed](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitSpeed creates a new UnitSpeed instance.
func NewUnitSpeed() UnitSpeed {
	return getUnitSpeedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitSpeed */
// A unit of measure for speed.
//
// You typically use instances of to represent specific quantities of speed using the class.


// A unit of measure for speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitSpeed
type UnitSpeed struct {
	Dimension
}

// UnitSpeedFrom constructs a [UnitSpeed] from an unsafe.Pointer.
//
// A unit of measure for speed.
func UnitSpeedFrom(ptr unsafe.Pointer) UnitSpeed {
	return UnitSpeed{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitSpeed *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitSpeed */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitSpeed */

// The meter per second unit of speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitSpeed/metersPerSecond
func (uc _UnitSpeedClass) MetersPerSecond() UnitSpeed {
	rv := objc.Send[UnitSpeed](objc.ID(uc.class), objc.Sel("metersPerSecond"))
	return rv
}/* debug [class_properties_class/property]: metersPerSecond */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitSpeed */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitSpeed */

// The meter per second unit of speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitSpeed/metersPerSecond
func (u_ UnitSpeed) MetersPerSecond() IUnitSpeed {
	rv := objc.Send[UnitSpeed](u_.ID, objc.Sel("metersPerSecond"))
	return rv
}/* debug [instance_properties/getter]: metersPerSecond */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitSpeed */



