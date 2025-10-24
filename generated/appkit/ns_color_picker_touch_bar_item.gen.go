// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSColorPickerTouchBarItem */


/* debug [class_header]: Header for NSColorPickerTouchBarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ColorPickerTouchBarItem */
// An interface definition for the [ColorPickerTouchBarItem] class.
type IColorPickerTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for ColorPickerTouchBarItem */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	AllowedColorSpaces() []ColorSpace
	SetAllowedColorSpaces(value []ColorSpace)
	Color() IColor
	SetColor(value IColor)
	ColorList() IColorList
	SetColorList(value IColorList)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Enabled() bool
	SetEnabled(value bool)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	Target() objc.ID
	SetTarget(value objc.ID)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ColorPickerTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ColorPickerTouchBarItem */
// Alloc allocates a new instance without initialization.
func (cc _ColorPickerTouchBarItemClass) Alloc() ColorPickerTouchBarItem {
	rv := objc.Send[ColorPickerTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ColorPickerTouchBarItem */
// A bar item that provides a system-defined color picker.
//
// For design guidance, see .


// A bar item that provides a system-defined color picker.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ColorPickerTouchBarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ColorPickerTouchBarItem */

// Creates a bar item with the standard color picker icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:)
func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifier(identifier TouchBarItemIdentifier /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorPickerWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorPickerWithIdentifier) */


// Creates a color picker bar item using the supplied image as its icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:buttonImage:)
func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifierButtonImage(identifier TouchBarItemIdentifier /* typedef */, image IImage) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorPickerWithIdentifier:buttonImage:"), identifier, image)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorPickerWithIdentifierButtonImage) */


// Creates a bar item with the standard stroke color picker icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/strokeColorPicker(withIdentifier:)
func (cc _ColorPickerTouchBarItemClass) StrokeColorPickerWithIdentifier(identifier TouchBarItemIdentifier /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("strokeColorPickerWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StrokeColorPickerWithIdentifier) */


// Creates a bar item with the standard text color picker icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/textColorPicker(withIdentifier:)
func (cc _ColorPickerTouchBarItemClass) TextColorPickerWithIdentifier(identifier TouchBarItemIdentifier /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("textColorPickerWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextColorPickerWithIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ColorPickerTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ColorPickerTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ColorPickerTouchBarItem */

// The selector on the target object that is invoked when a user interacts with the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/action
func (c_ ColorPickerTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The selector on the target object that is invoked when a user interacts with the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/action
func (c_ ColorPickerTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// Controls the color spaces that the color picker can produce.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/allowedColorSpaces
func (c_ ColorPickerTouchBarItem) AllowedColorSpaces() []ColorSpace {
	rv := objc.Send[[]ColorSpace](c_.ID, objc.Sel("allowedColorSpaces"))
	return rv
}/* debug [instance_properties/getter]: allowedColorSpaces */


// Controls the color spaces that the color picker can produce.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/allowedColorSpaces
func (c_ ColorPickerTouchBarItem) SetAllowedColorSpaces(value []ColorSpace) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedColorSpaces:"), nsArray)
}/* debug [instance_properties/setter]: allowedColorSpaces */


// The picker’s currently selected color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/color
func (c_ ColorPickerTouchBarItem) Color() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The picker’s currently selected color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/color
func (c_ ColorPickerTouchBarItem) SetColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// The list of colors displayed in the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c_ ColorPickerTouchBarItem) ColorList() IColorList {
	rv := objc.Send[ColorList](c_.ID, objc.Sel("colorList"))
	return rv
}/* debug [instance_properties/getter]: colorList */


// The list of colors displayed in the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c_ ColorPickerTouchBarItem) SetColorList(value IColorList) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorList:"), value)
}/* debug [instance_properties/setter]: colorList */


// The user-visible string identifying this item during touch bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/customizationLabel
func (c_ ColorPickerTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// The user-visible string identifying this item during touch bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/customizationLabel
func (c_ ColorPickerTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// A Boolean value that determines whether the color picker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/isEnabled
func (c_ ColorPickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that determines whether the color picker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/isEnabled
func (c_ ColorPickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/showsAlpha
func (c_ ColorPickerTouchBarItem) ShowsAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsAlpha"))
	return rv
}/* debug [instance_properties/getter]: showsAlpha */


// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/showsAlpha
func (c_ ColorPickerTouchBarItem) SetShowsAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsAlpha:"), value)
}/* debug [instance_properties/setter]: showsAlpha */


// An object that is notified when a user interacts with the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/target
func (c_ ColorPickerTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// An object that is notified when a user interacts with the color picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/target
func (c_ ColorPickerTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// A Boolean value that determines whether the color picker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/isenabled
func (c_ ColorPickerTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that determines whether the color picker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/isenabled
func (c_ ColorPickerTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSColorPickerTouchBarItem */



