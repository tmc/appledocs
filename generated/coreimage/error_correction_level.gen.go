// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [errorCorrectionLevel] class.
var errorCorrectionLevelClass = _errorCorrectionLevelClass{objc.GetClass("errorCorrectionLevel")}

type _errorCorrectionLevelClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-c.ivar

type errorCorrectionLevel struct {
	objectivec.Object
}

// errorCorrectionLevelFrom constructs a [errorCorrectionLevel] from an unsafe.Pointer.
func errorCorrectionLevelFrom(ptr unsafe.Pointer) errorCorrectionLevel {
	return errorCorrectionLevel{objectivec.Object{objc.ID(ptr)}}
}



