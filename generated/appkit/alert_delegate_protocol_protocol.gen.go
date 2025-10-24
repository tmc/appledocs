// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAlertDelegate is the NSAlertDelegate protocol interface.
//
// A set of optional methods implemented by the delegate of an   object to respond to a user’s request for help.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAlertDelegate
type PAlertDelegate interface {
	// Optional methods
	AlertShowHelp(alert IAlert) bool
	HasAlertShowHelp() bool
}

// AlertDelegate is a delegate implementation builder for the PAlertDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AlertDelegate struct {
	_AlertShowHelp func(alert IAlert) bool
}

// SetAlertShowHelp sets the handler for the AlertShowHelp delegate method.
//
// Sent to the delegate when the user clicks the alert’s help button. The delegate causes help to be displayed for an alert, directly or indirectly.
func (d *AlertDelegate) SetAlertShowHelp(f func(alert IAlert) bool) {
	d._AlertShowHelp = f
}

// AlertShowHelp implements the PAlertDelegate interface.
func (d *AlertDelegate) AlertShowHelp(alert IAlert) bool {
	if d._AlertShowHelp != nil {
		return d._AlertShowHelp(alert)
	}
	var zero bool
	return zero
}

// HasAlertShowHelp returns true if a handler for AlertShowHelp has been set.
func (d *AlertDelegate) HasAlertShowHelp() bool {
	return d._AlertShowHelp != nil
}
