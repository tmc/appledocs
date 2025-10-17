// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MachPort] class.
var MachPortClass = _MachPortClass{objc.GetClass("NSMachPort")}

type _MachPortClass struct {
	class objc.Class
}

type MachPort struct {
	objc.ID
}

func MachPortFrom(ptr unsafe.Pointer) MachPort {
	return MachPort{
		ID: objc.ID(ptr),
	}
}




