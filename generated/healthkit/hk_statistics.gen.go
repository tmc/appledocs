// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKStatistics] class.
type IHKStatistics interface {
	objectivec.IObject
}

// An object that represents the result of calculating the minimum, maximum, average, or sum over a set of samples from the HealthKit store.
//
// HealthKit creates statistic objects using either a statistics query or a statistics collection query. For the statistics query, it performs the specified calculations over all the samples that match the query. For the statistics collection query, it partitions the matching samples into a set of time intervals and performs the calculations over each interval separately. By default, these queries automatically merge the data from all of your data sources before performing the calculations. If you want to merge the data yourself, you can set the option. You can then request the statistical data for each source separately. When requesting data from a statistics object, your request must match the options you used when creating the query. For example, if you create a query using the option, you must access the results using the method. For more information on calculating statistical data, see Class Reference. To calculate the statistics over a series of time intervals, see the Class Reference.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsClass) Alloc() HKStatistics {
	rv := objc.Send[HKStatistics](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




