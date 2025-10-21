// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKQuantitySeriesSampleQuery] class.
var (
	HKQuantitySeriesSampleQueryClass     _HKQuantitySeriesSampleQueryClass
	HKQuantitySeriesSampleQueryClassOnce sync.Once
)

func getHKQuantitySeriesSampleQueryClass() _HKQuantitySeriesSampleQueryClass {
	HKQuantitySeriesSampleQueryClassOnce.Do(func() {
		HKQuantitySeriesSampleQueryClass = _HKQuantitySeriesSampleQueryClass{objc.GetClass("HKQuantitySeriesSampleQuery")}
	})
	return HKQuantitySeriesSampleQueryClass
}

type _HKQuantitySeriesSampleQueryClass struct {
	class objc.Class
}

// An interface definition for the [HKQuantitySeriesSampleQuery] class.
type IHKQuantitySeriesSampleQuery interface {
	IHKQuery
}

// A query that accesses the series data associated with a quantity sample.
//
// Use a series query to access the individual objects added to a sample using an .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantitySeriesSampleQuery
type HKQuantitySeriesSampleQuery struct {
	HKQuery
}

// HKQuantitySeriesSampleQueryFrom constructs a [HKQuantitySeriesSampleQuery] from an unsafe.Pointer.
//
// A query that accesses the series data associated with a quantity sample.
func HKQuantitySeriesSampleQueryFrom(ptr unsafe.Pointer) HKQuantitySeriesSampleQuery {
	return HKQuantitySeriesSampleQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQuantitySeriesSampleQueryClass) Alloc() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQuantitySeriesSampleQueryClass) New() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantitySeriesSampleQuery) Init() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantitySeriesSampleQuery) Autorelease() HKQuantitySeriesSampleQuery {
	rv := objc.Send[HKQuantitySeriesSampleQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantitySeriesSampleQuery creates a new HKQuantitySeriesSampleQuery instance.
func NewHKQuantitySeriesSampleQuery() HKQuantitySeriesSampleQuery {
	return getHKQuantitySeriesSampleQueryClass().New()
}


// A Boolean value that determines whether the query should return the series sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplequery/includesample
func (h_ HKQuantitySeriesSampleQuery) IncludeSample() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("includeSample"))
	return rv
}


// SetIncludeSample sets the value of the includeSample property.
// A Boolean value that determines whether the query should return the series sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplequery/includesample
func (h_ HKQuantitySeriesSampleQuery) SetIncludeSample(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIncludeSample:"), value)
}

// A Boolean value that determines whether the query groups the results based on the quantity sample’s start date.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplequery/orderbyquantitysamplestartdate
func (h_ HKQuantitySeriesSampleQuery) OrderByQuantitySampleStartDate() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("orderByQuantitySampleStartDate"))
	return rv
}


// SetOrderByQuantitySampleStartDate sets the value of the orderByQuantitySampleStartDate property.
// A Boolean value that determines whether the query groups the results based on the quantity sample’s start date.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquantityseriessamplequery/orderbyquantitysamplestartdate
func (h_ HKQuantitySeriesSampleQuery) SetOrderByQuantitySampleStartDate(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setOrderByQuantitySampleStartDate:"), value)
}



