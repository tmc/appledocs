// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPopoverTouchBarItem */


/* debug [class_header]: Header for NSPopoverTouchBarItem */
// The class instance for the [PopoverTouchBarItem] class.
var (
	PopoverTouchBarItemClass     _PopoverTouchBarItemClass
	PopoverTouchBarItemClassOnce sync.Once
)

func getPopoverTouchBarItemClass() _PopoverTouchBarItemClass {
	PopoverTouchBarItemClassOnce.Do(func() {
		PopoverTouchBarItemClass = _PopoverTouchBarItemClass{objc.GetClass("NSPopoverTouchBarItem")}
	})
	return PopoverTouchBarItemClass
}

type _PopoverTouchBarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PopoverTouchBarItem */
// An interface definition for the [PopoverTouchBarItem] class.
type IPopoverTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for PopoverTouchBarItem */
	// properties:
	CollapsedRepresentation() IView
	SetCollapsedRepresentation(value IView)
	CollapsedRepresentationImage() IImage
	SetCollapsedRepresentationImage(value IImage)
	CollapsedRepresentationLabel() objc.IObject /* cross-framework: NSString */
	SetCollapsedRepresentationLabel(value objc.IObject /* cross-framework: NSString */)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	PopoverTouchBar() objc.IObject /* cross-framework: TouchBar */
	SetPopoverTouchBar(value objc.IObject /* cross-framework: TouchBar */)
	PressAndHoldTouchBar() objc.IObject /* cross-framework: TouchBar */
	SetPressAndHoldTouchBar(value objc.IObject /* cross-framework: TouchBar */)
	ShowsCloseButton() bool
	SetShowsCloseButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PopoverTouchBarItem */
	// methods:
	DismissPopover(sender objc.IObject)
	MakeStandardActivatePopoverGestureRecognizer() objc.IObject /* cross-framework: GestureRecognizer */
	ShowPopover(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PopoverTouchBarItem */
// Alloc allocates a new instance without initialization.
func (pc _PopoverTouchBarItemClass) Alloc() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PopoverTouchBarItemClass) New() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PopoverTouchBarItem) Init() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PopoverTouchBarItem) Autorelease() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopoverTouchBarItem creates a new PopoverTouchBarItem instance.
func NewPopoverTouchBarItem() PopoverTouchBarItem {
	return getPopoverTouchBarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PopoverTouchBarItem */
// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns.


// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem
type PopoverTouchBarItem struct {
	TouchBarItem
}

// PopoverTouchBarItemFrom constructs a [PopoverTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns.
func PopoverTouchBarItemFrom(ptr unsafe.Pointer) PopoverTouchBarItem {
	return PopoverTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PopoverTouchBarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PopoverTouchBarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PopoverTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PopoverTouchBarItem */

// Restores the previously visible main bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/dismissPopover(_:)
func (p_ PopoverTouchBarItem) DismissPopover(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("dismissPopover:"), sender)
}/* debug [instance_methods/method]: DismissPopover */


// Returns a gesture recognizer, configured to invoke the method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/makeStandardActivatePopoverGestureRecognizer()
func (p_ PopoverTouchBarItem) MakeStandardActivatePopoverGestureRecognizer() objc.IObject /* cross-framework: GestureRecognizer */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("makeStandardActivatePopoverGestureRecognizer"))
	return rv
}/* debug [instance_methods/method]: MakeStandardActivatePopoverGestureRecognizer */


// Replaces the main bar with this item’s popover bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/showPopover(_:)
func (p_ PopoverTouchBarItem) ShowPopover(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("showPopover:"), sender)
}/* debug [instance_methods/method]: ShowPopover */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PopoverTouchBarItem */

// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentation
func (p_ PopoverTouchBarItem) CollapsedRepresentation() IView {
	rv := objc.Send[View](p_.ID, objc.Sel("collapsedRepresentation"))
	return rv
}/* debug [instance_properties/getter]: collapsedRepresentation */


// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentation
func (p_ PopoverTouchBarItem) SetCollapsedRepresentation(value IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentation:"), value)
}/* debug [instance_properties/setter]: collapsedRepresentation */


// The image displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationImage
func (p_ PopoverTouchBarItem) CollapsedRepresentationImage() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("collapsedRepresentationImage"))
	return rv
}/* debug [instance_properties/getter]: collapsedRepresentationImage */


// The image displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationImage
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}/* debug [instance_properties/setter]: collapsedRepresentationImage */


// The localized string displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationLabel
func (p_ PopoverTouchBarItem) CollapsedRepresentationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("collapsedRepresentationLabel"))
	return rv
}/* debug [instance_properties/getter]: collapsedRepresentationLabel */


// The localized string displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationLabel
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationLabel:"), value)
}/* debug [instance_properties/setter]: collapsedRepresentationLabel */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/customizationLabel
func (p_ PopoverTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/customizationLabel
func (p_ PopoverTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p_ PopoverTouchBarItem) PopoverTouchBar() objc.IObject /* cross-framework: TouchBar */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("popoverTouchBar"))
	return rv
}/* debug [instance_properties/getter]: popoverTouchBar */


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p_ PopoverTouchBarItem) SetPopoverTouchBar(value objc.IObject /* cross-framework: TouchBar */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPopoverTouchBar:"), value)
}/* debug [instance_properties/setter]: popoverTouchBar */


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p_ PopoverTouchBarItem) PressAndHoldTouchBar() objc.IObject /* cross-framework: TouchBar */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}/* debug [instance_properties/getter]: pressAndHoldTouchBar */


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p_ PopoverTouchBarItem) SetPressAndHoldTouchBar(value objc.IObject /* cross-framework: TouchBar */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}/* debug [instance_properties/setter]: pressAndHoldTouchBar */


// A Boolean value that determines whether a close button should be shown on the popover bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/showsCloseButton
func (p_ PopoverTouchBarItem) ShowsCloseButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsCloseButton"))
	return rv
}/* debug [instance_properties/getter]: showsCloseButton */


// A Boolean value that determines whether a close button should be shown on the popover bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/showsCloseButton
func (p_ PopoverTouchBarItem) SetShowsCloseButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsCloseButton:"), value)
}/* debug [instance_properties/setter]: showsCloseButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPopoverTouchBarItem */



