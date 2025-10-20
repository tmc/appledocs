// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var reservedClass _reservedClass

func init() {
	reservedClass = _reservedClass{objc.GetClass("reserved")}
}

type _reservedClass struct {
	class objc.Class
}

type reserved struct {
	objc.ID
}

func reservedFrom(ptr unsafe.Pointer) reserved {
	return reserved{
		ID: objc.ID(ptr),
	}
}




