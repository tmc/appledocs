// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorPickerTouchBarItem] class.
var (
	ColorPickerTouchBarItemClass     _ColorPickerTouchBarItemClass
	ColorPickerTouchBarItemClassOnce sync.Once
)

func getColorPickerTouchBarItemClass() _ColorPickerTouchBarItemClass {
	ColorPickerTouchBarItemClassOnce.Do(func() {
		ColorPickerTouchBarItemClass = _ColorPickerTouchBarItemClass{objc.GetClass("NSColorPickerTouchBarItem")}
	})
	return ColorPickerTouchBarItemClass
}

type _ColorPickerTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [ColorPickerTouchBarItem] class.
type IColorPickerTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a system-defined color picker.
//
// For design guidance, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem
type ColorPickerTouchBarItem struct {
	TouchBarItem
}

// ColorPickerTouchBarItemFrom constructs a [ColorPickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a system-defined color picker.
func ColorPickerTouchBarItemFrom(ptr unsafe.Pointer) ColorPickerTouchBarItem {
	return ColorPickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPickerTouchBarItemClass) Alloc() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorPickerTouchBarItemClass) New() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorPickerTouchBarItem) Init() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorPickerTouchBarItem) Autorelease() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorPickerTouchBarItem creates a new ColorPickerTouchBarItem instance.
func NewColorPickerTouchBarItem() ColorPickerTouchBarItem {
	return getColorPickerTouchBarItemClass().New()
}

// Creates a bar item with the standard color picker icon.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:)
func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifier(identifier unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorPickerWithIdentifier:"), identifier)
	return rv
}

// Creates a color picker bar item using the supplied image as its icon.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:buttonImage:)
func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifierButtonImage(identifier unsafe.Pointer, image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorPickerWithIdentifier:buttonImage:"), identifier, image)
	return rv
}

// Controls the color spaces that the color picker can produce.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/allowedColorSpaces
func (c_ ColorPickerTouchBarItem) AllowedColorSpaces() []ColorSpace {
	rv := objc.Send[[]ColorSpace](c_.ID, objc.Sel("allowedColorSpaces"))
	return rv
}

// SetAllowedColorSpaces sets the value of the allowedColorSpaces property.
// Controls the color spaces that the color picker can produce.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/allowedColorSpaces
func (c_ ColorPickerTouchBarItem) SetAllowedColorSpaces(value []ColorSpace) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedColorSpaces:"), value)
}

// The list of colors displayed in the color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c_ ColorPickerTouchBarItem) ColorList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorList"))
	return rv
}

// SetColorList sets the value of the colorList property.
// The list of colors displayed in the color picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c_ ColorPickerTouchBarItem) SetColorList(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorList:"), value)
}

// A Boolean value that determines whether the color picker is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/isEnabled
func (c_ ColorPickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}

// SetEnabled sets the value of the enabled property.
// A Boolean value that determines whether the color picker is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/isEnabled
func (c_ ColorPickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}
