// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinator] class.
var (
	WritingToolsCoordinatorClass     _WritingToolsCoordinatorClass
	WritingToolsCoordinatorClassOnce sync.Once
)

func getWritingToolsCoordinatorClass() _WritingToolsCoordinatorClass {
	WritingToolsCoordinatorClassOnce.Do(func() {
		WritingToolsCoordinatorClass = _WritingToolsCoordinatorClass{objc.GetClass("NSWritingToolsCoordinator")}
	})
	return WritingToolsCoordinatorClass
}

type _WritingToolsCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [WritingToolsCoordinator] class.
type IWritingToolsCoordinator interface {
	objectivec.IObject
	// properties:
	Behavior() WritingToolsBehavior
	DecorationContainerView() IView
	SetDecorationContainerView(value IView)
	Delegate() objc.ID
	EffectContainerView() IView
	SetEffectContainerView(value IView)
	IncludesTextListMarkers() bool /* primitive/slice/pointer. */
	SetIncludesTextListMarkers(value bool /* primitive/slice/pointer. */)
	PreferredBehavior() WritingToolsBehavior
	SetPreferredBehavior(value WritingToolsBehavior)
	PreferredResultOptions() WritingToolsResultOptions
	SetPreferredResultOptions(value WritingToolsResultOptions)
	ResultOptions() WritingToolsResultOptions
	State() WritingToolsCoordinatorState
	View() IView
	WritingToolsCoordinator() IWritingToolsCoordinator
	SetWritingToolsCoordinator(value IWritingToolsCoordinator)
	// methods:
	StopWritingTools()
	UpdateForReflowedTextInContextWithIdentifier(contextID objc.IObject /* cross-framework UUID */)
	UpdateRangeWithTextReasonForContextWithIdentifier(range_ objc.IObject /* cross-framework Range */, replacementText objc.IObject /* cross-framework AttributedString */, reason WritingToolsCoordinatorTextUpdateReason, contextID objc.IObject /* cross-framework UUID */)
}

// An object that manages interactions between Writing Tools and your custom text view.
//
// Add a object to a custom view when you want to add Writing Tools support to that view. The coordinator manages interactions between your view and the Writing Tools UI and back-end capabilities. When creating a coordinator, you supply a delegate object to respond to requests from the system and provide needed information. Your delegate delivers your view’s text to Writing Tools, incorporates suggested changes back into your text storage, and supports the animations that Writing Tools creates to show the state of an operation. Create the object when setting up your UI, and initialize it with a custom object that adopts the protocol. Add the coordinator to the property of your view. When a coordinator is present on a view, the system adds UI elements to initiate Writing Tools operations. When defining the delegate, choose an object from your app that has access to your view and its text storage. You can adopt the protocol in the view itself, or in another type that your view uses to manage content. During the interactions with Writing Tools, the delegate gets and sets the contents of the view’s text storage and supports Writing Tools behaviors.


// An object that manages interactions between Writing Tools and your custom text view.
//
// [Full Topic]
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



// Creates a writing tools coordinator and assigns the specified delegate object to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/init(delegate:)
func NewWritingToolsCoordinatorWithDelegate(delegate objectivec.IObject) WritingToolsCoordinator {
	instance := getWritingToolsCoordinatorClass().Alloc()
	rv := objc.Send[WritingToolsCoordinator](instance.ID, objc.Sel("initWithDelegate:"), delegate)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates whether Writing Tools features are currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/isWritingToolsAvailable
func (wc _WritingToolsCoordinatorClass) IsWritingToolsAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("isWritingToolsAvailable"))
	return rv
}

// Stops the current Writing Tools operation and dismisses the system UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/stopWritingTools()
func (w_ WritingToolsCoordinator) StopWritingTools() {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopWritingTools"))
}


// Informs the coordinator that a change occurred to the view or its text that requires a layout update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w_ WritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID objc.IObject /* cross-framework UUID */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateForReflowedTextInContextWithIdentifier:"), contextID)
}


// Informs the coordinator about changes your app made to the text in the specified context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w_ WritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ objc.IObject /* cross-framework Range */, replacementText objc.IObject /* cross-framework AttributedString */, reason WritingToolsCoordinatorTextUpdateReason, contextID objc.IObject /* cross-framework UUID */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateRange:withText:reason:forContextWithIdentifier:"), range_, replacementText, reason, contextID)
}


// The actual level of Writing Tools support the system provides for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/behavior
func (w_ WritingToolsCoordinator) Behavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](w_.ID, objc.Sel("behavior"))
	return rv
}


// The view that Writing Tools uses to display background decorations such as proofreading marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/decorationContainerView
func (w_ WritingToolsCoordinator) DecorationContainerView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("decorationContainerView"))
	return rv
}


// The view that Writing Tools uses to display background decorations such as proofreading marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/decorationContainerView
func (w_ WritingToolsCoordinator) SetDecorationContainerView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDecorationContainerView:"), value)
}


// The object that handles Writing Tools interactions for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/delegate-swift.property
func (w_ WritingToolsCoordinator) Delegate() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("delegate"))
	return rv
}


// The view that Writing Tools uses to display visual effects during the text-rewriting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/effectContainerView
func (w_ WritingToolsCoordinator) EffectContainerView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("effectContainerView"))
	return rv
}


// The view that Writing Tools uses to display visual effects during the text-rewriting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/effectContainerView
func (w_ WritingToolsCoordinator) SetEffectContainerView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEffectContainerView:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/includesTextListMarkers
func (w_ WritingToolsCoordinator) IncludesTextListMarkers() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("includesTextListMarkers"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/includesTextListMarkers
func (w_ WritingToolsCoordinator) SetIncludesTextListMarkers(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIncludesTextListMarkers:"), value)
}


// A Boolean value that indicates whether Writing Tools features are currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/isWritingToolsAvailable
func (w_ WritingToolsCoordinator) IsWritingToolsAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isWritingToolsAvailable"))
	return rv
}


// The level of Writing Tools support you want the system to provide for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) PreferredBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](w_.ID, objc.Sel("preferredBehavior"))
	return rv
}


// The level of Writing Tools support you want the system to provide for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) SetPreferredBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredBehavior:"), value)
}


// The type of content you allow Writing Tools to generate for your custom text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) PreferredResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](w_.ID, objc.Sel("preferredResultOptions"))
	return rv
}


// The type of content you allow Writing Tools to generate for your custom text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) SetPreferredResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredResultOptions:"), value)
}


// The type of content the system generates for your custom text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/resultOptions
func (w_ WritingToolsCoordinator) ResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](w_.ID, objc.Sel("resultOptions"))
	return rv
}


// The current level of Writing Tools activity in your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/state-swift.property
func (w_ WritingToolsCoordinator) State() WritingToolsCoordinatorState {
	rv := objc.Send[WritingToolsCoordinatorState](w_.ID, objc.Sel("state"))
	return rv
}


// The view that currently uses the writing tools coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/view
func (w_ WritingToolsCoordinator) View() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("view"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (w_ WritingToolsCoordinator) WritingToolsCoordinator() IWritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](w_.ID, objc.Sel("writingToolsCoordinator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (w_ WritingToolsCoordinator) SetWritingToolsCoordinator(value IWritingToolsCoordinator) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}


