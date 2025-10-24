// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapSnapshotOptions */


/* debug [class_header]: Header for MKMapSnapshotOptions */
// The class instance for the [MKMapSnapshotOptions] class.
var (
	MKMapSnapshotOptionsClass     _MKMapSnapshotOptionsClass
	MKMapSnapshotOptionsClassOnce sync.Once
)

func getMKMapSnapshotOptionsClass() _MKMapSnapshotOptionsClass {
	MKMapSnapshotOptionsClassOnce.Do(func() {
		MKMapSnapshotOptionsClass = _MKMapSnapshotOptionsClass{objc.GetClass("MKMapSnapshotOptions")}
	})
	return MKMapSnapshotOptionsClass
}

type _MKMapSnapshotOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapSnapshotOptions */
// An interface definition for the [MKMapSnapshotOptions] class.
type IMKMapSnapshotOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapSnapshotOptions */
	// properties:
	Appearance() appkit.Appearance
	SetAppearance(value appkit.Appearance)
	Camera() IMKMapCamera
	SetCamera(value IMKMapCamera)
	MapRect() objc.IObject /* cross-framework: MKMapRect */
	SetMapRect(value objc.IObject /* cross-framework: MKMapRect */)
	MapType() MKMapType
	SetMapType(value MKMapType)
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	PreferredConfiguration() IMKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	Region() objc.IObject /* cross-framework: MKCoordinateRegion */
	SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */)
	ShowsBuildings() bool
	SetShowsBuildings(value bool)
	ShowsPointsOfInterest() bool
	SetShowsPointsOfInterest(value bool)
	Size() Size /* not a class type */
	SetSize(value Size /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapSnapshotOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapSnapshotOptions */
// Alloc allocates a new instance without initialization.
func (mc _MKMapSnapshotOptionsClass) Alloc() MKMapSnapshotOptions {
	rv := objc.Send[MKMapSnapshotOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapSnapshotOptionsClass) New() MKMapSnapshotOptions {
	rv := objc.Send[MKMapSnapshotOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapSnapshotOptions) Init() MKMapSnapshotOptions {
	rv := objc.Send[MKMapSnapshotOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapSnapshotOptions) Autorelease() MKMapSnapshotOptions {
	rv := objc.Send[MKMapSnapshotOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapSnapshotOptions creates a new MKMapSnapshotOptions instance.
func NewMKMapSnapshotOptions() MKMapSnapshotOptions {
	return getMKMapSnapshotOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapSnapshotOptions */
// The options the snapshotter initializer uses to create a snapshotter to capture map-based imagery.
//
// After creating and configuring an instance of this class, you pass that instance to an object. The snapshotter uses the configuration options to determine which portion of the map to capture, the viewing angle to use for the camera, and the map’s overall appearance. In macOS 10.14 and later, you can apply a light or dark appearance to your map snapshots by modifying the property of your snapshot options. Even if you specify a custom appearance, users can use the Maps app to force all maps to adopt a light appearance.


// The options the snapshotter initializer uses to create a snapshotter to capture map-based imagery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options
type MKMapSnapshotOptions struct {
	objectivec.Object
}

// MKMapSnapshotOptionsFrom constructs a [MKMapSnapshotOptions] from an unsafe.Pointer.
//
// The options the snapshotter initializer uses to create a snapshotter to capture map-based imagery.
func MKMapSnapshotOptionsFrom(ptr unsafe.Pointer) MKMapSnapshotOptions {
	return MKMapSnapshotOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapSnapshotOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapSnapshotOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapSnapshotOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapSnapshotOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapSnapshotOptions */

// The visual style (light or dark) to apply to the map when rendering the snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/appearance
func (m_ MKMapSnapshotOptions) Appearance() appkit.Appearance {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("appearance"))
	return rv
}/* debug [instance_properties/getter]: appearance */


// The visual style (light or dark) to apply to the map when rendering the snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/appearance
func (m_ MKMapSnapshotOptions) SetAppearance(value appkit.Appearance) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAppearance:"), value)
}/* debug [instance_properties/setter]: appearance */


// The camera to use when taking the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/camera
func (m_ MKMapSnapshotOptions) Camera() IMKMapCamera {
	rv := objc.Send[MKMapCamera](m_.ID, objc.Sel("camera"))
	return rv
}/* debug [instance_properties/getter]: camera */


// The camera to use when taking the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/camera
func (m_ MKMapSnapshotOptions) SetCamera(value IMKMapCamera) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:"), value)
}/* debug [instance_properties/setter]: camera */


// The map rectangle that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapRect
func (m_ MKMapSnapshotOptions) MapRect() objc.IObject /* cross-framework: MKMapRect */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("mapRect"))
	return rv
}/* debug [instance_properties/getter]: mapRect */


// The map rectangle that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapRect
func (m_ MKMapSnapshotOptions) SetMapRect(value objc.IObject /* cross-framework: MKMapRect */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapRect:"), value)
}/* debug [instance_properties/setter]: mapRect */


// The map’s visual style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapType
func (m_ MKMapSnapshotOptions) MapType() MKMapType {
	rv := objc.Send[MKMapType](m_.ID, objc.Sel("mapType"))
	return rv
}/* debug [instance_properties/getter]: mapType */


// The map’s visual style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapType
func (m_ MKMapSnapshotOptions) SetMapType(value MKMapType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapType:"), value)
}/* debug [instance_properties/setter]: mapType */


// The filter to use for determining the points of interest that appear in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/pointOfInterestFilter
func (m_ MKMapSnapshotOptions) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// The filter to use for determining the points of interest that appear in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/pointOfInterestFilter
func (m_ MKMapSnapshotOptions) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// The map configuration style to use for snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/preferredConfiguration
func (m_ MKMapSnapshotOptions) PreferredConfiguration() IMKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}/* debug [instance_properties/getter]: preferredConfiguration */


// The map configuration style to use for snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/preferredConfiguration
func (m_ MKMapSnapshotOptions) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}/* debug [instance_properties/setter]: preferredConfiguration */


// The area of the map that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/region
func (m_ MKMapSnapshotOptions) Region() objc.IObject /* cross-framework: MKCoordinateRegion */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("region"))
	return rv
}/* debug [instance_properties/getter]: region */


// The area of the map that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/region
func (m_ MKMapSnapshotOptions) SetRegion(value objc.IObject /* cross-framework: MKCoordinateRegion */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}/* debug [instance_properties/setter]: region */


// A Boolean that indicates whether the map displays extruded building information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsBuildings
func (m_ MKMapSnapshotOptions) ShowsBuildings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsBuildings"))
	return rv
}/* debug [instance_properties/getter]: showsBuildings */


// A Boolean that indicates whether the map displays extruded building information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsBuildings
func (m_ MKMapSnapshotOptions) SetShowsBuildings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsBuildings:"), value)
}/* debug [instance_properties/setter]: showsBuildings */


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsPointsOfInterest
func (m_ MKMapSnapshotOptions) ShowsPointsOfInterest() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPointsOfInterest"))
	return rv
}/* debug [instance_properties/getter]: showsPointsOfInterest */


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsPointsOfInterest
func (m_ MKMapSnapshotOptions) SetShowsPointsOfInterest(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPointsOfInterest:"), value)
}/* debug [instance_properties/setter]: showsPointsOfInterest */


// The size of the image that you want to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/size
func (m_ MKMapSnapshotOptions) Size() Size /* not a class type */ {
	rv := objc.Send[Size](m_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The size of the image that you want to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/size
func (m_ MKMapSnapshotOptions) SetSize(value Size /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapSnapshotOptions */


