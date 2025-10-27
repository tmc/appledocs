// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
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
	SpringLoadingActivatedDraggingInfo(activated bool, draggingInfo unsafe.Pointer)
	SpringLoadingHighlightChanged(draggingInfo unsafe.Pointer)
	// Optional methods
	DraggingEnded(draggingInfo unsafe.Pointer)
	HasDraggingEnded() bool
	SpringLoadingEntered(draggingInfo unsafe.Pointer) SpringLoadingOptions
	HasSpringLoadingEntered() bool
	SpringLoadingExited(draggingInfo unsafe.Pointer)
	HasSpringLoadingExited() bool
	SpringLoadingUpdated(draggingInfo unsafe.Pointer) SpringLoadingOptions
	HasSpringLoadingUpdated() bool
}
