// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKOverlayRenderer] class.
var (
	MKOverlayRendererClass     _MKOverlayRendererClass
	MKOverlayRendererClassOnce sync.Once
)

func getMKOverlayRendererClass() _MKOverlayRendererClass {
	MKOverlayRendererClassOnce.Do(func() {
		MKOverlayRendererClass = _MKOverlayRendererClass{objc.GetClass("MKOverlayRenderer")}
	})
	return MKOverlayRendererClass
}

type _MKOverlayRendererClass struct {
	class objc.Class
}

// An interface definition for the [MKOverlayRenderer] class.
type IMKOverlayRenderer interface {
	objectivec.IObject
}

// The shared infrastructure for drawing overlays on the map surface.
//
// An overlay renderer draws the visual representation of an overlay object — that is, an object that conforms to the protocol. This class defines the drawing infrastructure the map view uses. Subclasses need to override the method to draw the contents of the overlay. The MapKit framework provides several concrete instances of overlay renderers. Specifically, it provides renderers for each of the concrete overlay objects. You can use one of these existing renderers or define your own subclasses if you want to draw the overlay contents differently. You can subclass to create overlays based on custom shapes, content, or drawing techniques. The only method subclasses need to override is the method. However, if your class contains content that may not be ready for drawing right away, you need to also override the method and use it to report when your class is ready and able to draw. The map view may tile large overlays and distribute the rendering of each tile to separate threads. Therefore, the implementation of your method needs to be safe to run from background threads and from multiple threads simultaneously.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer
type MKOverlayRenderer struct {
	objectivec.Object
}

// MKOverlayRendererFrom constructs a [MKOverlayRenderer] from an unsafe.Pointer.
//
// The shared infrastructure for drawing overlays on the map surface.
func MKOverlayRendererFrom(ptr unsafe.Pointer) MKOverlayRenderer {
	return MKOverlayRenderer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKOverlayRendererClass) Alloc() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKOverlayRendererClass) New() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKOverlayRenderer) Init() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKOverlayRenderer) Autorelease() MKOverlayRenderer {
	rv := objc.Send[MKOverlayRenderer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKOverlayRenderer creates a new MKOverlayRenderer instance.
func NewMKOverlayRenderer() MKOverlayRenderer {
	return getMKOverlayRendererClass().New()
}




