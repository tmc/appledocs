// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

// Enum types and constants
// CXCallDirectoryEnabledStatus - The enabled status of a Call Directory app extension, as reported by the 
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/EnabledStatus
type CXCallDirectoryEnabledStatus uint

const (
	// CXCallDirectoryEnabledStatusDisabled - Indicates that the extension is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/EnabledStatus/disabled
	CXCallDirectoryEnabledStatusDisabled CXCallDirectoryEnabledStatus = 0
	// CXCallDirectoryEnabledStatusEnabled - Indicates that the extension is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/EnabledStatus/enabled
	CXCallDirectoryEnabledStatusEnabled CXCallDirectoryEnabledStatus = 0
	// CXCallDirectoryEnabledStatusUnknown - Indicates that the enabled status for the extension is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/EnabledStatus/unknown
	CXCallDirectoryEnabledStatusUnknown CXCallDirectoryEnabledStatus = 0
)

// CXCallEndedReason - The reason that a call ended.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason
type CXCallEndedReason uint

const (
	// CXCallEndedReasonAnsweredElsewhere - Another device answered the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/answeredElsewhere
	CXCallEndedReasonAnsweredElsewhere CXCallEndedReason = 0
	// CXCallEndedReasonDeclinedElsewhere - Another device declined the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/declinedElsewhere
	CXCallEndedReasonDeclinedElsewhere CXCallEndedReason = 0
	// CXCallEndedReasonFailed - An error occurred while attempting to service the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/failed
	CXCallEndedReasonFailed CXCallEndedReason = 0
	// CXCallEndedReasonRemoteEnded - The remote party explicitly ended the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/remoteEnded
	CXCallEndedReasonRemoteEnded CXCallEndedReason = 0
	// CXCallEndedReasonUnanswered - The call never started connecting and was never explicitly ended, such as when an outgoing or incoming call times out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/unanswered
	CXCallEndedReasonUnanswered CXCallEndedReason = 0
)

// CXErrorCode - Error codes for the CallKit errors.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXError/Code
type CXErrorCode uint

// CXErrorCodeCallDirectoryManagerError - Error codes the CallKit framework returns.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code
type CXErrorCodeCallDirectoryManagerError uint

// CXErrorCodeIncomingCallError - Codes for errors that occur during incoming calls.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code
type CXErrorCodeIncomingCallError uint

// CXErrorCodeNotificationServiceExtensionError - Constants for errors returned when reporting new, incoming VoIP calls.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeNotificationServiceExtensionError-swift.struct/Code
type CXErrorCodeNotificationServiceExtensionError uint

const (
	// CXErrorCodeNotificationServiceExtensionErrorInvalidClientProcess - An error indicating that an invalid client process reported the incoming call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeNotificationServiceExtensionError-swift.struct/Code/invalidClientProcess
	CXErrorCodeNotificationServiceExtensionErrorInvalidClientProcess CXErrorCodeNotificationServiceExtensionError = 0
	// CXErrorCodeNotificationServiceExtensionErrorMissingNotificationFilteringEntitlement - An error indicating that the notification service extension is missing the required filtering entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeNotificationServiceExtensionError-swift.struct/Code/missingNotificationFilteringEntitlement
	CXErrorCodeNotificationServiceExtensionErrorMissingNotificationFilteringEntitlement CXErrorCodeNotificationServiceExtensionError = 0
	// CXErrorCodeNotificationServiceExtensionErrorUnknown - An error that occurs when there is an unknown problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeNotificationServiceExtensionError-swift.struct/Code/unknown
	CXErrorCodeNotificationServiceExtensionErrorUnknown CXErrorCodeNotificationServiceExtensionError = 0
)

// CXErrorCodeRequestTransactionError - Error codes for the CallKit error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code
type CXErrorCodeRequestTransactionError uint

// CXHandleType - The possible types of handles.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType
type CXHandleType uint

const (
	// CXHandleTypeEmailAddress - An email address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType/emailAddress
	CXHandleTypeEmailAddress CXHandleType = 0
	// CXHandleTypeGeneric - An unspecified type of handle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType/generic
	CXHandleTypeGeneric CXHandleType = 0
	// CXHandleTypePhoneNumber - A phone number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType/phoneNumber
	CXHandleTypePhoneNumber CXHandleType = 0
)

// CXPlayDTMFCallActionType - The types of events that generate dial tones.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/ActionType
type CXPlayDTMFCallActionType uint

// CXTranslationEngine - Values that describe the translation engine that provided a translation.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTranslationEngine
type CXTranslationEngine uint


