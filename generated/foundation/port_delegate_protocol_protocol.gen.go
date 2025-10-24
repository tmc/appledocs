// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
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
