// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTouchBarItem */


/* debug [class_header]: Header for NSTouchBarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TouchBarItem */
// An interface definition for the [TouchBarItem] class.
type ITouchBarItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TouchBarItem */
	// properties:
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	Identifier() TouchBarItemIdentifier /* typedef */
	Visible() bool
	View() IView
	ViewController() IViewController
	VisibilityPriority() TouchBarItemPriority /* typedef */
	SetVisibilityPriority(value TouchBarItemPriority /* typedef */)
	IsContinuous() bool
	SetIsContinuous(value bool)
	GroupTouchBar() objc.IObject /* cross-framework: TouchBar */
	SetGroupTouchBar(value objc.IObject /* cross-framework: TouchBar */)
	CollapsedRepresentation() IView
	SetCollapsedRepresentation(value IView)
	PopoverTouchBar() objc.IObject /* cross-framework: TouchBar */
	SetPopoverTouchBar(value objc.IObject /* cross-framework: TouchBar */)
	PressAndHoldTouchBar() objc.IObject /* cross-framework: TouchBar */
	SetPressAndHoldTouchBar(value objc.IObject /* cross-framework: TouchBar */)
	TrackingMode() objectivec.IObject
	SetTrackingMode(value objectivec.IObject)
	PrincipalItemIdentifier() objectivec.IObject
	SetPrincipalItemIdentifier(value objectivec.IObject)
	IsVisible() bool
	SetIsVisible(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TouchBarItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TouchBarItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TouchBarItem */

// Initializes and returns a new item from a storyboard or nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewTouchBarItemWithCoder(coder foundation.Coder) TouchBarItem {
	instance := getTouchBarItemClass().Alloc()
	rv := objc.Send[TouchBarItem](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTouchBarItemWithCoder */


// Creates a new item with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewTouchBarItemWithIdentifier(identifier TouchBarItemIdentifier /* typedef */) TouchBarItem {
	instance := getTouchBarItemClass().Alloc()
	rv := objc.Send[TouchBarItem](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTouchBarItemWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TouchBarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TouchBarItem */

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/customizationLabel
func (t_ TouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The identifier for this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/identifier-swift.property
func (t_ TouchBarItem) Identifier() TouchBarItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/isVisible
func (t_ TouchBarItem) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}/* debug [instance_properties/getter]: visible */


// The view associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/view
func (t_ TouchBarItem) View() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// The view controller associated with this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/viewController
func (t_ TouchBarItem) ViewController() IViewController {
	rv := objc.Send[ViewController](t_.ID, objc.Sel("viewController"))
	return rv
}/* debug [instance_properties/getter]: viewController */


// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) VisibilityPriority() TouchBarItemPriority /* typedef */ {
	rv := objc.Send[float32](t_.ID, objc.Sel("visibilityPriority"))
	return rv
}/* debug [instance_properties/getter]: visibilityPriority */


// Determines which items are shown in a bar when space is limited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t_ TouchBarItem) SetVisibilityPriority(value TouchBarItemPriority /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVisibilityPriority:"), value)
}/* debug [instance_properties/setter]: visibilityPriority */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t_ TouchBarItem) IsContinuous() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuous"))
	return rv
}/* debug [instance_properties/getter]: isContinuous */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t_ TouchBarItem) SetIsContinuous(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuous:"), value)
}/* debug [instance_properties/setter]: isContinuous */


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) GroupTouchBar() objc.IObject /* cross-framework: TouchBar */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("groupTouchBar"))
	return rv
}/* debug [instance_properties/getter]: groupTouchBar */


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBarItem) SetGroupTouchBar(value objc.IObject /* cross-framework: TouchBar */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupTouchBar:"), value)
}/* debug [instance_properties/setter]: groupTouchBar */


// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t_ TouchBarItem) CollapsedRepresentation() IView {
	rv := objc.Send[View](t_.ID, objc.Sel("collapsedRepresentation"))
	return rv
}/* debug [instance_properties/getter]: collapsedRepresentation */


// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t_ TouchBarItem) SetCollapsedRepresentation(value IView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCollapsedRepresentation:"), value)
}/* debug [instance_properties/setter]: collapsedRepresentation */


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) PopoverTouchBar() objc.IObject /* cross-framework: TouchBar */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("popoverTouchBar"))
	return rv
}/* debug [instance_properties/getter]: popoverTouchBar */


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBarItem) SetPopoverTouchBar(value objc.IObject /* cross-framework: TouchBar */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPopoverTouchBar:"), value)
}/* debug [instance_properties/setter]: popoverTouchBar */


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) PressAndHoldTouchBar() objc.IObject /* cross-framework: TouchBar */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}/* debug [instance_properties/getter]: pressAndHoldTouchBar */


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBarItem) SetPressAndHoldTouchBar(value objc.IObject /* cross-framework: TouchBar */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}/* debug [instance_properties/setter]: pressAndHoldTouchBar */


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) TrackingMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("trackingMode"))
	return rv
}/* debug [instance_properties/getter]: trackingMode */


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t_ TouchBarItem) SetTrackingMode(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackingMode:"), value)
}/* debug [instance_properties/setter]: trackingMode */


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) PrincipalItemIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("principalItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: principalItemIdentifier */


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBarItem) SetPrincipalItemIdentifier(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrincipalItemIdentifier:"), value)
}/* debug [instance_properties/setter]: principalItemIdentifier */


// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/isvisible
func (t_ TouchBarItem) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that reflects whether or not the item is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/isvisible
func (t_ TouchBarItem) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTouchBarItem */


