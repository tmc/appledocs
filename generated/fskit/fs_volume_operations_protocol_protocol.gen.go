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
	ActivateWithOptionsReplyHandler(options IFSTaskOptions, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: ActivateWithOptionsReplyHandler */
	CreateItemNamedTypeInDirectoryAttributesReplyHandler(name IFSFileName, type_ FSItemType, directory IFSItem, newAttributes IFSItemSetAttributesRequest, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CreateItemNamedTypeInDirectoryAttributesReplyHandler */
	CreateLinkToItemNamedInDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CreateLinkToItemNamedInDirectoryReplyHandler */
	CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler(name IFSFileName, directory IFSItem, newAttributes IFSItemSetAttributesRequest, contents IFSFileName, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler */
	DeactivateWithOptionsReplyHandler(options FSDeactivateOptions, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: DeactivateWithOptionsReplyHandler */
	EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler(directory IFSItem, cookie FSDirectoryCookie /* typedef */, verifier FSDirectoryVerifier /* typedef */, attributes IFSItemGetAttributesRequest, packer IFSDirectoryEntryPacker, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler */
	GetAttributesOfItemReplyHandler(desiredAttributes IFSItemGetAttributesRequest, item IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: GetAttributesOfItemReplyHandler */
	LookupItemNamedInDirectoryReplyHandler(name IFSFileName, directory IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: LookupItemNamedInDirectoryReplyHandler */
	MountWithOptionsReplyHandler(options IFSTaskOptions, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: MountWithOptionsReplyHandler */
	ReadSymbolicLinkReplyHandler(item IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: ReadSymbolicLinkReplyHandler */
	ReclaimItemReplyHandler(item IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: ReclaimItemReplyHandler */
	RemoveItemNamedFromDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: RemoveItemNamedFromDirectoryReplyHandler */
	RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler(item IFSItem, sourceDirectory IFSItem, sourceName IFSFileName, destinationName IFSFileName, destinationDirectory IFSItem, overItem IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler */
	SetAttributesOnItemReplyHandler(newAttributes IFSItemSetAttributesRequest, item IFSItem, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: SetAttributesOnItemReplyHandler */
	SynchronizeWithFlagsReplyHandler(flags FSSyncFlags, reply unsafe.Pointer)/* debug [protocol_interface/required_method]: SynchronizeWithFlagsReplyHandler */
	UnmountWithReplyHandler(reply unsafe.Pointer)/* debug [protocol_interface/required_method]: UnmountWithReplyHandler */
}
