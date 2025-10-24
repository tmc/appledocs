// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKWorkoutRouteBuilder] class.
type IHKWorkoutRouteBuilder interface {
	IHKSeriesBuilder
	// properties:
	HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKWorkoutRouteBuilderClass) Alloc() HKWorkoutRouteBuilder {
	rv := objc.Send[HKWorkoutRouteBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKWorkoutRouteBuilder) HKWorkoutRouteTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}



