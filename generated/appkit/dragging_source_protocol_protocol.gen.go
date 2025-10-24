// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PDraggingSource is the NSDraggingSource protocol interface.
//
// A set of methods that are implemented by the source object in a dragging session.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSDraggingSource
type PDraggingSource interface {
	// Required methods
	DraggingSessionSourceOperationMaskForDraggingContext(session IDraggingSession, context DraggingContext) DragOperation
	// Optional methods
	DraggingSessionEndedAtPointOperation(session IDraggingSession, screenPoint objc.IObject /* cross-framework: Point */, operation DragOperation)
	HasDraggingSessionEndedAtPointOperation() bool
	DraggingSessionMovedToPoint(session IDraggingSession, screenPoint objc.IObject /* cross-framework: Point */)
	HasDraggingSessionMovedToPoint() bool
	DraggingSessionWillBeginAtPoint(session IDraggingSession, screenPoint objc.IObject /* cross-framework: Point */)
	HasDraggingSessionWillBeginAtPoint() bool
	IgnoreModifierKeysForDraggingSession(session IDraggingSession) bool
	HasIgnoreModifierKeysForDraggingSession() bool
}
