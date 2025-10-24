// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PSBApplicationDelegate is the SBApplicationDelegate protocol interface.
//
// This informal protocol defines a delegation method for handling Apple event   errors that are sent from a target application to an     object.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.5+
//
// See: doc://com.apple.Scripting-Bridge/documentation/ScriptingBridge/SBApplicationDelegate
type PSBApplicationDelegate interface {
	// Required methods
	EventDidFailWithError(event unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */) objc.ID/* debug [protocol_interface/required_method]: EventDidFailWithError */
}

// SBApplicationDelegate is a delegate implementation builder for the PSBApplicationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SBApplicationDelegate struct {
	_EventDidFailWithError func(event unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */) objc.ID
}

// SetEventDidFailWithError sets the handler for the EventDidFailWithError delegate method.
//
// Sent by an   object when a target application returns an error   Apple event.
func (d *SBApplicationDelegate) SetEventDidFailWithError(f func(event unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */) objc.ID) {
	d._EventDidFailWithError = f
}

// EventDidFailWithError implements the PSBApplicationDelegate interface.
func (d *SBApplicationDelegate) EventDidFailWithError(event unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */) objc.ID {
	if d._EventDidFailWithError != nil {
		return d._EventDidFailWithError(event, error_)
	}
	var zero objc.ID
	return zero
}

// HasEventDidFailWithError returns true if a handler for EventDidFailWithError has been set.
func (d *SBApplicationDelegate) HasEventDidFailWithError() bool {
	return d._EventDidFailWithError != nil
}
