// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MachPort] class.
var machPortClass = _MachPortClass{objc.GetClass("NSMachPort")}

type _MachPortClass struct {
	class objc.Class
}

// An interface definition for the [MachPort] class.
type IMachPort interface {
	IPort
}

// A port that can be used as an endpoint for distributed object connections (or raw messaging). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort

type MachPort struct {
	Port
}

// MachPortFrom constructs a [MachPort] from an unsafe.Pointer.
//
// A port that can be used as an endpoint for distributed object connections (or raw messaging).
func MachPortFrom(ptr unsafe.Pointer) MachPort {
	return MachPort{
		Port: PortFrom(ptr),
	}
}



