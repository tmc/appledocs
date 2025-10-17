// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XPCInterface] class.
var XPCInterfaceClass objc.Class

func init() {
	XPCInterfaceClass = objc.GetClass("NSXPCInterface")
}

type XPCInterface struct {
	objc.ID
}

func XPCInterfaceFrom(ptr unsafe.Pointer) XPCInterface {
	return XPCInterface{
		ID: objc.ID(ptr),
	}
}




