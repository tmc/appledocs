// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Port] class.
var PortClass objc.Class

func init() {
	PortClass = objc.GetClass("NSPort")
}

type Port struct {
	objc.ID
}

func PortFrom(ptr unsafe.Pointer) Port {
	return Port{
		ID: objc.ID(ptr),
	}
}




