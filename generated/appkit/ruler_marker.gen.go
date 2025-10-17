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



