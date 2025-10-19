// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasMouthPosition] class.
var hasMouthPositionClass = _hasMouthPositionClass{objc.GetClass("hasMouthPosition")}

type _hasMouthPositionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-c.ivar

type hasMouthPosition struct {
	objectivec.Object
}

// hasMouthPositionFrom constructs a [hasMouthPosition] from an unsafe.Pointer.
func hasMouthPositionFrom(ptr unsafe.Pointer) hasMouthPosition {
	return hasMouthPosition{objectivec.Object{objc.ID(ptr)}}
}



