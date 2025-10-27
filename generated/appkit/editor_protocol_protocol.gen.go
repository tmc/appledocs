// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PEditor is the NSEditor protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSEditor
type PEditor interface {
	// Required methods
	CommitEditing() bool
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject)
	CommitEditingAndReturnError(error_ foundation.foundation.INSError) bool
	DiscardEditing()
}
