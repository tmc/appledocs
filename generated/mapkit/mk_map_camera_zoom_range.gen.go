// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapCameraZoomRange */


/* debug [class_header]: Header for MKMapCameraZoomRange */
// The class instance for the [MKMapCameraZoomRange] class.
var (
	MKMapCameraZoomRangeClass     _MKMapCameraZoomRangeClass
	MKMapCameraZoomRangeClassOnce sync.Once
)

func getMKMapCameraZoomRangeClass() _MKMapCameraZoomRangeClass {
	MKMapCameraZoomRangeClassOnce.Do(func() {
		MKMapCameraZoomRangeClass = _MKMapCameraZoomRangeClass{objc.GetClass("MKMapCameraZoomRange")}
	})
	return MKMapCameraZoomRangeClass
}

type _MKMapCameraZoomRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapCameraZoomRange */
// An interface definition for the [MKMapCameraZoomRange] class.
type IMKMapCameraZoomRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapCameraZoomRange */
	// properties:
	MaxCenterCoordinateDistance() LocationDistance /* not a class type */
	MinCenterCoordinateDistance() LocationDistance /* not a class type */
	MKMapCameraZoomDefault() LocationDistance /* not a class type */
	CameraBoundary() IMKMapCameraBoundary
	SetCameraBoundary(value IMKMapCameraBoundary)
	CameraZoomRange() IMKMapCameraZoomRange
	SetCameraZoomRange(value IMKMapCameraZoomRange)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapCameraZoomRange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapCameraZoomRange */
// Alloc allocates a new instance without initialization.
func (mc _MKMapCameraZoomRangeClass) Alloc() MKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKMapCameraZoomRangeClass) New() MKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapCameraZoomRange) Init() MKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapCameraZoomRange) Autorelease() MKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapCameraZoomRange creates a new MKMapCameraZoomRange instance.
func NewMKMapCameraZoomRange() MKMapCameraZoomRange {
	return getMKMapCameraZoomRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapCameraZoomRange */
// A camera zoom range that limits the distances to which the user can zoom.
//
// Create a camera zoom range to limit the distance to which the user can zoom. After you create the camera zoom range, you can apply it to multiple map views. If you don’t create a camera zoom range, your map view allows the user to zoom to MapKit’s capabilities.


// A camera zoom range that limits the distances to which the user can zoom.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraZoomRange-swift.class
type MKMapCameraZoomRange struct {
	objectivec.Object
}

// MKMapCameraZoomRangeFrom constructs a [MKMapCameraZoomRange] from an unsafe.Pointer.
//
// A camera zoom range that limits the distances to which the user can zoom.
func MKMapCameraZoomRangeFrom(ptr unsafe.Pointer) MKMapCameraZoomRange {
	return MKMapCameraZoomRange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapCameraZoomRange */

// Create a camera zoom range by specifying the maximum distance from your map view’s center coordinate, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraZoomRange-swift.class/init(maxCenterCoordinateDistance:)
func NewMKMapCameraZoomRangeWithMaxCenterCoordinateDistance(maxDistance LocationDistance /* not a class type */) MKMapCameraZoomRange {
	instance := getMKMapCameraZoomRangeClass().Alloc()
	rv := objc.Send[MKMapCameraZoomRange](instance.ID, objc.Sel("initWithMaxCenterCoordinateDistance:"), maxDistance)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraZoomRangeWithMaxCenterCoordinateDistance */


// Create a camera zoom range by specifying the minimum distance from your map view’s center coordinate, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraZoomRange-swift.class/init(minCenterCoordinateDistance:)
func NewMKMapCameraZoomRangeWithMinCenterCoordinateDistance(minDistance LocationDistance /* not a class type */) MKMapCameraZoomRange {
	instance := getMKMapCameraZoomRangeClass().Alloc()
	rv := objc.Send[MKMapCameraZoomRange](instance.ID, objc.Sel("initWithMinCenterCoordinateDistance:"), minDistance)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraZoomRangeWithMinCenterCoordinateDistance */


// Create a camera zoom range by specifying a minimum and maximum distance from your map view’s center coordinates, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraZoomRange-swift.class/init(minCenterCoordinateDistance:maxCenterCoordinateDistance:)
func NewMKMapCameraZoomRangeWithMinCenterCoordinateDistanceMaxCenterCoordinateDistance(minDistance LocationDistance /* not a class type */, maxDistance LocationDistance /* not a class type */) MKMapCameraZoomRange {
	instance := getMKMapCameraZoomRangeClass().Alloc()
	rv := objc.Send[MKMapCameraZoomRange](instance.ID, objc.Sel("initWithMinCenterCoordinateDistance:maxCenterCoordinateDistance:"), minDistance, maxDistance)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKMapCameraZoomRangeWithMinCenterCoordinateDistanceMaxCenterCoordinateDistance */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapCameraZoomRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapCameraZoomRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapCameraZoomRange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapCameraZoomRange */

// The maximum distance of the camera to the center of the map, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraZoomRange-swift.class/maxCenterCoordinateDistance
func (m_ MKMapCameraZoomRange) MaxCenterCoordinateDistance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("maxCenterCoordinateDistance"))
	return rv
}/* debug [instance_properties/getter]: maxCenterCoordinateDistance */


// The minimum distance of the camera to the center of the map, measured in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraZoomRange-swift.class/minCenterCoordinateDistance
func (m_ MKMapCameraZoomRange) MinCenterCoordinateDistance() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("minCenterCoordinateDistance"))
	return rv
}/* debug [instance_properties/getter]: minCenterCoordinateDistance */


// A constant value used to represent the default value for zooming in or out on a map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamerazoomdefault
func (m_ MKMapCameraZoomRange) MKMapCameraZoomDefault() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("MKMapCameraZoomDefault"))
	return rv
}/* debug [instance_properties/getter]: MKMapCameraZoomDefault */


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraZoomRange) CameraBoundary() IMKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}/* debug [instance_properties/getter]: cameraBoundary */


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraZoomRange) SetCameraBoundary(value IMKMapCameraBoundary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}/* debug [instance_properties/setter]: cameraBoundary */


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraZoomRange) CameraZoomRange() IMKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}/* debug [instance_properties/getter]: cameraZoomRange */


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraZoomRange) SetCameraZoomRange(value IMKMapCameraZoomRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}/* debug [instance_properties/setter]: cameraZoomRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapCameraZoomRange */


