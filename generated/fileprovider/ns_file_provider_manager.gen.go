// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderManager */


/* debug [class_header]: Header for NSFileProviderManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderManager */
// An interface definition for the [FileProviderManager] class.
type IFileProviderManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderManager */
	// methods:
	ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders IFileProviderKnownFolderLocations, localizedReason objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	DisconnectWithReasonOptionsCompletionHandler(localizedReason objc.IObject /* cross-framework: NSString */, options FileProviderManagerDisconnectionOptions, completionHandler unsafe.Pointer)
	EnumeratorForMaterializedItems() unsafe.Pointer
	EnumeratorForPendingItems() unsafe.Pointer
	EvictItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer)
	GetServiceWithNameItemIdentifierCompletionHandler(serviceName FileProviderServiceName /* not a class type */, itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer)
	GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer)
	GlobalProgressForKind(kind ProgressFileOperationKind /* not a class type */) foundation.Progress
	ListAvailableTestingOperationsWithError(error_ unsafe.Pointer) []objc.ID
	ReconnectWithCompletionHandler(completionHandler unsafe.Pointer)
	RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task avfoundation.URLSessionTask, identifier FileProviderItemIdentifier /* typedef */, completion unsafe.Pointer)
	ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer)
	ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders FileProviderKnownFolders, localizedReason objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, errorReason objc.IObject /* cross-framework: Error */, completionHandler unsafe.Pointer)
	RequestDownloadForItemWithIdentifierRequestedRangeCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, rangeToMaterialize corefoundation.Range, completionHandler unsafe.Pointer)
	RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields FileProviderItemFields, itemIdentifier FileProviderItemIdentifier /* typedef */, options FileProviderModifyItemOptions, completionHandler unsafe.Pointer)
	RunTestingOperationsError(operations []objc.ID, error_ unsafe.Pointer) foundation.IDictionary
	SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier FileProviderItemIdentifier /* typedef */, completion unsafe.Pointer)
	SignalErrorResolvedCompletionHandler(error_ objc.IObject /* cross-framework: Error */, completionHandler unsafe.Pointer)
	StateDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL
	TemporaryDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL
	WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer)
	WaitForStabilizationWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderManager */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderManagerClass) Alloc() FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderManager */
// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.


// A manager object that you use to communicate with the file provider from either your app or your File Provider extension.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderManager */

// Returns a newly created file provider manager for the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/init(for:)
func NewFileProviderManagerForDomain(domain IFileProviderDomain) FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(getFileProviderManagerClass().class), objc.Sel("managerForDomain:"), domain)
	return rv
}/* debug [class_init_methods/constructor]: NewFileProviderManagerForDomain */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderManager */

// Adds a domain to the File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/add(_:completionHandler:)
func (fc _FileProviderManagerClass) AddDomainCompletionHandler(domain IFileProviderDomain, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("addDomain:completionHandler:"), domain, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AddDomainCompletionHandler) */


// Check if a URL is eligible for storing a domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/checkDomainsCanBeStored:onVolumeAtURL:unsupportedReason:error:
func (fc _FileProviderManagerClass) CheckDomainsCanBeStoredOnVolumeAtURLUnsupportedReasonError(eligible bool, url objc.IObject /* cross-framework: NSURL */, unsupportedReason FileProviderVolumeUnsupportedReason, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("checkDomainsCanBeStored:onVolumeAtURL:unsupportedReason:error:"), eligible, url, unsupportedReason, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckDomainsCanBeStoredOnVolumeAtURLUnsupportedReasonError) */


// Returns all of the File Provider extension’s domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getDomainsWithCompletionHandler(_:)
func (fc _FileProviderManagerClass) GetDomainsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("getDomainsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetDomainsWithCompletionHandler) */


// Returns the identifier and domain for a user-visible URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getIdentifierForUserVisibleFile(at:completionHandler:)
func (fc _FileProviderManagerClass) GetIdentifierForUserVisibleFileAtURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("getIdentifierForUserVisibleFileAtURL:completionHandler:"), url, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetIdentifierForUserVisibleFileAtURLCompletionHandler) */


// Creates a new domain that takes ownership of on-disk data that your app previously managed without a file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/import(_:fromDirectoryAt:completionHandler:)
func (fc _FileProviderManagerClass) ImportDomainFromDirectoryAtURLCompletionHandler(domain IFileProviderDomain, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("importDomain:fromDirectoryAtURL:completionHandler:"), domain, url, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImportDomainFromDirectoryAtURLCompletionHandler) */


// Returns a newly created file provider manager for the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/init(for:)
func (fc _FileProviderManagerClass) ManagerForDomain(domain IFileProviderDomain) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("managerForDomain:"), domain)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ManagerForDomain) */


// Returns a placeholder URL for a given document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/placeholderURL(for:)
func (fc _FileProviderManagerClass) PlaceholderURLForURL(url objc.IObject /* cross-framework: NSURL */) foundation.URL {
	rv := objc.Send[foundation.URL](objc.ID(fc.class), objc.Sel("placeholderURLForURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlaceholderURLForURL) */


// Removes a domain from the File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/remove(_:completionHandler:)
func (fc _FileProviderManagerClass) RemoveDomainCompletionHandler(domain IFileProviderDomain, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeDomain:completionHandler:"), domain, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveDomainCompletionHandler) */


// Removes a domain from the File Provider extension using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/remove(_:mode:completionHandler:)
func (fc _FileProviderManagerClass) RemoveDomainModeCompletionHandler(domain IFileProviderDomain, mode FileProviderDomainRemovalMode, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeDomain:mode:completionHandler:"), domain, mode, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveDomainModeCompletionHandler) */


// Removes all domains from the File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/removeAllDomains(completionHandler:)
func (fc _FileProviderManagerClass) RemoveAllDomainsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("removeAllDomainsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveAllDomainsWithCompletionHandler) */


// Writes a document placeholder with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/writePlaceholder(at:withMetadata:)
func (fc _FileProviderManagerClass) WritePlaceholderAtURLWithMetadataError(placeholderURL objc.IObject /* cross-framework: NSURL */, metadata FileProviderItem /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WritePlaceholderAtURLWithMetadataError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderManager */

// A property that returns the shared file provider manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/default
func (fc _FileProviderManagerClass) DefaultManager() FileProviderManager {
	rv := objc.Send[FileProviderManager](objc.ID(fc.class), objc.Sel("defaultManager"))
	return rv
}/* debug [class_properties_class/property]: defaultManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderManager */

// Asks the domain to sync the specified known folders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/claimKnownFolders(_:localizedReason:completionHandler:)
func (f_ FileProviderManager) ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders IFileProviderKnownFolderLocations, localizedReason objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("claimKnownFolders:localizedReason:completionHandler:"), knownFolders, localizedReason, completionHandler)
}/* debug [instance_methods/method]: ClaimKnownFoldersLocalizedReasonCompletionHandler */


// Disconnects the domain from the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/disconnect(reason:options:completionHandler:)
func (f_ FileProviderManager) DisconnectWithReasonOptionsCompletionHandler(localizedReason objc.IObject /* cross-framework: NSString */, options FileProviderManagerDisconnectionOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("disconnectWithReason:options:completionHandler:"), localizedReason, options, completionHandler)
}/* debug [instance_methods/method]: DisconnectWithReasonOptionsCompletionHandler */


// Returns an enumerator for all the items the system currently stores on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/enumeratorForMaterializedItems()
func (f_ FileProviderManager) EnumeratorForMaterializedItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorForMaterializedItems"))
	return rv
}/* debug [instance_methods/method]: EnumeratorForMaterializedItems */


// Returns an enumerator for the set of pending items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/enumeratorForPendingItems()
func (f_ FileProviderManager) EnumeratorForPendingItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorForPendingItems"))
	return rv
}/* debug [instance_methods/method]: EnumeratorForPendingItems */


// Asks the system to remove an item from its cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/evictItem(identifier:completionHandler:)
func (f_ FileProviderManager) EvictItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("evictItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}/* debug [instance_methods/method]: EvictItemWithIdentifierCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getService(named:for:completionHandler:)
func (f_ FileProviderManager) GetServiceWithNameItemIdentifierCompletionHandler(serviceName FileProviderServiceName /* not a class type */, itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getServiceWithName:itemIdentifier:completionHandler:"), serviceName, itemIdentifier, completionHandler)
}/* debug [instance_methods/method]: GetServiceWithNameItemIdentifierCompletionHandler */


// Returns the user-visible URL for an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getUserVisibleURL(for:completionHandler:)
func (f_ FileProviderManager) GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getUserVisibleURLForItemIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}/* debug [instance_methods/method]: GetUserVisibleURLForItemIdentifierCompletionHandler */


// Returns a progress object that tracks either the uploading or downloading of items from the File Provider extension’s remote storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/globalProgress(for:)
func (f_ FileProviderManager) GlobalProgressForKind(kind ProgressFileOperationKind /* not a class type */) foundation.Progress {
	rv := objc.Send[foundation.Progress](f_.ID, objc.Sel("globalProgressForKind:"), kind)
	return rv
}/* debug [instance_methods/method]: GlobalProgressForKind */


// Lists all the operations that are ready for scheduling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/listAvailableTestingOperations()
func (f_ FileProviderManager) ListAvailableTestingOperationsWithError(error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("listAvailableTestingOperationsWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: ListAvailableTestingOperationsWithError */


// Reconnects the domain with the extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/reconnect(completionHandler:)
func (f_ FileProviderManager) ReconnectWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reconnectWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ReconnectWithCompletionHandler */


// Registers the URL session task responsible for the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/register(_:forItemWithIdentifier:completionHandler:)
func (f_ FileProviderManager) RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task avfoundation.URLSessionTask, identifier FileProviderItemIdentifier /* typedef */, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("registerURLSessionTask:forItemWithIdentifier:completionHandler:"), task, identifier, completion)
}/* debug [instance_methods/method]: RegisterURLSessionTaskForItemWithIdentifierCompletionHandler */


// Tells the system to reimport the item and its content recursively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/reimportItems(below:completionHandler:)
func (f_ FileProviderManager) ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reimportItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}/* debug [instance_methods/method]: ReimportItemsBelowItemWithIdentifierCompletionHandler */


// Asks the system to stop replicating the specified known folders in the domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/releaseKnownFolders(_:localizedReason:completionHandler:)
func (f_ FileProviderManager) ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders FileProviderKnownFolders, localizedReason objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("releaseKnownFolders:localizedReason:completionHandler:"), knownFolders, localizedReason, completionHandler)
}/* debug [instance_methods/method]: ReleaseKnownFoldersLocalizedReasonCompletionHandler */


// Requests a diagnostics collection for use when working directly with Apple to improve sync behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestDiagnosticCollection(for:errorReason:completionHandler:)
func (f_ FileProviderManager) RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, errorReason objc.IObject /* cross-framework: Error */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestDiagnosticCollectionForItemWithIdentifier:errorReason:completionHandler:"), itemIdentifier, errorReason, completionHandler)
}/* debug [instance_methods/method]: RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestDownloadForItemWithIdentifier:requestedRange:completionHandler:
func (f_ FileProviderManager) RequestDownloadForItemWithIdentifierRequestedRangeCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, rangeToMaterialize corefoundation.Range, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestDownloadForItemWithIdentifier:requestedRange:completionHandler:"), itemIdentifier, rangeToMaterialize, completionHandler)
}/* debug [instance_methods/method]: RequestDownloadForItemWithIdentifierRequestedRangeCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestModification(of:forItemWithIdentifier:options:completionHandler:)
func (f_ FileProviderManager) RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields FileProviderItemFields, itemIdentifier FileProviderItemIdentifier /* typedef */, options FileProviderModifyItemOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestModificationOfFields:forItemWithIdentifier:options:completionHandler:"), fields, itemIdentifier, options, completionHandler)
}/* debug [instance_methods/method]: RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler */


// Asks the system to schedule and execute the specified operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/run(_:)
func (f_ FileProviderManager) RunTestingOperationsError(operations []objc.ID, error_ unsafe.Pointer) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](f_.ID, objc.Sel("runTestingOperations:error:"), operations, error_)
	return rv
}/* debug [instance_methods/method]: RunTestingOperationsError */


// Alerts the system to changes in the specified folder’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/signalEnumerator(for:completionHandler:)
func (f_ FileProviderManager) SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier FileProviderItemIdentifier /* typedef */, completion unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("signalEnumeratorForContainerItemIdentifier:completionHandler:"), containerItemIdentifier, completion)
}/* debug [instance_methods/method]: SignalEnumeratorForContainerItemIdentifierCompletionHandler */


// Indicates a resolved error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/signalErrorResolved(_:completionHandler:)
func (f_ FileProviderManager) SignalErrorResolvedCompletionHandler(error_ objc.IObject /* cross-framework: Error */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("signalErrorResolved:completionHandler:"), error_, completionHandler)
}/* debug [instance_methods/method]: SignalErrorResolvedCompletionHandler */


// Returns a URL for a directory for storing state information for the domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/stateDirectoryURL()
func (f_ FileProviderManager) StateDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("stateDirectoryURLWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: StateDirectoryURLWithError */


// Returns the URL of a directory that the File Provider extension can use to temporarily store files before passing them to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/temporaryDirectoryURL()
func (f_ FileProviderManager) TemporaryDirectoryURLWithError(error_ unsafe.Pointer) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("temporaryDirectoryURLWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: TemporaryDirectoryURLWithError */


// Requests a notification after the system completes all the specified changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/waitForChanges(below:completionHandler:)
func (f_ FileProviderManager) WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("waitForChangesOnItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}/* debug [instance_methods/method]: WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler */


// Requests a notification after the domain stabilizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/waitForStabilization(completionHandler:)
func (f_ FileProviderManager) WaitForStabilizationWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("waitForStabilizationWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: WaitForStabilizationWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderManager */


