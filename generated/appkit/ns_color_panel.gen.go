// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	Color() IColor
	SetColor(value IColor)
	Continuous() bool
	SetContinuous(value bool)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	Alpha() float64
	SetAlpha(value float64)
	IsContinuous() bool
	SetIsContinuous(value bool)
	MaximumLinearExposure() float64
	SetMaximumLinearExposure(value float64)
	Mode() objectivec.IObject
	SetMode(value objectivec.IObject)


	

	// methods:
	SetAction(selector objc.SEL)
	SetTarget(target objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.Send[ColorPanel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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










// Drags a color into a destination view from the specified source view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/dragColor(_:with:from:)
func (cc _ColorPanelClass) DragColorWithEventFromView(color IColor, event IEvent, sourceView IView) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("dragColor:withEvent:fromView:"), color, event, sourceView)
	return rv
}












// Sets the color panel’s action message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/setAction(_:)
func (c_ ColorPanel) SetAction(selector objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), selector)
}


// Sets the target of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/setTarget(_:)
func (c_ ColorPanel) SetTarget(target objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), target)
}







// The accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/accessoryView
func (c_ ColorPanel) AccessoryView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("accessoryView"))
	return rv
}


// The accessory view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/accessoryView
func (c_ ColorPanel) SetAccessoryView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccessoryView:"), value)
}


// The color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/color
func (c_ ColorPanel) Color() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("color"))
	return rv
}


// The color of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/color
func (c_ ColorPanel) SetColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColor:"), value)
}


// A Boolean value indicating whether the receiver continuously sends the action message to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/isContinuous
func (c_ ColorPanel) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}


// A Boolean value indicating whether the receiver continuously sends the action message to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/isContinuous
func (c_ ColorPanel) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}


// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/showsAlpha
func (c_ ColorPanel) ShowsAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsAlpha"))
	return rv
}


// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/showsAlpha
func (c_ ColorPanel) SetShowsAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsAlpha:"), value)
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
func (c_ ColorPanel) Mode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("mode"))
	return rv
}


// The mode of the receiver the mode is one of the modes allowed by the color mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpanel/mode-swift.property
func (c_ ColorPanel) SetMode(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMode:"), value)
}








