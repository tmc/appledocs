// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKQuantitySample] class.
var (
	HKQuantitySampleClass     _HKQuantitySampleClass
	HKQuantitySampleClassOnce sync.Once
)

func getHKQuantitySampleClass() _HKQuantitySampleClass {
	HKQuantitySampleClassOnce.Do(func() {
		HKQuantitySampleClass = _HKQuantitySampleClass{objc.GetClass("HKQuantitySample")}
	})
	return HKQuantitySampleClass
}

type _HKQuantitySampleClass struct {
	class objc.Class
}

// An interface definition for the [HKQuantitySample] class.
type IHKQuantitySample interface {
	IHKSample
}

// A sample that represents a quantity, including the value and the units.
//
// A quantity sample contains one or more objects. Each quantity represents a single piece of data with a single numeric value and the value’s associated units. For example, you can use quantity samples to record the user’s height, the user’s current heart rate, or the number of calories in a hamburger. HealthKit provides a wide range of quantity types, letting you track many different health and fitness features. The class is a subclass of the class. Quantity samples are immutable; you set the sample’s properties when you create it, and they cannot change. In iOS 13 and later and watchOS 6 and later, is an abstract superclass for the and concrete subclasses. The system automatically selects the correct subclass based on the object used to create the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample
type HKQuantitySample struct {
	HKSample
}

// HKQuantitySampleFrom constructs a [HKQuantitySample] from an unsafe.Pointer.
//
// A sample that represents a quantity, including the value and the units.
func HKQuantitySampleFrom(ptr unsafe.Pointer) HKQuantitySample {
	return HKQuantitySample{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQuantitySampleClass) Alloc() HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQuantitySampleClass) New() HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantitySample) Init() HKQuantitySample {
	rv := objc.Send[HKQuantitySample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantitySample) Autorelease() HKQuantitySample {
	rv := objc.Send[HKQuantitySample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantitySample creates a new HKQuantitySample instance.
func NewHKQuantitySample() HKQuantitySample {
	return getHKQuantitySampleClass().New()
}




// Returns a sample containing a numeric measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDate(quantityType unsafe.Pointer, quantity unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:"), quantityType, quantity, startDate, endDate)
	return rv
}



// Returns a sample containing a numeric measurement with the provided device and metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:device:metadata:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDateDeviceMetadata(quantityType unsafe.Pointer, quantity unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, device unsafe.Pointer, metadata unsafe.Pointer) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:device:metadata:"), quantityType, quantity, startDate, endDate, device, metadata)
	return rv
}



// Returns a sample containing a numeric measurement with the provided metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:metadata:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDateMetadata(quantityType unsafe.Pointer, quantity unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, metadata unsafe.Pointer) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), quantityType, quantity, startDate, endDate, metadata)
	return rv
}


// Returns a sample containing a numeric measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDate(quantityType unsafe.Pointer, quantity unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:"), quantityType, quantity, startDate, endDate)
	return rv
}

// Returns a sample containing a numeric measurement with the provided device and metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:device:metadata:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDateDeviceMetadata(quantityType unsafe.Pointer, quantity unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, device unsafe.Pointer, metadata unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:device:metadata:"), quantityType, quantity, startDate, endDate, device, metadata)
	return rv
}

// Returns a sample containing a numeric measurement with the provided metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:metadata:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDateMetadata(quantityType unsafe.Pointer, quantity unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, metadata unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), quantityType, quantity, startDate, endDate, metadata)
	return rv
}

// The number of quantities contained in this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/count
func (h_ HKQuantitySample) Count() int {
	rv := objc.Send[int](h_.ID, objc.Sel("count"))
	return rv
}

// The quantity for this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/quantity
func (h_ HKQuantitySample) Quantity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("quantity"))
	return rv
}

// The quantity type for this sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/quantityType
func (h_ HKQuantitySample) QuantityType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("quantityType"))
	return rv
}

// A key path for the sample’s count.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcount
func (h_ HKQuantitySample) HKPredicateKeyPathCount() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathCount"))
	return rv
}

// The key path for accessing the sample’s quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathquantity
func (h_ HKQuantitySample) HKPredicateKeyPathQuantity() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathQuantity"))
	return rv
}


