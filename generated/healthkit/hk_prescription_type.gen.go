// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKPrescriptionType] class.
var (
	HKPrescriptionTypeClass     _HKPrescriptionTypeClass
	HKPrescriptionTypeClassOnce sync.Once
)

func getHKPrescriptionTypeClass() _HKPrescriptionTypeClass {
	HKPrescriptionTypeClassOnce.Do(func() {
		HKPrescriptionTypeClass = _HKPrescriptionTypeClass{objc.GetClass("HKPrescriptionType")}
	})
	return HKPrescriptionTypeClass
}

type _HKPrescriptionTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKPrescriptionType] class.
type IHKPrescriptionType interface {
	IHKSampleType
}

// A type that identifies samples that store a prescription.
//
// The class is a concrete subclass of the class. To create a vision prescription type instances, use the convenience method. Use this data type to request permission to save vision prescriptions to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrescriptionType
type HKPrescriptionType struct {
	HKSampleType
}

// HKPrescriptionTypeFrom constructs a [HKPrescriptionType] from an unsafe.Pointer.
//
// A type that identifies samples that store a prescription.
func HKPrescriptionTypeFrom(ptr unsafe.Pointer) HKPrescriptionType {
	return HKPrescriptionType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKPrescriptionTypeClass) Alloc() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKPrescriptionTypeClass) New() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKPrescriptionType) Init() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKPrescriptionType) Autorelease() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKPrescriptionType creates a new HKPrescriptionType instance.
func NewHKPrescriptionType() HKPrescriptionType {
	return getHKPrescriptionTypeClass().New()
}




