// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	AlphaComponent() float64 /* primitive/slice/pointer. */
	BlackComponent() float64 /* primitive/slice/pointer. */
	BlueComponent() float64 /* primitive/slice/pointer. */
	BrightnessComponent() float64 /* primitive/slice/pointer. */
	CatalogNameComponent() objc.IObject /* cross-framework: ColorListName */
	CGColor() ColorRef /* not a class type */
	ColorNameComponent() objc.IObject /* cross-framework: ColorName */
	ColorSpace() IColorSpace
	ColorSpaceName() objc.IObject /* cross-framework: ColorSpaceName */
	CyanComponent() float64 /* primitive/slice/pointer. */
	GreenComponent() float64 /* primitive/slice/pointer. */
	HueComponent() float64 /* primitive/slice/pointer. */
	LinearExposure() float64 /* primitive/slice/pointer. */
	LocalizedCatalogNameComponent() objc.IObject /* cross-framework: NSString */
	LocalizedColorNameComponent() objc.IObject /* cross-framework: NSString */
	MagentaComponent() float64 /* primitive/slice/pointer. */
	NumberOfComponents() int /* primitive/slice/pointer. */
	PatternImage() IImage
	RedComponent() float64 /* primitive/slice/pointer. */
	SaturationComponent() float64 /* primitive/slice/pointer. */
	StandardDynamicRangeColor() IColor
	Type() ColorType
	WhiteComponent() float64 /* primitive/slice/pointer. */
	YellowComponent() float64 /* primitive/slice/pointer. */
	StandardDynamicRange() IColor
	SetStandardDynamicRange(value IColor)
	// methods:
	ColorByApplyingContentHeadroom(contentHeadroom float64 /* primitive/slice/pointer. */) IColor
	BlendedColorWithFractionOfColor(fraction float64 /* primitive/slice/pointer. */, color IColor) IColor
	DrawSwatchInRect(rect objc.IObject /* cross-framework Rect */)
	GetComponents(components corefoundation.CGFloat)
	GetCyanMagentaYellowBlackAlpha(cyan corefoundation.CGFloat, magenta corefoundation.CGFloat, yellow corefoundation.CGFloat, black corefoundation.CGFloat, alpha corefoundation.CGFloat)
	GetHueSaturationBrightnessAlpha(hue corefoundation.CGFloat, saturation corefoundation.CGFloat, brightness corefoundation.CGFloat, alpha corefoundation.CGFloat)
	GetRedGreenBlueAlpha(red corefoundation.CGFloat, green corefoundation.CGFloat, blue corefoundation.CGFloat, alpha corefoundation.CGFloat)
	GetWhiteAlpha(white corefoundation.CGFloat, alpha corefoundation.CGFloat)
	HighlightWithLevel(val float64 /* primitive/slice/pointer. */) IColor
	Set()
	SetFill()
	SetStroke()
	ShadowWithLevel(val float64 /* primitive/slice/pointer. */) IColor
	ColorUsingColorSpace(space IColorSpace) IColor
	ColorUsingType(type_ ColorType) IColor
	ColorWithAlphaComponent(alpha float64 /* primitive/slice/pointer. */) IColor
	ColorWithSystemEffect(systemEffect ColorSystemEffect) IColor
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



// Creates a color object from color data currently on the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(from:)
func NewColorFromPasteboard(pasteBoard IPasteboard) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorFromPasteboard:"), pasteBoard)
	return rv
}


// Creates a color object from the provided name, which corresponds to a color in the default asset catalog of the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(named:bundle:)
func NewColorNamedBundle(name objc.IObject /* cross-framework ColorName */, bundle objc.IObject /* cross-framework Bundle */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorNamed:bundle:"), name, bundle)
	return rv
}


// Creates a color object using the specified Core Graphics color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(cgColor:)
func NewColorWithCGColor(cgColor ColorRef /* not a class type */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithCGColor:"), cgColor)
	return rv
}


// Creates a color object using the specified asset catalog and color names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(catalogName:colorName:)
func NewColorWithCatalogNameColorName(listName objc.IObject /* cross-framework ColorListName */, colorName objc.IObject /* cross-framework ColorName */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}


// Creates a color object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(coder:)
func NewColorWithCoder(coder objc.IObject /* cross-framework Coder */) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Creates a color object with the specified color space, hue, saturation, brightness, and alpha channel values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:hue:saturation:brightness:alpha:)
func NewColorWithColorSpaceHueSaturationBrightnessAlpha(space IColorSpace, hue float64 /* primitive/slice/pointer. */, saturation float64 /* primitive/slice/pointer. */, brightness float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithColorSpace:hue:saturation:brightness:alpha:"), space, hue, saturation, brightness, alpha)
	return rv
}


// Creates a color object using the given opacity value and CMYK components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceCyan:magenta:yellow:black:alpha:)
func NewColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64 /* primitive/slice/pointer. */, magenta float64 /* primitive/slice/pointer. */, yellow float64 /* primitive/slice/pointer. */, black float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithDeviceCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
	return rv
}


// Creates a color object using the given opacity value and RGB components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceRed:green:blue:alpha:)
func NewColorWithDeviceRedGreenBlueAlpha(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithDeviceRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}


// Creates a color object using the given opacity and grayscale values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceWhite:alpha:)
func NewColorWithDeviceWhiteAlpha(white float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithDeviceWhite:alpha:"), white, alpha)
	return rv
}


// Creates a color object from the specified components in the Display P3 color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(displayP3Red:green:blue:alpha:)
func NewColorWithDisplayP3RedGreenBlueAlpha(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithDisplayP3Red:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}


// Creates a dynamic catalog color with a provider that’s used to resolve the exact color value, calculated on first use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(name:dynamicProvider:)
func NewColorWithNameDynamicProvider(colorName objc.IObject /* cross-framework ColorName */, dynamicProvider unsafe.Pointer) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithName:dynamicProvider:"), colorName, dynamicProvider)
	return rv
}


// Creates a color object that uses the specified image pattern to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(patternImage:)
func NewColorWithPatternImage(image IImage) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithPatternImage:"), image)
	return rv
}


// Generates an HDR color in the extended sRGB colorspace by applying an exposure to the SDR color defined by the red, green, and blue components. The , , and components have a nominal range of [0..1], is a value >= 0. To produce an HDR color, we process the given color in a linear color space, multiplying component values by . The produced color will have a equal to the linearized exposure value. Each whole value of exposure produces a color that is twice as bright.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:exposure:)
func NewColorWithRedGreenBlueAlphaExposure(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */, exposure float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithRed:green:blue:alpha:exposure:"), red, green, blue, alpha, exposure)
	return rv
}


// Generates an HDR color in the extended sRGB colorspace by applying an exposure to the SDR color defined by the red, green, and blue components. The , , and components have a nominal range of [0..1], is a value >= 1. To produce an HDR color, we process the given color in a linear color space, multiplying component values by . The produced color will have a equal to . Each doubling of produces a color that is twice as bright.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:linearExposure:)
func NewColorWithRedGreenBlueAlphaLinearExposure(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */, linearExposure float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithRed:green:blue:alpha:linearExposure:"), red, green, blue, alpha, linearExposure)
	return rv
}


// Creates a color object with the specified brightness and alpha channel values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(white:alpha:)
func NewColorWithWhiteAlpha(white float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithWhite:alpha:"), white, alpha)
	return rv
}



// Creates a color object using the specified asset catalog and color names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(catalogName:colorName:)
func (cc _ColorClass) ColorWithCatalogNameColorName(listName objc.IObject /* cross-framework ColorListName */, colorName objc.IObject /* cross-framework ColorName */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}


// Creates a color object using the specified Core Graphics color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(cgColor:)
func (cc _ColorClass) ColorWithCGColor(cgColor ColorRef /* not a class type */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithCGColor:"), cgColor)
	return rv
}


// Creates a color object with the specified color space, hue, saturation, brightness, and alpha channel values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:hue:saturation:brightness:alpha:)
func (cc _ColorClass) ColorWithColorSpaceHueSaturationBrightnessAlpha(space IColorSpace, hue float64 /* primitive/slice/pointer. */, saturation float64 /* primitive/slice/pointer. */, brightness float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithColorSpace:hue:saturation:brightness:alpha:"), space, hue, saturation, brightness, alpha)
	return rv
}


// Creates a color object using the given opacity value and CMYK components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceCyan:magenta:yellow:black:alpha:)
func (cc _ColorClass) ColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64 /* primitive/slice/pointer. */, magenta float64 /* primitive/slice/pointer. */, yellow float64 /* primitive/slice/pointer. */, black float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithDeviceCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
	return rv
}


// Creates a color object using the given opacity value and RGB components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceRed:green:blue:alpha:)
func (cc _ColorClass) ColorWithDeviceRedGreenBlueAlpha(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithDeviceRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}


// Creates a color object using the given opacity and grayscale values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceWhite:alpha:)
func (cc _ColorClass) ColorWithDeviceWhiteAlpha(white float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithDeviceWhite:alpha:"), white, alpha)
	return rv
}


// Creates a color object from the specified components in the Display P3 color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(displayP3Red:green:blue:alpha:)
func (cc _ColorClass) ColorWithDisplayP3RedGreenBlueAlpha(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithDisplayP3Red:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}


// Creates a color object from color data currently on the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(from:)
func (cc _ColorClass) ColorFromPasteboard(pasteBoard IPasteboard) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorFromPasteboard:"), pasteBoard)
	return rv
}


// Creates a dynamic catalog color with a provider that’s used to resolve the exact color value, calculated on first use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(name:dynamicProvider:)
func (cc _ColorClass) ColorWithNameDynamicProvider(colorName objc.IObject /* cross-framework ColorName */, dynamicProvider unsafe.Pointer) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithName:dynamicProvider:"), colorName, dynamicProvider)
	return rv
}


// Creates a color object from the provided name, which corresponds to a color in the default asset catalog of the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(named:bundle:)
func (cc _ColorClass) ColorNamedBundle(name objc.IObject /* cross-framework ColorName */, bundle objc.IObject /* cross-framework Bundle */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorNamed:bundle:"), name, bundle)
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


// Generates an HDR color in the extended sRGB colorspace by applying an exposure to the SDR color defined by the red, green, and blue components. The , , and components have a nominal range of [0..1], is a value >= 0. To produce an HDR color, we process the given color in a linear color space, multiplying component values by . The produced color will have a equal to the linearized exposure value. Each whole value of exposure produces a color that is twice as bright.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:exposure:)
func (cc _ColorClass) ColorWithRedGreenBlueAlphaExposure(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */, exposure float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:alpha:exposure:"), red, green, blue, alpha, exposure)
	return rv
}


// Generates an HDR color in the extended sRGB colorspace by applying an exposure to the SDR color defined by the red, green, and blue components. The , , and components have a nominal range of [0..1], is a value >= 1. To produce an HDR color, we process the given color in a linear color space, multiplying component values by . The produced color will have a equal to . Each doubling of produces a color that is twice as bright.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:linearExposure:)
func (cc _ColorClass) ColorWithRedGreenBlueAlphaLinearExposure(red float64 /* primitive/slice/pointer. */, green float64 /* primitive/slice/pointer. */, blue float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */, linearExposure float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:alpha:linearExposure:"), red, green, blue, alpha, linearExposure)
	return rv
}


// Creates a color object with the specified brightness and alpha channel values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(white:alpha:)
func (cc _ColorClass) ColorWithWhiteAlpha(white float64 /* primitive/slice/pointer. */, alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithWhite:alpha:"), white, alpha)
	return rv
}


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/black
func (cc _ColorClass) BlackColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("blackColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blue
func (cc _ColorClass) BlueColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("blueColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brown
func (cc _ColorClass) BrownColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("brownColor"))
	return rv
}

// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/clear
func (cc _ColorClass) ClearColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("clearColor"))
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

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyan
func (cc _ColorClass) CyanColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("cyanColor"))
	return rv
}

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/darkGray
func (cc _ColorClass) DarkGrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("darkGrayColor"))
	return rv
}

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/gray
func (cc _ColorClass) GrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("grayColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/green
func (cc _ColorClass) GreenColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("greenColor"))
	return rv
}

// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (cc _ColorClass) IgnoresAlpha() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("ignoresAlpha"))
	return rv
}

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/lightGray
func (cc _ColorClass) LightGrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("lightGrayColor"))
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

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/orange
func (cc _ColorClass) OrangeColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("orangeColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/purple
func (cc _ColorClass) PurpleColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("purpleColor"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quinaryLabel
func (cc _ColorClass) QuinaryLabelColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("quinaryLabelColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/red
func (cc _ColorClass) RedColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("redColor"))
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemFill
func (cc _ColorClass) SystemFillColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("systemFillColor"))
	return rv
}

// Returns a color object for yellow that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemYellow
func (cc _ColorClass) SystemYellowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("systemYellowColor"))
	return rv
}

// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/white
func (cc _ColorClass) WhiteColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("whiteColor"))
	return rv
}

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellow
func (cc _ColorClass) YellowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("yellowColor"))
	return rv
}

// Reinterpret the color by applying a new without changing the color components. Changing the redefines the color relative to a different peak white, changing its behavior under tone mapping and the result of calling . The new color will have a >= 1.0. If called on a color with a color space that does not support extended range, or does not have an equivalent extended range counterpart, this will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/applyingContentHeadroom(_:)
func (c_ Color) ColorByApplyingContentHeadroom(contentHeadroom float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorByApplyingContentHeadroom:"), contentHeadroom)
	return rv
}


// Creates a new color object whose component values are a weighted sum of the current color object and the specified color object’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blended(withFraction:of:)
func (c_ Color) BlendedColorWithFractionOfColor(fraction float64 /* primitive/slice/pointer. */, color IColor) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blendedColorWithFraction:ofColor:"), fraction, color)
	return rv
}


// Draws the current color in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/drawSwatch(in:)
func (c_ Color) DrawSwatchInRect(rect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawSwatchInRect:"), rect)
}


// Returns the components of the color as an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getComponents(_:)
func (c_ Color) GetComponents(components corefoundation.CGFloat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getComponents:"), components)
}


// Returns the color object’s CMYK and opacity values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getCyan(_:magenta:yellow:black:alpha:)
func (c_ Color) GetCyanMagentaYellowBlackAlpha(cyan corefoundation.CGFloat, magenta corefoundation.CGFloat, yellow corefoundation.CGFloat, black corefoundation.CGFloat, alpha corefoundation.CGFloat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}


// Returns the color object’s HSB component and opacity values in the respective arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getHue(_:saturation:brightness:alpha:)
func (c_ Color) GetHueSaturationBrightnessAlpha(hue corefoundation.CGFloat, saturation corefoundation.CGFloat, brightness corefoundation.CGFloat, alpha corefoundation.CGFloat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
}


// Returns the color object’s RGB component and opacity values in the respective arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getRed(_:green:blue:alpha:)
func (c_ Color) GetRedGreenBlueAlpha(red corefoundation.CGFloat, green corefoundation.CGFloat, blue corefoundation.CGFloat, alpha corefoundation.CGFloat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getRed:green:blue:alpha:"), red, green, blue, alpha)
}


// Returns the grayscale and alpha values of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/getWhite(_:alpha:)
func (c_ Color) GetWhiteAlpha(white corefoundation.CGFloat, alpha corefoundation.CGFloat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getWhite:alpha:"), white, alpha)
}


// Creates a new color object that represents a blend between the current color and the highlight color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/highlight(withLevel:)
func (c_ Color) HighlightWithLevel(val float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("highlightWithLevel:"), val)
	return rv
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
func (c_ Color) ShadowWithLevel(val float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("shadowWithLevel:"), val)
	return rv
}


// Creates a new color object representing the color of the current color object in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpace(_:)
func (c_ Color) ColorUsingColorSpace(space IColorSpace) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorUsingColorSpace:"), space)
	return rv
}


// Returns a version of the color object that is compatible with the specified color type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/usingType(_:)
func (c_ Color) ColorUsingType(type_ ColorType) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorUsingType:"), type_)
	return rv
}


// Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/withAlphaComponent(_:)
func (c_ Color) ColorWithAlphaComponent(alpha float64 /* primitive/slice/pointer. */) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("colorWithAlphaComponent:"), alpha)
	return rv
}


// Returns a new color object that represents the current color modified to include the specified visual effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/withSystemEffect(_:)
func (c_ Color) ColorWithSystemEffect(systemEffect ColorSystemEffect) IColor {
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


// The alpha (opacity) component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alphaComponent
func (c_ Color) AlphaComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("alphaComponent"))
	return rv
}


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/black
func (c_ Color) BlackColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blackColor"))
	return rv
}


// The black component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blackComponent
func (c_ Color) BlackComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("blackComponent"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blue
func (c_ Color) BlueColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blueColor"))
	return rv
}


// The blue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blueComponent
func (c_ Color) BlueComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("blueComponent"))
	return rv
}


// The brightness component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brightnessComponent
func (c_ Color) BrightnessComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("brightnessComponent"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brown
func (c_ Color) BrownColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("brownColor"))
	return rv
}


// The catalog containing the color’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/catalogNameComponent
func (c_ Color) CatalogNameComponent() objc.IObject /* cross-framework: ColorListName */ {
	rv := objc.Send[ColorListName](c_.ID, objc.Sel("catalogNameComponent"))
	return rv
}


// The Core Graphics color object corresponding to the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cgColor
func (c_ Color) CGColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](c_.ID, objc.Sel("CGColor"))
	return rv
}


// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/clear
func (c_ Color) ClearColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("clearColor"))
	return rv
}


// The name of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorNameComponent
func (c_ Color) ColorNameComponent() objc.IObject /* cross-framework: ColorName */ {
	rv := objc.Send[ColorName](c_.ID, objc.Sel("colorNameComponent"))
	return rv
}


// The color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorSpace
func (c_ Color) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// The name of the color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorSpaceName
func (c_ Color) ColorSpaceName() objc.IObject /* cross-framework: ColorSpaceName */ {
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


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyan
func (c_ Color) CyanColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("cyanColor"))
	return rv
}


// The cyan component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyanComponent
func (c_ Color) CyanComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("cyanComponent"))
	return rv
}


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/darkGray
func (c_ Color) DarkGrayColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("darkGrayColor"))
	return rv
}


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/gray
func (c_ Color) GrayColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("grayColor"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/green
func (c_ Color) GreenColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("greenColor"))
	return rv
}


// The green component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/greenComponent
func (c_ Color) GreenComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("greenComponent"))
	return rv
}


// The hue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/hueComponent
func (c_ Color) HueComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("hueComponent"))
	return rv
}


// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (c_ Color) IgnoresAlpha() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("ignoresAlpha"))
	return rv
}


// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (c_ Color) SetIgnoresAlpha(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIgnoresAlpha:"), value)
}


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/lightGray
func (c_ Color) LightGrayColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("lightGrayColor"))
	return rv
}


// For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/linearExposure
func (c_ Color) LinearExposure() float64 /* primitive/slice/pointer. */ {
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


// The localized version of the catalog name containing the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/localizedCatalogNameComponent
func (c_ Color) LocalizedCatalogNameComponent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedCatalogNameComponent"))
	return rv
}


// The localized version of the color name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/localizedColorNameComponent
func (c_ Color) LocalizedColorNameComponent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedColorNameComponent"))
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


// The magenta component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/magentaComponent
func (c_ Color) MagentaComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("magentaComponent"))
	return rv
}


// The number of components in the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/numberOfComponents
func (c_ Color) NumberOfComponents() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfComponents"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/orange
func (c_ Color) OrangeColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("orangeColor"))
	return rv
}


// The pattern image used to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/patternImage
func (c_ Color) PatternImage() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("patternImage"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/purple
func (c_ Color) PurpleColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("purpleColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quinaryLabel
func (c_ Color) QuinaryLabelColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("quinaryLabelColor"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/red
func (c_ Color) RedColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("redColor"))
	return rv
}


// The red component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/redComponent
func (c_ Color) RedComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("redComponent"))
	return rv
}


// The saturation component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/saturationComponent
func (c_ Color) SaturationComponent() float64 /* primitive/slice/pointer. */ {
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


// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s is > 1, then this will return the base SDR color. If the color is not an HDR color, this will return .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/standardDynamicRange
func (c_ Color) StandardDynamicRangeColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("standardDynamicRangeColor"))
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemFill
func (c_ Color) SystemFillColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("systemFillColor"))
	return rv
}


// Returns a color object for yellow that automatically adapts to vibrancy and accessibility settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemYellow
func (c_ Color) SystemYellowColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("systemYellowColor"))
	return rv
}


// The type of the color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/type
func (c_ Color) Type() ColorType {
	rv := objc.Send[ColorType](c_.ID, objc.Sel("type"))
	return rv
}


// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/white
func (c_ Color) WhiteColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("whiteColor"))
	return rv
}


// The white component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/whiteComponent
func (c_ Color) WhiteComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("whiteComponent"))
	return rv
}


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellow
func (c_ Color) YellowColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("yellowColor"))
	return rv
}


// The yellow component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellowComponent
func (c_ Color) YellowComponent() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("yellowComponent"))
	return rv
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


