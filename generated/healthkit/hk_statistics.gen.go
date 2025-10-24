// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKStatistics */


/* debug [class_header]: Header for HKStatistics */
// The class instance for the [HKStatistics] class.
var (
	HKStatisticsClass     _HKStatisticsClass
	HKStatisticsClassOnce sync.Once
)

func getHKStatisticsClass() _HKStatisticsClass {
	HKStatisticsClassOnce.Do(func() {
		HKStatisticsClass = _HKStatisticsClass{objc.GetClass("HKStatistics")}
	})
	return HKStatisticsClass
}

type _HKStatisticsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKStatistics */
// An interface definition for the [HKStatistics] class.
type IHKStatistics interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKStatistics */
	// properties:
	EndDate() objc.IObject /* cross-framework: NSDate */
	QuantityType() IHKQuantityType
	Sources() []HKSource
	StartDate() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKStatistics */
	// methods:
	AverageQuantity() IHKQuantity
	AverageQuantityForSource(source IHKSource) IHKQuantity
	Duration() IHKQuantity
	DurationForSource(source IHKSource) IHKQuantity
	MaximumQuantity() IHKQuantity
	MaximumQuantityForSource(source IHKSource) IHKQuantity
	MinimumQuantity() IHKQuantity
	MinimumQuantityForSource(source IHKSource) IHKQuantity
	MostRecentQuantity() IHKQuantity
	MostRecentQuantityForSource(source IHKSource) IHKQuantity
	MostRecentQuantityDateInterval() foundation.DateInterval
	MostRecentQuantityDateIntervalForSource(source IHKSource) foundation.DateInterval
	SumQuantity() IHKQuantity
	SumQuantityForSource(source IHKSource) IHKQuantity
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKStatistics */
// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsClass) Alloc() HKStatistics {
	rv := objc.Send[HKStatistics](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKStatisticsClass) New() HKStatistics {
	rv := objc.Send[HKStatistics](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStatistics) Init() HKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStatistics) Autorelease() HKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStatistics creates a new HKStatistics instance.
func NewHKStatistics() HKStatistics {
	return getHKStatisticsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKStatistics */
// An object that represents the result of calculating the minimum, maximum, average, or sum over a set of samples from the HealthKit store.
//
// HealthKit creates statistic objects using either a statistics query or a statistics collection query. For the statistics query, it performs the specified calculations over all the samples that match the query. For the statistics collection query, it partitions the matching samples into a set of time intervals and performs the calculations over each interval separately. By default, these queries automatically merge the data from all of your data sources before performing the calculations. If you want to merge the data yourself, you can set the option. You can then request the statistical data for each source separately. When requesting data from a statistics object, your request must match the options you used when creating the query. For example, if you create a query using the option, you must access the results using the method. For more information on calculating statistical data, see Class Reference. To calculate the statistics over a series of time intervals, see the Class Reference.


// An object that represents the result of calculating the minimum, maximum, average, or sum over a set of samples from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics
type HKStatistics struct {
	objectivec.Object
}

// HKStatisticsFrom constructs a [HKStatistics] from an unsafe.Pointer.
//
// An object that represents the result of calculating the minimum, maximum, average, or sum over a set of samples from the HealthKit store.
func HKStatisticsFrom(ptr unsafe.Pointer) HKStatistics {
	return HKStatistics{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKStatistics *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKStatistics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKStatistics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKStatistics */

// Returns the average value from all the samples that match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/averageQuantity()
func (h_ HKStatistics) AverageQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("averageQuantity"))
	return rv
}/* debug [instance_methods/method]: AverageQuantity */


// Returns the average value from all the samples that match the query and that were created by the specified source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/averageQuantity(for:)
func (h_ HKStatistics) AverageQuantityForSource(source IHKSource) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("averageQuantityForSource:"), source)
	return rv
}/* debug [instance_methods/method]: AverageQuantityForSource */


// Returns the total duration covering all the samples that match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/duration()
func (h_ HKStatistics) Duration() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_methods/method]: Duration */


// Returns the total duration covering all the samples created by the specified source that also match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/duration(for:)
func (h_ HKStatistics) DurationForSource(source IHKSource) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("durationForSource:"), source)
	return rv
}/* debug [instance_methods/method]: DurationForSource */


// Returns the maximum value from all the samples that match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/maximumQuantity()
func (h_ HKStatistics) MaximumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("maximumQuantity"))
	return rv
}/* debug [instance_methods/method]: MaximumQuantity */


// Returns the maximum value from all the samples that match the query and that were created by the specified source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/maximumQuantity(for:)
func (h_ HKStatistics) MaximumQuantityForSource(source IHKSource) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("maximumQuantityForSource:"), source)
	return rv
}/* debug [instance_methods/method]: MaximumQuantityForSource */


// Returns the minimum value from all the samples that match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/minimumQuantity()
func (h_ HKStatistics) MinimumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("minimumQuantity"))
	return rv
}/* debug [instance_methods/method]: MinimumQuantity */


// Returns the minimum value from all the samples that match the query and that were created by the specified source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/minimumQuantity(for:)
func (h_ HKStatistics) MinimumQuantityForSource(source IHKSource) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("minimumQuantityForSource:"), source)
	return rv
}/* debug [instance_methods/method]: MinimumQuantityForSource */


// Returns the most recent value from all the samples that match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/mostRecentQuantity()
func (h_ HKStatistics) MostRecentQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("mostRecentQuantity"))
	return rv
}/* debug [instance_methods/method]: MostRecentQuantity */


// Returns the most recent value from all the samples that match the query and were created by the specified source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/mostRecentQuantity(for:)
func (h_ HKStatistics) MostRecentQuantityForSource(source IHKSource) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("mostRecentQuantityForSource:"), source)
	return rv
}/* debug [instance_methods/method]: MostRecentQuantityForSource */


// Returns the date interval of the most recent sample that matches the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/mostRecentQuantityDateInterval()
func (h_ HKStatistics) MostRecentQuantityDateInterval() foundation.DateInterval {
	rv := objc.Send[foundation.DateInterval](h_.ID, objc.Sel("mostRecentQuantityDateInterval"))
	return rv
}/* debug [instance_methods/method]: MostRecentQuantityDateInterval */


// Returns the date interval of the most recent sample that matches the query and was created by the specified source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/mostRecentQuantityDateInterval(for:)
func (h_ HKStatistics) MostRecentQuantityDateIntervalForSource(source IHKSource) foundation.DateInterval {
	rv := objc.Send[foundation.DateInterval](h_.ID, objc.Sel("mostRecentQuantityDateIntervalForSource:"), source)
	return rv
}/* debug [instance_methods/method]: MostRecentQuantityDateIntervalForSource */


// Returns the sum of all the samples that match the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/sumQuantity()
func (h_ HKStatistics) SumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sumQuantity"))
	return rv
}/* debug [instance_methods/method]: SumQuantity */


// Returns the sum of all the samples that match the query and that were created by the specified source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/sumQuantity(for:)
func (h_ HKStatistics) SumQuantityForSource(source IHKSource) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sumQuantityForSource:"), source)
	return rv
}/* debug [instance_methods/method]: SumQuantityForSource */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKStatistics */

// The end of the time period included in these statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/endDate
func (h_ HKStatistics) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The quantity type of the samples used to calculate these statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/quantityType
func (h_ HKStatistics) QuantityType() IHKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("quantityType"))
	return rv
}/* debug [instance_properties/getter]: quantityType */


// An array containing all the sources contributing to these statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/sources
func (h_ HKStatistics) Sources() []HKSource {
	rv := objc.Send[[]HKSource](h_.ID, objc.Sel("sources"))
	return rv
}/* debug [instance_properties/getter]: sources */


// The start of the time period included in these statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatistics/startDate
func (h_ HKStatistics) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKStatistics */



