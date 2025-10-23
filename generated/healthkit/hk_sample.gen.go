// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKSample] class.
var (
	HKSampleClass     _HKSampleClass
	HKSampleClassOnce sync.Once
)

func getHKSampleClass() _HKSampleClass {
	HKSampleClassOnce.Do(func() {
		HKSampleClass = _HKSampleClass{objc.GetClass("HKSample")}
	})
	return HKSampleClass
}

type _HKSampleClass struct {
	class objc.Class
}

// An interface definition for the [HKSample] class.
type IHKSample interface {
	IHKObject
	EndDate() foundation.NSDate
	HasUndeterminedDuration() bool
	SampleType() HKSampleType
	StartDate() foundation.NSDate
	HKPredicateKeyPathEndDate() string
	HKPredicateKeyPathStartDate() string
	HKSampleSortIdentifierEndDate() string
	HKSampleSortIdentifierStartDate() string
}

// A HealthKit sample represents a piece of data associated with a start and end time.
//
// The class is an abstract class. You should never instantiate a object directly. Instead, you always work with one of its concrete subclasses: , , , or classes. HealthKit samples are all immutable: You set the sample’s properties when you create it, and they cannot change. If the sample represents data over a duration, the start time must be earlier than the end time. If the sample represents data at a particular instant, the start and end times can be the same.


// A HealthKit sample represents a piece of data associated with a start and end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample
type HKSample struct {
	HKObject
}

// HKSampleFrom constructs a [HKSample] from an unsafe.Pointer.
//
// A HealthKit sample represents a piece of data associated with a start and end time.
func HKSampleFrom(ptr unsafe.Pointer) HKSample {
	return HKSample{
		HKObject: HKObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKSampleClass) Alloc() HKSample {
	rv := objc.Send[HKSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKSampleClass) New() HKSample {
	rv := objc.Send[HKSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSample) Init() HKSample {
	rv := objc.Send[HKSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSample) Autorelease() HKSample {
	rv := objc.Send[HKSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSample creates a new HKSample instance.
func NewHKSample() HKSample {
	return getHKSampleClass().New()
}



// The sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/endDate
func (h_ HKSample) EndDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("endDate"))
	return rv
}


// Indicates whether the sample has an unknown duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/hasUndeterminedDuration
func (h_ HKSample) HasUndeterminedDuration() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("hasUndeterminedDuration"))
	return rv
}


// The sample type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/sampleType
func (h_ HKSample) SampleType() HKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("sampleType"))
	return rv
}


// The sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/startDate
func (h_ HKSample) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}


// The key path for accessing the sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathenddate
func (h_ HKSample) HKPredicateKeyPathEndDate() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathEndDate"))
	return rv
}


// The key path for accessing the sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathstartdate
func (h_ HKSample) HKPredicateKeyPathStartDate() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathStartDate"))
	return rv
}


// A constant for sorting samples based on their end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplesortidentifierenddate
func (h_ HKSample) HKSampleSortIdentifierEndDate() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKSampleSortIdentifierEndDate"))
	return rv
}


// A constant for sorting samples based on their start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplesortidentifierstartdate
func (h_ HKSample) HKSampleSortIdentifierStartDate() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKSampleSortIdentifierStartDate"))
	return rv
}



