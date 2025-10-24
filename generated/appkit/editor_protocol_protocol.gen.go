// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

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
	CommitEditing() bool/* debug [protocol_interface/required_method]: CommitEditing */
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject)/* debug [protocol_interface/required_method]: CommitEditingWithDelegateDidCommitSelectorContextInfo */
	CommitEditingAndReturnError(error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: CommitEditingAndReturnError */
	DiscardEditing()/* debug [protocol_interface/required_method]: DiscardEditing */
}
