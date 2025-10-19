// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [leftEyeClosed] class.
var leftEyeClosedClass = _leftEyeClosedClass{objc.GetClass("leftEyeClosed")}

type _leftEyeClosedClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-c.ivar

type leftEyeClosed struct {
	objectivec.Object
}

// leftEyeClosedFrom constructs a [leftEyeClosed] from an unsafe.Pointer.
func leftEyeClosedFrom(ptr unsafe.Pointer) leftEyeClosed {
	return leftEyeClosed{objectivec.Object{objc.ID(ptr)}}
}



