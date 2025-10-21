// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKMedicationConcept] class.
var (
	HKMedicationConceptClass     _HKMedicationConceptClass
	HKMedicationConceptClassOnce sync.Once
)

func getHKMedicationConceptClass() _HKMedicationConceptClass {
	HKMedicationConceptClassOnce.Do(func() {
		HKMedicationConceptClass = _HKMedicationConceptClass{objc.GetClass("HKMedicationConcept")}
	})
	return HKMedicationConceptClass
}

type _HKMedicationConceptClass struct {
	class objc.Class
}

// An interface definition for the [HKMedicationConcept] class.
type IHKMedicationConcept interface {
	objectivec.IObject
}

// An object that describes a specific medication concept.
//
// A medication concept represents the idea of a medication, like ibuprofen or insulin. It can have clinical significance, or can be created by the person using your app.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept
type HKMedicationConcept struct {
	objectivec.Object
}

// HKMedicationConceptFrom constructs a [HKMedicationConcept] from an unsafe.Pointer.
//
// An object that describes a specific medication concept.
func HKMedicationConceptFrom(ptr unsafe.Pointer) HKMedicationConcept {
	return HKMedicationConcept{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKMedicationConceptClass) Alloc() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKMedicationConceptClass) New() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKMedicationConcept) Init() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKMedicationConcept) Autorelease() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKMedicationConcept creates a new HKMedicationConcept instance.
func NewHKMedicationConcept() HKMedicationConcept {
	return getHKMedicationConceptClass().New()
}


// The display name for this medication.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/displayText
func (h_ HKMedicationConcept) DisplayText() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("displayText"))
	return rv
}

// The general form the medication is manufactured in.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/generalForm
func (h_ HKMedicationConcept) GeneralForm() HKMedicationGeneralForm {
	rv := objc.Send[HKMedicationGeneralForm](h_.ID, objc.Sel("generalForm"))
	return rv
}

// The unique identifier for the specific medication concept.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/identifier
func (h_ HKMedicationConcept) Identifier() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("identifier"))
	return rv
}

// The set of related clinical codings for the medication.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/relatedCodings
func (h_ HKMedicationConcept) RelatedCodings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("relatedCodings"))
	return rv
}



