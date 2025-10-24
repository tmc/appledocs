// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSColorPicker */


/* debug [class_header]: Header for NSColorPicker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ColorPicker */
// An interface definition for the [ColorPicker] class.
type IColorPicker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ColorPicker */
	// properties:
	ButtonToolTip() objc.IObject /* cross-framework: NSString */
	SetButtonToolTip(value objc.IObject /* cross-framework: NSString */)
	ColorPanel() IColorPanel
	SetColorPanel(value IColorPanel)
	MinContentSize() Size /* not a class type */
	SetMinContentSize(value Size /* not a class type */)
	ProvideNewButtonImage() IImage
	SetProvideNewButtonImage(value IImage)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ColorPicker */
	// methods:
	SetMode(mode ColorPanelMode)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ColorPicker */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ColorPicker */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ColorPicker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ColorPicker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ColorPicker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ColorPicker */

// Overriden to set the color picker’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPicker/setMode(_:)
func (c_ ColorPicker) SetMode(mode ColorPanelMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMode:"), mode)
}/* debug [instance_methods/method]: SetMode */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ColorPicker */

// The tool tip that is shown when the mouse cursor is over the color picker’s button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/buttontooltip
func (c_ ColorPicker) ButtonToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("buttonToolTip"))
	return rv
}/* debug [instance_properties/getter]: buttonToolTip */


// The tool tip that is shown when the mouse cursor is over the color picker’s button image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/buttontooltip
func (c_ ColorPicker) SetButtonToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonToolTip:"), value)
}/* debug [instance_properties/setter]: buttonToolTip */


// The color panel instance that owns the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/colorpanel
func (c_ ColorPicker) ColorPanel() IColorPanel {
	rv := objc.Send[ColorPanel](c_.ID, objc.Sel("colorPanel"))
	return rv
}/* debug [instance_properties/getter]: colorPanel */


// The color panel instance that owns the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/colorpanel
func (c_ ColorPicker) SetColorPanel(value IColorPanel) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorPanel:"), value)
}/* debug [instance_properties/setter]: colorPanel */


// The minimum content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/mincontentsize
func (c_ ColorPicker) MinContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("minContentSize"))
	return rv
}/* debug [instance_properties/getter]: minContentSize */


// The minimum content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/mincontentsize
func (c_ ColorPicker) SetMinContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinContentSize:"), value)
}/* debug [instance_properties/setter]: minContentSize */


// The button image used by the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/providenewbuttonimage
func (c_ ColorPicker) ProvideNewButtonImage() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("provideNewButtonImage"))
	return rv
}/* debug [instance_properties/getter]: provideNewButtonImage */


// The button image used by the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpicker/providenewbuttonimage
func (c_ ColorPicker) SetProvideNewButtonImage(value IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProvideNewButtonImage:"), value)
}/* debug [instance_properties/setter]: provideNewButtonImage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSColorPicker */



