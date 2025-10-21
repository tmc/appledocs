// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKPolygon] class.
var (
	MKPolygonClass     _MKPolygonClass
	MKPolygonClassOnce sync.Once
)

func getMKPolygonClass() _MKPolygonClass {
	MKPolygonClassOnce.Do(func() {
		MKPolygonClass = _MKPolygonClass{objc.GetClass("MKPolygon")}
	})
	return MKPolygonClass
}

type _MKPolygonClass struct {
	class objc.Class
}

// An interface definition for the [MKPolygon] class.
type IMKPolygon interface {
	IMKMultiPoint
}

// A closed polygon overlay.
//
// The points you add to this overlay connect end-to-end in the order you provide them. The first and last points connect to each other to create a closed shape. When creating a polygon, you can mask out portions of the polygon by specifying one or more interior polygons. For the polygons you specify, this class uses the even-odd fill rule to determine the final occupied area. When applied to overlapping polygons, this rule can cause the framework to mask specific regions out and thereby remove them from the total occupied area. For more information about how fill rules apply to paths, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygon
type MKPolygon struct {
	MKMultiPoint
}

// MKPolygonFrom constructs a [MKPolygon] from an unsafe.Pointer.
//
// A closed polygon overlay.
func MKPolygonFrom(ptr unsafe.Pointer) MKPolygon {
	return MKPolygon{
		MKMultiPoint: MKMultiPointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKPolygonClass) Alloc() MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKPolygonClass) New() MKPolygon {
	rv := objc.Send[MKPolygon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolygon) Init() MKPolygon {
	rv := objc.Send[MKPolygon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolygon) Autorelease() MKPolygon {
	rv := objc.Send[MKPolygon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolygon creates a new MKPolygon instance.
func NewMKPolygon() MKPolygon {
	return getMKPolygonClass().New()
}




