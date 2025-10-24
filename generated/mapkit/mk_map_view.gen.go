// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapView */


/* debug [class_header]: Header for MKMapView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapView */
// An interface definition for the [MKMapView] class.
type IMKMapView interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKMapView */
	// properties:
	Overlays() objectivec.IObject
	SetOverlays(value objectivec.IObject)
	Annotations() []objc.ID
	AnnotationVisibleRect() corefoundation.CGRect
	Camera() IMKMapCamera
	SetCamera(value IMKMapCamera)
	CameraBoundary() IMKMapCameraBoundary
	SetCameraBoundary(value IMKMapCameraBoundary)
	CameraZoomRange() IMKMapCameraZoomRange
	SetCameraZoomRange(value IMKMapCameraZoomRange)
	CenterCoordinate() LocationCoordinate2D /* not a class type */
	SetCenterCoordinate(value LocationCoordinate2D /* not a class type */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	PitchEnabled() bool
	SetPitchEnabled(value bool)
	RotateEnabled() bool
	SetRotateEnabled(value bool)
	ScrollEnabled() bool
	SetScrollEnabled(value bool)
	UserLocationVisible() bool
	ZoomEnabled() bool
	SetZoomEnabled(value bool)
	MapType() MKMapType
	SetMapType(value MKMapType)
	PitchButtonVisibility() MKFeatureVisibility
	SetPitchButtonVisibility(value MKFeatureVisibility)
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	PreferredConfiguration() IMKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
	SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */)
	SelectedAnnotations() []objc.ID
	SetSelectedAnnotations(value []objc.ID)
	ShowsBuildings() bool
	SetShowsBuildings(value bool)
	ShowsCompass() bool
	SetShowsCompass(value bool)
	ShowsPitchControl() bool
	SetShowsPitchControl(value bool)
	ShowsPointsOfInterest() bool
	SetShowsPointsOfInterest(value bool)
	ShowsScale() bool
	SetShowsScale(value bool)
	ShowsTraffic() bool
	SetShowsTraffic(value bool)
	ShowsUserLocation() bool
	SetShowsUserLocation(value bool)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
	ShowsZoomControls() bool
	SetShowsZoomControls(value bool)
	UserLocation() IMKUserLocation
	UserTrackingMode() MKUserTrackingMode
	SetUserTrackingMode(value MKUserTrackingMode)
	VisibleMapRect() objc.IObject /* cross-framework: MKMapRect */
	SetVisibleMapRect(value objc.IObject /* cross-framework: MKMapRect */)
	Appearance() appkit.Appearance
	SetAppearance(value appkit.Appearance)
	EffectiveAppearance() appkit.Appearance
	SetEffectiveAppearance(value appkit.Appearance)
	IsPitchEnabled() bool
	SetIsPitchEnabled(value bool)
	IsRotateEnabled() bool
	SetIsRotateEnabled(value bool)
	IsScrollEnabled() bool
	SetIsScrollEnabled(value bool)
	IsUserLocationVisible() bool
	SetIsUserLocationVisible(value bool)
	IsZoomEnabled() bool
	SetIsZoomEnabled(value bool)
	MKMapViewDefaultAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */
	MKMapViewDefaultClusterAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapView */
	// methods:
	AddOverlay()
	AddAnnotation(annotation unsafe.Pointer)
	AddAnnotations(annotations []objc.ID)
	AddOverlayWithOverlay(overlay unsafe.Pointer)
	AddOverlayLevel(overlay unsafe.Pointer, level MKOverlayLevel)
	AddOverlays(overlays []objc.ID)
	AddOverlaysLevel(overlays []objc.ID, level MKOverlayLevel)
	AnnotationsInMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) unsafe.Pointer
	ConvertPointToCoordinateFromView(point corefoundation.CGPoint, view objc.IObject /* cross-framework: View */) LocationCoordinate2D /* not a class type */
	ConvertCoordinateToPointToView(coordinate LocationCoordinate2D /* not a class type */, view objc.IObject /* cross-framework: View */) corefoundation.CGPoint
	ConvertRegionToRectToView(region objc.IObject /* cross-framework: MKCoordinateRegion */, view objc.IObject /* cross-framework: View */) corefoundation.CGRect
	ConvertRectToRegionFromView(rect corefoundation.CGRect, view objc.IObject /* cross-framework: View */) objc.IObject /* cross-framework: MKCoordinateRegion */
	DequeueReusableAnnotationViewWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: MKAnnotationView */
	DequeueReusableAnnotationViewWithIdentifierForAnnotation(identifier objc.IObject /* cross-framework: NSString */, annotation unsafe.Pointer) objc.IObject /* cross-framework: MKAnnotationView */
	DeselectAnnotationAnimated(annotation unsafe.Pointer, animated bool)
	ExchangeOverlayWithOverlay(overlay1 unsafe.Pointer, overlay2 unsafe.Pointer)
	ExchangeOverlayAtIndexWithOverlayAtIndex(index1 uint, index2 uint)
	InsertOverlayAboveOverlay(overlay unsafe.Pointer, sibling unsafe.Pointer)
	InsertOverlayAtIndex(overlay unsafe.Pointer, index uint)
	InsertOverlayAtIndexLevel(overlay unsafe.Pointer, index uint, level MKOverlayLevel)
	InsertOverlayBelowOverlay(overlay unsafe.Pointer, sibling unsafe.Pointer)
	MapRectThatFits(mapRect objc.IObject /* cross-framework: MKMapRect */) objc.IObject /* cross-framework: MKMapRect */
	MapRectThatFitsEdgePadding(mapRect objc.IObject /* cross-framework: MKMapRect */, insets foundation.EdgeInsets) objc.IObject /* cross-framework: MKMapRect */
	OverlaysInLevel(level MKOverlayLevel) []objc.ID
	RegionThatFits(region objc.IObject /* cross-framework: MKCoordinateRegion */) objc.IObject /* cross-framework: MKCoordinateRegion */
	RegisterClassForAnnotationViewWithReuseIdentifier(viewClass objc.Class, identifier objc.IObject /* cross-framework: NSString */)
	RemoveAnnotation(annotation unsafe.Pointer)
	RemoveAnnotations(annotations []objc.ID)
	RemoveOverlay(overlay unsafe.Pointer)
	RemoveOverlays(overlays []objc.ID)
	RendererForOverlay(overlay unsafe.Pointer) IMKOverlayRenderer
	SelectAnnotationAnimated(annotation unsafe.Pointer, animated bool)
	SetCameraAnimated(camera IMKMapCamera, animated bool)
	SetCameraBoundaryAnimated(cameraBoundary IMKMapCameraBoundary, animated bool)
	SetCameraZoomRangeAnimated(cameraZoomRange IMKMapCameraZoomRange, animated bool)
	SetCenterCoordinateAnimated(coordinate LocationCoordinate2D /* not a class type */, animated bool)
	SetRegionAnimated(region objc.IObject /* cross-framework: MKCoordinateRegion */, animated bool)
	SetUserTrackingModeAnimated(mode MKUserTrackingMode, animated bool)
	SetVisibleMapRectAnimated(mapRect objc.IObject /* cross-framework: MKMapRect */, animate bool)
	SetVisibleMapRectEdgePaddingAnimated(mapRect objc.IObject /* cross-framework: MKMapRect */, insets foundation.EdgeInsets, animate bool)
	ShowAnnotationsAnimated(annotations []objc.ID, animated bool)
	ViewForAnnotation(annotation unsafe.Pointer) objc.IObject /* cross-framework: MKAnnotationView */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapView */
// Alloc allocates a new instance without initialization.
func (mc _MKMapViewClass) Alloc() MKMapView {
	rv := objc.Send[MKMapView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapView */
// An embeddable map interface, similar to the one that the Maps app provides.
//
// Use this class as-is to display map information and to manipulate the map contents from your app. The map view supports several display styles, including the that provides rich 2D and 3D presentations, an that provides a hybrid satellite map presentation, and that provides an imagery-based map presentation. Each of these map configurations support customization properties that refine specific elements of the map’s presentation. You can center the map on specific coordinates, specify the size of the area you want to display, and annotate the map with custom information. When you initialize a map view, you specify the initial region for that map to display by setting the property of the map. MapKit defines a region by a center point and a horizontal and vertical distance, referred to as the . The defines how much of the map is visible, and is also how you set the zoom level. For example, specifying a large span results in the user seeing a wide geographical area at a low zoom level, whereas specifying a small span results in a more narrow geographical area and a higher zoom level. In addition to setting the span programmatically, the class supports many standard interactions for changing the position and zoom level of the map. In particular, map views support flick and pinch gestures for scrolling around the map and zooming in and out. The map view enables support for these gestures by default. You can enable and disable them using the and properties. You can also use projected map coordinates instead of regions to specify some values. When you project the curved surface of the globe onto a flat surface, you get a two-dimensional version of a map where longitude lines appear to be parallel. To specify locations and distances, you use the , , and data types. Don’t subclass the class itself. You can get information about the map view’s behavior by providing a delegate object. The map view calls the methods of your custom delegate to let it know about changes in the map status and to coordinate the display of custom annotations. The delegate object can be any object in your app as long as it conforms to the protocol. For more information about implementing the delegate object, see . In macOS 10.14 and later, you can apply a light or dark appearance to your maps by modifying the property of your map view (or one of its ancestor views). Even if you specify a custom appearance, users can use the Maps app to force all maps to adopt a light appearance. Use the map view’s property to determine the actual appearance of your map. For information about how to set view appearances, see .


// An embeddable map interface, similar to the one that the Maps app provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView
type MKMapView struct {
	View
}

// MKMapViewFrom constructs a [MKMapView] from an unsafe.Pointer.
//
// An embeddable map interface, similar to the one that the Maps app provides.
func MKMapViewFrom(ptr unsafe.Pointer) MKMapView {
	return MKMapView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapView */

// Adds the overlay object to the map at the specified level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/1452635-addoverlay
func (m_ MKMapView) AddOverlay() {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlay"))
}/* debug [instance_methods/method]: AddOverlay */


// Adds the specified annotation to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addAnnotation(_:)
func (m_ MKMapView) AddAnnotation(annotation unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAnnotation:"), annotation)
}/* debug [instance_methods/method]: AddAnnotation */


// Adds an array of annotation objects to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addAnnotations(_:)
func (m_ MKMapView) AddAnnotations(annotations []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addAnnotations:"), annotations)
}/* debug [instance_methods/method]: AddAnnotations */


// Adds a single overlay object to the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlay(_:)
func (m_ MKMapView) AddOverlayWithOverlay(overlay unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlay:"), overlay)
}/* debug [instance_methods/method]: AddOverlayWithOverlay */


// Adds the overlay object to the map at the specified level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlay(_:level:)
func (m_ MKMapView) AddOverlayLevel(overlay unsafe.Pointer, level MKOverlayLevel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlay:level:"), overlay, level)
}/* debug [instance_methods/method]: AddOverlayLevel */


// Adds an array of overlay objects to the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlays(_:)
func (m_ MKMapView) AddOverlays(overlays []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlays:"), overlays)
}/* debug [instance_methods/method]: AddOverlays */


// Adds an array of overlay objects to the map at the specified level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/addOverlays(_:level:)
func (m_ MKMapView) AddOverlaysLevel(overlays []objc.ID, level MKOverlayLevel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addOverlays:level:"), overlays, level)
}/* debug [instance_methods/method]: AddOverlaysLevel */


// Returns the annotation objects within the specified map rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/annotations(in:)
func (m_ MKMapView) AnnotationsInMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("annotationsInMapRect:"), mapRect)
	return rv
}/* debug [instance_methods/method]: AnnotationsInMapRect */


// Converts a point in the specified view’s coordinate system to a map coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toCoordinateFrom:)
func (m_ MKMapView) ConvertPointToCoordinateFromView(point corefoundation.CGPoint, view objc.IObject /* cross-framework: View */) LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("convertPoint:toCoordinateFromView:"), point, view)
	return rv
}/* debug [instance_methods/method]: ConvertPointToCoordinateFromView */


// Converts a map coordinate to a point in the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toPointTo:)
func (m_ MKMapView) ConvertCoordinateToPointToView(coordinate LocationCoordinate2D /* not a class type */, view objc.IObject /* cross-framework: View */) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](m_.ID, objc.Sel("convertCoordinate:toPointToView:"), coordinate, view)
	return rv
}/* debug [instance_methods/method]: ConvertCoordinateToPointToView */


// Converts a map region to a rectangle in the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toRectTo:)
func (m_ MKMapView) ConvertRegionToRectToView(region objc.IObject /* cross-framework: MKCoordinateRegion */, view objc.IObject /* cross-framework: View */) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("convertRegion:toRectToView:"), region, view)
	return rv
}/* debug [instance_methods/method]: ConvertRegionToRectToView */


// Converts a rectangle in the specified view’s coordinate system to a map region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/convert(_:toRegionFrom:)
func (m_ MKMapView) ConvertRectToRegionFromView(rect corefoundation.CGRect, view objc.IObject /* cross-framework: View */) objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("convertRect:toRegionFromView:"), rect, view)
	return rv
}/* debug [instance_methods/method]: ConvertRectToRegionFromView */


// Returns a reusable annotation view using its identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/dequeueReusableAnnotationView(withIdentifier:)
func (m_ MKMapView) DequeueReusableAnnotationViewWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: MKAnnotationView */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("dequeueReusableAnnotationViewWithIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: DequeueReusableAnnotationViewWithIdentifier */


// Returns a reusable annotation view using the specified identifier with a specified existing annotation view, if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/dequeueReusableAnnotationView(withIdentifier:for:)
func (m_ MKMapView) DequeueReusableAnnotationViewWithIdentifierForAnnotation(identifier objc.IObject /* cross-framework: NSString */, annotation unsafe.Pointer) objc.IObject /* cross-framework: MKAnnotationView */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("dequeueReusableAnnotationViewWithIdentifier:forAnnotation:"), identifier, annotation)
	return rv
}/* debug [instance_methods/method]: DequeueReusableAnnotationViewWithIdentifierForAnnotation */


// Deselects the specified annotation and hides its callout view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/deselectAnnotation(_:animated:)
func (m_ MKMapView) DeselectAnnotationAnimated(annotation unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deselectAnnotation:animated:"), annotation, animated)
}/* debug [instance_methods/method]: DeselectAnnotationAnimated */


// Exchanges the positions of two overlay objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/exchangeOverlay(_:with:)
func (m_ MKMapView) ExchangeOverlayWithOverlay(overlay1 unsafe.Pointer, overlay2 unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeOverlay:withOverlay:"), overlay1, overlay2)
}/* debug [instance_methods/method]: ExchangeOverlayWithOverlay */


// Exchanges the position of two overlay objects at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/exchangeOverlay(at:withOverlayAt:)
func (m_ MKMapView) ExchangeOverlayAtIndexWithOverlayAtIndex(index1 uint, index2 uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeOverlayAtIndex:withOverlayAtIndex:"), index1, index2)
}/* debug [instance_methods/method]: ExchangeOverlayAtIndexWithOverlayAtIndex */


// Inserts one overlay object above another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:above:)
func (m_ MKMapView) InsertOverlayAboveOverlay(overlay unsafe.Pointer, sibling unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:aboveOverlay:"), overlay, sibling)
}/* debug [instance_methods/method]: InsertOverlayAboveOverlay */


// Inserts an overlay object into the list associated with the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:at:)
func (m_ MKMapView) InsertOverlayAtIndex(overlay unsafe.Pointer, index uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:atIndex:"), overlay, index)
}/* debug [instance_methods/method]: InsertOverlayAtIndex */


// Inserts an overlay object into the level at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:at:level:)
func (m_ MKMapView) InsertOverlayAtIndexLevel(overlay unsafe.Pointer, index uint, level MKOverlayLevel) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:atIndex:level:"), overlay, index, level)
}/* debug [instance_methods/method]: InsertOverlayAtIndexLevel */


// Inserts one overlay object below another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/insertOverlay(_:below:)
func (m_ MKMapView) InsertOverlayBelowOverlay(overlay unsafe.Pointer, sibling unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertOverlay:belowOverlay:"), overlay, sibling)
}/* debug [instance_methods/method]: InsertOverlayBelowOverlay */


// Returns a centered map rectangle with the same aspect ratio as the map view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapRectThatFits(_:)
func (m_ MKMapView) MapRectThatFits(mapRect objc.IObject /* cross-framework: MKMapRect */) objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mapRectThatFits:"), mapRect)
	return rv
}/* debug [instance_methods/method]: MapRectThatFits */


// Returns a centered, inset map rectangle with the same aspect ratio as the map view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapRectThatFits(_:edgePadding:)
func (m_ MKMapView) MapRectThatFitsEdgePadding(mapRect objc.IObject /* cross-framework: MKMapRect */, insets foundation.EdgeInsets) objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mapRectThatFits:edgePadding:"), mapRect, insets)
	return rv
}/* debug [instance_methods/method]: MapRectThatFitsEdgePadding */


// Returns overlay objects in the specified level of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/overlays(in:)
func (m_ MKMapView) OverlaysInLevel(level MKOverlayLevel) []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("overlaysInLevel:"), level)
	return rv
}/* debug [instance_methods/method]: OverlaysInLevel */


// Adjusts the aspect ratio of the specified region to ensure that it fits in the map view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/regionThatFits(_:)
func (m_ MKMapView) RegionThatFits(region objc.IObject /* cross-framework: MKCoordinateRegion */) objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("regionThatFits:"), region)
	return rv
}/* debug [instance_methods/method]: RegionThatFits */


// Registers an annotation view class that the map can create automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/register(_:forAnnotationViewWithReuseIdentifier:)
func (m_ MKMapView) RegisterClassForAnnotationViewWithReuseIdentifier(viewClass objc.Class, identifier objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("registerClass:forAnnotationViewWithReuseIdentifier:"), viewClass, identifier)
}/* debug [instance_methods/method]: RegisterClassForAnnotationViewWithReuseIdentifier */


// Removes the specified annotation object from the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeAnnotation(_:)
func (m_ MKMapView) RemoveAnnotation(annotation unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAnnotation:"), annotation)
}/* debug [instance_methods/method]: RemoveAnnotation */


// Removes an array of annotation objects from the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeAnnotations(_:)
func (m_ MKMapView) RemoveAnnotations(annotations []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAnnotations:"), annotations)
}/* debug [instance_methods/method]: RemoveAnnotations */


// Removes a single overlay object from the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeOverlay(_:)
func (m_ MKMapView) RemoveOverlay(overlay unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeOverlay:"), overlay)
}/* debug [instance_methods/method]: RemoveOverlay */


// Removes one or more overlay objects from the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/removeOverlays(_:)
func (m_ MKMapView) RemoveOverlays(overlays []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeOverlays:"), overlays)
}/* debug [instance_methods/method]: RemoveOverlays */


// Returns the renderer object for drawing the contents of the specified overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/renderer(for:)
func (m_ MKMapView) RendererForOverlay(overlay unsafe.Pointer) IMKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](m_.ID, objc.Sel("rendererForOverlay:"), overlay)
	return rv
}/* debug [instance_methods/method]: RendererForOverlay */


// Selects the specified annotation and displays a callout view for it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectAnnotation(_:animated:)
func (m_ MKMapView) SelectAnnotationAnimated(annotation unsafe.Pointer, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectAnnotation:animated:"), annotation, animated)
}/* debug [instance_methods/method]: SelectAnnotationAnimated */


// Changes the camera to use for determining the map’s viewing parameters, and optionally animates the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCamera(_:animated:)
func (m_ MKMapView) SetCameraAnimated(camera IMKMapCamera, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:animated:"), camera, animated)
}/* debug [instance_methods/method]: SetCameraAnimated */


// Sets the camera boundary for the map view, specifying whether to use animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCameraBoundary(_:animated:)
func (m_ MKMapView) SetCameraBoundaryAnimated(cameraBoundary IMKMapCameraBoundary, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:animated:"), cameraBoundary, animated)
}/* debug [instance_methods/method]: SetCameraBoundaryAnimated */


// Sets the camera zoom range for the map view, specifying whether to use animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCameraZoomRange(_:animated:)
func (m_ MKMapView) SetCameraZoomRangeAnimated(cameraZoomRange IMKMapCameraZoomRange, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:animated:"), cameraZoomRange, animated)
}/* debug [instance_methods/method]: SetCameraZoomRangeAnimated */


// Changes the center coordinate of the map, and optionally animates the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setCenter(_:animated:)
func (m_ MKMapView) SetCenterCoordinateAnimated(coordinate LocationCoordinate2D /* not a class type */, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:animated:"), coordinate, animated)
}/* debug [instance_methods/method]: SetCenterCoordinateAnimated */


// Changes the currently visible region, and optionally animates the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setRegion(_:animated:)
func (m_ MKMapView) SetRegionAnimated(region objc.IObject /* cross-framework: MKCoordinateRegion */, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:animated:"), region, animated)
}/* debug [instance_methods/method]: SetRegionAnimated */


// Sets the mode to use for tracking the user’s location, with optional animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setUserTrackingMode(_:animated:)
func (m_ MKMapView) SetUserTrackingModeAnimated(mode MKUserTrackingMode, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserTrackingMode:animated:"), mode, animated)
}/* debug [instance_methods/method]: SetUserTrackingModeAnimated */


// Changes the currently visible portion of the map, and optionally animates the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setVisibleMapRect(_:animated:)
func (m_ MKMapView) SetVisibleMapRectAnimated(mapRect objc.IObject /* cross-framework: MKMapRect */, animate bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:animated:"), mapRect, animate)
}/* debug [instance_methods/method]: SetVisibleMapRectAnimated */


// Changes the currently visible portion of the map, allowing you to specify additional space around the edges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/setVisibleMapRect(_:edgePadding:animated:)
func (m_ MKMapView) SetVisibleMapRectEdgePaddingAnimated(mapRect objc.IObject /* cross-framework: MKMapRect */, insets foundation.EdgeInsets, animate bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:edgePadding:animated:"), mapRect, insets, animate)
}/* debug [instance_methods/method]: SetVisibleMapRectEdgePaddingAnimated */


// Sets the visible region so that the map displays the specified annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showAnnotations(_:animated:)
func (m_ MKMapView) ShowAnnotationsAnimated(annotations []objc.ID, animated bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("showAnnotations:animated:"), annotations, animated)
}/* debug [instance_methods/method]: ShowAnnotationsAnimated */


// Returns the annotation view associated with the specified annotation object, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/view(for:)-33w8k
func (m_ MKMapView) ViewForAnnotation(annotation unsafe.Pointer) objc.IObject /* cross-framework: MKAnnotationView */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("viewForAnnotation:"), annotation)
	return rv
}/* debug [instance_methods/method]: ViewForAnnotation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapView */

// The overlay objects associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/1452784-overlays
func (m_ MKMapView) Overlays() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("overlays"))
	return rv
}/* debug [instance_properties/getter]: overlays */


// The overlay objects associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/1452784-overlays
func (m_ MKMapView) SetOverlays(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOverlays:"), value)
}/* debug [instance_properties/setter]: overlays */


// The annotations associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/annotations
func (m_ MKMapView) Annotations() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("annotations"))
	return rv
}/* debug [instance_properties/getter]: annotations */


// The visible rectangle where the map is displaying annotation views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/annotationVisibleRect
func (m_ MKMapView) AnnotationVisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("annotationVisibleRect"))
	return rv
}/* debug [instance_properties/getter]: annotationVisibleRect */


// The camera to use for determining the appearance of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/camera
func (m_ MKMapView) Camera() IMKMapCamera {
	rv := objc.Send[MKMapCamera](m_.ID, objc.Sel("camera"))
	return rv
}/* debug [instance_properties/getter]: camera */


// The camera to use for determining the appearance of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/camera
func (m_ MKMapView) SetCamera(value IMKMapCamera) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:"), value)
}/* debug [instance_properties/setter]: camera */


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraBoundary-swift.property
func (m_ MKMapView) CameraBoundary() IMKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}/* debug [instance_properties/getter]: cameraBoundary */


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraBoundary-swift.property
func (m_ MKMapView) SetCameraBoundary(value IMKMapCameraBoundary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}/* debug [instance_properties/setter]: cameraBoundary */


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraZoomRange-swift.property
func (m_ MKMapView) CameraZoomRange() IMKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}/* debug [instance_properties/getter]: cameraZoomRange */


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/cameraZoomRange-swift.property
func (m_ MKMapView) SetCameraZoomRange(value IMKMapCameraZoomRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}/* debug [instance_properties/setter]: cameraZoomRange */


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/centerCoordinate
func (m_ MKMapView) CenterCoordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("centerCoordinate"))
	return rv
}/* debug [instance_properties/getter]: centerCoordinate */


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/centerCoordinate
func (m_ MKMapView) SetCenterCoordinate(value LocationCoordinate2D /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:"), value)
}/* debug [instance_properties/setter]: centerCoordinate */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/delegate
func (m_ MKMapView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/delegate
func (m_ MKMapView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isPitchEnabled
func (m_ MKMapView) PitchEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("pitchEnabled"))
	return rv
}/* debug [instance_properties/getter]: pitchEnabled */


// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isPitchEnabled
func (m_ MKMapView) SetPitchEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchEnabled:"), value)
}/* debug [instance_properties/setter]: pitchEnabled */


// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isRotateEnabled
func (m_ MKMapView) RotateEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("rotateEnabled"))
	return rv
}/* debug [instance_properties/getter]: rotateEnabled */


// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isRotateEnabled
func (m_ MKMapView) SetRotateEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRotateEnabled:"), value)
}/* debug [instance_properties/setter]: rotateEnabled */


// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isScrollEnabled
func (m_ MKMapView) ScrollEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("scrollEnabled"))
	return rv
}/* debug [instance_properties/getter]: scrollEnabled */


// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isScrollEnabled
func (m_ MKMapView) SetScrollEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScrollEnabled:"), value)
}/* debug [instance_properties/setter]: scrollEnabled */


// A Boolean value that indicates whether the user’s location is visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isUserLocationVisible
func (m_ MKMapView) UserLocationVisible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("userLocationVisible"))
	return rv
}/* debug [instance_properties/getter]: userLocationVisible */


// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isZoomEnabled
func (m_ MKMapView) ZoomEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("zoomEnabled"))
	return rv
}/* debug [instance_properties/getter]: zoomEnabled */


// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/isZoomEnabled
func (m_ MKMapView) SetZoomEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setZoomEnabled:"), value)
}/* debug [instance_properties/setter]: zoomEnabled */


// The type of data the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapType
func (m_ MKMapView) MapType() MKMapType {
	rv := objc.Send[MKMapType](m_.ID, objc.Sel("mapType"))
	return rv
}/* debug [instance_properties/getter]: mapType */


// The type of data the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/mapType
func (m_ MKMapView) SetMapType(value MKMapType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapType:"), value)
}/* debug [instance_properties/setter]: mapType */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pitchButtonVisibility
func (m_ MKMapView) PitchButtonVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}/* debug [instance_properties/getter]: pitchButtonVisibility */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pitchButtonVisibility
func (m_ MKMapView) SetPitchButtonVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}/* debug [instance_properties/setter]: pitchButtonVisibility */


// The filter to use for determining the points of interest that appear on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pointOfInterestFilter
func (m_ MKMapView) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// The filter to use for determining the points of interest that appear on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/pointOfInterestFilter
func (m_ MKMapView) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/preferredConfiguration
func (m_ MKMapView) PreferredConfiguration() IMKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}/* debug [instance_properties/getter]: preferredConfiguration */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/preferredConfiguration
func (m_ MKMapView) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}/* debug [instance_properties/setter]: preferredConfiguration */


// The area the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/region
func (m_ MKMapView) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// The area the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/region
func (m_ MKMapView) SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}/* debug [instance_properties/setter]: region */


// The selected annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectedAnnotations
func (m_ MKMapView) SelectedAnnotations() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("selectedAnnotations"))
	return rv
}/* debug [instance_properties/getter]: selectedAnnotations */


// The selected annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectedAnnotations
func (m_ MKMapView) SetSelectedAnnotations(value []objc.ID) {
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
}/* debug [instance_properties/setter]: selectedAnnotations */


// A Boolean value that indicates whether the map displays extruded building information on supported map types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsBuildings
func (m_ MKMapView) ShowsBuildings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsBuildings"))
	return rv
}/* debug [instance_properties/getter]: showsBuildings */


// A Boolean value that indicates whether the map displays extruded building information on supported map types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsBuildings
func (m_ MKMapView) SetShowsBuildings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsBuildings:"), value)
}/* debug [instance_properties/setter]: showsBuildings */


// A Boolean value that indicates whether the map displays a compass control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsCompass
func (m_ MKMapView) ShowsCompass() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsCompass"))
	return rv
}/* debug [instance_properties/getter]: showsCompass */


// A Boolean value that indicates whether the map displays a compass control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsCompass
func (m_ MKMapView) SetShowsCompass(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsCompass:"), value)
}/* debug [instance_properties/setter]: showsCompass */


// A Boolean value that indicates whether the map displays the pitch control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPitchControl
func (m_ MKMapView) ShowsPitchControl() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPitchControl"))
	return rv
}/* debug [instance_properties/getter]: showsPitchControl */


// A Boolean value that indicates whether the map displays the pitch control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPitchControl
func (m_ MKMapView) SetShowsPitchControl(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPitchControl:"), value)
}/* debug [instance_properties/setter]: showsPitchControl */


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPointsOfInterest
func (m_ MKMapView) ShowsPointsOfInterest() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPointsOfInterest"))
	return rv
}/* debug [instance_properties/getter]: showsPointsOfInterest */


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPointsOfInterest
func (m_ MKMapView) SetShowsPointsOfInterest(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPointsOfInterest:"), value)
}/* debug [instance_properties/setter]: showsPointsOfInterest */


// A Boolean value that indicates whether the map shows scale information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsScale
func (m_ MKMapView) ShowsScale() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsScale"))
	return rv
}/* debug [instance_properties/getter]: showsScale */


// A Boolean value that indicates whether the map shows scale information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsScale
func (m_ MKMapView) SetShowsScale(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsScale:"), value)
}/* debug [instance_properties/setter]: showsScale */


// A Boolean value that indicates whether the map displays traffic information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsTraffic
func (m_ MKMapView) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}/* debug [instance_properties/getter]: showsTraffic */


// A Boolean value that indicates whether the map displays traffic information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsTraffic
func (m_ MKMapView) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}/* debug [instance_properties/setter]: showsTraffic */


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserLocation
func (m_ MKMapView) ShowsUserLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserLocation"))
	return rv
}/* debug [instance_properties/getter]: showsUserLocation */


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserLocation
func (m_ MKMapView) SetShowsUserLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserLocation:"), value)
}/* debug [instance_properties/setter]: showsUserLocation */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserTrackingButton
func (m_ MKMapView) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}/* debug [instance_properties/getter]: showsUserTrackingButton */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserTrackingButton
func (m_ MKMapView) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}/* debug [instance_properties/setter]: showsUserTrackingButton */


// A Boolean value that indicates whether the map displays zoom controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsZoomControls
func (m_ MKMapView) ShowsZoomControls() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsZoomControls"))
	return rv
}/* debug [instance_properties/getter]: showsZoomControls */


// A Boolean value that indicates whether the map displays zoom controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsZoomControls
func (m_ MKMapView) SetShowsZoomControls(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsZoomControls:"), value)
}/* debug [instance_properties/setter]: showsZoomControls */


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/userLocation
func (m_ MKMapView) UserLocation() IMKUserLocation {
	rv := objc.Send[MKUserLocation](m_.ID, objc.Sel("userLocation"))
	return rv
}/* debug [instance_properties/getter]: userLocation */


// The mode to use for tracking the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/userTrackingMode
func (m_ MKMapView) UserTrackingMode() MKUserTrackingMode {
	rv := objc.Send[MKUserTrackingMode](m_.ID, objc.Sel("userTrackingMode"))
	return rv
}/* debug [instance_properties/getter]: userTrackingMode */


// The mode to use for tracking the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/userTrackingMode
func (m_ MKMapView) SetUserTrackingMode(value MKUserTrackingMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserTrackingMode:"), value)
}/* debug [instance_properties/setter]: userTrackingMode */


// The area visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/visibleMapRect
func (m_ MKMapView) VisibleMapRect() objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("visibleMapRect"))
	return rv
}/* debug [instance_properties/getter]: visibleMapRect */


// The area visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/visibleMapRect
func (m_ MKMapView) SetVisibleMapRect(value objc.IObject /* cross-framework: MKMapRect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:"), value)
}/* debug [instance_properties/setter]: visibleMapRect */


// The appearance of the receiver, in an `NSAppearance` object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (m_ MKMapView) Appearance() appkit.Appearance {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("appearance"))
	return rv
}/* debug [instance_properties/getter]: appearance */


// The appearance of the receiver, in an `NSAppearance` object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (m_ MKMapView) SetAppearance(value appkit.Appearance) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAppearance:"), value)
}/* debug [instance_properties/setter]: appearance */


// The appearance that will be used when the receiver is drawn onscreen, in an `NSAppearance` object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (m_ MKMapView) EffectiveAppearance() appkit.Appearance {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("effectiveAppearance"))
	return rv
}/* debug [instance_properties/getter]: effectiveAppearance */


// The appearance that will be used when the receiver is drawn onscreen, in an `NSAppearance` object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (m_ MKMapView) SetEffectiveAppearance(value appkit.Appearance) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectiveAppearance:"), value)
}/* debug [instance_properties/setter]: effectiveAppearance */


// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/ispitchenabled
func (m_ MKMapView) IsPitchEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPitchEnabled"))
	return rv
}/* debug [instance_properties/getter]: isPitchEnabled */


// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/ispitchenabled
func (m_ MKMapView) SetIsPitchEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPitchEnabled:"), value)
}/* debug [instance_properties/setter]: isPitchEnabled */


// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isrotateenabled
func (m_ MKMapView) IsRotateEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRotateEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRotateEnabled */


// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isrotateenabled
func (m_ MKMapView) SetIsRotateEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRotateEnabled:"), value)
}/* debug [instance_properties/setter]: isRotateEnabled */


// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isscrollenabled
func (m_ MKMapView) IsScrollEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isScrollEnabled"))
	return rv
}/* debug [instance_properties/getter]: isScrollEnabled */


// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isscrollenabled
func (m_ MKMapView) SetIsScrollEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsScrollEnabled:"), value)
}/* debug [instance_properties/setter]: isScrollEnabled */


// A Boolean value that indicates whether the user’s location is visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isuserlocationvisible
func (m_ MKMapView) IsUserLocationVisible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUserLocationVisible"))
	return rv
}/* debug [instance_properties/getter]: isUserLocationVisible */


// A Boolean value that indicates whether the user’s location is visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isuserlocationvisible
func (m_ MKMapView) SetIsUserLocationVisible(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUserLocationVisible:"), value)
}/* debug [instance_properties/setter]: isUserLocationVisible */


// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/iszoomenabled
func (m_ MKMapView) IsZoomEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isZoomEnabled"))
	return rv
}/* debug [instance_properties/getter]: isZoomEnabled */


// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/iszoomenabled
func (m_ MKMapView) SetIsZoomEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsZoomEnabled:"), value)
}/* debug [instance_properties/setter]: isZoomEnabled */


// The default reuse identifier for your map’s annotation views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapviewdefaultannotationviewreuseidentifier
func (m_ MKMapView) MKMapViewDefaultAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKMapViewDefaultAnnotationViewReuseIdentifier"))
	return rv
}/* debug [instance_properties/getter]: MKMapViewDefaultAnnotationViewReuseIdentifier */


// The default reuse identifier for the annotation view representing a cluster of annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapviewdefaultclusterannotationviewreuseidentifier
func (m_ MKMapView) MKMapViewDefaultClusterAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKMapViewDefaultClusterAnnotationViewReuseIdentifier"))
	return rv
}/* debug [instance_properties/getter]: MKMapViewDefaultClusterAnnotationViewReuseIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapView */


