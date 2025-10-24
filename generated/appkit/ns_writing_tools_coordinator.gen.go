// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSWritingToolsCoordinator */


/* debug [class_header]: Header for NSWritingToolsCoordinator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WritingToolsCoordinator */
// An interface definition for the [WritingToolsCoordinator] class.
type IWritingToolsCoordinator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WritingToolsCoordinator */
	// properties:
	Behavior() WritingToolsBehavior
	DecorationContainerView() IView
	SetDecorationContainerView(value IView)
	Delegate() unsafe.Pointer
	EffectContainerView() IView
	SetEffectContainerView(value IView)
	IncludesTextListMarkers() bool
	SetIncludesTextListMarkers(value bool)
	PreferredBehavior() WritingToolsBehavior
	SetPreferredBehavior(value WritingToolsBehavior)
	PreferredResultOptions() WritingToolsResultOptions
	SetPreferredResultOptions(value WritingToolsResultOptions)
	ResultOptions() WritingToolsResultOptions
	State() WritingToolsCoordinatorState
	View() IView
	WritingToolsCoordinator() IWritingToolsCoordinator
	SetWritingToolsCoordinator(value IWritingToolsCoordinator)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WritingToolsCoordinator */
	// methods:
	StopWritingTools()
	UpdateForReflowedTextInContextWithIdentifier(contextID foundation.UUID)
	UpdateRangeWithTextReasonForContextWithIdentifier(range_ corefoundation.Range, replacementText foundation.AttributedString, reason WritingToolsCoordinatorTextUpdateReason, contextID foundation.UUID)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WritingToolsCoordinator */
// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorClass) Alloc() WritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WritingToolsCoordinator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WritingToolsCoordinator */

// Creates a writing tools coordinator and assigns the specified delegate object to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/init(delegate:)
func NewWritingToolsCoordinatorWithDelegate(delegate unsafe.Pointer) WritingToolsCoordinator {
	instance := getWritingToolsCoordinatorClass().Alloc()
	rv := objc.Send[WritingToolsCoordinator](instance.ID, objc.Sel("initWithDelegate:"), delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWritingToolsCoordinatorWithDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WritingToolsCoordinator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WritingToolsCoordinator */

// A Boolean value that indicates whether Writing Tools features are currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/isWritingToolsAvailable
func (wc _WritingToolsCoordinatorClass) IsWritingToolsAvailable() bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("isWritingToolsAvailable"))
	return rv
}/* debug [class_properties_class/property]: isWritingToolsAvailable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WritingToolsCoordinator */

// Stops the current Writing Tools operation and dismisses the system UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/stopWritingTools()
func (w_ WritingToolsCoordinator) StopWritingTools() {
	objc.Send[objc.ID](w_.ID, objc.Sel("stopWritingTools"))
}/* debug [instance_methods/method]: StopWritingTools */


// Informs the coordinator that a change occurred to the view or its text that requires a layout update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w_ WritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID foundation.UUID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateForReflowedTextInContextWithIdentifier:"), contextID)
}/* debug [instance_methods/method]: UpdateForReflowedTextInContextWithIdentifier */


// Informs the coordinator about changes your app made to the text in the specified context object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w_ WritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ corefoundation.Range, replacementText foundation.AttributedString, reason WritingToolsCoordinatorTextUpdateReason, contextID foundation.UUID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateRange:withText:reason:forContextWithIdentifier:"), range_, replacementText, reason, contextID)
}/* debug [instance_methods/method]: UpdateRangeWithTextReasonForContextWithIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WritingToolsCoordinator */

// The actual level of Writing Tools support the system provides for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/behavior
func (w_ WritingToolsCoordinator) Behavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](w_.ID, objc.Sel("behavior"))
	return rv
}/* debug [instance_properties/getter]: behavior */


// The view that Writing Tools uses to display background decorations such as proofreading marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/decorationContainerView
func (w_ WritingToolsCoordinator) DecorationContainerView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("decorationContainerView"))
	return rv
}/* debug [instance_properties/getter]: decorationContainerView */


// The view that Writing Tools uses to display background decorations such as proofreading marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/decorationContainerView
func (w_ WritingToolsCoordinator) SetDecorationContainerView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDecorationContainerView:"), value)
}/* debug [instance_properties/setter]: decorationContainerView */


// The object that handles Writing Tools interactions for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/delegate-swift.property
func (w_ WritingToolsCoordinator) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The view that Writing Tools uses to display visual effects during the text-rewriting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/effectContainerView
func (w_ WritingToolsCoordinator) EffectContainerView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("effectContainerView"))
	return rv
}/* debug [instance_properties/getter]: effectContainerView */


// The view that Writing Tools uses to display visual effects during the text-rewriting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/effectContainerView
func (w_ WritingToolsCoordinator) SetEffectContainerView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setEffectContainerView:"), value)
}/* debug [instance_properties/setter]: effectContainerView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/includesTextListMarkers
func (w_ WritingToolsCoordinator) IncludesTextListMarkers() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("includesTextListMarkers"))
	return rv
}/* debug [instance_properties/getter]: includesTextListMarkers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/includesTextListMarkers
func (w_ WritingToolsCoordinator) SetIncludesTextListMarkers(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIncludesTextListMarkers:"), value)
}/* debug [instance_properties/setter]: includesTextListMarkers */


// A Boolean value that indicates whether Writing Tools features are currently available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/isWritingToolsAvailable
func (w_ WritingToolsCoordinator) IsWritingToolsAvailable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isWritingToolsAvailable"))
	return rv
}/* debug [instance_properties/getter]: isWritingToolsAvailable */


// The level of Writing Tools support you want the system to provide for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) PreferredBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](w_.ID, objc.Sel("preferredBehavior"))
	return rv
}/* debug [instance_properties/getter]: preferredBehavior */


// The level of Writing Tools support you want the system to provide for your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w_ WritingToolsCoordinator) SetPreferredBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredBehavior:"), value)
}/* debug [instance_properties/setter]: preferredBehavior */


// The type of content you allow Writing Tools to generate for your custom text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) PreferredResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](w_.ID, objc.Sel("preferredResultOptions"))
	return rv
}/* debug [instance_properties/getter]: preferredResultOptions */


// The type of content you allow Writing Tools to generate for your custom text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w_ WritingToolsCoordinator) SetPreferredResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredResultOptions:"), value)
}/* debug [instance_properties/setter]: preferredResultOptions */


// The type of content the system generates for your custom text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/resultOptions
func (w_ WritingToolsCoordinator) ResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](w_.ID, objc.Sel("resultOptions"))
	return rv
}/* debug [instance_properties/getter]: resultOptions */


// The current level of Writing Tools activity in your view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/state-swift.property
func (w_ WritingToolsCoordinator) State() WritingToolsCoordinatorState {
	rv := objc.Send[WritingToolsCoordinatorState](w_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The view that currently uses the writing tools coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/view
func (w_ WritingToolsCoordinator) View() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (w_ WritingToolsCoordinator) WritingToolsCoordinator() IWritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](w_.ID, objc.Sel("writingToolsCoordinator"))
	return rv
}/* debug [instance_properties/getter]: writingToolsCoordinator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (w_ WritingToolsCoordinator) SetWritingToolsCoordinator(value IWritingToolsCoordinator) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}/* debug [instance_properties/setter]: writingToolsCoordinator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSWritingToolsCoordinator */


