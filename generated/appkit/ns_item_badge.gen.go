// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ItemBadge] class.
var (
	ItemBadgeClass     _ItemBadgeClass
	ItemBadgeClassOnce sync.Once
)

func getItemBadgeClass() _ItemBadgeClass {
	ItemBadgeClassOnce.Do(func() {
		ItemBadgeClass = _ItemBadgeClass{objc.GetClass("NSItemBadge")}
	})
	return ItemBadgeClass
}

type _ItemBadgeClass struct {
	class objc.Class
}

// An interface definition for the [ItemBadge] class.
type IItemBadge interface {
	objectivec.IObject
}

// represents a badge that can be attached to an .
//
// This badge provides a way to display small visual indicators, such as counts and text labels, within a toolbar item. Badges can be used to highlight important information, such as unread notifications or status indicators.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class
type ItemBadge struct {
	objectivec.Object
}

// ItemBadgeFrom constructs a [ItemBadge] from an unsafe.Pointer.
//
// represents a badge that can be attached to an .
func ItemBadgeFrom(ptr unsafe.Pointer) ItemBadge {
	return ItemBadge{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ItemBadgeClass) Alloc() ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ItemBadgeClass) New() ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ItemBadge) Init() ItemBadge {
	rv := objc.Send[ItemBadge](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ItemBadge) Autorelease() ItemBadge {
	rv := objc.Send[ItemBadge](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewItemBadge creates a new ItemBadge instance.
func NewItemBadge() ItemBadge {
	return getItemBadgeClass().New()
}


// Creates a badge displaying a localized numerical count.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithCount:
func (ic _ItemBadgeClass) BadgeWithCount(count int) ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("badgeWithCount:"), count)
	return rv
}

// Creates a badge displaying a text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithText:
func (ic _ItemBadgeClass) BadgeWithText(text string) ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("badgeWithText:"), objc.String(text))
	return rv
}

// Creates a badge styled as an indicator. In this context, an indicator is simply a badge without any text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/indicatorBadge
func (ic _ItemBadgeClass) IndicatorBadge() ItemBadge {
	rv := objc.Send[ItemBadge](objc.ID(ic.class), objc.Sel("indicatorBadge"))
	return rv
}

// The text to be displayed within the badge.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/text
func (i_ ItemBadge) Text() string {
	rv := objc.Send[string](i_.ID, objc.Sel("text"))
	return rv
}

// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i_ ItemBadge) IsBordered() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isBordered"))
	return rv
}


// SetIsBordered sets the value of the isBordered property.
// A Boolean value that indicates whether the toolbar item has a bordered style.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i_ ItemBadge) SetIsBordered(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBordered:"), value)
}

// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i_ ItemBadge) IsEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates whether the item is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i_ ItemBadge) SetIsEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i_ ItemBadge) IsHidden() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i_ ItemBadge) SetIsHidden(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsHidden:"), value)
}

// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i_ ItemBadge) IsNavigational() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isNavigational"))
	return rv
}


// SetIsNavigational sets the value of the isNavigational property.
// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i_ ItemBadge) SetIsNavigational(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsNavigational:"), value)
}

// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i_ ItemBadge) IsVisible() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isVisible"))
	return rv
}


// SetIsVisible sets the value of the isVisible property.
// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i_ ItemBadge) SetIsVisible(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsVisible:"), value)
}

// Defines the toolbar item’s appearance. The default style is plain.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i_ ItemBadge) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Defines the toolbar item’s appearance. The default style is plain.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i_ ItemBadge) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStyle:"), value)
}

// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i_ ItemBadge) Tag() int {
	rv := objc.Send[int](i_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// An integer tag you can use to identify the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i_ ItemBadge) SetTag(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTag:"), value)
}

// The display priority associated with the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i_ ItemBadge) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// SetVisibilityPriority sets the value of the visibilityPriority property.
// The display priority associated with the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i_ ItemBadge) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVisibilityPriority:"), value)
}



