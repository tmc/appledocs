// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQuantitySeriesSampleBuilder */


/* debug [class_header]: Header for HKQuantitySeriesSampleBuilder */
// The class instance for the [HKQuantitySeriesSampleBuilder] class.
var (
	HKQuantitySeriesSampleBuilderClass     _HKQuantitySeriesSampleBuilderClass
	HKQuantitySeriesSampleBuilderClassOnce sync.Once
)

func getHKQuantitySeriesSampleBuilderClass() _HKQuantitySeriesSampleBuilderClass {
	HKQuantitySeriesSampleBuilderClassOnce.Do(func() {
		HKQuantitySeriesSampleBuilderClass = _HKQuantitySeriesSampleBuilderClass{objc.GetClass("HKQuantitySeriesSampleBuilder")}
	})
	return HKQuantitySeriesSampleBuilderClass
}

type _HKQuantitySeriesSampleBuilderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQuantitySeriesSampleBuilder */
// An interface definition for the [HKQuantitySeriesSampleBuilder] class.
type IHKQuantitySeriesSampleBuilder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKQuantitySeriesSampleBuilder */
	// properties:
	Device() IHKDevice
	QuantityType() IHKQuantityType
	StartDate() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQuantitySeriesSampleBuilder */
	// methods:
	Discard()
	FinishSeriesWithMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer)
	FinishSeriesWithMetadataEndDateCompletion(metadata foundation.IDictionary, endDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer)
	InsertQuantityDateError(quantity IHKQuantity, date objc.IObject /* cross-framework: NSDate */, error_ objectivec.IObject) bool
	InsertQuantityDateIntervalError(quantity IHKQuantity, dateInterval foundation.DateInterval, error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQuantitySeriesSampleBuilder */
// Alloc allocates a new instance without initialization.
func (hc _HKQuantitySeriesSampleBuilderClass) Alloc() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKQuantitySeriesSampleBuilderClass) New() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantitySeriesSampleBuilder) Init() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantitySeriesSampleBuilder) Autorelease() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantitySeriesSampleBuilder creates a new HKQuantitySeriesSampleBuilder instance.
func NewHKQuantitySeriesSampleBuilder() HKQuantitySeriesSampleBuilder {
	return getHKQuantitySeriesSampleBuilderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQuantitySeriesSampleBuilder */
// A builder object for incrementally building a sample that contains multiple quantities.


// A builder object for incrementally building a sample that contains multiple quantities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder
type HKQuantitySeriesSampleBuilder struct {
	objectivec.Object
}

// HKQuantitySeriesSampleBuilderFrom constructs a [HKQuantitySeriesSampleBuilder] from an unsafe.Pointer.
//
// A builder object for incrementally building a sample that contains multiple quantities.
func HKQuantitySeriesSampleBuilderFrom(ptr unsafe.Pointer) HKQuantitySeriesSampleBuilder {
	return HKQuantitySeriesSampleBuilder{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQuantitySeriesSampleBuilder */

// Creates a new quantity series builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/init(healthStore:quantityType:startDate:device:)
func NewHKQuantitySeriesSampleBuilderWithHealthStoreQuantityTypeStartDateDevice(healthStore IHKHealthStore, quantityType IHKQuantityType, startDate objc.IObject /* cross-framework: NSDate */, device IHKDevice) HKQuantitySeriesSampleBuilder {
	instance := getHKQuantitySeriesSampleBuilderClass().Alloc()
	rv := objc.Send[HKQuantitySeriesSampleBuilder](instance.ID, objc.Sel("initWithHealthStore:quantityType:startDate:device:"), healthStore, quantityType, startDate, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantitySeriesSampleBuilderWithHealthStoreQuantityTypeStartDateDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQuantitySeriesSampleBuilder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQuantitySeriesSampleBuilder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQuantitySeriesSampleBuilder */

// Discards all previously collected data and invalidates the builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/discard()
func (h_ HKQuantitySeriesSampleBuilder) Discard() {
	objc.Send[objc.ID](h_.ID, objc.Sel("discard"))
}/* debug [instance_methods/method]: Discard */


// Finalizes the series and returns the resulting quantity samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/finishSeries(metadata:completion:)
func (h_ HKQuantitySeriesSampleBuilder) FinishSeriesWithMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("finishSeriesWithMetadata:completion:"), metadata, completion)
}/* debug [instance_methods/method]: FinishSeriesWithMetadataCompletion */


// Finalizes the series with the provided end date, and returns the resulting quantity samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/finishSeries(metadata:endDate:completion:)
func (h_ HKQuantitySeriesSampleBuilder) FinishSeriesWithMetadataEndDateCompletion(metadata foundation.IDictionary, endDate objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("finishSeriesWithMetadata:endDate:completion:"), metadata, endDate, completion)
}/* debug [instance_methods/method]: FinishSeriesWithMetadataEndDateCompletion */


// Adds a new quantity to the series at the provided date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/insert(_:at:)
func (h_ HKQuantitySeriesSampleBuilder) InsertQuantityDateError(quantity IHKQuantity, date objc.IObject /* cross-framework: NSDate */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("insertQuantity:date:error:"), quantity, date, error_)
	return rv
}/* debug [instance_methods/method]: InsertQuantityDateError */


// Adds a new quantity to the series with the provided date interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/insert(_:for:)
func (h_ HKQuantitySeriesSampleBuilder) InsertQuantityDateIntervalError(quantity IHKQuantity, dateInterval foundation.DateInterval, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("insertQuantity:dateInterval:error:"), quantity, dateInterval, error_)
	return rv
}/* debug [instance_methods/method]: InsertQuantityDateIntervalError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQuantitySeriesSampleBuilder */

// The device providing the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/device
func (h_ HKQuantitySeriesSampleBuilder) Device() IHKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The quantity type for the series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/quantityType
func (h_ HKQuantitySeriesSampleBuilder) QuantityType() IHKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("quantityType"))
	return rv
}/* debug [instance_properties/getter]: quantityType */


// The starting date and time for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder/startDate
func (h_ HKQuantitySeriesSampleBuilder) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQuantitySeriesSampleBuilder */


