// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/appledocs/generated/objc"
)

// PVZNetworkBlockDeviceStorageDeviceAttachmentDelegate is the VZNetworkBlockDeviceStorageDeviceAttachmentDelegate protocol interface.
//
// Methods you implement to respond to changes to a network block device attachment.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.virtualization/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachmentDelegate
type PVZNetworkBlockDeviceStorageDeviceAttachmentDelegate interface {
	// Optional methods
	AttachmentDidEncounterError(attachment IVZNetworkBlockDeviceStorageDeviceAttachment, error_ objc.IObject /* cross-framework: Error */)
	HasAttachmentDidEncounterError() bool
	AttachmentWasConnected(attachment IVZNetworkBlockDeviceStorageDeviceAttachment)
	HasAttachmentWasConnected() bool
}

// VZNetworkBlockDeviceStorageDeviceAttachmentDelegate is a delegate implementation builder for the PVZNetworkBlockDeviceStorageDeviceAttachmentDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type VZNetworkBlockDeviceStorageDeviceAttachmentDelegate struct {
	_AttachmentDidEncounterError func(attachment IVZNetworkBlockDeviceStorageDeviceAttachment, error_ objc.IObject /* cross-framework: Error */)
	_AttachmentWasConnected      func(attachment IVZNetworkBlockDeviceStorageDeviceAttachment)
}

// SetAttachmentDidEncounterError sets the handler for the AttachmentDidEncounterError delegate method.
//
// The method the attachment object calls when the NBD client encounters a nonrecoverable error.
func (d *VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) SetAttachmentDidEncounterError(f func(attachment IVZNetworkBlockDeviceStorageDeviceAttachment, error_ objc.IObject /* cross-framework: Error */)) {
	d._AttachmentDidEncounterError = f
}

// SetAttachmentWasConnected sets the handler for the AttachmentWasConnected delegate method.
//
// The method the attachment object calls when the NBD client successfully connects or reconnects with the server.
func (d *VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) SetAttachmentWasConnected(f func(attachment IVZNetworkBlockDeviceStorageDeviceAttachment)) {
	d._AttachmentWasConnected = f
}

// AttachmentDidEncounterError implements the PVZNetworkBlockDeviceStorageDeviceAttachmentDelegate interface.
func (d *VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) AttachmentDidEncounterError(attachment IVZNetworkBlockDeviceStorageDeviceAttachment, error_ objc.IObject /* cross-framework: Error */) {
	if d._AttachmentDidEncounterError != nil {
		d._AttachmentDidEncounterError(attachment, error_)
	}
}

// HasAttachmentDidEncounterError returns true if a handler for AttachmentDidEncounterError has been set.
func (d *VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) HasAttachmentDidEncounterError() bool {
	return d._AttachmentDidEncounterError != nil
}

// AttachmentWasConnected implements the PVZNetworkBlockDeviceStorageDeviceAttachmentDelegate interface.
func (d *VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) AttachmentWasConnected(attachment IVZNetworkBlockDeviceStorageDeviceAttachment) {
	if d._AttachmentWasConnected != nil {
		d._AttachmentWasConnected(attachment)
	}
}

// HasAttachmentWasConnected returns true if a handler for AttachmentWasConnected has been set.
func (d *VZNetworkBlockDeviceStorageDeviceAttachmentDelegate) HasAttachmentWasConnected() bool {
	return d._AttachmentWasConnected != nil
}
