// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [WritingToolsCoordinator] class.
var (
	writingToolsCoordinatorClass     _WritingToolsCoordinatorClass
	writingToolsCoordinatorClassOnce sync.Once
)

func getWritingToolsCoordinatorClass() _WritingToolsCoordinatorClass {
	writingToolsCoordinatorClassOnce.Do(func() {
		writingToolsCoordinatorClass = _WritingToolsCoordinatorClass{objc.GetClass("NSWritingToolsCoordinator")}
	})
	return writingToolsCoordinatorClass
}

type _WritingToolsCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [WritingToolsCoordinator] class.
type IWritingToolsCoordinator interface {
	objectivec.IObject
	UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer)
	UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.Range, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer)
}

// An object that manages interactions between Writing Tools and your custom text view.
//
// Add a object to a custom view when you want to add Writing Tools support to that view. The coordinator manages interactions between your view and the Writing Tools UI and back-end capabilities. When creating a coordinator, you supply a delegate object to respond to requests from the system and provide needed information. Your delegate delivers your view’s text to Writing Tools, incorporates suggested changes back into your text storage, and supports the animations that Writing Tools creates to show the state of an operation. Create the object when setting up your UI, and initialize it with a custom object that adopts the protocol. Add the coordinator to the property of your view. When a coordinator is present on a view, the system adds UI elements to initiate Writing Tools operations. When defining the delegate, choose an object from your app that has access to your view and its text storage. You can adopt the protocol in the view itself, or in another type that your view uses to manage content. During the interactions with Writing Tools, the delegate gets and sets the contents of the view’s text storage and supports Writing Tools behaviors.
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

// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorClass) Alloc() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WritingToolsCoordinatorClass) New() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WritingToolsCoordinator) Init() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WritingToolsCoordinator) Autorelease() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWritingToolsCoordinator creates a new WritingToolsCoordinator instance.
func NewWritingToolsCoordinator() WritingToolsCoordinator {
	return getWritingToolsCoordinatorClass().New()
}


// Informs the coordinator that a change occurred to the view or its text that requires a layout update.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w_ WritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateForReflowedTextInContextWithIdentifier:"), contextID)
}

// Informs the coordinator about changes your app made to the text in the specified context object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w_ WritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.Range, replacementText unsafe.Pointer, reason unsafe.Pointer, contextID unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateRange:withText:reason:forContextWithIdentifier:"), range_, replacementText, reason, contextID)
}

// The level of Writing Tools support you want the system to provide for your view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) PreferredBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferredBehavior"))
	return rv
}


// SetPreferredBehavior sets the value of the preferredBehavior property.
// The level of Writing Tools support you want the system to provide for your view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) SetPreferredBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredBehavior:"), value)
}
// The type of content you allow Writing Tools to generate for your custom text view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) PreferredResultOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferredResultOptions"))
	return rv
}


// SetPreferredResultOptions sets the value of the preferredResultOptions property.
// The type of content you allow Writing Tools to generate for your custom text view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) SetPreferredResultOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredResultOptions:"), value)
}


