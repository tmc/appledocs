// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKGAD7Assessment] class.
type IHKGAD7Assessment interface {
	IHKScoredAssessment
	// properties:
	Answers() unsafe.Pointer
	SetAnswers(value unsafe.Pointer)
	Risk() unsafe.Pointer
	SetRisk(value unsafe.Pointer)
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (hc _HKGAD7AssessmentClass) Alloc() HKGAD7Assessment {
	rv := objc.Send[HKGAD7Assessment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkgad7assessment/answers-1zj1
func (h_ HKGAD7Assessment) Answers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("answers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkgad7assessment/answers-1zj1
func (h_ HKGAD7Assessment) SetAnswers(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAnswers:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkgad7assessment/risk-swift.property
func (h_ HKGAD7Assessment) Risk() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("risk"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkgad7assessment/risk-swift.property
func (h_ HKGAD7Assessment) SetRisk(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRisk:"), value)
}



