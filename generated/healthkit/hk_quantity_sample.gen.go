// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQuantitySample */


/* debug [class_header]: Header for HKQuantitySample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQuantitySample */
// An interface definition for the [HKQuantitySample] class.
type IHKQuantitySample interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKQuantitySample */
	// properties:
	Count() int
	Quantity() IHKQuantity
	QuantityType() IHKQuantityType
	HKPredicateKeyPathCount() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathQuantity() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQuantitySample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQuantitySample */
// Alloc allocates a new instance without initialization.
func (hc _HKQuantitySampleClass) Alloc() HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQuantitySample */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQuantitySample */

// Returns a sample containing a numeric measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDate(quantityType IHKQuantityType, quantity IHKQuantity, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:"), quantityType, quantity, startDate, endDate)
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantitySampleWithTypeQuantityStartDateEndDate */


// Returns a sample containing a numeric measurement with the provided device and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:device:metadata:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDateDeviceMetadata(quantityType IHKQuantityType, quantity IHKQuantity, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:device:metadata:"), quantityType, quantity, startDate, endDate, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantitySampleWithTypeQuantityStartDateEndDateDeviceMetadata */


// Returns a sample containing a numeric measurement with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:metadata:)
func NewHKQuantitySampleWithTypeQuantityStartDateEndDateMetadata(quantityType IHKQuantityType, quantity IHKQuantity, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) HKQuantitySample {
	rv := objc.Send[HKQuantitySample](objc.ID(getHKQuantitySampleClass().class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), quantityType, quantity, startDate, endDate, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantitySampleWithTypeQuantityStartDateEndDateMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQuantitySample */

// Returns a sample containing a numeric measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDate(quantityType IHKQuantityType, quantity IHKQuantity, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:"), quantityType, quantity, startDate, endDate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuantitySampleWithTypeQuantityStartDateEndDate) */


// Returns a sample containing a numeric measurement with the provided device and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:device:metadata:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDateDeviceMetadata(quantityType IHKQuantityType, quantity IHKQuantity, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:device:metadata:"), quantityType, quantity, startDate, endDate, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuantitySampleWithTypeQuantityStartDateEndDateDeviceMetadata) */


// Returns a sample containing a numeric measurement with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/init(type:quantity:start:end:metadata:)
func (hc _HKQuantitySampleClass) QuantitySampleWithTypeQuantityStartDateEndDateMetadata(quantityType IHKQuantityType, quantity IHKQuantity, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("quantitySampleWithType:quantity:startDate:endDate:metadata:"), quantityType, quantity, startDate, endDate, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuantitySampleWithTypeQuantityStartDateEndDateMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQuantitySample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQuantitySample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQuantitySample */

// The number of quantities contained in this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/count
func (h_ HKQuantitySample) Count() int {
	rv := objc.Send[int](h_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// The quantity for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/quantity
func (h_ HKQuantitySample) Quantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("quantity"))
	return rv
}/* debug [instance_properties/getter]: quantity */


// The quantity type for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySample/quantityType
func (h_ HKQuantitySample) QuantityType() IHKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("quantityType"))
	return rv
}/* debug [instance_properties/getter]: quantityType */


// A key path for the sample’s count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcount
func (h_ HKQuantitySample) HKPredicateKeyPathCount() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCount"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCount */


// The key path for accessing the sample’s quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathquantity
func (h_ HKQuantitySample) HKPredicateKeyPathQuantity() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathQuantity"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathQuantity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQuantitySample */


