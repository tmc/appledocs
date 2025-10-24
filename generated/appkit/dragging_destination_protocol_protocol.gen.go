// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
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
	ConcludeDragOperation(sender unsafe.Pointer)
	HasConcludeDragOperation() bool
	DraggingEnded(sender unsafe.Pointer)
	HasDraggingEnded() bool
	DraggingEntered(sender unsafe.Pointer) DragOperation
	HasDraggingEntered() bool
	DraggingExited(sender unsafe.Pointer)
	HasDraggingExited() bool
	DraggingUpdated(sender unsafe.Pointer) DragOperation
	HasDraggingUpdated() bool
	PerformDragOperation(sender unsafe.Pointer) bool
	HasPerformDragOperation() bool
	PrepareForDragOperation(sender unsafe.Pointer) bool
	HasPrepareForDragOperation() bool
	UpdateDraggingItemsForDrag(sender unsafe.Pointer)
	HasUpdateDraggingItemsForDrag() bool
	WantsPeriodicDraggingUpdates() bool
	HasWantsPeriodicDraggingUpdates() bool
}
