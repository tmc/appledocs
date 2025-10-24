// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

/* debug [enums.gen.go]: Generating 73 enums for HealthKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum HKErrorCode (17 cases) */
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
	// HKErrorDatabaseInaccessible - The HealthKit data is unavailable because it’s protected and the device is locked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorDatabaseInaccessible
	HKErrorDatabaseInaccessible HKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKError/Code/errorDataSizeExceeded
	HKErrorDataSizeExceeded HKErrorCode = 0
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

/* debug [enums.gen.go]: Processing enum HKMedicationDoseEventLogStatus (6 cases) */
// HKMedicationDoseEventLogStatus - The statuses the system assigns to a logged medication dose event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum
type HKMedicationDoseEventLogStatus uint

const (
	// HKMedicationDoseEventLogStatusNotificationNotSent - The system assigns this status when it fails to deliver a scheduled medication notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/notificationNotSent
	HKMedicationDoseEventLogStatusNotificationNotSent HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusNotInteracted - The person doesn’t interact with a scheduled medication reminder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/notInteracted
	HKMedicationDoseEventLogStatusNotInteracted HKMedicationDoseEventLogStatus = 0
	// HKMedicationDoseEventLogStatusNotLogged - The person undoes a previously logged medication status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/LogStatus-swift.enum/notLogged
	HKMedicationDoseEventLogStatusNotLogged HKMedicationDoseEventLogStatus = 0
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

/* debug [enums.gen.go]: Processing enum HKMedicationDoseEventScheduleType (2 cases) */
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

/* debug [enums.gen.go]: Processing enum HKActivityMoveMode (2 cases) */
// HKActivityMoveMode - Constants that specify the value measured by the Move ring on the user’s device.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivityMoveMode
type HKActivityMoveMode uint

const (
	// HKActivityMoveModeActiveEnergy - A value that indicates the Move ring measures active energy burned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivityMoveMode/activeEnergy
	HKActivityMoveModeActiveEnergy HKActivityMoveMode = 0
	// HKActivityMoveModeAppleMoveTime - A value that indicates the Activity app’s Move ring measures Apple Move Time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivityMoveMode/appleMoveTime
	HKActivityMoveModeAppleMoveTime HKActivityMoveMode = 0
)

/* debug [enums.gen.go]: Processing enum HKAppleECGAlgorithmVersion (2 cases) */
// HKAppleECGAlgorithmVersion - Version numbers for the algorithm Apple Watch uses to generate an ECG reading.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleECGAlgorithmVersion
type HKAppleECGAlgorithmVersion uint

const (
	// HKAppleECGAlgorithmVersion1 - The version 1 algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleECGAlgorithmVersion/version1
	HKAppleECGAlgorithmVersion1 HKAppleECGAlgorithmVersion = 0
	// HKAppleECGAlgorithmVersion2 - The version 2 algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleECGAlgorithmVersion/version2
	HKAppleECGAlgorithmVersion2 HKAppleECGAlgorithmVersion = 0
)

/* debug [enums.gen.go]: Processing enum HKAppleSleepingBreathingDisturbancesClassification (2 cases) */
// HKAppleSleepingBreathingDisturbancesClassification enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification
type HKAppleSleepingBreathingDisturbancesClassification uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification/elevated
	HKAppleSleepingBreathingDisturbancesClassificationElevated HKAppleSleepingBreathingDisturbancesClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification/notElevated
	HKAppleSleepingBreathingDisturbancesClassificationNotElevated HKAppleSleepingBreathingDisturbancesClassification = 0
)

/* debug [enums.gen.go]: Processing enum HKAppleWalkingSteadinessClassification (3 cases) */
// HKAppleWalkingSteadinessClassification - A classification of a score based on the steadiness of the user’s gait.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassification
type HKAppleWalkingSteadinessClassification uint

const (
	// HKAppleWalkingSteadinessClassificationLow - A classification indicating that the stability of the user’s gate is below normal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassification/low
	HKAppleWalkingSteadinessClassificationLow HKAppleWalkingSteadinessClassification = 0
	// HKAppleWalkingSteadinessClassificationOK - A classification indicating that the stability of the user’s gait is within the normal range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassification/ok
	HKAppleWalkingSteadinessClassificationOK HKAppleWalkingSteadinessClassification = 0
	// HKAppleWalkingSteadinessClassificationVeryLow - A classification indicating that the stability of the user’s gate is considerably below normal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassification/veryLow
	HKAppleWalkingSteadinessClassificationVeryLow HKAppleWalkingSteadinessClassification = 0
)

/* debug [enums.gen.go]: Processing enum HKAudiogramConductionType (1 cases) */
// HKAudiogramConductionType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramConductionType
type HKAudiogramConductionType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramConductionType/air
	HKAudiogramConductionTypeAir HKAudiogramConductionType = 0
)

/* debug [enums.gen.go]: Processing enum HKAudiogramSensitivityTestSide (2 cases) */
// HKAudiogramSensitivityTestSide enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTestSide
type HKAudiogramSensitivityTestSide uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTestSide/left
	HKAudiogramSensitivityTestSideLeft HKAudiogramSensitivityTestSide = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTestSide/right
	HKAudiogramSensitivityTestSideRight HKAudiogramSensitivityTestSide = 0
)

/* debug [enums.gen.go]: Processing enum HKAuthorizationRequestStatus (3 cases) */
// HKAuthorizationRequestStatus - Values that indicate whether your app needs to request authorization from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationRequestStatus
type HKAuthorizationRequestStatus uint

const (
	// HKAuthorizationRequestStatusShouldRequest - The application has not yet requested authorization for all the specified data types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationRequestStatus/shouldRequest
	HKAuthorizationRequestStatusShouldRequest HKAuthorizationRequestStatus = 0
	// HKAuthorizationRequestStatusUnknown - The authorization request status could not be determined because an error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationRequestStatus/unknown
	HKAuthorizationRequestStatusUnknown HKAuthorizationRequestStatus = 0
	// HKAuthorizationRequestStatusUnnecessary - The application has already requested authorization for all the specified data types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationRequestStatus/unnecessary
	HKAuthorizationRequestStatusUnnecessary HKAuthorizationRequestStatus = 0
)

/* debug [enums.gen.go]: Processing enum HKAuthorizationStatus (3 cases) */
// HKAuthorizationStatus - Constants indicating the authorization status for a particular data type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationStatus
type HKAuthorizationStatus uint

const (
	// HKAuthorizationStatusNotDetermined - The user has not yet chosen to authorize access to the specified data type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationStatus/notDetermined
	HKAuthorizationStatusNotDetermined HKAuthorizationStatus = 0
	// HKAuthorizationStatusSharingAuthorized - The user has explicitly authorized your app to save data of the specified type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationStatus/sharingAuthorized
	HKAuthorizationStatusSharingAuthorized HKAuthorizationStatus = 0
	// HKAuthorizationStatusSharingDenied - The user has explicitly denied your app permission to save data of the specified type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAuthorizationStatus/sharingDenied
	HKAuthorizationStatusSharingDenied HKAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum HKBiologicalSex (4 cases) */
// HKBiologicalSex - Constants indicating the user’s sex.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex
type HKBiologicalSex uint

const (
	// HKBiologicalSexFemale - A constant indicating that the user is female.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex/female
	HKBiologicalSexFemale HKBiologicalSex = 0
	// HKBiologicalSexMale - A constant indicating that the user is male.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex/male
	HKBiologicalSexMale HKBiologicalSex = 0
	// HKBiologicalSexNotSet - A constant indicating that either the user’s biological sex characteristic type is not set, or the user has not granted your app permission to read that characteristic type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex/notSet
	HKBiologicalSexNotSet HKBiologicalSex = 0
	// HKBiologicalSexOther - A constant indicating that the user is otherwise not categorized as either male or female.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex/other
	HKBiologicalSexOther HKBiologicalSex = 0
)

/* debug [enums.gen.go]: Processing enum HKBloodGlucoseMealTime (2 cases) */
// HKBloodGlucoseMealTime - Constants indicating the timing of a blood glucose sample relative to a meal.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodGlucoseMealTime
type HKBloodGlucoseMealTime uint

const (
	// HKBloodGlucoseMealTimePostprandial - A blood glucose sample taken just after eating a meal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodGlucoseMealTime/postprandial
	HKBloodGlucoseMealTimePostprandial HKBloodGlucoseMealTime = 0
	// HKBloodGlucoseMealTimePreprandial - A blood glucose sample taken just before eating a meal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodGlucoseMealTime/preprandial
	HKBloodGlucoseMealTimePreprandial HKBloodGlucoseMealTime = 0
)

/* debug [enums.gen.go]: Processing enum HKBloodType (9 cases) */
// HKBloodType - Constants indicating the user’s blood type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType
type HKBloodType uint

const (
	// HKBloodTypeABNegative - The user has an AB– blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/abNegative
	HKBloodTypeABNegative HKBloodType = 0
	// HKBloodTypeABPositive - The user has an AB+ blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/abPositive
	HKBloodTypeABPositive HKBloodType = 0
	// HKBloodTypeANegative - The user has an A– blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/aNegative
	HKBloodTypeANegative HKBloodType = 0
	// HKBloodTypeAPositive - The user has an A+ blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/aPositive
	HKBloodTypeAPositive HKBloodType = 0
	// HKBloodTypeBNegative - The user has an B– blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/bNegative
	HKBloodTypeBNegative HKBloodType = 0
	// HKBloodTypeBPositive - The user has an B+ blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/bPositive
	HKBloodTypeBPositive HKBloodType = 0
	// HKBloodTypeNotSet - Either the user’s blood type is not set, or the user has not granted your app permission to read the blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/notSet
	HKBloodTypeNotSet HKBloodType = 0
	// HKBloodTypeONegative - The user has an O– blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/oNegative
	HKBloodTypeONegative HKBloodType = 0
	// HKBloodTypeOPositive - The user has an O+ blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/oPositive
	HKBloodTypeOPositive HKBloodType = 0
)

/* debug [enums.gen.go]: Processing enum HKBodyTemperatureSensorLocation (12 cases) */
// HKBodyTemperatureSensorLocation - Constants that indicate where on the body a temperature reading was taken.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation
type HKBodyTemperatureSensorLocation uint

const (
	// HKBodyTemperatureSensorLocationArmpit - The temperature was taken in the armpit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/armpit
	HKBodyTemperatureSensorLocationArmpit HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationBody - The temperature was taken on the body.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/body
	HKBodyTemperatureSensorLocationBody HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationEar - The temperature was taken in the ear.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/ear
	HKBodyTemperatureSensorLocationEar HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationEarDrum - The temperature was taken on the eardrum.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/earDrum
	HKBodyTemperatureSensorLocationEarDrum HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationFinger - The temperature was taken at the finger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/finger
	HKBodyTemperatureSensorLocationFinger HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationForehead - The temperature was taken on the forehead.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/forehead
	HKBodyTemperatureSensorLocationForehead HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationGastroIntestinal - The temperature was taken inside the gastrointestinal tract.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/gastroIntestinal
	HKBodyTemperatureSensorLocationGastroIntestinal HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationMouth - The temperature was taken in the mouth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/mouth
	HKBodyTemperatureSensorLocationMouth HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationOther - The temperature was taken at a location that is not otherwise in this list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/other
	HKBodyTemperatureSensorLocationOther HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationRectum - The temperature was taken in the rectum.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/rectum
	HKBodyTemperatureSensorLocationRectum HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationTemporalArtery - The temperature was taken at the temporal artery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/temporalArtery
	HKBodyTemperatureSensorLocationTemporalArtery HKBodyTemperatureSensorLocation = 0
	// HKBodyTemperatureSensorLocationToe - The temperature was taken at the toe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBodyTemperatureSensorLocation/toe
	HKBodyTemperatureSensorLocationToe HKBodyTemperatureSensorLocation = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValue (1 cases) */
// HKCategoryValue - Categories that are undefined.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValue
type HKCategoryValue uint

const (
	// HKCategoryValueNotApplicable - A category value for types that don’t have a defined value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValue/notApplicable
	HKCategoryValueNotApplicable HKCategoryValue = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueAppetiteChanges (4 cases) */
// HKCategoryValueAppetiteChanges - Categories that represent change in appetite.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges
type HKCategoryValueAppetiteChanges uint

const (
	// HKCategoryValueAppetiteChangesDecreased - The user’s appetite decreased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges/decreased
	HKCategoryValueAppetiteChangesDecreased HKCategoryValueAppetiteChanges = 0
	// HKCategoryValueAppetiteChangesIncreased - The user’s appetite increased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges/increased
	HKCategoryValueAppetiteChangesIncreased HKCategoryValueAppetiteChanges = 0
	// HKCategoryValueAppetiteChangesNoChange - No change in the user’s appetite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges/noChange
	HKCategoryValueAppetiteChangesNoChange HKCategoryValueAppetiteChanges = 0
	// HKCategoryValueAppetiteChangesUnspecified - An unspecified change in appetite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges/unspecified
	HKCategoryValueAppetiteChangesUnspecified HKCategoryValueAppetiteChanges = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueAppleStandHour (2 cases) */
// HKCategoryValueAppleStandHour - Categories that the system used to indicate whether the user stood during the sample’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleStandHour
type HKCategoryValueAppleStandHour uint

const (
	// HKCategoryValueAppleStandHourIdle - The user didn’t stand up and move for at least one continuous minute during the sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleStandHour/idle
	HKCategoryValueAppleStandHourIdle HKCategoryValueAppleStandHour = 0
	// HKCategoryValueAppleStandHourStood - The user stood up and moved for at least one continuous minute during the sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleStandHour/stood
	HKCategoryValueAppleStandHourStood HKCategoryValueAppleStandHour = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueAppleWalkingSteadinessEvent (4 cases) */
// HKCategoryValueAppleWalkingSteadinessEvent - The value of an event triggered by a reduced score for the steadiness of the user’s gait.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleWalkingSteadinessEvent
type HKCategoryValueAppleWalkingSteadinessEvent uint

const (
	// HKCategoryValueAppleWalkingSteadinessEventInitialLow - The user received a below-normal steadiness score for their gait while walking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleWalkingSteadinessEvent/initialLow
	HKCategoryValueAppleWalkingSteadinessEventInitialLow HKCategoryValueAppleWalkingSteadinessEvent = 0
	// HKCategoryValueAppleWalkingSteadinessEventInitialVeryLow - The user received a steadiness score for their gait while walking that was considerably below normal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleWalkingSteadinessEvent/initialVeryLow
	HKCategoryValueAppleWalkingSteadinessEventInitialVeryLow HKCategoryValueAppleWalkingSteadinessEvent = 0
	// HKCategoryValueAppleWalkingSteadinessEventRepeatLow - The user’s below-normal score persists over a significant period of time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleWalkingSteadinessEvent/repeatLow
	HKCategoryValueAppleWalkingSteadinessEventRepeatLow HKCategoryValueAppleWalkingSteadinessEvent = 0
	// HKCategoryValueAppleWalkingSteadinessEventRepeatVeryLow - The user’s considerably below-normal score persists over a significant period of time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppleWalkingSteadinessEvent/repeatVeryLow
	HKCategoryValueAppleWalkingSteadinessEventRepeatVeryLow HKCategoryValueAppleWalkingSteadinessEvent = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueAudioExposureEvent (1 cases) */
// HKCategoryValueAudioExposureEvent - Categories that indicate audio exposure events.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAudioExposureEvent
type HKCategoryValueAudioExposureEvent uint

const (
	// HKCategoryValueAudioExposureEventLoudEnvironment - Exposure to a loud environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAudioExposureEvent/loudEnvironment
	HKCategoryValueAudioExposureEventLoudEnvironment HKCategoryValueAudioExposureEvent = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueCervicalMucusQuality (5 cases) */
// HKCategoryValueCervicalMucusQuality - Categories that represent the user’s cervical mucus quality.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueCervicalMucusQuality
type HKCategoryValueCervicalMucusQuality uint

const (
	// HKCategoryValueCervicalMucusQualityCreamy - Creamy mucus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueCervicalMucusQuality/creamy
	HKCategoryValueCervicalMucusQualityCreamy HKCategoryValueCervicalMucusQuality = 0
	// HKCategoryValueCervicalMucusQualityDry - Little or no mucus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueCervicalMucusQuality/dry
	HKCategoryValueCervicalMucusQualityDry HKCategoryValueCervicalMucusQuality = 0
	// HKCategoryValueCervicalMucusQualityEggWhite - Mucus the color and consistency of egg whites.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueCervicalMucusQuality/eggWhite
	HKCategoryValueCervicalMucusQualityEggWhite HKCategoryValueCervicalMucusQuality = 0
	// HKCategoryValueCervicalMucusQualitySticky - Sticky mucus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueCervicalMucusQuality/sticky
	HKCategoryValueCervicalMucusQualitySticky HKCategoryValueCervicalMucusQuality = 0
	// HKCategoryValueCervicalMucusQualityWatery - Watery mucus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueCervicalMucusQuality/watery
	HKCategoryValueCervicalMucusQualityWatery HKCategoryValueCervicalMucusQuality = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueContraceptive (7 cases) */
// HKCategoryValueContraceptive - The type of contraceptive.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive
type HKCategoryValueContraceptive uint

const (
	// HKCategoryValueContraceptiveImplant - A contraceptive implant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/implant
	HKCategoryValueContraceptiveImplant HKCategoryValueContraceptive = 0
	// HKCategoryValueContraceptiveInjection - An injectable contraceptive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/injection
	HKCategoryValueContraceptiveInjection HKCategoryValueContraceptive = 0
	// HKCategoryValueContraceptiveIntrauterineDevice - An intrauterine device (IUD).
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/intrauterineDevice
	HKCategoryValueContraceptiveIntrauterineDevice HKCategoryValueContraceptive = 0
	// HKCategoryValueContraceptiveIntravaginalRing - A contraceptive intravaginal ring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/intravaginalRing
	HKCategoryValueContraceptiveIntravaginalRing HKCategoryValueContraceptive = 0
	// HKCategoryValueContraceptiveOral - An oral contraceptive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/oral
	HKCategoryValueContraceptiveOral HKCategoryValueContraceptive = 0
	// HKCategoryValueContraceptivePatch - A contraceptive patch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/patch
	HKCategoryValueContraceptivePatch HKCategoryValueContraceptive = 0
	// HKCategoryValueContraceptiveUnspecified - An unspecified type of contraceptive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/unspecified
	HKCategoryValueContraceptiveUnspecified HKCategoryValueContraceptive = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueEnvironmentalAudioExposureEvent (1 cases) */
// HKCategoryValueEnvironmentalAudioExposureEvent - Exposure events for environmental audio.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueEnvironmentalAudioExposureEvent
type HKCategoryValueEnvironmentalAudioExposureEvent uint

const (
	// HKCategoryValueEnvironmentalAudioExposureEventMomentaryLimit - A brief exposure to a loud environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueEnvironmentalAudioExposureEvent/momentaryLimit
	HKCategoryValueEnvironmentalAudioExposureEventMomentaryLimit HKCategoryValueEnvironmentalAudioExposureEvent = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueHeadphoneAudioExposureEvent (1 cases) */
// HKCategoryValueHeadphoneAudioExposureEvent - Exposure events for headphone audio.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueHeadphoneAudioExposureEvent
type HKCategoryValueHeadphoneAudioExposureEvent uint

const (
	// HKCategoryValueHeadphoneAudioExposureEventSevenDayLimit - Exposure to significant audio levels from headphones over a seven-day period.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueHeadphoneAudioExposureEvent/sevenDayLimit
	HKCategoryValueHeadphoneAudioExposureEventSevenDayLimit HKCategoryValueHeadphoneAudioExposureEvent = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueLowCardioFitnessEvent (1 cases) */
// HKCategoryValueLowCardioFitnessEvent - A value that indicates a low-level cardio fitness event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueLowCardioFitnessEvent
type HKCategoryValueLowCardioFitnessEvent uint

const (
	// HKCategoryValueLowCardioFitnessEventLowFitness - An event that indicates a low level of cardio fitness, based on the VO2 max readings from the user’s Apple Watch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueLowCardioFitnessEvent/lowFitness
	HKCategoryValueLowCardioFitnessEventLowFitness HKCategoryValueLowCardioFitnessEvent = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueMenstrualFlow (5 cases) */
// HKCategoryValueMenstrualFlow - Categories that indicate the amount of menstrual flow for a given sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueMenstrualFlow
type HKCategoryValueMenstrualFlow uint

const (
	// HKCategoryValueMenstrualFlowHeavy - Heavy menstrual flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueMenstrualFlow/heavy
	HKCategoryValueMenstrualFlowHeavy HKCategoryValueMenstrualFlow = 0
	// HKCategoryValueMenstrualFlowLight - Light menstrual flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueMenstrualFlow/light
	HKCategoryValueMenstrualFlowLight HKCategoryValueMenstrualFlow = 0
	// HKCategoryValueMenstrualFlowMedium - Medium menstrual flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueMenstrualFlow/medium
	HKCategoryValueMenstrualFlowMedium HKCategoryValueMenstrualFlow = 0
	// HKCategoryValueMenstrualFlowNone - No menstrual flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueMenstrualFlow/none
	HKCategoryValueMenstrualFlowNone HKCategoryValueMenstrualFlow = 0
	// HKCategoryValueMenstrualFlowUnspecified - An unspecified amount of menstrual flow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueMenstrualFlow/unspecified
	HKCategoryValueMenstrualFlowUnspecified HKCategoryValueMenstrualFlow = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueOvulationTestResult (5 cases) */
// HKCategoryValueOvulationTestResult - Categories that represent the result of an ovulation home test.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult
type HKCategoryValueOvulationTestResult uint

const (
	// HKCategoryValueOvulationTestResultEstrogenSurge - The ovulation test detected a surge in estrogen. This value often refers to a   result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult/estrogenSurge
	HKCategoryValueOvulationTestResultEstrogenSurge HKCategoryValueOvulationTestResult = 0
	// HKCategoryValueOvulationTestResultIndeterminate - The ovulation test is inconclusive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult/indeterminate
	HKCategoryValueOvulationTestResultIndeterminate HKCategoryValueOvulationTestResult = 0
	// HKCategoryValueOvulationTestResultLuteinizingHormoneSurge - The ovulation test detected a surge in the luteinizing hormone. This value often refers to a   or   result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult/luteinizingHormoneSurge
	HKCategoryValueOvulationTestResultLuteinizingHormoneSurge HKCategoryValueOvulationTestResult = 0
	// HKCategoryValueOvulationTestResultNegative - The ovulation test is negative.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult/negative
	HKCategoryValueOvulationTestResultNegative HKCategoryValueOvulationTestResult = 0
	// HKCategoryValueOvulationTestResultPositive - The ovulation test is positive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult/positive
	HKCategoryValueOvulationTestResultPositive HKCategoryValueOvulationTestResult = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValuePregnancyTestResult (3 cases) */
// HKCategoryValuePregnancyTestResult - Category values that indicate the results of a home pregnancy test.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePregnancyTestResult
type HKCategoryValuePregnancyTestResult uint

const (
	// HKCategoryValuePregnancyTestResultIndeterminate - The test was inconclusive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePregnancyTestResult/indeterminate
	HKCategoryValuePregnancyTestResultIndeterminate HKCategoryValuePregnancyTestResult = 0
	// HKCategoryValuePregnancyTestResultNegative - The test returned a negative result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePregnancyTestResult/negative
	HKCategoryValuePregnancyTestResultNegative HKCategoryValuePregnancyTestResult = 0
	// HKCategoryValuePregnancyTestResultPositive - The test returned a positive result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePregnancyTestResult/positive
	HKCategoryValuePregnancyTestResultPositive HKCategoryValuePregnancyTestResult = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValuePresence (2 cases) */
// HKCategoryValuePresence - Categories that indicate whether a symptom is present.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePresence
type HKCategoryValuePresence uint

const (
	// HKCategoryValuePresenceNotPresent - The symptom isn’t present.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePresence/notPresent
	HKCategoryValuePresenceNotPresent HKCategoryValuePresence = 0
	// HKCategoryValuePresencePresent - The symptom is present.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValuePresence/present
	HKCategoryValuePresencePresent HKCategoryValuePresence = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueProgesteroneTestResult (3 cases) */
// HKCategoryValueProgesteroneTestResult - A category value that indicates the result from a home progesterone test.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueProgesteroneTestResult
type HKCategoryValueProgesteroneTestResult uint

const (
	// HKCategoryValueProgesteroneTestResultIndeterminate - The test was inconclusive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueProgesteroneTestResult/indeterminate
	HKCategoryValueProgesteroneTestResultIndeterminate HKCategoryValueProgesteroneTestResult = 0
	// HKCategoryValueProgesteroneTestResultNegative - The test returned a negative result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueProgesteroneTestResult/negative
	HKCategoryValueProgesteroneTestResultNegative HKCategoryValueProgesteroneTestResult = 0
	// HKCategoryValueProgesteroneTestResultPositive - The test returned a positive result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueProgesteroneTestResult/positive
	HKCategoryValueProgesteroneTestResultPositive HKCategoryValueProgesteroneTestResult = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueSeverity (5 cases) */
// HKCategoryValueSeverity - Categories that represent the severity of a symptom.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity
type HKCategoryValueSeverity uint

const (
	// HKCategoryValueSeverityMild - The symptom is mild.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/mild
	HKCategoryValueSeverityMild HKCategoryValueSeverity = 0
	// HKCategoryValueSeverityModerate - The symptom is moderate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/moderate
	HKCategoryValueSeverityModerate HKCategoryValueSeverity = 0
	// HKCategoryValueSeverityNotPresent - The symptom is not present.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/notPresent
	HKCategoryValueSeverityNotPresent HKCategoryValueSeverity = 0
	// HKCategoryValueSeveritySevere - The symptom is severe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/severe
	HKCategoryValueSeveritySevere HKCategoryValueSeverity = 0
	// HKCategoryValueSeverityUnspecified - The symptom’s severity is not specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/unspecified
	HKCategoryValueSeverityUnspecified HKCategoryValueSeverity = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueSleepAnalysis (7 cases) */
// HKCategoryValueSleepAnalysis - Categories that represent the result of a sleep analysis.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis
type HKCategoryValueSleepAnalysis uint

const (
	// HKCategoryValueSleepAnalysisAsleep - The user is sleeping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/asleep
	HKCategoryValueSleepAnalysisAsleep HKCategoryValueSleepAnalysis = 0
	// HKCategoryValueSleepAnalysisAsleepCore - The user is in light or intermediate sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/asleepCore
	HKCategoryValueSleepAnalysisAsleepCore HKCategoryValueSleepAnalysis = 0
	// HKCategoryValueSleepAnalysisAsleepDeep - The user is in deep sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/asleepDeep
	HKCategoryValueSleepAnalysisAsleepDeep HKCategoryValueSleepAnalysis = 0
	// HKCategoryValueSleepAnalysisAsleepREM - The user is in REM sleep.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/asleepREM
	HKCategoryValueSleepAnalysisAsleepREM HKCategoryValueSleepAnalysis = 0
	// HKCategoryValueSleepAnalysisAsleepUnspecified - The user is asleep, but the specific stage isn’t known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/asleepUnspecified
	HKCategoryValueSleepAnalysisAsleepUnspecified HKCategoryValueSleepAnalysis = 0
	// HKCategoryValueSleepAnalysisAwake - The user is awake.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/awake
	HKCategoryValueSleepAnalysisAwake HKCategoryValueSleepAnalysis = 0
	// HKCategoryValueSleepAnalysisInBed - The user is in bed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/inBed
	HKCategoryValueSleepAnalysisInBed HKCategoryValueSleepAnalysis = 0
)

/* debug [enums.gen.go]: Processing enum HKCategoryValueVaginalBleeding (5 cases) */
// HKCategoryValueVaginalBleeding enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding
type HKCategoryValueVaginalBleeding uint

const (
	// HKCategoryValueVaginalBleedingHeavy - Heavy vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/heavy
	HKCategoryValueVaginalBleedingHeavy HKCategoryValueVaginalBleeding = 0
	// HKCategoryValueVaginalBleedingLight - Light vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/light
	HKCategoryValueVaginalBleedingLight HKCategoryValueVaginalBleeding = 0
	// HKCategoryValueVaginalBleedingMedium - Medium vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/medium
	HKCategoryValueVaginalBleedingMedium HKCategoryValueVaginalBleeding = 0
	// HKCategoryValueVaginalBleedingNone - No vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/none
	HKCategoryValueVaginalBleedingNone HKCategoryValueVaginalBleeding = 0
	// HKCategoryValueVaginalBleedingUnspecified - An unspecified amount of vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/unspecified
	HKCategoryValueVaginalBleedingUnspecified HKCategoryValueVaginalBleeding = 0
)

/* debug [enums.gen.go]: Processing enum HKCyclingFunctionalThresholdPowerTestType (4 cases) */
// HKCyclingFunctionalThresholdPowerTestType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCyclingFunctionalThresholdPowerTestType
type HKCyclingFunctionalThresholdPowerTestType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCyclingFunctionalThresholdPowerTestType/maxExercise20Minute
	HKCyclingFunctionalThresholdPowerTestTypeMaxExercise20Minute HKCyclingFunctionalThresholdPowerTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCyclingFunctionalThresholdPowerTestType/maxExercise60Minute
	HKCyclingFunctionalThresholdPowerTestTypeMaxExercise60Minute HKCyclingFunctionalThresholdPowerTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCyclingFunctionalThresholdPowerTestType/predictionExercise
	HKCyclingFunctionalThresholdPowerTestTypePredictionExercise HKCyclingFunctionalThresholdPowerTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCyclingFunctionalThresholdPowerTestType/rampTest
	HKCyclingFunctionalThresholdPowerTestTypeRampTest HKCyclingFunctionalThresholdPowerTestType = 0
)

/* debug [enums.gen.go]: Processing enum HKDevicePlacementSide (4 cases) */
// HKDevicePlacementSide - Values that indicate the placement of the device that measured a sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevicePlacementSide
type HKDevicePlacementSide uint

const (
	// HKDevicePlacementSideCentral - A device predominately located near the center of the body.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevicePlacementSide/central
	HKDevicePlacementSideCentral HKDevicePlacementSide = 0
	// HKDevicePlacementSideLeft - A device predominately located on the left side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevicePlacementSide/left
	HKDevicePlacementSideLeft HKDevicePlacementSide = 0
	// HKDevicePlacementSideRight - A device predominately located on the right side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevicePlacementSide/right
	HKDevicePlacementSideRight HKDevicePlacementSide = 0
	// HKDevicePlacementSideUnknown - The system couldn’t determine the device’s placement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDevicePlacementSide/unknown
	HKDevicePlacementSideUnknown HKDevicePlacementSide = 0
)

/* debug [enums.gen.go]: Processing enum HKElectrocardiogramClassification (8 cases) */
// HKElectrocardiogramClassification - Classifications returned by Apple Watch’s ECG algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum
type HKElectrocardiogramClassification uint

const (
	// HKElectrocardiogramClassificationAtrialFibrillation - The sample exhibits signs of atrial fibrillation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/atrialFibrillation
	HKElectrocardiogramClassificationAtrialFibrillation HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationInconclusiveHighHeartRate - An unclassifiable sample caused by a rapid heart rate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/inconclusiveHighHeartRate
	HKElectrocardiogramClassificationInconclusiveHighHeartRate HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationInconclusiveLowHeartRate - An unclassifiable sample caused by a heart rate below 50 bpm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/inconclusiveLowHeartRate
	HKElectrocardiogramClassificationInconclusiveLowHeartRate HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationInconclusiveOther - An unclassifiable sample caused by an unknown issue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/inconclusiveOther
	HKElectrocardiogramClassificationInconclusiveOther HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationInconclusivePoorReading - An unclassifiable sample caused by an unclear signal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/inconclusivePoorReading
	HKElectrocardiogramClassificationInconclusivePoorReading HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationNotSet - A sample that doesn’t have an assigned classification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/notSet
	HKElectrocardiogramClassificationNotSet HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationSinusRhythm - The sample exhibits no signs of atrial fibrillation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/sinusRhythm
	HKElectrocardiogramClassificationSinusRhythm HKElectrocardiogramClassification = 0
	// HKElectrocardiogramClassificationUnrecognized - A sample classification that this version of HealthKit doesn’t recognize.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Classification-swift.enum/unrecognized
	HKElectrocardiogramClassificationUnrecognized HKElectrocardiogramClassification = 0
)

/* debug [enums.gen.go]: Processing enum HKElectrocardiogramLead (1 cases) */
// HKElectrocardiogramLead - The lead used to record a voltage measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Lead
type HKElectrocardiogramLead uint

const (
	// HKElectrocardiogramLeadAppleWatchSimilarToLeadI - Apple Watch Series 4 or later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/Lead/appleWatchSimilarToLeadI
	HKElectrocardiogramLeadAppleWatchSimilarToLeadI HKElectrocardiogramLead = 0
)

/* debug [enums.gen.go]: Processing enum HKElectrocardiogramSymptomsStatus (3 cases) */
// HKElectrocardiogramSymptomsStatus - Values indicating whether the user entered a symptom when they recorded the ECG.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/SymptomsStatus-swift.enum
type HKElectrocardiogramSymptomsStatus uint

const (
	// HKElectrocardiogramSymptomsStatusNone - The user didn’t experience any symptoms during the duration of the electrocardiogram reading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/SymptomsStatus-swift.enum/none
	HKElectrocardiogramSymptomsStatusNone HKElectrocardiogramSymptomsStatus = 0
	// HKElectrocardiogramSymptomsStatusNotSet - The user didn’t specify whether or not they experienced symptoms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/SymptomsStatus-swift.enum/notSet
	HKElectrocardiogramSymptomsStatusNotSet HKElectrocardiogramSymptomsStatus = 0
	// HKElectrocardiogramSymptomsStatusPresent - The user added a symptom when they recorded the ECG.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/SymptomsStatus-swift.enum/present
	HKElectrocardiogramSymptomsStatusPresent HKElectrocardiogramSymptomsStatus = 0
)

/* debug [enums.gen.go]: Processing enum HKFitzpatrickSkinType (7 cases) */
// HKFitzpatrickSkinType - Categories representing the user’s skin type based on the Fitzpatrick scale.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType
type HKFitzpatrickSkinType uint

const (
	// HKFitzpatrickSkinTypeI - Pale white skin that always burns easily in the sun and never tans.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/I
	HKFitzpatrickSkinTypeI HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeII - White skin that burns easily and tans minimally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/II
	HKFitzpatrickSkinTypeII HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeIII - White to light brown skin that burns moderately and tans uniformly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/III
	HKFitzpatrickSkinTypeIII HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeIV - Beige-olive, lightly tanned skin that burns minimally and tans moderately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/IV
	HKFitzpatrickSkinTypeIV HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeNotSet - Either the user’s skin type is not set, or the user has not granted your app permission to read the skin type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/notSet
	HKFitzpatrickSkinTypeNotSet HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeV - Brown skin that rarely burns and tans profusely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/V
	HKFitzpatrickSkinTypeV HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeVI - Dark brown to black skin that never burns and tans profusely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/VI
	HKFitzpatrickSkinTypeVI HKFitzpatrickSkinType = 0
)

/* debug [enums.gen.go]: Processing enum HKGAD7AssessmentAnswer (4 cases) */
// HKGAD7AssessmentAnswer enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer
type HKGAD7AssessmentAnswer uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer/moreThanHalfTheDays
	HKGAD7AssessmentAnswerMoreThanHalfTheDays HKGAD7AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer/nearlyEveryDay
	HKGAD7AssessmentAnswerNearlyEveryDay HKGAD7AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer/notAtAll
	HKGAD7AssessmentAnswerNotAtAll HKGAD7AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer/severalDays
	HKGAD7AssessmentAnswerSeveralDays HKGAD7AssessmentAnswer = 0
)

/* debug [enums.gen.go]: Processing enum HKGAD7AssessmentRisk (4 cases) */
// HKGAD7AssessmentRisk enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum
type HKGAD7AssessmentRisk uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/mild
	HKGAD7AssessmentRiskMild HKGAD7AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/moderate
	HKGAD7AssessmentRiskModerate HKGAD7AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/noneToMinimal
	HKGAD7AssessmentRiskNoneToMinimal HKGAD7AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/severe
	HKGAD7AssessmentRiskSevere HKGAD7AssessmentRisk = 0
)

/* debug [enums.gen.go]: Processing enum HKHeartRateMotionContext (3 cases) */
// HKHeartRateMotionContext - Values that indicate the user’s level of activity when the heart rate sample was measured.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateMotionContext
type HKHeartRateMotionContext uint

const (
	// HKHeartRateMotionContextActive - A value indicating that the user was in motion during the heart rate sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateMotionContext/active
	HKHeartRateMotionContextActive HKHeartRateMotionContext = 0
	// HKHeartRateMotionContextNotSet - A value indicating that the user’s activity level could not be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateMotionContext/notSet
	HKHeartRateMotionContextNotSet HKHeartRateMotionContext = 0
	// HKHeartRateMotionContextSedentary - A value indicating that the user has been still for at least 5 minutes prior to the heart rate sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateMotionContext/sedentary
	HKHeartRateMotionContextSedentary HKHeartRateMotionContext = 0
)

/* debug [enums.gen.go]: Processing enum HKHeartRateRecoveryTestType (3 cases) */
// HKHeartRateRecoveryTestType - The test that measured a person’s heart-rate recovery.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateRecoveryTestType
type HKHeartRateRecoveryTestType uint

const (
	// HKHeartRateRecoveryTestTypeMaxExercise - Measures a person’s actual heart-rate recovery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateRecoveryTestType/maxExercise
	HKHeartRateRecoveryTestTypeMaxExercise HKHeartRateRecoveryTestType = 0
	// HKHeartRateRecoveryTestTypePredictionNonExercise - A test that estimates a person’s heart-rate recovery without using exercise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateRecoveryTestType/predictionNonExercise
	HKHeartRateRecoveryTestTypePredictionNonExercise HKHeartRateRecoveryTestType = 0
	// HKHeartRateRecoveryTestTypePredictionSubMaxExercise - A test that estimates a person’s heart-rate recovery using lower-intensity exercise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateRecoveryTestType/predictionSubMaxExercise
	HKHeartRateRecoveryTestTypePredictionSubMaxExercise HKHeartRateRecoveryTestType = 0
)

/* debug [enums.gen.go]: Processing enum HKHeartRateSensorLocation (7 cases) */
// HKHeartRateSensorLocation - Constants that indicate where on the body the heart rate sensor is located.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation
type HKHeartRateSensorLocation uint

const (
	// HKHeartRateSensorLocationChest - The heart rate sensor is located on the user’s chest.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/chest
	HKHeartRateSensorLocationChest HKHeartRateSensorLocation = 0
	// HKHeartRateSensorLocationEarLobe - The heart rate sensor is located on the user’s earlobe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/earLobe
	HKHeartRateSensorLocationEarLobe HKHeartRateSensorLocation = 0
	// HKHeartRateSensorLocationFinger - The heart rate sensor is located on the user’s finger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/finger
	HKHeartRateSensorLocationFinger HKHeartRateSensorLocation = 0
	// HKHeartRateSensorLocationFoot - The heart rate sensor is located on the user’s foot.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/foot
	HKHeartRateSensorLocationFoot HKHeartRateSensorLocation = 0
	// HKHeartRateSensorLocationHand - The heart rate sensor is located on the user’s hand.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/hand
	HKHeartRateSensorLocationHand HKHeartRateSensorLocation = 0
	// HKHeartRateSensorLocationOther - The heart rate sensor’s location is not otherwise on this list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/other
	HKHeartRateSensorLocationOther HKHeartRateSensorLocation = 0
	// HKHeartRateSensorLocationWrist - The heart rate sensor is located on the user’s wrist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateSensorLocation/wrist
	HKHeartRateSensorLocationWrist HKHeartRateSensorLocation = 0
)

/* debug [enums.gen.go]: Processing enum HKInsulinDeliveryReason (2 cases) */
// HKInsulinDeliveryReason - Possible reasons for administering insulin.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKInsulinDeliveryReason
type HKInsulinDeliveryReason uint

const (
	// HKInsulinDeliveryReasonBasal - Insulin administered to meet the user’s basic metabolic needs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKInsulinDeliveryReason/basal
	HKInsulinDeliveryReasonBasal HKInsulinDeliveryReason = 0
	// HKInsulinDeliveryReasonBolus - Insulin administered to meet the user’s episodic requirements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKInsulinDeliveryReason/bolus
	HKInsulinDeliveryReasonBolus HKInsulinDeliveryReason = 0
)

/* debug [enums.gen.go]: Processing enum HKMetricPrefix (14 cases) */
// HKMetricPrefix - Prefixes that can be added to SI units to change the order of magnitude.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix
type HKMetricPrefix uint

const (
	// HKMetricPrefixCenti - A prefix that multiplies the base unit by 0.01.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/centi
	HKMetricPrefixCenti HKMetricPrefix = 0
	// HKMetricPrefixDeca - A prefix that multiplies the base unit by 10.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/deca
	HKMetricPrefixDeca HKMetricPrefix = 0
	// HKMetricPrefixDeci - A prefix that multiplies the base unit by 0.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/deci
	HKMetricPrefixDeci HKMetricPrefix = 0
	// HKMetricPrefixFemto - A prefix that multiplies the base unit by 1e-15.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/femto
	HKMetricPrefixFemto HKMetricPrefix = 0
	// HKMetricPrefixGiga - A prefix that multiplies the base unit by 1e9.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/giga
	HKMetricPrefixGiga HKMetricPrefix = 0
	// HKMetricPrefixHecto - A prefix that multiplies the base unit by 100.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/hecto
	HKMetricPrefixHecto HKMetricPrefix = 0
	// HKMetricPrefixKilo - A prefix that multiplies the base unit by 1000.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/kilo
	HKMetricPrefixKilo HKMetricPrefix = 0
	// HKMetricPrefixMega - A prefix that multiplies the base unit by 1e6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/mega
	HKMetricPrefixMega HKMetricPrefix = 0
	// HKMetricPrefixMicro - A prefix that multiplies the base unit by 1e-6.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/micro
	HKMetricPrefixMicro HKMetricPrefix = 0
	// HKMetricPrefixMilli - A prefix that multiplies the base unit by 0.001.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/milli
	HKMetricPrefixMilli HKMetricPrefix = 0
	// HKMetricPrefixNano - A prefix that multiplies the base unit by 1e-9.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/nano
	HKMetricPrefixNano HKMetricPrefix = 0
	// HKMetricPrefixNone - A prefix that does not modify the base unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/none
	HKMetricPrefixNone HKMetricPrefix = 0
	// HKMetricPrefixPico - A prefix that multiplies the base unit by 1e-12.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/pico
	HKMetricPrefixPico HKMetricPrefix = 0
	// HKMetricPrefixTera - A prefix that multiplies the base unit by 1e12.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMetricPrefix/tera
	HKMetricPrefixTera HKMetricPrefix = 0
)

/* debug [enums.gen.go]: Processing enum HKPHQ9AssessmentAnswer (5 cases) */
// HKPHQ9AssessmentAnswer enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer
type HKPHQ9AssessmentAnswer uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer/moreThanHalfTheDays
	HKPHQ9AssessmentAnswerMoreThanHalfTheDays HKPHQ9AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer/nearlyEveryDay
	HKPHQ9AssessmentAnswerNearlyEveryDay HKPHQ9AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer/notAtAll
	HKPHQ9AssessmentAnswerNotAtAll HKPHQ9AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer/preferNotToAnswer
	HKPHQ9AssessmentAnswerPreferNotToAnswer HKPHQ9AssessmentAnswer = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer/severalDays
	HKPHQ9AssessmentAnswerSeveralDays HKPHQ9AssessmentAnswer = 0
)

/* debug [enums.gen.go]: Processing enum HKPHQ9AssessmentRisk (5 cases) */
// HKPHQ9AssessmentRisk enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum
type HKPHQ9AssessmentRisk uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum/mild
	HKPHQ9AssessmentRiskMild HKPHQ9AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum/moderate
	HKPHQ9AssessmentRiskModerate HKPHQ9AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum/moderatelySevere
	HKPHQ9AssessmentRiskModeratelySevere HKPHQ9AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum/noneToMinimal
	HKPHQ9AssessmentRiskNoneToMinimal HKPHQ9AssessmentRisk = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum/severe
	HKPHQ9AssessmentRiskSevere HKPHQ9AssessmentRisk = 0
)

/* debug [enums.gen.go]: Processing enum HKPhysicalEffortEstimationType (2 cases) */
// HKPhysicalEffortEstimationType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPhysicalEffortEstimationType
type HKPhysicalEffortEstimationType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPhysicalEffortEstimationType/activityLookup
	HKPhysicalEffortEstimationTypeActivityLookup HKPhysicalEffortEstimationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPhysicalEffortEstimationType/deviceSensed
	HKPhysicalEffortEstimationTypeDeviceSensed HKPhysicalEffortEstimationType = 0
)

/* debug [enums.gen.go]: Processing enum HKPrismBase (5 cases) */
// HKPrismBase - The orientation of the prism correction, represented by the location of the prism’s base (the thickest part of the prism).
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrismBase
type HKPrismBase uint

const (
	// HKPrismBaseDown - The prism’s base is at the bottom of the lens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrismBase/down
	HKPrismBaseDown HKPrismBase = 0
	// HKPrismBaseIn - The prism base is on the inside edge of the lens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrismBase/in
	HKPrismBaseIn HKPrismBase = 0
	// HKPrismBaseNone - No prism correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrismBase/none
	HKPrismBaseNone HKPrismBase = 0
	// HKPrismBaseOut - The prism base is on the outside edge of the lens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrismBase/out
	HKPrismBaseOut HKPrismBase = 0
	// HKPrismBaseUp - The prism’s base is at the top of the lens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrismBase/up
	HKPrismBaseUp HKPrismBase = 0
)

/* debug [enums.gen.go]: Processing enum HKQuantityAggregationStyle (5 cases) */
// HKQuantityAggregationStyle - Constant values that describe how quantities can be aggregated over time.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityAggregationStyle
type HKQuantityAggregationStyle uint

const (
	// HKQuantityAggregationStyleCumulative - Cumulative samples that can be summed over time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityAggregationStyle/cumulative
	HKQuantityAggregationStyleCumulative HKQuantityAggregationStyle = 0
	// HKQuantityAggregationStyleDiscrete - Discrete samples may be averaged over time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityAggregationStyle/discrete
	HKQuantityAggregationStyleDiscrete HKQuantityAggregationStyle = 0
	// HKQuantityAggregationStyleDiscreteArithmetic - Discrete samples that can be averaged over time using an arithmetic mean.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityAggregationStyle/discreteArithmetic
	HKQuantityAggregationStyleDiscreteArithmetic HKQuantityAggregationStyle = 0
	// HKQuantityAggregationStyleDiscreteEquivalentContinuousLevel - Discrete samples that can be combined over a time interval by computing the equivalent continuous sound level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityAggregationStyle/discreteEquivalentContinuousLevel
	HKQuantityAggregationStyleDiscreteEquivalentContinuousLevel HKQuantityAggregationStyle = 0
	// HKQuantityAggregationStyleDiscreteTemporallyWeighted - Discrete samples that can be averaged over a time interval using a temporally weighted integration function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityAggregationStyle/discreteTemporallyWeighted
	HKQuantityAggregationStyleDiscreteTemporallyWeighted HKQuantityAggregationStyle = 0
)

/* debug [enums.gen.go]: Processing enum HKQueryOptions (3 cases) */
// HKQueryOptions - Constants that describe how a sample’s time period overlaps with the target time period.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryOptions
type HKQueryOptions uint

const (
	// HKQueryOptionNone - The sample’s time period must overlap part of the target time period.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryOptions/HKQueryOptionNone
	HKQueryOptionNone HKQueryOptions = 0
	// HKQueryOptionStrictEndDate - The sample’s end time must fall within the target time period.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryOptions/strictEndDate
	HKQueryOptionStrictEndDate HKQueryOptions = 0
	// HKQueryOptionStrictStartDate - The sample’s start time must fall within the target time period.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryOptions/strictStartDate
	HKQueryOptionStrictStartDate HKQueryOptions = 0
)

/* debug [enums.gen.go]: Processing enum HKStateOfMindAssociation (18 cases) */
// HKStateOfMindAssociation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association
type HKStateOfMindAssociation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/community
	HKStateOfMindAssociationCommunity HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/currentEvents
	HKStateOfMindAssociationCurrentEvents HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/dating
	HKStateOfMindAssociationDating HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/education
	HKStateOfMindAssociationEducation HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/family
	HKStateOfMindAssociationFamily HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/fitness
	HKStateOfMindAssociationFitness HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/friends
	HKStateOfMindAssociationFriends HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/health
	HKStateOfMindAssociationHealth HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/hobbies
	HKStateOfMindAssociationHobbies HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/identity
	HKStateOfMindAssociationIdentity HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/money
	HKStateOfMindAssociationMoney HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/partner
	HKStateOfMindAssociationPartner HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/selfCare
	HKStateOfMindAssociationSelfCare HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/spirituality
	HKStateOfMindAssociationSpirituality HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/tasks
	HKStateOfMindAssociationTasks HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/travel
	HKStateOfMindAssociationTravel HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/weather
	HKStateOfMindAssociationWeather HKStateOfMindAssociation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/work
	HKStateOfMindAssociationWork HKStateOfMindAssociation = 0
)

/* debug [enums.gen.go]: Processing enum HKStateOfMindKind (2 cases) */
// HKStateOfMindKind enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Kind-swift.enum
type HKStateOfMindKind uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Kind-swift.enum/dailyMood
	HKStateOfMindKindDailyMood HKStateOfMindKind = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Kind-swift.enum/momentaryEmotion
	HKStateOfMindKindMomentaryEmotion HKStateOfMindKind = 0
)

/* debug [enums.gen.go]: Processing enum HKStateOfMindLabel (38 cases) */
// HKStateOfMindLabel enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label
type HKStateOfMindLabel uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/amazed
	HKStateOfMindLabelAmazed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/amused
	HKStateOfMindLabelAmused HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/angry
	HKStateOfMindLabelAngry HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/annoyed
	HKStateOfMindLabelAnnoyed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/anxious
	HKStateOfMindLabelAnxious HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/ashamed
	HKStateOfMindLabelAshamed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/brave
	HKStateOfMindLabelBrave HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/calm
	HKStateOfMindLabelCalm HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/confident
	HKStateOfMindLabelConfident HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/content
	HKStateOfMindLabelContent HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/disappointed
	HKStateOfMindLabelDisappointed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/discouraged
	HKStateOfMindLabelDiscouraged HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/disgusted
	HKStateOfMindLabelDisgusted HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/drained
	HKStateOfMindLabelDrained HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/embarrassed
	HKStateOfMindLabelEmbarrassed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/excited
	HKStateOfMindLabelExcited HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/frustrated
	HKStateOfMindLabelFrustrated HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/grateful
	HKStateOfMindLabelGrateful HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/guilty
	HKStateOfMindLabelGuilty HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/happy
	HKStateOfMindLabelHappy HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/hopeful
	HKStateOfMindLabelHopeful HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/hopeless
	HKStateOfMindLabelHopeless HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/indifferent
	HKStateOfMindLabelIndifferent HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/irritated
	HKStateOfMindLabelIrritated HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/jealous
	HKStateOfMindLabelJealous HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/joyful
	HKStateOfMindLabelJoyful HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/lonely
	HKStateOfMindLabelLonely HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/overwhelmed
	HKStateOfMindLabelOverwhelmed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/passionate
	HKStateOfMindLabelPassionate HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/peaceful
	HKStateOfMindLabelPeaceful HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/proud
	HKStateOfMindLabelProud HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/relieved
	HKStateOfMindLabelRelieved HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/sad
	HKStateOfMindLabelSad HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/satisfied
	HKStateOfMindLabelSatisfied HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/scared
	HKStateOfMindLabelScared HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/stressed
	HKStateOfMindLabelStressed HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/surprised
	HKStateOfMindLabelSurprised HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/worried
	HKStateOfMindLabelWorried HKStateOfMindLabel = 0
)

/* debug [enums.gen.go]: Processing enum HKStateOfMindValenceClassification (7 cases) */
// HKStateOfMindValenceClassification enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum
type HKStateOfMindValenceClassification uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/neutral
	HKStateOfMindValenceClassificationNeutral HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/pleasant
	HKStateOfMindValenceClassificationPleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/slightlyPleasant
	HKStateOfMindValenceClassificationSlightlyPleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/slightlyUnpleasant
	HKStateOfMindValenceClassificationSlightlyUnpleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/unpleasant
	HKStateOfMindValenceClassificationUnpleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/veryPleasant
	HKStateOfMindValenceClassificationVeryPleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/veryUnpleasant
	HKStateOfMindValenceClassificationVeryUnpleasant HKStateOfMindValenceClassification = 0
)

/* debug [enums.gen.go]: Processing enum HKStatisticsOptions (9 cases) */
// HKStatisticsOptions - Options for specifying the statistic to calculate.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions
type HKStatisticsOptions uint

const (
	// HKStatisticsOptionCumulativeSum - An option indicating that the system calculates the sum of all the quantities for the samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/cumulativeSum
	HKStatisticsOptionCumulativeSum HKStatisticsOptions = 0
	// HKStatisticsOptionDiscreteAverage - An option indicating that the system calculates the average quantity for the samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/discreteAverage
	HKStatisticsOptionDiscreteAverage HKStatisticsOptions = 0
	// HKStatisticsOptionDiscreteMax - An option indicating that the system calculates the maximum quantity for the samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/discreteMax
	HKStatisticsOptionDiscreteMax HKStatisticsOptions = 0
	// HKStatisticsOptionDiscreteMin - An option indicating that the system calculates the minimum quantity for the samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/discreteMin
	HKStatisticsOptionDiscreteMin HKStatisticsOptions = 0
	// HKStatisticsOptionDiscreteMostRecent - An option indicating that the system returns the most recent quantity from the matching samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/discreteMostRecent
	HKStatisticsOptionDiscreteMostRecent HKStatisticsOptions = 0
	// HKStatisticsOptionDuration - An option indicating that the system calculates the total duration covering all the samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/duration
	HKStatisticsOptionDuration HKStatisticsOptions = 0
	// HKStatisticsOptionNone - An option indicating that the system will not calculate any statistics values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/HKStatisticsOptionNone
	HKStatisticsOptionNone HKStatisticsOptions = 0
	// HKStatisticsOptionMostRecent - An option indicating that the system returns the most recent quantity from the matching samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/mostRecent
	HKStatisticsOptionMostRecent HKStatisticsOptions = 0
	// HKStatisticsOptionSeparateBySource - An option indicating that the system calculates the specified statistics separately for each source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions/separateBySource
	HKStatisticsOptionSeparateBySource HKStatisticsOptions = 0
)

/* debug [enums.gen.go]: Processing enum HKSwimmingStrokeStyle (7 cases) */
// HKSwimmingStrokeStyle - The style of stroke while swimming.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle
type HKSwimmingStrokeStyle uint

const (
	// HKSwimmingStrokeStyleBackstroke - The user swam the backstroke.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/backstroke
	HKSwimmingStrokeStyleBackstroke HKSwimmingStrokeStyle = 0
	// HKSwimmingStrokeStyleBreaststroke - The user swam the breaststroke.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/breaststroke
	HKSwimmingStrokeStyleBreaststroke HKSwimmingStrokeStyle = 0
	// HKSwimmingStrokeStyleButterfly - The user swam the butterfly stroke.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/butterfly
	HKSwimmingStrokeStyleButterfly HKSwimmingStrokeStyle = 0
	// HKSwimmingStrokeStyleFreestyle - The user swam the freestyle stroke.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/freestyle
	HKSwimmingStrokeStyleFreestyle HKSwimmingStrokeStyle = 0
	// HKSwimmingStrokeStyleKickboard - The user swam using a kickboard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/kickboard
	HKSwimmingStrokeStyleKickboard HKSwimmingStrokeStyle = 0
	// HKSwimmingStrokeStyleMixed - The user swam a mixture of strokes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/mixed
	HKSwimmingStrokeStyleMixed HKSwimmingStrokeStyle = 0
	// HKSwimmingStrokeStyleUnknown - The user’s stroke could not be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSwimmingStrokeStyle/unknown
	HKSwimmingStrokeStyleUnknown HKSwimmingStrokeStyle = 0
)

/* debug [enums.gen.go]: Processing enum HKUpdateFrequency (4 cases) */
// HKUpdateFrequency - Constants that determine how often the system launches your app in response to changes to HealthKit data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency
type HKUpdateFrequency uint

const (
	// HKUpdateFrequencyDaily - The system launches your app at most once a day in response to changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/daily
	HKUpdateFrequencyDaily HKUpdateFrequency = 0
	// HKUpdateFrequencyHourly - The system launches your app at most once an hour in response to changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/hourly
	HKUpdateFrequencyHourly HKUpdateFrequency = 0
	// HKUpdateFrequencyImmediate - The system launches your app every time it detects a change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/immediate
	HKUpdateFrequencyImmediate HKUpdateFrequency = 0
	// HKUpdateFrequencyWeekly - The system launches your app at most once per week in response to changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/weekly
	HKUpdateFrequencyWeekly HKUpdateFrequency = 0
)

/* debug [enums.gen.go]: Processing enum HKUserMotionContext (3 cases) */
// HKUserMotionContext - The type of motion performed during the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserMotionContext
type HKUserMotionContext uint

const (
	// HKUserMotionContextActive - The person was active during the sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserMotionContext/active
	HKUserMotionContextActive HKUserMotionContext = 0
	// HKUserMotionContextNotSet - The person’s motion was not specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserMotionContext/notSet
	HKUserMotionContextNotSet HKUserMotionContext = 0
	// HKUserMotionContextStationary - The person was stationary during the sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserMotionContext/stationary
	HKUserMotionContextStationary HKUserMotionContext = 0
)

/* debug [enums.gen.go]: Processing enum HKVisionEye (2 cases) */
// HKVisionEye - A value that specifies the eye for a vision prescription.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionEye
type HKVisionEye uint

const (
	// HKVisionEyeLeft - The left eye.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionEye/left
	HKVisionEyeLeft HKVisionEye = 0
	// HKVisionEyeRight - The right eye.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionEye/right
	HKVisionEyeRight HKVisionEye = 0
)

/* debug [enums.gen.go]: Processing enum HKVisionPrescriptionType (2 cases) */
// HKVisionPrescriptionType - The type of vision prescription, for example a prescription for glasses or for contacts.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescriptionType
type HKVisionPrescriptionType uint

const (
	// HKVisionPrescriptionTypeContacts - A prescription for contacts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescriptionType/contacts
	HKVisionPrescriptionTypeContacts HKVisionPrescriptionType = 0
	// HKVisionPrescriptionTypeGlasses - A prescription for glasses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescriptionType/glasses
	HKVisionPrescriptionTypeGlasses HKVisionPrescriptionType = 0
)

/* debug [enums.gen.go]: Processing enum HKVO2MaxTestType (4 cases) */
// HKVO2MaxTestType - Methods for calculating the user’s VO2 max rate.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVO2MaxTestType
type HKVO2MaxTestType uint

const (
	// HKVO2MaxTestTypeMaxExercise - A test that measures VO2 max rate by monitoring exercise to the user’s physical limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVO2MaxTestType/maxExercise
	HKVO2MaxTestTypeMaxExercise HKVO2MaxTestType = 0
	// HKVO2MaxTestTypePredictionNonExercise - A calculation that estimates VO2 max rate without any exercise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVO2MaxTestType/predictionNonExercise
	HKVO2MaxTestTypePredictionNonExercise HKVO2MaxTestType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVO2MaxTestType/predictionStepTest
	HKVO2MaxTestTypePredictionStepTest HKVO2MaxTestType = 0
	// HKVO2MaxTestTypePredictionSubMaxExercise - A calculation that estimates VO2 max rate based on low-intensity exercise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVO2MaxTestType/predictionSubMaxExercise
	HKVO2MaxTestTypePredictionSubMaxExercise HKVO2MaxTestType = 0
)

/* debug [enums.gen.go]: Processing enum HKWaterSalinity (2 cases) */
// HKWaterSalinity enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWaterSalinity
type HKWaterSalinity uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWaterSalinity/freshWater
	HKWaterSalinityFreshWater HKWaterSalinity = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWaterSalinity/saltWater
	HKWaterSalinitySaltWater HKWaterSalinity = 0
)

/* debug [enums.gen.go]: Processing enum HKWeatherCondition (28 cases) */
// HKWeatherCondition - Constants that indicate a type of weather.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition
type HKWeatherCondition uint

const (
	// HKWeatherConditionBlustery - The weather condition is blustery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/blustery
	HKWeatherConditionBlustery HKWeatherCondition = 0
	// HKWeatherConditionClear - The weather condition is clear.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/clear
	HKWeatherConditionClear HKWeatherCondition = 0
	// HKWeatherConditionCloudy - The weather condition is cloudy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/cloudy
	HKWeatherConditionCloudy HKWeatherCondition = 0
	// HKWeatherConditionDrizzle - The weather condition is drizzle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/drizzle
	HKWeatherConditionDrizzle HKWeatherCondition = 0
	// HKWeatherConditionDust - The weather condition is dust.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/dust
	HKWeatherConditionDust HKWeatherCondition = 0
	// HKWeatherConditionFair - The weather condition is fair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/fair
	HKWeatherConditionFair HKWeatherCondition = 0
	// HKWeatherConditionFoggy - The weather condition is foggy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/foggy
	HKWeatherConditionFoggy HKWeatherCondition = 0
	// HKWeatherConditionFreezingDrizzle - The weather condition is freezing drizzle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/freezingDrizzle
	HKWeatherConditionFreezingDrizzle HKWeatherCondition = 0
	// HKWeatherConditionFreezingRain - The weather condition is freezing rain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/freezingRain
	HKWeatherConditionFreezingRain HKWeatherCondition = 0
	// HKWeatherConditionHail - The weather condition is hail.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/hail
	HKWeatherConditionHail HKWeatherCondition = 0
	// HKWeatherConditionHaze - The weather condition is hazy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/haze
	HKWeatherConditionHaze HKWeatherCondition = 0
	// HKWeatherConditionHurricane - The weather condition is hurricane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/hurricane
	HKWeatherConditionHurricane HKWeatherCondition = 0
	// HKWeatherConditionMixedRainAndHail - The weather condition is mixed rain and hail.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/mixedRainAndHail
	HKWeatherConditionMixedRainAndHail HKWeatherCondition = 0
	// HKWeatherConditionMixedRainAndSleet - The weather condition is mixed rain and sleet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/mixedRainAndSleet
	HKWeatherConditionMixedRainAndSleet HKWeatherCondition = 0
	// HKWeatherConditionMixedRainAndSnow - The weather condition is mixed rain and snow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/mixedRainAndSnow
	HKWeatherConditionMixedRainAndSnow HKWeatherCondition = 0
	// HKWeatherConditionMixedSnowAndSleet - The weather condition is mixed snow and sleet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/mixedSnowAndSleet
	HKWeatherConditionMixedSnowAndSleet HKWeatherCondition = 0
	// HKWeatherConditionMostlyCloudy - The weather condition is mostly cloudy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/mostlyCloudy
	HKWeatherConditionMostlyCloudy HKWeatherCondition = 0
	// HKWeatherConditionNone - The weather condition is unknown or irrelevant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/none
	HKWeatherConditionNone HKWeatherCondition = 0
	// HKWeatherConditionPartlyCloudy - The weather condition is partly cloudy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/partlyCloudy
	HKWeatherConditionPartlyCloudy HKWeatherCondition = 0
	// HKWeatherConditionScatteredShowers - The weather condition is scattered showers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/scatteredShowers
	HKWeatherConditionScatteredShowers HKWeatherCondition = 0
	// HKWeatherConditionShowers - The weather condition is showers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/showers
	HKWeatherConditionShowers HKWeatherCondition = 0
	// HKWeatherConditionSleet - The weather condition is sleet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/sleet
	HKWeatherConditionSleet HKWeatherCondition = 0
	// HKWeatherConditionSmoky - The weather condition is smoky.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/smoky
	HKWeatherConditionSmoky HKWeatherCondition = 0
	// HKWeatherConditionSnow - The weather condition is snow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/snow
	HKWeatherConditionSnow HKWeatherCondition = 0
	// HKWeatherConditionThunderstorms - The weather condition is thunderstorms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/thunderstorms
	HKWeatherConditionThunderstorms HKWeatherCondition = 0
	// HKWeatherConditionTornado - The weather condition is tornado.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/tornado
	HKWeatherConditionTornado HKWeatherCondition = 0
	// HKWeatherConditionTropicalStorm - The weather condition is tropical storm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/tropicalStorm
	HKWeatherConditionTropicalStorm HKWeatherCondition = 0
	// HKWeatherConditionWindy - The weather condition is windy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWeatherCondition/windy
	HKWeatherConditionWindy HKWeatherCondition = 0
)

/* debug [enums.gen.go]: Processing enum HKWheelchairUse (3 cases) */
// HKWheelchairUse - Constants indicating the user’s wheelchair use.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse
type HKWheelchairUse uint

const (
	// HKWheelchairUseNo - The user does not use a wheelchair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse/no
	HKWheelchairUseNo HKWheelchairUse = 0
	// HKWheelchairUseNotSet - Either the wheelchair use is not set or the user has not granted your app permission to read that information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse/notSet
	HKWheelchairUseNotSet HKWheelchairUse = 0
	// HKWheelchairUseYes - The user uses a wheelchair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse/yes
	HKWheelchairUseYes HKWheelchairUse = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutActivityType (84 cases) */
// HKWorkoutActivityType - The type of activity performed during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType
type HKWorkoutActivityType uint

const (
	// HKWorkoutActivityTypeAmericanFootball - The constant for playing American football.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/americanFootball
	HKWorkoutActivityTypeAmericanFootball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeArchery - The constant for shooting archery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/archery
	HKWorkoutActivityTypeArchery HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeAustralianFootball - The constant for playing Australian football.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/australianFootball
	HKWorkoutActivityTypeAustralianFootball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeBadminton - The constant for playing badminton.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/badminton
	HKWorkoutActivityTypeBadminton HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeBarre - The constant for barre workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/barre
	HKWorkoutActivityTypeBarre HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeBaseball - The constant for playing baseball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/baseball
	HKWorkoutActivityTypeBaseball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeBasketball - The constant for playing basketball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/basketball
	HKWorkoutActivityTypeBasketball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeBowling - The constant for bowling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/bowling
	HKWorkoutActivityTypeBowling HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeBoxing - The constant for boxing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/boxing
	HKWorkoutActivityTypeBoxing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCardioDance - The constant for cardiovascular dance workouts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/cardioDance
	HKWorkoutActivityTypeCardioDance HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeClimbing - The constant for climbing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/climbing
	HKWorkoutActivityTypeClimbing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCooldown - The constant for low intensity stretching and mobility exercises following a more vigorous workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/cooldown
	HKWorkoutActivityTypeCooldown HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCoreTraining - The constant for core training.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/coreTraining
	HKWorkoutActivityTypeCoreTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCricket - The constant for playing cricket.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/cricket
	HKWorkoutActivityTypeCricket HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCrossCountrySkiing - The constant for cross country skiing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/crossCountrySkiing
	HKWorkoutActivityTypeCrossCountrySkiing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCrossTraining - The constant for exercise that includes any mixture of cardio, strength, and/or flexibility training.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/crossTraining
	HKWorkoutActivityTypeCrossTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCurling - The constant for curling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/curling
	HKWorkoutActivityTypeCurling HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeCycling - The constant for cycling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/cycling
	HKWorkoutActivityTypeCycling HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeDance - The constant for dancing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/dance
	HKWorkoutActivityTypeDance HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeDanceInspiredTraining - The constant for workouts inspired by dance, including Pilates, Barre, and Feldenkrais.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/danceInspiredTraining
	HKWorkoutActivityTypeDanceInspiredTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeDiscSports - The constant for playing disc sports such as Ultimate and Disc Golf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/discSports
	HKWorkoutActivityTypeDiscSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeDownhillSkiing - The constant for downhill skiing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/downhillSkiing
	HKWorkoutActivityTypeDownhillSkiing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeElliptical - The constant for workouts on an elliptical machine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/elliptical
	HKWorkoutActivityTypeElliptical HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeEquestrianSports - The constant for activities that involve riding a horse, including polo, horse racing, and horse riding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/equestrianSports
	HKWorkoutActivityTypeEquestrianSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeFencing - The constant for fencing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/fencing
	HKWorkoutActivityTypeFencing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeFishing - The constant for fishing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/fishing
	HKWorkoutActivityTypeFishing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeFitnessGaming - The constant for playing fitness-based video games.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/fitnessGaming
	HKWorkoutActivityTypeFitnessGaming HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeFlexibility - The constant for a flexibility workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/flexibility
	HKWorkoutActivityTypeFlexibility HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeFunctionalStrengthTraining - The constant for strength training, primarily with free weights and body weight.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/functionalStrengthTraining
	HKWorkoutActivityTypeFunctionalStrengthTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeGolf - The constant for playing golf.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/golf
	HKWorkoutActivityTypeGolf HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeGymnastics - Performing gymnastics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/gymnastics
	HKWorkoutActivityTypeGymnastics HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHandball - The constant for playing handball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/handball
	HKWorkoutActivityTypeHandball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHandCycling - The constant for hand cycling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/handCycling
	HKWorkoutActivityTypeHandCycling HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHighIntensityIntervalTraining - The constant for high intensity interval training.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/highIntensityIntervalTraining
	HKWorkoutActivityTypeHighIntensityIntervalTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHiking - The constant for hiking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/hiking
	HKWorkoutActivityTypeHiking HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHockey - The constant for playing hockey, including ice hockey, field hockey, and related sports.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/hockey
	HKWorkoutActivityTypeHockey HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHunting - The constant for hunting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/hunting
	HKWorkoutActivityTypeHunting HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeJumpRope - The constant for jumping rope.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/jumpRope
	HKWorkoutActivityTypeJumpRope HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeKickboxing - The constant for kickboxing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/kickboxing
	HKWorkoutActivityTypeKickboxing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeLacrosse - The constant for playing lacrosse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/lacrosse
	HKWorkoutActivityTypeLacrosse HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeMartialArts - The constant for practicing martial arts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/martialArts
	HKWorkoutActivityTypeMartialArts HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeMindAndBody - The constant for performing activities like walking meditation, Gyrotonic exercise, and Qigong.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/mindAndBody
	HKWorkoutActivityTypeMindAndBody HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeMixedCardio - The constant for workouts that mix a variety of cardio exercise machines or modalities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/mixedCardio
	HKWorkoutActivityTypeMixedCardio HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeMixedMetabolicCardioTraining - The constant for performing any mix of cardio-focused exercises.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/mixedMetabolicCardioTraining
	HKWorkoutActivityTypeMixedMetabolicCardioTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeOther - The constant for a workout that does not match any of the other workout activity types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/other
	HKWorkoutActivityTypeOther HKWorkoutActivityType = 0
	// HKWorkoutActivityTypePaddleSports - The constant for canoeing, kayaking, paddling an outrigger, paddling a stand-up paddle board, and related sports.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/paddleSports
	HKWorkoutActivityTypePaddleSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypePickleball - The constant for playing pickleball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/pickleball
	HKWorkoutActivityTypePickleball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypePilates - The constant for a pilates workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/pilates
	HKWorkoutActivityTypePilates HKWorkoutActivityType = 0
	// HKWorkoutActivityTypePlay - The constant for play-based activities like tag, dodgeball, hopscotch, tetherball, and playing on a jungle gym.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/play
	HKWorkoutActivityTypePlay HKWorkoutActivityType = 0
	// HKWorkoutActivityTypePreparationAndRecovery - The constant for warm-up and therapeutic activities like foam rolling and stretching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/preparationAndRecovery
	HKWorkoutActivityTypePreparationAndRecovery HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeRacquetball - The constant for playing racquetball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/racquetball
	HKWorkoutActivityTypeRacquetball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeRowing - The constant for rowing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/rowing
	HKWorkoutActivityTypeRowing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeRugby - The constant for playing rugby.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/rugby
	HKWorkoutActivityTypeRugby HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeRunning - The constant for running and jogging.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/running
	HKWorkoutActivityTypeRunning HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSailing - The constant for sailing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/sailing
	HKWorkoutActivityTypeSailing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSkatingSports - The constant for skating activities, including ice skating, speed skating, inline skating, and skateboarding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/skatingSports
	HKWorkoutActivityTypeSkatingSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSnowboarding - The constant for snowboarding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/snowboarding
	HKWorkoutActivityTypeSnowboarding HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSnowSports - The constant for a variety of snow sports, including sledding, snowmobiling, or building a snowman.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/snowSports
	HKWorkoutActivityTypeSnowSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSoccer - The constant for playing soccer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/soccer
	HKWorkoutActivityTypeSoccer HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSocialDance - The constant for dancing with a partner or partners, such as swing, salsa, or folk dances.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/socialDance
	HKWorkoutActivityTypeSocialDance HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSoftball - The constant for playing softball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/softball
	HKWorkoutActivityTypeSoftball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSquash - The constant for playing squash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/squash
	HKWorkoutActivityTypeSquash HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeStairClimbing - The constant for workouts using a stair climbing machine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/stairClimbing
	HKWorkoutActivityTypeStairClimbing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeStairs - The constant for running, walking, or other drills using stairs (for example, in a stadium or inside a multilevel building).
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/stairs
	HKWorkoutActivityTypeStairs HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeStepTraining - The constant for training using a step bench.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/stepTraining
	HKWorkoutActivityTypeStepTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSurfingSports - The constant for a variety of surf sports, including surfing, kite surfing, and wind surfing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/surfingSports
	HKWorkoutActivityTypeSurfingSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSwimBikeRun - The constant for multisport activities like triathlons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/swimBikeRun
	HKWorkoutActivityTypeSwimBikeRun HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSwimming - The constant for swimming.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/swimming
	HKWorkoutActivityTypeSwimming HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTableTennis - The constant for playing table tennis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/tableTennis
	HKWorkoutActivityTypeTableTennis HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTaiChi - The constant for tai chi.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/taiChi
	HKWorkoutActivityTypeTaiChi HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTennis - The constant for playing tennis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/tennis
	HKWorkoutActivityTypeTennis HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTrackAndField - Participating in track and field events, including shot put, javelin, pole vaulting, and related sports.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/trackAndField
	HKWorkoutActivityTypeTrackAndField HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTraditionalStrengthTraining - The constant for strength training exercises primarily using machines or free weights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/traditionalStrengthTraining
	HKWorkoutActivityTypeTraditionalStrengthTraining HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTransition - A constant for the transition time between activities in a multisport workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/transition
	HKWorkoutActivityTypeTransition HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeUnderwaterDiving - The constant for underwater diving.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/underwaterDiving
	HKWorkoutActivityTypeUnderwaterDiving HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeVolleyball - The constant for playing volleyball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/volleyball
	HKWorkoutActivityTypeVolleyball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWalking - The constant for walking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/walking
	HKWorkoutActivityTypeWalking HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWaterFitness - The constant for aerobic exercise performed in shallow water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/waterFitness
	HKWorkoutActivityTypeWaterFitness HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWaterPolo - The constant for playing water polo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/waterPolo
	HKWorkoutActivityTypeWaterPolo HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWaterSports - The constant for a variety of water sports, including water skiing, wake boarding, and related activities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/waterSports
	HKWorkoutActivityTypeWaterSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWheelchairRunPace - The constant for wheelchair workout at running pace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/wheelchairRunPace
	HKWorkoutActivityTypeWheelchairRunPace HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWheelchairWalkPace - The constant for a wheelchair workout at walking pace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/wheelchairWalkPace
	HKWorkoutActivityTypeWheelchairWalkPace HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeWrestling - The constant for wrestling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/wrestling
	HKWorkoutActivityTypeWrestling HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeYoga - The constant for practicing yoga.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/yoga
	HKWorkoutActivityTypeYoga HKWorkoutActivityType = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutEffortRelationshipQueryOptions (2 cases) */
// HKWorkoutEffortRelationshipQueryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQueryOptions
type HKWorkoutEffortRelationshipQueryOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQueryOptions/default
	HKWorkoutEffortRelationshipQueryOptionsDefault HKWorkoutEffortRelationshipQueryOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQueryOptions/mostRelevant
	HKWorkoutEffortRelationshipQueryOptionsMostRelevant HKWorkoutEffortRelationshipQueryOptions = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutEventType (8 cases) */
// HKWorkoutEventType - Constants that represent events occurring during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType
type HKWorkoutEventType uint

const (
	// HKWorkoutEventTypeLap - A constant indicating a lap.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/lap
	HKWorkoutEventTypeLap HKWorkoutEventType = 0
	// HKWorkoutEventTypeMarker - A constant indicating a point of interest during a workout session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/marker
	HKWorkoutEventTypeMarker HKWorkoutEventType = 0
	// HKWorkoutEventTypeMotionPaused - A constant indicating that the system has automatically paused a workout session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/motionPaused
	HKWorkoutEventTypeMotionPaused HKWorkoutEventType = 0
	// HKWorkoutEventTypeMotionResumed - A constant indicating that the system has automatically resumed a workout session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/motionResumed
	HKWorkoutEventTypeMotionResumed HKWorkoutEventType = 0
	// HKWorkoutEventTypePause - A constant indicating that the workout has paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/pause
	HKWorkoutEventTypePause HKWorkoutEventType = 0
	// HKWorkoutEventTypePauseOrResumeRequest - A constant indicating that the user has requested a pause or resume.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/pauseOrResumeRequest
	HKWorkoutEventTypePauseOrResumeRequest HKWorkoutEventType = 0
	// HKWorkoutEventTypeResume - A constant indicating that the workout has resumed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/resume
	HKWorkoutEventTypeResume HKWorkoutEventType = 0
	// HKWorkoutEventTypeSegment - A constant indicating a period of time of interest during a workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEventType/segment
	HKWorkoutEventTypeSegment HKWorkoutEventType = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutSessionLocationType (3 cases) */
// HKWorkoutSessionLocationType - A constant indicating whether the workout session takes place indoors or outdoors.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionLocationType
type HKWorkoutSessionLocationType uint

const (
	// HKWorkoutSessionLocationTypeIndoor - The workout session is indoors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionLocationType/indoor
	HKWorkoutSessionLocationTypeIndoor HKWorkoutSessionLocationType = 0
	// HKWorkoutSessionLocationTypeOutdoor - The workout session is outdoors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionLocationType/outdoor
	HKWorkoutSessionLocationTypeOutdoor HKWorkoutSessionLocationType = 0
	// HKWorkoutSessionLocationTypeUnknown - It is not known whether the workout session is taking place indoors or outdoors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionLocationType/unknown
	HKWorkoutSessionLocationTypeUnknown HKWorkoutSessionLocationType = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutSessionState (6 cases) */
// HKWorkoutSessionState - A workout session’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState
type HKWorkoutSessionState uint

const (
	// HKWorkoutSessionStateEnded - The workout session has ended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/ended
	HKWorkoutSessionStateEnded HKWorkoutSessionState = 0
	// HKWorkoutSessionStateNotStarted - The workout session has not started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/notStarted
	HKWorkoutSessionStateNotStarted HKWorkoutSessionState = 0
	// HKWorkoutSessionStatePaused - The workout session has paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/paused
	HKWorkoutSessionStatePaused HKWorkoutSessionState = 0
	// HKWorkoutSessionStatePrepared - The session is ready but not yet running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/prepared
	HKWorkoutSessionStatePrepared HKWorkoutSessionState = 0
	// HKWorkoutSessionStateRunning - The workout session is running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/running
	HKWorkoutSessionStateRunning HKWorkoutSessionState = 0
	// HKWorkoutSessionStateStopped - The session has stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/stopped
	HKWorkoutSessionStateStopped HKWorkoutSessionState = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutSessionType (2 cases) */
// HKWorkoutSessionType - The type of session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionType
type HKWorkoutSessionType uint

const (
	// HKWorkoutSessionTypeMirrored - A mirrored session, running on the companion iOS device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionType/mirrored
	HKWorkoutSessionTypeMirrored HKWorkoutSessionType = 0
	// HKWorkoutSessionTypePrimary - A primary session running on watchOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionType/primary
	HKWorkoutSessionTypePrimary HKWorkoutSessionType = 0
)

/* debug [enums.gen.go]: Processing enum HKWorkoutSwimmingLocationType (3 cases) */
// HKWorkoutSwimmingLocationType - The possible locations for swimming.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSwimmingLocationType
type HKWorkoutSwimmingLocationType uint

const (
	// HKWorkoutSwimmingLocationTypeOpenWater - The user swam in open water like a lake or ocean.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSwimmingLocationType/openWater
	HKWorkoutSwimmingLocationTypeOpenWater HKWorkoutSwimmingLocationType = 0
	// HKWorkoutSwimmingLocationTypePool - The user swam in a pool.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSwimmingLocationType/pool
	HKWorkoutSwimmingLocationTypePool HKWorkoutSwimmingLocationType = 0
	// HKWorkoutSwimmingLocationTypeUnknown - The swimming location could not be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSwimmingLocationType/unknown
	HKWorkoutSwimmingLocationTypeUnknown HKWorkoutSwimmingLocationType = 0
)


