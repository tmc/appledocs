// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKUserAnnotatedMedicationType] class.
var (
	HKUserAnnotatedMedicationTypeClass     _HKUserAnnotatedMedicationTypeClass
	HKUserAnnotatedMedicationTypeClassOnce sync.Once
)

func getHKUserAnnotatedMedicationTypeClass() _HKUserAnnotatedMedicationTypeClass {
	HKUserAnnotatedMedicationTypeClassOnce.Do(func() {
		HKUserAnnotatedMedicationTypeClass = _HKUserAnnotatedMedicationTypeClass{objc.GetClass("HKUserAnnotatedMedicationType")}
	})
	return HKUserAnnotatedMedicationTypeClass
}

type _HKUserAnnotatedMedicationTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKUserAnnotatedMedicationType] class.
type IHKUserAnnotatedMedicationType interface {
	IHKObjectType
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedicationType

type HKUserAnnotatedMedicationType struct {
	HKObjectType
}

// HKUserAnnotatedMedicationTypeFrom constructs a [HKUserAnnotatedMedicationType] from an unsafe.Pointer.
func HKUserAnnotatedMedicationTypeFrom(ptr unsafe.Pointer) HKUserAnnotatedMedicationType {
	return HKUserAnnotatedMedicationType{
		HKObjectType: HKObjectTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKUserAnnotatedMedicationTypeClass) Alloc() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKUserAnnotatedMedicationTypeClass) New() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUserAnnotatedMedicationType) Init() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUserAnnotatedMedicationType) Autorelease() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUserAnnotatedMedicationType creates a new HKUserAnnotatedMedicationType instance.
func NewHKUserAnnotatedMedicationType() HKUserAnnotatedMedicationType {
	return getHKUserAnnotatedMedicationTypeClass().New()
}




