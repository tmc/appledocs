// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// The class instance for the [MKOverlayPathRenderer] class.
var (
	MKOverlayPathRendererClass     _MKOverlayPathRendererClass
	MKOverlayPathRendererClassOnce sync.Once
)

func getMKOverlayPathRendererClass() _MKOverlayPathRendererClass {
	MKOverlayPathRendererClassOnce.Do(func() {
		MKOverlayPathRendererClass = _MKOverlayPathRendererClass{objc.GetClass("MKOverlayPathRenderer")}
	})
	return MKOverlayPathRendererClass
}

type _MKOverlayPathRendererClass struct {
	class objc.Class
}

// An interface definition for the [MKOverlayPathRenderer] class.
type IMKOverlayPathRenderer interface {
	IMKOverlayRenderer
}

// The visual representation of a path-based overlay.
//
// Use this renderer when a object defines your overlay’s shape. By default, this renderer fills the overlay’s shape and represents the strokes of the path using its current attributes. You can use this class as-is or subclass it to define additional drawing behaviors. If you subclass it, override the method and use that method to build the appropriate path object. To change the path, invalidate it and recreate the path using the new data your subclass obtains.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer
type MKOverlayPathRenderer struct {
	MKOverlayRenderer
}

// MKOverlayPathRendererFrom constructs a [MKOverlayPathRenderer] from an unsafe.Pointer.
//
// The visual representation of a path-based overlay.
func MKOverlayPathRendererFrom(ptr unsafe.Pointer) MKOverlayPathRenderer {
	return MKOverlayPathRenderer{
		MKOverlayRenderer: MKOverlayRendererFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKOverlayPathRendererClass) Alloc() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKOverlayPathRendererClass) New() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKOverlayPathRenderer) Init() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKOverlayPathRenderer) Autorelease() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKOverlayPathRenderer creates a new MKOverlayPathRenderer instance.
func NewMKOverlayPathRenderer() MKOverlayPathRenderer {
	return getMKOverlayPathRendererClass().New()
}


// The fill color to use for the path.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/fillcolor
func (m_ MKOverlayPathRenderer) FillColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("fillColor"))
	return rv
}


// SetFillColor sets the value of the fillColor property.
// The fill color to use for the path.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/fillcolor
func (m_ MKOverlayPathRenderer) SetFillColor(value appkit.IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFillColor:"), value)
}

// The line cap style to apply to the open ends of the path.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linecap
func (m_ MKOverlayPathRenderer) LineCap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lineCap"))
	return rv
}


// SetLineCap sets the value of the lineCap property.
// The line cap style to apply to the open ends of the path.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linecap
func (m_ MKOverlayPathRenderer) SetLineCap(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineCap:"), value)
}

// An array of numbers specifying the dash pattern to use for the path.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linedashpattern
func (m_ MKOverlayPathRenderer) LineDashPattern() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lineDashPattern"))
	return rv
}


// SetLineDashPattern sets the value of the lineDashPattern property.
// An array of numbers specifying the dash pattern to use for the path.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linedashpattern
func (m_ MKOverlayPathRenderer) SetLineDashPattern(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineDashPattern:"), value)
}

// The offset (in points) at which to start drawing the dash pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linedashphase
func (m_ MKOverlayPathRenderer) LineDashPhase() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineDashPhase"))
	return rv
}


// SetLineDashPhase sets the value of the lineDashPhase property.
// The offset (in points) at which to start drawing the dash pattern.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linedashphase
func (m_ MKOverlayPathRenderer) SetLineDashPhase(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineDashPhase:"), value)
}

// The line join style to apply to the corners of the path.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linejoin
func (m_ MKOverlayPathRenderer) LineJoin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lineJoin"))
	return rv
}


// SetLineJoin sets the value of the lineJoin property.
// The line join style to apply to the corners of the path.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linejoin
func (m_ MKOverlayPathRenderer) SetLineJoin(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineJoin:"), value)
}

// The stroke width to use for the path.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linewidth
func (m_ MKOverlayPathRenderer) LineWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// The stroke width to use for the path.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/linewidth
func (m_ MKOverlayPathRenderer) SetLineWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineWidth:"), value)
}

// The limiting value that helps avoid spikes at junctions between connected line segments.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/miterlimit
func (m_ MKOverlayPathRenderer) MiterLimit() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("miterLimit"))
	return rv
}


// SetMiterLimit sets the value of the miterLimit property.
// The limiting value that helps avoid spikes at junctions between connected line segments.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/miterlimit
func (m_ MKOverlayPathRenderer) SetMiterLimit(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMiterLimit:"), value)
}

// The path representing the overlay’s shape.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/path
func (m_ MKOverlayPathRenderer) Path() gameplaykit.Path {
	rv := objc.Send[gameplaykit.Path](m_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The path representing the overlay’s shape.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/path
func (m_ MKOverlayPathRenderer) SetPath(value gameplaykit.IPath) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}

// A Boolean value that determines whether the overlay path renderer renders the overlay as a bitmap before compositing.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/shouldrasterize
func (m_ MKOverlayPathRenderer) ShouldRasterize() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldRasterize"))
	return rv
}


// SetShouldRasterize sets the value of the shouldRasterize property.
// A Boolean value that determines whether the overlay path renderer renders the overlay as a bitmap before compositing.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/shouldrasterize
func (m_ MKOverlayPathRenderer) SetShouldRasterize(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRasterize:"), value)
}

// The stroke color to use for the path.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/strokecolor
func (m_ MKOverlayPathRenderer) StrokeColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("strokeColor"))
	return rv
}


// SetStrokeColor sets the value of the strokeColor property.
// The stroke color to use for the path.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkoverlaypathrenderer/strokecolor
func (m_ MKOverlayPathRenderer) SetStrokeColor(value appkit.IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeColor:"), value)
}



