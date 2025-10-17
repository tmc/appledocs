// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XPCCoder] class.
var XPCCoderClass = _XPCCoderClass{objc.GetClass("NSXPCCoder")}

type _XPCCoderClass struct {
	class objc.Class
}

type XPCCoder struct {
	objc.ID
}

func XPCCoderFrom(ptr unsafe.Pointer) XPCCoder {
	return XPCCoder{
		ID: objc.ID(ptr),
	}
}




