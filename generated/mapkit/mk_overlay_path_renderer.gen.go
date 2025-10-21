// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKOverlayPathRenderer] class.
var (
	MKOverlayPathRendererClass     _MKOverlayPathRendererClass
	MKOverlayPathRendererClassOnce sync.Once
)

func getMKOverlayPathRendererClass() _MKOverlayPathRendererClass {
	MKOverlayPathRendererClassOnce.Do(func() {
		MKOverlayPathRendererClass = _MKOverlayPathRendererClass{objc.GetClass("MKOverlayPathRenderer")}
	})
	return MKOverlayPathRendererClass
}

type _MKOverlayPathRendererClass struct {
	class objc.Class
}

// An interface definition for the [MKOverlayPathRenderer] class.
type IMKOverlayPathRenderer interface {
	IMKOverlayRenderer
}

// The visual representation of a path-based overlay.
//
// Use this renderer when a object defines your overlay’s shape. By default, this renderer fills the overlay’s shape and represents the strokes of the path using its current attributes. You can use this class as-is or subclass it to define additional drawing behaviors. If you subclass it, override the method and use that method to build the appropriate path object. To change the path, invalidate it and recreate the path using the new data your subclass obtains.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathRenderer
type MKOverlayPathRenderer struct {
	MKOverlayRenderer
}

// MKOverlayPathRendererFrom constructs a [MKOverlayPathRenderer] from an unsafe.Pointer.
//
// The visual representation of a path-based overlay.
func MKOverlayPathRendererFrom(ptr unsafe.Pointer) MKOverlayPathRenderer {
	return MKOverlayPathRenderer{
		MKOverlayRenderer: MKOverlayRendererFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKOverlayPathRendererClass) Alloc() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKOverlayPathRendererClass) New() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKOverlayPathRenderer) Init() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKOverlayPathRenderer) Autorelease() MKOverlayPathRenderer {
	rv := objc.Send[MKOverlayPathRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKOverlayPathRenderer creates a new MKOverlayPathRenderer instance.
func NewMKOverlayPathRenderer() MKOverlayPathRenderer {
	return getMKOverlayPathRendererClass().New()
}




