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
// Alloc allocates a new instance without initialization.
func (mc _MachPortClass) Alloc() MachPort {
	rv := objc.Send[MachPort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MachPortClass) New() MachPort {
	rv := objc.Send[MachPort](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MachPort) Init() MachPort {
	rv := objc.Send[MachPort](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MachPort) Autorelease() MachPort {
	rv := objc.Send[MachPort](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMachPort creates a new MachPort instance.
func NewMachPort() MachPort {
	return machPortClass.New()
}




