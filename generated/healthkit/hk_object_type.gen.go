// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKObjectType] class.
var (
	HKObjectTypeClass     _HKObjectTypeClass
	HKObjectTypeClassOnce sync.Once
)

func getHKObjectTypeClass() _HKObjectTypeClass {
	HKObjectTypeClassOnce.Do(func() {
		HKObjectTypeClass = _HKObjectTypeClass{objc.GetClass("HKObjectType")}
	})
	return HKObjectTypeClass
}

type _HKObjectTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKObjectType] class.
type IHKObjectType interface {
	objectivec.IObject
	RequiresPerObjectAuthorization() bool
}

// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
//
// The class is an abstract class. Don’t instantiate an object directly. Instead, instantiate one of the following concrete subclasses: The class provides a convenience method to create each of these subclasses.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType
type HKObjectType struct {
	objectivec.Object
}

// HKObjectTypeFrom constructs a [HKObjectType] from an unsafe.Pointer.
//
// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
func HKObjectTypeFrom(ptr unsafe.Pointer) HKObjectType {
	return HKObjectType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKObjectTypeClass) Alloc() HKObjectType {
	rv := objc.Send[HKObjectType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKObjectTypeClass) New() HKObjectType {
	rv := objc.Send[HKObjectType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKObjectType) Init() HKObjectType {
	rv := objc.Send[HKObjectType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKObjectType) Autorelease() HKObjectType {
	rv := objc.Send[HKObjectType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKObjectType creates a new HKObjectType instance.
func NewHKObjectType() HKObjectType {
	return getHKObjectTypeClass().New()
}


// Returns the shared activity summary type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/activitySummaryType()
func (hc _HKObjectTypeClass) ActivitySummaryType() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](objc.ID(hc.class), objc.Sel("activitySummaryType"))
	return rv
}

// Returns an audiogram sample type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/audiogramSampleType()
func (hc _HKObjectTypeClass) AudiogramSampleType() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](objc.ID(hc.class), objc.Sel("audiogramSampleType"))
	return rv
}

// Returns the shared category type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/categoryType(forIdentifier:)
func (hc _HKObjectTypeClass) CategoryTypeForIdentifier(identifier IHKCategoryTypeIdentifier) HKCategoryType {
	rv := objc.Send[HKCategoryType](objc.ID(hc.class), objc.Sel("categoryTypeForIdentifier:"), identifier)
	return rv
}

// Returns the shared characteristic type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/characteristicType(forIdentifier:)
func (hc _HKObjectTypeClass) CharacteristicTypeForIdentifier(identifier IHKCharacteristicTypeIdentifier) HKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](objc.ID(hc.class), objc.Sel("characteristicTypeForIdentifier:"), identifier)
	return rv
}

// Returns the shared clinical type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/clinicalType(forIdentifier:)
func (hc _HKObjectTypeClass) ClinicalTypeForIdentifier(identifier IHKClinicalTypeIdentifier) HKClinicalType {
	rv := objc.Send[HKClinicalType](objc.ID(hc.class), objc.Sel("clinicalTypeForIdentifier:"), identifier)
	return rv
}

// Returns the shared correlation type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/correlationType(forIdentifier:)
func (hc _HKObjectTypeClass) CorrelationTypeForIdentifier(identifier IHKCorrelationTypeIdentifier) HKCorrelationType {
	rv := objc.Send[HKCorrelationType](objc.ID(hc.class), objc.Sel("correlationTypeForIdentifier:"), identifier)
	return rv
}

// Returns the shared document type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/documentType(forIdentifier:)
func (hc _HKObjectTypeClass) DocumentTypeForIdentifier(identifier IHKDocumentTypeIdentifier) HKDocumentType {
	rv := objc.Send[HKDocumentType](objc.ID(hc.class), objc.Sel("documentTypeForIdentifier:"), identifier)
	return rv
}

// Returns the shared electrocardiogram type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/electrocardiogramType()
func (hc _HKObjectTypeClass) ElectrocardiogramType() HKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](objc.ID(hc.class), objc.Sel("electrocardiogramType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/medicationDoseEventType()
func (hc _HKObjectTypeClass) MedicationDoseEventType() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](objc.ID(hc.class), objc.Sel("medicationDoseEventType"))
	return rv
}

// Returns the shared quantity type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/quantityType(forIdentifier:)
func (hc _HKObjectTypeClass) QuantityTypeForIdentifier(identifier IHKQuantityTypeIdentifier) HKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("quantityTypeForIdentifier:"), identifier)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/scoredAssessmentTypeForIdentifier:
func (hc _HKObjectTypeClass) ScoredAssessmentTypeForIdentifier(identifier IHKScoredAssessmentTypeIdentifier) HKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](objc.ID(hc.class), objc.Sel("scoredAssessmentTypeForIdentifier:"), identifier)
	return rv
}

// Returns the shared series type for the provided identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/seriesType(forIdentifier:)
func (hc _HKObjectTypeClass) SeriesTypeForIdentifier(identifier appkit.string) HKSeriesType {
	rv := objc.Send[HKSeriesType](objc.ID(hc.class), objc.Sel("seriesTypeForIdentifier:"), identifier)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/stateOfMindType()
func (hc _HKObjectTypeClass) StateOfMindType() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("stateOfMindType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/userAnnotatedMedicationType()
func (hc _HKObjectTypeClass) UserAnnotatedMedicationType() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](objc.ID(hc.class), objc.Sel("userAnnotatedMedicationType"))
	return rv
}

// Returns a shared vision prescription type object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/visionPrescriptionType()
func (hc _HKObjectTypeClass) VisionPrescriptionType() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](objc.ID(hc.class), objc.Sel("visionPrescriptionType"))
	return rv
}

// Returns the shared object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/workoutType()
func (hc _HKObjectTypeClass) WorkoutType() HKWorkoutType {
	rv := objc.Send[HKWorkoutType](objc.ID(hc.class), objc.Sel("workoutType"))
	return rv
}

// Returns a Boolean that indicates whether the data type requires per-object authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/requiresPerObjectAuthorization()
func (h_ HKObjectType) RequiresPerObjectAuthorization() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("requiresPerObjectAuthorization"))
	return rv
}

// A unique string identifying the HealthKit object type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/identifier
func (h_ HKObjectType) Identifier() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("identifier"))
	return rv
}



