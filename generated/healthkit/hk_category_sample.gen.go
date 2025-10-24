// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKCategorySample */


/* debug [class_header]: Header for HKCategorySample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCategorySample */
// An interface definition for the [HKCategorySample] class.
type IHKCategorySample interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKCategorySample */
	// properties:
	CategoryType() IHKCategoryType
	Value() int
	HKPredicateKeyPathCategoryValue() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCategorySample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCategorySample */
// Alloc allocates a new instance without initialization.
func (hc _HKCategorySampleClass) Alloc() HKCategorySample {
	rv := objc.Send[HKCategorySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCategorySample */
// A sample with values from a short list of possible values.
//
// You can use category samples to record data associated with a . The value for the sample must come from the appropriate category value enumeration. Each category type uses its own enumeration. Individual samples represent a value and time period. Samples with different values may have overlapping time intervals. The class is a concrete subclass of the class. Category samples are immutable: You set the sample’s properties when you create it, and they can’t change.


// A sample with values from a short list of possible values.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCategorySample */

// Creates a newly instantiated category sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/init(type:value:start:end:)
func NewHKCategorySampleWithTypeValueStartDateEndDate(type_ IHKCategoryType, value int, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) HKCategorySample {
	rv := objc.Send[HKCategorySample](objc.ID(getHKCategorySampleClass().class), objc.Sel("categorySampleWithType:value:startDate:endDate:"), type_, value, startDate, endDate)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCategorySampleWithTypeValueStartDateEndDate */


// Creates a newly instantiated category sample including the provided device and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/init(type:value:start:end:device:metadata:)
func NewHKCategorySampleWithTypeValueStartDateEndDateDeviceMetadata(type_ IHKCategoryType, value int, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) HKCategorySample {
	rv := objc.Send[HKCategorySample](objc.ID(getHKCategorySampleClass().class), objc.Sel("categorySampleWithType:value:startDate:endDate:device:metadata:"), type_, value, startDate, endDate, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCategorySampleWithTypeValueStartDateEndDateDeviceMetadata */


// Creates a newly instantiated category sample with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/init(type:value:start:end:metadata:)
func NewHKCategorySampleWithTypeValueStartDateEndDateMetadata(type_ IHKCategoryType, value int, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) HKCategorySample {
	rv := objc.Send[HKCategorySample](objc.ID(getHKCategorySampleClass().class), objc.Sel("categorySampleWithType:value:startDate:endDate:metadata:"), type_, value, startDate, endDate, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCategorySampleWithTypeValueStartDateEndDateMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCategorySample */

// Creates a newly instantiated category sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/init(type:value:start:end:)
func (hc _HKCategorySampleClass) CategorySampleWithTypeValueStartDateEndDate(type_ IHKCategoryType, value int, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("categorySampleWithType:value:startDate:endDate:"), type_, value, startDate, endDate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CategorySampleWithTypeValueStartDateEndDate) */


// Creates a newly instantiated category sample including the provided device and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/init(type:value:start:end:device:metadata:)
func (hc _HKCategorySampleClass) CategorySampleWithTypeValueStartDateEndDateDeviceMetadata(type_ IHKCategoryType, value int, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("categorySampleWithType:value:startDate:endDate:device:metadata:"), type_, value, startDate, endDate, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CategorySampleWithTypeValueStartDateEndDateDeviceMetadata) */


// Creates a newly instantiated category sample with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/init(type:value:start:end:metadata:)
func (hc _HKCategorySampleClass) CategorySampleWithTypeValueStartDateEndDateMetadata(type_ IHKCategoryType, value int, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("categorySampleWithType:value:startDate:endDate:metadata:"), type_, value, startDate, endDate, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CategorySampleWithTypeValueStartDateEndDateMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCategorySample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCategorySample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCategorySample */

// The category type for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/categoryType
func (h_ HKCategorySample) CategoryType() IHKCategoryType {
	rv := objc.Send[HKCategoryType](h_.ID, objc.Sel("categoryType"))
	return rv
}/* debug [instance_properties/getter]: categoryType */


// The category value for this sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategorySample/value
func (h_ HKCategorySample) Value() int {
	rv := objc.Send[int](h_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The key path for accessing the category sample’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcategoryvalue
func (h_ HKCategorySample) HKPredicateKeyPathCategoryValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCategoryValue"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCategoryValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCategorySample */


