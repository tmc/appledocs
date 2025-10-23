// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

// Enum types and constants
// HKAppleSleepingBreathingDisturbancesClassification enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification
type HKAppleSleepingBreathingDisturbancesClassification uint

// HKAppleWalkingSteadinessClassification - A classification of a score based on the steadiness of the user’s gait.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassification
type HKAppleWalkingSteadinessClassification uint

// HKAudiogramConductionType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramConductionType
type HKAudiogramConductionType uint

// HKAudiogramSensitivityTestSide enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTestSide
type HKAudiogramSensitivityTestSide uint

// HKAuthorizationRequestStatus - Values that indicate whether your app needs to request authorization from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationRequestStatus
type HKAuthorizationRequestStatus uint

// HKAuthorizationStatus - Constants indicating the authorization status for a particular data type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationStatus
type HKAuthorizationStatus uint

// HKBiologicalSex - Constants indicating the user’s sex.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex
type HKBiologicalSex uint

// HKBloodType - Constants indicating the user’s blood type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType
type HKBloodType uint

// HKCategoryValueSleepAnalysis - Categories that represent the result of a sleep analysis.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis
type HKCategoryValueSleepAnalysis uint

// HKCategoryValueVaginalBleeding enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding
type HKCategoryValueVaginalBleeding uint

// HKErrorCode - Error codes returned by HealthKit.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code
type HKErrorCode uint

const (
	// HKErrorAnotherWorkoutSessionStarted - Another app started a workout session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorAnotherWorkoutSessionStarted
	HKErrorAnotherWorkoutSessionStarted HKErrorCode = 0
	// HKErrorAuthorizationDenied - The user hasn’t given the app permission to save data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorAuthorizationDenied
	HKErrorAuthorizationDenied HKErrorCode = 0
	// HKErrorAuthorizationNotDetermined - The app hasn’t yet asked the user for the authorization required to complete the task.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorAuthorizationNotDetermined
	HKErrorAuthorizationNotDetermined HKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorBackgroundWorkoutSessionNotAllowed
	HKErrorBackgroundWorkoutSessionNotAllowed HKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorDataSizeExceeded
	HKErrorDataSizeExceeded HKErrorCode = 0
	// HKErrorDatabaseInaccessible - The HealthKit data is unavailable because it’s protected and the device is locked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorDatabaseInaccessible
	HKErrorDatabaseInaccessible HKErrorCode = 0
	// HKErrorHealthDataRestricted - A Mobile Device Management (MDM) profile restricts the use of HealthKit on this device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorHealthDataRestricted
	HKErrorHealthDataRestricted HKErrorCode = 0
	// HKErrorHealthDataUnavailable - HealthKit accessed on an unsupported device, such as an iPad.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorHealthDataUnavailable
	HKErrorHealthDataUnavailable HKErrorCode = 0
	// HKErrorInvalidArgument - The app passed an invalid argument to the HealthKit API.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorInvalidArgument
	HKErrorInvalidArgument HKErrorCode = 0
	// HKErrorNoData - Data is unavailable for the requested query and predicate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorNoData
	HKErrorNoData HKErrorCode = 0
	// HKErrorNotPermissibleForGuestUserMode - The app attempted to write HealthKit data while in a Guest User session in visionOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorNotPermissibleForGuestUserMode
	HKErrorNotPermissibleForGuestUserMode HKErrorCode = 0
	// HKErrorRequiredAuthorizationDenied - The user hasn’t granted the application authorization to access all the required clinical record types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorRequiredAuthorizationDenied
	HKErrorRequiredAuthorizationDenied HKErrorCode = 0
	// HKErrorUserCanceled - The user canceled the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorUserCanceled
	HKErrorUserCanceled HKErrorCode = 0
	// HKErrorUserExitedWorkoutSession - The user exited your application while a workout session was running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorUserExitedWorkoutSession
	HKErrorUserExitedWorkoutSession HKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorWorkoutActivityNotAllowed
	HKErrorWorkoutActivityNotAllowed HKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/unknownError
	HKUnknownError HKErrorCode = 0
	// HKNoError - No error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKErrorCode/HKNoError
	HKNoError HKErrorCode = 0
)

// HKFitzpatrickSkinType - Categories representing the user’s skin type based on the Fitzpatrick scale.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType
type HKFitzpatrickSkinType uint

// HKGAD7AssessmentAnswer enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer
type HKGAD7AssessmentAnswer uint

// HKGAD7AssessmentRisk enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum
type HKGAD7AssessmentRisk uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/noneToMinimal
	HKGAD7AssessmentRiskNoneToMinimal HKGAD7AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/severe
	HKGAD7AssessmentRiskSevere HKGAD7AssessmentRisk = 0
)

// HKMedicationDoseEventLogStatus - The statuses the system assigns to a logged medication dose event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum
type HKMedicationDoseEventLogStatus uint

const (
	// HKMedicationDoseEventLogStatusNotInteracted - The person doesn’t interact with a scheduled medication reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/notInteracted
	HKMedicationDoseEventLogStatusNotInteracted HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusNotLogged - The person undoes a previously logged medication status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/notLogged
	HKMedicationDoseEventLogStatusNotLogged HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusNotificationNotSent - The system assigns this status when it fails to deliver a scheduled medication notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/notificationNotSent
	HKMedicationDoseEventLogStatusNotificationNotSent HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusSkipped - The person logs that they skipped the medication dose.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/skipped
	HKMedicationDoseEventLogStatusSkipped HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusSnoozed - The person snoozes a scheduled medication notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/snoozed
	HKMedicationDoseEventLogStatusSnoozed HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusTaken - The person logs that they took the medication dose.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/taken
	HKMedicationDoseEventLogStatusTaken HKMedicationDoseEventLogStatus = 0
)

// HKMedicationDoseEventScheduleType - The kind of schedule the system associates with a logged medication dose event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/ScheduleType-swift.enum
type HKMedicationDoseEventScheduleType uint

const (
	// HKMedicationDoseEventScheduleTypeAsNeeded - The person logged this dose event ad-hoc, outside of any scheduled reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/ScheduleType-swift.enum/asNeeded
	HKMedicationDoseEventScheduleTypeAsNeeded HKMedicationDoseEventScheduleType = 0
	// HKMedicationDoseEventScheduleTypeSchedule - The person logged this dose event in response to a scheduled medication reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/ScheduleType-swift.enum/schedule
	HKMedicationDoseEventScheduleTypeSchedule HKMedicationDoseEventScheduleType = 0
)

// HKPHQ9AssessmentAnswer enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer
type HKPHQ9AssessmentAnswer uint

// HKPHQ9AssessmentRisk enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum
type HKPHQ9AssessmentRisk uint

// HKStateOfMindAssociation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association
type HKStateOfMindAssociation uint

// HKStateOfMindKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Kind-swift.enum
type HKStateOfMindKind uint

// HKStateOfMindLabel enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label
type HKStateOfMindLabel uint

// HKStateOfMindValenceClassification enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum
type HKStateOfMindValenceClassification uint

// HKStatisticsOptions - Options for specifying the statistic to calculate.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions
type HKStatisticsOptions uint

const (
	// HKStatisticsOptionSeparateBySource - An option indicating that the system calculates the specified statistics separately for each source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/separateBySource
	HKStatisticsOptionSeparateBySource HKStatisticsOptions = 0
)

// HKUpdateFrequency - Constants that determine how often the system launches your app in response to changes to HealthKit data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency
type HKUpdateFrequency uint

const (
	// HKUpdateFrequencyHourly - The system launches your app at most once an hour in response to changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/hourly
	HKUpdateFrequencyHourly HKUpdateFrequency = 0
	// HKUpdateFrequencyWeekly - The system launches your app at most once per week in response to changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/weekly
	HKUpdateFrequencyWeekly HKUpdateFrequency = 0
)

// HKWheelchairUse - Constants indicating the user’s wheelchair use.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse
type HKWheelchairUse uint

// HKWorkoutActivityType - The type of activity performed during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType
type HKWorkoutActivityType uint

const (
	// HKWorkoutActivityTypeSwimming - The constant for swimming.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/swimming
	HKWorkoutActivityTypeSwimming HKWorkoutActivityType = 0
)

// HKWorkoutEffortRelationshipQueryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQueryOptions
type HKWorkoutEffortRelationshipQueryOptions uint

// HKWorkoutSessionState - A workout session’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState
type HKWorkoutSessionState uint

const (
	// HKWorkoutSessionStateEnded - The workout session has ended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/ended
	HKWorkoutSessionStateEnded HKWorkoutSessionState = 0
)

// HKWorkoutSessionType - The type of session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionType
type HKWorkoutSessionType uint


