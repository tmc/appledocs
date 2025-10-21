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
	AttachColorList(colorList unsafe.Pointer)
	DetachColorList(colorList unsafe.Pointer)
	SetAction(selector objc.SEL)
	SetTarget(target objc.ID)
}

// A standard user interface for selecting color in an app.
//
// provides a number of standard color selection modes and, with the and protocols, allows an app to add its own color selection modes. It also allows the user to save swatches containing frequently used colors.
//
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

// Drags a color into a destination view from the specified source view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/dragColor(_:with:from:)
func (cc _ColorPanelClass) DragColorWithEventFromView(color unsafe.Pointer, event unsafe.Pointer, sourceView unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("dragColor:withEvent:fromView:"), color, event, sourceView)
	return rv
}

// Determines which color selection modes are available in an application’s .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/setPickerMask(_:)
func (cc _ColorPanelClass) SetPickerMask(mask unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("setPickerMask:"), mask)
}

// Specifies the color panel’s initial picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/setPickerMode(_:)
func (cc _ColorPanelClass) SetPickerMode(mode unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("setPickerMode:"), mode)
}

// Returns the shared instance, creating it if necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/shared
func (cc _ColorPanelClass) SharedColorPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sharedColorPanel"))
	return rv
}

// Returns a Boolean value indicating whether the has been created already.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/sharedColorPanelExists
func (cc _ColorPanelClass) SharedColorPanelExists() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("sharedColorPanelExists"))
	return rv
}

// Adds the list of objects specified to all the color pickers in the receiver that display color lists by invoking on all color pickers in the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/attachColorList(_:)
func (c_ ColorPanel) AttachColorList(colorList unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("attachColorList:"), colorList)
}

// Removes the list of colors from all the color pickers in the receiver that display color lists by invoking on all color pickers in the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/detachColorList(_:)
func (c_ ColorPanel) DetachColorList(colorList unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("detachColorList:"), colorList)
}

// Sets the color panel’s action message.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/setAction(_:)
func (c_ ColorPanel) SetAction(selector objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), selector)
}

// Sets the target of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/setTarget(_:)
func (c_ ColorPanel) SetTarget(target objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), target)
}

// The accessory view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/accessoryView
func (c_ ColorPanel) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("accessoryView"))
	return rv
}

// SetAccessoryView sets the value of the accessoryView property.
// The accessory view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/accessoryView
func (c_ ColorPanel) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccessoryView:"), value)
}

// The receiver’s current alpha value based on its opacity slider.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/alpha
func (c_ ColorPanel) Alpha() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alpha"))
	return rv
}

// The color of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/color
func (c_ ColorPanel) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("color"))
	return rv
}

// SetColor sets the value of the color property.
// The color of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/color
func (c_ ColorPanel) SetColor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColor:"), value)
}

// A Boolean value indicating whether the receiver continuously sends the action message to the target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/isContinuous
func (c_ ColorPanel) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}

// SetContinuous sets the value of the continuous property.
// A Boolean value indicating whether the receiver continuously sends the action message to the target.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/isContinuous
func (c_ ColorPanel) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}

// The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/maximumLinearExposure
func (c_ ColorPanel) MaximumLinearExposure() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("maximumLinearExposure"))
	return rv
}

// SetMaximumLinearExposure sets the value of the maximumLinearExposure property.
// The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/maximumLinearExposure
func (c_ ColorPanel) SetMaximumLinearExposure(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumLinearExposure:"), value)
}

// The mode of the receiver the mode is one of the modes allowed by the color mask.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/mode-swift.property
func (c_ ColorPanel) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mode"))
	return rv
}

// SetMode sets the value of the mode property.
// The mode of the receiver the mode is one of the modes allowed by the color mask.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/mode-swift.property
func (c_ ColorPanel) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMode:"), value)
}

// Returns the shared instance, creating it if necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/shared
func (c_ ColorPanel) SharedColorPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sharedColorPanel"))
	return rv
}

// Returns a Boolean value indicating whether the has been created already.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/sharedColorPanelExists
func (c_ ColorPanel) SharedColorPanelExists() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sharedColorPanelExists"))
	return rv
}

// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/showsAlpha
func (c_ ColorPanel) ShowsAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsAlpha"))
	return rv
}

// SetShowsAlpha sets the value of the showsAlpha property.
// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/showsAlpha
func (c_ ColorPanel) SetShowsAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsAlpha:"), value)
}
