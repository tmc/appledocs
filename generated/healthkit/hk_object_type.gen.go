// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKObjectType */


/* debug [class_header]: Header for HKObjectType */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKObjectType */
// An interface definition for the [HKObjectType] class.
type IHKObjectType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKObjectType */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKObjectType */
	// methods:
	RequiresPerObjectAuthorization() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKObjectType */
// Alloc allocates a new instance without initialization.
func (hc _HKObjectTypeClass) Alloc() HKObjectType {
	rv := objc.Send[HKObjectType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKObjectType */
// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
//
// The class is an abstract class. Don’t instantiate an object directly. Instead, instantiate one of the following concrete subclasses: The class provides a convenience method to create each of these subclasses.


// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKObjectType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKObjectType */

// Returns the shared activity summary type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/activitySummaryType()
func (hc _HKObjectTypeClass) ActivitySummaryType() IHKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](objc.ID(hc.class), objc.Sel("activitySummaryType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ActivitySummaryType) */


// Returns an audiogram sample type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/audiogramSampleType()
func (hc _HKObjectTypeClass) AudiogramSampleType() IHKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](objc.ID(hc.class), objc.Sel("audiogramSampleType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudiogramSampleType) */


// Returns the shared category type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/categoryType(forIdentifier:)
func (hc _HKObjectTypeClass) CategoryTypeForIdentifier(identifier HKCategoryTypeIdentifier /* typedef */) IHKCategoryType {
	rv := objc.Send[HKCategoryType](objc.ID(hc.class), objc.Sel("categoryTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CategoryTypeForIdentifier) */


// Returns the shared characteristic type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/characteristicType(forIdentifier:)
func (hc _HKObjectTypeClass) CharacteristicTypeForIdentifier(identifier HKCharacteristicTypeIdentifier /* typedef */) IHKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](objc.ID(hc.class), objc.Sel("characteristicTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacteristicTypeForIdentifier) */


// Returns the shared clinical type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/clinicalType(forIdentifier:)
func (hc _HKObjectTypeClass) ClinicalTypeForIdentifier(identifier HKClinicalTypeIdentifier /* typedef */) IHKClinicalType {
	rv := objc.Send[HKClinicalType](objc.ID(hc.class), objc.Sel("clinicalTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ClinicalTypeForIdentifier) */


// Returns the shared correlation type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/correlationType(forIdentifier:)
func (hc _HKObjectTypeClass) CorrelationTypeForIdentifier(identifier HKCorrelationTypeIdentifier /* typedef */) IHKCorrelationType {
	rv := objc.Send[HKCorrelationType](objc.ID(hc.class), objc.Sel("correlationTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CorrelationTypeForIdentifier) */


// Returns the shared document type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/documentType(forIdentifier:)
func (hc _HKObjectTypeClass) DocumentTypeForIdentifier(identifier HKDocumentTypeIdentifier /* typedef */) IHKDocumentType {
	rv := objc.Send[HKDocumentType](objc.ID(hc.class), objc.Sel("documentTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DocumentTypeForIdentifier) */


// Returns the shared electrocardiogram type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/electrocardiogramType()
func (hc _HKObjectTypeClass) ElectrocardiogramType() IHKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](objc.ID(hc.class), objc.Sel("electrocardiogramType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ElectrocardiogramType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/medicationDoseEventType()
func (hc _HKObjectTypeClass) MedicationDoseEventType() IHKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](objc.ID(hc.class), objc.Sel("medicationDoseEventType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MedicationDoseEventType) */


// Returns the shared quantity type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/quantityType(forIdentifier:)
func (hc _HKObjectTypeClass) QuantityTypeForIdentifier(identifier HKQuantityTypeIdentifier /* typedef */) IHKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("quantityTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuantityTypeForIdentifier) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/scoredAssessmentTypeForIdentifier:
func (hc _HKObjectTypeClass) ScoredAssessmentTypeForIdentifier(identifier HKScoredAssessmentTypeIdentifier /* typedef */) IHKScoredAssessmentType {
	rv := objc.Send[HKScoredAssessmentType](objc.ID(hc.class), objc.Sel("scoredAssessmentTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScoredAssessmentTypeForIdentifier) */


// Returns the shared series type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/seriesType(forIdentifier:)
func (hc _HKObjectTypeClass) SeriesTypeForIdentifier(identifier objc.IObject /* cross-framework: NSString */) IHKSeriesType {
	rv := objc.Send[HKSeriesType](objc.ID(hc.class), objc.Sel("seriesTypeForIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SeriesTypeForIdentifier) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/stateOfMindType()
func (hc _HKObjectTypeClass) StateOfMindType() IHKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("stateOfMindType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StateOfMindType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/userAnnotatedMedicationType()
func (hc _HKObjectTypeClass) UserAnnotatedMedicationType() IHKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](objc.ID(hc.class), objc.Sel("userAnnotatedMedicationType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UserAnnotatedMedicationType) */


// Returns a shared vision prescription type object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/visionPrescriptionType()
func (hc _HKObjectTypeClass) VisionPrescriptionType() IHKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](objc.ID(hc.class), objc.Sel("visionPrescriptionType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VisionPrescriptionType) */


// Returns the shared object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/workoutType()
func (hc _HKObjectTypeClass) WorkoutType() IHKWorkoutType {
	rv := objc.Send[HKWorkoutType](objc.ID(hc.class), objc.Sel("workoutType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorkoutType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKObjectType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKObjectType */

// Returns a Boolean that indicates whether the data type requires per-object authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/requiresPerObjectAuthorization()
func (h_ HKObjectType) RequiresPerObjectAuthorization() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("requiresPerObjectAuthorization"))
	return rv
}/* debug [instance_methods/method]: RequiresPerObjectAuthorization */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKObjectType */

// A unique string identifying the HealthKit object type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/identifier
func (h_ HKObjectType) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKObjectType */



