// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [rightEyeClosed] class.
var rightEyeClosedClass = _rightEyeClosedClass{objc.GetClass("rightEyeClosed")}

type _rightEyeClosedClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyeClosed-c.ivar

type rightEyeClosed struct {
	objectivec.Object
}

// rightEyeClosedFrom constructs a [rightEyeClosed] from an unsafe.Pointer.
func rightEyeClosedFrom(ptr unsafe.Pointer) rightEyeClosed {
	return rightEyeClosed{objectivec.Object{objc.ID(ptr)}}
}



