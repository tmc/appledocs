// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RulerMarker] class.
var rulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}

type _RulerMarkerClass struct {
	class objc.Class
}

// An interface definition for the [RulerMarker] class.
type IRulerMarker interface {
	objectivec.IObject
}

// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return rulerMarkerClass.New()
}




