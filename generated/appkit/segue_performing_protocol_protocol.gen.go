// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSeguePerforming is the NSSeguePerforming protocol interface.
//
// A set of methods that support the mediation of a custom segue.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSeguePerforming
type PSeguePerforming interface {
	// Optional methods
	PerformSegueWithIdentifierSender(identifier objc.IObject /* cross-framework: StoryboardSegueIdentifier */, sender objc.IObject)
	HasPerformSegueWithIdentifierSender() bool
	PrepareForSegueSender(segue IStoryboardSegue, sender objc.IObject)
	HasPrepareForSegueSender() bool
	ShouldPerformSegueWithIdentifierSender(identifier objc.IObject /* cross-framework: StoryboardSegueIdentifier */, sender objc.IObject) bool
	HasShouldPerformSegueWithIdentifierSender() bool
}
