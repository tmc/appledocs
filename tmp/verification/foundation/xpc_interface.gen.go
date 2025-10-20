// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var xPCInterfaceClass _XPCInterfaceClass

func init() {
	xPCInterfaceClass = _XPCInterfaceClass{objc.GetClass("NSXPCInterface")}
}

type _XPCInterfaceClass struct {
	class objc.Class
}

type XPCInterface struct {
	objc.ID
}

func XPCInterfaceFrom(ptr unsafe.Pointer) XPCInterface {
	return XPCInterface{
		ID: objc.ID(ptr),
	}
}




