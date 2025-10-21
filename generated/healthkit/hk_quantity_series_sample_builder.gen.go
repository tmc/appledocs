// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKQuantitySeriesSampleBuilder] class.
var (
	HKQuantitySeriesSampleBuilderClass     _HKQuantitySeriesSampleBuilderClass
	HKQuantitySeriesSampleBuilderClassOnce sync.Once
)

func getHKQuantitySeriesSampleBuilderClass() _HKQuantitySeriesSampleBuilderClass {
	HKQuantitySeriesSampleBuilderClassOnce.Do(func() {
		HKQuantitySeriesSampleBuilderClass = _HKQuantitySeriesSampleBuilderClass{objc.GetClass("HKQuantitySeriesSampleBuilder")}
	})
	return HKQuantitySeriesSampleBuilderClass
}

type _HKQuantitySeriesSampleBuilderClass struct {
	class objc.Class
}

// An interface definition for the [HKQuantitySeriesSampleBuilder] class.
type IHKQuantitySeriesSampleBuilder interface {
	objectivec.IObject
}

// A builder object for incrementally building a sample that contains multiple quantities.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleBuilder
type HKQuantitySeriesSampleBuilder struct {
	objectivec.Object
}

// HKQuantitySeriesSampleBuilderFrom constructs a [HKQuantitySeriesSampleBuilder] from an unsafe.Pointer.
//
// A builder object for incrementally building a sample that contains multiple quantities.
func HKQuantitySeriesSampleBuilderFrom(ptr unsafe.Pointer) HKQuantitySeriesSampleBuilder {
	return HKQuantitySeriesSampleBuilder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQuantitySeriesSampleBuilderClass) Alloc() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQuantitySeriesSampleBuilderClass) New() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantitySeriesSampleBuilder) Init() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantitySeriesSampleBuilder) Autorelease() HKQuantitySeriesSampleBuilder {
	rv := objc.Send[HKQuantitySeriesSampleBuilder](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantitySeriesSampleBuilder creates a new HKQuantitySeriesSampleBuilder instance.
func NewHKQuantitySeriesSampleBuilder() HKQuantitySeriesSampleBuilder {
	return getHKQuantitySeriesSampleBuilderClass().New()
}


// The device providing the data.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/device
func (h_ HKQuantitySeriesSampleBuilder) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The device providing the data.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/device
func (h_ HKQuantitySeriesSampleBuilder) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDevice:"), value)
}

// The quantity type for the series.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/quantitytype
func (h_ HKQuantitySeriesSampleBuilder) QuantityType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("quantityType"))
	return rv
}


// SetQuantityType sets the value of the quantityType property.
// The quantity type for the series.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/quantitytype
func (h_ HKQuantitySeriesSampleBuilder) SetQuantityType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setQuantityType:"), value)
}

// The starting date and time for the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/startdate
func (h_ HKQuantitySeriesSampleBuilder) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The starting date and time for the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/startdate
func (h_ HKQuantitySeriesSampleBuilder) SetStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}



