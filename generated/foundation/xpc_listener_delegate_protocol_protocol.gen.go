// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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

// XPCListenerDelegateObject wraps an existing Objective-C object that conforms to the PXPCListenerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type XPCListenerDelegateObject struct {
	objectivec.Object
}

// NewXPCListenerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSXPCListenerDelegate protocol.
func NewXPCListenerDelegateObject(obj objectivec.Object) *XPCListenerDelegateObject {
	return &XPCListenerDelegateObject{obj}
}

// Make sure XPCListenerDelegateObject implements PXPCListenerDelegate.
var _ PXPCListenerDelegate = (*XPCListenerDelegateObject)(nil)

// ListenerShouldAcceptNewConnection implements the PXPCListenerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *XPCListenerDelegateObject) ListenerShouldAcceptNewConnection(listener IXPCListener, newConnection IXPCConnection) bool {
	return objc.Send[bool](o.ID, objc.Sel("listener:shouldAcceptNewConnection:"), listener, newConnection)
}

// HasListenerShouldAcceptNewConnection returns true; this is a placeholder for optional method checks.
func (o *XPCListenerDelegateObject) HasListenerShouldAcceptNewConnection() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
