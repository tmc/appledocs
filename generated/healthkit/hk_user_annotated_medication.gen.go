// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKUserAnnotatedMedication] class.
var (
	HKUserAnnotatedMedicationClass     _HKUserAnnotatedMedicationClass
	HKUserAnnotatedMedicationClassOnce sync.Once
)

func getHKUserAnnotatedMedicationClass() _HKUserAnnotatedMedicationClass {
	HKUserAnnotatedMedicationClassOnce.Do(func() {
		HKUserAnnotatedMedicationClass = _HKUserAnnotatedMedicationClass{objc.GetClass("HKUserAnnotatedMedication")}
	})
	return HKUserAnnotatedMedicationClass
}

type _HKUserAnnotatedMedicationClass struct {
	class objc.Class
}

// An interface definition for the [HKUserAnnotatedMedication] class.
type IHKUserAnnotatedMedication interface {
	objectivec.IObject
}

// A reference to the tracked medication and the details a person can customize.
//
// The details are relevant to the medication tracking experience.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication
type HKUserAnnotatedMedication struct {
	objectivec.Object
}

// HKUserAnnotatedMedicationFrom constructs a [HKUserAnnotatedMedication] from an unsafe.Pointer.
//
// A reference to the tracked medication and the details a person can customize.
func HKUserAnnotatedMedicationFrom(ptr unsafe.Pointer) HKUserAnnotatedMedication {
	return HKUserAnnotatedMedication{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKUserAnnotatedMedicationClass) Alloc() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKUserAnnotatedMedicationClass) New() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUserAnnotatedMedication) Init() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUserAnnotatedMedication) Autorelease() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUserAnnotatedMedication creates a new HKUserAnnotatedMedication instance.
func NewHKUserAnnotatedMedication() HKUserAnnotatedMedication {
	return getHKUserAnnotatedMedicationClass().New()
}


// A Boolean value that indicates whether a medication has a schedule set up.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/hasSchedule
func (h_ HKUserAnnotatedMedication) HasSchedule() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("hasSchedule"))
	return rv
}

// A Boolean value that indicates whether a medication is archived.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/isArchived
func (h_ HKUserAnnotatedMedication) IsArchived() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isArchived"))
	return rv
}

// A reference to the specific medication a person is tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/medication
func (h_ HKUserAnnotatedMedication) Medication() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("medication"))
	return rv
}

// The nickname that a person added to a medication during the entry experience.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/nickname
func (h_ HKUserAnnotatedMedication) Nickname() string {
	rv := objc.Send[string](h_.ID, objc.Sel("nickname"))
	return rv
}



