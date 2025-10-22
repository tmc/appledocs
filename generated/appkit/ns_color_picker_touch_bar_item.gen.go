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
	AllowedColorSpaces() []ColorSpace
	SetAllowedColorSpaces(value []ColorSpace)
	ColorList() NSColorList
	SetColorList(value IColorList)
	Enabled() bool
	SetEnabled(value bool)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	Color() Color
	SetColor(value IColor)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	IsEnabled() bool
	SetIsEnabled(value bool)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
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
func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifier(identifier ITouchBarItemIdentifier) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorPickerWithIdentifier:"), identifier)
	return rv
}

// Creates a color picker bar item using the supplied image as its icon.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:buttonImage:)
func (cc _ColorPickerTouchBarItemClass) ColorPickerWithIdentifierButtonImage(identifier ITouchBarItemIdentifier, image IImage) unsafe.Pointer {
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
	// Convert Go slice to NSArray
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
}

// The list of colors displayed in the color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c_ ColorPickerTouchBarItem) ColorList() NSColorList {
	rv := objc.Send[NSColorList](c_.ID, objc.Sel("colorList"))
	return rv
}


// SetColorList sets the value of the colorList property.
// The list of colors displayed in the color picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c_ ColorPickerTouchBarItem) SetColorList(value IColorList) {
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

// The selector on the target object that is invoked when a user interacts with the color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/action
func (c_ ColorPickerTouchBarItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The selector on the target object that is invoked when a user interacts with the color picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/action
func (c_ ColorPickerTouchBarItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}

// The picker’s currently selected color.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/color
func (c_ ColorPickerTouchBarItem) Color() Color {
	rv := objc.Send[Color](c_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// The picker’s currently selected color.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/color
func (c_ ColorPickerTouchBarItem) SetColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColor:"), value)
}

// The user-visible string identifying this item during touch bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/customizationlabel
func (c_ ColorPickerTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during touch bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/customizationlabel
func (c_ ColorPickerTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}

// A Boolean value that determines whether the color picker is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/isenabled
func (c_ ColorPickerTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that determines whether the color picker is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/isenabled
func (c_ ColorPickerTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}

// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/showsalpha
func (c_ ColorPickerTouchBarItem) ShowsAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsAlpha"))
	return rv
}


// SetShowsAlpha sets the value of the showsAlpha property.
// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/showsalpha
func (c_ ColorPickerTouchBarItem) SetShowsAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsAlpha:"), value)
}

// An object that is notified when a user interacts with the color picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/target
func (c_ ColorPickerTouchBarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// An object that is notified when a user interacts with the color picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorpickertouchbaritem/target
func (c_ ColorPickerTouchBarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}



