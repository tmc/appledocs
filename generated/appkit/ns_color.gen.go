// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	BlackComponent() float64
	BlueComponent() float64
	BrightnessComponent() float64
	CatalogNameComponent() unsafe.Pointer
	ColorNameComponent() unsafe.Pointer
	ColorSpace() ColorSpace
	ColorSpaceName() ColorSpaceName
	CyanComponent() float64
	GreenComponent() float64
	HueComponent() float64
	LinearExposure() float64
	LocalizedColorNameComponent() string
	NumberOfComponents() int
	RedComponent() float64
	SaturationComponent() float64
	Type() NSColorType
	YellowComponent() float64
	AlphaComponent() float64
	SetAlphaComponent(value float64)
	CgColor() IColor
	SetCgColor(value IColor)
	LocalizedCatalogNameComponent() string
	SetLocalizedCatalogNameComponent(value string)
	MagentaComponent() float64
	SetMagentaComponent(value float64)
	StandardDynamicRange() IColor
	SetStandardDynamicRange(value IColor)
	WhiteComponent() float64
	SetWhiteComponent(value float64)
	ColorByApplyingContentHeadroom(contentHeadroom float64) IColor
	BlendedColorWithFractionOfColor(fraction float64, color IColor) IColor
	GetComponents(components coregraphics.float64)
	GetCyanMagentaYellowBlackAlpha(cyan coregraphics.float64, magenta coregraphics.float64, yellow coregraphics.float64, black coregraphics.float64, alpha coregraphics.float64)
	GetHueSaturationBrightnessAlpha(hue coregraphics.float64, saturation coregraphics.float64, brightness coregraphics.float64, alpha coregraphics.float64)
	GetRedGreenBlueAlpha(red coregraphics.float64, green coregraphics.float64, blue coregraphics.float64, alpha coregraphics.float64)
	Set()
	SetFill()
	SetStroke()
	ShadowWithLevel(val float64) IColor
	ColorUsingColorSpace(space ColorSpace) IColor
	ColorWithSystemEffect(systemEffect NSColorSystemEffect) IColor
	WriteToPasteboard(pasteBoard IPasteboard)
}

// An object that stores color data and sometimes opacity (alpha value).
//
// Many methods in AppKit require you to specify color data using an object; when drawing you use them to set the current fill and stroke colors. Color objects are immutable and thread-safe. You can create color objects in many ways: Load colors from an asset catalog. Colors created from assets can adapt automatically to system appearance changes. Use the semantic colors for custom UI elements, so that they match the appearance of other AppKit views; see . Use the adaptable system colors, such as , when you want a specific tint that looks correct in both light and dark environments. Create a color object from another object, such as a Core Graphics representation of a color, or a Core Image color. Create a color from an object, and paint a repeating pattern instead of using a solid color. Create a color by applying a transform to another object. For example, you might perform a blend operation between two colors, or you might create a color that represents the same color, but in a different color space. Create custom colors using raw component values, and a variety of color spaces, when you need to represent user-specified colors. For user-specified colors, you can also display a color panel and let the user specify the color. For information about color panels, see .


// An object that stores color data and sometimes opacity (alpha value).
//
// [Full Topic]
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



// Creates a color object that uses the specified image pattern to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(patternImage:)
func NewColorWithPatternImage(image IImage) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithPatternImage:"), image)
	return rv
}



// Creates a color object that uses the specified image pattern to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(patternImage:)
func (cc _ColorClass) ColorWithPatternImage(image IImage) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithPatternImage:"), image)
	return rv
}


// The color to use for the background of large controls, such as scroll views or table views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlBackgroundColor
func (cc _ColorClass) ControlBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlBackgroundColor"))
	return rv
}

// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (cc _ColorClass) IgnoresAlpha() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("ignoresAlpha"))
	return rv
}

// The color to use for links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/linkColor
func (cc _ColorClass) LinkColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("linkColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/magenta
func (cc _ColorClass) MagentaColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("magentaColor"))
	return rv
}

// The patterned color to use for the background of a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/scrubberTexturedBackground
func (cc _ColorClass) ScrubberTexturedBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("scrubberTexturedBackgroundColor"))
	return rv
}

// Returns a color object for blue that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemBlue
func (cc _ColorClass) SystemBlueColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("systemBlueColor"))
	return rv
}

// Reinterpret the color by applying a new without changing the color components. Changing the redefines the color relative to a different peak white, changing its behavior under tone mapping and the result of calling . The new color will have a >= 1.0. If called on a color with a color space that does not support extended range, or does not have an equivalent extended range counterpart, this will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/applyingContentHeadroom(_:)
func (c_ Color) ColorByApplyingContentHeadroom(contentHeadroom float64) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorByApplyingContentHeadroom:"), contentHeadroom)
	return rv
}


// Creates a new color object whose component values are a weighted sum of the current color object and the specified color object’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blended(withFraction:of:)
func (c_ Color) BlendedColorWithFractionOfColor(fraction float64, color IColor) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blendedColorWithFraction:ofColor:"), fraction, color)
	return rv
}


// Returns the components of the color as an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getComponents(_:)
func (c_ Color) GetComponents(components coregraphics.float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getComponents:"), components)
}


// Returns the color object’s CMYK and opacity values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getCyan(_:magenta:yellow:black:alpha:)
func (c_ Color) GetCyanMagentaYellowBlackAlpha(cyan coregraphics.float64, magenta coregraphics.float64, yellow coregraphics.float64, black coregraphics.float64, alpha coregraphics.float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}


// Returns the color object’s HSB component and opacity values in the respective arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getHue(_:saturation:brightness:alpha:)
func (c_ Color) GetHueSaturationBrightnessAlpha(hue coregraphics.float64, saturation coregraphics.float64, brightness coregraphics.float64, alpha coregraphics.float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
}


// Returns the color object’s RGB component and opacity values in the respective arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getRed(_:green:blue:alpha:)
func (c_ Color) GetRedGreenBlueAlpha(red coregraphics.float64, green coregraphics.float64, blue coregraphics.float64, alpha coregraphics.float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getRed:green:blue:alpha:"), red, green, blue, alpha)
}


// Sets the color of subsequent drawing to the color that the color object represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/set()
func (c_ Color) Set() {
	objc.Send[objc.ID](c_.ID, objc.Sel("set"))
}


// Sets the fill color of subsequent drawing to the color object’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/setFill()
func (c_ Color) SetFill() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFill"))
}


// Sets the stroke color of subsequent drawing to the color object’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/setStroke()
func (c_ Color) SetStroke() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStroke"))
}


// Creates a new color object that represents a blend between the current color and the shadow color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/shadow(withLevel:)
func (c_ Color) ShadowWithLevel(val float64) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("shadowWithLevel:"), val)
	return rv
}


// Creates a new color object representing the color of the current color object in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpace(_:)
func (c_ Color) ColorUsingColorSpace(space ColorSpace) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorUsingColorSpace:"), space)
	return rv
}


// Returns a new color object that represents the current color modified to include the specified visual effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/withSystemEffect(_:)
func (c_ Color) ColorWithSystemEffect(systemEffect NSColorSystemEffect) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorWithSystemEffect:"), systemEffect)
	return rv
}


// Writes the color object’s data to the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/write(to:)
func (c_ Color) WriteToPasteboard(pasteBoard IPasteboard) {
	objc.Send[objc.ID](c_.ID, objc.Sel("writeToPasteboard:"), pasteBoard)
}


// The black component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blackComponent
func (c_ Color) BlackComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blackComponent"))
	return rv
}


// The blue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blueComponent
func (c_ Color) BlueComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blueComponent"))
	return rv
}


// The brightness component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brightnessComponent
func (c_ Color) BrightnessComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("brightnessComponent"))
	return rv
}


// The catalog containing the color’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/catalogNameComponent
func (c_ Color) CatalogNameComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("catalogNameComponent"))
	return rv
}


// The name of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorNameComponent
func (c_ Color) ColorNameComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorNameComponent"))
	return rv
}


// The color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorSpace
func (c_ Color) ColorSpace() ColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// The name of the color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorSpaceName
func (c_ Color) ColorSpaceName() ColorSpaceName {
	rv := objc.Send[ColorSpaceName](c_.ID, objc.Sel("colorSpaceName"))
	return rv
}


// The color to use for the background of large controls, such as scroll views or table views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlBackgroundColor
func (c_ Color) ControlBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlBackgroundColor"))
	return rv
}


// The cyan component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyanComponent
func (c_ Color) CyanComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("cyanComponent"))
	return rv
}


// The green component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/greenComponent
func (c_ Color) GreenComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("greenComponent"))
	return rv
}


// The hue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/hueComponent
func (c_ Color) HueComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("hueComponent"))
	return rv
}


// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (c_ Color) IgnoresAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ignoresAlpha"))
	return rv
}


// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (c_ Color) SetIgnoresAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIgnoresAlpha:"), value)
}


// For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/linearExposure
func (c_ Color) LinearExposure() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("linearExposure"))
	return rv
}


// The color to use for links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/linkColor
func (c_ Color) LinkColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("linkColor"))
	return rv
}


// The localized version of the color name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/localizedColorNameComponent
func (c_ Color) LocalizedColorNameComponent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedColorNameComponent"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/magenta
func (c_ Color) MagentaColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("magentaColor"))
	return rv
}


// The number of components in the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/numberOfComponents
func (c_ Color) NumberOfComponents() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfComponents"))
	return rv
}


// The red component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/redComponent
func (c_ Color) RedComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("redComponent"))
	return rv
}


// The saturation component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/saturationComponent
func (c_ Color) SaturationComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("saturationComponent"))
	return rv
}


// The patterned color to use for the background of a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/scrubberTexturedBackground
func (c_ Color) ScrubberTexturedBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("scrubberTexturedBackgroundColor"))
	return rv
}


// Returns a color object for blue that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemBlue
func (c_ Color) SystemBlueColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("systemBlueColor"))
	return rv
}


// The type of the color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/type
func (c_ Color) Type() NSColorType {
	rv := objc.Send[NSColorType](c_.ID, objc.Sel("type"))
	return rv
}


// The yellow component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellowComponent
func (c_ Color) YellowComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("yellowComponent"))
	return rv
}


// The alpha (opacity) component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/alphacomponent
func (c_ Color) AlphaComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alphaComponent"))
	return rv
}


// The alpha (opacity) component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/alphacomponent
func (c_ Color) SetAlphaComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlphaComponent:"), value)
}


// The Core Graphics color object corresponding to the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/cgcolor
func (c_ Color) CgColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("cgColor"))
	return rv
}


// The Core Graphics color object corresponding to the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/cgcolor
func (c_ Color) SetCgColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCgColor:"), value)
}


// The localized version of the catalog name containing the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/localizedcatalognamecomponent
func (c_ Color) LocalizedCatalogNameComponent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedCatalogNameComponent"))
	return rv
}


// The localized version of the catalog name containing the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/localizedcatalognamecomponent
func (c_ Color) SetLocalizedCatalogNameComponent(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedCatalogNameComponent:"), objc.String(value))
}


// The magenta component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/magentacomponent
func (c_ Color) MagentaComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("magentaComponent"))
	return rv
}


// The magenta component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/magentacomponent
func (c_ Color) SetMagentaComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMagentaComponent:"), value)
}


// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/standarddynamicrange
func (c_ Color) StandardDynamicRange() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("standardDynamicRange"))
	return rv
}


// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/standarddynamicrange
func (c_ Color) SetStandardDynamicRange(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandardDynamicRange:"), value)
}


// The white component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/whitecomponent
func (c_ Color) WhiteComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("whiteComponent"))
	return rv
}


// The white component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/whitecomponent
func (c_ Color) SetWhiteComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteComponent:"), value)
}


