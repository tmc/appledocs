// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKPHQ9Assessment */


/* debug [class_header]: Header for HKPHQ9Assessment */
// The class instance for the [HKPHQ9Assessment] class.
var (
	HKPHQ9AssessmentClass     _HKPHQ9AssessmentClass
	HKPHQ9AssessmentClassOnce sync.Once
)

func getHKPHQ9AssessmentClass() _HKPHQ9AssessmentClass {
	HKPHQ9AssessmentClassOnce.Do(func() {
		HKPHQ9AssessmentClass = _HKPHQ9AssessmentClass{objc.GetClass("HKPHQ9Assessment")}
	})
	return HKPHQ9AssessmentClass
}

type _HKPHQ9AssessmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKPHQ9Assessment */
// An interface definition for the [HKPHQ9Assessment] class.
type IHKPHQ9Assessment interface {
	IHKScoredAssessment
	
/* debug [class_interface_properties]: Properties for HKPHQ9Assessment */
	// properties:
	Answers() []foundation.Number
	Risk() HKPHQ9AssessmentRisk
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKPHQ9Assessment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKPHQ9Assessment */
// Alloc allocates a new instance without initialization.
func (hc _HKPHQ9AssessmentClass) Alloc() HKPHQ9Assessment {
	rv := objc.Send[HKPHQ9Assessment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKPHQ9AssessmentClass) New() HKPHQ9Assessment {
	rv := objc.Send[HKPHQ9Assessment](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKPHQ9Assessment) Init() HKPHQ9Assessment {
	rv := objc.Send[HKPHQ9Assessment](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKPHQ9Assessment) Autorelease() HKPHQ9Assessment {
	rv := objc.Send[HKPHQ9Assessment](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKPHQ9Assessment creates a new HKPHQ9Assessment instance.
func NewHKPHQ9Assessment() HKPHQ9Assessment {
	return getHKPHQ9AssessmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKPHQ9Assessment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment
type HKPHQ9Assessment struct {
	HKScoredAssessment
}

// HKPHQ9AssessmentFrom constructs a [HKPHQ9Assessment] from an unsafe.Pointer.
func HKPHQ9AssessmentFrom(ptr unsafe.Pointer) HKPHQ9Assessment {
	return HKPHQ9Assessment{
		HKScoredAssessment: HKScoredAssessmentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKPHQ9Assessment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKPHQ9Assessment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/assessmentWithDate:answers:
func (hc _HKPHQ9AssessmentClass) AssessmentWithDateAnswers(date objc.IObject /* cross-framework: NSDate */, answers []foundation.Number) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("assessmentWithDate:answers:"), date, answers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssessmentWithDateAnswers) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/assessmentWithDate:answers:metadata:
func (hc _HKPHQ9AssessmentClass) AssessmentWithDateAnswersMetadata(date objc.IObject /* cross-framework: NSDate */, answers []foundation.Number, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("assessmentWithDate:answers:metadata:"), date, answers, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssessmentWithDateAnswersMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKPHQ9Assessment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKPHQ9Assessment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKPHQ9Assessment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/answers-439dt
func (h_ HKPHQ9Assessment) Answers() []foundation.Number {
	rv := objc.Send[[]foundation.Number](h_.ID, objc.Sel("answers"))
	return rv
}/* debug [instance_properties/getter]: answers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/risk-swift.property
func (h_ HKPHQ9Assessment) Risk() HKPHQ9AssessmentRisk {
	rv := objc.Send[HKPHQ9AssessmentRisk](h_.ID, objc.Sel("risk"))
	return rv
}/* debug [instance_properties/getter]: risk */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKPHQ9Assessment */



