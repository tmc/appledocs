// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapCamera] class.
var (
	MKMapCameraClass     _MKMapCameraClass
	MKMapCameraClassOnce sync.Once
)

func getMKMapCameraClass() _MKMapCameraClass {
	MKMapCameraClassOnce.Do(func() {
		MKMapCameraClass = _MKMapCameraClass{objc.GetClass("MKMapCamera")}
	})
	return MKMapCameraClass
}

type _MKMapCameraClass struct {
	class objc.Class
}

// An interface definition for the [MKMapCamera] class.
type IMKMapCamera interface {
	objectivec.IObject
	// properties:
	Altitude() LocationDistance /* not a class type */
	SetAltitude(value LocationDistance /* not a class type */)
	CenterCoordinate() objc.IObject /* cross-framework: LocationCoordinate2D */
	SetCenterCoordinate(value objc.IObject /* cross-framework: LocationCoordinate2D */)
	CenterCoordinateDistance() LocationDistance /* not a class type */
	SetCenterCoordinateDistance(value LocationDistance /* not a class type */)
	Heading() LocationDirection /* not a class type */
	SetHeading(value LocationDirection /* not a class type */)
	Pitch() float64
	SetPitch(value float64)
	// methods:
}

// A virtual camera for defining the appearance of the map.
//
// A camera object defines a virtual viewpoint above the map surface and affects how MapKit presents the map to the user. You use a camera object to specify the location of the camera on the map, the compass heading indicating the camera’s viewing direction, the pitch of the camera relative to the map perpendicular, and the camera’s altitude above the map. These factors create a map view with a three-dimensional perspective. After creating an instance of this class, configure it with the desired attributes and assign it to your map view. When you assign a camera to your map view, MapKit centers the map using the value in your camera object’s property, updating the map’s own region information in the process. The map also takes the camera’s pitch and altitude into account when calculating the visible region, ensuring that the region encompasses the visible content on the map.


// A virtual camera for defining the appearance of the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera
type MKMapCamera struct {
	objectivec.Object
}

// MKMapCameraFrom constructs a [MKMapCamera] from an unsafe.Pointer.
//
// A virtual camera for defining the appearance of the map.
func MKMapCameraFrom(ptr unsafe.Pointer) MKMapCamera {
	return MKMapCamera{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapCameraClass) Alloc() MKMapCamera {
	rv := objc.Send[MKMapCamera](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapCameraClass) New() MKMapCamera {
	rv := objc.Send[MKMapCamera](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapCamera) Init() MKMapCamera {
	rv := objc.Send[MKMapCamera](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapCamera) Autorelease() MKMapCamera {
	rv := objc.Send[MKMapCamera](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapCamera creates a new MKMapCamera instance.
func NewMKMapCamera() MKMapCamera {
	return getMKMapCameraClass().New()
}



// The altitude above the ground, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/altitude
func (m_ MKMapCamera) Altitude() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("altitude"))
	return rv
}


// The altitude above the ground, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/altitude
func (m_ MKMapCamera) SetAltitude(value LocationDistance /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAltitude:"), value)
}


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/centercoordinate
func (m_ MKMapCamera) CenterCoordinate() objc.IObject /* cross-framework: LocationCoordinate2D */ {
	rv := objc.Send[corelocation.LocationCoordinate2D](m_.ID, objc.Sel("centerCoordinate"))
	return rv
}


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/centercoordinate
func (m_ MKMapCamera) SetCenterCoordinate(value objc.IObject /* cross-framework: LocationCoordinate2D */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:"), value)
}


// The distance from the center point of the map to the camera, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/centercoordinatedistance
func (m_ MKMapCamera) CenterCoordinateDistance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("centerCoordinateDistance"))
	return rv
}


// The distance from the center point of the map to the camera, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/centercoordinatedistance
func (m_ MKMapCamera) SetCenterCoordinateDistance(value LocationDistance /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinateDistance:"), value)
}


// The heading of the camera (in degrees) relative to true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/heading
func (m_ MKMapCamera) Heading() LocationDirection /* not a class type */ {
	rv := objc.Send[LocationDirection](m_.ID, objc.Sel("heading"))
	return rv
}


// The heading of the camera (in degrees) relative to true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/heading
func (m_ MKMapCamera) SetHeading(value LocationDirection /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeading:"), value)
}


// The viewing angle of the camera, in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/pitch
func (m_ MKMapCamera) Pitch() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("pitch"))
	return rv
}


// The viewing angle of the camera, in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamera/pitch
func (m_ MKMapCamera) SetPitch(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitch:"), value)
}



