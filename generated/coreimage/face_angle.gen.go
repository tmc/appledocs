// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [faceAngle] class.
var faceAngleClass = _faceAngleClass{objc.GetClass("faceAngle")}

type _faceAngleClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-c.ivar

type faceAngle struct {
	objectivec.Object
}

// faceAngleFrom constructs a [faceAngle] from an unsafe.Pointer.
func faceAngleFrom(ptr unsafe.Pointer) faceAngle {
	return faceAngle{objectivec.Object{objc.ID(ptr)}}
}



