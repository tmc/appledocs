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
	CXCallDirectoryEnabledStatusDisabled CXCallDirectoryEnabledStatus = 1
	// CXCallDirectoryEnabledStatusEnabled - Indicates that the extension is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/EnabledStatus/enabled
	CXCallDirectoryEnabledStatusEnabled CXCallDirectoryEnabledStatus = 2
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
	CXCallEndedReasonAnsweredElsewhere CXCallEndedReason = 4
	// CXCallEndedReasonDeclinedElsewhere - Another device declined the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/declinedElsewhere
	CXCallEndedReasonDeclinedElsewhere CXCallEndedReason = 5
	// CXCallEndedReasonFailed - An error occurred while attempting to service the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/failed
	CXCallEndedReasonFailed CXCallEndedReason = 1
	// CXCallEndedReasonRemoteEnded - The remote party explicitly ended the call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/remoteEnded
	CXCallEndedReasonRemoteEnded CXCallEndedReason = 2
	// CXCallEndedReasonUnanswered - The call never started connecting and was never explicitly ended, such as when an outgoing or incoming call times out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallEndedReason/unanswered
	CXCallEndedReasonUnanswered CXCallEndedReason = 3
)

// CXErrorCode - Error codes for the CallKit errors.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXError/Code
type CXErrorCode uint

// CXErrorCodeCallDirectoryManagerError - Error codes the CallKit framework returns.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code
type CXErrorCodeCallDirectoryManagerError uint

const (
	// CXErrorCodeCallDirectoryManagerErrorCurrentlyLoading - The call directory manager is loading the app extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/currentlyLoading
	CXErrorCodeCallDirectoryManagerErrorCurrentlyLoading CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorDuplicateEntries - There are duplicate entries in the call directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/duplicateEntries
	CXErrorCodeCallDirectoryManagerErrorDuplicateEntries CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorEntriesOutOfOrder - The entries in the call directory are out of order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/entriesOutOfOrder
	CXErrorCodeCallDirectoryManagerErrorEntriesOutOfOrder CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorExtensionDisabled - The call directory extension isn’t enabled by the system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/extensionDisabled
	CXErrorCodeCallDirectoryManagerErrorExtensionDisabled CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorLoadingInterrupted - The call directory manager was interrupted while loading the app extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/loadingInterrupted
	CXErrorCodeCallDirectoryManagerErrorLoadingInterrupted CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorMaximumEntriesExceeded - There are too many entries in the call directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/maximumEntriesExceeded
	CXErrorCodeCallDirectoryManagerErrorMaximumEntriesExceeded CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorNoExtensionFound - The call directory manager could not find a corresponding app extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/noExtensionFound
	CXErrorCodeCallDirectoryManagerErrorNoExtensionFound CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorUnexpectedIncrementalRemoval - A request occurred before confirming incremental loading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/unexpectedIncrementalRemoval
	CXErrorCodeCallDirectoryManagerErrorUnexpectedIncrementalRemoval CXErrorCodeCallDirectoryManagerError = 0
	// CXErrorCodeCallDirectoryManagerErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeCallDirectoryManagerError-swift.struct/Code/unknown
	CXErrorCodeCallDirectoryManagerErrorUnknown CXErrorCodeCallDirectoryManagerError = 0
)

// CXErrorCodeIncomingCallError - Codes for errors that occur during incoming calls.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code
type CXErrorCodeIncomingCallError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/callIsProtected
	CXErrorCodeIncomingCallErrorCallIsProtected CXErrorCodeIncomingCallError = 0
	// CXErrorCodeIncomingCallErrorCallUUIDAlreadyExists - The incoming call UUID already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/callUUIDAlreadyExists
	CXErrorCodeIncomingCallErrorCallUUIDAlreadyExists CXErrorCodeIncomingCallError = 0
	// CXErrorCodeIncomingCallErrorFilteredByBlockList - The incoming call is filtered because the incoming caller has been blocked by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/filteredByBlockList
	CXErrorCodeIncomingCallErrorFilteredByBlockList CXErrorCodeIncomingCallError = 0
	// CXErrorCodeIncomingCallErrorFilteredByDoNotDisturb - The incoming call is filtered because Do Not Disturb is active and the incoming caller is not a VIP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/filteredByDoNotDisturb
	CXErrorCodeIncomingCallErrorFilteredByDoNotDisturb CXErrorCodeIncomingCallError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/filteredBySensitiveParticipants
	CXErrorCodeIncomingCallErrorFilteredBySensitiveParticipants CXErrorCodeIncomingCallError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/filteredDuringRestrictedSharingMode
	CXErrorCodeIncomingCallErrorFilteredDuringRestrictedSharingMode CXErrorCodeIncomingCallError = 0
	// CXErrorCodeIncomingCallErrorUnentitled - The app isn’t entitled to receive incoming calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/unentitled
	CXErrorCodeIncomingCallErrorUnentitled CXErrorCodeIncomingCallError = 0
	// CXErrorCodeIncomingCallErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeIncomingCallError-swift.struct/Code/unknown
	CXErrorCodeIncomingCallErrorUnknown CXErrorCodeIncomingCallError = 0
)

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

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/callIsProtected
	CXErrorCodeRequestTransactionErrorCallIsProtected CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorCallUUIDAlreadyExists - The requested transaction contains call actions that reference a UUID that already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/callUUIDAlreadyExists
	CXErrorCodeRequestTransactionErrorCallUUIDAlreadyExists CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorEmptyTransaction - The requested transaction contains no actions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/emptyTransaction
	CXErrorCodeRequestTransactionErrorEmptyTransaction CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorInvalidAction - The requested transaction contains an invalid action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/invalidAction
	CXErrorCodeRequestTransactionErrorInvalidAction CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorMaximumCallGroupsReached - The requested transaction contains actions that, if performed, would exceed the maximum number of call groups for the provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/maximumCallGroupsReached
	CXErrorCodeRequestTransactionErrorMaximumCallGroupsReached CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorUnentitled - The app isn’t entitled to perform the actions in the requested transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/unentitled
	CXErrorCodeRequestTransactionErrorUnentitled CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/unknown
	CXErrorCodeRequestTransactionErrorUnknown CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorUnknownCallProvider - The controller couldn’t find a call provider to perform the actions in the requested transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/unknownCallProvider
	CXErrorCodeRequestTransactionErrorUnknownCallProvider CXErrorCodeRequestTransactionError = 0
	// CXErrorCodeRequestTransactionErrorUnknownCallUUID - The requested transaction contains call actions that reference an unknown UUID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXErrorCodeRequestTransactionError-swift.struct/Code/unknownCallUUID
	CXErrorCodeRequestTransactionErrorUnknownCallUUID CXErrorCodeRequestTransactionError = 0
)

// CXHandleType - The possible types of handles.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType
type CXHandleType uint

const (
	// CXHandleTypeEmailAddress - An email address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType/emailAddress
	CXHandleTypeEmailAddress CXHandleType = 3
	// CXHandleTypeGeneric - An unspecified type of handle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType/generic
	CXHandleTypeGeneric CXHandleType = 1
	// CXHandleTypePhoneNumber - A phone number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/HandleType/phoneNumber
	CXHandleTypePhoneNumber CXHandleType = 2
)

// CXPlayDTMFCallActionType - The types of events that generate dial tones.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/ActionType
type CXPlayDTMFCallActionType uint

const (
	CXPlayDTMFCallActionTypeSingleTone CXPlayDTMFCallActionType = 1
	CXPlayDTMFCallActionTypeSoftPause CXPlayDTMFCallActionType = 2
	CXPlayDTMFCallActionTypeHardPause CXPlayDTMFCallActionType = 3
	CXCallEndedReasonFailed CXPlayDTMFCallActionType = 1
	CXCallEndedReasonRemoteEnded CXPlayDTMFCallActionType = 2
	CXCallEndedReasonUnanswered CXPlayDTMFCallActionType = 3
	CXCallEndedReasonAnsweredElsewhere CXPlayDTMFCallActionType = 4
	CXCallEndedReasonDeclinedElsewhere CXPlayDTMFCallActionType = 5
	CXCallDirectoryEnabledStatusUnknown CXPlayDTMFCallActionType = 0
	CXCallDirectoryEnabledStatusDisabled CXPlayDTMFCallActionType = 1
	CXCallDirectoryEnabledStatusEnabled CXPlayDTMFCallActionType = 2
)

// CXTranslationEngine - Values that describe the translation engine that provided a translation.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTranslationEngine
type CXTranslationEngine uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTranslationEngine/custom
	CXTranslationEngineCustom CXTranslationEngine = 0
	// CXTranslationEngineDefault - The translation was provided by the system’s default translation engine.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTranslationEngine/default
	CXTranslationEngineDefault CXTranslationEngine = 0
)


