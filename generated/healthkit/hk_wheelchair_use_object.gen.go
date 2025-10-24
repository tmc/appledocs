// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWheelchairUseObject */


/* debug [class_header]: Header for HKWheelchairUseObject */
// The class instance for the [HKWheelchairUseObject] class.
var (
	HKWheelchairUseObjectClass     _HKWheelchairUseObjectClass
	HKWheelchairUseObjectClassOnce sync.Once
)

func getHKWheelchairUseObjectClass() _HKWheelchairUseObjectClass {
	HKWheelchairUseObjectClassOnce.Do(func() {
		HKWheelchairUseObjectClass = _HKWheelchairUseObjectClass{objc.GetClass("HKWheelchairUseObject")}
	})
	return HKWheelchairUseObjectClass
}

type _HKWheelchairUseObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWheelchairUseObject */
// An interface definition for the [HKWheelchairUseObject] class.
type IHKWheelchairUseObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWheelchairUseObject */
	// properties:
	WheelchairUse() HKWheelchairUse
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWheelchairUseObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWheelchairUseObject */
// Alloc allocates a new instance without initialization.
func (hc _HKWheelchairUseObjectClass) Alloc() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWheelchairUseObjectClass) New() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWheelchairUseObject) Init() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWheelchairUseObject) Autorelease() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWheelchairUseObject creates a new HKWheelchairUseObject instance.
func NewHKWheelchairUseObject() HKWheelchairUseObject {
	return getHKWheelchairUseObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWheelchairUseObject */
// This class acts as a wrapper for the wheelchair use enumeration.


// This class acts as a wrapper for the wheelchair use enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUseObject
type HKWheelchairUseObject struct {
	objectivec.Object
}

// HKWheelchairUseObjectFrom constructs a [HKWheelchairUseObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the wheelchair use enumeration.
func HKWheelchairUseObjectFrom(ptr unsafe.Pointer) HKWheelchairUseObject {
	return HKWheelchairUseObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWheelchairUseObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWheelchairUseObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWheelchairUseObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWheelchairUseObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWheelchairUseObject */

// A value indicating the user’s wheelchair use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUseObject/wheelchairUse
func (h_ HKWheelchairUseObject) WheelchairUse() HKWheelchairUse {
	rv := objc.Send[HKWheelchairUse](h_.ID, objc.Sel("wheelchairUse"))
	return rv
}/* debug [instance_properties/getter]: wheelchairUse */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWheelchairUseObject */



