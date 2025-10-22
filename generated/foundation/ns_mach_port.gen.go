// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	Delegate() objc.ID
	RemoveFromRunLoopForMode(runLoop IRunLoop, mode IRunLoopMode)
	ScheduleInRunLoopForMode(runLoop IRunLoop, mode IRunLoopMode)
	SetDelegate(anObject objectivec.IObject)
	MachPort() uint32
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




// Initializes a newly allocated object with a given Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:)

func NewMachPortWithMachPort(machPort Iuint32) MachPort {
	instance := getMachPortClass().Alloc()
	rv := objc.Send[MachPort](instance.ID, objc.Sel("initWithMachPort:"), machPort)
	rv.Autorelease()
	return rv
}



// Initializes a newly allocated object with a given Mach port and the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:options:)

func NewMachPortWithMachPortOptions(machPort Iuint32, f IMachPortOptions) MachPort {
	instance := getMachPortClass().Alloc()
	rv := objc.Send[MachPort](instance.ID, objc.Sel("initWithMachPort:options:"), machPort, f)
	rv.Autorelease()
	return rv
}



// Creates and returns a port object configured with the given Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/port(withMachPort:)

func (mc _MachPortClass) PortWithMachPort(machPort Iuint32) Port {
	rv := objc.Send[Port](objc.ID(mc.class), objc.Sel("portWithMachPort:"), machPort)
	return rv
}


// Creates and returns a port object configured with the specified options and the given Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/port(withMachPort:options:)

func (mc _MachPortClass) PortWithMachPortOptions(machPort Iuint32, f IMachPortOptions) Port {
	rv := objc.Send[Port](objc.ID(mc.class), objc.Sel("portWithMachPort:options:"), machPort, f)
	return rv
}



// Returns the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/delegate()

func (m_ MachPort) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}



// Removes the receiver from the run loop mode of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/remove(from:forMode:)

func (m_ MachPort) RemoveFromRunLoopForMode(runLoop IRunLoop, mode IRunLoopMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeFromRunLoop:forMode:"), runLoop, mode)
}



// Schedules the receiver into the run loop mode of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/schedule(in:forMode:)

func (m_ MachPort) ScheduleInRunLoopForMode(runLoop IRunLoop, mode IRunLoopMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scheduleInRunLoop:forMode:"), runLoop, mode)
}



// Sets the receiver’s delegate to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/setDelegate(_:)

func (m_ MachPort) SetDelegate(anObject objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), anObject)
}


// The Mach port used by the receiver, represented as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/machPort

func (m_ MachPort) MachPort() uint32 {
	rv := objc.Send[uint32](m_.ID, objc.Sel("machPort"))
	return rv
}


