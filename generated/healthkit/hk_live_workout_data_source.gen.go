// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKLiveWorkoutDataSource] class.
var (
	HKLiveWorkoutDataSourceClass     _HKLiveWorkoutDataSourceClass
	HKLiveWorkoutDataSourceClassOnce sync.Once
)

func getHKLiveWorkoutDataSourceClass() _HKLiveWorkoutDataSourceClass {
	HKLiveWorkoutDataSourceClassOnce.Do(func() {
		HKLiveWorkoutDataSourceClass = _HKLiveWorkoutDataSourceClass{objc.GetClass("HKLiveWorkoutDataSource")}
	})
	return HKLiveWorkoutDataSourceClass
}

type _HKLiveWorkoutDataSourceClass struct {
	class objc.Class
}

// An interface definition for the [HKLiveWorkoutDataSource] class.
type IHKLiveWorkoutDataSource interface {
	objectivec.IObject
	DisableCollectionForType(quantityType HKQuantityType)
	EnableCollectionForTypePredicate(quantityType HKQuantityType, predicate foundation.IPredicate)
	TypesToCollect() unsafe.Pointer
}

// A data source that automatically provides live data from an active workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource
type HKLiveWorkoutDataSource struct {
	objectivec.Object
}

// HKLiveWorkoutDataSourceFrom constructs a [HKLiveWorkoutDataSource] from an unsafe.Pointer.
//
// A data source that automatically provides live data from an active workout session.
func HKLiveWorkoutDataSourceFrom(ptr unsafe.Pointer) HKLiveWorkoutDataSource {
	return HKLiveWorkoutDataSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKLiveWorkoutDataSourceClass) Alloc() HKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKLiveWorkoutDataSourceClass) New() HKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKLiveWorkoutDataSource) Init() HKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKLiveWorkoutDataSource) Autorelease() HKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKLiveWorkoutDataSource creates a new HKLiveWorkoutDataSource instance.
func NewHKLiveWorkoutDataSource() HKLiveWorkoutDataSource {
	return getHKLiveWorkoutDataSourceClass().New()
}




// Creates a new data source based on the provided workout configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/init(healthStore:workoutConfiguration:)
func NewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration(healthStore IHKHealthStore, configuration IHKWorkoutConfiguration) HKLiveWorkoutDataSource {
	instance := getHKLiveWorkoutDataSourceClass().Alloc()
	rv := objc.Send[HKLiveWorkoutDataSource](instance.ID, objc.Sel("initWithHealthStore:workoutConfiguration:"), healthStore, configuration)
	rv.Autorelease()
	return rv
}


// Stops automatically calculating statistics for the quantity type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/disableCollection(for:)
func (h_ HKLiveWorkoutDataSource) DisableCollectionForType(quantityType HKQuantityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableCollectionForType:"), quantityType)
}

// Begins automatically calculating statistics for samples that match the quantity type and predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/enableCollection(for:predicate:)
func (h_ HKLiveWorkoutDataSource) EnableCollectionForTypePredicate(quantityType HKQuantityType, predicate foundation.IPredicate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("enableCollectionForType:predicate:"), quantityType, predicate)
}

// The quantity type samples that the data source automatically sends to the workout builder.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/typesToCollect
func (h_ HKLiveWorkoutDataSource) TypesToCollect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("typesToCollect"))
	return rv
}


