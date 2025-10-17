// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MachPort] class.
var MachPortClass objc.Class

func init() {
	MachPortClass = objc.GetClass("NSMachPort")
}

type MachPort struct {
	objc.ID
}

func MachPortFrom(ptr unsafe.Pointer) MachPort {
	return MachPort{
		ID: objc.ID(ptr),
	}
}



