// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKLiveWorkoutDataSource */


/* debug [class_header]: Header for HKLiveWorkoutDataSource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKLiveWorkoutDataSource */
// An interface definition for the [HKLiveWorkoutDataSource] class.
type IHKLiveWorkoutDataSource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKLiveWorkoutDataSource */
	// properties:
	TypesToCollect() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKLiveWorkoutDataSource */
	// methods:
	DisableCollectionForType(quantityType IHKQuantityType)
	EnableCollectionForTypePredicate(quantityType IHKQuantityType, predicate foundation.Predicate)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKLiveWorkoutDataSource */
// Alloc allocates a new instance without initialization.
func (hc _HKLiveWorkoutDataSourceClass) Alloc() HKLiveWorkoutDataSource {
	rv := objc.Send[HKLiveWorkoutDataSource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKLiveWorkoutDataSource */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKLiveWorkoutDataSource */

// Creates a new data source based on the provided workout configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/init(healthStore:workoutConfiguration:)
func NewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration(healthStore IHKHealthStore, configuration IHKWorkoutConfiguration) HKLiveWorkoutDataSource {
	instance := getHKLiveWorkoutDataSourceClass().Alloc()
	rv := objc.Send[HKLiveWorkoutDataSource](instance.ID, objc.Sel("initWithHealthStore:workoutConfiguration:"), healthStore, configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKLiveWorkoutDataSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKLiveWorkoutDataSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKLiveWorkoutDataSource */

// Stops automatically calculating statistics for the quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/disableCollection(for:)
func (h_ HKLiveWorkoutDataSource) DisableCollectionForType(quantityType IHKQuantityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableCollectionForType:"), quantityType)
}/* debug [instance_methods/method]: DisableCollectionForType */


// Begins automatically calculating statistics for samples that match the quantity type and predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/enableCollection(for:predicate:)
func (h_ HKLiveWorkoutDataSource) EnableCollectionForTypePredicate(quantityType IHKQuantityType, predicate foundation.Predicate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("enableCollectionForType:predicate:"), quantityType, predicate)
}/* debug [instance_methods/method]: EnableCollectionForTypePredicate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKLiveWorkoutDataSource */

// The quantity type samples that the data source automatically sends to the workout builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLiveWorkoutDataSource/typesToCollect
func (h_ HKLiveWorkoutDataSource) TypesToCollect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("typesToCollect"))
	return rv
}/* debug [instance_properties/getter]: typesToCollect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKLiveWorkoutDataSource */


