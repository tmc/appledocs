// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.SEL, contextInfo unsafe.Pointer)
	CommitEditingAndReturnError(error_ unsafe.Pointer) bool
	DiscardEditing()
}
