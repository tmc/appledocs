// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [FileProviderExtension] class.
var (
	FileProviderExtensionClass     _FileProviderExtensionClass
	FileProviderExtensionClassOnce sync.Once
)

func getFileProviderExtensionClass() _FileProviderExtensionClass {
	FileProviderExtensionClassOnce.Do(func() {
		FileProviderExtensionClass = _FileProviderExtensionClass{objc.GetClass("NSFileProviderExtension")}
	})
	return FileProviderExtensionClass
}

type _FileProviderExtensionClass struct {
	class objc.Class
}

// An interface definition for the [FileProviderExtension] class.
type IFileProviderExtension interface {
	objectivec.IObject
	CreateDirectoryWithNameInParentItemIdentifierCompletionHandler(directoryName string, parentItemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	DeleteItemWithIdentifierCompletionHandler(itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	EnumeratorForContainerItemIdentifierError(containerItemIdentifier unsafe.Pointer, error_ unsafe.Pointer) objc.ID
	FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers unsafe.Pointer, size coregraphics.CGSize, perThumbnailCompletionHandler unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer
	ImportDocumentAtURLToParentItemIdentifierCompletionHandler(fileURL unsafe.Pointer, parentItemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	ItemForIdentifierError(identifier unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	ItemChangedAtURL(url unsafe.Pointer)
	PersistentIdentifierForItemAtURL(url unsafe.Pointer) unsafe.Pointer
	ProvidePlaceholderAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer)
	RenameItemWithIdentifierToNameCompletionHandler(itemIdentifier unsafe.Pointer, itemName string, completionHandler unsafe.Pointer)
	ReparentItemWithIdentifierToParentItemWithIdentifierNewNameCompletionHandler(itemIdentifier unsafe.Pointer, parentItemIdentifier unsafe.Pointer, newName string, completionHandler unsafe.Pointer)
	SetFavoriteRankForItemIdentifierCompletionHandler(favoriteRank unsafe.Pointer, itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	SetLastUsedDateForItemIdentifierCompletionHandler(lastUsedDate unsafe.Pointer, itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	SetTagDataForItemIdentifierCompletionHandler(tagData unsafe.Pointer, itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	StartProvidingItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer)
	StopProvidingItemAtURL(url unsafe.Pointer)
	SupportedServiceSourcesForItemIdentifierError(itemIdentifier unsafe.Pointer, error_ unsafe.Pointer) []objc.ID
	TrashItemWithIdentifierCompletionHandler(itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	UntrashItemWithIdentifierToParentItemIdentifierCompletionHandler(itemIdentifier unsafe.Pointer, parentItemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer)
	URLForItemWithPersistentIdentifier(identifier unsafe.Pointer) unsafe.Pointer
}

// The principal class for the nonreplicated File Provider extension.
//
// To create a nonreplicated File Provider extension, start by creating a subclass of the class. When implementing your subclass, remember: Override all of the extension’s methods (except the deprecated methods), even if your implementation is only an empty method. Use your method implementations to provide access to the documents and folders managed by your file provider. Don’t call in your method implementations. Don’t use the class in macOS. Instead, create an subclass that adopts the and protocols. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension
type FileProviderExtension struct {
	objectivec.Object
}

// FileProviderExtensionFrom constructs a [FileProviderExtension] from an unsafe.Pointer.
//
// The principal class for the nonreplicated File Provider extension.
func FileProviderExtensionFrom(ptr unsafe.Pointer) FileProviderExtension {
	return FileProviderExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileProviderExtensionClass) Alloc() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileProviderExtensionClass) New() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderExtension) Init() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderExtension) Autorelease() FileProviderExtension {
	rv := objc.Send[FileProviderExtension](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderExtension creates a new FileProviderExtension instance.
func NewFileProviderExtension() FileProviderExtension {
	return getFileProviderExtensionClass().New()
}


// Returns a placeholder URL for a given document URL.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/placeholderURL(for:)
func (fc _FileProviderExtensionClass) PlaceholderURLForURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("placeholderURLForURL:"), url)
	return rv
}

// Writes a document placeholder with the provided metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/writePlaceholder(at:withMetadata:)
func (fc _FileProviderExtensionClass) WritePlaceholderAtURLWithMetadataError(placeholderURL unsafe.Pointer, metadata unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, error_)
	return rv
}

// Creates a directory with the given name inside the given parent directory.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/createDirectory(withName:inParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) CreateDirectoryWithNameInParentItemIdentifierCompletionHandler(directoryName string, parentItemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("createDirectoryWithName:inParentItemIdentifier:completionHandler:"), objc.String(directoryName), parentItemIdentifier, completionHandler)
}

// Permanently deletes an item from the trash.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/deleteItem(withIdentifier:completionHandler:)
func (f_ FileProviderExtension) DeleteItemWithIdentifierCompletionHandler(itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("deleteItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Returns an enumerator for the specified item.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/enumerator(for:)
func (f_ FileProviderExtension) EnumeratorForContainerItemIdentifierError(containerItemIdentifier unsafe.Pointer, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("enumeratorForContainerItemIdentifier:error:"), containerItemIdentifier, error_)
	return rv
}

// Fetches the thumbnails for items that have been enumerated by the file provider.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/fetchThumbnails(for:requestedSize:perThumbnailCompletionHandler:completionHandler:)
func (f_ FileProviderExtension) FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers unsafe.Pointer, size coregraphics.CGSize, perThumbnailCompletionHandler unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"), itemIdentifiers, size, perThumbnailCompletionHandler, completionHandler)
	return rv
}

// Imports a file or package into the given parent directory.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/importDocument(at:toParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) ImportDocumentAtURLToParentItemIdentifierCompletionHandler(fileURL unsafe.Pointer, parentItemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("importDocumentAtURL:toParentItemIdentifier:completionHandler:"), fileURL, parentItemIdentifier, completionHandler)
}

// Returns a description of the item associated with the persistent identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/item(for:)
func (f_ FileProviderExtension) ItemForIdentifierError(identifier unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("itemForIdentifier:error:"), identifier, error_)
	return rv
}

// Tells the File Provider extension that a document has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/itemChanged(at:)
func (f_ FileProviderExtension) ItemChangedAtURL(url unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("itemChangedAtURL:"), url)
}

// Returns a unique identifier for the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/persistentIdentifierForItem(at:)
func (f_ FileProviderExtension) PersistentIdentifierForItemAtURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("persistentIdentifierForItemAtURL:"), url)
	return rv
}

// Triggers the creation of a placeholder for the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/providePlaceholder(at:completionHandler:)
func (f_ FileProviderExtension) ProvidePlaceholderAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("providePlaceholderAtURL:completionHandler:"), url, completionHandler)
}

// Renames a document or directory.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/renameItem(withIdentifier:toName:completionHandler:)
func (f_ FileProviderExtension) RenameItemWithIdentifierToNameCompletionHandler(itemIdentifier unsafe.Pointer, itemName string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("renameItemWithIdentifier:toName:completionHandler:"), itemIdentifier, objc.String(itemName), completionHandler)
}

// Moves the specified item into the given parent directory.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/reparentItem(withIdentifier:toParentItemWithIdentifier:newName:completionHandler:)
func (f_ FileProviderExtension) ReparentItemWithIdentifierToParentItemWithIdentifierNewNameCompletionHandler(itemIdentifier unsafe.Pointer, parentItemIdentifier unsafe.Pointer, newName string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reparentItemWithIdentifier:toParentItemWithIdentifier:newName:completionHandler:"), itemIdentifier, parentItemIdentifier, objc.String(newName), completionHandler)
}

// Marks a directory as a favorite and sets its relative order in the Favorites list.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setFavoriteRank(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetFavoriteRankForItemIdentifierCompletionHandler(favoriteRank unsafe.Pointer, itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFavoriteRank:forItemIdentifier:completionHandler:"), favoriteRank, itemIdentifier, completionHandler)
}

// Marks an item as recently used and sets its relative order in the Recents list.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setLastUsedDate(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetLastUsedDateForItemIdentifierCompletionHandler(lastUsedDate unsafe.Pointer, itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLastUsedDate:forItemIdentifier:completionHandler:"), lastUsedDate, itemIdentifier, completionHandler)
}

// Tags an item.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setTagData(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetTagDataForItemIdentifierCompletionHandler(tagData unsafe.Pointer, itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTagData:forItemIdentifier:completionHandler:"), tagData, itemIdentifier, completionHandler)
}

// Provides an actual file on disk for a placeholder.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/startProvidingItem(at:completionHandler:)
func (f_ FileProviderExtension) StartProvidingItemAtURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("startProvidingItemAtURL:completionHandler:"), url, completionHandler)
}

// Tells the File Provider extension that a given document is no longer being accessed.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/stopProvidingItem(at:)
func (f_ FileProviderExtension) StopProvidingItemAtURL(url unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("stopProvidingItemAtURL:"), url)
}

// Return an array of service sources that let the host app perform actions associated with the specified item.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/supportedServiceSources(for:)
func (f_ FileProviderExtension) SupportedServiceSourcesForItemIdentifierError(itemIdentifier unsafe.Pointer, error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("supportedServiceSourcesForItemIdentifier:error:"), itemIdentifier, error_)
	return rv
}

// Moves an item into the trash.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/trashItem(withIdentifier:completionHandler:)
func (f_ FileProviderExtension) TrashItemWithIdentifierCompletionHandler(itemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("trashItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}

// Moves an item out of the trash.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/untrashItem(withIdentifier:toParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) UntrashItemWithIdentifierToParentItemIdentifierCompletionHandler(itemIdentifier unsafe.Pointer, parentItemIdentifier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("untrashItemWithIdentifier:toParentItemIdentifier:completionHandler:"), itemIdentifier, parentItemIdentifier, completionHandler)
}

// Returns the URL for a given persistent identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/urlForItem(withPersistentIdentifier:)
func (f_ FileProviderExtension) URLForItemWithPersistentIdentifier(identifier unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("URLForItemWithPersistentIdentifier:"), identifier)
	return rv
}

// The root URL for all shared documents.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/documentStorageURL
func (f_ FileProviderExtension) DocumentStorageURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("documentStorageURL"))
	return rv
}

// The domain managed by this file provider object.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/domain
func (f_ FileProviderExtension) Domain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("domain"))
	return rv
}

// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/providerIdentifier
func (f_ FileProviderExtension) ProviderIdentifier() string {
	rv := objc.Send[string](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}



