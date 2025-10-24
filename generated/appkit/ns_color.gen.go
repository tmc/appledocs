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

/* debug [class.gen.go]: Generating class NSColor */


/* debug [class_header]: Header for NSColor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Color */
// An interface definition for the [Color] class.
type IColor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Color */
	// properties:
	ColorSpaceName() ColorSpaceName /* typedef */
	PatternImage() IImage
	AlphaComponent() float64
	SetAlphaComponent(value float64)
	BlackComponent() float64
	SetBlackComponent(value float64)
	BlueComponent() float64
	SetBlueComponent(value float64)
	BrightnessComponent() float64
	SetBrightnessComponent(value float64)
	CatalogNameComponent() objectivec.IObject
	SetCatalogNameComponent(value objectivec.IObject)
	CgColor() IColor
	SetCgColor(value IColor)
	ColorNameComponent() objectivec.IObject
	SetColorNameComponent(value objectivec.IObject)
	ColorSpace() IColorSpace
	SetColorSpace(value IColorSpace)
	CyanComponent() float64
	SetCyanComponent(value float64)
	GreenComponent() float64
	SetGreenComponent(value float64)
	HueComponent() float64
	SetHueComponent(value float64)
	LinearExposure() float64
	SetLinearExposure(value float64)
	LocalizedCatalogNameComponent() objc.IObject /* cross-framework: NSString */
	SetLocalizedCatalogNameComponent(value objc.IObject /* cross-framework: NSString */)
	LocalizedColorNameComponent() objc.IObject /* cross-framework: NSString */
	SetLocalizedColorNameComponent(value objc.IObject /* cross-framework: NSString */)
	MagentaComponent() float64
	SetMagentaComponent(value float64)
	NumberOfComponents() int
	SetNumberOfComponents(value int)
	RedComponent() float64
	SetRedComponent(value float64)
	SaturationComponent() float64
	SetSaturationComponent(value float64)
	StandardDynamicRange() IColor
	SetStandardDynamicRange(value IColor)
	Type() objectivec.IObject
	SetType(value objectivec.IObject)
	WhiteComponent() float64
	SetWhiteComponent(value float64)
	YellowComponent() float64
	SetYellowComponent(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Color */
	// methods:
	SetFill()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Color */
// Alloc allocates a new instance without initialization.
func (cc _ColorClass) Alloc() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Color */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Color */

// Returns the color object specified by the given control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(for:)
func NewColorForControlTint(controlTint ControlTint) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorForControlTint:"), controlTint)
	return rv
}/* debug [class_init_methods/constructor]: NewColorForControlTint */


// Creates a color object from the specified components of the given color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:components:count:)
func NewColorWithColorSpaceComponentsCount(space IColorSpace, components corefoundation.CGFloat, numberOfComponents int) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithColorSpace:components:count:"), space, components, numberOfComponents)
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithColorSpaceComponentsCount */


// Creates a color object that uses the specified image pattern to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(patternImage:)
func NewColorWithPatternImage(image IImage) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithPatternImage:"), image)
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithPatternImage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Color */

// Creates a color object from the specified components of the given color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:components:count:)
func (cc _ColorClass) ColorWithColorSpaceComponentsCount(space IColorSpace, components corefoundation.CGFloat, numberOfComponents int) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithColorSpace:components:count:"), space, components, numberOfComponents)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithColorSpaceComponentsCount) */


// Returns the color object specified by the given control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(for:)
func (cc _ColorClass) ColorForControlTint(controlTint ControlTint) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorForControlTint:"), controlTint)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorForControlTint) */


// Creates a color object that uses the specified image pattern to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/init(patternImage:)
func (cc _ColorClass) ColorWithPatternImage(image IImage) IColor {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("colorWithPatternImage:"), image)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithPatternImage) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Color */

// The system color used for the face of a selected control in a list or table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlColor
func (cc _ColorClass) AlternateSelectedControlColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("alternateSelectedControlColor"))
	return rv
}/* debug [class_properties_class/property]: alternateSelectedControlColor */

// The color to use for text in a selected control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlTextColor
func (cc _ColorClass) AlternateSelectedControlTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("alternateSelectedControlTextColor"))
	return rv
}/* debug [class_properties_class/property]: alternateSelectedControlTextColor */

// The colors to use for alternating content, typically found in table views and collection views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternatingContentBackgroundColors
func (cc _ColorClass) AlternatingContentBackgroundColors() []Color {
	rv := objc.Send[[]Color](objc.ID(cc.class), objc.Sel("alternatingContentBackgroundColors"))
	return rv
}/* debug [class_properties_class/property]: alternatingContentBackgroundColors */

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/black
func (cc _ColorClass) BlackColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("blackColor"))
	return rv
}/* debug [class_properties_class/property]: blackColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blue
func (cc _ColorClass) BlueColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("blueColor"))
	return rv
}/* debug [class_properties_class/property]: blueColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brown
func (cc _ColorClass) BrownColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("brownColor"))
	return rv
}/* debug [class_properties_class/property]: brownColor */

// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/clear
func (cc _ColorClass) ClearColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("clearColor"))
	return rv
}/* debug [class_properties_class/property]: clearColor */

// The user’s current accent color preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlAccentColor
func (cc _ColorClass) ControlAccentColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlAccentColor"))
	return rv
}/* debug [class_properties_class/property]: controlAccentColor */

// An array containing the system specified background colors for alternating rows in tables and lists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlAlternatingRowBackgroundColors
func (cc _ColorClass) ControlAlternatingRowBackgroundColors() []Color {
	rv := objc.Send[[]Color](objc.ID(cc.class), objc.Sel("controlAlternatingRowBackgroundColors"))
	return rv
}/* debug [class_properties_class/property]: controlAlternatingRowBackgroundColors */

// The color to use for the background of large controls, such as scroll views or table views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlBackgroundColor
func (cc _ColorClass) ControlBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: controlBackgroundColor */

// The color to use for the flat surfaces of a control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlColor
func (cc _ColorClass) ControlColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlColor"))
	return rv
}/* debug [class_properties_class/property]: controlColor */

// The system color used for the dark edge of the shadow dropped from controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlDarkShadowColor
func (cc _ColorClass) ControlDarkShadowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlDarkShadowColor"))
	return rv
}/* debug [class_properties_class/property]: controlDarkShadowColor */

// The system color used for the highlighted bezels of controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlHighlightColor
func (cc _ColorClass) ControlHighlightColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlHighlightColor"))
	return rv
}/* debug [class_properties_class/property]: controlHighlightColor */

// The system color used for light highlights in controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlLightHighlightColor
func (cc _ColorClass) ControlLightHighlightColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlLightHighlightColor"))
	return rv
}/* debug [class_properties_class/property]: controlLightHighlightColor */

// The system color used for the shadows dropped from controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlShadowColor
func (cc _ColorClass) ControlShadowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlShadowColor"))
	return rv
}/* debug [class_properties_class/property]: controlShadowColor */

// The color to use for text on enabled controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlTextColor
func (cc _ColorClass) ControlTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("controlTextColor"))
	return rv
}/* debug [class_properties_class/property]: controlTextColor */

// The current system control tint color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/currentControlTint
func (cc _ColorClass) CurrentControlTint() ControlTint {
	rv := objc.Send[ControlTint](objc.ID(cc.class), objc.Sel("currentControlTint"))
	return rv
}/* debug [class_properties_class/property]: currentControlTint */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyan
func (cc _ColorClass) CyanColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("cyanColor"))
	return rv
}/* debug [class_properties_class/property]: cyanColor */

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/darkGray
func (cc _ColorClass) DarkGrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("darkGrayColor"))
	return rv
}/* debug [class_properties_class/property]: darkGrayColor */

// The color to use for text on disabled controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/disabledControlTextColor
func (cc _ColorClass) DisabledControlTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("disabledControlTextColor"))
	return rv
}/* debug [class_properties_class/property]: disabledControlTextColor */

// The highlight color to use for the bubble that shows inline search result values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/findHighlightColor
func (cc _ColorClass) FindHighlightColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("findHighlightColor"))
	return rv
}/* debug [class_properties_class/property]: findHighlightColor */

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/gray
func (cc _ColorClass) GrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("grayColor"))
	return rv
}/* debug [class_properties_class/property]: grayColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/green
func (cc _ColorClass) GreenColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("greenColor"))
	return rv
}/* debug [class_properties_class/property]: greenColor */

// The color to use for the optional gridlines, such as those in a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/gridColor
func (cc _ColorClass) GridColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("gridColor"))
	return rv
}/* debug [class_properties_class/property]: gridColor */

// The system color used as the background color for header cells in table views and outline views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/headerColor
func (cc _ColorClass) HeaderColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("headerColor"))
	return rv
}/* debug [class_properties_class/property]: headerColor */

// The color to use for text in header cells in table views and outline views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/headerTextColor
func (cc _ColorClass) HeaderTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("headerTextColor"))
	return rv
}/* debug [class_properties_class/property]: headerTextColor */

// The color to use as a virtual light source on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/highlightColor
func (cc _ColorClass) HighlightColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("highlightColor"))
	return rv
}/* debug [class_properties_class/property]: highlightColor */

// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (cc _ColorClass) IgnoresAlpha() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("ignoresAlpha"))
	return rv
}/* debug [class_properties_class/property]: ignoresAlpha */

// The color to use for the keyboard focus ring around controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/keyboardFocusIndicatorColor
func (cc _ColorClass) KeyboardFocusIndicatorColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("keyboardFocusIndicatorColor"))
	return rv
}/* debug [class_properties_class/property]: keyboardFocusIndicatorColor */

// The system color used for the flat surface of a slider knob that hasn’t been selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/knobColor
func (cc _ColorClass) KnobColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("knobColor"))
	return rv
}/* debug [class_properties_class/property]: knobColor */

// The primary color to use for text labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
func (cc _ColorClass) LabelColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("labelColor"))
	return rv
}/* debug [class_properties_class/property]: labelColor */

// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/lightGray
func (cc _ColorClass) LightGrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("lightGrayColor"))
	return rv
}/* debug [class_properties_class/property]: lightGrayColor */

// The color to use for links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/linkColor
func (cc _ColorClass) LinkColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("linkColor"))
	return rv
}/* debug [class_properties_class/property]: linkColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/magenta
func (cc _ColorClass) MagentaColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("magentaColor"))
	return rv
}/* debug [class_properties_class/property]: magentaColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/orange
func (cc _ColorClass) OrangeColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("orangeColor"))
	return rv
}/* debug [class_properties_class/property]: orangeColor */

// The color to use for placeholder text in controls or text views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/placeholderTextColor
func (cc _ColorClass) PlaceholderTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("placeholderTextColor"))
	return rv
}/* debug [class_properties_class/property]: placeholderTextColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/purple
func (cc _ColorClass) PurpleColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("purpleColor"))
	return rv
}/* debug [class_properties_class/property]: purpleColor */

// The quaternary color to use for text labels and separators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
func (cc _ColorClass) QuaternaryLabelColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("quaternaryLabelColor"))
	return rv
}/* debug [class_properties_class/property]: quaternaryLabelColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quaternarySystemFill
func (cc _ColorClass) QuaternarySystemFillColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("quaternarySystemFillColor"))
	return rv
}/* debug [class_properties_class/property]: quaternarySystemFillColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quinaryLabel
func (cc _ColorClass) QuinaryLabelColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("quinaryLabelColor"))
	return rv
}/* debug [class_properties_class/property]: quinaryLabelColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quinarySystemFill
func (cc _ColorClass) QuinarySystemFillColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("quinarySystemFillColor"))
	return rv
}/* debug [class_properties_class/property]: quinarySystemFillColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/red
func (cc _ColorClass) RedColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("redColor"))
	return rv
}/* debug [class_properties_class/property]: redColor */

// The system color used for scroll “bars”—that is, for the groove in which a scroller’s knob moves
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/scrollBarColor
func (cc _ColorClass) ScrollBarColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("scrollBarColor"))
	return rv
}/* debug [class_properties_class/property]: scrollBarColor */

// The patterned color to use for the background of a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/scrubberTexturedBackground
func (cc _ColorClass) ScrubberTexturedBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("scrubberTexturedBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: scrubberTexturedBackgroundColor */

// The secondary color to use for text labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
func (cc _ColorClass) SecondaryLabelColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("secondaryLabelColor"))
	return rv
}/* debug [class_properties_class/property]: secondaryLabelColor */

// The color used for selected controls in non-key views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondarySelectedControlColor
func (cc _ColorClass) SecondarySelectedControlColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("secondarySelectedControlColor"))
	return rv
}/* debug [class_properties_class/property]: secondarySelectedControlColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondarySystemFill
func (cc _ColorClass) SecondarySystemFillColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("secondarySystemFillColor"))
	return rv
}/* debug [class_properties_class/property]: secondarySystemFillColor */

// The color to use for the background of selected and emphasized content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedContentBackgroundColor
func (cc _ColorClass) SelectedContentBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedContentBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: selectedContentBackgroundColor */

// The color to use for the face of a selected control—that is, a control that has been clicked or is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlColor
func (cc _ColorClass) SelectedControlColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedControlColor"))
	return rv
}/* debug [class_properties_class/property]: selectedControlColor */

// The color to use for text in a selected control—that is, a control being clicked or dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlTextColor
func (cc _ColorClass) SelectedControlTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedControlTextColor"))
	return rv
}/* debug [class_properties_class/property]: selectedControlTextColor */

// The system color used for the slider knob when it is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedKnobColor
func (cc _ColorClass) SelectedKnobColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedKnobColor"))
	return rv
}/* debug [class_properties_class/property]: selectedKnobColor */

// The color to use for the face of selected menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedMenuItemColor
func (cc _ColorClass) SelectedMenuItemColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedMenuItemColor"))
	return rv
}/* debug [class_properties_class/property]: selectedMenuItemColor */

// The color to use for the text in menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedMenuItemTextColor
func (cc _ColorClass) SelectedMenuItemTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedMenuItemTextColor"))
	return rv
}/* debug [class_properties_class/property]: selectedMenuItemTextColor */

// The color to use for the background of selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedTextBackgroundColor
func (cc _ColorClass) SelectedTextBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedTextBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: selectedTextBackgroundColor */

// The color to use for selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedTextColor
func (cc _ColorClass) SelectedTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("selectedTextColor"))
	return rv
}/* debug [class_properties_class/property]: selectedTextColor */

// The color to use for separators between different sections of content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/separatorColor
func (cc _ColorClass) SeparatorColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("separatorColor"))
	return rv
}/* debug [class_properties_class/property]: separatorColor */

// The color to use for virtual shadows cast by raised objects on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/shadowColor
func (cc _ColorClass) ShadowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("shadowColor"))
	return rv
}/* debug [class_properties_class/property]: shadowColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemFill
func (cc _ColorClass) SystemFillColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("systemFillColor"))
	return rv
}/* debug [class_properties_class/property]: systemFillColor */

// The tertiary color to use for text labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
func (cc _ColorClass) TertiaryLabelColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("tertiaryLabelColor"))
	return rv
}/* debug [class_properties_class/property]: tertiaryLabelColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/tertiarySystemFill
func (cc _ColorClass) TertiarySystemFillColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("tertiarySystemFillColor"))
	return rv
}/* debug [class_properties_class/property]: tertiarySystemFillColor */

// The color to use for the background area behind text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textBackgroundColor
func (cc _ColorClass) TextBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("textBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: textBackgroundColor */

// The color to use for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textColor
func (cc _ColorClass) TextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("textColor"))
	return rv
}/* debug [class_properties_class/property]: textColor */

// The color to use in the area beneath your window’s views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/underPageBackgroundColor
func (cc _ColorClass) UnderPageBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("underPageBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: underPageBackgroundColor */

// The color to use for selected and unemphasized content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedContentBackgroundColor
func (cc _ColorClass) UnemphasizedSelectedContentBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("unemphasizedSelectedContentBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: unemphasizedSelectedContentBackgroundColor */

// The color to use for the text background in an unemphasized context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedTextBackgroundColor
func (cc _ColorClass) UnemphasizedSelectedTextBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("unemphasizedSelectedTextBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: unemphasizedSelectedTextBackgroundColor */

// The color to use for selected text in an unemphasized context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedTextColor
func (cc _ColorClass) UnemphasizedSelectedTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("unemphasizedSelectedTextColor"))
	return rv
}/* debug [class_properties_class/property]: unemphasizedSelectedTextColor */

// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/white
func (cc _ColorClass) WhiteColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("whiteColor"))
	return rv
}/* debug [class_properties_class/property]: whiteColor */

// The color to use for the window background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/windowBackgroundColor
func (cc _ColorClass) WindowBackgroundColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("windowBackgroundColor"))
	return rv
}/* debug [class_properties_class/property]: windowBackgroundColor */

// The system color used for window frames, except for their text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/windowFrameColor
func (cc _ColorClass) WindowFrameColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("windowFrameColor"))
	return rv
}/* debug [class_properties_class/property]: windowFrameColor */

// The color to use for text in a window’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/windowFrameTextColor
func (cc _ColorClass) WindowFrameTextColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("windowFrameTextColor"))
	return rv
}/* debug [class_properties_class/property]: windowFrameTextColor */

// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellow
func (cc _ColorClass) YellowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("yellowColor"))
	return rv
}/* debug [class_properties_class/property]: yellowColor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Color */

// Sets the fill color of subsequent drawing to the color object’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/setFill()
func (c_ Color) SetFill() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFill"))
}/* debug [instance_methods/method]: SetFill */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Color */

// The system color used for the face of a selected control in a list or table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlColor
func (c_ Color) AlternateSelectedControlColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("alternateSelectedControlColor"))
	return rv
}/* debug [instance_properties/getter]: alternateSelectedControlColor */


// The color to use for text in a selected control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlTextColor
func (c_ Color) AlternateSelectedControlTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("alternateSelectedControlTextColor"))
	return rv
}/* debug [instance_properties/getter]: alternateSelectedControlTextColor */


// The colors to use for alternating content, typically found in table views and collection views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/alternatingContentBackgroundColors
func (c_ Color) AlternatingContentBackgroundColors() []Color {
	rv := objc.Send[[]Color](c_.ID, objc.Sel("alternatingContentBackgroundColors"))
	return rv
}/* debug [instance_properties/getter]: alternatingContentBackgroundColors */


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/black
func (c_ Color) BlackColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blackColor"))
	return rv
}/* debug [instance_properties/getter]: blackColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/blue
func (c_ Color) BlueColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blueColor"))
	return rv
}/* debug [instance_properties/getter]: blueColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/brown
func (c_ Color) BrownColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("brownColor"))
	return rv
}/* debug [instance_properties/getter]: brownColor */


// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/clear
func (c_ Color) ClearColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("clearColor"))
	return rv
}/* debug [instance_properties/getter]: clearColor */


// The name of the color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/colorSpaceName
func (c_ Color) ColorSpaceName() ColorSpaceName /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("colorSpaceName"))
	return rv
}/* debug [instance_properties/getter]: colorSpaceName */


// The user’s current accent color preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlAccentColor
func (c_ Color) ControlAccentColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlAccentColor"))
	return rv
}/* debug [instance_properties/getter]: controlAccentColor */


// An array containing the system specified background colors for alternating rows in tables and lists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlAlternatingRowBackgroundColors
func (c_ Color) ControlAlternatingRowBackgroundColors() []Color {
	rv := objc.Send[[]Color](c_.ID, objc.Sel("controlAlternatingRowBackgroundColors"))
	return rv
}/* debug [instance_properties/getter]: controlAlternatingRowBackgroundColors */


// The color to use for the background of large controls, such as scroll views or table views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlBackgroundColor
func (c_ Color) ControlBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: controlBackgroundColor */


// The color to use for the flat surfaces of a control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlColor
func (c_ Color) ControlColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlColor"))
	return rv
}/* debug [instance_properties/getter]: controlColor */


// The system color used for the dark edge of the shadow dropped from controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlDarkShadowColor
func (c_ Color) ControlDarkShadowColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlDarkShadowColor"))
	return rv
}/* debug [instance_properties/getter]: controlDarkShadowColor */


// The system color used for the highlighted bezels of controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlHighlightColor
func (c_ Color) ControlHighlightColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlHighlightColor"))
	return rv
}/* debug [instance_properties/getter]: controlHighlightColor */


// The system color used for light highlights in controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlLightHighlightColor
func (c_ Color) ControlLightHighlightColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlLightHighlightColor"))
	return rv
}/* debug [instance_properties/getter]: controlLightHighlightColor */


// The system color used for the shadows dropped from controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlShadowColor
func (c_ Color) ControlShadowColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlShadowColor"))
	return rv
}/* debug [instance_properties/getter]: controlShadowColor */


// The color to use for text on enabled controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/controlTextColor
func (c_ Color) ControlTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("controlTextColor"))
	return rv
}/* debug [instance_properties/getter]: controlTextColor */


// The current system control tint color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/currentControlTint
func (c_ Color) CurrentControlTint() ControlTint {
	rv := objc.Send[ControlTint](c_.ID, objc.Sel("currentControlTint"))
	return rv
}/* debug [instance_properties/getter]: currentControlTint */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/cyan
func (c_ Color) CyanColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("cyanColor"))
	return rv
}/* debug [instance_properties/getter]: cyanColor */


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/darkGray
func (c_ Color) DarkGrayColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("darkGrayColor"))
	return rv
}/* debug [instance_properties/getter]: darkGrayColor */


// The color to use for text on disabled controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/disabledControlTextColor
func (c_ Color) DisabledControlTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("disabledControlTextColor"))
	return rv
}/* debug [instance_properties/getter]: disabledControlTextColor */


// The highlight color to use for the bubble that shows inline search result values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/findHighlightColor
func (c_ Color) FindHighlightColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("findHighlightColor"))
	return rv
}/* debug [instance_properties/getter]: findHighlightColor */


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/gray
func (c_ Color) GrayColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("grayColor"))
	return rv
}/* debug [instance_properties/getter]: grayColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/green
func (c_ Color) GreenColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("greenColor"))
	return rv
}/* debug [instance_properties/getter]: greenColor */


// The color to use for the optional gridlines, such as those in a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/gridColor
func (c_ Color) GridColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("gridColor"))
	return rv
}/* debug [instance_properties/getter]: gridColor */


// The system color used as the background color for header cells in table views and outline views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/headerColor
func (c_ Color) HeaderColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("headerColor"))
	return rv
}/* debug [instance_properties/getter]: headerColor */


// The color to use for text in header cells in table views and outline views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/headerTextColor
func (c_ Color) HeaderTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("headerTextColor"))
	return rv
}/* debug [instance_properties/getter]: headerTextColor */


// The color to use as a virtual light source on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/highlightColor
func (c_ Color) HighlightColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("highlightColor"))
	return rv
}/* debug [instance_properties/getter]: highlightColor */


// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (c_ Color) IgnoresAlpha() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ignoresAlpha"))
	return rv
}/* debug [instance_properties/getter]: ignoresAlpha */


// A Boolean value that indicates whether the app supports alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (c_ Color) SetIgnoresAlpha(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIgnoresAlpha:"), value)
}/* debug [instance_properties/setter]: ignoresAlpha */


// The color to use for the keyboard focus ring around controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/keyboardFocusIndicatorColor
func (c_ Color) KeyboardFocusIndicatorColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("keyboardFocusIndicatorColor"))
	return rv
}/* debug [instance_properties/getter]: keyboardFocusIndicatorColor */


// The system color used for the flat surface of a slider knob that hasn’t been selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/knobColor
func (c_ Color) KnobColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("knobColor"))
	return rv
}/* debug [instance_properties/getter]: knobColor */


// The primary color to use for text labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
func (c_ Color) LabelColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("labelColor"))
	return rv
}/* debug [instance_properties/getter]: labelColor */


// Returns a color object whose grayscale value is and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/lightGray
func (c_ Color) LightGrayColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("lightGrayColor"))
	return rv
}/* debug [instance_properties/getter]: lightGrayColor */


// The color to use for links.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/linkColor
func (c_ Color) LinkColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("linkColor"))
	return rv
}/* debug [instance_properties/getter]: linkColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/magenta
func (c_ Color) MagentaColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("magentaColor"))
	return rv
}/* debug [instance_properties/getter]: magentaColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/orange
func (c_ Color) OrangeColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("orangeColor"))
	return rv
}/* debug [instance_properties/getter]: orangeColor */


// The pattern image used to paint the target area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/patternImage
func (c_ Color) PatternImage() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("patternImage"))
	return rv
}/* debug [instance_properties/getter]: patternImage */


// The color to use for placeholder text in controls or text views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/placeholderTextColor
func (c_ Color) PlaceholderTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("placeholderTextColor"))
	return rv
}/* debug [instance_properties/getter]: placeholderTextColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/purple
func (c_ Color) PurpleColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("purpleColor"))
	return rv
}/* debug [instance_properties/getter]: purpleColor */


// The quaternary color to use for text labels and separators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
func (c_ Color) QuaternaryLabelColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("quaternaryLabelColor"))
	return rv
}/* debug [instance_properties/getter]: quaternaryLabelColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quaternarySystemFill
func (c_ Color) QuaternarySystemFillColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("quaternarySystemFillColor"))
	return rv
}/* debug [instance_properties/getter]: quaternarySystemFillColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quinaryLabel
func (c_ Color) QuinaryLabelColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("quinaryLabelColor"))
	return rv
}/* debug [instance_properties/getter]: quinaryLabelColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/quinarySystemFill
func (c_ Color) QuinarySystemFillColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("quinarySystemFillColor"))
	return rv
}/* debug [instance_properties/getter]: quinarySystemFillColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/red
func (c_ Color) RedColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("redColor"))
	return rv
}/* debug [instance_properties/getter]: redColor */


// The system color used for scroll “bars”—that is, for the groove in which a scroller’s knob moves
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/scrollBarColor
func (c_ Color) ScrollBarColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("scrollBarColor"))
	return rv
}/* debug [instance_properties/getter]: scrollBarColor */


// The patterned color to use for the background of a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/scrubberTexturedBackground
func (c_ Color) ScrubberTexturedBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("scrubberTexturedBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: scrubberTexturedBackgroundColor */


// The secondary color to use for text labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
func (c_ Color) SecondaryLabelColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("secondaryLabelColor"))
	return rv
}/* debug [instance_properties/getter]: secondaryLabelColor */


// The color used for selected controls in non-key views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondarySelectedControlColor
func (c_ Color) SecondarySelectedControlColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("secondarySelectedControlColor"))
	return rv
}/* debug [instance_properties/getter]: secondarySelectedControlColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/secondarySystemFill
func (c_ Color) SecondarySystemFillColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("secondarySystemFillColor"))
	return rv
}/* debug [instance_properties/getter]: secondarySystemFillColor */


// The color to use for the background of selected and emphasized content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedContentBackgroundColor
func (c_ Color) SelectedContentBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedContentBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: selectedContentBackgroundColor */


// The color to use for the face of a selected control—that is, a control that has been clicked or is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlColor
func (c_ Color) SelectedControlColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedControlColor"))
	return rv
}/* debug [instance_properties/getter]: selectedControlColor */


// The color to use for text in a selected control—that is, a control being clicked or dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlTextColor
func (c_ Color) SelectedControlTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedControlTextColor"))
	return rv
}/* debug [instance_properties/getter]: selectedControlTextColor */


// The system color used for the slider knob when it is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedKnobColor
func (c_ Color) SelectedKnobColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedKnobColor"))
	return rv
}/* debug [instance_properties/getter]: selectedKnobColor */


// The color to use for the face of selected menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedMenuItemColor
func (c_ Color) SelectedMenuItemColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedMenuItemColor"))
	return rv
}/* debug [instance_properties/getter]: selectedMenuItemColor */


// The color to use for the text in menu items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedMenuItemTextColor
func (c_ Color) SelectedMenuItemTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedMenuItemTextColor"))
	return rv
}/* debug [instance_properties/getter]: selectedMenuItemTextColor */


// The color to use for the background of selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedTextBackgroundColor
func (c_ Color) SelectedTextBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedTextBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: selectedTextBackgroundColor */


// The color to use for selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/selectedTextColor
func (c_ Color) SelectedTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("selectedTextColor"))
	return rv
}/* debug [instance_properties/getter]: selectedTextColor */


// The color to use for separators between different sections of content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/separatorColor
func (c_ Color) SeparatorColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("separatorColor"))
	return rv
}/* debug [instance_properties/getter]: separatorColor */


// The color to use for virtual shadows cast by raised objects on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/shadowColor
func (c_ Color) ShadowColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("shadowColor"))
	return rv
}/* debug [instance_properties/getter]: shadowColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/systemFill
func (c_ Color) SystemFillColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("systemFillColor"))
	return rv
}/* debug [instance_properties/getter]: systemFillColor */


// The tertiary color to use for text labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
func (c_ Color) TertiaryLabelColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("tertiaryLabelColor"))
	return rv
}/* debug [instance_properties/getter]: tertiaryLabelColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/tertiarySystemFill
func (c_ Color) TertiarySystemFillColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("tertiarySystemFillColor"))
	return rv
}/* debug [instance_properties/getter]: tertiarySystemFillColor */


// The color to use for the background area behind text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textBackgroundColor
func (c_ Color) TextBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("textBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: textBackgroundColor */


// The color to use for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/textColor
func (c_ Color) TextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("textColor"))
	return rv
}/* debug [instance_properties/getter]: textColor */


// The color to use in the area beneath your window’s views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/underPageBackgroundColor
func (c_ Color) UnderPageBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("underPageBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: underPageBackgroundColor */


// The color to use for selected and unemphasized content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedContentBackgroundColor
func (c_ Color) UnemphasizedSelectedContentBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("unemphasizedSelectedContentBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: unemphasizedSelectedContentBackgroundColor */


// The color to use for the text background in an unemphasized context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedTextBackgroundColor
func (c_ Color) UnemphasizedSelectedTextBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("unemphasizedSelectedTextBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: unemphasizedSelectedTextBackgroundColor */


// The color to use for selected text in an unemphasized context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedTextColor
func (c_ Color) UnemphasizedSelectedTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("unemphasizedSelectedTextColor"))
	return rv
}/* debug [instance_properties/getter]: unemphasizedSelectedTextColor */


// Returns a color object whose grayscale and alpha values are both .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/white
func (c_ Color) WhiteColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("whiteColor"))
	return rv
}/* debug [instance_properties/getter]: whiteColor */


// The color to use for the window background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/windowBackgroundColor
func (c_ Color) WindowBackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("windowBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: windowBackgroundColor */


// The system color used for window frames, except for their text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/windowFrameColor
func (c_ Color) WindowFrameColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("windowFrameColor"))
	return rv
}/* debug [instance_properties/getter]: windowFrameColor */


// The color to use for text in a window’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/windowFrameTextColor
func (c_ Color) WindowFrameTextColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("windowFrameTextColor"))
	return rv
}/* debug [instance_properties/getter]: windowFrameTextColor */


// Returns a color object whose RGB value is , , and whose alpha value is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/yellow
func (c_ Color) YellowColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("yellowColor"))
	return rv
}/* debug [instance_properties/getter]: yellowColor */


// The alpha (opacity) component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/alphacomponent
func (c_ Color) AlphaComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alphaComponent"))
	return rv
}/* debug [instance_properties/getter]: alphaComponent */


// The alpha (opacity) component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/alphacomponent
func (c_ Color) SetAlphaComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlphaComponent:"), value)
}/* debug [instance_properties/setter]: alphaComponent */


// The black component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/blackcomponent
func (c_ Color) BlackComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blackComponent"))
	return rv
}/* debug [instance_properties/getter]: blackComponent */


// The black component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/blackcomponent
func (c_ Color) SetBlackComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBlackComponent:"), value)
}/* debug [instance_properties/setter]: blackComponent */


// The blue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/bluecomponent
func (c_ Color) BlueComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blueComponent"))
	return rv
}/* debug [instance_properties/getter]: blueComponent */


// The blue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/bluecomponent
func (c_ Color) SetBlueComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBlueComponent:"), value)
}/* debug [instance_properties/setter]: blueComponent */


// The brightness component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/brightnesscomponent
func (c_ Color) BrightnessComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("brightnessComponent"))
	return rv
}/* debug [instance_properties/getter]: brightnessComponent */


// The brightness component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/brightnesscomponent
func (c_ Color) SetBrightnessComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBrightnessComponent:"), value)
}/* debug [instance_properties/setter]: brightnessComponent */


// The catalog containing the color’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/catalognamecomponent
func (c_ Color) CatalogNameComponent() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("catalogNameComponent"))
	return rv
}/* debug [instance_properties/getter]: catalogNameComponent */


// The catalog containing the color’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/catalognamecomponent
func (c_ Color) SetCatalogNameComponent(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCatalogNameComponent:"), value)
}/* debug [instance_properties/setter]: catalogNameComponent */


// The Core Graphics color object corresponding to the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/cgcolor
func (c_ Color) CgColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("cgColor"))
	return rv
}/* debug [instance_properties/getter]: cgColor */


// The Core Graphics color object corresponding to the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/cgcolor
func (c_ Color) SetCgColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCgColor:"), value)
}/* debug [instance_properties/setter]: cgColor */


// The name of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/colornamecomponent
func (c_ Color) ColorNameComponent() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("colorNameComponent"))
	return rv
}/* debug [instance_properties/getter]: colorNameComponent */


// The name of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/colornamecomponent
func (c_ Color) SetColorNameComponent(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorNameComponent:"), value)
}/* debug [instance_properties/setter]: colorNameComponent */


// The color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/colorspace
func (c_ Color) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// The color space associated with the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/colorspace
func (c_ Color) SetColorSpace(value IColorSpace) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), value)
}/* debug [instance_properties/setter]: colorSpace */


// The cyan component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/cyancomponent
func (c_ Color) CyanComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("cyanComponent"))
	return rv
}/* debug [instance_properties/getter]: cyanComponent */


// The cyan component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/cyancomponent
func (c_ Color) SetCyanComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCyanComponent:"), value)
}/* debug [instance_properties/setter]: cyanComponent */


// The green component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/greencomponent
func (c_ Color) GreenComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("greenComponent"))
	return rv
}/* debug [instance_properties/getter]: greenComponent */


// The green component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/greencomponent
func (c_ Color) SetGreenComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGreenComponent:"), value)
}/* debug [instance_properties/setter]: greenComponent */


// The hue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/huecomponent
func (c_ Color) HueComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("hueComponent"))
	return rv
}/* debug [instance_properties/getter]: hueComponent */


// The hue component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/huecomponent
func (c_ Color) SetHueComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHueComponent:"), value)
}/* debug [instance_properties/setter]: hueComponent */


// For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/linearexposure
func (c_ Color) LinearExposure() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("linearExposure"))
	return rv
}/* debug [instance_properties/getter]: linearExposure */


// For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/linearexposure
func (c_ Color) SetLinearExposure(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLinearExposure:"), value)
}/* debug [instance_properties/setter]: linearExposure */


// The localized version of the catalog name containing the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/localizedcatalognamecomponent
func (c_ Color) LocalizedCatalogNameComponent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedCatalogNameComponent"))
	return rv
}/* debug [instance_properties/getter]: localizedCatalogNameComponent */


// The localized version of the catalog name containing the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/localizedcatalognamecomponent
func (c_ Color) SetLocalizedCatalogNameComponent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedCatalogNameComponent:"), value)
}/* debug [instance_properties/setter]: localizedCatalogNameComponent */


// The localized version of the color name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/localizedcolornamecomponent
func (c_ Color) LocalizedColorNameComponent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedColorNameComponent"))
	return rv
}/* debug [instance_properties/getter]: localizedColorNameComponent */


// The localized version of the color name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/localizedcolornamecomponent
func (c_ Color) SetLocalizedColorNameComponent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedColorNameComponent:"), value)
}/* debug [instance_properties/setter]: localizedColorNameComponent */


// The magenta component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/magentacomponent
func (c_ Color) MagentaComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("magentaComponent"))
	return rv
}/* debug [instance_properties/getter]: magentaComponent */


// The magenta component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/magentacomponent
func (c_ Color) SetMagentaComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMagentaComponent:"), value)
}/* debug [instance_properties/setter]: magentaComponent */


// The number of components in the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/numberofcomponents
func (c_ Color) NumberOfComponents() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfComponents"))
	return rv
}/* debug [instance_properties/getter]: numberOfComponents */


// The number of components in the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/numberofcomponents
func (c_ Color) SetNumberOfComponents(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfComponents:"), value)
}/* debug [instance_properties/setter]: numberOfComponents */


// The red component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/redcomponent
func (c_ Color) RedComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("redComponent"))
	return rv
}/* debug [instance_properties/getter]: redComponent */


// The red component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/redcomponent
func (c_ Color) SetRedComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedComponent:"), value)
}/* debug [instance_properties/setter]: redComponent */


// The saturation component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/saturationcomponent
func (c_ Color) SaturationComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("saturationComponent"))
	return rv
}/* debug [instance_properties/getter]: saturationComponent */


// The saturation component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/saturationcomponent
func (c_ Color) SetSaturationComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSaturationComponent:"), value)
}/* debug [instance_properties/setter]: saturationComponent */


// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/standarddynamicrange
func (c_ Color) StandardDynamicRange() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("standardDynamicRange"))
	return rv
}/* debug [instance_properties/getter]: standardDynamicRange */


// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/standarddynamicrange
func (c_ Color) SetStandardDynamicRange(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandardDynamicRange:"), value)
}/* debug [instance_properties/setter]: standardDynamicRange */


// The type of the color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/type
func (c_ Color) Type() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The type of the color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/type
func (c_ Color) SetType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// The white component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/whitecomponent
func (c_ Color) WhiteComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("whiteComponent"))
	return rv
}/* debug [instance_properties/getter]: whiteComponent */


// The white component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/whitecomponent
func (c_ Color) SetWhiteComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteComponent:"), value)
}/* debug [instance_properties/setter]: whiteComponent */


// The yellow component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/yellowcomponent
func (c_ Color) YellowComponent() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("yellowComponent"))
	return rv
}/* debug [instance_properties/getter]: yellowComponent */


// The yellow component value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolor/yellowcomponent
func (c_ Color) SetYellowComponent(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setYellowComponent:"), value)
}/* debug [instance_properties/setter]: yellowComponent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSColor */


