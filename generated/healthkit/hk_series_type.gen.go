// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKSeriesType */


/* debug [class_header]: Header for HKSeriesType */
// The class instance for the [HKSeriesType] class.
var (
	HKSeriesTypeClass     _HKSeriesTypeClass
	HKSeriesTypeClassOnce sync.Once
)

func getHKSeriesTypeClass() _HKSeriesTypeClass {
	HKSeriesTypeClassOnce.Do(func() {
		HKSeriesTypeClass = _HKSeriesTypeClass{objc.GetClass("HKSeriesType")}
	})
	return HKSeriesTypeClass
}

type _HKSeriesTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSeriesType */
// An interface definition for the [HKSeriesType] class.
type IHKSeriesType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKSeriesType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSeriesType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSeriesType */
// Alloc allocates a new instance without initialization.
func (hc _HKSeriesTypeClass) Alloc() HKSeriesType {
	rv := objc.Send[HKSeriesType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSeriesTypeClass) New() HKSeriesType {
	rv := objc.Send[HKSeriesType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSeriesType) Init() HKSeriesType {
	rv := objc.Send[HKSeriesType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSeriesType) Autorelease() HKSeriesType {
	rv := objc.Send[HKSeriesType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSeriesType creates a new HKSeriesType instance.
func NewHKSeriesType() HKSeriesType {
	return getHKSeriesTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSeriesType */
// A type that indicates the data stored in a series sample.


// A type that indicates the data stored in a series sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesType
type HKSeriesType struct {
	HKSampleType
}

// HKSeriesTypeFrom constructs a [HKSeriesType] from an unsafe.Pointer.
//
// A type that indicates the data stored in a series sample.
func HKSeriesTypeFrom(ptr unsafe.Pointer) HKSeriesType {
	return HKSeriesType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSeriesType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSeriesType */

// Returns a series type object for heartbeat data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesType/heartbeat()
func (hc _HKSeriesTypeClass) HeartbeatSeriesType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("heartbeatSeriesType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HeartbeatSeriesType) */


// Returns a series type object for workout routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesType/workoutRoute()
func (hc _HKSeriesTypeClass) WorkoutRouteType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("workoutRouteType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutRouteType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSeriesType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSeriesType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSeriesType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSeriesType */



