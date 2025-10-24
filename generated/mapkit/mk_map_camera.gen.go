// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapCamera */


/* debug [class_header]: Header for MKMapCamera */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapCamera */
// An interface definition for the [MKMapCamera] class.
type IMKMapCamera interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapCamera */
	// properties:
	Altitude() LocationDistance /* not a class type */
	SetAltitude(value LocationDistance /* not a class type */)
	CenterCoordinate() LocationCoordinate2D /* not a class type */
	SetCenterCoordinate(value LocationCoordinate2D /* not a class type */)
	CenterCoordinateDistance() LocationDistance /* not a class type */
	SetCenterCoordinateDistance(value LocationDistance /* not a class type */)
	Heading() LocationDirection /* not a class type */
	SetHeading(value LocationDirection /* not a class type */)
	Pitch() float64
	SetPitch(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapCamera */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapCamera */
// Alloc allocates a new instance without initialization.
func (mc _MKMapCameraClass) Alloc() MKMapCamera {
	rv := objc.Send[MKMapCamera](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapCamera */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapCamera */

// Returns a new camera object using the specified distance, pitch, and heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/init(lookingAtCenter:fromDistance:pitch:heading:)
func NewMKMapCameraLookingAtCenterCoordinateFromDistancePitchHeading(centerCoordinate LocationCoordinate2D /* not a class type */, distance LocationDistance /* not a class type */, pitch float64, heading LocationDirection /* not a class type */) MKMapCamera {
	rv := objc.Send[MKMapCamera](objc.ID(getMKMapCameraClass().class), objc.Sel("cameraLookingAtCenterCoordinate:fromDistance:pitch:heading:"), centerCoordinate, distance, pitch, heading)
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraLookingAtCenterCoordinateFromDistancePitchHeading */


// Returns a new camera object using the specified viewing angle information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/init(lookingAtCenter:fromEyeCoordinate:eyeAltitude:)
func NewMKMapCameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude(centerCoordinate LocationCoordinate2D /* not a class type */, eyeCoordinate LocationCoordinate2D /* not a class type */, eyeAltitude LocationDistance /* not a class type */) MKMapCamera {
	rv := objc.Send[MKMapCamera](objc.ID(getMKMapCameraClass().class), objc.Sel("cameraLookingAtCenterCoordinate:fromEyeCoordinate:eyeAltitude:"), centerCoordinate, eyeCoordinate, eyeAltitude)
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude */


// Returns a new camera object using the specified map item, view size, and pitch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/init(lookingAt:forViewSize:allowPitch:)
func NewMKMapCameraLookingAtMapItemForViewSizeAllowPitch(mapItem IMKMapItem, viewSize corefoundation.CGSize, allowPitch bool) MKMapCamera {
	rv := objc.Send[MKMapCamera](objc.ID(getMKMapCameraClass().class), objc.Sel("cameraLookingAtMapItem:forViewSize:allowPitch:"), mapItem, viewSize, allowPitch)
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraLookingAtMapItemForViewSizeAllowPitch */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapCamera */

// Returns a new camera object for you to configure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/camera
func (mc _MKMapCameraClass) Camera() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("camera"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Camera) */


// Returns a new camera object using the specified map item, view size, and pitch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/init(lookingAt:forViewSize:allowPitch:)
func (mc _MKMapCameraClass) CameraLookingAtMapItemForViewSizeAllowPitch(mapItem IMKMapItem, viewSize corefoundation.CGSize, allowPitch bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("cameraLookingAtMapItem:forViewSize:allowPitch:"), mapItem, viewSize, allowPitch)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CameraLookingAtMapItemForViewSizeAllowPitch) */


// Returns a new camera object using the specified distance, pitch, and heading information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/init(lookingAtCenter:fromDistance:pitch:heading:)
func (mc _MKMapCameraClass) CameraLookingAtCenterCoordinateFromDistancePitchHeading(centerCoordinate LocationCoordinate2D /* not a class type */, distance LocationDistance /* not a class type */, pitch float64, heading LocationDirection /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("cameraLookingAtCenterCoordinate:fromDistance:pitch:heading:"), centerCoordinate, distance, pitch, heading)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CameraLookingAtCenterCoordinateFromDistancePitchHeading) */


// Returns a new camera object using the specified viewing angle information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/init(lookingAtCenter:fromEyeCoordinate:eyeAltitude:)
func (mc _MKMapCameraClass) CameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude(centerCoordinate LocationCoordinate2D /* not a class type */, eyeCoordinate LocationCoordinate2D /* not a class type */, eyeAltitude LocationDistance /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("cameraLookingAtCenterCoordinate:fromEyeCoordinate:eyeAltitude:"), centerCoordinate, eyeCoordinate, eyeAltitude)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapCamera */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapCamera */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapCamera */

// The altitude above the ground, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/altitude
func (m_ MKMapCamera) Altitude() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("altitude"))
	return rv
}/* debug [instance_properties/getter]: altitude */


// The altitude above the ground, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/altitude
func (m_ MKMapCamera) SetAltitude(value LocationDistance /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAltitude:"), value)
}/* debug [instance_properties/setter]: altitude */


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/centerCoordinate
func (m_ MKMapCamera) CenterCoordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("centerCoordinate"))
	return rv
}/* debug [instance_properties/getter]: centerCoordinate */


// The map coordinate at the center of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/centerCoordinate
func (m_ MKMapCamera) SetCenterCoordinate(value LocationCoordinate2D /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinate:"), value)
}/* debug [instance_properties/setter]: centerCoordinate */


// The distance from the center point of the map to the camera, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/centerCoordinateDistance
func (m_ MKMapCamera) CenterCoordinateDistance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("centerCoordinateDistance"))
	return rv
}/* debug [instance_properties/getter]: centerCoordinateDistance */


// The distance from the center point of the map to the camera, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/centerCoordinateDistance
func (m_ MKMapCamera) SetCenterCoordinateDistance(value LocationDistance /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCenterCoordinateDistance:"), value)
}/* debug [instance_properties/setter]: centerCoordinateDistance */


// The heading of the camera (in degrees) relative to true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/heading
func (m_ MKMapCamera) Heading() LocationDirection /* not a class type */ {
	rv := objc.Send[LocationDirection](m_.ID, objc.Sel("heading"))
	return rv
}/* debug [instance_properties/getter]: heading */


// The heading of the camera (in degrees) relative to true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/heading
func (m_ MKMapCamera) SetHeading(value LocationDirection /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeading:"), value)
}/* debug [instance_properties/setter]: heading */


// The viewing angle of the camera, in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/pitch
func (m_ MKMapCamera) Pitch() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("pitch"))
	return rv
}/* debug [instance_properties/getter]: pitch */


// The viewing angle of the camera, in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapCamera/pitch
func (m_ MKMapCamera) SetPitch(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitch:"), value)
}/* debug [instance_properties/setter]: pitch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapCamera */


