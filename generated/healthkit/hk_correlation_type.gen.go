// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCorrelationType] class.
var (
	HKCorrelationTypeClass     _HKCorrelationTypeClass
	HKCorrelationTypeClassOnce sync.Once
)

func getHKCorrelationTypeClass() _HKCorrelationTypeClass {
	HKCorrelationTypeClassOnce.Do(func() {
		HKCorrelationTypeClass = _HKCorrelationTypeClass{objc.GetClass("HKCorrelationType")}
	})
	return HKCorrelationTypeClass
}

type _HKCorrelationTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKCorrelationType] class.
type IHKCorrelationType interface {
	IHKSampleType
}

// A type that identifies samples that group multiple subsamples.
//
// The class is a concrete subclass of the class. To create a correlation type instance, use the object type’s conveniance method. Use correlation types to: Request permission to read or write matching quantity samples. Create and share matching quantity samples. Query for matching quantity samples. HealthKit provides two correlation types: blood pressure and food.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationType
type HKCorrelationType struct {
	HKSampleType
}

// HKCorrelationTypeFrom constructs a [HKCorrelationType] from an unsafe.Pointer.
//
// A type that identifies samples that group multiple subsamples.
func HKCorrelationTypeFrom(ptr unsafe.Pointer) HKCorrelationType {
	return HKCorrelationType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCorrelationTypeClass) Alloc() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCorrelationTypeClass) New() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCorrelationType) Init() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCorrelationType) Autorelease() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCorrelationType creates a new HKCorrelationType instance.
func NewHKCorrelationType() HKCorrelationType {
	return getHKCorrelationTypeClass().New()
}




