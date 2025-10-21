// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

// Enum types and constants
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

// HKAppleECGAlgorithmVersion - Version numbers for the algorithm Apple Watch uses to generate an ECG reading.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleECGAlgorithmVersion
type HKAppleECGAlgorithmVersion uint

// HKAppleSleepingBreathingDisturbancesClassification enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification
type HKAppleSleepingBreathingDisturbancesClassification uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification/notElevated
	HKAppleSleepingBreathingDisturbancesClassificationNotElevated HKAppleSleepingBreathingDisturbancesClassification = 0
)

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

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSensitivityTestSide/left
	HKAudiogramSensitivityTestSideLeft HKAudiogramSensitivityTestSide = 0
)

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

// HKBiologicalSex - Constants indicating the user’s sex.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex
type HKBiologicalSex uint

const (
	// HKBiologicalSexFemale - A constant indicating that the user is female.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex/female
	HKBiologicalSexFemale HKBiologicalSex = 0
	// HKBiologicalSexOther - A constant indicating that the user is otherwise not categorized as either male or female.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSex/other
	HKBiologicalSexOther HKBiologicalSex = 0
)

// HKBloodType - Constants indicating the user’s blood type.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType
type HKBloodType uint

const (
	// HKBloodTypeANegative - The user has an A– blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/aNegative
	HKBloodTypeANegative HKBloodType = 0
	// HKBloodTypeAPositive - The user has an A+ blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/aPositive
	HKBloodTypeAPositive HKBloodType = 0
	// HKBloodTypeABNegative - The user has an AB– blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/abNegative
	HKBloodTypeABNegative HKBloodType = 0
	// HKBloodTypeABPositive - The user has an AB+ blood type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodType/abPositive
	HKBloodTypeABPositive HKBloodType = 0
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

// HKCategoryValueAppetiteChanges - Categories that represent change in appetite.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges
type HKCategoryValueAppetiteChanges uint

const (
	// HKCategoryValueAppetiteChangesUnspecified - An unspecified change in appetite.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueAppetiteChanges/unspecified
	HKCategoryValueAppetiteChangesUnspecified HKCategoryValueAppetiteChanges = 0
)

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

// HKCategoryValueContraceptive - The type of contraceptive.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive
type HKCategoryValueContraceptive uint

const (
	// HKCategoryValueContraceptiveInjection - An injectable contraceptive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueContraceptive/injection
	HKCategoryValueContraceptiveInjection HKCategoryValueContraceptive = 0
)

// HKCategoryValueEnvironmentalAudioExposureEvent - Exposure events for environmental audio.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueEnvironmentalAudioExposureEvent
type HKCategoryValueEnvironmentalAudioExposureEvent uint

// HKCategoryValueHeadphoneAudioExposureEvent - Exposure events for headphone audio.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueHeadphoneAudioExposureEvent
type HKCategoryValueHeadphoneAudioExposureEvent uint

// HKCategoryValueOvulationTestResult - Categories that represent the result of an ovulation home test.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult
type HKCategoryValueOvulationTestResult uint

const (
	// HKCategoryValueOvulationTestResultEstrogenSurge - The ovulation test detected a surge in estrogen. This value often refers to a   result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueOvulationTestResult/estrogenSurge
	HKCategoryValueOvulationTestResultEstrogenSurge HKCategoryValueOvulationTestResult = 0
)

// HKCategoryValueSeverity - Categories that represent the severity of a symptom.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity
type HKCategoryValueSeverity uint

const (
	// HKCategoryValueSeverityNotPresent - The symptom is not present.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/notPresent
	HKCategoryValueSeverityNotPresent HKCategoryValueSeverity = 0
	// HKCategoryValueSeveritySevere - The symptom is severe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSeverity/severe
	HKCategoryValueSeveritySevere HKCategoryValueSeverity = 0
)

// HKCategoryValueSleepAnalysis - Categories that represent the result of a sleep analysis.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis
type HKCategoryValueSleepAnalysis uint

const (
	// HKCategoryValueSleepAnalysisAsleepUnspecified - The user is asleep, but the specific stage isn’t known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysis/asleepUnspecified
	HKCategoryValueSleepAnalysisAsleepUnspecified HKCategoryValueSleepAnalysis = 0
)

// HKCategoryValueVaginalBleeding enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding
type HKCategoryValueVaginalBleeding uint

const (
	// HKCategoryValueVaginalBleedingMedium - Medium vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/medium
	HKCategoryValueVaginalBleedingMedium HKCategoryValueVaginalBleeding = 0
	// HKCategoryValueVaginalBleedingNone - No vaginal bleeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueVaginalBleeding/none
	HKCategoryValueVaginalBleedingNone HKCategoryValueVaginalBleeding = 0
)

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

const (
	// HKFitzpatrickSkinTypeII - White skin that burns easily and tans minimally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/II
	HKFitzpatrickSkinTypeII HKFitzpatrickSkinType = 0
	// HKFitzpatrickSkinTypeNotSet - Either the user’s skin type is not set, or the user has not granted your app permission to read the skin type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFitzpatrickSkinType/notSet
	HKFitzpatrickSkinTypeNotSet HKFitzpatrickSkinType = 0
)

// HKGAD7AssessmentAnswer enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer
type HKGAD7AssessmentAnswer uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Answer/moreThanHalfTheDays
	HKGAD7AssessmentAnswerMoreThanHalfTheDays HKGAD7AssessmentAnswer = 0
)

// HKGAD7AssessmentRisk enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum
type HKGAD7AssessmentRisk uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGAD7Assessment/Risk-swift.enum/mild
	HKGAD7AssessmentRiskMild HKGAD7AssessmentRisk = 0
)

// HKHeartRateMotionContext - Values that indicate the user’s level of activity when the heart rate sample was measured.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHeartRateMotionContext
type HKHeartRateMotionContext uint

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

// HKPHQ9AssessmentAnswer enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Answer
type HKPHQ9AssessmentAnswer uint

// HKPHQ9AssessmentRisk enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum
type HKPHQ9AssessmentRisk uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPHQ9Assessment/Risk-swift.enum/moderatelySevere
	HKPHQ9AssessmentRiskModeratelySevere HKPHQ9AssessmentRisk = 0
)

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

// HKStateOfMindAssociation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association
type HKStateOfMindAssociation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Association/community
	HKStateOfMindAssociationCommunity HKStateOfMindAssociation = 0
)

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

// HKStateOfMindLabel enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label
type HKStateOfMindLabel uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/content
	HKStateOfMindLabelContent HKStateOfMindLabel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/Label/disgusted
	HKStateOfMindLabelDisgusted HKStateOfMindLabel = 0
)

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
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/unpleasant
	HKStateOfMindValenceClassificationUnpleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/veryPleasant
	HKStateOfMindValenceClassificationVeryPleasant HKStateOfMindValenceClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/ValenceClassification-swift.enum/veryUnpleasant
	HKStateOfMindValenceClassificationVeryUnpleasant HKStateOfMindValenceClassification = 0
)

// HKStatisticsOptions - Options for specifying the statistic to calculate.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStatisticsOptions
type HKStatisticsOptions uint

// HKUpdateFrequency - Constants that determine how often the system launches your app in response to changes to HealthKit data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency
type HKUpdateFrequency uint

const (
	// HKUpdateFrequencyHourly - The system launches your app at most once an hour in response to changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/hourly
	HKUpdateFrequencyHourly HKUpdateFrequency = 0
	// HKUpdateFrequencyImmediate - The system launches your app every time it detects a change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUpdateFrequency/immediate
	HKUpdateFrequencyImmediate HKUpdateFrequency = 0
)

// HKWheelchairUse - Constants indicating the user’s wheelchair use.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse
type HKWheelchairUse uint

const (
	// HKWheelchairUseNo - The user does not use a wheelchair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUse/no
	HKWheelchairUseNo HKWheelchairUse = 0
)

// HKWorkoutActivityType - The type of activity performed during a workout.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType
type HKWorkoutActivityType uint

const (
	// HKWorkoutActivityTypeBaseball - The constant for playing baseball.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/baseball
	HKWorkoutActivityTypeBaseball HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeEquestrianSports - The constant for activities that involve riding a horse, including polo, horse racing, and horse riding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/equestrianSports
	HKWorkoutActivityTypeEquestrianSports HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeFencing - The constant for fencing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/fencing
	HKWorkoutActivityTypeFencing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeHockey - The constant for playing hockey, including ice hockey, field hockey, and related sports.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/hockey
	HKWorkoutActivityTypeHockey HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSailing - The constant for sailing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/sailing
	HKWorkoutActivityTypeSailing HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeSwimming - The constant for swimming.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/swimming
	HKWorkoutActivityTypeSwimming HKWorkoutActivityType = 0
	// HKWorkoutActivityTypeTransition - A constant for the transition time between activities in a multisport workout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutActivityType/transition
	HKWorkoutActivityTypeTransition HKWorkoutActivityType = 0
)

// HKWorkoutEffortRelationshipQueryOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutEffortRelationshipQueryOptions
type HKWorkoutEffortRelationshipQueryOptions uint

// HKWorkoutSessionLocationType - A constant indicating whether the workout session takes place indoors or outdoors.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionLocationType
type HKWorkoutSessionLocationType uint

const (
	// HKWorkoutSessionLocationTypeUnknown - It is not known whether the workout session is taking place indoors or outdoors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionLocationType/unknown
	HKWorkoutSessionLocationTypeUnknown HKWorkoutSessionLocationType = 0
)

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
	// HKWorkoutSessionStateRunning - The workout session is running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/running
	HKWorkoutSessionStateRunning HKWorkoutSessionState = 0
	// HKWorkoutSessionStateStopped - The session has stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionState/stopped
	HKWorkoutSessionStateStopped HKWorkoutSessionState = 0
)

// HKWorkoutSessionType - The type of session.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionType
type HKWorkoutSessionType uint

const (
	// HKWorkoutSessionTypePrimary - A primary session running on watchOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWorkoutSessionType/primary
	HKWorkoutSessionTypePrimary HKWorkoutSessionType = 0
)

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


