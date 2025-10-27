// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	CustomizationLabel() foundation.foundation.INSString
	Identifier() TouchBarItemIdentifier
	Visible() bool
	View() IView
	ViewController() IViewController
	VisibilityPriority() TouchBarItemPriority
	SetVisibilityPriority(value TouchBarItemPriority)
	IsContinuous() bool
	SetIsContinuous(value bool)
	GroupTouchBar() TouchBar /* not a class type */
	SetGroupTouchBar(value TouchBar /* not a class type */)
	CollapsedRepresentation() IView
	SetCollapsedRepresentation(value IView)
	PopoverTouchBar() TouchBar /* not a class type */
	SetPopoverTouchBar(value TouchBar /* not a class type */)
	PressAndHoldTouchBar() TouchBar /* not a class type */
	SetPressAndHoldTouchBar(value TouchBar /* not a class type */)
	TrackingMode() objectivec.IObject
	SetTrackingMode(value objectivec.IObject)
	PrincipalItemIdentifier() objectivec.IObject
	SetPrincipalItemIdentifier(value objectivec.IObject)
	IsVisible() bool
	SetIsVisible(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Initializes and returns a new item from a storyboard or nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewTouchBarItemWithCoder(coder foundation.foundation.INSCoder) TouchBarItem {
	instance := getTouchBarItemClass().Alloc()
	rv := objc.Send[TouchBarItem](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates a new item with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	instance := getTouchBarItemClass().Alloc()
	rv := objc.Send[TouchBarItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}






















// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/customizationLabel
func (t_ TouchBarItem) CustomizationLabel() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The identifier for this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/identifier-swift.property
func (t_ TouchBarItem) Identifier() TouchBarItemIdentifier {
	rv := objc.Send[TouchBarItemIdentifier](t_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/isVisible
func (t_ TouchBarItem) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}


// The view associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/view
func (t_ TouchBarItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}


// The view controller associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/viewController
func (t_ TouchBarItem) ViewController() IViewController {
	rv := objc.Send[ViewController](t_.ID, objc.Sel("viewController"))
	return rv
}


// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) VisibilityPriority() TouchBarItemPriority {
	rv := objc.Send[TouchBarItemPriority](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) SetVisibilityPriority(value TouchBarItemPriority) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
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
func (t_ TouchBarItem) GroupTouchBar() TouchBar /* not a class type */ {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) SetGroupTouchBar(value TouchBar /* not a class type */) {
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
func (t_ TouchBarItem) PopoverTouchBar() TouchBar /* not a class type */ {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) SetPopoverTouchBar(value TouchBar /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPopoverTouchBar:"), value)
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) PressAndHoldTouchBar() TouchBar /* not a class type */ {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) SetPressAndHoldTouchBar(value TouchBar /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) TrackingMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("trackingMode"))
	return rv
}


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) SetTrackingMode(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackingMode:"), value)
}


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) PrincipalItemIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("principalItemIdentifier"))
	return rv
}


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) SetPrincipalItemIdentifier(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrincipalItemIdentifier:"), value)
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







