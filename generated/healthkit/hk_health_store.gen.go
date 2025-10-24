// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKHealthStore] class.
var (
	HKHealthStoreClass     _HKHealthStoreClass
	HKHealthStoreClassOnce sync.Once
)

func getHKHealthStoreClass() _HKHealthStoreClass {
	HKHealthStoreClassOnce.Do(func() {
		HKHealthStoreClass = _HKHealthStoreClass{objc.GetClass("HKHealthStore")}
	})
	return HKHealthStoreClass
}

type _HKHealthStoreClass struct {
	class objc.Class
}

// An interface definition for the [HKHealthStore] class.
type IHKHealthStore interface {
	objectivec.IObject
	// properties:
	// methods:
	ActivityMoveModeWithError(error_ unsafe.Pointer) IHKActivityMoveModeObject
	AuthorizationStatusForType(type_ IHKObjectType) HKAuthorizationStatus
	BiologicalSexWithError(error_ unsafe.Pointer) IHKBiologicalSexObject
	BloodTypeWithError(error_ unsafe.Pointer) IHKBloodTypeObject
	DateOfBirthComponentsWithError(error_ unsafe.Pointer) objc.IObject /* cross-framework: DateComponents */
	DeleteObjectsWithCompletion(objects []IHKObject, completion unsafe.Pointer)
	DeleteObjectWithCompletion(object IHKObject, completion unsafe.Pointer)
	DeleteObjectsOfTypePredicateWithCompletion(objectType IHKObjectType, predicate objc.IObject /* cross-framework: Predicate */, completion unsafe.Pointer)
	DisableAllBackgroundDeliveryWithCompletion(completion unsafe.Pointer)
	DisableBackgroundDeliveryForTypeWithCompletion(type_ IHKObjectType, completion unsafe.Pointer)
	EarliestPermittedSampleDate() objc.IObject /* cross-framework: Date */
	EnableBackgroundDeliveryForTypeFrequencyWithCompletion(type_ IHKObjectType, frequency HKUpdateFrequency, completion unsafe.Pointer)
	ExecuteQuery(query IHKQuery)
	FitzpatrickSkinTypeWithError(error_ unsafe.Pointer) IHKFitzpatrickSkinTypeObject
	GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer)
	HandleAuthorizationForExtensionWithCompletion(completion unsafe.Pointer)
	PreferredUnitsForQuantityTypesCompletion(quantityTypes unsafe.Pointer, completion foundation.IDictionary)
	RecalibrateEstimatesForSampleTypeAtDateCompletion(sampleType IHKSampleType, date objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer)
	RelateWorkoutEffortSampleWithWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer)
	RequestAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer)
	RequestPerObjectReadAuthorizationForTypePredicateCompletion(objectType IHKObjectType, predicate objc.IObject /* cross-framework: Predicate */, completion unsafe.Pointer)
	SaveObjectsWithCompletion(objects []IHKObject, completion unsafe.Pointer)
	SaveObjectWithCompletion(object IHKObject, completion unsafe.Pointer)
	StartWatchAppWithWorkoutConfigurationCompletion(workoutConfiguration IHKWorkoutConfiguration, completion unsafe.Pointer)
	StopQuery(query IHKQuery)
	SupportsHealthRecords() bool
	UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer)
	WheelchairUseWithError(error_ unsafe.Pointer) IHKWheelchairUseObject
}

// The access point for all data managed by HealthKit.
//
// Use a object to request permission to share or read HealthKit data. After you have permission, you can use the HealthKit store to save new samples to the store, or to manage the samples that your app saved. Additionally, you can use the HealthKit store to start, stop, and manage queries. For more information, see .


// The access point for all data managed by HealthKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore
type HKHealthStore struct {
	objectivec.Object
}

// HKHealthStoreFrom constructs a [HKHealthStore] from an unsafe.Pointer.
//
// The access point for all data managed by HealthKit.
func HKHealthStoreFrom(ptr unsafe.Pointer) HKHealthStore {
	return HKHealthStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKHealthStoreClass) Alloc() HKHealthStore {
	rv := objc.Send[HKHealthStore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKHealthStoreClass) New() HKHealthStore {
	rv := objc.Send[HKHealthStore](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHealthStore) Init() HKHealthStore {
	rv := objc.Send[HKHealthStore](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHealthStore) Autorelease() HKHealthStore {
	rv := objc.Send[HKHealthStore](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHealthStore creates a new HKHealthStore instance.
func NewHKHealthStore() HKHealthStore {
	return getHKHealthStoreClass().New()
}



// Returns a Boolean value that indicates whether HealthKit is available on this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/isHealthDataAvailable()
func (hc _HKHealthStoreClass) IsHealthDataAvailable() bool {
	rv := objc.Send[bool](objc.ID(hc.class), objc.Sel("isHealthDataAvailable"))
	return rv
}


// Returns the activity move mode for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/activityMoveMode()
func (h_ HKHealthStore) ActivityMoveModeWithError(error_ unsafe.Pointer) IHKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](h_.ID, objc.Sel("activityMoveModeWithError:"), error_)
	return rv
}


// Returns the app’s authorization status for sharing the specified data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/authorizationStatus(for:)
func (h_ HKHealthStore) AuthorizationStatusForType(type_ IHKObjectType) HKAuthorizationStatus {
	rv := objc.Send[HKAuthorizationStatus](h_.ID, objc.Sel("authorizationStatusForType:"), type_)
	return rv
}


// Reads someone’s biological sex from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/biologicalSex()
func (h_ HKHealthStore) BiologicalSexWithError(error_ unsafe.Pointer) IHKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](h_.ID, objc.Sel("biologicalSexWithError:"), error_)
	return rv
}


// Reads the user’s blood type from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/bloodType()
func (h_ HKHealthStore) BloodTypeWithError(error_ unsafe.Pointer) IHKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](h_.ID, objc.Sel("bloodTypeWithError:"), error_)
	return rv
}


// Reads the user’s date of birth from the HealthKit store as date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/dateOfBirthComponents()
func (h_ HKHealthStore) DateOfBirthComponentsWithError(error_ unsafe.Pointer) objc.IObject /* cross-framework: DateComponents */ {
	rv := objc.Send[foundation.DateComponents](h_.ID, objc.Sel("dateOfBirthComponentsWithError:"), error_)
	return rv
}


// Deletes the specified objects from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/delete(_:withCompletion:)-17hzm
func (h_ HKHealthStore) DeleteObjectsWithCompletion(objects []IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObjects:withCompletion:"), objects, completion)
}


// Deletes the specified object from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/delete(_:withCompletion:)-78l1m
func (h_ HKHealthStore) DeleteObjectWithCompletion(object IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObject:withCompletion:"), object, completion)
}


// Deletes objects saved by this application that match the provided type and predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/deleteObjects(of:predicate:withCompletion:)
func (h_ HKHealthStore) DeleteObjectsOfTypePredicateWithCompletion(objectType IHKObjectType, predicate objc.IObject /* cross-framework: Predicate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObjectsOfType:predicate:withCompletion:"), objectType, predicate, completion)
}


// Disables all background deliveries of update notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/disableAllBackgroundDelivery(completion:)
func (h_ HKHealthStore) DisableAllBackgroundDeliveryWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableAllBackgroundDeliveryWithCompletion:"), completion)
}


// Disables background deliveries of update notifications for the specified data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/disableBackgroundDelivery(for:withCompletion:)
func (h_ HKHealthStore) DisableBackgroundDeliveryForTypeWithCompletion(type_ IHKObjectType, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableBackgroundDeliveryForType:withCompletion:"), type_, completion)
}


// Returns the earliest date permitted for samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/earliestPermittedSampleDate()
func (h_ HKHealthStore) EarliestPermittedSampleDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("earliestPermittedSampleDate"))
	return rv
}


// Enables the delivery of updates to an app running in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/enableBackgroundDelivery(for:frequency:withCompletion:)
func (h_ HKHealthStore) EnableBackgroundDeliveryForTypeFrequencyWithCompletion(type_ IHKObjectType, frequency HKUpdateFrequency, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("enableBackgroundDeliveryForType:frequency:withCompletion:"), type_, frequency, completion)
}


// Starts executing the provided query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/execute(_:)
func (h_ HKHealthStore) ExecuteQuery(query IHKQuery) {
	objc.Send[objc.ID](h_.ID, objc.Sel("executeQuery:"), query)
}


// Reads the user’s Fitzpatrick Skin Type from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/fitzpatrickSkinType()
func (h_ HKHealthStore) FitzpatrickSkinTypeWithError(error_ unsafe.Pointer) IHKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](h_.ID, objc.Sel("fitzpatrickSkinTypeWithError:"), error_)
	return rv
}


// Indicates whether the system presents the user with a permission sheet if your app requests authorization for the provided types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/getRequestStatusForAuthorization(toShare:read:completion:)
func (h_ HKHealthStore) GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getRequestStatusForAuthorizationToShareTypes:readTypes:completion:"), typesToShare, typesToRead, completion)
}


// Requests permission to save and read the data types specified by an extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/handleAuthorizationForExtension(completion:)
func (h_ HKHealthStore) HandleAuthorizationForExtensionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("handleAuthorizationForExtensionWithCompletion:"), completion)
}


// Returns the user’s preferred units for the given quantity types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/preferredUnits(for:completion:)
func (h_ HKHealthStore) PreferredUnitsForQuantityTypesCompletion(quantityTypes unsafe.Pointer, completion foundation.IDictionary) {
	objc.Send[objc.ID](h_.ID, objc.Sel("preferredUnitsForQuantityTypes:completion:"), quantityTypes, completion)
}


// Recalibrates the prediction algorithm used to calculate the specified sample type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/recalibrateEstimates(sampleType:date:completion:)
func (h_ HKHealthStore) RecalibrateEstimatesForSampleTypeAtDateCompletion(sampleType IHKSampleType, date objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("recalibrateEstimatesForSampleType:atDate:completion:"), sampleType, date, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/relateWorkoutEffortSample(_:with:activity:completion:)
func (h_ HKHealthStore) RelateWorkoutEffortSampleWithWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("relateWorkoutEffortSample:withWorkout:activity:completion:"), sample, workout, activity, completion)
}


// Requests permission to save and read the specified data types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/requestAuthorization(toShare:read:completion:)
func (h_ HKHealthStore) RequestAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("requestAuthorizationToShareTypes:readTypes:completion:"), typesToShare, typesToRead, completion)
}


// Asynchronously requests permission to read a data type that requires per-object authorization (such as vision prescriptions).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/requestPerObjectReadAuthorization(for:predicate:completion:)
func (h_ HKHealthStore) RequestPerObjectReadAuthorizationForTypePredicateCompletion(objectType IHKObjectType, predicate objc.IObject /* cross-framework: Predicate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("requestPerObjectReadAuthorizationForType:predicate:completion:"), objectType, predicate, completion)
}


// Saves an array of objects to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/save(_:withCompletion:)-47iwb
func (h_ HKHealthStore) SaveObjectsWithCompletion(objects []IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("saveObjects:withCompletion:"), objects, completion)
}


// Saves the provided object to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/save(_:withCompletion:)-6fmtg
func (h_ HKHealthStore) SaveObjectWithCompletion(object IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("saveObject:withCompletion:"), object, completion)
}


// Launches or wakes the companion watchOS app to create a new workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/startWatchApp(with:completion:)
func (h_ HKHealthStore) StartWatchAppWithWorkoutConfigurationCompletion(workoutConfiguration IHKWorkoutConfiguration, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startWatchAppWithWorkoutConfiguration:completion:"), workoutConfiguration, completion)
}


// Stops a long-running query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/stop(_:)
func (h_ HKHealthStore) StopQuery(query IHKQuery) {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopQuery:"), query)
}


// Returns a Boolean value that indicates whether the current device supports clinical records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/supportsHealthRecords()
func (h_ HKHealthStore) SupportsHealthRecords() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("supportsHealthRecords"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/unrelateWorkoutEffortSample(_:from:activity:completion:)
func (h_ HKHealthStore) UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("unrelateWorkoutEffortSample:fromWorkout:activity:completion:"), sample, workout, activity, completion)
}


// Reads the user’s wheelchair use from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/wheelchairUse()
func (h_ HKHealthStore) WheelchairUseWithError(error_ unsafe.Pointer) IHKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](h_.ID, objc.Sel("wheelchairUseWithError:"), error_)
	return rv
}


