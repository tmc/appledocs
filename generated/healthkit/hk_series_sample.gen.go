// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKSeriesSample] class.
type IHKSeriesSample interface {
	IHKSample
	Count() int
	SetCount(value int)
	HKWorkoutRouteTypeIdentifier() string
}

// An abstract base class that defines samples that contain a series of items.
//
// Never instantiate objects directly. Instead, user one of the concrete subclasses (for example, the class).
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKSeriesSampleClass) Alloc() HKSeriesSample {
	rv := objc.Send[HKSeriesSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The number of items in the series.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkseriessample/count
func (h_ HKSeriesSample) Count() int {
	rv := objc.Send[int](h_.ID, objc.Sel("count"))
	return rv
}


// SetCount sets the value of the count property.
// The number of items in the series.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkseriessample/count
func (h_ HKSeriesSample) SetCount(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCount:"), value)
}

// A series sample containing location data that defines the route the user took during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkworkoutroutetypeidentifier
func (h_ HKSeriesSample) HKWorkoutRouteTypeIdentifier() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKWorkoutRouteTypeIdentifier"))
	return rv
}



