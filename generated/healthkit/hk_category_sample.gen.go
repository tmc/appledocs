// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCategorySample] class.
var (
	HKCategorySampleClass     _HKCategorySampleClass
	HKCategorySampleClassOnce sync.Once
)

func getHKCategorySampleClass() _HKCategorySampleClass {
	HKCategorySampleClassOnce.Do(func() {
		HKCategorySampleClass = _HKCategorySampleClass{objc.GetClass("HKCategorySample")}
	})
	return HKCategorySampleClass
}

type _HKCategorySampleClass struct {
	class objc.Class
}

// An interface definition for the [HKCategorySample] class.
type IHKCategorySample interface {
	IHKSample
}

// A sample with values from a short list of possible values.
//
// You can use category samples to record data associated with a . The value for the sample must come from the appropriate category value enumeration. Each category type uses its own enumeration. Individual samples represent a value and time period. Samples with different values may have overlapping time intervals. The class is a concrete subclass of the class. Category samples are immutable: You set the sample’s properties when you create it, and they can’t change.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample
type HKCategorySample struct {
	HKSample
}

// HKCategorySampleFrom constructs a [HKCategorySample] from an unsafe.Pointer.
//
// A sample with values from a short list of possible values.
func HKCategorySampleFrom(ptr unsafe.Pointer) HKCategorySample {
	return HKCategorySample{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCategorySampleClass) Alloc() HKCategorySample {
	rv := objc.Send[HKCategorySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCategorySampleClass) New() HKCategorySample {
	rv := objc.Send[HKCategorySample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCategorySample) Init() HKCategorySample {
	rv := objc.Send[HKCategorySample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCategorySample) Autorelease() HKCategorySample {
	rv := objc.Send[HKCategorySample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCategorySample creates a new HKCategorySample instance.
func NewHKCategorySample() HKCategorySample {
	return getHKCategorySampleClass().New()
}


// The category type for this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/categoryType
func (h_ HKCategorySample) CategoryType() HKCategoryType {
	rv := objc.Send[HKCategoryType](h_.ID, objc.Sel("categoryType"))
	return rv
}

// The category value for this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcategorysample/value
func (h_ HKCategorySample) Value() int {
	rv := objc.Send[int](h_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The category value for this sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcategorysample/value
func (h_ HKCategorySample) SetValue(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setValue:"), value)
}

// The key path for accessing the category sample’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcategoryvalue
func (h_ HKCategorySample) HKPredicateKeyPathCategoryValue() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathCategoryValue"))
	return rv
}



