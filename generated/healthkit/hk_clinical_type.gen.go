// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKClinicalType] class.
var (
	HKClinicalTypeClass     _HKClinicalTypeClass
	HKClinicalTypeClassOnce sync.Once
)

func getHKClinicalTypeClass() _HKClinicalTypeClass {
	HKClinicalTypeClassOnce.Do(func() {
		HKClinicalTypeClass = _HKClinicalTypeClass{objc.GetClass("HKClinicalType")}
	})
	return HKClinicalTypeClass
}

type _HKClinicalTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKClinicalType] class.
type IHKClinicalType interface {
	IHKSampleType
}

// A type that identifies samples that contain clinical record data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalType
type HKClinicalType struct {
	HKSampleType
}

// HKClinicalTypeFrom constructs a [HKClinicalType] from an unsafe.Pointer.
//
// A type that identifies samples that contain clinical record data.
func HKClinicalTypeFrom(ptr unsafe.Pointer) HKClinicalType {
	return HKClinicalType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKClinicalTypeClass) Alloc() HKClinicalType {
	rv := objc.Send[HKClinicalType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKClinicalTypeClass) New() HKClinicalType {
	rv := objc.Send[HKClinicalType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKClinicalType) Init() HKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKClinicalType) Autorelease() HKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKClinicalType creates a new HKClinicalType instance.
func NewHKClinicalType() HKClinicalType {
	return getHKClinicalTypeClass().New()
}




