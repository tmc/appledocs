// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMachPort */


/* debug [class_header]: Header for NSMachPort */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MachPort */
// An interface definition for the [MachPort] class.
type IMachPort interface {
	IPort
	
/* debug [class_interface_properties]: Properties for MachPort */
	// properties:
	MachPort() uint32 /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MachPort */
	// methods:
	Delegate() unsafe.Pointer
	RemoveFromRunLoopForMode(runLoop IRunLoop, mode RunLoopMode)
	ScheduleInRunLoopForMode(runLoop IRunLoop, mode RunLoopMode)
	SetDelegate(anObject unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MachPort */
// Alloc allocates a new instance without initialization.
func (mc _MachPortClass) Alloc() MachPort {
	rv := objc.Send[MachPort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MachPort */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MachPort */

// Initializes a newly allocated object with a given Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:)
func NewMachPortWithMachPort(machPort uint32 /* not a class type */) MachPort {
	instance := getMachPortClass().Alloc()
	rv := objc.Send[MachPort](instance.ID, objc.Sel("initWithMachPort:"), machPort)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMachPortWithMachPort */


// Initializes a newly allocated object with a given Mach port and the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:options:)
func NewMachPortWithMachPortOptions(machPort uint32 /* not a class type */, f MachPortOptions) MachPort {
	instance := getMachPortClass().Alloc()
	rv := objc.Send[MachPort](instance.ID, objc.Sel("initWithMachPort:options:"), machPort, f)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMachPortWithMachPortOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MachPort */

// Creates and returns a port object configured with the given Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/port(withMachPort:)
func (mc _MachPortClass) PortWithMachPort(machPort uint32 /* not a class type */) IPort {
	rv := objc.Send[Port](objc.ID(mc.class), objc.Sel("portWithMachPort:"), machPort)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PortWithMachPort) */


// Creates and returns a port object configured with the specified options and the given Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/port(withMachPort:options:)
func (mc _MachPortClass) PortWithMachPortOptions(machPort uint32 /* not a class type */, f MachPortOptions) IPort {
	rv := objc.Send[Port](objc.ID(mc.class), objc.Sel("portWithMachPort:options:"), machPort, f)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PortWithMachPortOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MachPort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MachPort */

// Returns the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/delegate()
func (m_ MachPort) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_methods/method]: Delegate */


// Removes the receiver from the run loop mode of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/remove(from:forMode:)
func (m_ MachPort) RemoveFromRunLoopForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeFromRunLoop:forMode:"), runLoop, mode)
}/* debug [instance_methods/method]: RemoveFromRunLoopForMode */


// Schedules the receiver into the run loop mode of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/schedule(in:forMode:)
func (m_ MachPort) ScheduleInRunLoopForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scheduleInRunLoop:forMode:"), runLoop, mode)
}/* debug [instance_methods/method]: ScheduleInRunLoopForMode */


// Sets the receiver’s delegate to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/setDelegate(_:)
func (m_ MachPort) SetDelegate(anObject unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), anObject)
}/* debug [instance_methods/method]: SetDelegate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MachPort */

// The Mach port used by the receiver, represented as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachPort/machPort
func (m_ MachPort) MachPort() uint32 /* not a class type */ {
	rv := objc.Send[uint32](m_.ID, objc.Sel("machPort"))
	return rv
}/* debug [instance_properties/getter]: machPort */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMachPort */


