// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKOverlayRenderer */


/* debug [class_header]: Header for MKOverlayRenderer */
// The class instance for the [MKOverlayRenderer] class.
var (
	MKOverlayRendererClass     _MKOverlayRendererClass
	MKOverlayRendererClassOnce sync.Once
)

func getMKOverlayRendererClass() _MKOverlayRendererClass {
	MKOverlayRendererClassOnce.Do(func() {
		MKOverlayRendererClass = _MKOverlayRendererClass{objc.GetClass("MKOverlayRenderer")}
	})
	return MKOverlayRendererClass
}

type _MKOverlayRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKOverlayRenderer */
// An interface definition for the [MKOverlayRenderer] class.
type IMKOverlayRenderer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKOverlayRenderer */
	// properties:
	Alpha() float64
	SetAlpha(value float64)
	ContentScaleFactor() float64
	Overlay() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKOverlayRenderer */
	// methods:
	CanDrawMapRectZoomScale(mapRect objc.IObject /* cross-framework: MKMapRect */, zoomScale MKZoomScale /* typedef */) bool
	DrawMapRectZoomScaleInContext(mapRect objc.IObject /* cross-framework: MKMapRect */, zoomScale MKZoomScale /* typedef */, context ContextRef /* not a class type */)
	MapPointForPoint(point corefoundation.CGPoint) objc.IObject /* cross-framework: MKMapPoint */
	MapRectForRect(rect corefoundation.CGRect) objc.IObject /* cross-framework: MKMapRect */
	PointForMapPoint(mapPoint objc.IObject /* cross-framework: MKMapPoint */) corefoundation.CGPoint
	RectForMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) corefoundation.CGRect
	SetNeedsDisplay()
	SetNeedsDisplayInMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */)
	SetNeedsDisplayInMapRectZoomScale(mapRect objc.IObject /* cross-framework: MKMapRect */, zoomScale MKZoomScale /* typedef */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKOverlayRenderer */
// Alloc allocates a new instance without initialization.
func (mc _MKOverlayRendererClass) Alloc() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKOverlayRendererClass) New() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKOverlayRenderer) Init() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKOverlayRenderer) Autorelease() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKOverlayRenderer creates a new MKOverlayRenderer instance.
func NewMKOverlayRenderer() MKOverlayRenderer {
	return getMKOverlayRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKOverlayRenderer */
// The shared infrastructure for drawing overlays on the map surface.
//
// An overlay renderer draws the visual representation of an overlay object — that is, an object that conforms to the protocol. This class defines the drawing infrastructure the map view uses. Subclasses need to override the method to draw the contents of the overlay. The MapKit framework provides several concrete instances of overlay renderers. Specifically, it provides renderers for each of the concrete overlay objects. You can use one of these existing renderers or define your own subclasses if you want to draw the overlay contents differently. You can subclass to create overlays based on custom shapes, content, or drawing techniques. The only method subclasses need to override is the method. However, if your class contains content that may not be ready for drawing right away, you need to also override the method and use it to report when your class is ready and able to draw. The map view may tile large overlays and distribute the rendering of each tile to separate threads. Therefore, the implementation of your method needs to be safe to run from background threads and from multiple threads simultaneously.


// The shared infrastructure for drawing overlays on the map surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer
type MKOverlayRenderer struct {
	objectivec.Object
}

// MKOverlayRendererFrom constructs a [MKOverlayRenderer] from an unsafe.Pointer.
//
// The shared infrastructure for drawing overlays on the map surface.
func MKOverlayRendererFrom(ptr unsafe.Pointer) MKOverlayRenderer {
	return MKOverlayRenderer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKOverlayRenderer */

// Creates and returns the overlay renderer and associates it with the specified overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/init(overlay:)
func NewMKOverlayRendererWithOverlay(overlay unsafe.Pointer) MKOverlayRenderer {
	instance := getMKOverlayRendererClass().Alloc()
	rv := objc.Send[MKOverlayRenderer](instance.ID, objc.Sel("initWithOverlay:"), overlay)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKOverlayRendererWithOverlay */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKOverlayRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKOverlayRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKOverlayRenderer */

// Returns a Boolean value that indicates whether the overlay view is ready to draw its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/canDraw(_:zoomScale:)
func (m_ MKOverlayRenderer) CanDrawMapRectZoomScale(mapRect objc.IObject /* cross-framework: MKMapRect */, zoomScale MKZoomScale /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canDrawMapRect:zoomScale:"), mapRect, zoomScale)
	return rv
}/* debug [instance_methods/method]: CanDrawMapRectZoomScale */


// Draws the overlay’s contents at the specified location on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/draw(_:zoomScale:in:)
func (m_ MKOverlayRenderer) DrawMapRectZoomScaleInContext(mapRect objc.IObject /* cross-framework: MKMapRect */, zoomScale MKZoomScale /* typedef */, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawMapRect:zoomScale:inContext:"), mapRect, zoomScale, context)
}/* debug [instance_methods/method]: DrawMapRectZoomScaleInContext */


// Returns the point on the map that corresponds to the specified point in the overlay renderer’s drawing area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/mapPoint(for:)
func (m_ MKOverlayRenderer) MapPointForPoint(point corefoundation.CGPoint) objc.IObject /* cross-framework: MKMapPoint */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mapPointForPoint:"), point)
	return rv
}/* debug [instance_methods/method]: MapPointForPoint */


// Returns the rectangle on the map that corresponds to the specified rectangle in the overlay renderer’s drawing area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/mapRect(for:)
func (m_ MKOverlayRenderer) MapRectForRect(rect corefoundation.CGRect) objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mapRectForRect:"), rect)
	return rv
}/* debug [instance_methods/method]: MapRectForRect */


// Returns the point in the overlay renderer’s drawing area corresponding to the specified point on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/point(for:)
func (m_ MKOverlayRenderer) PointForMapPoint(mapPoint objc.IObject /* cross-framework: MKMapPoint */) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](m_.ID, objc.Sel("pointForMapPoint:"), mapPoint)
	return rv
}/* debug [instance_methods/method]: PointForMapPoint */


// Returns the rectangle in the overlay renderer’s drawing area corresponding to the specified rectangle on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/rect(for:)
func (m_ MKOverlayRenderer) RectForMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("rectForMapRect:"), mapRect)
	return rv
}/* debug [instance_methods/method]: RectForMapRect */


// Invalidates the entire contents of the overlay for all zoom scales.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/setNeedsDisplay()
func (m_ MKOverlayRenderer) SetNeedsDisplay() {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsDisplay"))
}/* debug [instance_methods/method]: SetNeedsDisplay */


// Invalidates the specified portion of the overlay at all zoom scales.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/setNeedsDisplay(_:)
func (m_ MKOverlayRenderer) SetNeedsDisplayInMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsDisplayInMapRect:"), mapRect)
}/* debug [instance_methods/method]: SetNeedsDisplayInMapRect */


// Invalidates the specified portion of the overlay, but only at the specified zoom scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/setNeedsDisplay(_:zoomScale:)
func (m_ MKOverlayRenderer) SetNeedsDisplayInMapRectZoomScale(mapRect objc.IObject /* cross-framework: MKMapRect */, zoomScale MKZoomScale /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeedsDisplayInMapRect:zoomScale:"), mapRect, zoomScale)
}/* debug [instance_methods/method]: SetNeedsDisplayInMapRectZoomScale */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKOverlayRenderer */

// The amount of transparency to apply to the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/alpha
func (m_ MKOverlayRenderer) Alpha() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// The amount of transparency to apply to the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/alpha
func (m_ MKOverlayRenderer) SetAlpha(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlpha:"), value)
}/* debug [instance_properties/setter]: alpha */


// The scale factor for drawing the overlay’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/contentScaleFactor
func (m_ MKOverlayRenderer) ContentScaleFactor() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("contentScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: contentScaleFactor */


// The overlay object containing the data for drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/overlay
func (m_ MKOverlayRenderer) Overlay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("overlay"))
	return rv
}/* debug [instance_properties/getter]: overlay */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKOverlayRenderer */


