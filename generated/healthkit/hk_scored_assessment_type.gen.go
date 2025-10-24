// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKScoredAssessmentType */


/* debug [class_header]: Header for HKScoredAssessmentType */
// The class instance for the [HKScoredAssessmentType] class.
var (
	HKScoredAssessmentTypeClass     _HKScoredAssessmentTypeClass
	HKScoredAssessmentTypeClassOnce sync.Once
)

func getHKScoredAssessmentTypeClass() _HKScoredAssessmentTypeClass {
	HKScoredAssessmentTypeClassOnce.Do(func() {
		HKScoredAssessmentTypeClass = _HKScoredAssessmentTypeClass{objc.GetClass("HKScoredAssessmentType")}
	})
	return HKScoredAssessmentTypeClass
}

type _HKScoredAssessmentTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKScoredAssessmentType */
// An interface definition for the [HKScoredAssessmentType] class.
type IHKScoredAssessmentType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKScoredAssessmentType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKScoredAssessmentType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKScoredAssessmentType */
// Alloc allocates a new instance without initialization.
func (hc _HKScoredAssessmentTypeClass) Alloc() HKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKScoredAssessmentTypeClass) New() HKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKScoredAssessmentType) Init() HKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKScoredAssessmentType) Autorelease() HKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKScoredAssessmentType creates a new HKScoredAssessmentType instance.
func NewHKScoredAssessmentType() HKScoredAssessmentType {
	return getHKScoredAssessmentTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKScoredAssessmentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKScoredAssessmentType
type HKScoredAssessmentType struct {
	HKSampleType
}

// HKScoredAssessmentTypeFrom constructs a [HKScoredAssessmentType] from an unsafe.Pointer.
func HKScoredAssessmentTypeFrom(ptr unsafe.Pointer) HKScoredAssessmentType {
	return HKScoredAssessmentType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKScoredAssessmentType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKScoredAssessmentType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKScoredAssessmentType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKScoredAssessmentType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKScoredAssessmentType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKScoredAssessmentType */



