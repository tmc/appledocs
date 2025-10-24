// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKSample */


/* debug [class_header]: Header for HKSample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSample */
// An interface definition for the [HKSample] class.
type IHKSample interface {
	IHKObject
	
/* debug [class_interface_properties]: Properties for HKSample */
	// properties:
	EndDate() objc.IObject /* cross-framework: NSDate */
	HasUndeterminedDuration() bool
	SampleType() IHKSampleType
	StartDate() objc.IObject /* cross-framework: NSDate */
	HKPredicateKeyPathEndDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathStartDate() objc.IObject /* cross-framework: NSString */
	HKSampleSortIdentifierEndDate() objc.IObject /* cross-framework: NSString */
	HKSampleSortIdentifierStartDate() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSample */
// Alloc allocates a new instance without initialization.
func (hc _HKSampleClass) Alloc() HKSample {
	rv := objc.Send[HKSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSample */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSample *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSample */

// The sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/endDate
func (h_ HKSample) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// Indicates whether the sample has an unknown duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/hasUndeterminedDuration
func (h_ HKSample) HasUndeterminedDuration() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("hasUndeterminedDuration"))
	return rv
}/* debug [instance_properties/getter]: hasUndeterminedDuration */


// The sample type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/sampleType
func (h_ HKSample) SampleType() IHKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("sampleType"))
	return rv
}/* debug [instance_properties/getter]: sampleType */


// The sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSample/startDate
func (h_ HKSample) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The key path for accessing the sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathenddate
func (h_ HKSample) HKPredicateKeyPathEndDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathEndDate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathEndDate */


// The key path for accessing the sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathstartdate
func (h_ HKSample) HKPredicateKeyPathStartDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathStartDate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathStartDate */


// A constant for sorting samples based on their end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplesortidentifierenddate
func (h_ HKSample) HKSampleSortIdentifierEndDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKSampleSortIdentifierEndDate"))
	return rv
}/* debug [instance_properties/getter]: HKSampleSortIdentifierEndDate */


// A constant for sorting samples based on their start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksamplesortidentifierstartdate
func (h_ HKSample) HKSampleSortIdentifierStartDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKSampleSortIdentifierStartDate"))
	return rv
}/* debug [instance_properties/getter]: HKSampleSortIdentifierStartDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSample */



