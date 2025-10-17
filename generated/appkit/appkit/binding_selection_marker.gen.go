// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BindingSelectionMarker] class.
var BindingSelectionMarkerClass objc.Class

func init() {
	BindingSelectionMarkerClass = objc.GetClass("NSBindingSelectionMarker")
}

type BindingSelectionMarker struct {
	objc.ID
}

func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{
		ID: objc.ID(ptr),
	}
}




