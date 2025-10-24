// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

/* debug [enums.gen.go]: Generating 17 enums for FileProvider */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NSFileProviderDomainTestingModes (2 cases) */
// NSFileProviderDomainTestingModes - Modes that modify the system’s behavior while testing.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/TestingModes-swift.struct
type NSFileProviderDomainTestingModes uint

const (
	// NSFileProviderDomainTestingModeAlwaysEnabled - A testing mode that automatically enables the domain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/TestingModes-swift.struct/alwaysEnabled
	NSFileProviderDomainTestingModeAlwaysEnabled NSFileProviderDomainTestingModes = 1
	// NSFileProviderDomainTestingModeInteractive - A testing mode where the extension can deterministically test asynchronous operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/TestingModes-swift.struct/interactive
	NSFileProviderDomainTestingModeInteractive NSFileProviderDomainTestingModes = 2
)

/* debug [enums.gen.go]: Processing enum NSFileProviderErrorCode (24 cases) */
// NSFileProviderErrorCode - The error codes for the File Provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code
type NSFileProviderErrorCode int

const (
	// NSFileProviderErrorApplicationExtensionNotFound - An error indicating that there isn’t an app extension within the app bundle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/applicationExtensionNotFound
	NSFileProviderErrorApplicationExtensionNotFound NSFileProviderErrorCode = -989
	// NSFileProviderErrorCannotSynchronize - An error indicating a failed sync attempt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/cannotSynchronize
	NSFileProviderErrorCannotSynchronize NSFileProviderErrorCode = -998
	// NSFileProviderErrorDeletionRejected - An error indicating a failed deletion action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/deletionRejected
	NSFileProviderErrorDeletionRejected NSFileProviderErrorCode = -1004
	// NSFileProviderErrorDirectoryNotEmpty - An error indicating an attempt to nonrecursively delete a directory that isn’t empty.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/directoryNotEmpty
	NSFileProviderErrorDirectoryNotEmpty NSFileProviderErrorCode = -1003
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/domainDisabled
	NSFileProviderErrorDomainDisabled NSFileProviderErrorCode = -992
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/excludedFromSync
	NSFileProviderErrorExcludedFromSync NSFileProviderErrorCode = -993
	// NSFileProviderErrorFilenameCollision - An error indicating that an item with the same name already exists in the same directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/filenameCollision
	NSFileProviderErrorFilenameCollision NSFileProviderErrorCode = -1001
	// NSFileProviderErrorInsufficientQuota - An error indicating that the File Provider extension can’t upload the item because it would push the account over its quota.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/insufficientQuota
	NSFileProviderErrorInsufficientQuota NSFileProviderErrorCode = -1003
	// NSFileProviderErrorLocalVersionConflictingWithServer - Returned by createItemBasedOnTemplate or modifyItem if the provider does not wish to sync the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/localVersionConflictingWithServer
	NSFileProviderErrorLocalVersionConflictingWithServer NSFileProviderErrorCode = -988
	// NSFileProviderErrorNewerExtensionVersionFound - An error indicating that the registered provider in the system is a newer version than the one the app uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/newerExtensionVersionFound
	NSFileProviderErrorNewerExtensionVersionFound NSFileProviderErrorCode = -999
	// NSFileProviderErrorNonEvictable - An error indicating that the File Provider extension can’t evict an item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictable
	NSFileProviderErrorNonEvictable NSFileProviderErrorCode = -995
	// NSFileProviderErrorNonEvictableChildren - An error indicating that the File Provider extension can’t evict a directory because it contains nonevictable items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictableChildren
	NSFileProviderErrorNonEvictableChildren NSFileProviderErrorCode = -997
	// NSFileProviderErrorNoSuchItem - An error indicating that the specified item doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
	NSFileProviderErrorNoSuchItem NSFileProviderErrorCode = -1005
	// NSFileProviderErrorNotAuthenticated - An error indicating that you can’t verify the user’s credentials.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
	NSFileProviderErrorNotAuthenticated NSFileProviderErrorCode = -1000
	// NSFileProviderErrorOlderExtensionVersionRunning - An error indicating that the registered provider in the system is an older version than the one the app uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/olderExtensionVersionRunning
	NSFileProviderErrorOlderExtensionVersionRunning NSFileProviderErrorCode = -1000
	// NSFileProviderErrorProviderDomainNotFound - An error indicating that there isn’t a registered domain for the corresponding identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerDomainNotFound
	NSFileProviderErrorProviderDomainNotFound NSFileProviderErrorCode = -990
	// NSFileProviderErrorProviderDomainTemporarilyUnavailable - An error indicating that the system is unable to service requests for the domain temporarily, and you can try again later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerDomainTemporarilyUnavailable
	NSFileProviderErrorProviderDomainTemporarilyUnavailable NSFileProviderErrorCode = -991
	// NSFileProviderErrorProviderNotFound - An error indicating that the File Provider manager can’t find the specified provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerNotFound
	NSFileProviderErrorProviderNotFound NSFileProviderErrorCode = -1002
	// NSFileProviderErrorProviderTranslocated - An error indicating the File Provider extension is in a disabled state due to Gatekeeper’s restrictions for apps from outside the App Store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerTranslocated
	NSFileProviderErrorProviderTranslocated NSFileProviderErrorCode = -1001
	// NSFileProviderErrorServerUnreachable - An error indicating that the File Provider extension can’t reach the remote server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
	NSFileProviderErrorServerUnreachable NSFileProviderErrorCode = -1004
	// NSFileProviderErrorSyncAnchorExpired - An error indicating that the sync anchor is too old, and that the system must restart the sync operation from the beginning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/syncAnchorExpired
	NSFileProviderErrorSyncAnchorExpired NSFileProviderErrorCode = -1002
	// NSFileProviderErrorUnsyncedEdits - An error indicating that the item contains unsynced changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/unsyncedEdits
	NSFileProviderErrorUnsyncedEdits NSFileProviderErrorCode = -996
	// NSFileProviderErrorVersionNoLongerAvailable - An error indicating that the specified version is no longer available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/versionNoLongerAvailable
	NSFileProviderErrorVersionNoLongerAvailable NSFileProviderErrorCode = -994
	// NSFileProviderErrorPageExpired - An error indicating that the page is too old, and that the system must restart the enumeration operation from the beginning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderErrorCode/NSFileProviderErrorPageExpired
	NSFileProviderErrorPageExpired NSFileProviderErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum NSFileProviderItemCapabilities (11 cases) */
// NSFileProviderItemCapabilities - An item’s capabilities, which define the actions that the user can perform in the document browser.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities
type NSFileProviderItemCapabilities uint

const (
	// NSFileProviderItemCapabilitiesAllowsAddingSubItems - A value indicating that the user can add subitems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsAddingSubItems
	NSFileProviderItemCapabilitiesAllowsAddingSubItems NSFileProviderItemCapabilities = 0
	// NSFileProviderItemCapabilitiesAllowsAll - A convenience value for enabling all capabilities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsAll
	NSFileProviderItemCapabilitiesAllowsAll NSFileProviderItemCapabilities = 35
	// NSFileProviderItemCapabilitiesAllowsContentEnumerating - A value indicating that the item’s contents can be enumerated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsContentEnumerating
	NSFileProviderItemCapabilitiesAllowsContentEnumerating NSFileProviderItemCapabilities = 0
	// NSFileProviderItemCapabilitiesAllowsDeleting - A value indicating that the item can be deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsDeleting
	NSFileProviderItemCapabilitiesAllowsDeleting NSFileProviderItemCapabilities = 32
	// NSFileProviderItemCapabilitiesAllowsEvicting - A value indicating that the system can delete the local copy of the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsEvicting
	NSFileProviderItemCapabilitiesAllowsEvicting NSFileProviderItemCapabilities = 33
	// NSFileProviderItemCapabilitiesAllowsExcludingFromSync - A value indicating that the user can exclude the item from sync operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsExcludingFromSync
	NSFileProviderItemCapabilitiesAllowsExcludingFromSync NSFileProviderItemCapabilities = 34
	// NSFileProviderItemCapabilitiesAllowsReading - A value indicating that the value can be read from.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsReading
	NSFileProviderItemCapabilitiesAllowsReading NSFileProviderItemCapabilities = 36
	// NSFileProviderItemCapabilitiesAllowsRenaming - A value indicating that the item can be renamed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsRenaming
	NSFileProviderItemCapabilitiesAllowsRenaming NSFileProviderItemCapabilities = 8
	// NSFileProviderItemCapabilitiesAllowsReparenting - A value indicating that the item can be moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsReparenting
	NSFileProviderItemCapabilitiesAllowsReparenting NSFileProviderItemCapabilities = 4
	// NSFileProviderItemCapabilitiesAllowsTrashing - A value indicating that the item can be moved to the trash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsTrashing
	NSFileProviderItemCapabilitiesAllowsTrashing NSFileProviderItemCapabilities = 16
	// NSFileProviderItemCapabilitiesAllowsWriting - A value indicating that the item can be written to.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsWriting
	NSFileProviderItemCapabilitiesAllowsWriting NSFileProviderItemCapabilities = 2
)

/* debug [enums.gen.go]: Processing enum NSFileProviderManagerDisconnectionOptions (1 cases) */
// NSFileProviderManagerDisconnectionOptions - Options for disconnecting a domain from the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DisconnectionOptions
type NSFileProviderManagerDisconnectionOptions uint

const (
	// NSFileProviderManagerDisconnectionOptionsTemporary - A temporary disconnection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DisconnectionOptions/temporary
	NSFileProviderManagerDisconnectionOptionsTemporary NSFileProviderManagerDisconnectionOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderDomainRemovalMode (3 cases) */
// NSFileProviderDomainRemovalMode - A mode indicating how the system handles user data when removing a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode
type NSFileProviderDomainRemovalMode uint

const (
	// NSFileProviderDomainRemovalModePreserveDirtyUserData - Deletes the domain but keeps any items with unsynced, local changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode/preserveDirtyUserData
	NSFileProviderDomainRemovalModePreserveDirtyUserData NSFileProviderDomainRemovalMode = 1
	// NSFileProviderDomainRemovalModePreserveDownloadedUserData - Deletes the domain, but keeps the downloaded user data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode/preserveDownloadedUserData
	NSFileProviderDomainRemovalModePreserveDownloadedUserData NSFileProviderDomainRemovalMode = 2
	// NSFileProviderDomainRemovalModeRemoveAll - Deletes all items in the domain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode/removeAll
	NSFileProviderDomainRemovalModeRemoveAll NSFileProviderDomainRemovalMode = 0
)

/* debug [enums.gen.go]: Processing enum NSFileProviderVolumeUnsupportedReason (7 cases) */
// NSFileProviderVolumeUnsupportedReason - Constants that describe why an external volume might not be eligible for storing a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason
type NSFileProviderVolumeUnsupportedReason uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/network
	NSFileProviderVolumeUnsupportedReasonNetwork NSFileProviderVolumeUnsupportedReason = 16
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/nonAPFS
	NSFileProviderVolumeUnsupportedReasonNonAPFS NSFileProviderVolumeUnsupportedReason = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/nonEncrypted
	NSFileProviderVolumeUnsupportedReasonNonEncrypted NSFileProviderVolumeUnsupportedReason = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/NSFileProviderVolumeUnsupportedReasonNone
	NSFileProviderVolumeUnsupportedReasonNone NSFileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/quarantined
	NSFileProviderVolumeUnsupportedReasonQuarantined NSFileProviderVolumeUnsupportedReason = 32
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/readOnly
	NSFileProviderVolumeUnsupportedReasonReadOnly NSFileProviderVolumeUnsupportedReason = 8
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/unknown
	NSFileProviderVolumeUnsupportedReasonUnknown NSFileProviderVolumeUnsupportedReason = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderContentPolicy (4 cases) */
// NSFileProviderContentPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy
type NSFileProviderContentPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy/downloadEagerlyAndKeepDownloaded
	NSFileProviderContentPolicyDownloadEagerlyAndKeepDownloaded NSFileProviderContentPolicy = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy/downloadLazily
	NSFileProviderContentPolicyDownloadLazily NSFileProviderContentPolicy = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy/downloadLazilyAndEvictOnRemoteUpdate
	NSFileProviderContentPolicyDownloadLazilyAndEvictOnRemoteUpdate NSFileProviderContentPolicy = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy/inherited
	NSFileProviderContentPolicyInherited NSFileProviderContentPolicy = 0
)

/* debug [enums.gen.go]: Processing enum NSFileProviderCreateItemOptions (2 cases) */
// NSFileProviderCreateItemOptions - Options for creating items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions
type NSFileProviderCreateItemOptions uint

const (
	// NSFileProviderCreateItemDeletionConflicted - A value indicating a conflict for a deleted item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions/deletionConflicted
	NSFileProviderCreateItemDeletionConflicted NSFileProviderCreateItemOptions = 2
	// NSFileProviderCreateItemMayAlreadyExist - An option indicating that the item may already exist in your remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions/mayAlreadyExist
	NSFileProviderCreateItemMayAlreadyExist NSFileProviderCreateItemOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderDeleteItemOptions (1 cases) */
// NSFileProviderDeleteItemOptions - Options for deleting items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDeleteItemOptions
type NSFileProviderDeleteItemOptions uint

const (
	// NSFileProviderDeleteItemRecursive - A value indicating that the delete operation removes the item and all of its children.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDeleteItemOptions/recursive
	NSFileProviderDeleteItemRecursive NSFileProviderDeleteItemOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderFetchContentsOptions (1 cases) */
// NSFileProviderFetchContentsOptions - Options for fetching a range of data from a file.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFetchContentsOptions
type NSFileProviderFetchContentsOptions uint

const (
	// NSFileProviderFetchContentsOptionsStrictVersioning - An option that indicates the system requires an exact match of the requested item’s version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFetchContentsOptions/strictVersioning
	NSFileProviderFetchContentsOptionsStrictVersioning NSFileProviderFetchContentsOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderFileSystemFlags (5 cases) */
// NSFileProviderFileSystemFlags - Flags that define an item’s on-disk properties and its appearance in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags
type NSFileProviderFileSystemFlags uint

const (
	// NSFileProviderFileSystemHidden - By default, the system hides the item when the user views the file system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags/hidden
	NSFileProviderFileSystemHidden NSFileProviderFileSystemFlags = 8
	// NSFileProviderFileSystemPathExtensionHidden - By default, the system hides the item’s extension when showing its filename.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags/pathExtensionHidden
	NSFileProviderFileSystemPathExtensionHidden NSFileProviderFileSystemFlags = 16
	// NSFileProviderFileSystemUserExecutable - The user can execute the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags/userExecutable
	NSFileProviderFileSystemUserExecutable NSFileProviderFileSystemFlags = 1
	// NSFileProviderFileSystemUserReadable - The user can read the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags/userReadable
	NSFileProviderFileSystemUserReadable NSFileProviderFileSystemFlags = 2
	// NSFileProviderFileSystemUserWritable - The user can modify the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags/userWritable
	NSFileProviderFileSystemUserWritable NSFileProviderFileSystemFlags = 4
)

/* debug [enums.gen.go]: Processing enum NSFileProviderItemFields (11 cases) */
// NSFileProviderItemFields - Fields that specify which of the item’s properties have changed.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields
type NSFileProviderItemFields uint

const (
	// NSFileProviderItemContentModificationDate - The item’s modification date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/contentModificationDate
	NSFileProviderItemContentModificationDate NSFileProviderItemFields = 128
	// NSFileProviderItemContents - The item’s content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/contents
	NSFileProviderItemContents NSFileProviderItemFields = 1
	// NSFileProviderItemCreationDate - The item’s creation date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/creationDate
	NSFileProviderItemCreationDate NSFileProviderItemFields = 64
	// NSFileProviderItemExtendedAttributes - The item’s extended attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/extendedAttributes
	NSFileProviderItemExtendedAttributes NSFileProviderItemFields = 512
	// NSFileProviderItemFavoriteRank - The item’s favorite rank.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/favoriteRank
	NSFileProviderItemFavoriteRank NSFileProviderItemFields = 32
	// NSFileProviderItemFilename - The item’s filename.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/filename
	NSFileProviderItemFilename NSFileProviderItemFields = 2
	// NSFileProviderItemFileSystemFlags - The flags describing the item’s on-disk representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/fileSystemFlags
	NSFileProviderItemFileSystemFlags NSFileProviderItemFields = 256
	// NSFileProviderItemLastUsedDate - The date the item was last used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/lastUsedDate
	NSFileProviderItemLastUsedDate NSFileProviderItemFields = 8
	// NSFileProviderItemParentItemIdentifier - The identity of the directory that contains the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/parentItemIdentifier
	NSFileProviderItemParentItemIdentifier NSFileProviderItemFields = 4
	// NSFileProviderItemTagData - The tags for the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/tagData
	NSFileProviderItemTagData NSFileProviderItemFields = 16
	// NSFileProviderItemTypeAndCreator - The file type and creator codes for the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields/typeAndCreator
	NSFileProviderItemTypeAndCreator NSFileProviderItemFields = 513
)

/* debug [enums.gen.go]: Processing enum NSFileProviderKnownFolders (2 cases) */
// NSFileProviderKnownFolders - Constants that identify known folders.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders
type NSFileProviderKnownFolders uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders/desktop
	NSFileProviderDesktop NSFileProviderKnownFolders = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders/documents
	NSFileProviderDocuments NSFileProviderKnownFolders = 2
)

/* debug [enums.gen.go]: Processing enum NSFileProviderMaterializationFlags (1 cases) */
// NSFileProviderMaterializationFlags - Flags that provides additional information about the provided content.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderMaterializationFlags
type NSFileProviderMaterializationFlags uint

const (
	// NSFileProviderMaterializationFlagsKnownSparseRanges - A flag indicating that the system should consider the file fully materialized, even if it’s a sparse file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderMaterializationFlags/knownSparseRanges
	NSFileProviderMaterializationFlagsKnownSparseRanges NSFileProviderMaterializationFlags = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderModifyItemOptions (3 cases) */
// NSFileProviderModifyItemOptions - Options for modifying items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions
type NSFileProviderModifyItemOptions uint

const (
	// NSFileProviderModifyItemFailOnConflict - An option to fail an upload in the event of a version conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions/failOnConflict
	NSFileProviderModifyItemFailOnConflict NSFileProviderModifyItemOptions = 2
	// NSFileProviderModifyItemIsImmediateUploadRequestByPresentingApplication - An option to require the upload to complete before calling the completion handler.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions/isImmediateUploadRequestByPresentingApplication
	NSFileProviderModifyItemIsImmediateUploadRequestByPresentingApplication NSFileProviderModifyItemOptions = 0
	// NSFileProviderModifyItemMayAlreadyExist - An option that indicates the changes may already exist in your remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions/mayAlreadyExist
	NSFileProviderModifyItemMayAlreadyExist NSFileProviderModifyItemOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderTestingOperationSide (2 cases) */
// NSFileProviderTestingOperationSide - The location where the operation takes place.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationSide
type NSFileProviderTestingOperationSide uint

const (
	// NSFileProviderTestingOperationSideDisk - The File Provider extension’s local storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationSide/disk
	NSFileProviderTestingOperationSideDisk NSFileProviderTestingOperationSide = 0
	// NSFileProviderTestingOperationSideFileProvider - The File Provider extension’s remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationSide/fileProvider
	NSFileProviderTestingOperationSideFileProvider NSFileProviderTestingOperationSide = 1
)

/* debug [enums.gen.go]: Processing enum NSFileProviderTestingOperationType (8 cases) */
// NSFileProviderTestingOperationType - The action that an operation performs.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType
type NSFileProviderTestingOperationType uint

const (
	// NSFileProviderTestingOperationTypeChildrenEnumeration - Lists an item’s content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/childrenEnumeration
	NSFileProviderTestingOperationTypeChildrenEnumeration NSFileProviderTestingOperationType = 6
	// NSFileProviderTestingOperationTypeCollisionResolution - Resolves a collision by renaming the new item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/collisionResolution
	NSFileProviderTestingOperationTypeCollisionResolution NSFileProviderTestingOperationType = 7
	// NSFileProviderTestingOperationTypeContentFetch - Fetches an item’s content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/contentFetch
	NSFileProviderTestingOperationTypeContentFetch NSFileProviderTestingOperationType = 5
	// NSFileProviderTestingOperationTypeCreation - Propagates the creation of a source item to the target location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/creation
	NSFileProviderTestingOperationTypeCreation NSFileProviderTestingOperationType = 2
	// NSFileProviderTestingOperationTypeDeletion - Propagates the deletion of the source item from the target location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/deletion
	NSFileProviderTestingOperationTypeDeletion NSFileProviderTestingOperationType = 4
	// NSFileProviderTestingOperationTypeIngestion - Alerts the system to changes to either the local or remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/ingestion
	NSFileProviderTestingOperationTypeIngestion NSFileProviderTestingOperationType = 0
	// NSFileProviderTestingOperationTypeLookup - Looks up an item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/lookup
	NSFileProviderTestingOperationTypeLookup NSFileProviderTestingOperationType = 1
	// NSFileProviderTestingOperationTypeModification - Propagates a change from the source item to the target location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType/modification
	NSFileProviderTestingOperationTypeModification NSFileProviderTestingOperationType = 3
)


