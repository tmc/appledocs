
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [WritingToolsCoordinator] class.
var WritingToolsCoordinatorClass _WritingToolsCoordinatorClass

func init() {
	WritingToolsCoordinatorClass = _WritingToolsCoordinatorClass{objc.GetClass("NSWritingToolsCoordinator")}
}

type _WritingToolsCoordinatorClass struct {
	objc.Class
}

// An interface definition for the [WritingToolsCoordinator] class.
type IWritingToolsCoordinator interface {
	ID() objc.ID
	UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer)
	UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.Range, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer)
}

type WritingToolsCoordinator struct {
	id objc.ID
}

func WritingToolsCoordinatorFrom(ptr unsafe.Pointer) WritingToolsCoordinator {
	return WritingToolsCoordinator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WritingToolsCoordinator) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorClass) Alloc() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WritingToolsCoordinatorClass) New() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWritingToolsCoordinator creates and returns a new initialized instance.
func NewWritingToolsCoordinator() WritingToolsCoordinator {
	return WritingToolsCoordinatorClass.New()
}

// Init initializes the instance.
func (w_ WritingToolsCoordinator) Init() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](w_.ID(), selInit)
	return rv
}
// Informs the coordinator that a change occurred to the view or its text   that requires a layout update. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w_ WritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("updateForReflowedTextInContextWithIdentifier:"), contextID)
}
// Informs the coordinator about changes your app made to the text   in the specified context object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w_ WritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.Range, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("updateRange:withText:reason:forContextWithIdentifier:"), range_, replacementText, reason, contextID)
}
// The level of Writing Tools support you want the system to provide   for your view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) PreferredBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("preferredBehavior"))
	return rv
}
// SetPreferredBehavior sets the value of the preferredBehavior property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) SetPreferredBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setPreferredBehavior:"), value)
}
// The type of content you allow Writing Tools to generate for your custom   text view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) PreferredResultOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("preferredResultOptions"))
	return rv
}
// SetPreferredResultOptions sets the value of the preferredResultOptions property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) SetPreferredResultOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setPreferredResultOptions:"), value)
}
