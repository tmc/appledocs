// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKWorkoutType */


/* debug [class_header]: Header for HKWorkoutType */
// The class instance for the [HKWorkoutType] class.
var (
	HKWorkoutTypeClass     _HKWorkoutTypeClass
	HKWorkoutTypeClassOnce sync.Once
)

func getHKWorkoutTypeClass() _HKWorkoutTypeClass {
	HKWorkoutTypeClassOnce.Do(func() {
		HKWorkoutTypeClass = _HKWorkoutTypeClass{objc.GetClass("HKWorkoutType")}
	})
	return HKWorkoutTypeClass
}

type _HKWorkoutTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutType */
// An interface definition for the [HKWorkoutType] class.
type IHKWorkoutType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKWorkoutType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutType */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutTypeClass) Alloc() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutTypeClass) New() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutType) Init() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutType) Autorelease() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutType creates a new HKWorkoutType instance.
func NewHKWorkoutType() HKWorkoutType {
	return getHKWorkoutTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutType */
// A type that identifies samples that store information about a workout.
//
// The class is a concrete subclass of the class. To create a workout type instances, use the convenience method. All workouts use the same workout type instance.


// A type that identifies samples that store information about a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutType
type HKWorkoutType struct {
	HKSampleType
}

// HKWorkoutTypeFrom constructs a [HKWorkoutType] from an unsafe.Pointer.
//
// A type that identifies samples that store information about a workout.
func HKWorkoutTypeFrom(ptr unsafe.Pointer) HKWorkoutType {
	return HKWorkoutType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutType */





