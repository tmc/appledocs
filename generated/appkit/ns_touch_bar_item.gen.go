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
}

// A UI control shown in the Touch Bar on supported models of MacBook Pro.
//
// An instance of the class is called an . It appears to the user on the Touch Bar, typically along with other items, within the (invisible) bounds of the view for an object, called a . You use an item by adding it or its identifier to one or another of a bar’s arrays, depending on your app’s architecture and on the user customization you want to support. Because of the close interaction between bars and items, be sure you have read the overview for the class before continuing here to learn about items. AppKit provides a rich set of subclasses of , each of which is described in the corresponding class reference document: An object (a ), along with its delegate, provides a list of textual suggestions for the current text view An object (a ) provides a system-defined color picker An object (a ) contains a responder of your choice, such as a view, a button, or a scrubber (an instance of the class) An object (a ) provides a bar to contain other items An object (a ) provides a two-state control that, when touched or pressed, expands into its second state, showing the contents of a bar it owns An object (a ), along with its delegate, provides a list of objects eligible for sharing An object (a ) provides a slider control for choosing a value in a range The two most commonly-used item classes are and . Refer to the following sample code projects which demonstrate how to use and related classes:
//
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


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/customizationLabel
func (t_ TouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](t_.ID, objc.Sel("customizationLabel"))
	return rv
}

// The view associated with this item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/view
func (t_ TouchBarItem) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("view"))
	return rv
}

// The view controller associated with this item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/viewController
func (t_ TouchBarItem) ViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("viewController"))
	return rv
}

// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// SetVisibilityPriority sets the value of the visibilityPriority property.
// Determines which items are shown in a bar when space is limited.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}

// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t_ TouchBarItem) IsContinuous() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuous"))
	return rv
}


// SetIsContinuous sets the value of the isContinuous property.
// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t_ TouchBarItem) SetIsContinuous(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuous:"), value)
}

// A bar that holds this group’s items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) GroupTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// SetGroupTouchBar sets the value of the groupTouchBar property.
// A bar that holds this group’s items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) SetGroupTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupTouchBar:"), value)
}

// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t_ TouchBarItem) CollapsedRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("collapsedRepresentation"))
	return rv
}


// SetCollapsedRepresentation sets the value of the collapsedRepresentation property.
// The view displayed when this item is displayed in its parent bar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t_ TouchBarItem) SetCollapsedRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCollapsedRepresentation:"), value)
}

// The bar displayed when this item is “popped.”
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) PopoverTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// SetPopoverTouchBar sets the value of the popoverTouchBar property.
// The bar displayed when this item is “popped.”

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) SetPopoverTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPopoverTouchBar:"), value)
}

// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) PressAndHoldTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// SetPressAndHoldTouchBar sets the value of the pressAndHoldTouchBar property.
// The bar that is displayed when a user press-and-holds on the popover item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) SetPressAndHoldTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}

// The type of tracking behavior the control exhibits.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) TrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("trackingMode"))
	return rv
}


// SetTrackingMode sets the value of the trackingMode property.
// The type of tracking behavior the control exhibits.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) SetTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackingMode:"), value)
}

// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) PrincipalItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("principalItemIdentifier"))
	return rv
}


// SetPrincipalItemIdentifier sets the value of the principalItemIdentifier property.
// The identifier of an item you want the system to center in the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) SetPrincipalItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrincipalItemIdentifier:"), value)
}

// The identifier for this item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/identifier-swift.property
func (t_ TouchBarItem) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The identifier for this item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/identifier-swift.property
func (t_ TouchBarItem) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}

// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/isvisible
func (t_ TouchBarItem) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}


// SetIsVisible sets the value of the isVisible property.
// A Boolean value that reflects whether or not the item is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/isvisible
func (t_ TouchBarItem) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}



