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

// CKApplicationPermissionStatus - Constants that represent the status of a permission.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/ApplicationPermissionStatus
type CKApplicationPermissionStatus uint

// CKApplicationPermissions - Constants that represent the permissions that a user grants.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/ApplicationPermissions
type CKApplicationPermissions uint

const (
// CKApplicationPermissionUserDiscoverability - The user is discoverable using their email address.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKContainer/ApplicationPermissions/userDiscoverability
CKApplicationPermissionUserDiscoverability CKApplicationPermissions = 0
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

// CKRecordSavePolicy - Constants that indicate which policy to apply when saving records.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy
type CKRecordSavePolicy uint

const (
// CKRecordSaveIfServerRecordUnchanged - A policy that instructs CloudKit to only proceed if the record’s change tag matches that of the server’s copy.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy/ifServerRecordUnchanged
CKRecordSaveIfServerRecordUnchanged CKRecordSavePolicy = 0
)

// CKNotificationType - Constants that indicate the type of event that generates the push notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/NotificationType-swift.enum
type CKNotificationType uint

// CKOperationGroupTransferSize - Constants that represent possible data transfer sizes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize
type CKOperationGroupTransferSize uint

const (
// CKOperationGroupTransferSizeGigabytes - A transfer size that represents 1 or more gigabytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/gigabytes
CKOperationGroupTransferSizeGigabytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeHundredsOfGigabytes - A transfer size that represents hundreds of gigabytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/hundredsOfGigabytes
CKOperationGroupTransferSizeHundredsOfGigabytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeHundredsOfMegabytes - A transfer size that represents hundreds of megabytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/hundredsOfMegabytes
CKOperationGroupTransferSizeHundredsOfMegabytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeKilobytes - A transfer size that represents 1 or more kilobytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/kilobytes
CKOperationGroupTransferSizeKilobytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeMegabytes - A transfer size that represents 1 or more megabytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/megabytes
CKOperationGroupTransferSizeMegabytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeTensOfGigabytes - A transfer size that represents tens of gigabytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/tensOfGigabytes
CKOperationGroupTransferSizeTensOfGigabytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeTensOfMegabytes - A transfer size that represents tens of megabytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/tensOfMegabytes
CKOperationGroupTransferSizeTensOfMegabytes CKOperationGroupTransferSize = 0
// CKOperationGroupTransferSizeUnknown - An unknown transfer size.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKOperationGroup/TransferSize/unknown
CKOperationGroupTransferSizeUnknown CKOperationGroupTransferSize = 0
)

// CKQuerySubscriptionOptions - Configuration options for a query subscription.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/Options
type CKQuerySubscriptionOptions uint

const (
// CKQuerySubscriptionOptionsFiresOnRecordCreation - An option that instructs CloudKit to send a push notification when it creates a record that matches a subscription’s criteria.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/Options/firesOnRecordCreation
CKQuerySubscriptionOptionsFiresOnRecordCreation CKQuerySubscriptionOptions = 0
// CKQuerySubscriptionOptionsFiresOnRecordDeletion - An option that instructs CloudKit to send a push notification when it deletes a record that matches a subscription’s criteria.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/Options/firesOnRecordDeletion
CKQuerySubscriptionOptionsFiresOnRecordDeletion CKQuerySubscriptionOptions = 0
// CKQuerySubscriptionOptionsFiresOnRecordUpdate - An option that instructs CloudKit to send a push notification when it modifies a record that matches a subscription’s criteria.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/Options/firesOnRecordUpdate
CKQuerySubscriptionOptionsFiresOnRecordUpdate CKQuerySubscriptionOptions = 0
// CKQuerySubscriptionOptionsFiresOnce - An option that instructs CloudKit to send a push notification only once.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/Options/firesOnce
CKQuerySubscriptionOptionsFiresOnce CKQuerySubscriptionOptions = 0
)

// CKReferenceAction - Constants that indicate the behavior when deleting a referenced record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction
type CKReferenceAction uint

const (
// CKReferenceActionDeleteSelf - A reference action that cascades deletions.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
CKReferenceActionDeleteSelf CKReferenceAction = 0
// CKReferenceActionNone - A reference action that has no cascading behavior.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/none
CKReferenceActionNone CKReferenceAction = 0
)

// CKRecordZoneCapabilities - The capabilities that a record zone supports.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct
type CKRecordZoneCapabilities uint

const (
// CKRecordZoneCapabilityAtomic - A capability that allows atomic changes of multiple records.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct/atomic
CKRecordZoneCapabilityAtomic CKRecordZoneCapabilities = 0
// CKRecordZoneCapabilityFetchChanges - A capability for fetching only the changed records from a zone.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct/fetchChanges
CKRecordZoneCapabilityFetchChanges CKRecordZoneCapabilities = 0
// CKRecordZoneCapabilitySharing - A capability for sharing a specific hierarchy of records.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct/sharing
CKRecordZoneCapabilitySharing CKRecordZoneCapabilities = 0
// CKRecordZoneCapabilityZoneWideSharing - A capability for sharing the entire contents of a record zone.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct/zoneWideSharing
CKRecordZoneCapabilityZoneWideSharing CKRecordZoneCapabilities = 0
)

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

// CKShareParticipantAcceptanceStatus - Constants that represent the status of a participant.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus
type CKShareParticipantAcceptanceStatus uint

const (
// CKShareParticipantAcceptanceStatusAccepted - The participant accepted the share request.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus/accepted
CKShareParticipantAcceptanceStatusAccepted CKShareParticipantAcceptanceStatus = 0
// CKShareParticipantAcceptanceStatusPending - The participant’s acceptance of the share request is pending.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus/pending
CKShareParticipantAcceptanceStatusPending CKShareParticipantAcceptanceStatus = 0
// CKShareParticipantAcceptanceStatusRemoved - The system removed the participant from the share.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus/removed
CKShareParticipantAcceptanceStatusRemoved CKShareParticipantAcceptanceStatus = 0
// CKShareParticipantAcceptanceStatusUnknown - The participant’s status is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus/unknown
CKShareParticipantAcceptanceStatusUnknown CKShareParticipantAcceptanceStatus = 0
)

// CKShareParticipantPermission - Constants that represent the permissions to grant to a share participant.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission
type CKShareParticipantPermission uint

const (
// CKShareParticipantPermissionNone - The participant doesn’t have any permissions for the share.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/none
CKShareParticipantPermissionNone CKShareParticipantPermission = 0
// CKShareParticipantPermissionReadOnly - The participant has read-only permissions for the share.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/readOnly
CKShareParticipantPermissionReadOnly CKShareParticipantPermission = 0
// CKShareParticipantPermissionReadWrite - The participant has read-and-write permissions for the share.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/readWrite
CKShareParticipantPermissionReadWrite CKShareParticipantPermission = 0
// CKShareParticipantPermissionUnknown - The participant’s permissions are unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/unknown
CKShareParticipantPermissionUnknown CKShareParticipantPermission = 0
)

// CKShareParticipantRole - Constants that represent the role of a share’s participant.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole
type CKShareParticipantRole uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/administrator
CKShareParticipantRoleAdministrator CKShareParticipantRole = 0
// CKShareParticipantRoleOwner - The participant is the share’s owner.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/owner
CKShareParticipantRoleOwner CKShareParticipantRole = 0
// CKShareParticipantRolePrivateUser - The participant has the private role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/privateUser
CKShareParticipantRolePrivateUser CKShareParticipantRole = 0
// CKShareParticipantRolePublicUser - The participant has the public role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/publicUser
CKShareParticipantRolePublicUser CKShareParticipantRole = 0
// CKShareParticipantRoleUnknown - The participant’s role is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/unknown
CKShareParticipantRoleUnknown CKShareParticipantRole = 0
)

// CKShareParticipantType - The role of a participant.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantType
type CKShareParticipantType uint

const (
// CKShareParticipantTypeOwner - The type of an owner.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantType/owner
CKShareParticipantTypeOwner CKShareParticipantType = 0
// CKShareParticipantTypePrivateUser - The type of a private user.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantType/privateUser
CKShareParticipantTypePrivateUser CKShareParticipantType = 0
// CKShareParticipantTypePublicUser - The type of a public owner.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantType/publicUser
CKShareParticipantTypePublicUser CKShareParticipantType = 0
// CKShareParticipantTypeUnknown - An unknown role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantType/unknown
CKShareParticipantTypeUnknown CKShareParticipantType = 0
)

// CKSyncEngineEventType - Describes an event that occurs during a sync operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType
type CKSyncEngineEventType uint

const (
// CKSyncEngineEventTypeAccountChange - An event indicating a change to the device’s iCloud account.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/accountChange
CKSyncEngineEventTypeAccountChange CKSyncEngineEventType = 0
// CKSyncEngineEventTypeDidFetchChanges - An event that indicates the database fetch is done.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/didFetchChanges
CKSyncEngineEventTypeDidFetchChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeDidFetchRecordZoneChanges - An event that indicates the record zone fetch is done.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/didFetchRecordZoneChanges
CKSyncEngineEventTypeDidFetchRecordZoneChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeDidSendChanges - An event that indicates a finished send operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/didSendChanges
CKSyncEngineEventTypeDidSendChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeFetchedDatabaseChanges - An event indicating there are fetched database changes to process.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/fetchedDatabaseChanges
CKSyncEngineEventTypeFetchedDatabaseChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeFetchedRecordZoneChanges - An event indicating there are fetched record zone changes to process.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/fetchedRecordZoneChanges
CKSyncEngineEventTypeFetchedRecordZoneChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeSentDatabaseChanges - An event indicating a sent batch of database changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/sentDatabaseChanges
CKSyncEngineEventTypeSentDatabaseChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeSentRecordZoneChanges - An event indicating a sent batch of record zone changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/sentRecordZoneChanges
CKSyncEngineEventTypeSentRecordZoneChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeStateUpdate - An event indicating an update to the sync engine’s state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/stateUpdate
CKSyncEngineEventTypeStateUpdate CKSyncEngineEventType = 0
// CKSyncEngineEventTypeWillFetchChanges - An event indicating an imminent database fetch.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/willFetchChanges
CKSyncEngineEventTypeWillFetchChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeWillFetchRecordZoneChanges - An event indicating an imminent fetch of changes in a record zone.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/willFetchRecordZoneChanges
CKSyncEngineEventTypeWillFetchRecordZoneChanges CKSyncEngineEventType = 0
// CKSyncEngineEventTypeWillSendChanges - An event indicating an imminent send of local changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEventType/willSendChanges
CKSyncEngineEventTypeWillSendChanges CKSyncEngineEventType = 0
)

// CKSyncEnginePendingDatabaseChangeType - Describes the type of a pending database change.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChangeType
type CKSyncEnginePendingDatabaseChangeType uint

// CKSyncEnginePendingRecordZoneChangeType - Describes a type of modification a record zone change makes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChangeType
type CKSyncEnginePendingRecordZoneChangeType uint

// CKSyncEngineSyncReason - Describes the reason for a sync operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSyncReason
type CKSyncEngineSyncReason uint

const (
// CKSyncEngineSyncReasonManual - A manual sync operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSyncReason/manual
CKSyncEngineSyncReasonManual CKSyncEngineSyncReason = 0
// CKSyncEngineSyncReasonScheduled - A scheduled sync operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSyncReason/scheduled
CKSyncEngineSyncReasonScheduled CKSyncEngineSyncReason = 0
)


