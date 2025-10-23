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
	// properties:
	IsBordered() bool /* primitive/slice/pointer. */
	SetIsBordered(value bool /* primitive/slice/pointer. */)
	IsEnabled() bool /* primitive/slice/pointer. */
	SetIsEnabled(value bool /* primitive/slice/pointer. */)
	IsHidden() bool /* primitive/slice/pointer. */
	SetIsHidden(value bool /* primitive/slice/pointer. */)
	IsNavigational() bool /* primitive/slice/pointer. */
	SetIsNavigational(value bool /* primitive/slice/pointer. */)
	IsVisible() bool /* primitive/slice/pointer. */
	SetIsVisible(value bool /* primitive/slice/pointer. */)
	Style() unsafe.Pointer
	SetStyle(value unsafe.Pointer)
	Tag() int /* primitive/slice/pointer. */
	SetTag(value int /* primitive/slice/pointer. */)
	VisibilityPriority() unsafe.Pointer
	SetVisibilityPriority(value unsafe.Pointer)
	// methods:
}

// represents a badge that can be attached to an .
//
// This badge provides a way to display small visual indicators, such as counts and text labels, within a toolbar item. Badges can be used to highlight important information, such as unread notifications or status indicators.


// represents a badge that can be attached to an .
//
// [Full Topic]
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



// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i_ ItemBadge) IsBordered() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isBordered"))
	return rv
}


// A Boolean value that indicates whether the toolbar item has a bordered style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isbordered
func (i_ ItemBadge) SetIsBordered(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsBordered:"), value)
}


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i_ ItemBadge) IsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the item is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isenabled
func (i_ ItemBadge) SetIsEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i_ ItemBadge) IsHidden() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/ishidden
func (i_ ItemBadge) SetIsHidden(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsHidden:"), value)
}


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i_ ItemBadge) IsNavigational() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isNavigational"))
	return rv
}


// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isnavigational
func (i_ ItemBadge) SetIsNavigational(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsNavigational:"), value)
}


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i_ ItemBadge) IsVisible() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/isvisible
func (i_ ItemBadge) SetIsVisible(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsVisible:"), value)
}


// Defines the toolbar item’s appearance. The default style is plain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i_ ItemBadge) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("style"))
	return rv
}


// Defines the toolbar item’s appearance. The default style is plain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/style-swift.property
func (i_ ItemBadge) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStyle:"), value)
}


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i_ ItemBadge) Tag() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](i_.ID, objc.Sel("tag"))
	return rv
}


// An integer tag you can use to identify the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/tag
func (i_ ItemBadge) SetTag(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTag:"), value)
}


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i_ ItemBadge) VisibilityPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("visibilityPriority"))
	return rv
}


// The display priority associated with the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstoolbaritem/visibilitypriority-swift.property
func (i_ ItemBadge) SetVisibilityPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVisibilityPriority:"), value)
}



