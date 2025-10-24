// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PVZVirtualMachineDelegate is the VZVirtualMachineDelegate protocol interface.
//
// The methods you use to respond to changes in the state of the VM.
//
// Availability:
//   - macOS 11.0+
//
// See: doc://com.apple.virtualization/documentation/Virtualization/VZVirtualMachineDelegate
type PVZVirtualMachineDelegate interface {
	// Optional methods
	GuestDidStopVirtualMachine(virtualMachine IVZVirtualMachine)
	HasGuestDidStopVirtualMachine() bool
	VirtualMachineDidStopWithError(virtualMachine IVZVirtualMachine, error_ objc.IObject /* cross-framework: Error */)
	HasVirtualMachineDidStopWithError() bool
	VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError(virtualMachine IVZVirtualMachine, networkDevice IVZNetworkDevice, error_ objc.IObject /* cross-framework: Error */)
	HasVirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError() bool
}

// VZVirtualMachineDelegate is a delegate implementation builder for the PVZVirtualMachineDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type VZVirtualMachineDelegate struct {
	_GuestDidStopVirtualMachine func(virtualMachine IVZVirtualMachine)
	_VirtualMachineDidStopWithError func(virtualMachine IVZVirtualMachine, error_ objc.IObject /* cross-framework: Error */)
	_VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError func(virtualMachine IVZVirtualMachine, networkDevice IVZNetworkDevice, error_ objc.IObject /* cross-framework: Error */)
}

// SetGuestDidStopVirtualMachine sets the handler for the GuestDidStopVirtualMachine delegate method.
//
// Tells the delegate that the guest operating system stopped the VM.
func (d *VZVirtualMachineDelegate) SetGuestDidStopVirtualMachine(f func(virtualMachine IVZVirtualMachine)) {
	d._GuestDidStopVirtualMachine = f
}

// SetVirtualMachineDidStopWithError sets the handler for the VirtualMachineDidStopWithError delegate method.
//
// Tells the delegate that the VM stopped because of an error.
func (d *VZVirtualMachineDelegate) SetVirtualMachineDidStopWithError(f func(virtualMachine IVZVirtualMachine, error_ objc.IObject /* cross-framework: Error */)) {
	d._VirtualMachineDidStopWithError = f
}

// SetVirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError sets the handler for the VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError delegate method.
//
// The method the framework calls when an error causes a VM’s network attachment to disconnect.
func (d *VZVirtualMachineDelegate) SetVirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError(f func(virtualMachine IVZVirtualMachine, networkDevice IVZNetworkDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError = f
}

// GuestDidStopVirtualMachine implements the PVZVirtualMachineDelegate interface.
func (d *VZVirtualMachineDelegate) GuestDidStopVirtualMachine(virtualMachine IVZVirtualMachine) {
	if d._GuestDidStopVirtualMachine != nil {
		d._GuestDidStopVirtualMachine(virtualMachine)
	}
}

// HasGuestDidStopVirtualMachine returns true if a handler for GuestDidStopVirtualMachine has been set.
func (d *VZVirtualMachineDelegate) HasGuestDidStopVirtualMachine() bool {
	return d._GuestDidStopVirtualMachine != nil
}

// VirtualMachineDidStopWithError implements the PVZVirtualMachineDelegate interface.
func (d *VZVirtualMachineDelegate) VirtualMachineDidStopWithError(virtualMachine IVZVirtualMachine, error_ objc.IObject /* cross-framework: Error */) {
	if d._VirtualMachineDidStopWithError != nil {
		d._VirtualMachineDidStopWithError(virtualMachine, error_)
	}
}

// HasVirtualMachineDidStopWithError returns true if a handler for VirtualMachineDidStopWithError has been set.
func (d *VZVirtualMachineDelegate) HasVirtualMachineDidStopWithError() bool {
	return d._VirtualMachineDidStopWithError != nil
}

// VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError implements the PVZVirtualMachineDelegate interface.
func (d *VZVirtualMachineDelegate) VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError(virtualMachine IVZVirtualMachine, networkDevice IVZNetworkDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError != nil {
		d._VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError(virtualMachine, networkDevice, error_)
	}
}

// HasVirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError returns true if a handler for VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError has been set.
func (d *VZVirtualMachineDelegate) HasVirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError() bool {
	return d._VirtualMachineNetworkDeviceAttachmentWasDisconnectedWithError != nil
}
