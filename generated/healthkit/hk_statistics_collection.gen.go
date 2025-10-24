// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKStatisticsCollection */


/* debug [class_header]: Header for HKStatisticsCollection */
// The class instance for the [HKStatisticsCollection] class.
var (
	HKStatisticsCollectionClass     _HKStatisticsCollectionClass
	HKStatisticsCollectionClassOnce sync.Once
)

func getHKStatisticsCollectionClass() _HKStatisticsCollectionClass {
	HKStatisticsCollectionClassOnce.Do(func() {
		HKStatisticsCollectionClass = _HKStatisticsCollectionClass{objc.GetClass("HKStatisticsCollection")}
	})
	return HKStatisticsCollectionClass
}

type _HKStatisticsCollectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKStatisticsCollection */
// An interface definition for the [HKStatisticsCollection] class.
type IHKStatisticsCollection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKStatisticsCollection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKStatisticsCollection */
	// methods:
	EnumerateStatisticsFromDateToDateWithBlock(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, block unsafe.Pointer)
	Sources() unsafe.Pointer
	Statistics() []HKStatistics
	StatisticsForDate(date objc.IObject /* cross-framework: NSDate */) IHKStatistics
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKStatisticsCollection */
// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsCollectionClass) Alloc() HKStatisticsCollection {
	rv := objc.Send[HKStatisticsCollection](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKStatisticsCollectionClass) New() HKStatisticsCollection {
	rv := objc.Send[HKStatisticsCollection](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStatisticsCollection) Init() HKStatisticsCollection {
	rv := objc.Send[HKStatisticsCollection](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStatisticsCollection) Autorelease() HKStatisticsCollection {
	rv := objc.Send[HKStatisticsCollection](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStatisticsCollection creates a new HKStatisticsCollection instance.
func NewHKStatisticsCollection() HKStatisticsCollection {
	return getHKStatisticsCollectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKStatisticsCollection */
// An object that manages a collection of statistics, representing the results calculated over separate time intervals.
//
// For more information on statistics objects, see . For more information on calculating statistics over consecutive time intervals, see .


// An object that manages a collection of statistics, representing the results calculated over separate time intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollection
type HKStatisticsCollection struct {
	objectivec.Object
}

// HKStatisticsCollectionFrom constructs a [HKStatisticsCollection] from an unsafe.Pointer.
//
// An object that manages a collection of statistics, representing the results calculated over separate time intervals.
func HKStatisticsCollectionFrom(ptr unsafe.Pointer) HKStatisticsCollection {
	return HKStatisticsCollection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKStatisticsCollection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKStatisticsCollection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKStatisticsCollection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKStatisticsCollection */

// Enumerates the statistics objects for all the time intervals from the start date until the end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollection/enumerateStatistics(from:to:with:)
func (h_ HKStatisticsCollection) EnumerateStatisticsFromDateToDateWithBlock(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, block unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("enumerateStatisticsFromDate:toDate:withBlock:"), startDate, endDate, block)
}/* debug [instance_methods/method]: EnumerateStatisticsFromDateToDateWithBlock */


// Returns a set containing all the sources that had samples matched by the statistics collection query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollection/sources()
func (h_ HKStatisticsCollection) Sources() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sources"))
	return rv
}/* debug [instance_methods/method]: Sources */


// Returns an array of statistics objects representing the populated time intervals covered by the statistics collection query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollection/statistics()
func (h_ HKStatisticsCollection) Statistics() []HKStatistics {
	rv := objc.Send[[]HKStatistics](h_.ID, objc.Sel("statistics"))
	return rv
}/* debug [instance_methods/method]: Statistics */


// Returns the statistics object for the time interval that contains the provided date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollection/statistics(for:)
func (h_ HKStatisticsCollection) StatisticsForDate(date objc.IObject /* cross-framework: NSDate */) IHKStatistics {
	rv := objc.Send[HKStatistics](h_.ID, objc.Sel("statisticsForDate:"), date)
	return rv
}/* debug [instance_methods/method]: StatisticsForDate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKStatisticsCollection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKStatisticsCollection */



