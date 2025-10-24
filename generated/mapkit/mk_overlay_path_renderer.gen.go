// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class MKOverlayPathRenderer */


/* debug [class_header]: Header for MKOverlayPathRenderer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKOverlayPathRenderer */
// An interface definition for the [MKOverlayPathRenderer] class.
type IMKOverlayPathRenderer interface {
	IMKOverlayRenderer
	
/* debug [class_interface_properties]: Properties for MKOverlayPathRenderer */
	// properties:
	FillColor() appkit.Color
	SetFillColor(value appkit.Color)
	LineCap() LineCap /* not a class type */
	SetLineCap(value LineCap /* not a class type */)
	LineDashPattern() []foundation.Number
	SetLineDashPattern(value []foundation.Number)
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	LineJoin() LineJoin /* not a class type */
	SetLineJoin(value LineJoin /* not a class type */)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	Path() PathRef /* not a class type */
	SetPath(value PathRef /* not a class type */)
	ShouldRasterize() bool
	SetShouldRasterize(value bool)
	StrokeColor() appkit.Color
	SetStrokeColor(value appkit.Color)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKOverlayPathRenderer */
	// methods:
	ApplyFillPropertiesToContextAtZoomScale(context ContextRef /* not a class type */, zoomScale MKZoomScale /* typedef */)
	ApplyStrokePropertiesToContextAtZoomScale(context ContextRef /* not a class type */, zoomScale MKZoomScale /* typedef */)
	CreatePath()
	FillPathInContext(path PathRef /* not a class type */, context ContextRef /* not a class type */)
	InvalidatePath()
	StrokePathInContext(path PathRef /* not a class type */, context ContextRef /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKOverlayPathRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKOverlayPathRendererClass) Alloc() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKOverlayPathRenderer */
// The visual representation of a path-based overlay.
//
// Use this renderer when a object defines your overlay’s shape. By default, this renderer fills the overlay’s shape and represents the strokes of the path using its current attributes. You can use this class as-is or subclass it to define additional drawing behaviors. If you subclass it, override the method and use that method to build the appropriate path object. To change the path, invalidate it and recreate the path using the new data your subclass obtains.


// The visual representation of a path-based overlay.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKOverlayPathRenderer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKOverlayPathRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKOverlayPathRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKOverlayPathRenderer */

// Applies the receiver’s fill-related drawing properties to the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/applyFillProperties(to:atZoomScale:)
func (m_ MKOverlayPathRenderer) ApplyFillPropertiesToContextAtZoomScale(context ContextRef /* not a class type */, zoomScale MKZoomScale /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyFillPropertiesToContext:atZoomScale:"), context, zoomScale)
}/* debug [instance_methods/method]: ApplyFillPropertiesToContextAtZoomScale */


// Applies the renderer’s stroke-related drawing properties to the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/applyStrokeProperties(to:atZoomScale:)
func (m_ MKOverlayPathRenderer) ApplyStrokePropertiesToContextAtZoomScale(context ContextRef /* not a class type */, zoomScale MKZoomScale /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyStrokePropertiesToContext:atZoomScale:"), context, zoomScale)
}/* debug [instance_methods/method]: ApplyStrokePropertiesToContextAtZoomScale */


// Creates the path for the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/createPath()
func (m_ MKOverlayPathRenderer) CreatePath() {
	objc.Send[objc.ID](m_.ID, objc.Sel("createPath"))
}/* debug [instance_methods/method]: CreatePath */


// Fills the area that the specified path encloses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/fillPath(_:in:)
func (m_ MKOverlayPathRenderer) FillPathInContext(path PathRef /* not a class type */, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("fillPath:inContext:"), path, context)
}/* debug [instance_methods/method]: FillPathInContext */


// Updates the path associated with the overlay renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/invalidatePath()
func (m_ MKOverlayPathRenderer) InvalidatePath() {
	objc.Send[objc.ID](m_.ID, objc.Sel("invalidatePath"))
}/* debug [instance_methods/method]: InvalidatePath */


// Draws a line along the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/strokePath(_:in:)
func (m_ MKOverlayPathRenderer) StrokePathInContext(path PathRef /* not a class type */, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("strokePath:inContext:"), path, context)
}/* debug [instance_methods/method]: StrokePathInContext */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKOverlayPathRenderer */

// The fill color to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/fillColor
func (m_ MKOverlayPathRenderer) FillColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("fillColor"))
	return rv
}/* debug [instance_properties/getter]: fillColor */


// The fill color to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/fillColor
func (m_ MKOverlayPathRenderer) SetFillColor(value appkit.Color) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFillColor:"), value)
}/* debug [instance_properties/setter]: fillColor */


// The line cap style to apply to the open ends of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineCap
func (m_ MKOverlayPathRenderer) LineCap() LineCap /* not a class type */ {
	rv := objc.Send[LineCap](m_.ID, objc.Sel("lineCap"))
	return rv
}/* debug [instance_properties/getter]: lineCap */


// The line cap style to apply to the open ends of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineCap
func (m_ MKOverlayPathRenderer) SetLineCap(value LineCap /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineCap:"), value)
}/* debug [instance_properties/setter]: lineCap */


// An array of numbers specifying the dash pattern to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineDashPattern
func (m_ MKOverlayPathRenderer) LineDashPattern() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("lineDashPattern"))
	return rv
}/* debug [instance_properties/getter]: lineDashPattern */


// An array of numbers specifying the dash pattern to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineDashPattern
func (m_ MKOverlayPathRenderer) SetLineDashPattern(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineDashPattern:"), nsArray)
}/* debug [instance_properties/setter]: lineDashPattern */


// The offset (in points) at which to start drawing the dash pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineDashPhase
func (m_ MKOverlayPathRenderer) LineDashPhase() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineDashPhase"))
	return rv
}/* debug [instance_properties/getter]: lineDashPhase */


// The offset (in points) at which to start drawing the dash pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineDashPhase
func (m_ MKOverlayPathRenderer) SetLineDashPhase(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineDashPhase:"), value)
}/* debug [instance_properties/setter]: lineDashPhase */


// The line join style to apply to the corners of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineJoin
func (m_ MKOverlayPathRenderer) LineJoin() LineJoin /* not a class type */ {
	rv := objc.Send[LineJoin](m_.ID, objc.Sel("lineJoin"))
	return rv
}/* debug [instance_properties/getter]: lineJoin */


// The line join style to apply to the corners of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineJoin
func (m_ MKOverlayPathRenderer) SetLineJoin(value LineJoin /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineJoin:"), value)
}/* debug [instance_properties/setter]: lineJoin */


// The stroke width to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineWidth
func (m_ MKOverlayPathRenderer) LineWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// The stroke width to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/lineWidth
func (m_ MKOverlayPathRenderer) SetLineWidth(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// The limiting value that helps avoid spikes at junctions between connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/miterLimit
func (m_ MKOverlayPathRenderer) MiterLimit() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("miterLimit"))
	return rv
}/* debug [instance_properties/getter]: miterLimit */


// The limiting value that helps avoid spikes at junctions between connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/miterLimit
func (m_ MKOverlayPathRenderer) SetMiterLimit(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMiterLimit:"), value)
}/* debug [instance_properties/setter]: miterLimit */


// The path representing the overlay’s shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/path
func (m_ MKOverlayPathRenderer) Path() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](m_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// The path representing the overlay’s shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/path
func (m_ MKOverlayPathRenderer) SetPath(value PathRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}/* debug [instance_properties/setter]: path */


// A Boolean value that determines whether the overlay path renderer renders the overlay as a bitmap before compositing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/shouldRasterize
func (m_ MKOverlayPathRenderer) ShouldRasterize() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldRasterize"))
	return rv
}/* debug [instance_properties/getter]: shouldRasterize */


// A Boolean value that determines whether the overlay path renderer renders the overlay as a bitmap before compositing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/shouldRasterize
func (m_ MKOverlayPathRenderer) SetShouldRasterize(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRasterize:"), value)
}/* debug [instance_properties/setter]: shouldRasterize */


// The stroke color to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/strokeColor
func (m_ MKOverlayPathRenderer) StrokeColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("strokeColor"))
	return rv
}/* debug [instance_properties/getter]: strokeColor */


// The stroke color to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer/strokeColor
func (m_ MKOverlayPathRenderer) SetStrokeColor(value appkit.Color) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStrokeColor:"), value)
}/* debug [instance_properties/setter]: strokeColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKOverlayPathRenderer */



