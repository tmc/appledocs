// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKGeodesicPolyline] class.
var (
	MKGeodesicPolylineClass     _MKGeodesicPolylineClass
	MKGeodesicPolylineClassOnce sync.Once
)

func getMKGeodesicPolylineClass() _MKGeodesicPolylineClass {
	MKGeodesicPolylineClassOnce.Do(func() {
		MKGeodesicPolylineClass = _MKGeodesicPolylineClass{objc.GetClass("MKGeodesicPolyline")}
	})
	return MKGeodesicPolylineClass
}

type _MKGeodesicPolylineClass struct {
	class objc.Class
}

// An interface definition for the [MKGeodesicPolyline] class.
type IMKGeodesicPolyline interface {
	IMKPolyline
}

// An open polygon overlay consisting of line segments that follow the contours of the Earth to create the shortest path between the specified points.
//
// A geodesic polyline contains a set of points that connect end-to-end in the order that you provide them. The first and last points don’t automatically connect to each other. When displaying on a two-dimensional map view, the line segment between any two points may appear curved.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeodesicPolyline
type MKGeodesicPolyline struct {
	MKPolyline
}

// MKGeodesicPolylineFrom constructs a [MKGeodesicPolyline] from an unsafe.Pointer.
//
// An open polygon overlay consisting of line segments that follow the contours of the Earth to create the shortest path between the specified points.
func MKGeodesicPolylineFrom(ptr unsafe.Pointer) MKGeodesicPolyline {
	return MKGeodesicPolyline{
		MKPolyline: MKPolylineFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKGeodesicPolylineClass) Alloc() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKGeodesicPolylineClass) New() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeodesicPolyline) Init() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeodesicPolyline) Autorelease() MKGeodesicPolyline {
	rv := objc.Send[MKGeodesicPolyline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeodesicPolyline creates a new MKGeodesicPolyline instance.
func NewMKGeodesicPolyline() MKGeodesicPolyline {
	return getMKGeodesicPolylineClass().New()
}




