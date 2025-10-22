// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKScoredAssessmentType] class.
type IHKScoredAssessmentType interface {
	IHKSampleType
}



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

// Alloc allocates a new instance without initialization.
func (hc _HKScoredAssessmentTypeClass) Alloc() HKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




