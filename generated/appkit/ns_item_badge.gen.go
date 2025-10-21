// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (ic _ItemBadgeClass) BadgeWithCount(count int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("badgeWithCount:"), count)
	return rv
}

// Creates a badge displaying a text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/badgeWithText:
func (ic _ItemBadgeClass) BadgeWithText(text string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("badgeWithText:"), objc.String(text))
	return rv
}

// Creates a badge styled as an indicator. In this context, an indicator is simply a badge without any text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/indicatorBadge
func (ic _ItemBadgeClass) IndicatorBadge() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("indicatorBadge"))
	return rv
}

// The text to be displayed within the badge.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSItemBadge-c.class/text
func (i_ ItemBadge) Text() string {
	rv := objc.Send[string](i_.ID, objc.Sel("text"))
	return rv
}



