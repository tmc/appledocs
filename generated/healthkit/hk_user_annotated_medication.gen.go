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
	// properties:
	HasSchedule() bool /* primitive/slice/pointer. */
	SetHasSchedule(value bool /* primitive/slice/pointer. */)
	IsArchived() bool /* primitive/slice/pointer. */
	SetIsArchived(value bool /* primitive/slice/pointer. */)
	Medication() IHKMedicationConcept
	SetMedication(value IHKMedicationConcept)
	Nickname() string /* primitive/slice/pointer. */
	SetNickname(value string /* primitive/slice/pointer. */)
	// methods:
}

// A reference to the tracked medication and the details a person can customize.
//
// The details are relevant to the medication tracking experience.


// A reference to the tracked medication and the details a person can customize.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/hasschedule
func (h_ HKUserAnnotatedMedication) HasSchedule() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("hasSchedule"))
	return rv
}


// A Boolean value that indicates whether a medication has a schedule set up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/hasschedule
func (h_ HKUserAnnotatedMedication) SetHasSchedule(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHasSchedule:"), value)
}


// A Boolean value that indicates whether a medication is archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/isarchived
func (h_ HKUserAnnotatedMedication) IsArchived() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("isArchived"))
	return rv
}


// A Boolean value that indicates whether a medication is archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/isarchived
func (h_ HKUserAnnotatedMedication) SetIsArchived(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsArchived:"), value)
}


// A reference to the specific medication a person is tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/medication
func (h_ HKUserAnnotatedMedication) Medication() IHKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](h_.ID, objc.Sel("medication"))
	return rv
}


// A reference to the specific medication a person is tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/medication
func (h_ HKUserAnnotatedMedication) SetMedication(value IHKMedicationConcept) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMedication:"), value)
}


// The nickname that a person added to a medication during the entry experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/nickname
func (h_ HKUserAnnotatedMedication) Nickname() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("nickname"))
	return rv
}


// The nickname that a person added to a medication during the entry experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkuserannotatedmedication/nickname
func (h_ HKUserAnnotatedMedication) SetNickname(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setNickname:"), objc.String(value))
}



