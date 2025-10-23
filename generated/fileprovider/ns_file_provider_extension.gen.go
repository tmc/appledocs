// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	DocumentStorageURL() foundation.objc.IObject /* cross-framework: URL */
	Domain() IFileProviderDomain
	ProviderIdentifier() string /* primitive/slice/pointer. */
	// methods:
	CreateDirectoryWithNameInParentItemIdentifierCompletionHandler(directoryName string /* primitive/slice/pointer. */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	DeleteItemWithIdentifierCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	EnumeratorForContainerItemIdentifierError(containerItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, error_ unsafe.Pointer) objc.ID
	FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []string /* primitive/slice/pointer. */, size coregraphics.CGSize, perThumbnailCompletionHandler unsafe.Pointer, completionHandler unsafe.Pointer) Progress /* not a class type */
	ImportDocumentAtURLToParentItemIdentifierCompletionHandler(fileURL foundation.objc.IObject /* cross-framework URL */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	ItemForIdentifierError(identifier objc.IObject /* cross-framework FileProviderItemIdentifier */, error_ unsafe.Pointer) objc.IObject /* cross-framework: FileProviderItem */
	ItemChangedAtURL(url foundation.objc.IObject /* cross-framework URL */)
	PersistentIdentifierForItemAtURL(url foundation.objc.IObject /* cross-framework URL */) objc.IObject /* cross-framework: FileProviderItemIdentifier */
	ProvidePlaceholderAtURLCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer)
	RenameItemWithIdentifierToNameCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, itemName string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer)
	ReparentItemWithIdentifierToParentItemWithIdentifierNewNameCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, newName string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer)
	SetFavoriteRankForItemIdentifierCompletionHandler(favoriteRank foundation.objc.IObject /* cross-framework Number */, itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	SetLastUsedDateForItemIdentifierCompletionHandler(lastUsedDate foundation.objc.IObject /* cross-framework NSDate */, itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	SetTagDataForItemIdentifierCompletionHandler(tagData foundation.objc.IObject /* cross-framework NSData */, itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	StartProvidingItemAtURLCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer)
	StopProvidingItemAtURL(url foundation.objc.IObject /* cross-framework URL */)
	SupportedServiceSourcesForItemIdentifierError(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, error_ unsafe.Pointer) []objc.ID /* already interface */
	TrashItemWithIdentifierCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	UntrashItemWithIdentifierToParentItemIdentifierCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer)
	URLForItemWithPersistentIdentifier(identifier objc.IObject /* cross-framework FileProviderItemIdentifier */) foundation.objc.IObject /* cross-framework: URL */
}

// The principal class for the nonreplicated File Provider extension.
//
// To create a nonreplicated File Provider extension, start by creating a subclass of the class. When implementing your subclass, remember: Override all of the extension’s methods (except the deprecated methods), even if your implementation is only an empty method. Use your method implementations to provide access to the documents and folders managed by your file provider. Don’t call in your method implementations. Don’t use the class in macOS. Instead, create an subclass that adopts the and protocols. For more information, see .


// The principal class for the nonreplicated File Provider extension.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/placeholderURL(for:)
func (fc _FileProviderExtensionClass) PlaceholderURLForURL(url foundation.objc.IObject /* cross-framework URL */) foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](objc.ID(fc.class), objc.Sel("placeholderURLForURL:"), url)
	return rv
}


// Writes a document placeholder with the provided metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/writePlaceholder(at:withMetadata:)
func (fc _FileProviderExtensionClass) WritePlaceholderAtURLWithMetadataError(placeholderURL foundation.objc.IObject /* cross-framework URL */, metadata foundation.IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, error_)
	return rv
}


// Creates a directory with the given name inside the given parent directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/createDirectory(withName:inParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) CreateDirectoryWithNameInParentItemIdentifierCompletionHandler(directoryName string /* primitive/slice/pointer. */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("createDirectoryWithName:inParentItemIdentifier:completionHandler:"), objc.String(directoryName), parentItemIdentifier, completionHandler)
}


// Permanently deletes an item from the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/deleteItem(withIdentifier:completionHandler:)
func (f_ FileProviderExtension) DeleteItemWithIdentifierCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("deleteItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}


// Returns an enumerator for the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/enumerator(for:)
func (f_ FileProviderExtension) EnumeratorForContainerItemIdentifierError(containerItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("enumeratorForContainerItemIdentifier:error:"), containerItemIdentifier, error_)
	return rv
}


// Fetches the thumbnails for items that have been enumerated by the file provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/fetchThumbnails(for:requestedSize:perThumbnailCompletionHandler:completionHandler:)
func (f_ FileProviderExtension) FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []string /* primitive/slice/pointer. */, size coregraphics.CGSize, perThumbnailCompletionHandler unsafe.Pointer, completionHandler unsafe.Pointer) Progress /* not a class type */ {
	rv := objc.Send[Progress](f_.ID, objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"), itemIdentifiers, size, perThumbnailCompletionHandler, completionHandler)
	return rv
}


// Imports a file or package into the given parent directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/importDocument(at:toParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) ImportDocumentAtURLToParentItemIdentifierCompletionHandler(fileURL foundation.objc.IObject /* cross-framework URL */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("importDocumentAtURL:toParentItemIdentifier:completionHandler:"), fileURL, parentItemIdentifier, completionHandler)
}


// Returns a description of the item associated with the persistent identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/item(for:)
func (f_ FileProviderExtension) ItemForIdentifierError(identifier objc.IObject /* cross-framework FileProviderItemIdentifier */, error_ unsafe.Pointer) objc.IObject /* cross-framework: FileProviderItem */ {
	rv := objc.Send[FileProviderItem](f_.ID, objc.Sel("itemForIdentifier:error:"), identifier, error_)
	return rv
}


// Tells the File Provider extension that a document has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/itemChanged(at:)
func (f_ FileProviderExtension) ItemChangedAtURL(url foundation.objc.IObject /* cross-framework URL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("itemChangedAtURL:"), url)
}


// Returns a unique identifier for the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/persistentIdentifierForItem(at:)
func (f_ FileProviderExtension) PersistentIdentifierForItemAtURL(url foundation.objc.IObject /* cross-framework URL */) objc.IObject /* cross-framework: FileProviderItemIdentifier */ {
	rv := objc.Send[FileProviderItemIdentifier](f_.ID, objc.Sel("persistentIdentifierForItemAtURL:"), url)
	return rv
}


// Triggers the creation of a placeholder for the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/providePlaceholder(at:completionHandler:)
func (f_ FileProviderExtension) ProvidePlaceholderAtURLCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("providePlaceholderAtURL:completionHandler:"), url, completionHandler)
}


// Renames a document or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/renameItem(withIdentifier:toName:completionHandler:)
func (f_ FileProviderExtension) RenameItemWithIdentifierToNameCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, itemName string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("renameItemWithIdentifier:toName:completionHandler:"), itemIdentifier, objc.String(itemName), completionHandler)
}


// Moves the specified item into the given parent directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/reparentItem(withIdentifier:toParentItemWithIdentifier:newName:completionHandler:)
func (f_ FileProviderExtension) ReparentItemWithIdentifierToParentItemWithIdentifierNewNameCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, newName string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("reparentItemWithIdentifier:toParentItemWithIdentifier:newName:completionHandler:"), itemIdentifier, parentItemIdentifier, objc.String(newName), completionHandler)
}


// Marks a directory as a favorite and sets its relative order in the Favorites list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setFavoriteRank(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetFavoriteRankForItemIdentifierCompletionHandler(favoriteRank foundation.objc.IObject /* cross-framework Number */, itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFavoriteRank:forItemIdentifier:completionHandler:"), favoriteRank, itemIdentifier, completionHandler)
}


// Marks an item as recently used and sets its relative order in the Recents list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setLastUsedDate(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetLastUsedDateForItemIdentifierCompletionHandler(lastUsedDate foundation.objc.IObject /* cross-framework NSDate */, itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLastUsedDate:forItemIdentifier:completionHandler:"), lastUsedDate, itemIdentifier, completionHandler)
}


// Tags an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/setTagData(_:forItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) SetTagDataForItemIdentifierCompletionHandler(tagData foundation.objc.IObject /* cross-framework NSData */, itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTagData:forItemIdentifier:completionHandler:"), tagData, itemIdentifier, completionHandler)
}


// Provides an actual file on disk for a placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/startProvidingItem(at:completionHandler:)
func (f_ FileProviderExtension) StartProvidingItemAtURLCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("startProvidingItemAtURL:completionHandler:"), url, completionHandler)
}


// Tells the File Provider extension that a given document is no longer being accessed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/stopProvidingItem(at:)
func (f_ FileProviderExtension) StopProvidingItemAtURL(url foundation.objc.IObject /* cross-framework URL */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("stopProvidingItemAtURL:"), url)
}


// Return an array of service sources that let the host app perform actions associated with the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/supportedServiceSources(for:)
func (f_ FileProviderExtension) SupportedServiceSourcesForItemIdentifierError(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, error_ unsafe.Pointer) []objc.ID /* already interface */ {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("supportedServiceSourcesForItemIdentifier:error:"), itemIdentifier, error_)
	return rv
}


// Moves an item into the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/trashItem(withIdentifier:completionHandler:)
func (f_ FileProviderExtension) TrashItemWithIdentifierCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("trashItemWithIdentifier:completionHandler:"), itemIdentifier, completionHandler)
}


// Moves an item out of the trash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/untrashItem(withIdentifier:toParentItemIdentifier:completionHandler:)
func (f_ FileProviderExtension) UntrashItemWithIdentifierToParentItemIdentifierCompletionHandler(itemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, parentItemIdentifier objc.IObject /* cross-framework FileProviderItemIdentifier */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("untrashItemWithIdentifier:toParentItemIdentifier:completionHandler:"), itemIdentifier, parentItemIdentifier, completionHandler)
}


// Returns the URL for a given persistent identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/urlForItem(withPersistentIdentifier:)
func (f_ FileProviderExtension) URLForItemWithPersistentIdentifier(identifier objc.IObject /* cross-framework FileProviderItemIdentifier */) foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("URLForItemWithPersistentIdentifier:"), identifier)
	return rv
}


// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/documentStorageURL
func (f_ FileProviderExtension) DocumentStorageURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("documentStorageURL"))
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
func (f_ FileProviderExtension) ProviderIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}



