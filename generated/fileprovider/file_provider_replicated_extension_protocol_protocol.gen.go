// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderReplicatedExtension is the NSFileProviderReplicatedExtension protocol interface.
//
// A File Provider extension in which the system replicates the contents on disk.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderReplicatedExtension
type PFileProviderReplicatedExtension interface {
	// Required methods
	CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate FileProviderItem /* typedef */, fields FileProviderItemFields, url objc.IObject /* cross-framework: NSURL */, options FileProviderCreateItemOptions, request IFileProviderRequest, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler */
	DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier FileProviderItemIdentifier /* typedef */, version IFileProviderItemVersion, options FileProviderDeleteItemOptions, request IFileProviderRequest, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler */
	FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier FileProviderItemIdentifier /* typedef */, requestedVersion IFileProviderItemVersion, request IFileProviderRequest, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: FetchContentsForItemWithIdentifierVersionRequestCompletionHandler */
	InitWithDomain(domain IFileProviderDomain) unsafe.Pointer/* debug [protocol_interface/required_method]: InitWithDomain */
	Invalidate()/* debug [protocol_interface/required_method]: Invalidate */
	ItemForIdentifierRequestCompletionHandler(identifier FileProviderItemIdentifier /* typedef */, request IFileProviderRequest, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: ItemForIdentifierRequestCompletionHandler */
	ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item FileProviderItem /* typedef */, version IFileProviderItemVersion, changedFields FileProviderItemFields, newContents objc.IObject /* cross-framework: NSURL */, options FileProviderModifyItemOptions, request IFileProviderRequest, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler */
	// Optional methods
	ImportDidFinishWithCompletionHandler(completionHandler unsafe.Pointer)
	HasImportDidFinishWithCompletionHandler() bool
	MaterializedItemsDidChangeWithCompletionHandler(completionHandler unsafe.Pointer)
	HasMaterializedItemsDidChangeWithCompletionHandler() bool
	PendingItemsDidChangeWithCompletionHandler(completionHandler unsafe.Pointer)
	HasPendingItemsDidChangeWithCompletionHandler() bool
}
