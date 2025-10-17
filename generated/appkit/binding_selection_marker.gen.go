// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BindingSelectionMarker] class.
var bindingSelectionMarkerClass = _BindingSelectionMarkerClass{objc.GetClass("NSBindingSelectionMarker")}

type _BindingSelectionMarkerClass struct {
	class objc.Class
}

// An interface definition for the [BindingSelectionMarker] class.
type IBindingSelectionMarker interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker

type BindingSelectionMarker struct {
	objectivec.Object
}

// BindingSelectionMarkerFrom constructs a [BindingSelectionMarker] from an unsafe.Pointer.
func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{objectivec.Object{objc.ID(ptr)}}
}



