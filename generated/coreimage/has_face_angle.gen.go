// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasFaceAngle] class.
var hasFaceAngleClass = _hasFaceAngleClass{objc.GetClass("hasFaceAngle")}

type _hasFaceAngleClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-c.ivar

type hasFaceAngle struct {
	objectivec.Object
}

// hasFaceAngleFrom constructs a [hasFaceAngle] from an unsafe.Pointer.
func hasFaceAngleFrom(ptr unsafe.Pointer) hasFaceAngle {
	return hasFaceAngle{objectivec.Object{objc.ID(ptr)}}
}



