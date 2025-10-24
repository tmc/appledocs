// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PEditorRegistration is the NSEditorRegistration protocol interface.
//
// A set of methods that controllers can implement to enable an editor view to inform the controller when it has uncommitted changes.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSEditorRegistration
type PEditorRegistration interface {
	// Optional methods
	ObjectDidBeginEditing(editor objc.IObject)
	HasObjectDidBeginEditing() bool
	ObjectDidEndEditing(editor objc.IObject)
	HasObjectDidEndEditing() bool
}
