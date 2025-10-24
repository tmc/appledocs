// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ShowsPointsOfInterest() bool
	SetShowsPointsOfInterest(value bool)
	ShowsUserLocation() bool
	SetShowsUserLocation(value bool)
	Appearance() objc.IObject /* cross-framework: Appearance */
	SetAppearance(value objc.IObject /* cross-framework: Appearance */)
	EffectiveAppearance() objc.IObject /* cross-framework: Appearance */
	SetEffectiveAppearance(value objc.IObject /* cross-framework: Appearance */)
	AnnotationVisibleRect() objc.IObject /* cross-framework: Rect */
	SetAnnotationVisibleRect(value objc.IObject /* cross-framework: Rect */)
	Annotations() unsafe.Pointer
	SetAnnotations(value unsafe.Pointer)
	Camera() IMKMapCamera
	SetCamera(value IMKMapCamera)
	CameraBoundary() objc.IObject /* cross-framework: MKMapCameraBoundary */
	SetCameraBoundary(value objc.IObject /* cross-framework: MKMapCameraBoundary */)
	CameraZoomRange() objc.IObject /* cross-framework: MKMapCameraZoomRange */
	SetCameraZoomRange(value objc.IObject /* cross-framework: MKMapCameraZoomRange */)
	CenterCoordinate() objc.IObject /* cross-framework: LocationCoordinate2D */
	SetCenterCoordinate(value objc.IObject /* cross-framework: LocationCoordinate2D */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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
	MapType() unsafe.Pointer
	SetMapType(value unsafe.Pointer)
	Overlays() unsafe.Pointer
	SetOverlays(value unsafe.Pointer)
	PitchButtonVisibility() unsafe.Pointer
	SetPitchButtonVisibility(value unsafe.Pointer)
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	PreferredConfiguration() objc.IObject /* cross-framework: MKMapConfiguration */
	SetPreferredConfiguration(value objc.IObject /* cross-framework: MKMapConfiguration */)
	Region() unsafe.Pointer
	SetRegion(value unsafe.Pointer)
	SelectableMapFeatures() unsafe.Pointer
	SetSelectableMapFeatures(value unsafe.Pointer)
	SelectedAnnotations() unsafe.Pointer
	SetSelectedAnnotations(value unsafe.Pointer)
	ShowsBuildings() bool
	SetShowsBuildings(value bool)
	ShowsCompass() bool
	SetShowsCompass(value bool)
	ShowsPitchControl() bool
	SetShowsPitchControl(value bool)
	ShowsScale() bool
	SetShowsScale(value bool)
	ShowsTraffic() bool
	SetShowsTraffic(value bool)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
	ShowsZoomControls() bool
	SetShowsZoomControls(value bool)
	UserLocation() IMKUserLocation
	SetUserLocation(value IMKUserLocation)
	UserTrackingMode() unsafe.Pointer
	SetUserTrackingMode(value unsafe.Pointer)
	VisibleMapRect() unsafe.Pointer
	SetVisibleMapRect(value unsafe.Pointer)
	MKMapViewDefaultAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */
	MKMapViewDefaultClusterAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An embeddable map interface, similar to the one that the Maps app provides.
//
// Use this class as-is to display map information and to manipulate the map contents from your app. The map view supports several display styles, including the that provides rich 2D and 3D presentations, an that provides a hybrid satellite map presentation, and that provides an imagery-based map presentation. Each of these map configurations support customization properties that refine specific elements of the map’s presentation. You can center the map on specific coordinates, specify the size of the area you want to display, and annotate the map with custom information. When you initialize a map view, you specify the initial region for that map to display by setting the property of the map. MapKit defines a region by a center point and a horizontal and vertical distance, referred to as the . The defines how much of the map is visible, and is also how you set the zoom level. For example, specifying a large span results in the user seeing a wide geographical area at a low zoom level, whereas specifying a small span results in a more narrow geographical area and a higher zoom level. In addition to setting the span programmatically, the class supports many standard interactions for changing the position and zoom level of the map. In particular, map views support flick and pinch gestures for scrolling around the map and zooming in and out. The map view enables support for these gestures by default. You can enable and disable them using the and properties. You can also use projected map coordinates instead of regions to specify some values. When you project the curved surface of the globe onto a flat surface, you get a two-dimensional version of a map where longitude lines appear to be parallel. To specify locations and distances, you use the , , and data types. Don’t subclass the class itself. You can get information about the map view’s behavior by providing a delegate object. The map view calls the methods of your custom delegate to let it know about changes in the map status and to coordinate the display of custom annotations. The delegate object can be any object in your app as long as it conforms to the protocol. For more information about implementing the delegate object, see . In macOS 10.14 and later, you can apply a light or dark appearance to your maps by modifying the property of your map view (or one of its ancestor views). Even if you specify a custom appearance, users can use the Maps app to force all maps to adopt a light appearance. Use the map view’s property to determine the actual appearance of your map. For information about how to set view appearances, see .


// An embeddable map interface, similar to the one that the Maps app provides.
//
// [Full Topic]
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



// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPointsOfInterest
func (m_ MKMapView) ShowsPointsOfInterest() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPointsOfInterest"))
	return rv
}


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsPointsOfInterest
func (m_ MKMapView) SetShowsPointsOfInterest(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPointsOfInterest:"), value)
}


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserLocation
func (m_ MKMapView) ShowsUserLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserLocation"))
	return rv
}


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/showsUserLocation
func (m_ MKMapView) SetShowsUserLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserLocation:"), value)
}


// The appearance of the receiver, in an `NSAppearance` object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (m_ MKMapView) Appearance() objc.IObject /* cross-framework: Appearance */ {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("appearance"))
	return rv
}


// The appearance of the receiver, in an `NSAppearance` object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (m_ MKMapView) SetAppearance(value objc.IObject /* cross-framework: Appearance */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAppearance:"), value)
}


// The appearance that will be used when the receiver is drawn onscreen, in an `NSAppearance` object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (m_ MKMapView) EffectiveAppearance() objc.IObject /* cross-framework: Appearance */ {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("effectiveAppearance"))
	return rv
}


// The appearance that will be used when the receiver is drawn onscreen, in an `NSAppearance` object. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (m_ MKMapView) SetEffectiveAppearance(value objc.IObject /* cross-framework: Appearance */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectiveAppearance:"), value)
}


// The visible rectangle where the map is displaying annotation views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/annotationvisiblerect
func (m_ MKMapView) AnnotationVisibleRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("annotationVisibleRect"))
	return rv
}


// The visible rectangle where the map is displaying annotation views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/annotationvisiblerect
func (m_ MKMapView) SetAnnotationVisibleRect(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnotationVisibleRect:"), value)
}


// The annotations associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/annotations
func (m_ MKMapView) Annotations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("annotations"))
	return rv
}


// The annotations associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/annotations
func (m_ MKMapView) SetAnnotations(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnnotations:"), value)
}


// The camera to use for determining the appearance of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camera
func (m_ MKMapView) Camera() IMKMapCamera {
	rv := objc.Send[MKMapCamera](m_.ID, objc.Sel("camera"))
	return rv
}


// The camera to use for determining the appearance of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camera
func (m_ MKMapView) SetCamera(value IMKMapCamera) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:"), value)
}


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapView) CameraBoundary() objc.IObject /* cross-framework: MKMapCameraBoundary */ {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapView) SetCameraBoundary(value objc.IObject /* cross-framework: MKMapCameraBoundary */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapView) CameraZoomRange() objc.IObject /* cross-framework: MKMapCameraZoomRange */ {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapView) SetCameraZoomRange(value objc.IObject /* cross-framework: MKMapCameraZoomRange */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/centercoordinate
func (m_ MKMapView) CenterCoordinate() objc.IObject /* cross-framework: LocationCoordinate2D */ {
	rv := objc.Send[corelocation.LocationCoordinate2D](m_.ID, objc.Sel("centerCoordinate"))
	return rv
}


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/centercoordinate
func (m_ MKMapView) SetCenterCoordinate(value objc.IObject /* cross-framework: LocationCoordinate2D */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:"), value)
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/delegate
func (m_ MKMapView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/delegate
func (m_ MKMapView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/ispitchenabled
func (m_ MKMapView) IsPitchEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPitchEnabled"))
	return rv
}


// A Boolean value that indicates whether the map uses the camera’s pitch information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/ispitchenabled
func (m_ MKMapView) SetIsPitchEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPitchEnabled:"), value)
}


// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isrotateenabled
func (m_ MKMapView) IsRotateEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRotateEnabled"))
	return rv
}


// A Boolean value that indicates whether the map uses the camera’s heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isrotateenabled
func (m_ MKMapView) SetIsRotateEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRotateEnabled:"), value)
}


// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isscrollenabled
func (m_ MKMapView) IsScrollEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isScrollEnabled"))
	return rv
}


// A Boolean value that determines whether the user may scroll around the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isscrollenabled
func (m_ MKMapView) SetIsScrollEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsScrollEnabled:"), value)
}


// A Boolean value that indicates whether the user’s location is visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isuserlocationvisible
func (m_ MKMapView) IsUserLocationVisible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUserLocationVisible"))
	return rv
}


// A Boolean value that indicates whether the user’s location is visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/isuserlocationvisible
func (m_ MKMapView) SetIsUserLocationVisible(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUserLocationVisible:"), value)
}


// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/iszoomenabled
func (m_ MKMapView) IsZoomEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isZoomEnabled"))
	return rv
}


// A Boolean value that determines whether the user may use pinch gestures to zoom in and out of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/iszoomenabled
func (m_ MKMapView) SetIsZoomEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsZoomEnabled:"), value)
}


// The type of data the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/maptype
func (m_ MKMapView) MapType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapType"))
	return rv
}


// The type of data the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/maptype
func (m_ MKMapView) SetMapType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapType:"), value)
}


// The overlay objects associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/overlays
func (m_ MKMapView) Overlays() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("overlays"))
	return rv
}


// The overlay objects associated with the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/overlays
func (m_ MKMapView) SetOverlays(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOverlays:"), value)
}


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKMapView) PitchButtonVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKMapView) SetPitchButtonVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}


// The filter to use for determining the points of interest that appear on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pointofinterestfilter
func (m_ MKMapView) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// The filter to use for determining the points of interest that appear on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pointofinterestfilter
func (m_ MKMapView) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKMapView) PreferredConfiguration() objc.IObject /* cross-framework: MKMapConfiguration */ {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKMapView) SetPreferredConfiguration(value objc.IObject /* cross-framework: MKMapConfiguration */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}


// The area the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/region
func (m_ MKMapView) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// The area the map view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/region
func (m_ MKMapView) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}


// The property that describes which selectable features the map responds to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/selectablemapfeatures
func (m_ MKMapView) SelectableMapFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("selectableMapFeatures"))
	return rv
}


// The property that describes which selectable features the map responds to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/selectablemapfeatures
func (m_ MKMapView) SetSelectableMapFeatures(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectableMapFeatures:"), value)
}


// The selected annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/selectedannotations
func (m_ MKMapView) SelectedAnnotations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("selectedAnnotations"))
	return rv
}


// The selected annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/selectedannotations
func (m_ MKMapView) SetSelectedAnnotations(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedAnnotations:"), value)
}


// A Boolean value that indicates whether the map displays extruded building information on supported map types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsbuildings
func (m_ MKMapView) ShowsBuildings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsBuildings"))
	return rv
}


// A Boolean value that indicates whether the map displays extruded building information on supported map types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsbuildings
func (m_ MKMapView) SetShowsBuildings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsBuildings:"), value)
}


// A Boolean value that indicates whether the map displays a compass control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showscompass
func (m_ MKMapView) ShowsCompass() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsCompass"))
	return rv
}


// A Boolean value that indicates whether the map displays a compass control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showscompass
func (m_ MKMapView) SetShowsCompass(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsCompass:"), value)
}


// A Boolean value that indicates whether the map displays the pitch control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showspitchcontrol
func (m_ MKMapView) ShowsPitchControl() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPitchControl"))
	return rv
}


// A Boolean value that indicates whether the map displays the pitch control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showspitchcontrol
func (m_ MKMapView) SetShowsPitchControl(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPitchControl:"), value)
}


// A Boolean value that indicates whether the map shows scale information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsscale
func (m_ MKMapView) ShowsScale() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsScale"))
	return rv
}


// A Boolean value that indicates whether the map shows scale information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsscale
func (m_ MKMapView) SetShowsScale(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsScale:"), value)
}


// A Boolean value that indicates whether the map displays traffic information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showstraffic
func (m_ MKMapView) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}


// A Boolean value that indicates whether the map displays traffic information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showstraffic
func (m_ MKMapView) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKMapView) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKMapView) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}


// A Boolean value that indicates whether the map displays zoom controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showszoomcontrols
func (m_ MKMapView) ShowsZoomControls() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsZoomControls"))
	return rv
}


// A Boolean value that indicates whether the map displays zoom controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showszoomcontrols
func (m_ MKMapView) SetShowsZoomControls(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsZoomControls:"), value)
}


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKMapView) UserLocation() IMKUserLocation {
	rv := objc.Send[MKUserLocation](m_.ID, objc.Sel("userLocation"))
	return rv
}


// The annotation object that represents the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/userlocation
func (m_ MKMapView) SetUserLocation(value IMKUserLocation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserLocation:"), value)
}


// The mode to use for tracking the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/usertrackingmode
func (m_ MKMapView) UserTrackingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("userTrackingMode"))
	return rv
}


// The mode to use for tracking the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/usertrackingmode
func (m_ MKMapView) SetUserTrackingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserTrackingMode:"), value)
}


// The area visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/visiblemaprect
func (m_ MKMapView) VisibleMapRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("visibleMapRect"))
	return rv
}


// The area visible in the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/visiblemaprect
func (m_ MKMapView) SetVisibleMapRect(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisibleMapRect:"), value)
}


// The default reuse identifier for your map’s annotation views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapviewdefaultannotationviewreuseidentifier
func (m_ MKMapView) MKMapViewDefaultAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKMapViewDefaultAnnotationViewReuseIdentifier"))
	return rv
}


// The default reuse identifier for the annotation view representing a cluster of annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapviewdefaultclusterannotationviewreuseidentifier
func (m_ MKMapView) MKMapViewDefaultClusterAnnotationViewReuseIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MKMapViewDefaultClusterAnnotationViewReuseIdentifier"))
	return rv
}



