// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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
	PerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objectivec.IObject)
	HasPerformSegueWithIdentifierSender() bool
	PrepareForSegueSender(segue IStoryboardSegue, sender objectivec.IObject)
	HasPrepareForSegueSender() bool
	ShouldPerformSegueWithIdentifierSender(identifier StoryboardSegueIdentifier, sender objectivec.IObject) bool
	HasShouldPerformSegueWithIdentifierSender() bool
}
