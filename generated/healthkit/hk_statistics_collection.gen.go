// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKStatisticsCollection] class.
type IHKStatisticsCollection interface {
	objectivec.IObject
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKStatisticsCollectionClass) Alloc() HKStatisticsCollection {
	rv := objc.Send[HKStatisticsCollection](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




