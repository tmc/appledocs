// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKSeriesSample */


/* debug [class_header]: Header for HKSeriesSample */
// The class instance for the [HKSeriesSample] class.
var (
	HKSeriesSampleClass     _HKSeriesSampleClass
	HKSeriesSampleClassOnce sync.Once
)

func getHKSeriesSampleClass() _HKSeriesSampleClass {
	HKSeriesSampleClassOnce.Do(func() {
		HKSeriesSampleClass = _HKSeriesSampleClass{objc.GetClass("HKSeriesSample")}
	})
	return HKSeriesSampleClass
}

type _HKSeriesSampleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSeriesSample */
// An interface definition for the [HKSeriesSample] class.
type IHKSeriesSample interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKSeriesSample */
	// properties:
	Count() uint
	HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSeriesSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSeriesSample */
// Alloc allocates a new instance without initialization.
func (hc _HKSeriesSampleClass) Alloc() HKSeriesSample {
	rv := objc.Send[HKSeriesSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSeriesSampleClass) New() HKSeriesSample {
	rv := objc.Send[HKSeriesSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSeriesSample) Init() HKSeriesSample {
	rv := objc.Send[HKSeriesSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSeriesSample) Autorelease() HKSeriesSample {
	rv := objc.Send[HKSeriesSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSeriesSample creates a new HKSeriesSample instance.
func NewHKSeriesSample() HKSeriesSample {
	return getHKSeriesSampleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSeriesSample */
// An abstract base class that defines samples that contain a series of items.
//
// Never instantiate objects directly. Instead, user one of the concrete subclasses (for example, the class).


// An abstract base class that defines samples that contain a series of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesSample
type HKSeriesSample struct {
	HKSample
}

// HKSeriesSampleFrom constructs a [HKSeriesSample] from an unsafe.Pointer.
//
// An abstract base class that defines samples that contain a series of items.
func HKSeriesSampleFrom(ptr unsafe.Pointer) HKSeriesSample {
	return HKSeriesSample{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSeriesSample *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSeriesSample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSeriesSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSeriesSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSeriesSample */

// The number of items in the series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesSample/count
func (h_ HKSeriesSample) Count() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKSeriesSample) HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutRouteTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSeriesSample */



