// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

// Enum types and constants
// NSFileProviderContentPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderContentPolicy
type FileProviderContentPolicy uint

// NSFileProviderCreateItemOptions - Options for creating items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions
type FileProviderCreateItemOptions uint

const (
	// FileProviderCreateItemMayAlreadyExist - An option indicating that the item may already exist in your remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderCreateItemOptions/mayAlreadyExist
	FileProviderCreateItemMayAlreadyExist FileProviderCreateItemOptions = 0
)

// NSFileProviderDeleteItemOptions - Options for deleting items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDeleteItemOptions
type FileProviderDeleteItemOptions uint

// NSFileProviderDomainTestingModes - Modes that modify the system’s behavior while testing.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/TestingModes-swift.struct
type FileProviderDomainTestingModes uint

// NSFileProviderErrorCode - The error codes for the File Provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code
type FileProviderErrorCode uint

const (
	// FileProviderErrorApplicationExtensionNotFound - An error indicating that there isn’t an app extension within the app bundle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/applicationExtensionNotFound
	FileProviderErrorApplicationExtensionNotFound FileProviderErrorCode = 0
	// FileProviderErrorCannotSynchronize - An error indicating a failed sync attempt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/cannotSynchronize
	FileProviderErrorCannotSynchronize FileProviderErrorCode = 0
	// FileProviderErrorDeletionRejected - An error indicating a failed deletion action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/deletionRejected
	FileProviderErrorDeletionRejected FileProviderErrorCode = 0
	// FileProviderErrorDirectoryNotEmpty - An error indicating an attempt to nonrecursively delete a directory that isn’t empty.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/directoryNotEmpty
	FileProviderErrorDirectoryNotEmpty FileProviderErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/domainDisabled
	FileProviderErrorDomainDisabled FileProviderErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/excludedFromSync
	FileProviderErrorExcludedFromSync FileProviderErrorCode = 0
	// FileProviderErrorFilenameCollision - An error indicating that an item with the same name already exists in the same directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/filenameCollision
	FileProviderErrorFilenameCollision FileProviderErrorCode = 0
	// FileProviderErrorInsufficientQuota - An error indicating that the File Provider extension can’t upload the item because it would push the account over its quota.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/insufficientQuota
	FileProviderErrorInsufficientQuota FileProviderErrorCode = 0
	// FileProviderErrorLocalVersionConflictingWithServer - Returned by createItemBasedOnTemplate or modifyItem if the provider does not wish to sync the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/localVersionConflictingWithServer
	FileProviderErrorLocalVersionConflictingWithServer FileProviderErrorCode = 0
	// FileProviderErrorNewerExtensionVersionFound - An error indicating that the registered provider in the system is a newer version than the one the app uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/newerExtensionVersionFound
	FileProviderErrorNewerExtensionVersionFound FileProviderErrorCode = 0
	// FileProviderErrorNoSuchItem - An error indicating that the specified item doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
	FileProviderErrorNoSuchItem FileProviderErrorCode = 0
	// FileProviderErrorNonEvictable - An error indicating that the File Provider extension can’t evict an item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictable
	FileProviderErrorNonEvictable FileProviderErrorCode = 0
	// FileProviderErrorNonEvictableChildren - An error indicating that the File Provider extension can’t evict a directory because it contains nonevictable items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictableChildren
	FileProviderErrorNonEvictableChildren FileProviderErrorCode = 0
	// FileProviderErrorNotAuthenticated - An error indicating that you can’t verify the user’s credentials.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
	FileProviderErrorNotAuthenticated FileProviderErrorCode = 0
	// FileProviderErrorOlderExtensionVersionRunning - An error indicating that the registered provider in the system is an older version than the one the app uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/olderExtensionVersionRunning
	FileProviderErrorOlderExtensionVersionRunning FileProviderErrorCode = 0
	// FileProviderErrorProviderDomainNotFound - An error indicating that there isn’t a registered domain for the corresponding identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerDomainNotFound
	FileProviderErrorProviderDomainNotFound FileProviderErrorCode = 0
	// FileProviderErrorProviderDomainTemporarilyUnavailable - An error indicating that the system is unable to service requests for the domain temporarily, and you can try again later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerDomainTemporarilyUnavailable
	FileProviderErrorProviderDomainTemporarilyUnavailable FileProviderErrorCode = 0
	// FileProviderErrorProviderNotFound - An error indicating that the File Provider manager can’t find the specified provider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerNotFound
	FileProviderErrorProviderNotFound FileProviderErrorCode = 0
	// FileProviderErrorProviderTranslocated - An error indicating the File Provider extension is in a disabled state due to Gatekeeper’s restrictions for apps from outside the App Store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/providerTranslocated
	FileProviderErrorProviderTranslocated FileProviderErrorCode = 0
	// FileProviderErrorServerUnreachable - An error indicating that the File Provider extension can’t reach the remote server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
	FileProviderErrorServerUnreachable FileProviderErrorCode = 0
	// FileProviderErrorSyncAnchorExpired - An error indicating that the sync anchor is too old, and that the system must restart the sync operation from the beginning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/syncAnchorExpired
	FileProviderErrorSyncAnchorExpired FileProviderErrorCode = 0
	// FileProviderErrorUnsyncedEdits - An error indicating that the item contains unsynced changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/unsyncedEdits
	FileProviderErrorUnsyncedEdits FileProviderErrorCode = 0
	// FileProviderErrorVersionNoLongerAvailable - An error indicating that the specified version is no longer available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/versionNoLongerAvailable
	FileProviderErrorVersionNoLongerAvailable FileProviderErrorCode = 0
	// FileProviderErrorPageExpired - An error indicating that the page is too old, and that the system must restart the enumeration operation from the beginning.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderErrorCode/NSFileProviderErrorPageExpired
	FileProviderErrorPageExpired FileProviderErrorCode = 0
)

// NSFileProviderFileSystemFlags - Flags that define an item’s on-disk properties and its appearance in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderFileSystemFlags
type FileProviderFileSystemFlags uint

// NSFileProviderItemCapabilities - An item’s capabilities, which define the actions that the user can perform in the document browser.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities
type FileProviderItemCapabilities uint

const (
	// FileProviderItemCapabilitiesAllowsAddingSubItems - A value indicating that the user can add subitems.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsAddingSubItems
	FileProviderItemCapabilitiesAllowsAddingSubItems FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsAll - A convenience value for enabling all capabilities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsAll
	FileProviderItemCapabilitiesAllowsAll FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsContentEnumerating - A value indicating that the item’s contents can be enumerated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsContentEnumerating
	FileProviderItemCapabilitiesAllowsContentEnumerating FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsDeleting - A value indicating that the item can be deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsDeleting
	FileProviderItemCapabilitiesAllowsDeleting FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsEvicting - A value indicating that the system can delete the local copy of the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsEvicting
	FileProviderItemCapabilitiesAllowsEvicting FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsExcludingFromSync - A value indicating that the user can exclude the item from sync operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsExcludingFromSync
	FileProviderItemCapabilitiesAllowsExcludingFromSync FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsReading - A value indicating that the value can be read from.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsReading
	FileProviderItemCapabilitiesAllowsReading FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsRenaming - A value indicating that the item can be renamed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsRenaming
	FileProviderItemCapabilitiesAllowsRenaming FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsReparenting - A value indicating that the item can be moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsReparenting
	FileProviderItemCapabilitiesAllowsReparenting FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsTrashing - A value indicating that the item can be moved to the trash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsTrashing
	FileProviderItemCapabilitiesAllowsTrashing FileProviderItemCapabilities = 0
	// FileProviderItemCapabilitiesAllowsWriting - A value indicating that the item can be written to.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemCapabilities/allowsWriting
	FileProviderItemCapabilitiesAllowsWriting FileProviderItemCapabilities = 0
)

// NSFileProviderItemFields - Fields that specify which of the item’s properties have changed.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemFields
type FileProviderItemFields uint

// NSFileProviderKnownFolders - Constants that identify known folders.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders
type FileProviderKnownFolders uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders/desktop
	FileProviderDesktop FileProviderKnownFolders = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolders/documents
	FileProviderDocuments FileProviderKnownFolders = 0
)

// NSFileProviderManagerDisconnectionOptions - Options for disconnecting a domain from the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DisconnectionOptions
type FileProviderManagerDisconnectionOptions uint

// NSFileProviderDomainRemovalMode - A mode indicating how the system handles user data when removing a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode
type FileProviderDomainRemovalMode uint

// NSFileProviderModifyItemOptions - Options for modifying items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions
type FileProviderModifyItemOptions uint

const (
	// FileProviderModifyItemMayAlreadyExist - An option that indicates the changes may already exist in your remote storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderModifyItemOptions/mayAlreadyExist
	FileProviderModifyItemMayAlreadyExist FileProviderModifyItemOptions = 0
)

// NSFileProviderTestingOperationSide - The location where the operation takes place.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationSide
type FileProviderTestingOperationSide uint

// NSFileProviderTestingOperationType - The action that an operation performs.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperationType
type FileProviderTestingOperationType uint

// NSFileProviderVolumeUnsupportedReason - Constants that describe why an external volume might not be eligible for storing a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason
type FileProviderVolumeUnsupportedReason uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/NSFileProviderVolumeUnsupportedReasonNone
	FileProviderVolumeUnsupportedReasonNone FileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/network
	FileProviderVolumeUnsupportedReasonNetwork FileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/nonAPFS
	FileProviderVolumeUnsupportedReasonNonAPFS FileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/nonEncrypted
	FileProviderVolumeUnsupportedReasonNonEncrypted FileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/quarantined
	FileProviderVolumeUnsupportedReasonQuarantined FileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/readOnly
	FileProviderVolumeUnsupportedReasonReadOnly FileProviderVolumeUnsupportedReason = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderVolumeUnsupportedReason/unknown
	FileProviderVolumeUnsupportedReasonUnknown FileProviderVolumeUnsupportedReason = 0
)


