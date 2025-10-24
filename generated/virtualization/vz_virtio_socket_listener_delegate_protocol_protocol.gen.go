// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

// PVZVirtioSocketListenerDelegate is the VZVirtioSocketListenerDelegate protocol interface.
//
// An interface you use to manage connections between the guest operating system and host computer.
//
// Availability:
//   - macOS 11.0+
//
// See: doc://com.apple.virtualization/documentation/Virtualization/VZVirtioSocketListenerDelegate
type PVZVirtioSocketListenerDelegate interface {
	// Optional methods
	ListenerShouldAcceptNewConnectionFromSocketDevice(listener IVZVirtioSocketListener, connection IVZVirtioSocketConnection, socketDevice IVZVirtioSocketDevice) bool
	HasListenerShouldAcceptNewConnectionFromSocketDevice() bool
}

// VZVirtioSocketListenerDelegate is a delegate implementation builder for the PVZVirtioSocketListenerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type VZVirtioSocketListenerDelegate struct {
	_ListenerShouldAcceptNewConnectionFromSocketDevice func(listener IVZVirtioSocketListener, connection IVZVirtioSocketConnection, socketDevice IVZVirtioSocketDevice) bool
}

// SetListenerShouldAcceptNewConnectionFromSocketDevice sets the handler for the ListenerShouldAcceptNewConnectionFromSocketDevice delegate method.
//
// Returns a Boolean value that indicates whether to accept a new connection from the guest operating system.
func (d *VZVirtioSocketListenerDelegate) SetListenerShouldAcceptNewConnectionFromSocketDevice(f func(listener IVZVirtioSocketListener, connection IVZVirtioSocketConnection, socketDevice IVZVirtioSocketDevice) bool) {
	d._ListenerShouldAcceptNewConnectionFromSocketDevice = f
}

// ListenerShouldAcceptNewConnectionFromSocketDevice implements the PVZVirtioSocketListenerDelegate interface.
func (d *VZVirtioSocketListenerDelegate) ListenerShouldAcceptNewConnectionFromSocketDevice(listener IVZVirtioSocketListener, connection IVZVirtioSocketConnection, socketDevice IVZVirtioSocketDevice) bool {
	if d._ListenerShouldAcceptNewConnectionFromSocketDevice != nil {
		return d._ListenerShouldAcceptNewConnectionFromSocketDevice(listener, connection, socketDevice)
	}
	var zero bool
	return zero
}

// HasListenerShouldAcceptNewConnectionFromSocketDevice returns true if a handler for ListenerShouldAcceptNewConnectionFromSocketDevice has been set.
func (d *VZVirtioSocketListenerDelegate) HasListenerShouldAcceptNewConnectionFromSocketDevice() bool {
	return d._ListenerShouldAcceptNewConnectionFromSocketDevice != nil
}
