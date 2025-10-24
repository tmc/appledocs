// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIColor */


/* debug [class_header]: Header for CIColor */
// The class instance for the [Color] class.
var (
	ColorClass     _ColorClass
	ColorClassOnce sync.Once
)

func getColorClass() _ColorClass {
	ColorClassOnce.Do(func() {
		ColorClass = _ColorClass{objc.GetClass("CIColor")}
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
	Alpha() float64
	Blue() float64
	ColorSpace() ColorSpaceRef /* not a class type */
	Components() corefoundation.CGFloat
	Green() float64
	NumberOfComponents() uintptr /* not a class type */
	Red() float64
	StringRepresentation() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Color */
	// methods:
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
// The Core Image class that defines a color object.
//
// Use instances in conjunction with other Core Image classes, such as and . Many of the built-in Core Image filters have one or more inputs that you can set to affect the filter’s behavior.


// The Core Image class that defines a color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor
type Color struct {
	objectivec.Object
}

// ColorFrom constructs a [Color] from an unsafe.Pointer.
//
// The Core Image class that defines a color object.
func ColorFrom(ptr unsafe.Pointer) Color {
	return Color{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Color */

// Create a Core Image color object with a Core Graphics color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(cgColor:)
func NewColorWithCGColor(color ColorRef /* not a class type */) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithCGColor:"), color)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithCGColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(color:)
func NewColorWithColor(color IColor) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithColor:"), color)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithColor */


// Initialize a Core Image color object in the sRGB color space with the specified red, green, and blue component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/initWithRed:green:blue:
func NewColorWithRedGreenBlue(red float64, green float64, blue float64) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithRedGreenBlue */


// Initialize a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:)
func NewColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:alpha:"), red, green, blue, alpha)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithRedGreenBlueAlpha */


// Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:colorSpace:)
func NewColorWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace ColorSpaceRef /* not a class type */) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithRedGreenBlueAlphaColorSpace */


// Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:colorSpace:)
func NewColorWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace ColorSpaceRef /* not a class type */) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithRedGreenBlueColorSpace */


// Create a Core Image color object in the sRGB color space using a string containing the RGBA color component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(string:)
func NewColorWithString(representation objc.IObject /* cross-framework: NSString */) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithString:"), representation)
	return rv
}/* debug [class_init_methods/constructor]: NewColorWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Color */

// Create a Core Image color object with a Core Graphics color object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithCGColor:
func (cc _ColorClass) ColorWithCGColor(color ColorRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorWithCGColor:"), color)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithCGColor) */


// Create a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:alpha:
func (cc _ColorClass) ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithRedGreenBlueAlpha) */


// Create a Core Image color object with the specified red, green, blue, and alpha component values as measured in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:alpha:colorSpace:
func (cc _ColorClass) ColorWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace ColorSpaceRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithRedGreenBlueAlphaColorSpace) */


// Create a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:colorSpace:
func (cc _ColorClass) ColorWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace ColorSpaceRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithRedGreenBlueColorSpace) */


// Create a Core Image color object in the sRGB color space with the specified red, green, and blue component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:)
func (cc _ColorClass) ColorWithRedGreenBlue(red float64, green float64, blue float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:"), red, green, blue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithRedGreenBlue) */


// Create a Core Image color object in the sRGB color space using a string containing the RGBA color component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(string:)
func (cc _ColorClass) ColorWithString(representation objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("colorWithString:"), representation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ColorWithString) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Color */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/black
func (cc _ColorClass) BlackColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("blackColor"))
	return rv
}/* debug [class_properties_class/property]: blackColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.type.property
func (cc _ColorClass) BlueColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("blueColor"))
	return rv
}/* debug [class_properties_class/property]: blueColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/clear
func (cc _ColorClass) ClearColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("clearColor"))
	return rv
}/* debug [class_properties_class/property]: clearColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/cyan
func (cc _ColorClass) CyanColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("cyanColor"))
	return rv
}/* debug [class_properties_class/property]: cyanColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/gray
func (cc _ColorClass) GrayColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("grayColor"))
	return rv
}/* debug [class_properties_class/property]: grayColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.type.property
func (cc _ColorClass) GreenColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("greenColor"))
	return rv
}/* debug [class_properties_class/property]: greenColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/magenta
func (cc _ColorClass) MagentaColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("magentaColor"))
	return rv
}/* debug [class_properties_class/property]: magentaColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.type.property
func (cc _ColorClass) RedColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("redColor"))
	return rv
}/* debug [class_properties_class/property]: redColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/white
func (cc _ColorClass) WhiteColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("whiteColor"))
	return rv
}/* debug [class_properties_class/property]: whiteColor */

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/yellow
func (cc _ColorClass) YellowColor() Color {
	rv := objc.Send[Color](objc.ID(cc.class), objc.Sel("yellowColor"))
	return rv
}/* debug [class_properties_class/property]: yellowColor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Color */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Color */

// Returns the alpha value of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/alpha
func (c_ Color) Alpha() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/black
func (c_ Color) BlackColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blackColor"))
	return rv
}/* debug [instance_properties/getter]: blackColor */


// Returns the unpremultiplied blue component of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.property
func (c_ Color) Blue() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blue"))
	return rv
}/* debug [instance_properties/getter]: blue */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.type.property
func (c_ Color) BlueColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("blueColor"))
	return rv
}/* debug [instance_properties/getter]: blueColor */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/clear
func (c_ Color) ClearColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("clearColor"))
	return rv
}/* debug [instance_properties/getter]: clearColor */


// Returns the associated with the color
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorSpace
func (c_ Color) ColorSpace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](c_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// Return a pointer to an array of values including alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/components
func (c_ Color) Components() corefoundation.CGFloat {
	rv := objc.Send[corefoundation.CGFloat](c_.ID, objc.Sel("components"))
	return rv
}/* debug [instance_properties/getter]: components */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/cyan
func (c_ Color) CyanColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("cyanColor"))
	return rv
}/* debug [instance_properties/getter]: cyanColor */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/gray
func (c_ Color) GrayColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("grayColor"))
	return rv
}/* debug [instance_properties/getter]: grayColor */


// Returns the unpremultiplied green component of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.property
func (c_ Color) Green() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("green"))
	return rv
}/* debug [instance_properties/getter]: green */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.type.property
func (c_ Color) GreenColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("greenColor"))
	return rv
}/* debug [instance_properties/getter]: greenColor */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/magenta
func (c_ Color) MagentaColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("magentaColor"))
	return rv
}/* debug [instance_properties/getter]: magentaColor */


// Returns the color components of the color including alpha.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/numberOfComponents
func (c_ Color) NumberOfComponents() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](c_.ID, objc.Sel("numberOfComponents"))
	return rv
}/* debug [instance_properties/getter]: numberOfComponents */


// Returns the unpremultiplied red component of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.property
func (c_ Color) Red() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("red"))
	return rv
}/* debug [instance_properties/getter]: red */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.type.property
func (c_ Color) RedColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("redColor"))
	return rv
}/* debug [instance_properties/getter]: redColor */


// Returns a formatted string with the unpremultiplied color and alpha components of the color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/stringRepresentation
func (c_ Color) StringRepresentation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("stringRepresentation"))
	return rv
}/* debug [instance_properties/getter]: stringRepresentation */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/white
func (c_ Color) WhiteColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("whiteColor"))
	return rv
}/* debug [instance_properties/getter]: whiteColor */


// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/yellow
func (c_ Color) YellowColor() ICIColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("yellowColor"))
	return rv
}/* debug [instance_properties/getter]: yellowColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIColor */


