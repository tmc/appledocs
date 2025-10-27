// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"
)

// PFSVolumeOperations is the FSVolumeOperations protocol interface.
//
// Methods that all volumes implement to provide required capabilities.
//
// Availability:
//   - macOS 15.4+
//
// See: doc://FSKit/documentation/FSKit/FSVolume/Operations
type PFSVolumeOperations interface {
	// Required methods
	ActivateWithOptionsReplyHandler(options IFSTaskOptions, reply unsafe.Pointer)
	CreateItemNamedTypeInDirectoryAttributesReplyHandler(name IFSFileName, type_ FSItemType, directory IFSItem, newAttributes IFSItemSetAttributesRequest, reply unsafe.Pointer)
	CreateLinkToItemNamedInDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply unsafe.Pointer)
	CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler(name IFSFileName, directory IFSItem, newAttributes IFSItemSetAttributesRequest, contents IFSFileName, reply unsafe.Pointer)
	DeactivateWithOptionsReplyHandler(options FSDeactivateOptions, reply unsafe.Pointer)
	EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler(directory IFSItem, cookie FSDirectoryCookie, verifier FSDirectoryVerifier, attributes IFSItemGetAttributesRequest, packer IFSDirectoryEntryPacker, reply unsafe.Pointer)
	GetAttributesOfItemReplyHandler(desiredAttributes IFSItemGetAttributesRequest, item IFSItem, reply unsafe.Pointer)
	LookupItemNamedInDirectoryReplyHandler(name IFSFileName, directory IFSItem, reply unsafe.Pointer)
	MountWithOptionsReplyHandler(options IFSTaskOptions, reply unsafe.Pointer)
	ReadSymbolicLinkReplyHandler(item IFSItem, reply unsafe.Pointer)
	ReclaimItemReplyHandler(item IFSItem, reply unsafe.Pointer)
	RemoveItemNamedFromDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply unsafe.Pointer)
	RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler(item IFSItem, sourceDirectory IFSItem, sourceName IFSFileName, destinationName IFSFileName, destinationDirectory IFSItem, overItem IFSItem, reply unsafe.Pointer)
	SetAttributesOnItemReplyHandler(newAttributes IFSItemSetAttributesRequest, item IFSItem, reply unsafe.Pointer)
	SynchronizeWithFlagsReplyHandler(flags FSSyncFlags, reply unsafe.Pointer)
	UnmountWithReplyHandler(reply unsafe.Pointer)
}
