// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Color] class.
var (
	ColorClass     _ColorClass
	ColorClassOnce sync.Once
)

func getColorClass() _ColorClass {
	ColorClassOnce.Do(func() {
		ColorClass = _ColorClass{objc.GetClass("NSColor")}
	})
	return ColorClass
}

type _ColorClass struct {
	class objc.Class
}

// An interface definition for the [Color] class.
type IColor interface {
	objectivec.IObject
	ColorByApplyingContentHeadroom(contentHeadroom float64) unsafe.Pointer
	BlendedColorWithFractionOfColor(fraction float64, color unsafe.Pointer) unsafe.Pointer
	GetComponents(components float64)
	GetCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64)
	GetRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64)
	Set()
	SetFill()
	ColorUsingColorSpace(space unsafe.Pointer) unsafe.Pointer
	ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer
	ColorUsingType(type_ unsafe.Pointer) unsafe.Pointer
	ColorWithAlphaComponent(alpha float64) unsafe.Pointer
	ColorWithSystemEffect(systemEffect unsafe.Pointer) unsafe.Pointer
}

// An object that stores color data and sometimes opacity (alpha value).
//
// Many methods in AppKit require you to specify color data using an object; when drawing you use them to set the current fill and stroke colors. Color objects are immutable and thread-safe. You can create color objects in many ways: Load colors from an asset catalog. Colors created from assets can adapt automatically to system appearance changes. Use the semantic colors for custom UI elements, so that they match the appearance of other AppKit views; see . Use the adaptable system colors, such as , when you want a specific tint that looks correct in both light and dark environments. Create a color object from another object, such as a Core Graphics representation of a color, or a Core Image color. Create a color from an object, and paint a repeating pattern instead of using a solid color. Create a color by applying a transform to another object. For example, you might perform a blend operation between two colors, or you might create a color that represents the same color, but in a different color space. Create custom colors using raw component values, and a variety of color spaces, when you need to represent user-specified colors. For user-specified colors, you can also display a color panel and let the user specify the color. For information about color panels, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor
type Color struct {
	objectivec.Object
}

// ColorFrom constructs a [Color] from an unsafe.Pointer.
//
// An object that stores color data and sometimes opacity (alpha value).
func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorClass) Alloc() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorClass) New() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Color) Init() Color {
	rv := objc.Send[Color](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Color) Autorelease() Color {
	rv := objc.Send[Color](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColor creates a new Color instance.
func NewColor() Color {
	return getColorClass().New()
}




// Returns the color object specified by the given control tint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(for:)
func NewColorForControlTint(controlTint unsafe.Pointer) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorForControlTint:"), controlTint)
	return rv
}



// Creates a color object using the specified asset catalog and color names.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(catalogName:colorName:)
func NewColorWithCatalogNameColorName(listName unsafe.Pointer, colorName unsafe.Pointer) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}



// Creates a color object from the specified components of the given color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:components:count:)
func NewColorWithColorSpaceComponentsCount(space unsafe.Pointer, components unsafe.Pointer, numberOfComponents int) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithColorSpace:components:count:"), space, components, numberOfComponents)
	return rv
}


// Creates a color object using the specified asset catalog and color names.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(catalogName:colorName:)
func (cc _ColorClass) ColorWithCatalogNameColorName(listName unsafe.Pointer, colorName unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}

// Creates a color object from the specified components of the given color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:components:count:)
func (cc _ColorClass) ColorWithColorSpaceComponentsCount(space unsafe.Pointer, components unsafe.Pointer, numberOfComponents int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithColorSpace:components:count:"), space, components, numberOfComponents)
	return rv
}

// Returns the color object specified by the given control tint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(for:)
func (cc _ColorClass) ColorForControlTint(controlTint unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorForControlTint:"), controlTint)
	return rv
}

// The color to use for text in a selected control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlTextColor
func (cc _ColorClass) AlternateSelectedControlTextColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("alternateSelectedControlTextColor"))
	return rv
}
// The system color used for the dark edge of the shadow dropped from controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlDarkShadowColor
func (cc _ColorClass) ControlDarkShadowColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("controlDarkShadowColor"))
	return rv
}
// The current system control tint color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/currentControlTint
func (cc _ColorClass) CurrentControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("currentControlTint"))
	return rv
}
// The primary color to use for text labels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
func (cc _ColorClass) LabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("labelColor"))
	return rv
}
// The quaternary color to use for text labels and separators.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
func (cc _ColorClass) QuaternaryLabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("quaternaryLabelColor"))
	return rv
}
// The secondary color to use for text labels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
func (cc _ColorClass) SecondaryLabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("secondaryLabelColor"))
	return rv
}
// Returns a color object for blue that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemBlue
func (cc _ColorClass) SystemBlueColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("systemBlueColor"))
	return rv
}
// Returns a color object for mint that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemMint
func (cc _ColorClass) SystemMintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("systemMintColor"))
	return rv
}
// Returns a color object for yellow that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemYellow
func (cc _ColorClass) SystemYellowColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("systemYellowColor"))
	return rv
}
// The tertiary color to use for text labels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
func (cc _ColorClass) TertiaryLabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tertiaryLabelColor"))
	return rv
}
// The color to use for the background area behind text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textBackgroundColor
func (cc _ColorClass) TextBackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("textBackgroundColor"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textInsertionPointColor
func (cc _ColorClass) TextInsertionPointColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("textInsertionPointColor"))
	return rv
}
// Reinterpret the color by applying a new without changing the color components. Changing the redefines the color relative to a different peak white, changing its behavior under tone mapping and the result of calling . The new color will have a >= 1.0. If called on a color with a color space that does not support extended range, or does not have an equivalent extended range counterpart, this will return .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/applyingContentHeadroom(_:)
func (c_ Color) ColorByApplyingContentHeadroom(contentHeadroom float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorByApplyingContentHeadroom:"), contentHeadroom)
	return rv
}

// Creates a new color object whose component values are a weighted sum of the current color object and the specified color object’s.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blended(withFraction:of:)
func (c_ Color) BlendedColorWithFractionOfColor(fraction float64, color unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("blendedColorWithFraction:ofColor:"), fraction, color)
	return rv
}

// Returns the components of the color as an array.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getComponents(_:)
func (c_ Color) GetComponents(components float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getComponents:"), components)
}

// Returns the color object’s CMYK and opacity values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getCyan(_:magenta:yellow:black:alpha:)
func (c_ Color) GetCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}

// Returns the color object’s RGB component and opacity values in the respective arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getRed(_:green:blue:alpha:)
func (c_ Color) GetRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getRed:green:blue:alpha:"), red, green, blue, alpha)
}

// Sets the color of subsequent drawing to the color that the color object represents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/set()
func (c_ Color) Set() {
	objc.Send[objc.ID](c_.ID, objc.Sel("set"))
}

// Sets the fill color of subsequent drawing to the color object’s color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/setFill()
func (c_ Color) SetFill() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFill"))
}

// Creates a new color object representing the color of the current color object in the specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpace(_:)
func (c_ Color) ColorUsingColorSpace(space unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorUsingColorSpace:"), space)
	return rv
}

// Creates a new color object whose color is the same as the receiver’s, except that the new color object is in the specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpaceName(_:)
func (c_ Color) ColorUsingColorSpaceName(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorUsingColorSpaceName:"), name)
	return rv
}

// Returns a version of the color object that is compatible with the specified color type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingType(_:)
func (c_ Color) ColorUsingType(type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorUsingType:"), type_)
	return rv
}

// Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/withAlphaComponent(_:)
func (c_ Color) ColorWithAlphaComponent(alpha float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorWithAlphaComponent:"), alpha)
	return rv
}

// Returns a new color object that represents the current color modified to include the specified visual effect.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/withSystemEffect(_:)
func (c_ Color) ColorWithSystemEffect(systemEffect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorWithSystemEffect:"), systemEffect)
	return rv
}

// The color to use for text in a selected control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlTextColor
func (c_ Color) AlternateSelectedControlTextColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("alternateSelectedControlTextColor"))
	return rv
}

// The black component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blackComponent
func (c_ Color) BlackComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blackComponent"))
	return rv
}

// The brightness component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brightnessComponent
func (c_ Color) BrightnessComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("brightnessComponent"))
	return rv
}

// The catalog containing the color’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/catalogNameComponent
func (c_ Color) CatalogNameComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("catalogNameComponent"))
	return rv
}

// The color space associated with the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorSpace
func (c_ Color) ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorSpace"))
	return rv
}

// The system color used for the dark edge of the shadow dropped from controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlDarkShadowColor
func (c_ Color) ControlDarkShadowColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlDarkShadowColor"))
	return rv
}

// The current system control tint color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/currentControlTint
func (c_ Color) CurrentControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentControlTint"))
	return rv
}

// The cyan component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyanComponent
func (c_ Color) CyanComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("cyanComponent"))
	return rv
}

// The hue component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/hueComponent
func (c_ Color) HueComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("hueComponent"))
	return rv
}

// The primary color to use for text labels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
func (c_ Color) LabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labelColor"))
	return rv
}

// The localized version of the catalog name containing the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/localizedCatalogNameComponent
func (c_ Color) LocalizedCatalogNameComponent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedCatalogNameComponent"))
	return rv
}

// The number of components in the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/numberOfComponents
func (c_ Color) NumberOfComponents() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfComponents"))
	return rv
}

// The pattern image used to paint the target area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/patternImage
func (c_ Color) PatternImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("patternImage"))
	return rv
}

// The quaternary color to use for text labels and separators.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
func (c_ Color) QuaternaryLabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("quaternaryLabelColor"))
	return rv
}

// The red component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/redComponent
func (c_ Color) RedComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("redComponent"))
	return rv
}

// The saturation component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/saturationComponent
func (c_ Color) SaturationComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("saturationComponent"))
	return rv
}

// The secondary color to use for text labels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
func (c_ Color) SecondaryLabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("secondaryLabelColor"))
	return rv
}

// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s is > 1, then this will return the base SDR color. If the color is not an HDR color, this will return .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/standardDynamicRange
func (c_ Color) StandardDynamicRangeColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("standardDynamicRangeColor"))
	return rv
}

// Returns a color object for blue that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemBlue
func (c_ Color) SystemBlueColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemBlueColor"))
	return rv
}

// Returns a color object for mint that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemMint
func (c_ Color) SystemMintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemMintColor"))
	return rv
}

// Returns a color object for yellow that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemYellow
func (c_ Color) SystemYellowColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemYellowColor"))
	return rv
}

// The tertiary color to use for text labels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
func (c_ Color) TertiaryLabelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tertiaryLabelColor"))
	return rv
}

// The color to use for the background area behind text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textBackgroundColor
func (c_ Color) TextBackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("textBackgroundColor"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textInsertionPointColor
func (c_ Color) TextInsertionPointColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("textInsertionPointColor"))
	return rv
}

// The type of the color object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/type
func (c_ Color) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}

// The yellow component value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellowComponent
func (c_ Color) YellowComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("yellowComponent"))
	return rv
}


