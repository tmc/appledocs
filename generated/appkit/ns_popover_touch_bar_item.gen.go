// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	CollapsedRepresentation() IView
	SetCollapsedRepresentation(value IView)
	CollapsedRepresentationImage() IImage
	SetCollapsedRepresentationImage(value IImage)
	CollapsedRepresentationLabel() objc.IObject /* cross-framework: NSString */
	SetCollapsedRepresentationLabel(value objc.IObject /* cross-framework: NSString */)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	PopoverTouchBar() ITouchBar
	SetPopoverTouchBar(value ITouchBar)
	PressAndHoldTouchBar() ITouchBar
	SetPressAndHoldTouchBar(value ITouchBar)
	ShowsCloseButton() bool
	SetShowsCloseButton(value bool)
	// methods:
}

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



// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (p_ PopoverTouchBarItem) CollapsedRepresentation() IView {
	rv := objc.Send[View](p_.ID, objc.Sel("collapsedRepresentation"))
	return rv
}


// The view displayed when this item is displayed in its parent bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (p_ PopoverTouchBarItem) SetCollapsedRepresentation(value IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentation:"), value)
}


// The image displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentationimage
func (p_ PopoverTouchBarItem) CollapsedRepresentationImage() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("collapsedRepresentationImage"))
	return rv
}


// The image displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentationimage
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}


// The localized string displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentationlabel
func (p_ PopoverTouchBarItem) CollapsedRepresentationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("collapsedRepresentationLabel"))
	return rv
}


// The localized string displayed by the button for the default collapsed representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentationlabel
func (p_ PopoverTouchBarItem) SetCollapsedRepresentationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationLabel:"), value)
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/customizationlabel
func (p_ PopoverTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/customizationlabel
func (p_ PopoverTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (p_ PopoverTouchBarItem) PopoverTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](p_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (p_ PopoverTouchBarItem) SetPopoverTouchBar(value ITouchBar) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPopoverTouchBar:"), value)
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (p_ PopoverTouchBarItem) PressAndHoldTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](p_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (p_ PopoverTouchBarItem) SetPressAndHoldTouchBar(value ITouchBar) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}


// A Boolean value that determines whether a close button should be shown on the popover bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/showsclosebutton
func (p_ PopoverTouchBarItem) ShowsCloseButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsCloseButton"))
	return rv
}


// A Boolean value that determines whether a close button should be shown on the popover bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/showsclosebutton
func (p_ PopoverTouchBarItem) SetShowsCloseButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsCloseButton:"), value)
}



