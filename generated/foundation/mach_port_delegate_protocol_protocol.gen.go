// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PMachPortDelegate is the NSMachPortDelegate protocol interface.
//
// An interface for handling incoming Mach messages.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSMachPortDelegate
type PMachPortDelegate interface {
	// Optional methods
	HandleMachMessage(msg objectivec.IObject)
	HasHandleMachMessage() bool
}

// MachPortDelegate is a delegate implementation builder for the PMachPortDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MachPortDelegate struct {
	_HandleMachMessage func(msg objectivec.IObject)
}

// SetHandleMachMessage sets the handler for the HandleMachMessage delegate method.
//
// Process an incoming Mach message.
func (d *MachPortDelegate) SetHandleMachMessage(f func(msg objectivec.IObject)) {
	d._HandleMachMessage = f
}

// HandleMachMessage implements the PMachPortDelegate interface.
func (d *MachPortDelegate) HandleMachMessage(msg objectivec.IObject) {
	if d._HandleMachMessage != nil {
		d._HandleMachMessage(msg)
	}
}

// HasHandleMachMessage returns true if a handler for HandleMachMessage has been set.
func (d *MachPortDelegate) HasHandleMachMessage() bool {
	return d._HandleMachMessage != nil
}
