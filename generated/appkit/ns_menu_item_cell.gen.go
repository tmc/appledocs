// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	ImageWidth() float64
	KeyEquivalentWidth() float64
	MenuItem() IMenuItem
	SetMenuItem(value IMenuItem)
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	NeedsSizing() bool
	SetNeedsSizing(value bool)
	StateImageWidth() float64
	Tag() int
	SetTag(value int)
	TitleWidth() float64


	

	// methods:
	CalcSize()
	DrawBorderAndBackgroundWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DrawImageWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DrawKeyEquivalentWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DrawSeparatorItemWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DrawStateImageWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	DrawTitleWithFrameInView(cellFrame corefoundation.CGRect, controlView IView)
	KeyEquivalentRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect
	StateImageRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect
	TitleRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect


}





// Alloc allocates a new instance without initialization.
func (mc _MenuItemCellClass) Alloc() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that handles the measurement and display of a single menu item in its encompassing frame.


// An object that handles the measurement and display of a single menu item in its encompassing frame.
//
// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(textCell:)
func NewMenuItemCellTextCell(string_ foundation.foundation.INSString) MenuItemCell {
	instance := getMenuItemCellClass().Alloc()
	rv := objc.Send[MenuItemCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(coder:)
func NewMenuItemCellWithCoder(coder foundation.foundation.INSCoder) MenuItemCell {
	instance := getMenuItemCellClass().Alloc()
	rv := objc.Send[MenuItemCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

















// Calculates the minimum required width and height of the receiver’s menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/calcSize()
func (m_ MenuItemCell) CalcSize() {
	objc.Send[objc.ID](m_.ID, objc.Sel("calcSize"))
}


// Draws the borders and background associated with the receiver’s menu item (if any).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawBorderAndBackground(withFrame:in:)
func (m_ MenuItemCell) DrawBorderAndBackgroundWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawBorderAndBackgroundWithFrame:inView:"), cellFrame, controlView)
}


// Draws the image associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawImage(withFrame:in:)
func (m_ MenuItemCell) DrawImageWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawImageWithFrame:inView:"), cellFrame, controlView)
}


// Draws the key equivalent associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawKeyEquivalent(withFrame:in:)
func (m_ MenuItemCell) DrawKeyEquivalentWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawKeyEquivalentWithFrame:inView:"), cellFrame, controlView)
}


// Draws a menu item separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawSeparatorItem(withFrame:in:)
func (m_ MenuItemCell) DrawSeparatorItemWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawSeparatorItemWithFrame:inView:"), cellFrame, controlView)
}


// Draws the state image associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawStateImage(withFrame:in:)
func (m_ MenuItemCell) DrawStateImageWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawStateImageWithFrame:inView:"), cellFrame, controlView)
}


// Draws the title associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawTitle(withFrame:in:)
func (m_ MenuItemCell) DrawTitleWithFrameInView(cellFrame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawTitleWithFrame:inView:"), cellFrame, controlView)
}


// Returns the rectangle into which the menu item’s key equivalent should be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/keyEquivalentRect(forBounds:)
func (m_ MenuItemCell) KeyEquivalentRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("keyEquivalentRectForBounds:"), cellFrame)
	return rv
}


// Returns the rectangle into which the menu item’s state image should be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/stateImageRect(forBounds:)
func (m_ MenuItemCell) StateImageRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("stateImageRectForBounds:"), cellFrame)
	return rv
}


// Returns the rectangle into which the menu item’s title should be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/titleRect(forBounds:)
func (m_ MenuItemCell) TitleRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("titleRectForBounds:"), cellFrame)
	return rv
}







// The width of the image associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/imageWidth
func (m_ MenuItemCell) ImageWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("imageWidth"))
	return rv
}


// The width of the menu item’s key equivalent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/keyEquivalentWidth
func (m_ MenuItemCell) KeyEquivalentWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("keyEquivalentWidth"))
	return rv
}


// The menu item object associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/menuItem
func (m_ MenuItemCell) MenuItem() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("menuItem"))
	return rv
}


// The menu item object associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/menuItem
func (m_ MenuItemCell) SetMenuItem(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuItem:"), value)
}


// A Boolean value indicating whether the menu item needs to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsDisplay
func (m_ MenuItemCell) NeedsDisplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("needsDisplay"))
	return rv
}


// A Boolean value indicating whether the menu item needs to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsDisplay
func (m_ MenuItemCell) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsDisplay:"), value)
}


// A Boolean value indicating whether the size of the menu needs to be calculated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsSizing
func (m_ MenuItemCell) NeedsSizing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("needsSizing"))
	return rv
}


// A Boolean value indicating whether the size of the menu needs to be calculated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsSizing
func (m_ MenuItemCell) SetNeedsSizing(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsSizing:"), value)
}


// The width of the image used to indicate the state of the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/stateImageWidth
func (m_ MenuItemCell) StateImageWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("stateImageWidth"))
	return rv
}


// The integer tag of the selected menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/tag
func (m_ MenuItemCell) Tag() int {
	rv := objc.Send[int](m_.ID, objc.Sel("tag"))
	return rv
}


// The integer tag of the selected menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/tag
func (m_ MenuItemCell) SetTag(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}


// The width of the menu item’s text, measured in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/titleWidth
func (m_ MenuItemCell) TitleWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("titleWidth"))
	return rv
}







