// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKSeriesType] class.
type IHKSeriesType interface {
	IHKSampleType
	// properties:
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKSeriesTypeClass) Alloc() HKSeriesType {
	rv := objc.Send[HKSeriesType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




