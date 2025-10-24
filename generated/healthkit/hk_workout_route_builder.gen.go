// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKWorkoutRouteBuilder */


/* debug [class_header]: Header for HKWorkoutRouteBuilder */
// The class instance for the [HKWorkoutRouteBuilder] class.
var (
	HKWorkoutRouteBuilderClass     _HKWorkoutRouteBuilderClass
	HKWorkoutRouteBuilderClassOnce sync.Once
)

func getHKWorkoutRouteBuilderClass() _HKWorkoutRouteBuilderClass {
	HKWorkoutRouteBuilderClassOnce.Do(func() {
		HKWorkoutRouteBuilderClass = _HKWorkoutRouteBuilderClass{objc.GetClass("HKWorkoutRouteBuilder")}
	})
	return HKWorkoutRouteBuilderClass
}

type _HKWorkoutRouteBuilderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKWorkoutRouteBuilder */
// An interface definition for the [HKWorkoutRouteBuilder] class.
type IHKWorkoutRouteBuilder interface {
	IHKSeriesBuilder
	
/* debug [class_interface_properties]: Properties for HKWorkoutRouteBuilder */
	// properties:
	HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKWorkoutRouteBuilder */
	// methods:
	AddMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer)
	FinishRouteWithWorkoutMetadataCompletion(workout IHKWorkout, metadata foundation.IDictionary, completion unsafe.Pointer)
	InsertRouteDataCompletion(routeData []corelocation.Location, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKWorkoutRouteBuilder */
// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutRouteBuilderClass) Alloc() HKWorkoutRouteBuilder {
	rv := objc.Send[HKWorkoutRouteBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKWorkoutRouteBuilderClass) New() HKWorkoutRouteBuilder {
	rv := objc.Send[HKWorkoutRouteBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWorkoutRouteBuilder) Init() HKWorkoutRouteBuilder {
	rv := objc.Send[HKWorkoutRouteBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWorkoutRouteBuilder) Autorelease() HKWorkoutRouteBuilder {
	rv := objc.Send[HKWorkoutRouteBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWorkoutRouteBuilder creates a new HKWorkoutRouteBuilder instance.
func NewHKWorkoutRouteBuilder() HKWorkoutRouteBuilder {
	return getHKWorkoutRouteBuilderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKWorkoutRouteBuilder */
// A builder object that incrementally constructs a workout route.
//
// To create a workout route, use to instantiate a , and provide it with location data throughout the workout. After the workout ends, call the builder’s method to construct the route. Instantiating a directly is discouraged. For detailed instructions, see .


// A builder object that incrementally constructs a workout route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteBuilder
type HKWorkoutRouteBuilder struct {
	HKSeriesBuilder
}

// HKWorkoutRouteBuilderFrom constructs a [HKWorkoutRouteBuilder] from an unsafe.Pointer.
//
// A builder object that incrementally constructs a workout route.
func HKWorkoutRouteBuilderFrom(ptr unsafe.Pointer) HKWorkoutRouteBuilder {
	return HKWorkoutRouteBuilder{
		HKSeriesBuilder: HKSeriesBuilderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKWorkoutRouteBuilder */

// Creates and returns a new workout route builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteBuilder/init(healthStore:device:)
func NewHKWorkoutRouteBuilderWithHealthStoreDevice(healthStore IHKHealthStore, device IHKDevice) HKWorkoutRouteBuilder {
	instance := getHKWorkoutRouteBuilderClass().Alloc()
	rv := objc.Send[HKWorkoutRouteBuilder](instance.ID, objc.Sel("initWithHealthStore:device:"), healthStore, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKWorkoutRouteBuilderWithHealthStoreDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKWorkoutRouteBuilder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKWorkoutRouteBuilder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKWorkoutRouteBuilder */

// Adds metadata to the builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteBuilder/addMetadata(_:completion:)
func (h_ HKWorkoutRouteBuilder) AddMetadataCompletion(metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addMetadata:completion:"), metadata, completion)
}/* debug [instance_methods/method]: AddMetadataCompletion */


// Creates, saves, and associates the route with the provided workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteBuilder/finishRoute(with:metadata:completion:)
func (h_ HKWorkoutRouteBuilder) FinishRouteWithWorkoutMetadataCompletion(workout IHKWorkout, metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("finishRouteWithWorkout:metadata:completion:"), workout, metadata, completion)
}/* debug [instance_methods/method]: FinishRouteWithWorkoutMetadataCompletion */


// Adds route data to the builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutRouteBuilder/insertRouteData(_:completion:)
func (h_ HKWorkoutRouteBuilder) InsertRouteDataCompletion(routeData []corelocation.Location, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("insertRouteData:completion:"), routeData, completion)
}/* debug [instance_methods/method]: InsertRouteDataCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKWorkoutRouteBuilder */

// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKWorkoutRouteBuilder) HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutRouteTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKWorkoutRouteBuilder */


