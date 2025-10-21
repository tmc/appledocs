// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Gradient] class.
var (
	GradientClass     _GradientClass
	GradientClassOnce sync.Once
)

func getGradientClass() _GradientClass {
	GradientClassOnce.Do(func() {
		GradientClass = _GradientClass{objc.GetClass("NSGradient")}
	})
	return GradientClass
}

type _GradientClass struct {
	class objc.Class
}

// An interface definition for the [Gradient] class.
type IGradient interface {
	objectivec.IObject
	DrawFromPointToPointOptions(startingPoint coregraphics.CGPoint, endingPoint coregraphics.CGPoint, options unsafe.Pointer)
	DrawFromCenterRadiusToCenterRadiusOptions(startCenter coregraphics.CGPoint, startRadius float64, endCenter coregraphics.CGPoint, endRadius float64, options unsafe.Pointer)
	DrawInBezierPathAngle(path unsafe.Pointer, angle float64)
	DrawInRectRelativeCenterPosition(rect coregraphics.CGRect, relativeCenterPosition coregraphics.CGPoint)
}

// An object that can draw gradient fill colors
//
// This class provides convenience methods for drawing radial or linear (axial) gradients for rectangles and objects. It also supports primitive methods that let you customize the shape of the gradient fill. A gradient consists of two or more color changes over the range of the gradient shape. When creating a gradient object, you specify the colors and their locations relative to the start and end of the gradient. This combination of color and location is known as a . During drawing, the object uses the color stop information to compute color changes for you and passes that information to the Quartz shading functions. Because the class uses Quartz shadings, drawing is handled by computing the colors at a given point mathematically. This technique results in smooth gradients regardless of the resolution of the target device. For more information about gradients and their appearance, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient
type Gradient struct {
	objectivec.Object
}

// GradientFrom constructs a [Gradient] from an unsafe.Pointer.
//
// An object that can draw gradient fill colors
func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GradientClass) Alloc() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GradientClass) New() Gradient {
	rv := objc.Send[Gradient](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ Gradient) Init() Gradient {
	rv := objc.Send[Gradient](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ Gradient) Autorelease() Gradient {
	rv := objc.Send[Gradient](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGradient creates a new Gradient instance.
func NewGradient() Gradient {
	return getGradientClass().New()
}




// Initializes a newly allocated gradient object with the specified colors, color locations, and color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/init(colors:atLocations:colorSpace:)
func NewGradientWithColorsAtLocationsColorSpace(colorArray unsafe.Pointer, locations unsafe.Pointer, colorSpace unsafe.Pointer) Gradient {
	instance := getGradientClass().Alloc()
	rv := objc.Send[Gradient](instance.ID, objc.Sel("initWithColors:atLocations:colorSpace:"), colorArray, locations, colorSpace)
	rv.Autorelease()
	return rv
}


// Draws a linear gradient between the specified start and end points.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/draw(from:to:options:)
func (g_ Gradient) DrawFromPointToPointOptions(startingPoint coregraphics.CGPoint, endingPoint coregraphics.CGPoint, options unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("drawFromPoint:toPoint:options:"), startingPoint, endingPoint, options)
}

// Draws a radial gradient between the specified circles.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/draw(fromCenter:radius:toCenter:radius:options:)
func (g_ Gradient) DrawFromCenterRadiusToCenterRadiusOptions(startCenter coregraphics.CGPoint, startRadius float64, endCenter coregraphics.CGPoint, endRadius float64, options unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("drawFromCenter:radius:toCenter:radius:options:"), startCenter, startRadius, endCenter, endRadius, options)
}

// Fills the specified path with a linear gradient.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/draw(in:angle:)-68adz
func (g_ Gradient) DrawInBezierPathAngle(path unsafe.Pointer, angle float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("drawInBezierPath:angle:"), path, angle)
}

// Draws a radial gradient starting at the center of the specified rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/draw(in:relativeCenterPosition:)-3a83
func (g_ Gradient) DrawInRectRelativeCenterPosition(rect coregraphics.CGRect, relativeCenterPosition coregraphics.CGPoint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("drawInRect:relativeCenterPosition:"), rect, relativeCenterPosition)
}

// The number of color stops associated with the gradient.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/numberOfColorStops
func (g_ Gradient) NumberOfColorStops() int {
	rv := objc.Send[int](g_.ID, objc.Sel("numberOfColorStops"))
	return rv
}

// The color space of the colors associated with the gradient.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/colorspace
func (g_ Gradient) ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("colorSpace"))
	return rv
}


// SetColorSpace sets the value of the colorSpace property.
// The color space of the colors associated with the gradient.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/colorspace
func (g_ Gradient) SetColorSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColorSpace:"), value)
}


