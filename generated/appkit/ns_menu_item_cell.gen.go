// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMenuItemCell */


/* debug [class_header]: Header for NSMenuItemCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MenuItemCell */
// An interface definition for the [MenuItemCell] class.
type IMenuItemCell interface {
	IButtonCell
	
/* debug [class_interface_properties]: Properties for MenuItemCell */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MenuItemCell */
	// methods:
	CalcSize()
	DrawBorderAndBackgroundWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawImageWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawKeyEquivalentWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawSeparatorItemWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawStateImageWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawTitleWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	KeyEquivalentRectForBounds(cellFrame Rect /* not a class type */) Rect /* not a class type */
	StateImageRectForBounds(cellFrame Rect /* not a class type */) Rect /* not a class type */
	TitleRectForBounds(cellFrame Rect /* not a class type */) Rect /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MenuItemCell */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MenuItemCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MenuItemCell */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(textCell:)
func NewMenuItemCellTextCell(string_ objc.IObject /* cross-framework: NSString */) MenuItemCell {
	instance := getMenuItemCellClass().Alloc()
	rv := objc.Send[MenuItemCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMenuItemCellTextCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(coder:)
func NewMenuItemCellWithCoder(coder foundation.Coder) MenuItemCell {
	instance := getMenuItemCellClass().Alloc()
	rv := objc.Send[MenuItemCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMenuItemCellWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MenuItemCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MenuItemCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MenuItemCell */

// Calculates the minimum required width and height of the receiver’s menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/calcSize()
func (m_ MenuItemCell) CalcSize() {
	objc.Send[objc.ID](m_.ID, objc.Sel("calcSize"))
}/* debug [instance_methods/method]: CalcSize */


// Draws the borders and background associated with the receiver’s menu item (if any).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawBorderAndBackground(withFrame:in:)
func (m_ MenuItemCell) DrawBorderAndBackgroundWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawBorderAndBackgroundWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawBorderAndBackgroundWithFrameInView */


// Draws the image associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawImage(withFrame:in:)
func (m_ MenuItemCell) DrawImageWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawImageWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawImageWithFrameInView */


// Draws the key equivalent associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawKeyEquivalent(withFrame:in:)
func (m_ MenuItemCell) DrawKeyEquivalentWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawKeyEquivalentWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawKeyEquivalentWithFrameInView */


// Draws a menu item separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawSeparatorItem(withFrame:in:)
func (m_ MenuItemCell) DrawSeparatorItemWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawSeparatorItemWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawSeparatorItemWithFrameInView */


// Draws the state image associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawStateImage(withFrame:in:)
func (m_ MenuItemCell) DrawStateImageWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawStateImageWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawStateImageWithFrameInView */


// Draws the title associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawTitle(withFrame:in:)
func (m_ MenuItemCell) DrawTitleWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawTitleWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawTitleWithFrameInView */


// Returns the rectangle into which the menu item’s key equivalent should be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/keyEquivalentRect(forBounds:)
func (m_ MenuItemCell) KeyEquivalentRectForBounds(cellFrame Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](m_.ID, objc.Sel("keyEquivalentRectForBounds:"), cellFrame)
	return rv
}/* debug [instance_methods/method]: KeyEquivalentRectForBounds */


// Returns the rectangle into which the menu item’s state image should be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/stateImageRect(forBounds:)
func (m_ MenuItemCell) StateImageRectForBounds(cellFrame Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](m_.ID, objc.Sel("stateImageRectForBounds:"), cellFrame)
	return rv
}/* debug [instance_methods/method]: StateImageRectForBounds */


// Returns the rectangle into which the menu item’s title should be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/titleRect(forBounds:)
func (m_ MenuItemCell) TitleRectForBounds(cellFrame Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](m_.ID, objc.Sel("titleRectForBounds:"), cellFrame)
	return rv
}/* debug [instance_methods/method]: TitleRectForBounds */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MenuItemCell */

// The width of the image associated with the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/imageWidth
func (m_ MenuItemCell) ImageWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("imageWidth"))
	return rv
}/* debug [instance_properties/getter]: imageWidth */


// The width of the menu item’s key equivalent string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/keyEquivalentWidth
func (m_ MenuItemCell) KeyEquivalentWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("keyEquivalentWidth"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalentWidth */


// The menu item object associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/menuItem
func (m_ MenuItemCell) MenuItem() IMenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("menuItem"))
	return rv
}/* debug [instance_properties/getter]: menuItem */


// The menu item object associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/menuItem
func (m_ MenuItemCell) SetMenuItem(value IMenuItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenuItem:"), value)
}/* debug [instance_properties/setter]: menuItem */


// A Boolean value indicating whether the menu item needs to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsDisplay
func (m_ MenuItemCell) NeedsDisplay() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("needsDisplay"))
	return rv
}/* debug [instance_properties/getter]: needsDisplay */


// A Boolean value indicating whether the menu item needs to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsDisplay
func (m_ MenuItemCell) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsDisplay:"), value)
}/* debug [instance_properties/setter]: needsDisplay */


// A Boolean value indicating whether the size of the menu needs to be calculated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsSizing
func (m_ MenuItemCell) NeedsSizing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("needsSizing"))
	return rv
}/* debug [instance_properties/getter]: needsSizing */


// A Boolean value indicating whether the size of the menu needs to be calculated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsSizing
func (m_ MenuItemCell) SetNeedsSizing(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsSizing:"), value)
}/* debug [instance_properties/setter]: needsSizing */


// The width of the image used to indicate the state of the menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/stateImageWidth
func (m_ MenuItemCell) StateImageWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("stateImageWidth"))
	return rv
}/* debug [instance_properties/getter]: stateImageWidth */


// The integer tag of the selected menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/tag
func (m_ MenuItemCell) Tag() int {
	rv := objc.Send[int](m_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// The integer tag of the selected menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/tag
func (m_ MenuItemCell) SetTag(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTag:"), value)
}/* debug [instance_properties/setter]: tag */


// The width of the menu item’s text, measured in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/titleWidth
func (m_ MenuItemCell) TitleWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("titleWidth"))
	return rv
}/* debug [instance_properties/getter]: titleWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMenuItemCell */


