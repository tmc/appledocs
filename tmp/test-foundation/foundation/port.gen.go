// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var PortClass _PortClass

func init() {
	PortClass = _PortClass{objc.GetClass("NSPort")}
}

type _PortClass struct {
	class objc.Class
}

type Port struct {
	objc.ID
}

func PortFrom(ptr unsafe.Pointer) Port {
	return Port{
		ID: objc.ID(ptr),
	}
}




