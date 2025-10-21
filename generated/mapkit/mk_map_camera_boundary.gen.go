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
}

// A boundary of an area within which the map’s center needs to remain.
//
// The constraints of the camera boundary restrict the center point of your map.
//
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




