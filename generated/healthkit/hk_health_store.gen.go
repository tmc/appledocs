// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ActivityMoveModeWithError(error_ unsafe.Pointer) unsafe.Pointer
	AddSamplesToWorkoutCompletion(samples unsafe.Pointer, workout unsafe.Pointer, completion unsafe.Pointer)
	AuthorizationStatusForType(type_ unsafe.Pointer) unsafe.Pointer
	BiologicalSexWithError(error_ unsafe.Pointer) unsafe.Pointer
	BloodTypeWithError(error_ unsafe.Pointer) unsafe.Pointer
	DateOfBirthWithError(error_ unsafe.Pointer) unsafe.Pointer
	DateOfBirthComponentsWithError(error_ unsafe.Pointer) unsafe.Pointer
	DeleteObjectsWithCompletion(objects unsafe.Pointer, completion unsafe.Pointer)
	DeleteObjectWithCompletion(object unsafe.Pointer, completion unsafe.Pointer)
	DeleteObjectsOfTypePredicateWithCompletion(objectType unsafe.Pointer, predicate unsafe.Pointer, completion unsafe.Pointer)
	DisableAllBackgroundDeliveryWithCompletion(completion unsafe.Pointer)
	DisableBackgroundDeliveryForTypeWithCompletion(type_ unsafe.Pointer, completion unsafe.Pointer)
	EarliestPermittedSampleDate() unsafe.Pointer
	EnableBackgroundDeliveryForTypeFrequencyWithCompletion(type_ unsafe.Pointer, frequency unsafe.Pointer, completion unsafe.Pointer)
	EndWorkoutSession(workoutSession unsafe.Pointer)
	ExecuteQuery(query unsafe.Pointer)
	FitzpatrickSkinTypeWithError(error_ unsafe.Pointer) unsafe.Pointer
	GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer)
	HandleAuthorizationForExtensionWithCompletion(completion unsafe.Pointer)
	PauseWorkoutSession(workoutSession unsafe.Pointer)
	PreferredUnitsForQuantityTypesCompletion(quantityTypes unsafe.Pointer, completion unsafe.Pointer)
	RecalibrateEstimatesForSampleTypeAtDateCompletion(sampleType unsafe.Pointer, date unsafe.Pointer, completion unsafe.Pointer)
	RecoverActiveWorkoutSessionWithCompletion(completion unsafe.Pointer)
	RelateWorkoutEffortSampleWithWorkoutActivityCompletion(sample unsafe.Pointer, workout unsafe.Pointer, activity unsafe.Pointer, completion unsafe.Pointer)
	RequestAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer)
	RequestPerObjectReadAuthorizationForTypePredicateCompletion(objectType unsafe.Pointer, predicate unsafe.Pointer, completion unsafe.Pointer)
	ResumeWorkoutSession(workoutSession unsafe.Pointer)
	SaveObjectsWithCompletion(objects unsafe.Pointer, completion unsafe.Pointer)
	SaveObjectWithCompletion(object unsafe.Pointer, completion unsafe.Pointer)
	SplitTotalEnergyStartDateEndDateResultsHandler(totalEnergy unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, resultsHandler unsafe.Pointer)
	StartWorkoutSession(workoutSession unsafe.Pointer)
	StartWatchAppWithWorkoutConfigurationCompletion(workoutConfiguration unsafe.Pointer, completion unsafe.Pointer)
	StopQuery(query unsafe.Pointer)
	SupportsHealthRecords() bool
	UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion(sample unsafe.Pointer, workout unsafe.Pointer, activity unsafe.Pointer, completion unsafe.Pointer)
	WheelchairUseWithError(error_ unsafe.Pointer) unsafe.Pointer
}

// The access point for all data managed by HealthKit.
//
// Use a object to request permission to share or read HealthKit data. After you have permission, you can use the HealthKit store to save new samples to the store, or to manage the samples that your app saved. Additionally, you can use the HealthKit store to start, stop, and manage queries. For more information, see .
//
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
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/isHealthDataAvailable()
func (hc _HKHealthStoreClass) IsHealthDataAvailable() bool {
	rv := objc.Send[bool](objc.ID(hc.class), objc.Sel("isHealthDataAvailable"))
	return rv
}

// Returns the activity move mode for the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/activityMoveMode()
func (h_ HKHealthStore) ActivityMoveModeWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activityMoveModeWithError:"), error_)
	return rv
}

// Associates the provided samples with the specified workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/add(_:to:completion:)
func (h_ HKHealthStore) AddSamplesToWorkoutCompletion(samples unsafe.Pointer, workout unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addSamples:toWorkout:completion:"), samples, workout, completion)
}

// Returns the app’s authorization status for sharing the specified data type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/authorizationStatus(for:)
func (h_ HKHealthStore) AuthorizationStatusForType(type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("authorizationStatusForType:"), type_)
	return rv
}

// Reads someone’s biological sex from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/biologicalSex()
func (h_ HKHealthStore) BiologicalSexWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("biologicalSexWithError:"), error_)
	return rv
}

// Reads the user’s blood type from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/bloodType()
func (h_ HKHealthStore) BloodTypeWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("bloodTypeWithError:"), error_)
	return rv
}

// Reads the user’s date of birth from the HealthKit store as a date value.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/dateOfBirth()
func (h_ HKHealthStore) DateOfBirthWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("dateOfBirthWithError:"), error_)
	return rv
}

// Reads the user’s date of birth from the HealthKit store as date components.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/dateOfBirthComponents()
func (h_ HKHealthStore) DateOfBirthComponentsWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("dateOfBirthComponentsWithError:"), error_)
	return rv
}

// Deletes the specified objects from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/delete(_:withCompletion:)-17hzm
func (h_ HKHealthStore) DeleteObjectsWithCompletion(objects unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObjects:withCompletion:"), objects, completion)
}

// Deletes the specified object from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/delete(_:withCompletion:)-78l1m
func (h_ HKHealthStore) DeleteObjectWithCompletion(object unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObject:withCompletion:"), object, completion)
}

// Deletes objects saved by this application that match the provided type and predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/deleteObjects(of:predicate:withCompletion:)
func (h_ HKHealthStore) DeleteObjectsOfTypePredicateWithCompletion(objectType unsafe.Pointer, predicate unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObjectsOfType:predicate:withCompletion:"), objectType, predicate, completion)
}

// Disables all background deliveries of update notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/disableAllBackgroundDelivery(completion:)
func (h_ HKHealthStore) DisableAllBackgroundDeliveryWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableAllBackgroundDeliveryWithCompletion:"), completion)
}

// Disables background deliveries of update notifications for the specified data type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/disableBackgroundDelivery(for:withCompletion:)
func (h_ HKHealthStore) DisableBackgroundDeliveryForTypeWithCompletion(type_ unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableBackgroundDeliveryForType:withCompletion:"), type_, completion)
}

// Returns the earliest date permitted for samples.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/earliestPermittedSampleDate()
func (h_ HKHealthStore) EarliestPermittedSampleDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("earliestPermittedSampleDate"))
	return rv
}

// Enables the delivery of updates to an app running in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/enableBackgroundDelivery(for:frequency:withCompletion:)
func (h_ HKHealthStore) EnableBackgroundDeliveryForTypeFrequencyWithCompletion(type_ unsafe.Pointer, frequency unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("enableBackgroundDeliveryForType:frequency:withCompletion:"), type_, frequency, completion)
}

// Ends a workout session for the current app.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/end(_:)
func (h_ HKHealthStore) EndWorkoutSession(workoutSession unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("endWorkoutSession:"), workoutSession)
}

// Starts executing the provided query.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/execute(_:)
func (h_ HKHealthStore) ExecuteQuery(query unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("executeQuery:"), query)
}

// Reads the user’s Fitzpatrick Skin Type from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/fitzpatrickSkinType()
func (h_ HKHealthStore) FitzpatrickSkinTypeWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("fitzpatrickSkinTypeWithError:"), error_)
	return rv
}

// Indicates whether the system presents the user with a permission sheet if your app requests authorization for the provided types.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/getRequestStatusForAuthorization(toShare:read:completion:)
func (h_ HKHealthStore) GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getRequestStatusForAuthorizationToShareTypes:readTypes:completion:"), typesToShare, typesToRead, completion)
}

// Requests permission to save and read the data types specified by an extension.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/handleAuthorizationForExtension(completion:)
func (h_ HKHealthStore) HandleAuthorizationForExtensionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("handleAuthorizationForExtensionWithCompletion:"), completion)
}

// Pauses the provided workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/pause(_:)
func (h_ HKHealthStore) PauseWorkoutSession(workoutSession unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("pauseWorkoutSession:"), workoutSession)
}

// Returns the user’s preferred units for the given quantity types.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/preferredUnits(for:completion:)
func (h_ HKHealthStore) PreferredUnitsForQuantityTypesCompletion(quantityTypes unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("preferredUnitsForQuantityTypes:completion:"), quantityTypes, completion)
}

// Recalibrates the prediction algorithm used to calculate the specified sample type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/recalibrateEstimates(sampleType:date:completion:)
func (h_ HKHealthStore) RecalibrateEstimatesForSampleTypeAtDateCompletion(sampleType unsafe.Pointer, date unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("recalibrateEstimatesForSampleType:atDate:completion:"), sampleType, date, completion)
}

// Recovers an active workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/recoverActiveWorkoutSession(completion:)
func (h_ HKHealthStore) RecoverActiveWorkoutSessionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("recoverActiveWorkoutSessionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/relateWorkoutEffortSample(_:with:activity:completion:)
func (h_ HKHealthStore) RelateWorkoutEffortSampleWithWorkoutActivityCompletion(sample unsafe.Pointer, workout unsafe.Pointer, activity unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("relateWorkoutEffortSample:withWorkout:activity:completion:"), sample, workout, activity, completion)
}

// Requests permission to save and read the specified data types.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/requestAuthorization(toShare:read:completion:)
func (h_ HKHealthStore) RequestAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("requestAuthorizationToShareTypes:readTypes:completion:"), typesToShare, typesToRead, completion)
}

// Asynchronously requests permission to read a data type that requires per-object authorization (such as vision prescriptions).
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/requestPerObjectReadAuthorization(for:predicate:completion:)
func (h_ HKHealthStore) RequestPerObjectReadAuthorizationForTypePredicateCompletion(objectType unsafe.Pointer, predicate unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("requestPerObjectReadAuthorizationForType:predicate:completion:"), objectType, predicate, completion)
}

// Resumes the provided workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/resumeWorkoutSession(_:)
func (h_ HKHealthStore) ResumeWorkoutSession(workoutSession unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("resumeWorkoutSession:"), workoutSession)
}

// Saves an array of objects to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/save(_:withCompletion:)-47iwb
func (h_ HKHealthStore) SaveObjectsWithCompletion(objects unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("saveObjects:withCompletion:"), objects, completion)
}

// Saves the provided object to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/save(_:withCompletion:)-6fmtg
func (h_ HKHealthStore) SaveObjectWithCompletion(object unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("saveObject:withCompletion:"), object, completion)
}

// Calculates the active and resting energy burned based on the total energy burned over the given duration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/splitTotalEnergy(_:start:end:resultsHandler:)
func (h_ HKHealthStore) SplitTotalEnergyStartDateEndDateResultsHandler(totalEnergy unsafe.Pointer, startDate unsafe.Pointer, endDate unsafe.Pointer, resultsHandler unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("splitTotalEnergy:startDate:endDate:resultsHandler:"), totalEnergy, startDate, endDate, resultsHandler)
}

// Starts a workout session for the current app.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/start(_:)
func (h_ HKHealthStore) StartWorkoutSession(workoutSession unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startWorkoutSession:"), workoutSession)
}

// Launches or wakes the companion watchOS app to create a new workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/startWatchApp(with:completion:)
func (h_ HKHealthStore) StartWatchAppWithWorkoutConfigurationCompletion(workoutConfiguration unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startWatchAppWithWorkoutConfiguration:completion:"), workoutConfiguration, completion)
}

// Stops a long-running query.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/stop(_:)
func (h_ HKHealthStore) StopQuery(query unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopQuery:"), query)
}

// Returns a Boolean value that indicates whether the current device supports clinical records.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/supportsHealthRecords()
func (h_ HKHealthStore) SupportsHealthRecords() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("supportsHealthRecords"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/unrelateWorkoutEffortSample(_:from:activity:completion:)
func (h_ HKHealthStore) UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion(sample unsafe.Pointer, workout unsafe.Pointer, activity unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("unrelateWorkoutEffortSample:fromWorkout:activity:completion:"), sample, workout, activity, completion)
}

// Reads the user’s wheelchair use from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/wheelchairUse()
func (h_ HKHealthStore) WheelchairUseWithError(error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("wheelchairUseWithError:"), error_)
	return rv
}

// The view controller that presents HealthKit authorization sheets.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/authorizationViewControllerPresenter
func (h_ HKHealthStore) AuthorizationViewControllerPresenter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("authorizationViewControllerPresenter"))
	return rv
}


// SetAuthorizationViewControllerPresenter sets the value of the authorizationViewControllerPresenter property.
// The view controller that presents HealthKit authorization sheets.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/authorizationViewControllerPresenter
func (h_ HKHealthStore) SetAuthorizationViewControllerPresenter(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAuthorizationViewControllerPresenter:"), value)
}

// A block that the system calls when it starts a mirrored workout session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/workoutSessionMirroringStartHandler
func (h_ HKHealthStore) WorkoutSessionMirroringStartHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("workoutSessionMirroringStartHandler"))
	return rv
}


// SetWorkoutSessionMirroringStartHandler sets the value of the workoutSessionMirroringStartHandler property.
// A block that the system calls when it starts a mirrored workout session.

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/workoutSessionMirroringStartHandler
func (h_ HKHealthStore) SetWorkoutSessionMirroringStartHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutSessionMirroringStartHandler:"), value)
}



