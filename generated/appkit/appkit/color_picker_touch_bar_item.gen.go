// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorPickerTouchBarItem] class.
var (
	colorPickerTouchBarItemClass     _ColorPickerTouchBarItemClass
	colorPickerTouchBarItemClassOnce sync.Once
)

func getColorPickerTouchBarItemClass() _ColorPickerTouchBarItemClass {
	colorPickerTouchBarItemClassOnce.Do(func() {
		colorPickerTouchBarItemClass = _ColorPickerTouchBarItemClass{objc.GetClass("NSColorPickerTouchBarItem")}
	})
	return colorPickerTouchBarItemClass
}

type _ColorPickerTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [ColorPickerTouchBarItem] class.
type IColorPickerTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a system-defined color picker. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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




