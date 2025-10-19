// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [errorCorrectedPayload] class.
var errorCorrectedPayloadClass = _errorCorrectedPayloadClass{objc.GetClass("errorCorrectedPayload")}

type _errorCorrectedPayloadClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectedPayload-c.ivar

type errorCorrectedPayload struct {
	objectivec.Object
}

// errorCorrectedPayloadFrom constructs a [errorCorrectedPayload] from an unsafe.Pointer.
func errorCorrectedPayloadFrom(ptr unsafe.Pointer) errorCorrectedPayload {
	return errorCorrectedPayload{objectivec.Object{objc.ID(ptr)}}
}



