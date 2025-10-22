// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKScoredAssessment] class.
type IHKScoredAssessment interface {
	IHKSample
	Score() int
	SetScore(value int)
}



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

// Alloc allocates a new instance without initialization.
func (hc _HKScoredAssessmentClass) Alloc() HKScoredAssessment {
	rv := objc.Send[HKScoredAssessment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkscoredassessment/score

func (h_ HKScoredAssessment) Score() int {
	rv := objc.Send[int](h_.ID, objc.Sel("score"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkscoredassessment/score

func (h_ HKScoredAssessment) SetScore(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setScore:"), value)
}



