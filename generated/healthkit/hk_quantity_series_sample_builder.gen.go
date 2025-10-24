// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Device() IHKDevice
	SetDevice(value IHKDevice)
	QuantityType() IHKQuantityType
	SetQuantityType(value IHKQuantityType)
	StartDate() objc.IObject /* cross-framework: Date */
	SetStartDate(value objc.IObject /* cross-framework: Date */)
	// methods:
}

// A builder object for incrementally building a sample that contains multiple quantities.


// A builder object for incrementally building a sample that contains multiple quantities.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/device
func (h_ HKQuantitySeriesSampleBuilder) Device() IHKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("device"))
	return rv
}


// The device providing the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/device
func (h_ HKQuantitySeriesSampleBuilder) SetDevice(value IHKDevice) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDevice:"), value)
}


// The quantity type for the series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/quantitytype
func (h_ HKQuantitySeriesSampleBuilder) QuantityType() IHKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("quantityType"))
	return rv
}


// The quantity type for the series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/quantitytype
func (h_ HKQuantitySeriesSampleBuilder) SetQuantityType(value IHKQuantityType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setQuantityType:"), value)
}


// The starting date and time for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/startdate
func (h_ HKQuantitySeriesSampleBuilder) StartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}


// The starting date and time for the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplebuilder/startdate
func (h_ HKQuantitySeriesSampleBuilder) SetStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}



