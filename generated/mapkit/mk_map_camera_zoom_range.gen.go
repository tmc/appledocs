// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MKMapCameraZoomRange] class.
type IMKMapCameraZoomRange interface {
	objectivec.IObject
}

// A camera zoom range that limits the distances to which the user can zoom.
//
// Create a camera zoom range to limit the distance to which the user can zoom. After you create the camera zoom range, you can apply it to multiple map views. If you don’t create a camera zoom range, your map view allows the user to zoom to MapKit’s capabilities.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MKMapCameraZoomRangeClass) Alloc() MKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The zoom range to apply to the map view.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraZoomRange) CameraZoomRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}


// SetCameraZoomRange sets the value of the cameraZoomRange property.
// The zoom range to apply to the map view.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraZoomRange) SetCameraZoomRange(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}

// The minimum distance of the camera to the center of the map, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.class/mincentercoordinatedistance
func (m_ MKMapCameraZoomRange) MinCenterCoordinateDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minCenterCoordinateDistance"))
	return rv
}


// SetMinCenterCoordinateDistance sets the value of the minCenterCoordinateDistance property.
// The minimum distance of the camera to the center of the map, measured in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.class/mincentercoordinatedistance
func (m_ MKMapCameraZoomRange) SetMinCenterCoordinateDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinCenterCoordinateDistance:"), value)
}

// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraZoomRange) CameraBoundary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}


// SetCameraBoundary sets the value of the cameraBoundary property.
// The boundary of the area within which the map view’s center needs to remain.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraZoomRange) SetCameraBoundary(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}

// The maximum distance of the camera to the center of the map, measured in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.class/maxcentercoordinatedistance
func (m_ MKMapCameraZoomRange) MaxCenterCoordinateDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxCenterCoordinateDistance"))
	return rv
}


// SetMaxCenterCoordinateDistance sets the value of the maxCenterCoordinateDistance property.
// The maximum distance of the camera to the center of the map, measured in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.class/maxcentercoordinatedistance
func (m_ MKMapCameraZoomRange) SetMaxCenterCoordinateDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxCenterCoordinateDistance:"), value)
}

// A constant value used to represent the default value for zooming in or out on a map.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapcamerazoomdefault
func (m_ MKMapCameraZoomRange) MKMapCameraZoomDefault() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("MKMapCameraZoomDefault"))
	return rv
}




