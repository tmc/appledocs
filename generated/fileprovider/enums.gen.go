// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

// Enum types and constants
// NSFileProviderContentPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy
type NSFileProviderContentPolicy uint

const (
	NSFileProviderContentPolicyInherited NSFileProviderContentPolicy = 0
	NSFileProviderContentPolicyDownloadLazily NSFileProviderContentPolicy = 1
	NSFileProviderContentPolicyDownloadLazilyAndEvictOnRemoteUpdate NSFileProviderContentPolicy = 2
	NSFileProviderContentPolicyDownloadEagerlyAndKeepDownloaded NSFileProviderContentPolicy = 3
)

// NSFileProviderCreateItemOptions - Options for creating items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions
type NSFileProviderCreateItemOptions uint

const (
	// NSFileProviderCreateItemMayAlreadyExist - An option indicating that the item may already exist in your remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions/mayAlreadyExist
	NSFileProviderCreateItemMayAlreadyExist NSFileProviderCreateItemOptions = 1
)

// NSFileProviderDeleteItemOptions - Options for deleting items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDeleteItemOptions
type NSFileProviderDeleteItemOptions uint

const (
	NSFileProviderDeleteItemRecursive NSFileProviderDeleteItemOptions = 1
)

// NSFileProviderDomainTestingModes - Modes that modify the system’s behavior while testing.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/TestingModes-swift.struct
type NSFileProviderDomainTestingModes uint

const (
	NSFileProviderDomainTestingModeAlwaysEnabled NSFileProviderDomainTestingModes = 1
	NSFileProviderDomainTestingModeInteractive NSFileProviderDomainTestingModes = 2
)

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
	// NSFileProviderErrorNoSuchItem - An error indicating that the specified item doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
	NSFileProviderErrorNoSuchItem NSFileProviderErrorCode = -1005
	// NSFileProviderErrorNonEvictable - An error indicating that the File Provider extension can’t evict an item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictable
	NSFileProviderErrorNonEvictable NSFileProviderErrorCode = -995
	// NSFileProviderErrorNonEvictableChildren - An error indicating that the File Provider extension can’t evict a directory because it contains nonevictable items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictableChildren
	NSFileProviderErrorNonEvictableChildren NSFileProviderErrorCode = -997
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

// NSFileProviderFileSystemFlags - Flags that define an item’s on-disk properties and its appearance in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags
type NSFileProviderFileSystemFlags uint

const (
	NSFileProviderFileSystemUserExecutable NSFileProviderFileSystemFlags = 1
	NSFileProviderFileSystemUserReadable NSFileProviderFileSystemFlags = 2
	NSFileProviderFileSystemUserWritable NSFileProviderFileSystemFlags = 4
	NSFileProviderFileSystemHidden NSFileProviderFileSystemFlags = 8
	NSFileProviderFileSystemPathExtensionHidden NSFileProviderFileSystemFlags = 16
)

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

// NSFileProviderItemFields - Fields that specify which of the item’s properties have changed.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields
type NSFileProviderItemFields uint

const (
	NSFileProviderItemContents NSFileProviderItemFields = 1
	NSFileProviderItemFilename NSFileProviderItemFields = 2
	NSFileProviderItemParentItemIdentifier NSFileProviderItemFields = 4
	NSFileProviderItemLastUsedDate NSFileProviderItemFields = 8
	NSFileProviderItemTagData NSFileProviderItemFields = 16
	NSFileProviderItemFavoriteRank NSFileProviderItemFields = 32
	NSFileProviderItemCreationDate NSFileProviderItemFields = 64
	NSFileProviderItemContentModificationDate NSFileProviderItemFields = 128
	NSFileProviderItemFileSystemFlags NSFileProviderItemFields = 256
	NSFileProviderItemExtendedAttributes NSFileProviderItemFields = 512
	NSFileProviderItemTypeAndCreator NSFileProviderItemFields = 513
)

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

// NSFileProviderModifyItemOptions - Options for modifying items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions
type NSFileProviderModifyItemOptions uint

const (
	NSFileProviderModifyItemMayAlreadyExist NSFileProviderModifyItemOptions = 1
	NSFileProviderModifyItemFailOnConflict NSFileProviderModifyItemOptions = 2
)

// NSFileProviderTestingOperationSide - The location where the operation takes place.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationSide
type NSFileProviderTestingOperationSide uint

const (
	NSFileProviderTestingOperationSideDisk NSFileProviderTestingOperationSide = 0
	NSFileProviderTestingOperationSideFileProvider NSFileProviderTestingOperationSide = 1
)

// NSFileProviderTestingOperationType - The action that an operation performs.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType
type NSFileProviderTestingOperationType uint

const (
	NSFileProviderTestingOperationTypeIngestion NSFileProviderTestingOperationType = 0
	NSFileProviderTestingOperationTypeLookup NSFileProviderTestingOperationType = 1
	NSFileProviderTestingOperationTypeCreation NSFileProviderTestingOperationType = 2
	NSFileProviderTestingOperationTypeModification NSFileProviderTestingOperationType = 3
	NSFileProviderTestingOperationTypeDeletion NSFileProviderTestingOperationType = 4
	NSFileProviderTestingOperationTypeContentFetch NSFileProviderTestingOperationType = 5
	NSFileProviderTestingOperationTypeChildrenEnumeration NSFileProviderTestingOperationType = 6
	NSFileProviderTestingOperationTypeCollisionResolution NSFileProviderTestingOperationType = 7
)

// NSFileProviderVolumeUnsupportedReason - Constants that describe why an external volume might not be eligible for storing a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason
type NSFileProviderVolumeUnsupportedReason uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/NSFileProviderVolumeUnsupportedReasonNone
	NSFileProviderVolumeUnsupportedReasonNone NSFileProviderVolumeUnsupportedReason = 0
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
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/quarantined
	NSFileProviderVolumeUnsupportedReasonQuarantined NSFileProviderVolumeUnsupportedReason = 32
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/readOnly
	NSFileProviderVolumeUnsupportedReasonReadOnly NSFileProviderVolumeUnsupportedReason = 8
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/unknown
	NSFileProviderVolumeUnsupportedReasonUnknown NSFileProviderVolumeUnsupportedReason = 1
)


