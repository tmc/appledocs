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





