// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mouthPosition] class.
var mouthPositionClass = _mouthPositionClass{objc.GetClass("mouthPosition")}

type _mouthPositionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-c.ivar

type mouthPosition struct {
	objectivec.Object
}

// mouthPositionFrom constructs a [mouthPosition] from an unsafe.Pointer.
func mouthPositionFrom(ptr unsafe.Pointer) mouthPosition {
	return mouthPosition{objectivec.Object{objc.ID(ptr)}}
}



