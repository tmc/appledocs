// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQuery */


/* debug [class_header]: Header for HKQuery */
// The class instance for the [HKQuery] class.
var (
	HKQueryClass     _HKQueryClass
	HKQueryClassOnce sync.Once
)

func getHKQueryClass() _HKQueryClass {
	HKQueryClassOnce.Do(func() {
		HKQueryClass = _HKQueryClass{objc.GetClass("HKQuery")}
	})
	return HKQueryClass
}

type _HKQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQuery */
// An interface definition for the [HKQuery] class.
type IHKQuery interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKQuery */
	// properties:
	ObjectType() IHKObjectType
	Predicate() foundation.Predicate
	SampleType() IHKSampleType
	HKPredicateKeyPathMetadata() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathUUID() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKQueryClass) Alloc() HKQuery {
	rv := objc.Send[HKQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKQueryClass) New() HKQuery {
	rv := objc.Send[HKQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuery) Init() HKQuery {
	rv := objc.Send[HKQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuery) Autorelease() HKQuery {
	rv := objc.Send[HKQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuery creates a new HKQuery instance.
func NewHKQuery() HKQuery {
	return getHKQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQuery */
// An abstract class for all the query classes in HealthKit.
//
// The class is the basis for all the query objects that retrieve data from the HealthKit store. The class is an abstract class. You should never instantiate it directly. Instead, you always work with one of its concrete subclasses.


// An abstract class for all the query classes in HealthKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery
type HKQuery struct {
	objectivec.Object
}

// HKQueryFrom constructs a [HKQuery] from an unsafe.Pointer.
//
// An abstract class for all the query classes in HealthKit.
func HKQueryFrom(ptr unsafe.Pointer) HKQuery {
	return HKQuery{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQuery *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQuery */

// Returns a predicate for matching all the activity summaries that fall between the days identified by the start and end date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicate(forActivitySummariesBetweenStart:end:)
func (hc _HKQueryClass) PredicateForActivitySummariesBetweenStartDateComponentsEndDateComponents(startDateComponents foundation.DateComponents, endDateComponents foundation.DateComponents) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForActivitySummariesBetweenStartDateComponents:endDateComponents:"), startDateComponents, endDateComponents)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForActivitySummariesBetweenStartDateComponentsEndDateComponents) */


// Returns a predicate that matches the activity summary for the specified day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForActivitySummary(with:)
func (hc _HKQueryClass) PredicateForActivitySummaryWithDateComponents(dateComponents foundation.DateComponents) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForActivitySummaryWithDateComponents:"), dateComponents)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForActivitySummaryWithDateComponents) */


// Returns a predicate that checks a category sample’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForCategorySamples(with:value:)
func (hc _HKQueryClass) PredicateForCategorySamplesWithOperatorTypeValue(operatorType PredicateOperatorType /* not a class type */, value int) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForCategorySamplesWithOperatorType:value:"), operatorType, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForCategorySamplesWithOperatorTypeValue) */


// A predicate that returns category samples with a matching value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForCategorySamplesEqualToValues:
func (hc _HKQueryClass) PredicateForCategorySamplesEqualToValues(values unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForCategorySamplesEqualToValues:"), values)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForCategorySamplesEqualToValues) */


// Returns a predicate for a specific FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForClinicalRecords(from:fhirResourceType:identifier:)
func (hc _HKQueryClass) PredicateForClinicalRecordsFromSourceFHIRResourceTypeIdentifier(source IHKSource, resourceType HKFHIRResourceType /* typedef */, identifier objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForClinicalRecordsFromSource:FHIRResourceType:identifier:"), source, resourceType, identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForClinicalRecordsFromSourceFHIRResourceTypeIdentifier) */


// Returns a predicate for a specific FHIR type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForClinicalRecords(withFHIRResourceType:)
func (hc _HKQueryClass) PredicateForClinicalRecordsWithFHIRResourceType(resourceType HKFHIRResourceType /* typedef */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForClinicalRecordsWithFHIRResourceType:"), resourceType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForClinicalRecordsWithFHIRResourceType) */


// Returns a predicate that matches electrocardiogram samples with the specified classification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForElectrocardiograms(classification:)
func (hc _HKQueryClass) PredicateForElectrocardiogramsWithClassification(classification HKElectrocardiogramClassification) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForElectrocardiogramsWithClassification:"), classification)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForElectrocardiogramsWithClassification) */


// Returns a predicate that matches electrocardiogram samples with the specified symptom status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForElectrocardiograms(symptomsStatus:)
func (hc _HKQueryClass) PredicateForElectrocardiogramsWithSymptomsStatus(symptomsStatus HKElectrocardiogramSymptomsStatus) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForElectrocardiogramsWithSymptomsStatus:"), symptomsStatus)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForElectrocardiogramsWithSymptomsStatus) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(medicationConceptIdentifier:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithMedicationConceptIdentifier(medicationConceptIdentifier IHKHealthConceptIdentifier) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithMedicationConceptIdentifier:"), medicationConceptIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithMedicationConceptIdentifier) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(medicationConceptIdentifiers:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithMedicationConceptIdentifiers(medicationConceptIdentifiers unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithMedicationConceptIdentifiers:"), medicationConceptIdentifiers)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithMedicationConceptIdentifiers) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(scheduledDate:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithScheduledDate(scheduledDate objc.IObject /* cross-framework: NSDate */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithScheduledDate:"), scheduledDate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithScheduledDate) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(scheduledDates:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithScheduledDates(scheduledDates unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithScheduledDates:"), scheduledDates)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithScheduledDates) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(scheduledStart:end:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithScheduledStartDateEndDate(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithScheduledStartDate:endDate:"), startDate, endDate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithScheduledStartDateEndDate) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(status:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithStatus(status HKMedicationDoseEventLogStatus) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithStatus:"), status)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithStatus) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForMedicationDoseEvent(statuses:)
func (hc _HKQueryClass) PredicateForMedicationDoseEventWithStatuses(statuses unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForMedicationDoseEventWithStatuses:"), statuses)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForMedicationDoseEventWithStatuses) */


// Returns a predicate that matches an object with the specified universally unique identifier (UUID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObject(with:)
func (hc _HKQueryClass) PredicateForObjectWithUUID(UUID foundation.UUID) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectWithUUID:"), UUID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectWithUUID) */


// Returns a predicate that matches all the objects that were created by any of the provided source revisions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(from:)-1ar4g
func (hc _HKQueryClass) PredicateForObjectsFromSourceRevisions(sourceRevisions unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsFromSourceRevisions:"), sourceRevisions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsFromSourceRevisions) */


// Returns a predicate that matches any objects that have been associated with the provided workout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(from:)-5irg9
func (hc _HKQueryClass) PredicateForObjectsFromWorkout(workout IHKWorkout) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsFromWorkout:"), workout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsFromWorkout) */


// Returns a predicate that matches all the objects that were created by the provided source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(from:)-7j3p2
func (hc _HKQueryClass) PredicateForObjectsFromSource(source IHKSource) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsFromSource:"), source)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsFromSource) */


// Returns a predicate that matches all the objects that were created by any of the provided sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(from:)-89b4t
func (hc _HKQueryClass) PredicateForObjectsFromSources(sources unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsFromSources:"), sources)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsFromSources) */


// Returns a predicate that matches all the objects that were created by any of the provided devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(from:)-9h87f
func (hc _HKQueryClass) PredicateForObjectsFromDevices(devices unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsFromDevices:"), devices)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsFromDevices) */


// Returns a predicate that matches the objects with the specified universally unique identifiers (UUIDs).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(with:)
func (hc _HKQueryClass) PredicateForObjectsWithUUIDs(UUIDs unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsWithUUIDs:"), UUIDs)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsWithUUIDs) */


// Returns a predicate that matches all objects created by devices with the specified properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(withDeviceProperty:allowedValues:)
func (hc _HKQueryClass) PredicateForObjectsWithDevicePropertyAllowedValues(key objc.IObject /* cross-framework: NSString */, allowedValues unsafe.Pointer) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsWithDeviceProperty:allowedValues:"), key, allowedValues)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsWithDevicePropertyAllowedValues) */


// Returns a predicate that matches any object whose metadata contains the provided key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(withMetadataKey:)
func (hc _HKQueryClass) PredicateForObjectsWithMetadataKey(key objc.IObject /* cross-framework: NSString */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsWithMetadataKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsWithMetadataKey) */


// Returns a predicate that matches objects based on the provided metadata key and an array of target values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(withMetadataKey:allowedValues:)
func (hc _HKQueryClass) PredicateForObjectsWithMetadataKeyAllowedValues(key objc.IObject /* cross-framework: NSString */, allowedValues objc.IObject /* cross-framework: NSArray */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsWithMetadataKey:allowedValues:"), key, allowedValues)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsWithMetadataKeyAllowedValues) */


// Returns a predicate that matches objects based on the provided metadata key, value, and operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjects(withMetadataKey:operatorType:value:)
func (hc _HKQueryClass) PredicateForObjectsWithMetadataKeyOperatorTypeValue(key objc.IObject /* cross-framework: NSString */, operatorType PredicateOperatorType /* not a class type */, value objc.IObject) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsWithMetadataKey:operatorType:value:"), key, operatorType, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsWithMetadataKeyOperatorTypeValue) */


// Returns a predicate that matches symptom samples associated with the specified electrocardiogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjectsAssociated(electrocardiogram:)
func (hc _HKQueryClass) PredicateForObjectsAssociatedWithElectrocardiogram(electrocardiogram IHKElectrocardiogram) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsAssociatedWithElectrocardiogram:"), electrocardiogram)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsAssociatedWithElectrocardiogram) */


// Returns a predicate that matches all objects that are not associated with a HealthKit correlation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForObjectsWithNoCorrelation()
func (hc _HKQueryClass) PredicateForObjectsWithNoCorrelation() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForObjectsWithNoCorrelation"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForObjectsWithNoCorrelation) */


// Returns a predicate that matches samples based on the target quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForQuantitySamples(with:quantity:)
func (hc _HKQueryClass) PredicateForQuantitySamplesWithOperatorTypeQuantity(operatorType PredicateOperatorType /* not a class type */, quantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForQuantitySamplesWithOperatorType:quantity:"), operatorType, quantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForQuantitySamplesWithOperatorTypeQuantity) */


// Returns a predicate for samples whose start and end dates fall within the specified time interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForSamples(withStart:end:options:)
func (hc _HKQueryClass) PredicateForSamplesWithStartDateEndDateOptions(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, options HKQueryOptions) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForSamplesWithStartDate:endDate:options:"), startDate, endDate, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForSamplesWithStartDateEndDateOptions) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForStatesOfMind(with:)-3iyym
func (hc _HKQueryClass) PredicateForStatesOfMindWithLabel(label HKStateOfMindLabel) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForStatesOfMindWithLabel:"), label)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForStatesOfMindWithLabel) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForStatesOfMind(with:)-6obe4
func (hc _HKQueryClass) PredicateForStatesOfMindWithKind(kind HKStateOfMindKind) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForStatesOfMindWithKind:"), kind)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForStatesOfMindWithKind) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForStatesOfMind(with:)-9fny6
func (hc _HKQueryClass) PredicateForStatesOfMindWithAssociation(association HKStateOfMindAssociation) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForStatesOfMindWithAssociation:"), association)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForStatesOfMindWithAssociation) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForStatesOfMind(withValence:operatorType:)
func (hc _HKQueryClass) PredicateForStatesOfMindWithValenceOperatorType(valence float64, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForStatesOfMindWithValence:operatorType:"), valence, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForStatesOfMindWithValenceOperatorType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForUserAnnotatedMedications(hasSchedule:)
func (hc _HKQueryClass) PredicateForUserAnnotatedMedicationsWithHasSchedule(hasSchedule bool) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForUserAnnotatedMedicationsWithHasSchedule:"), hasSchedule)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForUserAnnotatedMedicationsWithHasSchedule) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForUserAnnotatedMedications(isArchived:)
func (hc _HKQueryClass) PredicateForUserAnnotatedMedicationsWithIsArchived(isArchived bool) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForUserAnnotatedMedicationsWithIsArchived:"), isArchived)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForUserAnnotatedMedicationsWithIsArchived) */


// Returns a predicate that finds verifiable health records with a relevant date within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForVerifiableClinicalRecords(withRelevantDateWithin:)
func (hc _HKQueryClass) PredicateForVerifiableClinicalRecordsWithRelevantDateWithinDateInterval(dateInterval foundation.DateInterval) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForVerifiableClinicalRecordsWithRelevantDateWithinDateInterval:"), dateInterval)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForVerifiableClinicalRecordsWithRelevantDateWithinDateInterval) */


// Returns a predicate for matching workout activities based on their duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(operatorType:duration:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithOperatorTypeDuration(operatorType PredicateOperatorType /* not a class type */, duration float64) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithOperatorType:duration:"), operatorType, duration)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithOperatorTypeDuration) */


// Returns a predicate for matching workout activities based the average value of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(operatorType:quantityType:averageQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeAverageQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, averageQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithOperatorType:quantityType:averageQuantity:"), operatorType, quantityType, averageQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeAverageQuantity) */


// Returns a predicate for matching workout activities based the maximum value of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(operatorType:quantityType:maximumQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeMaximumQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, maximumQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithOperatorType:quantityType:maximumQuantity:"), operatorType, quantityType, maximumQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeMaximumQuantity) */


// Returns a predicate for matching workout activities based the minimum value of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(operatorType:quantityType:minimumQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeMinimumQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, minimumQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithOperatorType:quantityType:minimumQuantity:"), operatorType, quantityType, minimumQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeMinimumQuantity) */


// Returns a predicate for matching workout activities based the sum of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(operatorType:quantityType:sumQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeSumQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, sumQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithOperatorType:quantityType:sumQuantity:"), operatorType, quantityType, sumQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithOperatorTypeQuantityTypeSumQuantity) */


// Returns a predicate for workout activities that occur between the start and end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(start:end:options:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithStartDateEndDateOptions(startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, options HKQueryOptions) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithStartDate:endDate:options:"), startDate, endDate, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithStartDateEndDateOptions) */


// Returns a predicate for workout activities based on the type of activity performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutActivities(workoutActivityType:)
func (hc _HKQueryClass) PredicateForWorkoutActivitiesWithWorkoutActivityType(workoutActivityType HKWorkoutActivityType) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutActivitiesWithWorkoutActivityType:"), workoutActivityType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutActivitiesWithWorkoutActivityType) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkoutEffortSamplesRelated(workout:activity:)
func (hc _HKQueryClass) PredicateForWorkoutEffortSamplesRelatedToWorkoutActivity(workout IHKWorkout, activity IHKWorkoutActivity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutEffortSamplesRelatedToWorkout:activity:"), workout, activity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutEffortSamplesRelatedToWorkoutActivity) */


// Returns a predicate for matching workouts based on the associated workout activities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(activityPredicate:)
func (hc _HKQueryClass) PredicateForWorkoutsWithActivityPredicate(activityPredicate foundation.Predicate) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithActivityPredicate:"), activityPredicate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithActivityPredicate) */


// Returns a predicate for matching workouts based the average value of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(operatorType:quantityType:averageQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeQuantityTypeAverageQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, averageQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:quantityType:averageQuantity:"), operatorType, quantityType, averageQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeQuantityTypeAverageQuantity) */


// Returns a predicate for matching workout activities based the maximum value of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(operatorType:quantityType:maximumQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeQuantityTypeMaximumQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, maximumQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:quantityType:maximumQuantity:"), operatorType, quantityType, maximumQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeQuantityTypeMaximumQuantity) */


// Returns a predicate for matching workout activities based the minimum value of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(operatorType:quantityType:minimumQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeQuantityTypeMinimumQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, minimumQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:quantityType:minimumQuantity:"), operatorType, quantityType, minimumQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeQuantityTypeMinimumQuantity) */


// Returns a predicate for matching workout activities based the sum of an associated quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(operatorType:quantityType:sumQuantity:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeQuantityTypeSumQuantity(operatorType PredicateOperatorType /* not a class type */, quantityType IHKQuantityType, sumQuantity IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:quantityType:sumQuantity:"), operatorType, quantityType, sumQuantity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeQuantityTypeSumQuantity) */


// Returns a predicate for matching workouts based on the type of activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(with:)
func (hc _HKQueryClass) PredicateForWorkoutsWithWorkoutActivityType(workoutActivityType HKWorkoutActivityType) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithWorkoutActivityType:"), workoutActivityType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithWorkoutActivityType) */


// Returns a predicate for matching workouts based on their duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(with:duration:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeDuration(operatorType PredicateOperatorType /* not a class type */, duration float64) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:duration:"), operatorType, duration)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeDuration) */


// Returns a predicate for matching workouts based on the total distance traveled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(with:totalDistance:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeTotalDistance(operatorType PredicateOperatorType /* not a class type */, totalDistance IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:totalDistance:"), operatorType, totalDistance)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeTotalDistance) */


// Returns a predicate for matching workouts based on the total energy burned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(with:totalEnergyBurned:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeTotalEnergyBurned(operatorType PredicateOperatorType /* not a class type */, totalEnergyBurned IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:totalEnergyBurned:"), operatorType, totalEnergyBurned)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeTotalEnergyBurned) */


// Returns a predicate that matches workout samples based on the number of flights climbed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(with:totalFlightsClimbed:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeTotalFlightsClimbed(operatorType PredicateOperatorType /* not a class type */, totalFlightsClimbed IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:totalFlightsClimbed:"), operatorType, totalFlightsClimbed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeTotalFlightsClimbed) */


// Returns a predicate that matches workout samples based on the number of strokes while swimming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicateForWorkouts(with:totalSwimmingStrokeCount:)
func (hc _HKQueryClass) PredicateForWorkoutsWithOperatorTypeTotalSwimmingStrokeCount(operatorType PredicateOperatorType /* not a class type */, totalSwimmingStrokeCount IHKQuantity) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(hc.class), objc.Sel("predicateForWorkoutsWithOperatorType:totalSwimmingStrokeCount:"), operatorType, totalSwimmingStrokeCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForWorkoutsWithOperatorTypeTotalSwimmingStrokeCount) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQuery */

// The type of objects being queried.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/objectType
func (h_ HKQuery) ObjectType() IHKObjectType {
	rv := objc.Send[HKObjectType](h_.ID, objc.Sel("objectType"))
	return rv
}/* debug [instance_properties/getter]: objectType */


// A predicate used to filter the objects returned from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/predicate
func (h_ HKQuery) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](h_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// The type of objects being queried.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuery/sampleType
func (h_ HKQuery) SampleType() IHKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("sampleType"))
	return rv
}/* debug [instance_properties/getter]: sampleType */


// The key path for accessing the object’s metadata dictionary inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmetadata
func (h_ HKQuery) HKPredicateKeyPathMetadata() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMetadata"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMetadata */


// The key path for accessing the object’s UUID inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathuuid
func (h_ HKQuery) HKPredicateKeyPathUUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathUUID"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathUUID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQuery */



