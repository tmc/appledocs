// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AttachColorList(colorList IColorList)
	ViewSizeChanged(sender objectivec.IObject)
	ColorPanel() NSColorPanel
	ProvideNewButtonImage() Image
	ButtonToolTip() string
	SetButtonToolTip(value string)
	MinContentSize() coregraphics.CGSize
	SetMinContentSize(value coregraphics.CGSize)
}

// An abstract superclass that implements the default color picking protocol.
//
// The and protocols define a way to add color pickers (custom user interfaces for color selection) to the color panel.


// An abstract superclass that implements the default color picking protocol.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/attachColorList(_:)

func (c_ ColorPicker) AttachColorList(colorList IColorList) {
	objc.Send[objc.ID](c_.ID, objc.Sel("attachColorList:"), colorList)
}



// Overriden to respond to a size change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/viewSizeChanged(_:)

func (c_ ColorPicker) ViewSizeChanged(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewSizeChanged:"), sender)
}


// The color panel instance that owns the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/colorPanel

func (c_ ColorPicker) ColorPanel() NSColorPanel {
	rv := objc.Send[NSColorPanel](c_.ID, objc.Sel("colorPanel"))
	return rv
}


// The button image used by the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/provideNewButtonImage

func (c_ ColorPicker) ProvideNewButtonImage() Image {
	rv := objc.Send[Image](c_.ID, objc.Sel("provideNewButtonImage"))
	return rv
}


// The tool tip that is shown when the mouse cursor is over the color picker’s button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/buttontooltip

func (c_ ColorPicker) ButtonToolTip() string {
	rv := objc.Send[string](c_.ID, objc.Sel("buttonToolTip"))
	return rv
}


// The tool tip that is shown when the mouse cursor is over the color picker’s button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/buttontooltip

func (c_ ColorPicker) SetButtonToolTip(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonToolTip:"), objc.String(value))
}


// The minimum content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/mincontentsize

func (c_ ColorPicker) MinContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("minContentSize"))
	return rv
}


// The minimum content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/mincontentsize

func (c_ ColorPicker) SetMinContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinContentSize:"), value)
}



