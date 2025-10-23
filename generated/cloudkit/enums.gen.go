// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

// Enum types and constants
// CKAccountStatus - Constants that indicate the availability of the user’s iCloud account.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAccountStatus
type CKAccountStatus uint

const (
	// CKAccountStatusAvailable - The user’s iCloud account is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAccountStatus/available
	CKAccountStatusAvailable CKAccountStatus = 0
	// CKAccountStatusCouldNotDetermine - CloudKit can’t determine the status of the user’s iCloud account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAccountStatus/couldNotDetermine
	CKAccountStatusCouldNotDetermine CKAccountStatus = 0
	// CKAccountStatusNoAccount - The device doesn’t have an iCloud account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAccountStatus/noAccount
	CKAccountStatusNoAccount CKAccountStatus = 0
	// CKAccountStatusRestricted - The system denies access to the user’s iCloud account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAccountStatus/restricted
	CKAccountStatusRestricted CKAccountStatus = 0
	// CKAccountStatusTemporarilyUnavailable - The user’s iCloud account is temporarily unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAccountStatus/temporarilyUnavailable
	CKAccountStatusTemporarilyUnavailable CKAccountStatus = 0
)

// CKDatabaseScope - Constants that represent the scope of a database.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope
type CKDatabaseScope uint

const (
	// CKDatabaseScopePublic - The public database.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope/public
	CKDatabaseScopePublic CKDatabaseScope = 0
	// CKDatabaseScopeShared - The shared database.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabase/Scope/shared
	CKDatabaseScopeShared CKDatabaseScope = 0
)

// CKErrorCode - The error codes that CloudKit returns.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code
type CKErrorCode uint

const (
	// CKErrorAccountTemporarilyUnavailable - An error that occurs when the user’s iCloud account is temporarily unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/accountTemporarilyUnavailable
	CKErrorAccountTemporarilyUnavailable CKErrorCode = 0
	// CKErrorAlreadyShared - An error that occurs when CloudKit attempts to share a record with an existing share.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/alreadyShared
	CKErrorAlreadyShared CKErrorCode = 0
	// CKErrorAssetFileModified - An error that occurs when the system modifies an asset while saving it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/assetFileModified
	CKErrorAssetFileModified CKErrorCode = 0
	// CKErrorAssetFileNotFound - An error that occurs when the system can’t find the specified asset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/assetFileNotFound
	CKErrorAssetFileNotFound CKErrorCode = 0
	// CKErrorAssetNotAvailable - An error that occurs when the system can’t access the specified asset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/assetNotAvailable
	CKErrorAssetNotAvailable CKErrorCode = 0
	// CKErrorBadContainer - An error that occurs when you use an unknown or unauthorized container.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/badContainer
	CKErrorBadContainer CKErrorCode = 0
	// CKErrorBadDatabase - An error that occurs when the operation can’t complete for the specified database.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/badDatabase
	CKErrorBadDatabase CKErrorCode = 0
	// CKErrorBatchRequestFailed - An error that occurs when the system rejects the entire batch of changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/batchRequestFailed
	CKErrorBatchRequestFailed CKErrorCode = 0
	// CKErrorChangeTokenExpired - An error that occurs when the change token expires.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/changeTokenExpired
	CKErrorChangeTokenExpired CKErrorCode = 0
	// CKErrorConstraintViolation - An error that occurs when the server rejects the request because of a unique constraint violation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/constraintViolation
	CKErrorConstraintViolation CKErrorCode = 0
	// CKErrorIncompatibleVersion - An error that occurs when the current app version is older than the oldest allowed version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/incompatibleVersion
	CKErrorIncompatibleVersion CKErrorCode = 0
	// CKErrorInternalError - A nonrecoverable error that CloudKit encounters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/internalError
	CKErrorInternalError CKErrorCode = 0
	// CKErrorInvalidArguments - An error that occurs when the request contains invalid information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/invalidArguments
	CKErrorInvalidArguments CKErrorCode = 0
	// CKErrorLimitExceeded - An error that occurs when a request’s size exceeds the limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/limitExceeded
	CKErrorLimitExceeded CKErrorCode = 0
	// CKErrorManagedAccountRestricted - An error that occurs when CloudKit rejects a request due to a managed-account restriction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/managedAccountRestricted
	CKErrorManagedAccountRestricted CKErrorCode = 0
	// CKErrorMissingEntitlement - An error that occurs when the app is missing a required entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/missingEntitlement
	CKErrorMissingEntitlement CKErrorCode = 0
	// CKErrorNetworkFailure - An error that occurs when a network is available, but CloudKit is inaccessible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/networkFailure
	CKErrorNetworkFailure CKErrorCode = 0
	// CKErrorNetworkUnavailable - An error that occurs when the network is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/networkUnavailable
	CKErrorNetworkUnavailable CKErrorCode = 0
	// CKErrorNotAuthenticated - An error that occurs when the user is unauthenticated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/notAuthenticated
	CKErrorNotAuthenticated CKErrorCode = 0
	// CKErrorOperationCancelled - An error that occurs when an operation cancels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/operationCancelled
	CKErrorOperationCancelled CKErrorCode = 0
	// CKErrorPartialFailure - An error that occurs when an operation completes with partial failures.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/partialFailure
	CKErrorPartialFailure CKErrorCode = 0
	// CKErrorParticipantAlreadyInvited - The user is already an invited participant on this share. They must accept the existing share invitation before continuing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/participantAlreadyInvited
	CKErrorParticipantAlreadyInvited CKErrorCode = 0
	// CKErrorParticipantMayNeedVerification - An error that occurs when the user isn’t a participant of the share.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/participantMayNeedVerification
	CKErrorParticipantMayNeedVerification CKErrorCode = 0
	// CKErrorPermissionFailure - An error that occurs when the user doesn’t have permission to save or fetch data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/permissionFailure
	CKErrorPermissionFailure CKErrorCode = 0
	// CKErrorQuotaExceeded - An error that occurs when saving a record exceeds the user’s storage quota.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/quotaExceeded
	CKErrorQuotaExceeded CKErrorCode = 0
	// CKErrorReferenceViolation - An error that occurs when CloudKit can’t find the target of a reference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/referenceViolation
	CKErrorReferenceViolation CKErrorCode = 0
	// CKErrorRequestRateLimited - An error that occurs when CloudKit rate-limits requests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/requestRateLimited
	CKErrorRequestRateLimited CKErrorCode = 0
	// CKErrorResultsTruncated - An error that occurs when CloudKit truncates a query’s results.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/resultsTruncated
	CKErrorResultsTruncated CKErrorCode = 0
	// CKErrorServerRecordChanged - An error that occurs when CloudKit rejects a record because the server’s version is different.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverRecordChanged
	CKErrorServerRecordChanged CKErrorCode = 0
	// CKErrorServerRejectedRequest - An error that occurs when CloudKit rejects the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverRejectedRequest
	CKErrorServerRejectedRequest CKErrorCode = 0
	// CKErrorServerResponseLost - An error that occurs when CloudKit is unable to maintain the network connection and provide a response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverResponseLost
	CKErrorServerResponseLost CKErrorCode = 0
	// CKErrorServiceUnavailable - An error that occurs when CloudKit is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serviceUnavailable
	CKErrorServiceUnavailable CKErrorCode = 0
	// CKErrorTooManyParticipants - An error that occurs when a share has too many participants.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/tooManyParticipants
	CKErrorTooManyParticipants CKErrorCode = 0
	// CKErrorUnknownItem - An error that occurs when the specified record doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/unknownItem
	CKErrorUnknownItem CKErrorCode = 0
	// CKErrorUserDeletedZone - An error that occurs when the user deletes a record zone using the Settings app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/userDeletedZone
	CKErrorUserDeletedZone CKErrorCode = 0
	// CKErrorZoneBusy - An error that occurs when the server is too busy to handle the record zone operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/zoneBusy
	CKErrorZoneBusy CKErrorCode = 0
	// CKErrorZoneNotFound - An error that occurs when the specified record zone doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKError/Code/zoneNotFound
	CKErrorZoneNotFound CKErrorCode = 0
)

// CKOperationGroupTransferSize - Constants that represent possible data transfer sizes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize
type CKOperationGroupTransferSize uint

// CKRecordZoneEncryptionScope enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/EncryptionScope-swift.enum
type CKRecordZoneEncryptionScope uint

const (
	// CKRecordZoneEncryptionScopePerRecord - Zone uses per-record encryption keys for any encrypted values on a record or share.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/EncryptionScope-swift.enum/perRecord
	CKRecordZoneEncryptionScopePerRecord CKRecordZoneEncryptionScope = 0
	// CKRecordZoneEncryptionScopePerZone - Zone uses per-zone encryption keys for encrypted values across all records and the zone-wide share, if present.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/EncryptionScope-swift.enum/perZone
	CKRecordZoneEncryptionScopePerZone CKRecordZoneEncryptionScope = 0
)

// CKSharingParticipantPermissionOption - An object that controls participant permission options.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantPermissionOption
type CKSharingParticipantPermissionOption uint


