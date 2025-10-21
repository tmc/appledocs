// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [MKMapView] class.
var (
	MKMapViewClass     _MKMapViewClass
	MKMapViewClassOnce sync.Once
)

func getMKMapViewClass() _MKMapViewClass {
	MKMapViewClassOnce.Do(func() {
		MKMapViewClass = _MKMapViewClass{objc.GetClass("MKMapView")}
	})
	return MKMapViewClass
}

type _MKMapViewClass struct {
	class objc.Class
}

// An interface definition for the [MKMapView] class.
type IMKMapView interface {
	appkit.IView
	AddAnnotation(annotation objc.ID)
	AddAnnotations(annotations unsafe.Pointer)
	AddOverlay(overlay objc.ID)
	AddOverlayLevel(overlay objc.ID, level unsafe.Pointer)
	AddOverlays(overlays unsafe.Pointer)
	AddOverlaysLevel(overlays unsafe.Pointer, level unsafe.Pointer)
	AnnotationsInMapRect(mapRect unsafe.Pointer) unsafe.Pointer
	ConvertPointToCoordinateFromView(point coregraphics.CGPoint, view unsafe.Pointer) unsafe.Pointer
	ConvertCoordinateToPointToView(coordinate unsafe.Pointer, view unsafe.Pointer) coregraphics.CGPoint
	ConvertRegionToRectToView(region unsafe.Pointer, view unsafe.Pointer) coregraphics.CGRect
	ConvertRectToRegionFromView(rect coregraphics.CGRect, view unsafe.Pointer) unsafe.Pointer
	DequeueReusableAnnotationViewWithIdentifier(identifier string) unsafe.Pointer
	DequeueReusableAnnotationViewWithIdentifierForAnnotation(identifier string, annotation objc.ID) unsafe.Pointer
	DeselectAnnotationAnimated(annotation objc.ID, animated bool)
	ExchangeOverlayWithOverlay(overlay1 objc.ID, overlay2 objc.ID)
	ExchangeOverlayAtIndexWithOverlayAtIndex(index1 uint, index2 uint)
	InsertOverlayAboveOverlay(overlay objc.ID, sibling objc.ID)
	InsertOverlayAtIndex(overlay objc.ID, index uint)
	InsertOverlayAtIndexLevel(overlay objc.ID, index uint, level unsafe.Pointer)
	InsertOverlayBelowOverlay(overlay objc.ID, sibling objc.ID)
	MapRectThatFits(mapRect unsafe.Pointer) unsafe.Pointer
	MapRectThatFitsEdgePadding(mapRect unsafe.Pointer, insets unsafe.Pointer) unsafe.Pointer
	OverlaysInLevel(level unsafe.Pointer) []objc.ID
	RegionThatFits(region unsafe.Pointer) unsafe.Pointer
	RegisterClassForAnnotationViewWithReuseIdentifier(viewClass objc.Class, identifier string)
	RemoveAnnotation(annotation objc.ID)
	RemoveAnnotations(annotations unsafe.Pointer)
	RemoveOverlay(overlay objc.ID)
	RemoveOverlays(overlays unsafe.Pointer)
	RendererForOverlay(overlay objc.ID) unsafe.Pointer
	SelectAnnotationAnimated(annotation objc.ID, animated bool)
	SetCameraAnimated(camera unsafe.Pointer, animated bool)
	SetCameraBoundaryAnimated(cameraBoundary unsafe.Pointer, animated bool)
	SetCameraZoomRangeAnimated(cameraZoomRange unsafe.Pointer, animated bool)
	SetCenterCoordinateAnimated(coordinate unsafe.Pointer, animated bool)
	SetRegionAnimated(region unsafe.Pointer, animated bool)
	SetUserTrackingModeAnimated(mode unsafe.Pointer, animated bool)
	SetVisibleMapRectAnimated(mapRect unsafe.Pointer, animate bool)
	SetVisibleMapRectEdgePaddingAnimated(mapRect unsafe.Pointer, insets unsafe.Pointer, animate bool)
	ShowAnnotationsAnimated(annotations unsafe.Pointer, animated bool)
	ViewForAnnotation(annotation objc.ID) unsafe.Pointer
	ViewForOverlay(overlay objc.ID) unsafe.Pointer
}

// An embeddable map interface, similar to the one that the Maps app provides.
//
// Use this class as-is to display map information and to manipulate the map contents from your app. The map view supports several display styles, including the that provides rich 2D and 3D presentations, an that provides a hybrid satellite map presentation, and that provides an imagery-based map presentation. Each of these map configurations support customization properties that refine specific elements of the map’s presentation. You can center the map on specific coordinates, specify the size of the area you want to display, and annotate the map with custom information. When you initialize a map view, you specify the initial region for that map to display by setting the property of the map. MapKit defines a region by a center point and a horizontal and vertical distance, referred to as the . The defines how much of the map is visible, and is also how you set the zoom level. For example, specifying a large span results in the user seeing a wide geographical area at a low zoom level, whereas specifying a small span results in a more narrow geographical area and a higher zoom level. In addition to setting the span programmatically, the class supports many standard interactions for changing the position and zoom level of the map. In particular, map views support flick and pinch gestures for scrolling around the map and zooming in and out. The map view enables support for these gestures by default. You can enable and disable them using the and properties. You can also use projected map coordinates instead of regions to specify some values. When you project the curved surface of the globe onto a flat surface, you get a two-dimensional version of a map where longitude lines appear to be parallel. To specify locations and distances, you use the , , and data types. Don’t subclass the class itself. You can get information about the map view’s behavior by providing a delegate object. The map view calls the methods of your custom delegate to let it know about changes in the map status and to coordinate the display of custom annotations. The delegate object can be any object in your app as long as it conforms to the protocol. For more information about implementing the delegate object, see . In macOS 10.14 and later, you can apply a light or dark appearance to your maps by modifying the property of your map view (or one of its ancestor views). Even if you specify a custom appearance, users can use the Maps app to force all maps to adopt a light appearance. Use the map view’s property to determine the actual appearance of your map. For information about how to set view appearances, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView
type MKMapView struct {
	appkit.View
}

// MKMapViewFrom constructs a [MKMapView] from an unsafe.Pointer.
//
// An embeddable map interface, similar to the one that the Maps app provides.
func MKMapViewFrom(ptr unsafe.Pointer) MKMapView {
	return MKMapView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapViewClass) Alloc() MKMapView {
	rv := objc.Send[MKMapView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapViewClass) New() MKMapView {
	rv := objc.Send[MKMapView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapView) Init() MKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapView) Autorelease() MKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapView creates a new MKMapView instance.
func NewMKMapView() MKMapView {
	return getMKMapViewClass().New()
}


// Adds the specified annotation to the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addAnnotation(_:)
func (m_ MKMapView) AddAnnotation(annotation objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAnnotation:"), annotation)
}

// Adds an array of annotation objects to the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addAnnotations(_:)
func (m_ MKMapView) AddAnnotations(annotations unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAnnotations:"), annotations)
}

// Adds a single overlay object to the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlay(_:)
func (m_ MKMapView) AddOverlay(overlay objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlay:"), overlay)
}

// Adds the overlay object to the map at the specified level.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlay(_:level:)
func (m_ MKMapView) AddOverlayLevel(overlay objc.ID, level unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlay:level:"), overlay, level)
}

// Adds an array of overlay objects to the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlays(_:)
func (m_ MKMapView) AddOverlays(overlays unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlays:"), overlays)
}

// Adds an array of overlay objects to the map at the specified level.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlays(_:level:)
func (m_ MKMapView) AddOverlaysLevel(overlays unsafe.Pointer, level unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlays:level:"), overlays, level)
}

// Returns the annotation objects within the specified map rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/annotations(in:)
func (m_ MKMapView) AnnotationsInMapRect(mapRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("annotationsInMapRect:"), mapRect)
	return rv
}

// Converts a point in the specified view’s coordinate system to a map coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toCoordinateFrom:)
func (m_ MKMapView) ConvertPointToCoordinateFromView(point coregraphics.CGPoint, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("convertPoint:toCoordinateFromView:"), point, view)
	return rv
}

// Converts a map coordinate to a point in the specified view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toPointTo:)
func (m_ MKMapView) ConvertCoordinateToPointToView(coordinate unsafe.Pointer, view unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](m_.ID, objc.Sel("convertCoordinate:toPointToView:"), coordinate, view)
	return rv
}

// Converts a map region to a rectangle in the specified view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toRectTo:)
func (m_ MKMapView) ConvertRegionToRectToView(region unsafe.Pointer, view unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("convertRegion:toRectToView:"), region, view)
	return rv
}

// Converts a rectangle in the specified view’s coordinate system to a map region.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toRegionFrom:)
func (m_ MKMapView) ConvertRectToRegionFromView(rect coregraphics.CGRect, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("convertRect:toRegionFromView:"), rect, view)
	return rv
}

// Returns a reusable annotation view using its identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/dequeueReusableAnnotationView(withIdentifier:)
func (m_ MKMapView) DequeueReusableAnnotationViewWithIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dequeueReusableAnnotationViewWithIdentifier:"), objc.String(identifier))
	return rv
}

// Returns a reusable annotation view using the specified identifier with a specified existing annotation view, if possible.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/dequeueReusableAnnotationView(withIdentifier:for:)
func (m_ MKMapView) DequeueReusableAnnotationViewWithIdentifierForAnnotation(identifier string, annotation objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dequeueReusableAnnotationViewWithIdentifier:forAnnotation:"), objc.String(identifier), annotation)
	return rv
}

// Deselects the specified annotation and hides its callout view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/deselectAnnotation(_:animated:)
func (m_ MKMapView) DeselectAnnotationAnimated(annotation objc.ID, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deselectAnnotation:animated:"), annotation, animated)
}

// Exchanges the positions of two overlay objects.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/exchangeOverlay(_:with:)
func (m_ MKMapView) ExchangeOverlayWithOverlay(overlay1 objc.ID, overlay2 objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeOverlay:withOverlay:"), overlay1, overlay2)
}

// Exchanges the position of two overlay objects at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/exchangeOverlay(at:withOverlayAt:)
func (m_ MKMapView) ExchangeOverlayAtIndexWithOverlayAtIndex(index1 uint, index2 uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeOverlayAtIndex:withOverlayAtIndex:"), index1, index2)
}

// Inserts one overlay object above another.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:above:)
func (m_ MKMapView) InsertOverlayAboveOverlay(overlay objc.ID, sibling objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:aboveOverlay:"), overlay, sibling)
}

// Inserts an overlay object into the list associated with the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:at:)
func (m_ MKMapView) InsertOverlayAtIndex(overlay objc.ID, index uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:atIndex:"), overlay, index)
}

// Inserts an overlay object into the level at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:at:level:)
func (m_ MKMapView) InsertOverlayAtIndexLevel(overlay objc.ID, index uint, level unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:atIndex:level:"), overlay, index, level)
}

// Inserts one overlay object below another.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:below:)
func (m_ MKMapView) InsertOverlayBelowOverlay(overlay objc.ID, sibling objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:belowOverlay:"), overlay, sibling)
}

// Returns a centered map rectangle with the same aspect ratio as the map view’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapRectThatFits(_:)
func (m_ MKMapView) MapRectThatFits(mapRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapRectThatFits:"), mapRect)
	return rv
}

// Returns a centered, inset map rectangle with the same aspect ratio as the map view’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapRectThatFits(_:edgePadding:)
func (m_ MKMapView) MapRectThatFitsEdgePadding(mapRect unsafe.Pointer, insets unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapRectThatFits:edgePadding:"), mapRect, insets)
	return rv
}

// Returns overlay objects in the specified level of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/overlays(in:)
func (m_ MKMapView) OverlaysInLevel(level unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("overlaysInLevel:"), level)
	return rv
}

// Adjusts the aspect ratio of the specified region to ensure that it fits in the map view’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/regionThatFits(_:)
func (m_ MKMapView) RegionThatFits(region unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("regionThatFits:"), region)
	return rv
}

// Registers an annotation view class that the map can create automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/register(_:forAnnotationViewWithReuseIdentifier:)
func (m_ MKMapView) RegisterClassForAnnotationViewWithReuseIdentifier(viewClass objc.Class, identifier string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("registerClass:forAnnotationViewWithReuseIdentifier:"), viewClass, objc.String(identifier))
}

// Removes the specified annotation object from the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeAnnotation(_:)
func (m_ MKMapView) RemoveAnnotation(annotation objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAnnotation:"), annotation)
}

// Removes an array of annotation objects from the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeAnnotations(_:)
func (m_ MKMapView) RemoveAnnotations(annotations unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAnnotations:"), annotations)
}

// Removes a single overlay object from the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeOverlay(_:)
func (m_ MKMapView) RemoveOverlay(overlay objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeOverlay:"), overlay)
}

// Removes one or more overlay objects from the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeOverlays(_:)
func (m_ MKMapView) RemoveOverlays(overlays unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeOverlays:"), overlays)
}

// Returns the renderer object for drawing the contents of the specified overlay object.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/renderer(for:)
func (m_ MKMapView) RendererForOverlay(overlay objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rendererForOverlay:"), overlay)
	return rv
}

// Selects the specified annotation and displays a callout view for it.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectAnnotation(_:animated:)
func (m_ MKMapView) SelectAnnotationAnimated(annotation objc.ID, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectAnnotation:animated:"), annotation, animated)
}

// Changes the camera to use for determining the map’s viewing parameters, and optionally animates the change.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCamera(_:animated:)
func (m_ MKMapView) SetCameraAnimated(camera unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:animated:"), camera, animated)
}

// Sets the camera boundary for the map view, specifying whether to use animation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCameraBoundary(_:animated:)
func (m_ MKMapView) SetCameraBoundaryAnimated(cameraBoundary unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:animated:"), cameraBoundary, animated)
}

// Sets the camera zoom range for the map view, specifying whether to use animation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCameraZoomRange(_:animated:)
func (m_ MKMapView) SetCameraZoomRangeAnimated(cameraZoomRange unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:animated:"), cameraZoomRange, animated)
}

// Changes the center coordinate of the map, and optionally animates the change.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCenter(_:animated:)
func (m_ MKMapView) SetCenterCoordinateAnimated(coordinate unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:animated:"), coordinate, animated)
}

// Changes the currently visible region, and optionally animates the change.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setRegion(_:animated:)
func (m_ MKMapView) SetRegionAnimated(region unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:animated:"), region, animated)
}

// Sets the mode to use for tracking the user’s location, with optional animation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setUserTrackingMode(_:animated:)
func (m_ MKMapView) SetUserTrackingModeAnimated(mode unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserTrackingMode:animated:"), mode, animated)
}

// Changes the currently visible portion of the map, and optionally animates the change.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setVisibleMapRect(_:animated:)
func (m_ MKMapView) SetVisibleMapRectAnimated(mapRect unsafe.Pointer, animate bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:animated:"), mapRect, animate)
}

// Changes the currently visible portion of the map, allowing you to specify additional space around the edges.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setVisibleMapRect(_:edgePadding:animated:)
func (m_ MKMapView) SetVisibleMapRectEdgePaddingAnimated(mapRect unsafe.Pointer, insets unsafe.Pointer, animate bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:edgePadding:animated:"), mapRect, insets, animate)
}

// Sets the visible region so that the map displays the specified annotations.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showAnnotations(_:animated:)
func (m_ MKMapView) ShowAnnotationsAnimated(annotations unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("showAnnotations:animated:"), annotations, animated)
}

// Returns the annotation view associated with the specified annotation object, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/view(for:)-33w8k
func (m_ MKMapView) ViewForAnnotation(annotation objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("viewForAnnotation:"), annotation)
	return rv
}

// Returns the view associated with the overlay object, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/view(for:)-38z60
func (m_ MKMapView) ViewForOverlay(overlay objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("viewForOverlay:"), overlay)
	return rv
}

// The visible rectangle where the map is displaying annotation views.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/annotationVisibleRect
func (m_ MKMapView) AnnotationVisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](m_.ID, objc.Sel("annotationVisibleRect"))
	return rv
}

// The annotations associated with the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/annotations
func (m_ MKMapView) Annotations() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("annotations"))
	return rv
}

// The camera to use for determining the appearance of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/camera
func (m_ MKMapView) Camera() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("camera"))
	return rv
}


// SetCamera sets the value of the camera property.
// The camera to use for determining the appearance of the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/camera
func (m_ MKMapView) SetCamera(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:"), value)
}

// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraBoundary-swift.property
func (m_ MKMapView) CameraBoundary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}


// SetCameraBoundary sets the value of the cameraBoundary property.
// The boundary of the area within which the map view’s center needs to remain.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraBoundary-swift.property
func (m_ MKMapView) SetCameraBoundary(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}

// The zoom range to apply to the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraZoomRange-swift.property
func (m_ MKMapView) CameraZoomRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}


// SetCameraZoomRange sets the value of the cameraZoomRange property.
// The zoom range to apply to the map view.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraZoomRange-swift.property
func (m_ MKMapView) SetCameraZoomRange(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}

// The map coordinate at the center of the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/centerCoordinate
func (m_ MKMapView) CenterCoordinate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("centerCoordinate"))
	return rv
}


// SetCenterCoordinate sets the value of the centerCoordinate property.
// The map coordinate at the center of the map view.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/centerCoordinate
func (m_ MKMapView) SetCenterCoordinate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:"), value)
}

// The receiver’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/delegate
func (m_ MKMapView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The receiver’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/delegate
func (m_ MKMapView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isPitchEnabled
func (m_ MKMapView) PitchEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("pitchEnabled"))
	return rv
}


// SetPitchEnabled sets the value of the pitchEnabled property.
// A Boolean value that indicates whether the map uses the camera’s pitch information.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isPitchEnabled
func (m_ MKMapView) SetPitchEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchEnabled:"), value)
}

// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isRotateEnabled
func (m_ MKMapView) RotateEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rotateEnabled"))
	return rv
}


// SetRotateEnabled sets the value of the rotateEnabled property.
// A Boolean value that indicates whether the map uses the camera’s heading information.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isRotateEnabled
func (m_ MKMapView) SetRotateEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRotateEnabled:"), value)
}

// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isScrollEnabled
func (m_ MKMapView) ScrollEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("scrollEnabled"))
	return rv
}


// SetScrollEnabled sets the value of the scrollEnabled property.
// A Boolean value that determines whether the user may scroll around the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isScrollEnabled
func (m_ MKMapView) SetScrollEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScrollEnabled:"), value)
}

// A Boolean value that indicates whether the user’s location is visible in the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isUserLocationVisible
func (m_ MKMapView) UserLocationVisible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("userLocationVisible"))
	return rv
}

// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isZoomEnabled
func (m_ MKMapView) ZoomEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("zoomEnabled"))
	return rv
}


// SetZoomEnabled sets the value of the zoomEnabled property.
// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isZoomEnabled
func (m_ MKMapView) SetZoomEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setZoomEnabled:"), value)
}

// The type of data the map view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapType
func (m_ MKMapView) MapType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapType"))
	return rv
}


// SetMapType sets the value of the mapType property.
// The type of data the map view displays.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapType
func (m_ MKMapView) SetMapType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapType:"), value)
}

// The overlay objects associated with the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/overlays
func (m_ MKMapView) Overlays() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("overlays"))
	return rv
}

// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pitchButtonVisibility
func (m_ MKMapView) PitchButtonVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}


// SetPitchButtonVisibility sets the value of the pitchButtonVisibility property.
// A value that indicates whether the map’s pitch button is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pitchButtonVisibility
func (m_ MKMapView) SetPitchButtonVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}

// The filter to use for determining the points of interest that appear on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pointOfInterestFilter
func (m_ MKMapView) PointOfInterestFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// The filter to use for determining the points of interest that appear on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pointOfInterestFilter
func (m_ MKMapView) SetPointOfInterestFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}

// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/preferredConfiguration
func (m_ MKMapView) PreferredConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// SetPreferredConfiguration sets the value of the preferredConfiguration property.
// The characteristics of the map view, including the map type and features the map displays.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/preferredConfiguration
func (m_ MKMapView) SetPreferredConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}

// The area the map view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/region
func (m_ MKMapView) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// SetRegion sets the value of the region property.
// The area the map view displays.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/region
func (m_ MKMapView) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}

// The property that describes which selectable features the map responds to.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectableMapFeatures
func (m_ MKMapView) SelectableMapFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("selectableMapFeatures"))
	return rv
}


// SetSelectableMapFeatures sets the value of the selectableMapFeatures property.
// The property that describes which selectable features the map responds to.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectableMapFeatures
func (m_ MKMapView) SetSelectableMapFeatures(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectableMapFeatures:"), value)
}

// The selected annotations.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectedAnnotations
func (m_ MKMapView) SelectedAnnotations() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("selectedAnnotations"))
	return rv
}


// SetSelectedAnnotations sets the value of the selectedAnnotations property.
// The selected annotations.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectedAnnotations
func (m_ MKMapView) SetSelectedAnnotations(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedAnnotations:"), nsArray)
}

// A Boolean value that indicates whether the map displays extruded building information on supported map types.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsBuildings
func (m_ MKMapView) ShowsBuildings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsBuildings"))
	return rv
}


// SetShowsBuildings sets the value of the showsBuildings property.
// A Boolean value that indicates whether the map displays extruded building information on supported map types.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsBuildings
func (m_ MKMapView) SetShowsBuildings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsBuildings:"), value)
}

// A Boolean value that indicates whether the map displays a compass control.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsCompass
func (m_ MKMapView) ShowsCompass() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsCompass"))
	return rv
}


// SetShowsCompass sets the value of the showsCompass property.
// A Boolean value that indicates whether the map displays a compass control.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsCompass
func (m_ MKMapView) SetShowsCompass(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsCompass:"), value)
}

// A Boolean value that indicates whether the map displays the pitch control.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPitchControl
func (m_ MKMapView) ShowsPitchControl() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPitchControl"))
	return rv
}


// SetShowsPitchControl sets the value of the showsPitchControl property.
// A Boolean value that indicates whether the map displays the pitch control.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPitchControl
func (m_ MKMapView) SetShowsPitchControl(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPitchControl:"), value)
}

// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPointsOfInterest
func (m_ MKMapView) ShowsPointsOfInterest() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPointsOfInterest"))
	return rv
}


// SetShowsPointsOfInterest sets the value of the showsPointsOfInterest property.
// A Boolean value that indicates whether the map displays point-of-interest information.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPointsOfInterest
func (m_ MKMapView) SetShowsPointsOfInterest(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPointsOfInterest:"), value)
}

// A Boolean value that indicates whether the map shows scale information.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsScale
func (m_ MKMapView) ShowsScale() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsScale"))
	return rv
}


// SetShowsScale sets the value of the showsScale property.
// A Boolean value that indicates whether the map shows scale information.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsScale
func (m_ MKMapView) SetShowsScale(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsScale:"), value)
}

// A Boolean value that indicates whether the map displays traffic information.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsTraffic
func (m_ MKMapView) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}


// SetShowsTraffic sets the value of the showsTraffic property.
// A Boolean value that indicates whether the map displays traffic information.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsTraffic
func (m_ MKMapView) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}

// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserLocation
func (m_ MKMapView) ShowsUserLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserLocation"))
	return rv
}


// SetShowsUserLocation sets the value of the showsUserLocation property.
// A Boolean value that indicates whether the map tries to display the user’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserLocation
func (m_ MKMapView) SetShowsUserLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserLocation:"), value)
}

// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserTrackingButton
func (m_ MKMapView) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}


// SetShowsUserTrackingButton sets the value of the showsUserTrackingButton property.
// A Boolean value that indicates whether the map displays the user tracking button.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserTrackingButton
func (m_ MKMapView) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}

// A Boolean value that indicates whether the map displays zoom controls.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsZoomControls
func (m_ MKMapView) ShowsZoomControls() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsZoomControls"))
	return rv
}


// SetShowsZoomControls sets the value of the showsZoomControls property.
// A Boolean value that indicates whether the map displays zoom controls.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsZoomControls
func (m_ MKMapView) SetShowsZoomControls(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsZoomControls:"), value)
}

// The annotation object that represents the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/userLocation
func (m_ MKMapView) UserLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("userLocation"))
	return rv
}

// The mode to use for tracking the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/userTrackingMode
func (m_ MKMapView) UserTrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("userTrackingMode"))
	return rv
}


// SetUserTrackingMode sets the value of the userTrackingMode property.
// The mode to use for tracking the user’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/userTrackingMode
func (m_ MKMapView) SetUserTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserTrackingMode:"), value)
}

// The area visible in the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/visibleMapRect
func (m_ MKMapView) VisibleMapRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("visibleMapRect"))
	return rv
}


// SetVisibleMapRect sets the value of the visibleMapRect property.
// The area visible in the map view.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/visibleMapRect
func (m_ MKMapView) SetVisibleMapRect(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:"), value)
}



