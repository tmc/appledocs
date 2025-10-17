// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RulerMarker] class.
var RulerMarkerClass objc.Class

func init() {
	RulerMarkerClass = objc.GetClass("NSRulerMarker")
}

type RulerMarker struct {
	objc.ID
}

func RulerMarkerFrom(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{
		ID: objc.ID(ptr),
	}
}



