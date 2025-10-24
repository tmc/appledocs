// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PXPCListenerDelegate is the NSXPCListenerDelegate protocol interface.
//
// The protocol that delegates to the XPC listener use to accept or reject new connections.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSXPCListenerDelegate
type PXPCListenerDelegate interface {
	// Optional methods
	ListenerShouldAcceptNewConnection(listener IXPCListener, newConnection IXPCConnection) bool
	HasListenerShouldAcceptNewConnection() bool
}

// XPCListenerDelegate is a delegate implementation builder for the PXPCListenerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type XPCListenerDelegate struct {
	_ListenerShouldAcceptNewConnection func(listener IXPCListener, newConnection IXPCConnection) bool
}

// SetListenerShouldAcceptNewConnection sets the handler for the ListenerShouldAcceptNewConnection delegate method.
//
// Accepts or rejects a new connection to the listener.
func (d *XPCListenerDelegate) SetListenerShouldAcceptNewConnection(f func(listener IXPCListener, newConnection IXPCConnection) bool) {
	d._ListenerShouldAcceptNewConnection = f
}

// ListenerShouldAcceptNewConnection implements the PXPCListenerDelegate interface.
func (d *XPCListenerDelegate) ListenerShouldAcceptNewConnection(listener IXPCListener, newConnection IXPCConnection) bool {
	if d._ListenerShouldAcceptNewConnection != nil {
		return d._ListenerShouldAcceptNewConnection(listener, newConnection)
	}
	var zero bool
	return zero
}

// HasListenerShouldAcceptNewConnection returns true if a handler for ListenerShouldAcceptNewConnection has been set.
func (d *XPCListenerDelegate) HasListenerShouldAcceptNewConnection() bool {
	return d._ListenerShouldAcceptNewConnection != nil
}
