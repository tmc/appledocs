// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileProviderManager] class.
var (
	FileProviderManagerClass     _FileProviderManagerClass
	FileProviderManagerClassOnce sync.Once
)

func getFileProviderManagerClass() _FileProviderManagerClass {
	FileProviderManagerClassOnce.Do(func() {
		FileProviderManagerClass = _FileProviderManagerClass{objc.GetClass("NSFileProviderManager")}
	})
	return FileProviderManagerClass
}

type _FileProviderManagerClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderManager] class.
type IFileProviderManager interface {
	objectivec.IObject
	ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders IFileProviderKnownFolderLocations, localizedReason appkit.string, completionHandler unsafe.Pointer)
	DisconnectWithReasonOptionsCompletionHandler(localizedReason appkit.string, options FileProviderManagerDisconnectionOptions, completionHandler unsafe.Pointer)
	EnumeratorForMaterializedItems() objc.ID
	EnumeratorForPendingItems() objc.ID
	EvictItemWithIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer)
	GetServiceWithNameItemIdentifierCompletionHandler(serviceName unsafe.Pointer, itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer)
	GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer)
	GlobalProgressForKind(kind unsafe.Pointer) foundation.Progress
	ListAvailableTestingOperationsWithError(error_ unsafe.Pointer) []objc.ID
	ReconnectWithCompletionHandler(completionHandler unsafe.Pointer)
	RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task foundation.IURLSessionTask, identifier IFileProviderItemIdentifier, completion unsafe.Pointer)
	ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer)
	ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders IFileProviderKnownFolders, localizedReason appkit.string, completionHandler unsafe.Pointer)
	RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier IFileProviderItemIdentifier, errorReason foundation.IError, completionHandler unsafe.Pointer)
	RequestDownloadForItemWithIdentifierRequestedRangeCompletionHandler(itemIdentifier IFileProviderItemIdentifier, rangeToMaterialize foundation.IRange, completionHandler unsafe.Pointer)
	RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields IFileProviderItemFields, itemIdentifier IFileProviderItemIdentifier, options FileProviderModifyItemOptions, completionHandler unsafe.Pointer)
	RunTestingOperationsError(operations []objc.ID, error_ unsafe.Pointer) unsafe.Pointer
	SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier IFileProviderItemIdentifier, completion unsafe.Pointer)
	SignalErrorResolvedCompletionHandler(error_ foundation.IError, completionHandler unsafe.Pointer)
	StateDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL
	TemporaryDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL
	WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer)
	WaitForStabilizationWithCompletionHandler(completionHandler unsafe.Pointer)
}

// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager
type FileProviderManager struct {
	objectivec.Object
}

// FileProviderManagerFrom constructs a [FileProviderManager] from an unsafe.Pointer.
//
// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.
func FileProviderManagerFrom(ptr unsafe.Pointer) FileProviderManager {
	return FileProviderManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderManagerClass) Alloc() FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderManagerClass) New() FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderManager) Init() FileProviderManager {
	rv := objc.Send[FileProviderManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderManager) Autorelease() FileProviderManager {
	rv := objc.Send[FileProviderManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderManager creates a new FileProviderManager instance.
func NewFileProviderManager() FileProviderManager {
	return getFileProviderManagerClass().New()
}




// Returns a newly created file provider manager for the specified domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/init(for:)
func NewFileProviderManagerForDomain(domain IFileProviderDomain) FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(getFileProviderManagerClass().class), objc.Sel("managerForDomain:"), domain)
	return rv
}


// Adds a domain to the File Provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/add(_:completionHandler:)
func (fc _FileProviderManagerClass) AddDomainCompletionHandler(domain IFileProviderDomain, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("addDomain:completionHandler:"), domain, completionHandler)
}

// Check if a URL is eligible for storing a domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/checkDomainsCanBeStored:onVolumeAtURL:unsupportedReason:error:
func (fc _FileProviderManagerClass) CheckDomainsCanBeStoredOnVolumeAtURLUnsupportedReasonError(eligible unsafe.Pointer, url foundation.IURL, unsupportedReason IFileProviderVolumeUnsupportedReason, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("checkDomainsCanBeStored:onVolumeAtURL:unsupportedReason:error:"), eligible, url, unsupportedReason, error_)
	return rv
}

// Returns all of the File Provider extension’s domains.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getDomainsWithCompletionHandler(_:)
func (fc _FileProviderManagerClass) GetDomainsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("getDomainsWithCompletionHandler:"), completionHandler)
}

// Returns the identifier and domain for a user-visible URL.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getIdentifierForUserVisibleFile(at:completionHandler:)
func (fc _FileProviderManagerClass) GetIdentifierForUserVisibleFileAtURLCompletionHandler(url foundation.IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("getIdentifierForUserVisibleFileAtURL:completionHandler:"), url, completionHandler)
}

// Creates a new domain that takes ownership of on-disk data that your app previously managed without a file provider.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/import(_:fromDirectoryAt:completionHandler:)
func (fc _FileProviderManagerClass) ImportDomainFromDirectoryAtURLCompletionHandler(domain IFileProviderDomain, url foundation.IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("importDomain:fromDirectoryAtURL:completionHandler:"), domain, url, completionHandler)
}

// Returns a newly created file provider manager for the specified domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/init(for:)
func (fc _FileProviderManagerClass) ManagerForDomain(domain IFileProviderDomain) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("managerForDomain:"), domain)
	return rv
}

// Returns a placeholder URL for a given document URL.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/placeholderURL(for:)
func (fc _FileProviderManagerClass) PlaceholderURLForURL(url foundation.IURL) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(fc.class), objc.Sel("placeholderURLForURL:"), url)
	return rv
}

// Removes a domain from the File Provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/remove(_:completionHandler:)
func (fc _FileProviderManagerClass) RemoveDomainCompletionHandler(domain IFileProviderDomain, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeDomain:completionHandler:"), domain, completionHandler)
}

// Removes a domain from the File Provider extension using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/remove(_:mode:completionHandler:)
func (fc _FileProviderManagerClass) RemoveDomainModeCompletionHandler(domain IFileProviderDomain, mode FileProviderDomainRemovalMode, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeDomain:mode:completionHandler:"), domain, mode, completionHandler)
}

// Removes all domains from the File Provider extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/removeAllDomains(completionHandler:)
func (fc _FileProviderManagerClass) RemoveAllDomainsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeAllDomainsWithCompletionHandler:"), completionHandler)
}

// Writes a document placeholder with the provided metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/writePlaceholder(at:withMetadata:)
func (fc _FileProviderManagerClass) WritePlaceholderAtURLWithMetadataError(placeholderURL foundation.IURL, metadata IFileProviderItem, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, error_)
	return rv
}

// A property that returns the shared file provider manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/default
func (fc _FileProviderManagerClass) DefaultManager() FileProviderManager {
	rv := objc.Send[NSFileProviderManager](objc.ID(fc.class), objc.Sel("defaultManager"))
	return rv
}
// Asks the domain to sync the specified known folders.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/claimKnownFolders(_:localizedReason:completionHandler:)
func (f_ FileProviderManager) ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders IFileProviderKnownFolderLocations, localizedReason appkit.string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("claimKnownFolders:localizedReason:completionHandler:"), knownFolders, localizedReason, completionHandler)
}

// Disconnects the domain from the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/disconnect(reason:options:completionHandler:)
func (f_ FileProviderManager) DisconnectWithReasonOptionsCompletionHandler(localizedReason appkit.string, options FileProviderManagerDisconnectionOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("disconnectWithReason:options:completionHandler:"), localizedReason, options, completionHandler)
}

// Returns an enumerator for all the items the system currently stores on disk.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/enumeratorForMaterializedItems()
func (f_ FileProviderManager) EnumeratorForMaterializedItems() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("enumeratorForMaterializedItems"))
	return rv
}

// Returns an enumerator for the set of pending items.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/enumeratorForPendingItems()
func (f_ FileProviderManager) EnumeratorForPendingItems() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("enumeratorForPendingItems"))
	return rv
}

// Asks the system to remove an item from its cache.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/evictItem(identifier:completionHandler:)
func (f_ FileProviderManager) EvictItemWithIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("evictItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getService(named:for:completionHandler:)
func (f_ FileProviderManager) GetServiceWithNameItemIdentifierCompletionHandler(serviceName unsafe.Pointer, itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getServiceWithName:itemIdentifier:completionHandler:"), serviceName, itemIdentifier, completionHandler)
}

// Returns the user-visible URL for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getUserVisibleURL(for:completionHandler:)
func (f_ FileProviderManager) GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getUserVisibleURLForItemIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Returns a progress object that tracks either the uploading or downloading of items from the File Provider extension’s remote storage.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/globalProgress(for:)
func (f_ FileProviderManager) GlobalProgressForKind(kind unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](f_.ID, objc.Sel("globalProgressForKind:"), kind)
	return rv
}

// Lists all the operations that are ready for scheduling.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/listAvailableTestingOperations()
func (f_ FileProviderManager) ListAvailableTestingOperationsWithError(error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("listAvailableTestingOperationsWithError:"), error_)
	return rv
}

// Reconnects the domain with the extension.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/reconnect(completionHandler:)
func (f_ FileProviderManager) ReconnectWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reconnectWithCompletionHandler:"), completionHandler)
}

// Registers the URL session task responsible for the specified item.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/register(_:forItemWithIdentifier:completionHandler:)
func (f_ FileProviderManager) RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task foundation.IURLSessionTask, identifier IFileProviderItemIdentifier, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("registerURLSessionTask:forItemWithIdentifier:completionHandler:"), task, identifier, completion)
}

// Tells the system to reimport the item and its content recursively.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/reimportItems(below:completionHandler:)
func (f_ FileProviderManager) ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reimportItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Asks the system to stop replicating the specified known folders in the domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/releaseKnownFolders(_:localizedReason:completionHandler:)
func (f_ FileProviderManager) ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders IFileProviderKnownFolders, localizedReason appkit.string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("releaseKnownFolders:localizedReason:completionHandler:"), knownFolders, localizedReason, completionHandler)
}

// Requests a diagnostics collection for use when working directly with Apple to improve sync behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestDiagnosticCollection(for:errorReason:completionHandler:)
func (f_ FileProviderManager) RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier IFileProviderItemIdentifier, errorReason foundation.IError, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestDiagnosticCollectionForItemWithIdentifier:errorReason:completionHandler:"), itemIdentifier, errorReason, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestDownloadForItemWithIdentifier:requestedRange:completionHandler:
func (f_ FileProviderManager) RequestDownloadForItemWithIdentifierRequestedRangeCompletionHandler(itemIdentifier IFileProviderItemIdentifier, rangeToMaterialize foundation.IRange, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestDownloadForItemWithIdentifier:requestedRange:completionHandler:"), itemIdentifier, rangeToMaterialize, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestModification(of:forItemWithIdentifier:options:completionHandler:)
func (f_ FileProviderManager) RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields IFileProviderItemFields, itemIdentifier IFileProviderItemIdentifier, options FileProviderModifyItemOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestModificationOfFields:forItemWithIdentifier:options:completionHandler:"), fields, itemIdentifier, options, completionHandler)
}

// Asks the system to schedule and execute the specified operations.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/run(_:)
func (f_ FileProviderManager) RunTestingOperationsError(operations []objc.ID, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("runTestingOperations:error:"), operations, error_)
	return rv
}

// Alerts the system to changes in the specified folder’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/signalEnumerator(for:completionHandler:)
func (f_ FileProviderManager) SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier IFileProviderItemIdentifier, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("signalEnumeratorForContainerItemIdentifier:completionHandler:"), containerItemIdentifier, completion)
}

// Indicates a resolved error.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/signalErrorResolved(_:completionHandler:)
func (f_ FileProviderManager) SignalErrorResolvedCompletionHandler(error_ foundation.IError, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("signalErrorResolved:completionHandler:"), error_, completionHandler)
}

// Returns a URL for a directory for storing state information for the domain.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/stateDirectoryURL()
func (f_ FileProviderManager) StateDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("stateDirectoryURLWithError:"), error_)
	return rv
}

// Returns the URL of a directory that the File Provider extension can use to temporarily store files before passing them to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/temporaryDirectoryURL()
func (f_ FileProviderManager) TemporaryDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("temporaryDirectoryURLWithError:"), error_)
	return rv
}

// Requests a notification after the system completes all the specified changes.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/waitForChanges(below:completionHandler:)
func (f_ FileProviderManager) WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier IFileProviderItemIdentifier, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("waitForChangesOnItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Requests a notification after the domain stabilizes.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/waitForStabilization(completionHandler:)
func (f_ FileProviderManager) WaitForStabilizationWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("waitForStabilizationWithCompletionHandler:"), completionHandler)
}

// A property that returns the shared file provider manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/default
func (f_ FileProviderManager) DefaultManager() NSFileProviderManager {
	rv := objc.Send[NSFileProviderManager](f_.ID, objc.Sel("defaultManager"))
	return rv
}

// The root URL for all shared documents.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/documentStorageURL
func (f_ FileProviderManager) DocumentStorageURL() foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("documentStorageURL"))
	return rv
}

// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/providerIdentifier
func (f_ FileProviderManager) ProviderIdentifier() appkit.string {
	rv := objc.Send[appkit.string](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}


