// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	TypesToCollect() IHKQuantityType
	SetTypesToCollect(value IHKQuantityType)
	// methods:
}

// A data source that automatically provides live data from an active workout session.


// A data source that automatically provides live data from an active workout session.
//
// [Full Topic]
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



// The quantity type samples that the data source automatically sends to the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutdatasource/typestocollect
func (h_ HKLiveWorkoutDataSource) TypesToCollect() IHKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("typesToCollect"))
	return rv
}


// The quantity type samples that the data source automatically sends to the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkliveworkoutdatasource/typestocollect
func (h_ HKLiveWorkoutDataSource) SetTypesToCollect(value IHKQuantityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setTypesToCollect:"), value)
}



