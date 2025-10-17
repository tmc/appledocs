// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XPCCoder] class.
var XPCCoderClass objc.Class

func init() {
	XPCCoderClass = objc.GetClass("NSXPCCoder")
}

type XPCCoder struct {
	objc.ID
}

func XPCCoderFrom(ptr unsafe.Pointer) XPCCoder {
	return XPCCoder{
		ID: objc.ID(ptr),
	}
}



