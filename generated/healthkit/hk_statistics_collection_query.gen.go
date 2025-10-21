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




