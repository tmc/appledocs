// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKSeriesBuilder */


/* debug [class_header]: Header for HKSeriesBuilder */
// The class instance for the [HKSeriesBuilder] class.
var (
	HKSeriesBuilderClass     _HKSeriesBuilderClass
	HKSeriesBuilderClassOnce sync.Once
)

func getHKSeriesBuilderClass() _HKSeriesBuilderClass {
	HKSeriesBuilderClassOnce.Do(func() {
		HKSeriesBuilderClass = _HKSeriesBuilderClass{objc.GetClass("HKSeriesBuilder")}
	})
	return HKSeriesBuilderClass
}

type _HKSeriesBuilderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSeriesBuilder */
// An interface definition for the [HKSeriesBuilder] class.
type IHKSeriesBuilder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKSeriesBuilder */
	// properties:
	HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSeriesBuilder */
	// methods:
	Discard()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSeriesBuilder */
// Alloc allocates a new instance without initialization.
func (hc _HKSeriesBuilderClass) Alloc() HKSeriesBuilder {
	rv := objc.Send[HKSeriesBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSeriesBuilderClass) New() HKSeriesBuilder {
	rv := objc.Send[HKSeriesBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSeriesBuilder) Init() HKSeriesBuilder {
	rv := objc.Send[HKSeriesBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSeriesBuilder) Autorelease() HKSeriesBuilder {
	rv := objc.Send[HKSeriesBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSeriesBuilder creates a new HKSeriesBuilder instance.
func NewHKSeriesBuilder() HKSeriesBuilder {
	return getHKSeriesBuilderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSeriesBuilder */
// An abstract base class for building series samples.
//
// Never instantiate objects directly. Instead, user one of the concrete subclasses (for example, the class).


// An abstract base class for building series samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesBuilder
type HKSeriesBuilder struct {
	objectivec.Object
}

// HKSeriesBuilderFrom constructs a [HKSeriesBuilder] from an unsafe.Pointer.
//
// An abstract base class for building series samples.
func HKSeriesBuilderFrom(ptr unsafe.Pointer) HKSeriesBuilder {
	return HKSeriesBuilder{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSeriesBuilder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSeriesBuilder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSeriesBuilder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSeriesBuilder */

// Invalidates the builder and discards the collected data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesBuilder/discard()
func (h_ HKSeriesBuilder) Discard() {
	objc.Send[objc.ID](h_.ID, objc.Sel("discard"))
}/* debug [instance_methods/method]: Discard */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSeriesBuilder */

// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKSeriesBuilder) HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: HKWorkoutRouteTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSeriesBuilder */



