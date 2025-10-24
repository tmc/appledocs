//go:build darwin && ios

// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FileProviderExtension


// Creates a directory with the given name inside the given parent directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/createDirectory(withName:inParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) CreateDirectoryWithNameInParentItemIdentifierCompletionHandler(directoryName objc.IObject /* cross-framework: NSString */, parentItemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("createDirectoryWithName:inParentItemIdentifier:completionHandler:"), directoryName, parentItemIdentifier, completionHandler)
}

// Permanently deletes an item from the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/deleteItem(withIdentifier:completionHandler:)
func (f_ FileProviderExtension) DeleteItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("deleteItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Returns an enumerator for the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/enumerator(for:)
func (f_ FileProviderExtension) EnumeratorForContainerItemIdentifierError(containerItemIdentifier FileProviderItemIdentifier /* typedef */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("enumeratorForContainerItemIdentifier:error:"), containerItemIdentifier, error_)
	return rv
}

// Fetches the thumbnails for items that have been enumerated by the file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/fetchThumbnails(for:requestedSize:perThumbnailCompletionHandler:completionHandler:)
func (f_ FileProviderExtension) FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []string, size corefoundation.CGSize, perThumbnailCompletionHandler unsafe.Pointer, completionHandler unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](f_.ID, objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"), itemIdentifiers, size, perThumbnailCompletionHandler, completionHandler)
	return rv
}

// Imports a file or package into the given parent directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/importDocument(at:toParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) ImportDocumentAtURLToParentItemIdentifierCompletionHandler(fileURL objc.IObject /* cross-framework: NSURL */, parentItemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("importDocumentAtURL:toParentItemIdentifier:completionHandler:"), fileURL, parentItemIdentifier, completionHandler)
}

// Returns a description of the item associated with the persistent identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/item(for:)
func (f_ FileProviderExtension) ItemForIdentifierError(identifier FileProviderItemIdentifier /* typedef */, error_ unsafe.Pointer) FileProviderItem /* typedef */ {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("itemForIdentifier:error:"), identifier, error_)
	return rv
}

// Tells the File Provider extension that a document has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/itemChanged(at:)
func (f_ FileProviderExtension) ItemChangedAtURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("itemChangedAtURL:"), url)
}

// Returns a unique identifier for the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/persistentIdentifierForItem(at:)
func (f_ FileProviderExtension) PersistentIdentifierForItemAtURL(url objc.IObject /* cross-framework: NSURL */) FileProviderItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("persistentIdentifierForItemAtURL:"), url)
	return rv
}

// Triggers the creation of a placeholder for the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/providePlaceholder(at:completionHandler:)
func (f_ FileProviderExtension) ProvidePlaceholderAtURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("providePlaceholderAtURL:completionHandler:"), url, completionHandler)
}

// Renames a document or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/renameItem(withIdentifier:toName:completionHandler:)
func (f_ FileProviderExtension) RenameItemWithIdentifierToNameCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, itemName objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("renameItemWithIdentifier:toName:completionHandler:"), itemIdentifier, itemName, completionHandler)
}

// Moves the specified item into the given parent directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/reparentItem(withIdentifier:toParentItemWithIdentifier:newName:completionHandler:)
func (f_ FileProviderExtension) ReparentItemWithIdentifierToParentItemWithIdentifierNewNameCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, parentItemIdentifier FileProviderItemIdentifier /* typedef */, newName objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reparentItemWithIdentifier:toParentItemWithIdentifier:newName:completionHandler:"), itemIdentifier, parentItemIdentifier, newName, completionHandler)
}

// Marks a directory as a favorite and sets its relative order in the Favorites list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setFavoriteRank(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetFavoriteRankForItemIdentifierCompletionHandler(favoriteRank objc.IObject /* cross-framework: NSNumber */, itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFavoriteRank:forItemIdentifier:completionHandler:"), favoriteRank, itemIdentifier, completionHandler)
}

// Marks an item as recently used and sets its relative order in the Recents list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setLastUsedDate(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetLastUsedDateForItemIdentifierCompletionHandler(lastUsedDate objc.IObject /* cross-framework: NSDate */, itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLastUsedDate:forItemIdentifier:completionHandler:"), lastUsedDate, itemIdentifier, completionHandler)
}

// Tags an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setTagData(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetTagDataForItemIdentifierCompletionHandler(tagData objc.IObject /* cross-framework: NSData */, itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTagData:forItemIdentifier:completionHandler:"), tagData, itemIdentifier, completionHandler)
}

// Provides an actual file on disk for a placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/startProvidingItem(at:completionHandler:)
func (f_ FileProviderExtension) StartProvidingItemAtURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("startProvidingItemAtURL:completionHandler:"), url, completionHandler)
}

// Tells the File Provider extension that a given document is no longer being accessed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/stopProvidingItem(at:)
func (f_ FileProviderExtension) StopProvidingItemAtURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("stopProvidingItemAtURL:"), url)
}

// Return an array of service sources that let the host app perform actions associated with the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/supportedServiceSources(for:)
func (f_ FileProviderExtension) SupportedServiceSourcesForItemIdentifierError(itemIdentifier FileProviderItemIdentifier /* typedef */, error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("supportedServiceSourcesForItemIdentifier:error:"), itemIdentifier, error_)
	return rv
}

// Moves an item into the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/trashItem(withIdentifier:completionHandler:)
func (f_ FileProviderExtension) TrashItemWithIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("trashItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Moves an item out of the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/untrashItem(withIdentifier:toParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) UntrashItemWithIdentifierToParentItemIdentifierCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, parentItemIdentifier FileProviderItemIdentifier /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("untrashItemWithIdentifier:toParentItemIdentifier:completionHandler:"), itemIdentifier, parentItemIdentifier, completionHandler)
}

// Returns the URL for a given persistent identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/urlForItem(withPersistentIdentifier:)
func (f_ FileProviderExtension) URLForItemWithPersistentIdentifier(identifier FileProviderItemIdentifier /* typedef */) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("URLForItemWithPersistentIdentifier:"), identifier)
	return rv
}

// iOS-only properties

// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/documentStorageURL
func (f_ FileProviderExtension) DocumentStorageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("documentStorageURL"))
	return rv
}

// The domain managed by this file provider object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/domain
func (f_ FileProviderExtension) Domain() IFileProviderDomain {
	rv := objc.Send[FileProviderDomain](f_.ID, objc.Sel("domain"))
	return rv
}

// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/providerIdentifier
func (f_ FileProviderExtension) ProviderIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}





