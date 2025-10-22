// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItemBadge] class.
var (
	MenuItemBadgeClass     _MenuItemBadgeClass
	MenuItemBadgeClassOnce sync.Once
)

func getMenuItemBadgeClass() _MenuItemBadgeClass {
	MenuItemBadgeClassOnce.Do(func() {
		MenuItemBadgeClass = _MenuItemBadgeClass{objc.GetClass("NSMenuItemBadge")}
	})
	return MenuItemBadgeClass
}

type _MenuItemBadgeClass struct {
	class objc.Class
}

// An interface definition for the [MenuItemBadge] class.
type IMenuItemBadge interface {
	objectivec.IObject
	ItemCount() int
	StringValue() string
	Type() MenuItemBadgeType
	Badge() NSMenuItemBadge
	SetBadge(value IMenuItemBadge)
}

// A control that provides additional quantitative information specific to a menu item, such as the number of available updates.
//
// You create a badge using an initializer or a predefined factory method, and then you assign it to the property of a for display. For example, to display a badge with a count, use the initalizer, passing in the value of as an . To display a badge with a custom string, use the initializer, passing in the string you want to display. To display a badge using a predefined , use a factory method such as , passing in the of the badge to display. The default value of this property is .


// A control that provides additional quantitative information specific to a menu item, such as the number of available updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge

type MenuItemBadge struct {
	objectivec.Object
}

// MenuItemBadgeFrom constructs a [MenuItemBadge] from an unsafe.Pointer.
//
// A control that provides additional quantitative information specific to a menu item, such as the number of available updates.
func MenuItemBadgeFrom(ptr unsafe.Pointer) MenuItemBadge {
	return MenuItemBadge{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemBadgeClass) Alloc() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemBadgeClass) New() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItemBadge) Init() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItemBadge) Autorelease() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItemBadge creates a new MenuItemBadge instance.
func NewMenuItemBadge() MenuItemBadge {
	return getMenuItemBadgeClass().New()
}




// Creates a badge with a count and an empty string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/init(count:)

func NewMenuItemBadgeWithCount(itemCount int) MenuItemBadge {
	instance := getMenuItemBadgeClass().Alloc()
	rv := objc.Send[MenuItemBadge](instance.ID, objc.Sel("initWithCount:"), itemCount)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/initWithCount:type:

func NewMenuItemBadgeWithCountType(itemCount int, type_ MenuItemBadgeType) MenuItemBadge {
	instance := getMenuItemBadgeClass().Alloc()
	rv := objc.Send[MenuItemBadge](instance.ID, objc.Sel("initWithCount:type:"), itemCount, type_)
	rv.Autorelease()
	return rv
}



// Creates a badge with the provided custom string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/init(string:)

func NewMenuItemBadgeWithString(string_ string) MenuItemBadge {
	instance := getMenuItemBadgeClass().Alloc()
	rv := objc.Send[MenuItemBadge](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}



// Creates an alert-style badge with an integer count and a predefined label that represents the number of alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/alerts(count:)

func (mc _MenuItemBadgeClass) AlertsWithCount(itemCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("alertsWithCount:"), itemCount)
	return rv
}


// Creates a new item-style badge with an integer count and a predefined label that represents the number of new items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/newItems(count:)

func (mc _MenuItemBadgeClass) NewItemsWithCount(itemCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("newItemsWithCount:"), itemCount)
	return rv
}


// Creates an update-style badge with an integer count and a predefined label that represents the number of available updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/updates(count:)

func (mc _MenuItemBadgeClass) UpdatesWithCount(itemCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("updatesWithCount:"), itemCount)
	return rv
}


// The number of items the badge displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/itemCount

func (m_ MenuItemBadge) ItemCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("itemCount"))
	return rv
}


// The string representation of the badge when it displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/stringValue-fc9f

func (m_ MenuItemBadge) StringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("stringValue"))
	return rv
}


// The type of items the badge displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/type

func (m_ MenuItemBadge) Type() MenuItemBadgeType {
	rv := objc.Send[MenuItemBadgeType](m_.ID, objc.Sel("type"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/badge

func (m_ MenuItemBadge) Badge() NSMenuItemBadge {
	rv := objc.Send[NSMenuItemBadge](m_.ID, objc.Sel("badge"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/badge

func (m_ MenuItemBadge) SetBadge(value IMenuItemBadge) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadge:"), value)
}


