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
}

// A control that provides additional quantitative information specific to a menu item, such as the number of available updates.
//
// You create a badge using an initializer or a predefined factory method, and then you assign it to the property of a for display. For example, to display a badge with a count, use the initalizer, passing in the value of as an . To display a badge with a custom string, use the initializer, passing in the string you want to display. To display a badge using a predefined , use a factory method such as , passing in the of the badge to display. The default value of this property is .
//
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


// The type of items the badge displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/type
func (m_ MenuItemBadge) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("type"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/badge
func (m_ MenuItemBadge) Badge() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("badge"))
	return rv
}


// SetBadge sets the value of the badge property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/badge
func (m_ MenuItemBadge) SetBadge(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadge:"), value)
}

// The number of items the badge displays.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitembadge/itemcount
func (m_ MenuItemBadge) ItemCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("itemCount"))
	return rv
}


// SetItemCount sets the value of the itemCount property.
// The number of items the badge displays.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitembadge/itemcount
func (m_ MenuItemBadge) SetItemCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setItemCount:"), value)
}

// The string representation of the badge as it would appear when
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitembadge/stringvalue-32sbt
func (m_ MenuItemBadge) StringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("stringValue"))
	return rv
}


// SetStringValue sets the value of the stringValue property.
// The string representation of the badge as it would appear when

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitembadge/stringvalue-32sbt
func (m_ MenuItemBadge) SetStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), objc.String(value))
}



