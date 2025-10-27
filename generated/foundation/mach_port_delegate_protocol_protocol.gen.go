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

// MachPortDelegateObject wraps an existing Objective-C object that conforms to the PMachPortDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type MachPortDelegateObject struct {
	objectivec.Object
}

// NewMachPortDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSMachPortDelegate protocol.
func NewMachPortDelegateObject(obj objectivec.Object) *MachPortDelegateObject {
	return &MachPortDelegateObject{obj}
}

// Make sure MachPortDelegateObject implements PMachPortDelegate.
var _ PMachPortDelegate = (*MachPortDelegateObject)(nil)

// HandleMachMessage implements the PMachPortDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *MachPortDelegateObject) HandleMachMessage(msg objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("handleMachMessage:"), msg)
}

// HasHandleMachMessage returns true; this is a placeholder for optional method checks.
func (o *MachPortDelegateObject) HasHandleMachMessage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
