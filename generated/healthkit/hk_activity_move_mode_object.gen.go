// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKActivityMoveModeObject */


/* debug [class_header]: Header for HKActivityMoveModeObject */
// The class instance for the [HKActivityMoveModeObject] class.
var (
	HKActivityMoveModeObjectClass     _HKActivityMoveModeObjectClass
	HKActivityMoveModeObjectClassOnce sync.Once
)

func getHKActivityMoveModeObjectClass() _HKActivityMoveModeObjectClass {
	HKActivityMoveModeObjectClassOnce.Do(func() {
		HKActivityMoveModeObjectClass = _HKActivityMoveModeObjectClass{objc.GetClass("HKActivityMoveModeObject")}
	})
	return HKActivityMoveModeObjectClass
}

type _HKActivityMoveModeObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKActivityMoveModeObject */
// An interface definition for the [HKActivityMoveModeObject] class.
type IHKActivityMoveModeObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKActivityMoveModeObject */
	// properties:
	ActivityMoveMode() HKActivityMoveMode
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKActivityMoveModeObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKActivityMoveModeObject */
// Alloc allocates a new instance without initialization.
func (hc _HKActivityMoveModeObjectClass) Alloc() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKActivityMoveModeObjectClass) New() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivityMoveModeObject) Init() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivityMoveModeObject) Autorelease() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivityMoveModeObject creates a new HKActivityMoveModeObject instance.
func NewHKActivityMoveModeObject() HKActivityMoveModeObject {
	return getHKActivityMoveModeObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKActivityMoveModeObject */
// An object that contains a movement mode value.


// An object that contains a movement mode value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivityMoveModeObject
type HKActivityMoveModeObject struct {
	objectivec.Object
}

// HKActivityMoveModeObjectFrom constructs a [HKActivityMoveModeObject] from an unsafe.Pointer.
//
// An object that contains a movement mode value.
func HKActivityMoveModeObjectFrom(ptr unsafe.Pointer) HKActivityMoveModeObject {
	return HKActivityMoveModeObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKActivityMoveModeObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKActivityMoveModeObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKActivityMoveModeObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKActivityMoveModeObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKActivityMoveModeObject */

// A property that contains the movement mode value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivityMoveModeObject/activityMoveMode
func (h_ HKActivityMoveModeObject) ActivityMoveMode() HKActivityMoveMode {
	rv := objc.Send[HKActivityMoveMode](h_.ID, objc.Sel("activityMoveMode"))
	return rv
}/* debug [instance_properties/getter]: activityMoveMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKActivityMoveModeObject */



