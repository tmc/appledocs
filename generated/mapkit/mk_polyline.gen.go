// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKPolyline] class.
var (
	MKPolylineClass     _MKPolylineClass
	MKPolylineClassOnce sync.Once
)

func getMKPolylineClass() _MKPolylineClass {
	MKPolylineClassOnce.Do(func() {
		MKPolylineClass = _MKPolylineClass{objc.GetClass("MKPolyline")}
	})
	return MKPolylineClass
}

type _MKPolylineClass struct {
	class objc.Class
}

// An interface definition for the [MKPolyline] class.
type IMKPolyline interface {
	IMKMultiPoint
}

// An open polygon overlay consisting of one or more connected line segments.
//
// The points connect end-to-end in the order that you provide them. The first and last points don’t automatically connect to each other.


// An open polygon overlay consisting of one or more connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolyline
type MKPolyline struct {
	MKMultiPoint
}

// MKPolylineFrom constructs a [MKPolyline] from an unsafe.Pointer.
//
// An open polygon overlay consisting of one or more connected line segments.
func MKPolylineFrom(ptr unsafe.Pointer) MKPolyline {
	return MKPolyline{
		MKMultiPoint: MKMultiPointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKPolylineClass) Alloc() MKPolyline {
	rv := objc.Send[MKPolyline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKPolylineClass) New() MKPolyline {
	rv := objc.Send[MKPolyline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolyline) Init() MKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolyline) Autorelease() MKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolyline creates a new MKPolyline instance.
func NewMKPolyline() MKPolyline {
	return getMKPolylineClass().New()
}




