// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ColorPanel] class.
var (
	ColorPanelClass     _ColorPanelClass
	ColorPanelClassOnce sync.Once
)

func getColorPanelClass() _ColorPanelClass {
	ColorPanelClassOnce.Do(func() {
		ColorPanelClass = _ColorPanelClass{objc.GetClass("NSColorPanel")}
	})
	return ColorPanelClass
}

type _ColorPanelClass struct {
	class objc.Class
}

// An interface definition for the [ColorPanel] class.
type IColorPanel interface {
	IPanel
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	Alpha() float64
	SetAlpha(value float64)
	Color() IColor
	SetColor(value IColor)
	IsContinuous() bool
	SetIsContinuous(value bool)
	MaximumLinearExposure() float64
	SetMaximumLinearExposure(value float64)
	Mode() unsafe.Pointer
	SetMode(value unsafe.Pointer)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	// methods:
}

// A standard user interface for selecting color in an app.
//
// provides a number of standard color selection modes and, with the and protocols, allows an app to add its own color selection modes. It also allows the user to save swatches containing frequently used colors.


// A standard user interface for selecting color in an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel
type ColorPanel struct {
	Panel
}

// ColorPanelFrom constructs a [ColorPanel] from an unsafe.Pointer.
//
// A standard user interface for selecting color in an app.
func ColorPanelFrom(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		Panel: PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorPanelClass) New() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorPanel) Init() ColorPanel {
	rv := objc.Send[ColorPanel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorPanel) Autorelease() ColorPanel {
	rv := objc.Send[ColorPanel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorPanel creates a new ColorPanel instance.
func NewColorPanel() ColorPanel {
	return getColorPanelClass().New()
}



// The accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/accessoryview
func (c_ ColorPanel) AccessoryView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("accessoryView"))
	return rv
}


// The accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/accessoryview
func (c_ ColorPanel) SetAccessoryView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccessoryView:"), value)
}


// The receiver’s current alpha value based on its opacity slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/alpha
func (c_ ColorPanel) Alpha() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alpha"))
	return rv
}


// The receiver’s current alpha value based on its opacity slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/alpha
func (c_ ColorPanel) SetAlpha(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}


// The color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/color
func (c_ ColorPanel) Color() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("color"))
	return rv
}


// The color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/color
func (c_ ColorPanel) SetColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColor:"), value)
}


// A Boolean value indicating whether the receiver continuously sends the action message to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/iscontinuous
func (c_ ColorPanel) IsContinuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuous"))
	return rv
}


// A Boolean value indicating whether the receiver continuously sends the action message to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/iscontinuous
func (c_ ColorPanel) SetIsContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuous:"), value)
}


// The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/maximumlinearexposure
func (c_ ColorPanel) MaximumLinearExposure() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("maximumLinearExposure"))
	return rv
}


// The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/maximumlinearexposure
func (c_ ColorPanel) SetMaximumLinearExposure(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumLinearExposure:"), value)
}


// The mode of the receiver the mode is one of the modes allowed by the color mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/mode-swift.property
func (c_ ColorPanel) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mode"))
	return rv
}


// The mode of the receiver the mode is one of the modes allowed by the color mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/mode-swift.property
func (c_ ColorPanel) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMode:"), value)
}


// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/showsalpha
func (c_ ColorPanel) ShowsAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsAlpha"))
	return rv
}


// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/showsalpha
func (c_ ColorPanel) SetShowsAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsAlpha:"), value)
}



