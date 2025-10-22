// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MenuItemCell] class.
var (
	MenuItemCellClass     _MenuItemCellClass
	MenuItemCellClassOnce sync.Once
)

func getMenuItemCellClass() _MenuItemCellClass {
	MenuItemCellClassOnce.Do(func() {
		MenuItemCellClass = _MenuItemCellClass{objc.GetClass("NSMenuItemCell")}
	})
	return MenuItemCellClass
}

type _MenuItemCellClass struct {
	class objc.Class
}

// An interface definition for the [MenuItemCell] class.
type IMenuItemCell interface {
	IButtonCell
	ImageWidth() float64
	SetImageWidth(value float64)
	KeyEquivalentWidth() float64
	SetKeyEquivalentWidth(value float64)
	MenuItem() NSMenuItem
	SetMenuItem(value IMenuItem)
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	NeedsSizing() bool
	SetNeedsSizing(value bool)
	StateImageWidth() float64
	SetStateImageWidth(value float64)
	Tag() int
	SetTag(value int)
	TitleWidth() float64
	SetTitleWidth(value float64)
}

// An object that handles the measurement and display of a single menu item in its encompassing frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell
type MenuItemCell struct {
	ButtonCell
}

// MenuItemCellFrom constructs a [MenuItemCell] from an unsafe.Pointer.
//
// An object that handles the measurement and display of a single menu item in its encompassing frame.
func MenuItemCellFrom(ptr unsafe.Pointer) MenuItemCell {
	return MenuItemCell{
		ButtonCell: ButtonCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemCellClass) Alloc() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemCellClass) New() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItemCell) Init() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItemCell) Autorelease() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItemCell creates a new MenuItemCell instance.
func NewMenuItemCell() MenuItemCell {
	return getMenuItemCellClass().New()
}


// The width of the image associated with the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/imagewidth
func (m_ MenuItemCell) ImageWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("imageWidth"))
	return rv
}


// SetImageWidth sets the value of the imageWidth property.
// The width of the image associated with the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/imagewidth
func (m_ MenuItemCell) SetImageWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageWidth:"), value)
}

// The width of the menu item’s key equivalent string.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/keyequivalentwidth
func (m_ MenuItemCell) KeyEquivalentWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("keyEquivalentWidth"))
	return rv
}


// SetKeyEquivalentWidth sets the value of the keyEquivalentWidth property.
// The width of the menu item’s key equivalent string.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/keyequivalentwidth
func (m_ MenuItemCell) SetKeyEquivalentWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyEquivalentWidth:"), value)
}

// The menu item object associated with the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/menuitem
func (m_ MenuItemCell) MenuItem() NSMenuItem {
	rv := objc.Send[NSMenuItem](m_.ID, objc.Sel("menuItem"))
	return rv
}


// SetMenuItem sets the value of the menuItem property.
// The menu item object associated with the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/menuitem
func (m_ MenuItemCell) SetMenuItem(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuItem:"), value)
}

// A Boolean value indicating whether the menu item needs to be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/needsdisplay
func (m_ MenuItemCell) NeedsDisplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("needsDisplay"))
	return rv
}


// SetNeedsDisplay sets the value of the needsDisplay property.
// A Boolean value indicating whether the menu item needs to be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/needsdisplay
func (m_ MenuItemCell) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsDisplay:"), value)
}

// A Boolean value indicating whether the size of the menu needs to be calculated.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/needssizing
func (m_ MenuItemCell) NeedsSizing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("needsSizing"))
	return rv
}


// SetNeedsSizing sets the value of the needsSizing property.
// A Boolean value indicating whether the size of the menu needs to be calculated.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/needssizing
func (m_ MenuItemCell) SetNeedsSizing(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsSizing:"), value)
}

// The width of the image used to indicate the state of the menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/stateimagewidth
func (m_ MenuItemCell) StateImageWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("stateImageWidth"))
	return rv
}


// SetStateImageWidth sets the value of the stateImageWidth property.
// The width of the image used to indicate the state of the menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/stateimagewidth
func (m_ MenuItemCell) SetStateImageWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStateImageWidth:"), value)
}

// The integer tag of the selected menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/tag
func (m_ MenuItemCell) Tag() int {
	rv := objc.Send[int](m_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// The integer tag of the selected menu item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/tag
func (m_ MenuItemCell) SetTag(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}

// The width of the menu item’s text, measured in points.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/titlewidth
func (m_ MenuItemCell) TitleWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("titleWidth"))
	return rv
}


// SetTitleWidth sets the value of the titleWidth property.
// The width of the menu item’s text, measured in points.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitemcell/titlewidth
func (m_ MenuItemCell) SetTitleWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitleWidth:"), value)
}



