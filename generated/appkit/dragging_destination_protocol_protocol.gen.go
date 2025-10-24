// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PDraggingDestination is the NSDraggingDestination protocol interface.
//
// A set of methods that the destination object (or recipient) of a dragged image must implement.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSDraggingDestination
type PDraggingDestination interface {
	// Optional methods
	ConcludeDragOperation(sender objc.IObject)
	HasConcludeDragOperation() bool
	DraggingEnded(sender objc.IObject)
	HasDraggingEnded() bool
	DraggingEntered(sender objc.IObject) DragOperation
	HasDraggingEntered() bool
	DraggingExited(sender objc.IObject)
	HasDraggingExited() bool
	DraggingUpdated(sender objc.IObject) DragOperation
	HasDraggingUpdated() bool
	PerformDragOperation(sender objc.IObject) bool
	HasPerformDragOperation() bool
	PrepareForDragOperation(sender objc.IObject) bool
	HasPrepareForDragOperation() bool
	UpdateDraggingItemsForDrag(sender objc.IObject)
	HasUpdateDraggingItemsForDrag() bool
	WantsPeriodicDraggingUpdates() bool
	HasWantsPeriodicDraggingUpdates() bool
}
