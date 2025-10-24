// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKHeartbeatSeriesBuilder */


/* debug [class_header]: Header for HKHeartbeatSeriesBuilder */
// The class instance for the [HKHeartbeatSeriesBuilder] class.
var (
	HKHeartbeatSeriesBuilderClass     _HKHeartbeatSeriesBuilderClass
	HKHeartbeatSeriesBuilderClassOnce sync.Once
)

func getHKHeartbeatSeriesBuilderClass() _HKHeartbeatSeriesBuilderClass {
	HKHeartbeatSeriesBuilderClassOnce.Do(func() {
		HKHeartbeatSeriesBuilderClass = _HKHeartbeatSeriesBuilderClass{objc.GetClass("HKHeartbeatSeriesBuilder")}
	})
	return HKHeartbeatSeriesBuilderClass
}

type _HKHeartbeatSeriesBuilderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKHeartbeatSeriesBuilder */
// An interface definition for the [HKHeartbeatSeriesBuilder] class.
type IHKHeartbeatSeriesBuilder interface {
	IHKSeriesBuilder
	
/* debug [class_interface_properties]: Properties for HKHeartbeatSeriesBuilder */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKHeartbeatSeriesBuilder */
	// methods:
	AddHeartbeatWithTimeIntervalSinceSeriesStartDatePrecededByGapCompletion(timeIntervalSinceStart float64, precededByGap bool, completion unsafe.Pointer)
	AddMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer)
	FinishSeriesWithCompletion(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKHeartbeatSeriesBuilder */
// Alloc allocates a new instance without initialization.
func (hc _HKHeartbeatSeriesBuilderClass) Alloc() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKHeartbeatSeriesBuilderClass) New() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHeartbeatSeriesBuilder) Init() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHeartbeatSeriesBuilder) Autorelease() HKHeartbeatSeriesBuilder {
	rv := objc.Send[HKHeartbeatSeriesBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHeartbeatSeriesBuilder creates a new HKHeartbeatSeriesBuilder instance.
func NewHKHeartbeatSeriesBuilder() HKHeartbeatSeriesBuilder {
	return getHKHeartbeatSeriesBuilderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKHeartbeatSeriesBuilder */
// A builder object for incrementally building a heartbeat series.


// A builder object for incrementally building a heartbeat series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder
type HKHeartbeatSeriesBuilder struct {
	HKSeriesBuilder
}

// HKHeartbeatSeriesBuilderFrom constructs a [HKHeartbeatSeriesBuilder] from an unsafe.Pointer.
//
// A builder object for incrementally building a heartbeat series.
func HKHeartbeatSeriesBuilderFrom(ptr unsafe.Pointer) HKHeartbeatSeriesBuilder {
	return HKHeartbeatSeriesBuilder{
		HKSeriesBuilder: HKSeriesBuilderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKHeartbeatSeriesBuilder */

// Creates a new heartbeat series builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder/init(healthStore:device:start:)
func NewHKHeartbeatSeriesBuilderWithHealthStoreDeviceStartDate(healthStore IHKHealthStore, device IHKDevice, startDate objc.IObject /* cross-framework: NSDate */) HKHeartbeatSeriesBuilder {
	instance := getHKHeartbeatSeriesBuilderClass().Alloc()
	rv := objc.Send[HKHeartbeatSeriesBuilder](instance.ID, objc.Sel("initWithHealthStore:device:startDate:"), healthStore, device, startDate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKHeartbeatSeriesBuilderWithHealthStoreDeviceStartDate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKHeartbeatSeriesBuilder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKHeartbeatSeriesBuilder */

// The maximum number of heartbeats you can add to the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder/maximumCount
func (hc _HKHeartbeatSeriesBuilderClass) MaximumCount() uint {
	rv := objc.Send[uint](objc.ID(hc.class), objc.Sel("maximumCount"))
	return rv
}/* debug [class_properties_class/property]: maximumCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKHeartbeatSeriesBuilder */

// Adds a heartbeat to the series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder/addHeartbeatWithTimeInterval(sinceSeriesStartDate:precededByGap:completion:)
func (h_ HKHeartbeatSeriesBuilder) AddHeartbeatWithTimeIntervalSinceSeriesStartDatePrecededByGapCompletion(timeIntervalSinceStart float64, precededByGap bool, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addHeartbeatWithTimeIntervalSinceSeriesStartDate:precededByGap:completion:"), timeIntervalSinceStart, precededByGap, completion)
}/* debug [instance_methods/method]: AddHeartbeatWithTimeIntervalSinceSeriesStartDatePrecededByGapCompletion */


// Adds metadata to the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder/addMetadata(_:completion:)
func (h_ HKHeartbeatSeriesBuilder) AddMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addMetadata:completion:"), metadata, completion)
}/* debug [instance_methods/method]: AddMetadataCompletion */


// Finalizes the series and returns the resulting heartbeat series sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder/finishSeries(completion:)
func (h_ HKHeartbeatSeriesBuilder) FinishSeriesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("finishSeriesWithCompletion:"), completion)
}/* debug [instance_methods/method]: FinishSeriesWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKHeartbeatSeriesBuilder */

// The maximum number of heartbeats you can add to the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartbeatSeriesBuilder/maximumCount
func (h_ HKHeartbeatSeriesBuilder) MaximumCount() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("maximumCount"))
	return rv
}/* debug [instance_properties/getter]: maximumCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKHeartbeatSeriesBuilder */


