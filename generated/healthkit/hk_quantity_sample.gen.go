// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	HKPredicateKeyPathCount() string /* primitive/slice/pointer. */
	HKPredicateKeyPathQuantity() string /* primitive/slice/pointer. */
	Count() int /* primitive/slice/pointer. */
	SetCount(value int /* primitive/slice/pointer. */)
	Quantity() IHKQuantity
	SetQuantity(value IHKQuantity)
	QuantityType() IHKQuantityType
	SetQuantityType(value IHKQuantityType)
	// methods:
}

// A sample that represents a quantity, including the value and the units.
//
// A quantity sample contains one or more objects. Each quantity represents a single piece of data with a single numeric value and the value’s associated units. For example, you can use quantity samples to record the user’s height, the user’s current heart rate, or the number of calories in a hamburger. HealthKit provides a wide range of quantity types, letting you track many different health and fitness features. The class is a subclass of the class. Quantity samples are immutable; you set the sample’s properties when you create it, and they cannot change. In iOS 13 and later and watchOS 6 and later, is an abstract superclass for the and concrete subclasses. The system automatically selects the correct subclass based on the object used to create the sample.


// A sample that represents a quantity, including the value and the units.
//
// [Full Topic]
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



// Returns a sample containing a numeric measurement with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:metadata:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDateMetadata(quantityType IHKQuantityType, quantity IHKQuantity, startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, metadata foundation.IDictionary /* already interface */) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), quantityType, quantity, startDate, endDate, metadata)
	return rv
}



// Returns a sample containing a numeric measurement with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:metadata:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDateMetadata(quantityType IHKQuantityType, quantity IHKQuantity, startDate foundation.objc.IObject /* cross-framework NSDate */, endDate foundation.objc.IObject /* cross-framework NSDate */, metadata foundation.IDictionary /* already interface */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), quantityType, quantity, startDate, endDate, metadata)
	return rv
}


// A key path for the sample’s count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcount
func (h_ HKQuantitySample) HKPredicateKeyPathCount() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathCount"))
	return rv
}


// The key path for accessing the sample’s quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathquantity
func (h_ HKQuantitySample) HKPredicateKeyPathQuantity() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathQuantity"))
	return rv
}


// The number of quantities contained in this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantitysample/count
func (h_ HKQuantitySample) Count() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](h_.ID, objc.Sel("count"))
	return rv
}


// The number of quantities contained in this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantitysample/count
func (h_ HKQuantitySample) SetCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCount:"), value)
}


// The quantity for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantitysample/quantity
func (h_ HKQuantitySample) Quantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("quantity"))
	return rv
}


// The quantity for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantitysample/quantity
func (h_ HKQuantitySample) SetQuantity(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setQuantity:"), value)
}


// The quantity type for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantitysample/quantitytype
func (h_ HKQuantitySample) QuantityType() IHKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("quantityType"))
	return rv
}


// The quantity type for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantitysample/quantitytype
func (h_ HKQuantitySample) SetQuantityType(value IHKQuantityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setQuantityType:"), value)
}


