// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Answers() unsafe.Pointer
	SetAnswers(value unsafe.Pointer)
	Risk() unsafe.Pointer
	SetRisk(value unsafe.Pointer)
	// methods:
}



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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkphq9assessment/answers-4y95e
func (h_ HKPHQ9Assessment) Answers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("answers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkphq9assessment/answers-4y95e
func (h_ HKPHQ9Assessment) SetAnswers(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAnswers:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkphq9assessment/risk-swift.property
func (h_ HKPHQ9Assessment) Risk() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("risk"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkphq9assessment/risk-swift.property
func (h_ HKPHQ9Assessment) SetRisk(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRisk:"), value)
}



