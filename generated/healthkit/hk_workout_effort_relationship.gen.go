// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutEffortRelationship */


/* debug [class_header]: Header for HKWorkoutEffortRelationship */
// The class instance for the [HKWorkoutEffortRelationship] class.
var (
	HKWorkoutEffortRelationshipClass     _HKWorkoutEffortRelationshipClass
	HKWorkoutEffortRelationshipClassOnce sync.Once
)

func getHKWorkoutEffortRelationshipClass() _HKWorkoutEffortRelationshipClass {
	HKWorkoutEffortRelationshipClassOnce.Do(func() {
		HKWorkoutEffortRelationshipClass = _HKWorkoutEffortRelationshipClass{objc.GetClass("HKWorkoutEffortRelationship")}
	})
	return HKWorkoutEffortRelationshipClass
}

type _HKWorkoutEffortRelationshipClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutEffortRelationship */
// An interface definition for the [HKWorkoutEffortRelationship] class.
type IHKWorkoutEffortRelationship interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKWorkoutEffortRelationship */
	// properties:
	Activity() IHKWorkoutActivity
	Samples() []HKSample
	Workout() IHKWorkout
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutEffortRelationship */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutEffortRelationship */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutEffortRelationshipClass) Alloc() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutEffortRelationshipClass) New() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutEffortRelationship) Init() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutEffortRelationship) Autorelease() HKWorkoutEffortRelationship {
	rv := objc.Send[HKWorkoutEffortRelationship](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutEffortRelationship creates a new HKWorkoutEffortRelationship instance.
func NewHKWorkoutEffortRelationship() HKWorkoutEffortRelationship {
	return getHKWorkoutEffortRelationshipClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutEffortRelationship */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationship
type HKWorkoutEffortRelationship struct {
	objectivec.Object
}

// HKWorkoutEffortRelationshipFrom constructs a [HKWorkoutEffortRelationship] from an unsafe.Pointer.
func HKWorkoutEffortRelationshipFrom(ptr unsafe.Pointer) HKWorkoutEffortRelationship {
	return HKWorkoutEffortRelationship{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutEffortRelationship *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutEffortRelationship */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutEffortRelationship */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutEffortRelationship */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutEffortRelationship */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationship/activity
func (h_ HKWorkoutEffortRelationship) Activity() IHKWorkoutActivity {
	rv := objc.Send[HKWorkoutActivity](h_.ID, objc.Sel("activity"))
	return rv
}/* debug [instance_properties/getter]: activity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationship/samples
func (h_ HKWorkoutEffortRelationship) Samples() []HKSample {
	rv := objc.Send[[]HKSample](h_.ID, objc.Sel("samples"))
	return rv
}/* debug [instance_properties/getter]: samples */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationship/workout
func (h_ HKWorkoutEffortRelationship) Workout() IHKWorkout {
	rv := objc.Send[HKWorkout](h_.ID, objc.Sel("workout"))
	return rv
}/* debug [instance_properties/getter]: workout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutEffortRelationship */



