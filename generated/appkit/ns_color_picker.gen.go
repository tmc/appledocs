// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	ButtonToolTip() foundation.foundation.INSString
	SetButtonToolTip(value foundation.foundation.INSString)
	ColorPanel() IColorPanel
	SetColorPanel(value IColorPanel)
	MinContentSize() corefoundation.CGSize
	SetMinContentSize(value corefoundation.CGSize)
	ProvideNewButtonImage() IImage
	SetProvideNewButtonImage(value IImage)


	

	// methods:
	SetMode(mode ColorPanelMode)


}





// Alloc allocates a new instance without initialization.
func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.Send[ColorPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Overriden to set the color picker’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/setMode(_:)
func (c_ ColorPicker) SetMode(mode ColorPanelMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMode:"), mode)
}







// The tool tip that is shown when the mouse cursor is over the color picker’s button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/buttontooltip
func (c_ ColorPicker) ButtonToolTip() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("buttonToolTip"))
	return rv
}


// The tool tip that is shown when the mouse cursor is over the color picker’s button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/buttontooltip
func (c_ ColorPicker) SetButtonToolTip(value foundation.foundation.INSString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonToolTip:"), value)
}


// The color panel instance that owns the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/colorpanel
func (c_ ColorPicker) ColorPanel() IColorPanel {
	rv := objc.Send[ColorPanel](c_.ID, objc.Sel("colorPanel"))
	return rv
}


// The color panel instance that owns the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/colorpanel
func (c_ ColorPicker) SetColorPanel(value IColorPanel) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorPanel:"), value)
}


// The minimum content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/mincontentsize
func (c_ ColorPicker) MinContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("minContentSize"))
	return rv
}


// The minimum content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/mincontentsize
func (c_ ColorPicker) SetMinContentSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinContentSize:"), value)
}


// The button image used by the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/providenewbuttonimage
func (c_ ColorPicker) ProvideNewButtonImage() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("provideNewButtonImage"))
	return rv
}


// The button image used by the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/providenewbuttonimage
func (c_ ColorPicker) SetProvideNewButtonImage(value IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProvideNewButtonImage:"), value)
}








