// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKScoredAssessment */


/* debug [class_header]: Header for HKScoredAssessment */
// The class instance for the [HKScoredAssessment] class.
var (
	HKScoredAssessmentClass     _HKScoredAssessmentClass
	HKScoredAssessmentClassOnce sync.Once
)

func getHKScoredAssessmentClass() _HKScoredAssessmentClass {
	HKScoredAssessmentClassOnce.Do(func() {
		HKScoredAssessmentClass = _HKScoredAssessmentClass{objc.GetClass("HKScoredAssessment")}
	})
	return HKScoredAssessmentClass
}

type _HKScoredAssessmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKScoredAssessment */
// An interface definition for the [HKScoredAssessment] class.
type IHKScoredAssessment interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKScoredAssessment */
	// properties:
	Score() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKScoredAssessment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKScoredAssessment */
// Alloc allocates a new instance without initialization.
func (hc _HKScoredAssessmentClass) Alloc() HKScoredAssessment {
	rv := objc.Send[HKScoredAssessment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKScoredAssessmentClass) New() HKScoredAssessment {
	rv := objc.Send[HKScoredAssessment](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKScoredAssessment) Init() HKScoredAssessment {
	rv := objc.Send[HKScoredAssessment](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKScoredAssessment) Autorelease() HKScoredAssessment {
	rv := objc.Send[HKScoredAssessment](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKScoredAssessment creates a new HKScoredAssessment instance.
func NewHKScoredAssessment() HKScoredAssessment {
	return getHKScoredAssessmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKScoredAssessment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKScoredAssessment
type HKScoredAssessment struct {
	HKSample
}

// HKScoredAssessmentFrom constructs a [HKScoredAssessment] from an unsafe.Pointer.
func HKScoredAssessmentFrom(ptr unsafe.Pointer) HKScoredAssessment {
	return HKScoredAssessment{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKScoredAssessment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKScoredAssessment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKScoredAssessment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKScoredAssessment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKScoredAssessment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKScoredAssessment/score
func (h_ HKScoredAssessment) Score() int {
	rv := objc.Send[int](h_.ID, objc.Sel("score"))
	return rv
}/* debug [instance_properties/getter]: score */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKScoredAssessment */



