// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKGAD7Assessment */


/* debug [class_header]: Header for HKGAD7Assessment */
// The class instance for the [HKGAD7Assessment] class.
var (
	HKGAD7AssessmentClass     _HKGAD7AssessmentClass
	HKGAD7AssessmentClassOnce sync.Once
)

func getHKGAD7AssessmentClass() _HKGAD7AssessmentClass {
	HKGAD7AssessmentClassOnce.Do(func() {
		HKGAD7AssessmentClass = _HKGAD7AssessmentClass{objc.GetClass("HKGAD7Assessment")}
	})
	return HKGAD7AssessmentClass
}

type _HKGAD7AssessmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKGAD7Assessment */
// An interface definition for the [HKGAD7Assessment] class.
type IHKGAD7Assessment interface {
	IHKScoredAssessment
	
/* debug [class_interface_properties]: Properties for HKGAD7Assessment */
	// properties:
	Answers() []foundation.Number
	Risk() HKGAD7AssessmentRisk
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKGAD7Assessment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKGAD7Assessment */
// Alloc allocates a new instance without initialization.
func (hc _HKGAD7AssessmentClass) Alloc() HKGAD7Assessment {
	rv := objc.Send[HKGAD7Assessment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKGAD7AssessmentClass) New() HKGAD7Assessment {
	rv := objc.Send[HKGAD7Assessment](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKGAD7Assessment) Init() HKGAD7Assessment {
	rv := objc.Send[HKGAD7Assessment](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKGAD7Assessment) Autorelease() HKGAD7Assessment {
	rv := objc.Send[HKGAD7Assessment](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKGAD7Assessment creates a new HKGAD7Assessment instance.
func NewHKGAD7Assessment() HKGAD7Assessment {
	return getHKGAD7AssessmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKGAD7Assessment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment
type HKGAD7Assessment struct {
	HKScoredAssessment
}

// HKGAD7AssessmentFrom constructs a [HKGAD7Assessment] from an unsafe.Pointer.
func HKGAD7AssessmentFrom(ptr unsafe.Pointer) HKGAD7Assessment {
	return HKGAD7Assessment{
		HKScoredAssessment: HKScoredAssessmentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKGAD7Assessment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKGAD7Assessment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/assessmentWithDate:answers:
func (hc _HKGAD7AssessmentClass) AssessmentWithDateAnswers(date objc.IObject /* cross-framework: NSDate */, answers []foundation.Number) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("assessmentWithDate:answers:"), date, answers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssessmentWithDateAnswers) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/assessmentWithDate:answers:metadata:
func (hc _HKGAD7AssessmentClass) AssessmentWithDateAnswersMetadata(date objc.IObject /* cross-framework: NSDate */, answers []foundation.Number, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("assessmentWithDate:answers:metadata:"), date, answers, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssessmentWithDateAnswersMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKGAD7Assessment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKGAD7Assessment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKGAD7Assessment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/answers-7k5w8
func (h_ HKGAD7Assessment) Answers() []foundation.Number {
	rv := objc.Send[[]foundation.Number](h_.ID, objc.Sel("answers"))
	return rv
}/* debug [instance_properties/getter]: answers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/risk-swift.property
func (h_ HKGAD7Assessment) Risk() HKGAD7AssessmentRisk {
	rv := objc.Send[HKGAD7AssessmentRisk](h_.ID, objc.Sel("risk"))
	return rv
}/* debug [instance_properties/getter]: risk */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKGAD7Assessment */



