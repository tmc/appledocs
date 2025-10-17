// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WritingToolsCoordinator] class.
var WritingToolsCoordinatorClass objc.Class

func init() {
	WritingToolsCoordinatorClass = objc.GetClass("NSWritingToolsCoordinator")
}

type WritingToolsCoordinator struct {
	objc.ID
}

func WritingToolsCoordinatorFrom(ptr unsafe.Pointer) WritingToolsCoordinator {
	return WritingToolsCoordinator{
		ID: objc.ID(ptr),
	}
}


// Informs the coordinator that a change occurred to the view or its text   that requires a layout update. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w_ WritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer) {
	sel := objc.RegisterName("updateForReflowedTextInContextWithIdentifier:")
	w_.ID.Send(sel, contextID)
}
// Informs the coordinator about changes your app made to the text   in the specified context object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w_ WritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.Range, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer) {
	sel := objc.RegisterName("updateRange:withText:reason:forContextWithIdentifier:")
	w_.ID.Send(sel, range_, replacementText, reason, contextID)
}

