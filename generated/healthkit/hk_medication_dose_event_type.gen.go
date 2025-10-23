// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKMedicationDoseEventType] class.
var (
	HKMedicationDoseEventTypeClass     _HKMedicationDoseEventTypeClass
	HKMedicationDoseEventTypeClassOnce sync.Once
)

func getHKMedicationDoseEventTypeClass() _HKMedicationDoseEventTypeClass {
	HKMedicationDoseEventTypeClassOnce.Do(func() {
		HKMedicationDoseEventTypeClass = _HKMedicationDoseEventTypeClass{objc.GetClass("HKMedicationDoseEventType")}
	})
	return HKMedicationDoseEventTypeClass
}

type _HKMedicationDoseEventTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKMedicationDoseEventType] class.
type IHKMedicationDoseEventType interface {
	IHKSampleType
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEventType
type HKMedicationDoseEventType struct {
	HKSampleType
}

// HKMedicationDoseEventTypeFrom constructs a [HKMedicationDoseEventType] from an unsafe.Pointer.
func HKMedicationDoseEventTypeFrom(ptr unsafe.Pointer) HKMedicationDoseEventType {
	return HKMedicationDoseEventType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKMedicationDoseEventTypeClass) Alloc() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKMedicationDoseEventTypeClass) New() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKMedicationDoseEventType) Init() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKMedicationDoseEventType) Autorelease() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKMedicationDoseEventType creates a new HKMedicationDoseEventType instance.
func NewHKMedicationDoseEventType() HKMedicationDoseEventType {
	return getHKMedicationDoseEventTypeClass().New()
}




