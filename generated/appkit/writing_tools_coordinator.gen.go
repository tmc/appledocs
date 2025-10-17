// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinator] class.
var writingToolsCoordinatorClass = _WritingToolsCoordinatorClass{objc.GetClass("NSWritingToolsCoordinator")}

type _WritingToolsCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [WritingToolsCoordinator] class.
type IWritingToolsCoordinator interface {
	objectivec.IObject
	UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer)
	UpdateRangeWithTextReasonForContextWithIdentifier(range_ unsafe.Pointer, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer)
}

// An object that manages interactions between Writing Tools and your custom text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator

type WritingToolsCoordinator struct {
	objectivec.Object
}

// WritingToolsCoordinatorFrom constructs a [WritingToolsCoordinator] from an unsafe.Pointer.
//
// An object that manages interactions between Writing Tools and your custom text view.
func WritingToolsCoordinatorFrom(ptr unsafe.Pointer) WritingToolsCoordinator {
	return WritingToolsCoordinator{objectivec.Object{objc.ID(ptr)}}
}

// Informs the coordinator that a change occurred to the view or its text that requires a layout update. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w_ WritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateForReflowedTextInContextWithIdentifier:"), contextID)
}
// Informs the coordinator about changes your app made to the text in the specified context object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w_ WritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ unsafe.Pointer, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateRange:withText:reason:forContextWithIdentifier:"), range_, replacementText, reason, contextID)
}


