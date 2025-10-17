// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TapDescription] class.
var tapDescriptionClass = _TapDescriptionClass{objc.GetClass("CATapDescription")}

type _TapDescriptionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapDescription

type TapDescription struct {
	objectivec.Object
}

// TapDescriptionFrom constructs a [TapDescription] from an unsafe.Pointer.
func TapDescriptionFrom(ptr unsafe.Pointer) TapDescription {
	return TapDescription{objectivec.Object{objc.ID(ptr)}}
}


