// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MachPort] class.
var (
	MachPortClass     _MachPortClass
	MachPortClassOnce sync.Once
)

func getMachPortClass() _MachPortClass {
	MachPortClassOnce.Do(func() {
		MachPortClass = _MachPortClass{objc.GetClass("NSMachPort")}
	})
	return MachPortClass
}

type _MachPortClass struct {
	class objc.Class
}

// An interface definition for the [MachPort] class.
type IMachPort interface {
	IPort
	// properties:
	MachPort() unsafe.Pointer
	SetMachPort(value unsafe.Pointer)
	// methods:
}

// A port that can be used as an endpoint for distributed object connections (or raw messaging).
//
// is a subclass of that wraps a Mach port, the fundamental communication port in macOS. allows for local (on the same machine) communication only. A companion class, , allows for both local and remote distributed object communication, but may be more expensive than for the local case. To use effectively, you should be familiar with Mach ports, port access rights, and Mach messages. See the Mach OS documentation for more information.


// A port that can be used as an endpoint for distributed object connections (or raw messaging).
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getMachPortClass().New()
}



// The Mach port used by the receiver, represented as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport/machport
func (m_ MachPort) MachPort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("machPort"))
	return rv
}


// The Mach port used by the receiver, represented as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport/machport
func (m_ MachPort) SetMachPort(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMachPort:"), value)
}



