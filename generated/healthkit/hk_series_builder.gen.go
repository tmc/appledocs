// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKSeriesBuilder] class.
type IHKSeriesBuilder interface {
	objectivec.IObject
	Discard()
}

// An abstract base class for building series samples.
//
// Never instantiate objects directly. Instead, user one of the concrete subclasses (for example, the class).
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKSeriesBuilderClass) Alloc() HKSeriesBuilder {
	rv := objc.Send[HKSeriesBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Invalidates the builder and discards the collected data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSeriesBuilder/discard()
func (h_ HKSeriesBuilder) Discard() {
	objc.Send[objc.ID](h_.ID, objc.Sel("discard"))
}

// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKSeriesBuilder) HKWorkoutRouteTypeIdentifier() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}



