// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PPortDelegate is the NSPortDelegate protocol interface.
//
// An interface for handling incoming messages.
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
// See: doc://com.apple.foundation/documentation/Foundation/PortDelegate
type PPortDelegate interface {
	// Optional methods
	HandlePortMessage(message IPortMessage)
	HasHandlePortMessage() bool
}

// PortDelegate is a delegate implementation builder for the PPortDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PortDelegate struct {
	_HandlePortMessage func(message IPortMessage)
}

// SetHandlePortMessage sets the handler for the HandlePortMessage delegate method.
//
// Processes a given incoming message on the port.
func (d *PortDelegate) SetHandlePortMessage(f func(message IPortMessage)) {
	d._HandlePortMessage = f
}

// HandlePortMessage implements the PPortDelegate interface.
func (d *PortDelegate) HandlePortMessage(message IPortMessage) {
	if d._HandlePortMessage != nil {
		d._HandlePortMessage(message)
	}
}

// HasHandlePortMessage returns true if a handler for HandlePortMessage has been set.
func (d *PortDelegate) HasHandlePortMessage() bool {
	return d._HandlePortMessage != nil
}

// PortDelegateObject wraps an existing Objective-C object that conforms to the PPortDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type PortDelegateObject struct {
	objectivec.Object
}

// NewPortDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSPortDelegate protocol.
func NewPortDelegateObject(obj objectivec.Object) *PortDelegateObject {
	return &PortDelegateObject{obj}
}

// Make sure PortDelegateObject implements PPortDelegate.
var _ PPortDelegate = (*PortDelegateObject)(nil)

// HandlePortMessage implements the PPortDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PortDelegateObject) HandlePortMessage(message IPortMessage) {
	objc.Send[objc.ID](o.ID, objc.Sel("handlePortMessage:"), message)
}

// HasHandlePortMessage returns true; this is a placeholder for optional method checks.
func (o *PortDelegateObject) HasHandlePortMessage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
