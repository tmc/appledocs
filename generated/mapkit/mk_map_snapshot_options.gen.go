// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKMapSnapshotOptions] class.
type IMKMapSnapshotOptions interface {
	objectivec.IObject
	Appearance() appkit.Appearance
	SetAppearance(value appkit.IAppearance)
	Camera() unsafe.Pointer
	SetCamera(value unsafe.Pointer)
	MapRect() unsafe.Pointer
	SetMapRect(value unsafe.Pointer)
	MapType() MKMapType
	SetMapType(value MKMapType)
	PointOfInterestFilter() MKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	PreferredConfiguration() MKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	Region() unsafe.Pointer
	SetRegion(value unsafe.Pointer)
	Scale() float64
	SetScale(value float64)
	ShowsBuildings() bool
	SetShowsBuildings(value bool)
	ShowsPointsOfInterest() bool
	SetShowsPointsOfInterest(value bool)
	Size() foundation.Size
	SetSize(value foundation.Size)
	TraitCollection() unsafe.Pointer
	SetTraitCollection(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (mc _MKMapSnapshotOptionsClass) Alloc() MKMapSnapshotOptions {
	rv := objc.Send[MKMapSnapshotOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The visual style (light or dark) to apply to the map when rendering the snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/appearance
func (m_ MKMapSnapshotOptions) Appearance() appkit.Appearance {
	rv := objc.Send[appkit.Appearance](m_.ID, objc.Sel("appearance"))
	return rv
}


// The visual style (light or dark) to apply to the map when rendering the snapshot image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/appearance
func (m_ MKMapSnapshotOptions) SetAppearance(value appkit.IAppearance) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAppearance:"), value)
}


// The camera to use when taking the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/camera
func (m_ MKMapSnapshotOptions) Camera() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("camera"))
	return rv
}


// The camera to use when taking the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/camera
func (m_ MKMapSnapshotOptions) SetCamera(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCamera:"), value)
}


// The map rectangle that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapRect
func (m_ MKMapSnapshotOptions) MapRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapRect"))
	return rv
}


// The map rectangle that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapRect
func (m_ MKMapSnapshotOptions) SetMapRect(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapRect:"), value)
}


// The map’s visual style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapType
func (m_ MKMapSnapshotOptions) MapType() MKMapType {
	rv := objc.Send[MKMapType](m_.ID, objc.Sel("mapType"))
	return rv
}


// The map’s visual style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/mapType
func (m_ MKMapSnapshotOptions) SetMapType(value MKMapType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapType:"), value)
}


// The filter to use for determining the points of interest that appear in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/pointOfInterestFilter
func (m_ MKMapSnapshotOptions) PointOfInterestFilter() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// The filter to use for determining the points of interest that appear in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/pointOfInterestFilter
func (m_ MKMapSnapshotOptions) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}


// The map configuration style to use for snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/preferredConfiguration
func (m_ MKMapSnapshotOptions) PreferredConfiguration() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// The map configuration style to use for snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/preferredConfiguration
func (m_ MKMapSnapshotOptions) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}


// The area of the map that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/region
func (m_ MKMapSnapshotOptions) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// The area of the map that you want to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/region
func (m_ MKMapSnapshotOptions) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}


// The scale factor to use when creating the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/scale
func (m_ MKMapSnapshotOptions) Scale() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("scale"))
	return rv
}


// The scale factor to use when creating the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/scale
func (m_ MKMapSnapshotOptions) SetScale(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScale:"), value)
}


// A Boolean that indicates whether the map displays extruded building information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsBuildings
func (m_ MKMapSnapshotOptions) ShowsBuildings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsBuildings"))
	return rv
}


// A Boolean that indicates whether the map displays extruded building information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsBuildings
func (m_ MKMapSnapshotOptions) SetShowsBuildings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsBuildings:"), value)
}


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsPointsOfInterest
func (m_ MKMapSnapshotOptions) ShowsPointsOfInterest() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsPointsOfInterest"))
	return rv
}


// A Boolean value that indicates whether the map displays point-of-interest information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/showsPointsOfInterest
func (m_ MKMapSnapshotOptions) SetShowsPointsOfInterest(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsPointsOfInterest:"), value)
}


// The size of the image that you want to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/size
func (m_ MKMapSnapshotOptions) Size() foundation.Size {
	rv := objc.Send[foundation.Size](m_.ID, objc.Sel("size"))
	return rv
}


// The size of the image that you want to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/size
func (m_ MKMapSnapshotOptions) SetSize(value foundation.Size) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}


// Traits that determine the appearance of the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/traitCollection
func (m_ MKMapSnapshotOptions) TraitCollection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("traitCollection"))
	return rv
}


// Traits that determine the appearance of the map snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSnapshotter/Options/traitCollection
func (m_ MKMapSnapshotOptions) SetTraitCollection(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTraitCollection:"), value)
}



