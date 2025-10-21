// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ColorPicker] class.
var (
	ColorPickerClass     _ColorPickerClass
	ColorPickerClassOnce sync.Once
)

func getColorPickerClass() _ColorPickerClass {
	ColorPickerClassOnce.Do(func() {
		ColorPickerClass = _ColorPickerClass{objc.GetClass("NSColorPicker")}
	})
	return ColorPickerClass
}

type _ColorPickerClass struct {
	class objc.Class
}

// An interface definition for the [ColorPicker] class.
type IColorPicker interface {
	objectivec.IObject
	AttachColorList(colorList unsafe.Pointer)
	ViewSizeChanged(sender objc.ID)
}

// An abstract superclass that implements the default color picking protocol.
//
// The and protocols define a way to add color pickers (custom user interfaces for color selection) to the color panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker
type ColorPicker struct {
	objectivec.Object
}

// ColorPickerFrom constructs a [ColorPicker] from an unsafe.Pointer.
//
// An abstract superclass that implements the default color picking protocol.
func ColorPickerFrom(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorPickerClass) New() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorPicker) Init() ColorPicker {
	rv := objc.Send[ColorPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorPicker) Autorelease() ColorPicker {
	rv := objc.Send[ColorPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorPicker creates a new ColorPicker instance.
func NewColorPicker() ColorPicker {
	return getColorPickerClass().New()
}


// Overriden to attach a color list to a color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/attachColorList(_:)
func (c_ ColorPicker) AttachColorList(colorList unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("attachColorList:"), colorList)
}

// Overriden to respond to a size change.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/viewSizeChanged(_:)
func (c_ ColorPicker) ViewSizeChanged(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewSizeChanged:"), sender)
}

// The color panel instance that owns the color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/colorPanel
func (c_ ColorPicker) ColorPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorPanel"))
	return rv
}

// The button image used by the color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/provideNewButtonImage
func (c_ ColorPicker) ProvideNewButtonImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("provideNewButtonImage"))
	return rv
}



