// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Color] class.
type IColor interface {
	objectivec.IObject
}

// The Core Image class that defines a color object.
//
// Use instances in conjunction with other Core Image classes, such as and . Many of the built-in Core Image filters have one or more inputs that you can set to affect the filter’s behavior.
//
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




// Create a Core Image color object with a Core Graphics color object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(cgColor:)
func NewColorWithCGColor(color coregraphics.CGColorRef) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithCGColor:"), color)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(color:)
func NewColorWithColor(color unsafe.Pointer) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithColor:"), color)
	rv.Autorelease()
	return rv
}



// Initialize a Core Image color object in the sRGB color space with the specified red, green, and blue component values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/initWithRed:green:blue:
func NewColorWithRedGreenBlue(red float64, green float64, blue float64) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	rv.Autorelease()
	return rv
}



// Initialize a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:)
func NewColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:alpha:"), red, green, blue, alpha)
	rv.Autorelease()
	return rv
}



// Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:colorSpace:)
func NewColorWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace coregraphics.CGColorSpaceRef) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	rv.Autorelease()
	return rv
}



// Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:colorSpace:)
func NewColorWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace coregraphics.CGColorSpaceRef) Color {
	instance := getColorClass().Alloc()
	rv := objc.Send[Color](instance.ID, objc.Sel("initWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	rv.Autorelease()
	return rv
}



// Create a Core Image color object in the sRGB color space using a string containing the RGBA color component values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(string:)
func NewColorWithString(representation string) Color {
	rv := objc.Send[Color](objc.ID(getColorClass().class), objc.Sel("colorWithString:"), objc.String(representation))
	return rv
}


// Create a Core Image color object with a Core Graphics color object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithCGColor:
func (cc _ColorClass) ColorWithCGColor(color coregraphics.CGColorRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithCGColor:"), color)
	return rv
}

// Create a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:alpha:
func (cc _ColorClass) ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

// Create a Core Image color object with the specified red, green, blue, and alpha component values as measured in the specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:alpha:colorSpace:
func (cc _ColorClass) ColorWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace coregraphics.CGColorSpaceRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	return rv
}

// Create a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:colorSpace:
func (cc _ColorClass) ColorWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace coregraphics.CGColorSpaceRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	return rv
}

// Create a Core Image color object in the sRGB color space with the specified red, green, and blue component values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:)
func (cc _ColorClass) ColorWithRedGreenBlue(red float64, green float64, blue float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithRed:green:blue:"), red, green, blue)
	return rv
}

// Create a Core Image color object in the sRGB color space using a string containing the RGBA color component values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/init(string:)
func (cc _ColorClass) ColorWithString(representation string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("colorWithString:"), objc.String(representation))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/black
func (cc _ColorClass) BlackColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("blackColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.type.property
func (cc _ColorClass) BlueColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("blueColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/clear
func (cc _ColorClass) ClearColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("clearColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/cyan
func (cc _ColorClass) CyanColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("cyanColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/gray
func (cc _ColorClass) GrayColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("grayColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.type.property
func (cc _ColorClass) GreenColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("greenColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/magenta
func (cc _ColorClass) MagentaColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("magentaColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.type.property
func (cc _ColorClass) RedColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("redColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/white
func (cc _ColorClass) WhiteColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("whiteColor"))
	return rv
}
// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/yellow
func (cc _ColorClass) YellowColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("yellowColor"))
	return rv
}
// Returns the alpha value of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/alpha
func (c_ Color) Alpha() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("alpha"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/black
func (c_ Color) BlackColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("blackColor"))
	return rv
}

// Returns the unpremultiplied blue component of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.property
func (c_ Color) Blue() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("blue"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.type.property
func (c_ Color) BlueColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("blueColor"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/clear
func (c_ Color) ClearColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("clearColor"))
	return rv
}

// Returns the associated with the color
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/colorSpace
func (c_ Color) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](c_.ID, objc.Sel("colorSpace"))
	return rv
}

// Return a pointer to an array of values including alpha.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/components
func (c_ Color) Components() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("components"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/cyan
func (c_ Color) CyanColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cyanColor"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/gray
func (c_ Color) GrayColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("grayColor"))
	return rv
}

// Returns the unpremultiplied green component of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.property
func (c_ Color) Green() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("green"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.type.property
func (c_ Color) GreenColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("greenColor"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/magenta
func (c_ Color) MagentaColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("magentaColor"))
	return rv
}

// Returns the color components of the color including alpha.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/numberOfComponents
func (c_ Color) NumberOfComponents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("numberOfComponents"))
	return rv
}

// Returns the unpremultiplied red component of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.property
func (c_ Color) Red() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("red"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.type.property
func (c_ Color) RedColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("redColor"))
	return rv
}

// Returns a formatted string with the unpremultiplied color and alpha components of the color.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/stringRepresentation
func (c_ Color) StringRepresentation() string {
	rv := objc.Send[string](c_.ID, objc.Sel("stringRepresentation"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/white
func (c_ Color) WhiteColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("whiteColor"))
	return rv
}

// Returns a singleton Core Image color instance in the sRGB color space with RGB values and alpha value .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIColor/yellow
func (c_ Color) YellowColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("yellowColor"))
	return rv
}


