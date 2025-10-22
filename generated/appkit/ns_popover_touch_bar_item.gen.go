// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PopoverTouchBarItem] class.
type IPopoverTouchBarItem interface {
	ITouchBarItem
	DismissPopover(sender objectivec.IObject)
	MakeStandardActivatePopoverGestureRecognizer() GestureRecognizer
	ShowPopover(sender objectivec.IObject)
	CollapsedRepresentation() NSView
	SetCollapsedRepresentation(value IView)
	CollapsedRepresentationImage() Image
	SetCollapsedRepresentationImage(value IImage)
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	PopoverTouchBar() NSTouchBar
	SetPopoverTouchBar(value ITouchBar)
	PressAndHoldTouchBar() NSTouchBar
	SetPressAndHoldTouchBar(value ITouchBar)
	ShowsCloseButton() bool
	SetShowsCloseButton(value bool)
}

// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PopoverTouchBarItemClass) Alloc() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Restores the previously visible main bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/dismissPopover(_:)
func (p_ PopoverTouchBarItem) DismissPopover(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("dismissPopover:"), sender)
}

// Returns a gesture recognizer, configured to invoke the method.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/makeStandardActivatePopoverGestureRecognizer()
func (p_ PopoverTouchBarItem) MakeStandardActivatePopoverGestureRecognizer() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](p_.ID, objc.Sel("makeStandardActivatePopoverGestureRecognizer"))
	return rv
}

// Replaces the main bar with this item’s popover bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/showPopover(_:)
func (p_ PopoverTouchBarItem) ShowPopover(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("showPopover:"), sender)
}

// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentation
func (p_ PopoverTouchBarItem) CollapsedRepresentation() NSView {
	rv := objc.Send[NSView](p_.ID, objc.Sel("collapsedRepresentation"))
	return rv
}


// SetCollapsedRepresentation sets the value of the collapsedRepresentation property.
// The view displayed when this item is displayed in its parent bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentation
func (p_ PopoverTouchBarItem) SetCollapsedRepresentation(value IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentation:"), value)
}

// The image displayed by the button for the default collapsed representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationImage
func (p_ PopoverTouchBarItem) CollapsedRepresentationImage() Image {
	rv := objc.Send[Image](p_.ID, objc.Sel("collapsedRepresentationImage"))
	return rv
}


// SetCollapsedRepresentationImage sets the value of the collapsedRepresentationImage property.
// The image displayed by the button for the default collapsed representation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationImage
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}

// The localized string displayed by the button for the default collapsed representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationLabel
func (p_ PopoverTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Send[string](p_.ID, objc.Sel("collapsedRepresentationLabel"))
	return rv
}


// SetCollapsedRepresentationLabel sets the value of the collapsedRepresentationLabel property.
// The localized string displayed by the button for the default collapsed representation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationLabel
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationLabel:"), objc.String(value))
}

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/customizationLabel
func (p_ PopoverTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](p_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/customizationLabel
func (p_ PopoverTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}

// The bar displayed when this item is “popped.”
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p_ PopoverTouchBarItem) PopoverTouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](p_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// SetPopoverTouchBar sets the value of the popoverTouchBar property.
// The bar displayed when this item is “popped.”

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p_ PopoverTouchBarItem) SetPopoverTouchBar(value ITouchBar) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPopoverTouchBar:"), value)
}

// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p_ PopoverTouchBarItem) PressAndHoldTouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](p_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// SetPressAndHoldTouchBar sets the value of the pressAndHoldTouchBar property.
// The bar that is displayed when a user press-and-holds on the popover item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p_ PopoverTouchBarItem) SetPressAndHoldTouchBar(value ITouchBar) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}

// A Boolean value that determines whether a close button should be shown on the popover bar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/showsclosebutton
func (p_ PopoverTouchBarItem) ShowsCloseButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsCloseButton"))
	return rv
}


// SetShowsCloseButton sets the value of the showsCloseButton property.
// A Boolean value that determines whether a close button should be shown on the popover bar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/showsclosebutton
func (p_ PopoverTouchBarItem) SetShowsCloseButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsCloseButton:"), value)
}



