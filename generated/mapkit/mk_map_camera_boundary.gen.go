// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapCameraBoundary] class.
var (
	MKMapCameraBoundaryClass     _MKMapCameraBoundaryClass
	MKMapCameraBoundaryClassOnce sync.Once
)

func getMKMapCameraBoundaryClass() _MKMapCameraBoundaryClass {
	MKMapCameraBoundaryClassOnce.Do(func() {
		MKMapCameraBoundaryClass = _MKMapCameraBoundaryClass{objc.GetClass("MKMapCameraBoundary")}
	})
	return MKMapCameraBoundaryClass
}

type _MKMapCameraBoundaryClass struct {
	class objc.Class
}

// An interface definition for the [MKMapCameraBoundary] class.
type IMKMapCameraBoundary interface {
	objectivec.IObject
	MapRect() unsafe.Pointer
	SetMapRect(value unsafe.Pointer)
	Region() unsafe.Pointer
	SetRegion(value unsafe.Pointer)
	CameraBoundary() MKMapCameraBoundary
	SetCameraBoundary(value IMKMapCameraBoundary)
	CameraZoomRange() MKMapCameraZoomRange
	SetCameraZoomRange(value IMKMapCameraZoomRange)
}

// A boundary of an area within which the map’s center needs to remain.
//
// The constraints of the camera boundary restrict the center point of your map.


// A boundary of an area within which the map’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/CameraBoundary-swift.class
type MKMapCameraBoundary struct {
	objectivec.Object
}

// MKMapCameraBoundaryFrom constructs a [MKMapCameraBoundary] from an unsafe.Pointer.
//
// A boundary of an area within which the map’s center needs to remain.
func MKMapCameraBoundaryFrom(ptr unsafe.Pointer) MKMapCameraBoundary {
	return MKMapCameraBoundary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapCameraBoundaryClass) Alloc() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapCameraBoundaryClass) New() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapCameraBoundary) Init() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapCameraBoundary) Autorelease() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapCameraBoundary creates a new MKMapCameraBoundary instance.
func NewMKMapCameraBoundary() MKMapCameraBoundary {
	return getMKMapCameraBoundaryClass().New()
}



// The map rectangle that describes the camera boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.class/maprect
func (m_ MKMapCameraBoundary) MapRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapRect"))
	return rv
}


// The map rectangle that describes the camera boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.class/maprect
func (m_ MKMapCameraBoundary) SetMapRect(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapRect:"), value)
}


// The coordinate region that describes the camera boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.class/region
func (m_ MKMapCameraBoundary) Region() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("region"))
	return rv
}


// The coordinate region that describes the camera boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.class/region
func (m_ MKMapCameraBoundary) SetRegion(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRegion:"), value)
}


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraBoundary) CameraBoundary() MKMapCameraBoundary {
	rv := objc.Send[MKMapCameraBoundary](m_.ID, objc.Sel("cameraBoundary"))
	return rv
}


// The boundary of the area within which the map view’s center needs to remain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/cameraboundary-swift.property
func (m_ MKMapCameraBoundary) SetCameraBoundary(value IMKMapCameraBoundary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraBoundary:"), value)
}


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraBoundary) CameraZoomRange() MKMapCameraZoomRange {
	rv := objc.Send[MKMapCameraZoomRange](m_.ID, objc.Sel("cameraZoomRange"))
	return rv
}


// The zoom range to apply to the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/camerazoomrange-swift.property
func (m_ MKMapCameraBoundary) SetCameraZoomRange(value IMKMapCameraZoomRange) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraZoomRange:"), value)
}



