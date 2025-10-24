// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TouchBarItem] class.
var (
	TouchBarItemClass     _TouchBarItemClass
	TouchBarItemClassOnce sync.Once
)

func getTouchBarItemClass() _TouchBarItemClass {
	TouchBarItemClassOnce.Do(func() {
		TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}
	})
	return TouchBarItemClass
}

type _TouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [TouchBarItem] class.
type ITouchBarItem interface {
	objectivec.IObject
	// properties:
	IsContinuous() bool
	SetIsContinuous(value bool)
	GroupTouchBar() ITouchBar
	SetGroupTouchBar(value ITouchBar)
	CollapsedRepresentation() IView
	SetCollapsedRepresentation(value IView)
	PopoverTouchBar() ITouchBar
	SetPopoverTouchBar(value ITouchBar)
	PressAndHoldTouchBar() ITouchBar
	SetPressAndHoldTouchBar(value ITouchBar)
	TrackingMode() unsafe.Pointer
	SetTrackingMode(value unsafe.Pointer)
	PrincipalItemIdentifier() unsafe.Pointer
	SetPrincipalItemIdentifier(value unsafe.Pointer)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Identifier() unsafe.Pointer
	SetIdentifier(value unsafe.Pointer)
	IsVisible() bool
	SetIsVisible(value bool)
	View() IView
	SetView(value IView)
	ViewController() objc.IObject /* cross-framework: ViewController */
	SetViewController(value objc.IObject /* cross-framework: ViewController */)
	VisibilityPriority() unsafe.Pointer
	SetVisibilityPriority(value unsafe.Pointer)
	// methods:
}

// A UI control shown in the Touch Bar on supported models of MacBook Pro.
//
// An instance of the class is called an . It appears to the user on the Touch Bar, typically along with other items, within the (invisible) bounds of the view for an object, called a . You use an item by adding it or its identifier to one or another of a bar’s arrays, depending on your app’s architecture and on the user customization you want to support. Because of the close interaction between bars and items, be sure you have read the overview for the class before continuing here to learn about items. AppKit provides a rich set of subclasses of , each of which is described in the corresponding class reference document: An object (a ), along with its delegate, provides a list of textual suggestions for the current text view An object (a ) provides a system-defined color picker An object (a ) contains a responder of your choice, such as a view, a button, or a scrubber (an instance of the class) An object (a ) provides a bar to contain other items An object (a ) provides a two-state control that, when touched or pressed, expands into its second state, showing the contents of a bar it owns An object (a ), along with its delegate, provides a list of objects eligible for sharing An object (a ) provides a slider control for choosing a value in a range The two most commonly-used item classes are and . Refer to the following sample code projects which demonstrate how to use and related classes:


// A UI control shown in the Touch Bar on supported models of MacBook Pro.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem
type TouchBarItem struct {
	objectivec.Object
}

// TouchBarItemFrom constructs a [TouchBarItem] from an unsafe.Pointer.
//
// A UI control shown in the Touch Bar on supported models of MacBook Pro.
func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TouchBarItem) Autorelease() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouchBarItem creates a new TouchBarItem instance.
func NewTouchBarItem() TouchBarItem {
	return getTouchBarItemClass().New()
}



// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t_ TouchBarItem) IsContinuous() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuous"))
	return rv
}


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t_ TouchBarItem) SetIsContinuous(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuous:"), value)
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) GroupTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) SetGroupTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupTouchBar:"), value)
}


// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t_ TouchBarItem) CollapsedRepresentation() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("collapsedRepresentation"))
	return rv
}


// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t_ TouchBarItem) SetCollapsedRepresentation(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCollapsedRepresentation:"), value)
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) PopoverTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) SetPopoverTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPopoverTouchBar:"), value)
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) PressAndHoldTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) SetPressAndHoldTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) TrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("trackingMode"))
	return rv
}


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) SetTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackingMode:"), value)
}


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) PrincipalItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("principalItemIdentifier"))
	return rv
}


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) SetPrincipalItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrincipalItemIdentifier:"), value)
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t_ TouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t_ TouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The identifier for this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/identifier-swift.property
func (t_ TouchBarItem) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier for this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/identifier-swift.property
func (t_ TouchBarItem) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}


// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/isvisible
func (t_ TouchBarItem) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/isvisible
func (t_ TouchBarItem) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}


// The view associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/view
func (t_ TouchBarItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}


// The view associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/view
func (t_ TouchBarItem) SetView(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setView:"), value)
}


// The view controller associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/viewcontroller
func (t_ TouchBarItem) ViewController() objc.IObject /* cross-framework: ViewController */ {
	rv := objc.Send[ViewController](t_.ID, objc.Sel("viewController"))
	return rv
}


// The view controller associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/viewcontroller
func (t_ TouchBarItem) SetViewController(value objc.IObject /* cross-framework: ViewController */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewController:"), value)
}


// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/visibilitypriority
func (t_ TouchBarItem) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/visibilitypriority
func (t_ TouchBarItem) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}



