// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMEContentBlocker is the MEContentBlocker protocol interface.
//
// An object that provides a set of rules to block content when displaying a message.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEContentBlocker
type PMEContentBlocker interface {
	// Required methods
	ContentRulesJSON() foundation.Data/* debug [protocol_interface/required_method]: ContentRulesJSON */
}
