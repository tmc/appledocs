// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSpringLoadingDestination is the NSSpringLoadingDestination protocol interface.
//
// A set of methods that the destination object (or recipient) of a dragged object can implement to support spring-loading.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSpringLoadingDestination
type PSpringLoadingDestination interface {
	// Required methods
	SpringLoadingActivatedDraggingInfo(activated bool, draggingInfo objc.IObject)
	SpringLoadingHighlightChanged(draggingInfo objc.IObject)
	// Optional methods
	DraggingEnded(draggingInfo objc.IObject)
	HasDraggingEnded() bool
	SpringLoadingEntered(draggingInfo objc.IObject) SpringLoadingOptions
	HasSpringLoadingEntered() bool
	SpringLoadingExited(draggingInfo objc.IObject)
	HasSpringLoadingExited() bool
	SpringLoadingUpdated(draggingInfo objc.IObject) SpringLoadingOptions
	HasSpringLoadingUpdated() bool
}
