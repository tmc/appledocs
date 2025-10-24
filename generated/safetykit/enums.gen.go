// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

/* debug [enums.gen.go]: Generating 4 enums for SafetyKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SAAuthorizationStatus (3 cases) */
// SAAuthorizationStatus - An enumeration that represents the current Crash Detection event authorization state.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAAuthorizationStatus
type SAAuthorizationStatus uint

const (
	// SAAuthorizationStatusAuthorized - This is the designated app for receiving Crash Detection events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAAuthorizationStatus/authorized
	SAAuthorizationStatusAuthorized SAAuthorizationStatus = 0
	// SAAuthorizationStatusDenied - The system denies the app from receiving Crash Detection events because another app has authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAAuthorizationStatus/denied
	SAAuthorizationStatusDenied SAAuthorizationStatus = 0
	// SAAuthorizationStatusNotDetermined - There isn’t a designated app for receiving Crash Detection events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAAuthorizationStatus/notDetermined
	SAAuthorizationStatusNotDetermined SAAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum SACrashDetectionEventResponse (2 cases) */
// SACrashDetectionEventResponse - An enumeration that defines possible emergency responses to a Crash Detection event.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/Response-swift.enum
type SACrashDetectionEventResponse uint

const (
	// SACrashDetectionEventResponseAttempted - The system attempted to dial the Emergency SOS - Call After Severe Crash provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/Response-swift.enum/attempted
	SACrashDetectionEventResponseAttempted SACrashDetectionEventResponse = 0
	// SACrashDetectionEventResponseDisabled - The system couldn’t contact the Emergency SOS - Call After Severe Crash provider because the feature is off in the Settings app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/Response-swift.enum/disabled
	SACrashDetectionEventResponseDisabled SACrashDetectionEventResponse = 0
)

/* debug [enums.gen.go]: Processing enum SAEmergencyResponseManagerVoiceCallStatus (4 cases) */
// SAEmergencyResponseManagerVoiceCallStatus - An enumeration that defines the status of a requested voice call.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/VoiceCallStatus
type SAEmergencyResponseManagerVoiceCallStatus uint

const (
	// SAEmergencyResponseManagerVoiceCallStatusActive - The system successfully placed a call to the desired contact and that call is currently active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/VoiceCallStatus/active
	SAEmergencyResponseManagerVoiceCallStatusActive SAEmergencyResponseManagerVoiceCallStatus = 0
	// SAEmergencyResponseManagerVoiceCallStatusDialing - The system is dialing the desired contact.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/VoiceCallStatus/dialing
	SAEmergencyResponseManagerVoiceCallStatusDialing SAEmergencyResponseManagerVoiceCallStatus = 0
	// SAEmergencyResponseManagerVoiceCallStatusDisconnected - The voice call to the desired contact disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/VoiceCallStatus/disconnected
	SAEmergencyResponseManagerVoiceCallStatusDisconnected SAEmergencyResponseManagerVoiceCallStatus = 0
	// SAEmergencyResponseManagerVoiceCallStatusFailed - The voice call failed to connect to the desired contact.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/VoiceCallStatus/failed
	SAEmergencyResponseManagerVoiceCallStatusFailed SAEmergencyResponseManagerVoiceCallStatus = 0
)

/* debug [enums.gen.go]: Processing enum SAErrorCode (4 cases) */
// SAErrorCode - Codes for identifying errors in SafetyKit.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAError/Code
type SAErrorCode uint

const (
	// SAErrorInvalidArgument - The passed argument is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAError/Code/invalidArgument
	SAErrorInvalidArgument SAErrorCode = 0
	// SAErrorNotAllowed - The system restricts the feature on this iPhone at the current time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAError/Code/notAllowed
	SAErrorNotAllowed SAErrorCode = 0
	// SAErrorNotAuthorized - The app isn’t authorized to perform the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAError/Code/notAuthorized
	SAErrorNotAuthorized SAErrorCode = 0
	// SAErrorOperationFailed - The requested operation failed; retrying may succeed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAError/Code/operationFailed
	SAErrorOperationFailed SAErrorCode = 0
)


