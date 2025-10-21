// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKMultiPolyline] class.
var (
	MKMultiPolylineClass     _MKMultiPolylineClass
	MKMultiPolylineClassOnce sync.Once
)

func getMKMultiPolylineClass() _MKMultiPolylineClass {
	MKMultiPolylineClassOnce.Do(func() {
		MKMultiPolylineClass = _MKMultiPolylineClass{objc.GetClass("MKMultiPolyline")}
	})
	return MKMultiPolylineClass
}

type _MKMultiPolylineClass struct {
	class objc.Class
}

// An interface definition for the [MKMultiPolyline] class.
type IMKMultiPolyline interface {
	IMKShape
}

// A collection of multipolyline shapes, each consisting of one or more connected line segments.
//
// Use a object when you have multiple distinct polyline shapes that you intend to render using the same style.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolyline
type MKMultiPolyline struct {
	MKShape
}

// MKMultiPolylineFrom constructs a [MKMultiPolyline] from an unsafe.Pointer.
//
// A collection of multipolyline shapes, each consisting of one or more connected line segments.
func MKMultiPolylineFrom(ptr unsafe.Pointer) MKMultiPolyline {
	return MKMultiPolyline{
		MKShape: MKShapeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMultiPolylineClass) Alloc() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMultiPolylineClass) New() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPolyline) Init() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPolyline) Autorelease() MKMultiPolyline {
	rv := objc.Send[MKMultiPolyline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPolyline creates a new MKMultiPolyline instance.
func NewMKMultiPolyline() MKMultiPolyline {
	return getMKMultiPolylineClass().New()
}


// An array containing the polyline objects that make up the multipolyline object.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmultipolyline/polylines
func (m_ MKMultiPolyline) Polylines() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("polylines"))
	return rv
}


// SetPolylines sets the value of the polylines property.
// An array containing the polyline objects that make up the multipolyline object.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmultipolyline/polylines
func (m_ MKMultiPolyline) SetPolylines(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPolylines:"), value)
}



