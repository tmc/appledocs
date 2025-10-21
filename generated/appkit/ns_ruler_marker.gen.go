// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RulerMarker] class.
var (
	RulerMarkerClass     _RulerMarkerClass
	RulerMarkerClassOnce sync.Once
)

func getRulerMarkerClass() _RulerMarkerClass {
	RulerMarkerClassOnce.Do(func() {
		RulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}
	})
	return RulerMarkerClass
}

type _RulerMarkerClass struct {
	class objc.Class
}

// An interface definition for the [RulerMarker] class.
type IRulerMarker interface {
	objectivec.IObject
	TrackMouseAdding(mouseDownEvent unsafe.Pointer, isAdding bool) bool
}

// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view.
//
// An example of a marker is the representation of a margin or tab setting, or the edges of a graphic on the page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker
type RulerMarker struct {
	objectivec.Object
}

// RulerMarkerFrom constructs a [RulerMarker] from an unsafe.Pointer.
//
// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view.
func RulerMarkerFrom(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.Send[RulerMarker](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RulerMarker) Autorelease() RulerMarker {
	rv := objc.Send[RulerMarker](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRulerMarker creates a new RulerMarker instance.
func NewRulerMarker() RulerMarker {
	return getRulerMarkerClass().New()
}


// Handles user manipulation of the receiver in its ruler view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/trackMouse(with:adding:)
func (r_ RulerMarker) TrackMouseAdding(mouseDownEvent unsafe.Pointer, isAdding bool) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMouse:adding:"), mouseDownEvent, isAdding)
	return rv
}

// The receiver’s image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r_ RulerMarker) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The receiver’s image.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r_ RulerMarker) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setImage:"), value)
}

// The receiver’s ruler view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/ruler
func (r_ RulerMarker) Ruler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("ruler"))
	return rv
}



