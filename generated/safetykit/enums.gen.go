// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

// Enum types and constants
// SAAuthorizationStatus - An enumeration that represents the current Crash Detection event authorization state.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAAuthorizationStatus
type SAAuthorizationStatus uint

// SACrashDetectionEventResponse - An enumeration that defines possible emergency responses to a Crash Detection event.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/Response-swift.enum
type SACrashDetectionEventResponse uint

const (
	// SACrashDetectionEventResponseAttempted - The system attempted to dial the Emergency SOS - Call After Severe Crash provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SACrashDetectionEvent/Response-swift.enum/attempted
	SACrashDetectionEventResponseAttempted SACrashDetectionEventResponse = 0
)

// SAEmergencyResponseManagerVoiceCallStatus - An enumeration that defines the status of a requested voice call.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit/SAEmergencyResponseManager/VoiceCallStatus
type SAEmergencyResponseManagerVoiceCallStatus uint

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
)


