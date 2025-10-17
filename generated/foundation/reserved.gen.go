// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [reserved] class.
var reservedClass objc.Class

func init() {
	reservedClass = objc.GetClass("reserved")
}

type reserved struct {
	objc.ID
}

func reservedFrom(ptr unsafe.Pointer) reserved {
	return reserved{
		ID: objc.ID(ptr),
	}
}



