// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKStatisticsCollectionQuery] class.
var (
	HKStatisticsCollectionQueryClass     _HKStatisticsCollectionQueryClass
	HKStatisticsCollectionQueryClassOnce sync.Once
)

func getHKStatisticsCollectionQueryClass() _HKStatisticsCollectionQueryClass {
	HKStatisticsCollectionQueryClassOnce.Do(func() {
		HKStatisticsCollectionQueryClass = _HKStatisticsCollectionQueryClass{objc.GetClass("HKStatisticsCollectionQuery")}
	})
	return HKStatisticsCollectionQueryClass
}

type _HKStatisticsCollectionQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKStatisticsCollectionQuery] class.
type IHKStatisticsCollectionQuery interface {
	IHKQuery
}

// A query that performs multiple statistics queries over a series of fixed-length time intervals.
//
// Statistics collection queries are often used to produce data for graphs and charts. For example, you might create a statistics collection query that calculates the total number of steps for each day or the average heart rate for each hour. Like observer queries, collection queries can also act as long-running queries, receiving updates when the HealthKit store’s content changes. Statistics collection queries are mostly immutable. You can assign the query’s and properties after instantiating the object. You must set all other properties when you instantiate the object, and they can’t change. For more information about statistics queries, see .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsCollectionQuery
type HKStatisticsCollectionQuery struct {
	HKQuery
}

// HKStatisticsCollectionQueryFrom constructs a [HKStatisticsCollectionQuery] from an unsafe.Pointer.
//
// A query that performs multiple statistics queries over a series of fixed-length time intervals.
func HKStatisticsCollectionQueryFrom(ptr unsafe.Pointer) HKStatisticsCollectionQuery {
	return HKStatisticsCollectionQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsCollectionQueryClass) Alloc() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKStatisticsCollectionQueryClass) New() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStatisticsCollectionQuery) Init() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStatisticsCollectionQuery) Autorelease() HKStatisticsCollectionQuery {
	rv := objc.Send[HKStatisticsCollectionQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStatisticsCollectionQuery creates a new HKStatisticsCollectionQuery instance.
func NewHKStatisticsCollectionQuery() HKStatisticsCollectionQuery {
	return getHKStatisticsCollectionQueryClass().New()
}


// The anchor date for the collection’s time intervals.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/anchordate
func (h_ HKStatisticsCollectionQuery) AnchorDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("anchorDate"))
	return rv
}


// SetAnchorDate sets the value of the anchorDate property.
// The anchor date for the collection’s time intervals.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/anchordate
func (h_ HKStatisticsCollectionQuery) SetAnchorDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAnchorDate:"), value)
}

// The results handler for the query’s initial results.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/initialresultshandler
func (h_ HKStatisticsCollectionQuery) InitialResultsHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("initialResultsHandler"))
	return rv
}


// SetInitialResultsHandler sets the value of the initialResultsHandler property.
// The results handler for the query’s initial results.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/initialresultshandler
func (h_ HKStatisticsCollectionQuery) SetInitialResultsHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setInitialResultsHandler:"), value)
}

// The date components that define the time interval for each statistics object in the collection.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/intervalcomponents
func (h_ HKStatisticsCollectionQuery) IntervalComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("intervalComponents"))
	return rv
}


// SetIntervalComponents sets the value of the intervalComponents property.
// The date components that define the time interval for each statistics object in the collection.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/intervalcomponents
func (h_ HKStatisticsCollectionQuery) SetIntervalComponents(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIntervalComponents:"), value)
}

// A list of options that define the type of statistical calculations performed and the way in which data from multiple sources are merged.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/options
func (h_ HKStatisticsCollectionQuery) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// A list of options that define the type of statistical calculations performed and the way in which data from multiple sources are merged.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/options
func (h_ HKStatisticsCollectionQuery) SetOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setOptions:"), value)
}

// The results handler for monitoring updates to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/statisticsupdatehandler
func (h_ HKStatisticsCollectionQuery) StatisticsUpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("statisticsUpdateHandler"))
	return rv
}


// SetStatisticsUpdateHandler sets the value of the statisticsUpdateHandler property.
// The results handler for monitoring updates to the HealthKit store.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstatisticscollectionquery/statisticsupdatehandler
func (h_ HKStatisticsCollectionQuery) SetStatisticsUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStatisticsUpdateHandler:"), value)
}



