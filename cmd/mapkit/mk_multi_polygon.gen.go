// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKMultiPolygon] class.
var (
	MKMultiPolygonClass     _MKMultiPolygonClass
	MKMultiPolygonClassOnce sync.Once
)

func getMKMultiPolygonClass() _MKMultiPolygonClass {
	MKMultiPolygonClassOnce.Do(func() {
		MKMultiPolygonClass = _MKMultiPolygonClass{objc.GetClass("MKMultiPolygon")}
	})
	return MKMultiPolygonClass
}

type _MKMultiPolygonClass struct {
	class objc.Class
}

// An interface definition for the [MKMultiPolygon] class.
type IMKMultiPolygon interface {
	IMKShape
}

// A collection of multiple closed polygon overlays.
//
// Use a when you have multiple distinct polygon shapes that you intend to render using the same style.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMultiPolygon
type MKMultiPolygon struct {
	MKShape
}

// MKMultiPolygonFrom constructs a [MKMultiPolygon] from an unsafe.Pointer.
//
// A collection of multiple closed polygon overlays.
func MKMultiPolygonFrom(ptr unsafe.Pointer) MKMultiPolygon {
	return MKMultiPolygon{
		MKShape: MKShapeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMultiPolygonClass) Alloc() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMultiPolygonClass) New() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMultiPolygon) Init() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMultiPolygon) Autorelease() MKMultiPolygon {
	rv := objc.Send[MKMultiPolygon](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMultiPolygon creates a new MKMultiPolygon instance.
func NewMKMultiPolygon() MKMultiPolygon {
	return getMKMultiPolygonClass().New()
}




