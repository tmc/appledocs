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

/* debug [class.gen.go]: Generating class HKHealthStore */


/* debug [class_header]: Header for HKHealthStore */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKHealthStore */
// An interface definition for the [HKHealthStore] class.
type IHKHealthStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKHealthStore */
	// properties:
	WorkoutSessionMirroringStartHandler() func(unsafe.Pointer)
	SetWorkoutSessionMirroringStartHandler(value func(unsafe.Pointer))
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKHealthStore */
	// methods:
	ActivityMoveModeWithError(error_ objectivec.IObject) IHKActivityMoveModeObject
	AuthorizationStatusForType(type_ IHKObjectType) HKAuthorizationStatus
	BiologicalSexWithError(error_ objectivec.IObject) IHKBiologicalSexObject
	BloodTypeWithError(error_ objectivec.IObject) IHKBloodTypeObject
	DateOfBirthComponentsWithError(error_ objectivec.IObject) foundation.DateComponents
	DeleteObjectsWithCompletion(objects []HKObject, completion unsafe.Pointer)
	DeleteObjectWithCompletion(object IHKObject, completion unsafe.Pointer)
	DeleteObjectsOfTypePredicateWithCompletion(objectType IHKObjectType, predicate foundation.Predicate, completion unsafe.Pointer)
	DisableAllBackgroundDeliveryWithCompletion(completion unsafe.Pointer)
	DisableBackgroundDeliveryForTypeWithCompletion(type_ IHKObjectType, completion unsafe.Pointer)
	EarliestPermittedSampleDate() foundation.Date
	EnableBackgroundDeliveryForTypeFrequencyWithCompletion(type_ IHKObjectType, frequency HKUpdateFrequency, completion unsafe.Pointer)
	ExecuteQuery(query IHKQuery)
	FitzpatrickSkinTypeWithError(error_ objectivec.IObject) IHKFitzpatrickSkinTypeObject
	GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer)
	HandleAuthorizationForExtensionWithCompletion(completion unsafe.Pointer)
	PreferredUnitsForQuantityTypesCompletion(quantityTypes unsafe.Pointer, completion unsafe.Pointer)
	RecalibrateEstimatesForSampleTypeAtDateCompletion(sampleType IHKSampleType, date objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer)
	RelateWorkoutEffortSampleWithWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer)
	RequestAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer)
	RequestPerObjectReadAuthorizationForTypePredicateCompletion(objectType IHKObjectType, predicate foundation.Predicate, completion unsafe.Pointer)
	SaveObjectsWithCompletion(objects []HKObject, completion unsafe.Pointer)
	SaveObjectWithCompletion(object IHKObject, completion unsafe.Pointer)
	StartWatchAppWithWorkoutConfigurationCompletion(workoutConfiguration IHKWorkoutConfiguration, completion unsafe.Pointer)
	StopQuery(query IHKQuery)
	SupportsHealthRecords() bool
	UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer)
	WheelchairUseWithError(error_ objectivec.IObject) IHKWheelchairUseObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKHealthStore */
// Alloc allocates a new instance without initialization.
func (hc _HKHealthStoreClass) Alloc() HKHealthStore {
	rv := objc.Send[HKHealthStore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKHealthStore */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKHealthStore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKHealthStore */

// Returns a Boolean value that indicates whether HealthKit is available on this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/isHealthDataAvailable()
func (hc _HKHealthStoreClass) IsHealthDataAvailable() bool {
	rv := objc.Send[bool](objc.ID(hc.class), objc.Sel("isHealthDataAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsHealthDataAvailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKHealthStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKHealthStore */

// Returns the activity move mode for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/activityMoveMode()
func (h_ HKHealthStore) ActivityMoveModeWithError(error_ objectivec.IObject) IHKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](h_.ID, objc.Sel("activityMoveModeWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: ActivityMoveModeWithError */


// Returns the app’s authorization status for sharing the specified data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/authorizationStatus(for:)
func (h_ HKHealthStore) AuthorizationStatusForType(type_ IHKObjectType) HKAuthorizationStatus {
	rv := objc.Send[HKAuthorizationStatus](h_.ID, objc.Sel("authorizationStatusForType:"), type_)
	return rv
}/* debug [instance_methods/method]: AuthorizationStatusForType */


// Reads someone’s biological sex from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/biologicalSex()
func (h_ HKHealthStore) BiologicalSexWithError(error_ objectivec.IObject) IHKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](h_.ID, objc.Sel("biologicalSexWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: BiologicalSexWithError */


// Reads the user’s blood type from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/bloodType()
func (h_ HKHealthStore) BloodTypeWithError(error_ objectivec.IObject) IHKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](h_.ID, objc.Sel("bloodTypeWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: BloodTypeWithError */


// Reads the user’s date of birth from the HealthKit store as date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/dateOfBirthComponents()
func (h_ HKHealthStore) DateOfBirthComponentsWithError(error_ objectivec.IObject) foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](h_.ID, objc.Sel("dateOfBirthComponentsWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: DateOfBirthComponentsWithError */


// Deletes the specified objects from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/delete(_:withCompletion:)-17hzm
func (h_ HKHealthStore) DeleteObjectsWithCompletion(objects []HKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObjects:withCompletion:"), objects, completion)
}/* debug [instance_methods/method]: DeleteObjectsWithCompletion */


// Deletes the specified object from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/delete(_:withCompletion:)-78l1m
func (h_ HKHealthStore) DeleteObjectWithCompletion(object IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObject:withCompletion:"), object, completion)
}/* debug [instance_methods/method]: DeleteObjectWithCompletion */


// Deletes objects saved by this application that match the provided type and predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/deleteObjects(of:predicate:withCompletion:)
func (h_ HKHealthStore) DeleteObjectsOfTypePredicateWithCompletion(objectType IHKObjectType, predicate foundation.Predicate, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteObjectsOfType:predicate:withCompletion:"), objectType, predicate, completion)
}/* debug [instance_methods/method]: DeleteObjectsOfTypePredicateWithCompletion */


// Disables all background deliveries of update notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/disableAllBackgroundDelivery(completion:)
func (h_ HKHealthStore) DisableAllBackgroundDeliveryWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableAllBackgroundDeliveryWithCompletion:"), completion)
}/* debug [instance_methods/method]: DisableAllBackgroundDeliveryWithCompletion */


// Disables background deliveries of update notifications for the specified data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/disableBackgroundDelivery(for:withCompletion:)
func (h_ HKHealthStore) DisableBackgroundDeliveryForTypeWithCompletion(type_ IHKObjectType, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("disableBackgroundDeliveryForType:withCompletion:"), type_, completion)
}/* debug [instance_methods/method]: DisableBackgroundDeliveryForTypeWithCompletion */


// Returns the earliest date permitted for samples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/earliestPermittedSampleDate()
func (h_ HKHealthStore) EarliestPermittedSampleDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("earliestPermittedSampleDate"))
	return rv
}/* debug [instance_methods/method]: EarliestPermittedSampleDate */


// Enables the delivery of updates to an app running in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/enableBackgroundDelivery(for:frequency:withCompletion:)
func (h_ HKHealthStore) EnableBackgroundDeliveryForTypeFrequencyWithCompletion(type_ IHKObjectType, frequency HKUpdateFrequency, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("enableBackgroundDeliveryForType:frequency:withCompletion:"), type_, frequency, completion)
}/* debug [instance_methods/method]: EnableBackgroundDeliveryForTypeFrequencyWithCompletion */


// Starts executing the provided query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/execute(_:)
func (h_ HKHealthStore) ExecuteQuery(query IHKQuery) {
	objc.Send[objc.ID](h_.ID, objc.Sel("executeQuery:"), query)
}/* debug [instance_methods/method]: ExecuteQuery */


// Reads the user’s Fitzpatrick Skin Type from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/fitzpatrickSkinType()
func (h_ HKHealthStore) FitzpatrickSkinTypeWithError(error_ objectivec.IObject) IHKFitzpatrickSkinTypeObject {
	rv := objc.Send[HKFitzpatrickSkinTypeObject](h_.ID, objc.Sel("fitzpatrickSkinTypeWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: FitzpatrickSkinTypeWithError */


// Indicates whether the system presents the user with a permission sheet if your app requests authorization for the provided types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/getRequestStatusForAuthorization(toShare:read:completion:)
func (h_ HKHealthStore) GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getRequestStatusForAuthorizationToShareTypes:readTypes:completion:"), typesToShare, typesToRead, completion)
}/* debug [instance_methods/method]: GetRequestStatusForAuthorizationToShareTypesReadTypesCompletion */


// Requests permission to save and read the data types specified by an extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/handleAuthorizationForExtension(completion:)
func (h_ HKHealthStore) HandleAuthorizationForExtensionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("handleAuthorizationForExtensionWithCompletion:"), completion)
}/* debug [instance_methods/method]: HandleAuthorizationForExtensionWithCompletion */


// Returns the user’s preferred units for the given quantity types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/preferredUnits(for:completion:)
func (h_ HKHealthStore) PreferredUnitsForQuantityTypesCompletion(quantityTypes unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("preferredUnitsForQuantityTypes:completion:"), quantityTypes, completion)
}/* debug [instance_methods/method]: PreferredUnitsForQuantityTypesCompletion */


// Recalibrates the prediction algorithm used to calculate the specified sample type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/recalibrateEstimates(sampleType:date:completion:)
func (h_ HKHealthStore) RecalibrateEstimatesForSampleTypeAtDateCompletion(sampleType IHKSampleType, date objc.IObject /* cross-framework: NSDate */, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("recalibrateEstimatesForSampleType:atDate:completion:"), sampleType, date, completion)
}/* debug [instance_methods/method]: RecalibrateEstimatesForSampleTypeAtDateCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/relateWorkoutEffortSample(_:with:activity:completion:)
func (h_ HKHealthStore) RelateWorkoutEffortSampleWithWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("relateWorkoutEffortSample:withWorkout:activity:completion:"), sample, workout, activity, completion)
}/* debug [instance_methods/method]: RelateWorkoutEffortSampleWithWorkoutActivityCompletion */


// Requests permission to save and read the specified data types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/requestAuthorization(toShare:read:completion:)
func (h_ HKHealthStore) RequestAuthorizationToShareTypesReadTypesCompletion(typesToShare unsafe.Pointer, typesToRead unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("requestAuthorizationToShareTypes:readTypes:completion:"), typesToShare, typesToRead, completion)
}/* debug [instance_methods/method]: RequestAuthorizationToShareTypesReadTypesCompletion */


// Asynchronously requests permission to read a data type that requires per-object authorization (such as vision prescriptions).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/requestPerObjectReadAuthorization(for:predicate:completion:)
func (h_ HKHealthStore) RequestPerObjectReadAuthorizationForTypePredicateCompletion(objectType IHKObjectType, predicate foundation.Predicate, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("requestPerObjectReadAuthorizationForType:predicate:completion:"), objectType, predicate, completion)
}/* debug [instance_methods/method]: RequestPerObjectReadAuthorizationForTypePredicateCompletion */


// Saves an array of objects to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/save(_:withCompletion:)-47iwb
func (h_ HKHealthStore) SaveObjectsWithCompletion(objects []HKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("saveObjects:withCompletion:"), objects, completion)
}/* debug [instance_methods/method]: SaveObjectsWithCompletion */


// Saves the provided object to the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/save(_:withCompletion:)-6fmtg
func (h_ HKHealthStore) SaveObjectWithCompletion(object IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("saveObject:withCompletion:"), object, completion)
}/* debug [instance_methods/method]: SaveObjectWithCompletion */


// Launches or wakes the companion watchOS app to create a new workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/startWatchApp(with:completion:)
func (h_ HKHealthStore) StartWatchAppWithWorkoutConfigurationCompletion(workoutConfiguration IHKWorkoutConfiguration, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("startWatchAppWithWorkoutConfiguration:completion:"), workoutConfiguration, completion)
}/* debug [instance_methods/method]: StartWatchAppWithWorkoutConfigurationCompletion */


// Stops a long-running query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/stop(_:)
func (h_ HKHealthStore) StopQuery(query IHKQuery) {
	objc.Send[objc.ID](h_.ID, objc.Sel("stopQuery:"), query)
}/* debug [instance_methods/method]: StopQuery */


// Returns a Boolean value that indicates whether the current device supports clinical records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/supportsHealthRecords()
func (h_ HKHealthStore) SupportsHealthRecords() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("supportsHealthRecords"))
	return rv
}/* debug [instance_methods/method]: SupportsHealthRecords */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/unrelateWorkoutEffortSample(_:from:activity:completion:)
func (h_ HKHealthStore) UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion(sample IHKSample, workout IHKWorkout, activity IHKWorkoutActivity, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("unrelateWorkoutEffortSample:fromWorkout:activity:completion:"), sample, workout, activity, completion)
}/* debug [instance_methods/method]: UnrelateWorkoutEffortSampleFromWorkoutActivityCompletion */


// Reads the user’s wheelchair use from the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/wheelchairUse()
func (h_ HKHealthStore) WheelchairUseWithError(error_ objectivec.IObject) IHKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](h_.ID, objc.Sel("wheelchairUseWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: WheelchairUseWithError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKHealthStore */

// A block that the system calls when it starts a mirrored workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/workoutSessionMirroringStartHandler
func (h_ HKHealthStore) WorkoutSessionMirroringStartHandler() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](h_.ID, objc.Sel("workoutSessionMirroringStartHandler"))
	return rv
}/* debug [instance_properties/getter]: workoutSessionMirroringStartHandler */


// A block that the system calls when it starts a mirrored workout session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthStore/workoutSessionMirroringStartHandler
func (h_ HKHealthStore) SetWorkoutSessionMirroringStartHandler(value func(unsafe.Pointer)) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWorkoutSessionMirroringStartHandler:"), value)
}/* debug [instance_properties/setter]: workoutSessionMirroringStartHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKHealthStore */


