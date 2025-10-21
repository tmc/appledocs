// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKPHQ9Assessment] class.
type IHKPHQ9Assessment interface {
	IHKScoredAssessment
}

//
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

// Alloc allocates a new instance without initialization.
func (hc _HKPHQ9AssessmentClass) Alloc() HKPHQ9Assessment {
	rv := objc.Send[HKPHQ9Assessment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/assessmentWithDate:answers:
func (hc _HKPHQ9AssessmentClass) AssessmentWithDateAnswers(date foundation.IDate, answers []foundation.INumber) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("assessmentWithDate:answers:"), date, answers)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/assessmentWithDate:answers:metadata:
func (hc _HKPHQ9AssessmentClass) AssessmentWithDateAnswersMetadata(date foundation.IDate, answers []foundation.INumber, metadata unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("assessmentWithDate:answers:metadata:"), date, answers, metadata)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/answers-439dt
func (h_ HKPHQ9Assessment) Answers() []foundation.Number {
	rv := objc.Send[[]foundation.Number](h_.ID, objc.Sel("answers"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/risk-swift.property
func (h_ HKPHQ9Assessment) Risk() HKPHQ9AssessmentRisk {
	rv := objc.Send[HKPHQ9AssessmentRisk](h_.ID, objc.Sel("risk"))
	return rv
}



