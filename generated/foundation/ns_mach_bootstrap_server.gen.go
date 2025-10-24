// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMachBootstrapServer */


/* debug [class_header]: Header for NSMachBootstrapServer */
// The class instance for the [MachBootstrapServer] class.
var (
	MachBootstrapServerClass     _MachBootstrapServerClass
	MachBootstrapServerClassOnce sync.Once
)

func getMachBootstrapServerClass() _MachBootstrapServerClass {
	MachBootstrapServerClassOnce.Do(func() {
		MachBootstrapServerClass = _MachBootstrapServerClass{objc.GetClass("NSMachBootstrapServer")}
	})
	return MachBootstrapServerClass
}

type _MachBootstrapServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MachBootstrapServer */
// An interface definition for the [MachBootstrapServer] class.
type IMachBootstrapServer interface {
	IPortNameServer
	
/* debug [class_interface_properties]: Properties for MachBootstrapServer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MachBootstrapServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MachBootstrapServer */
// Alloc allocates a new instance without initialization.
func (mc _MachBootstrapServerClass) Alloc() MachBootstrapServer {
	rv := objc.Send[MachBootstrapServer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MachBootstrapServerClass) New() MachBootstrapServer {
	rv := objc.Send[MachBootstrapServer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MachBootstrapServer) Init() MachBootstrapServer {
	rv := objc.Send[MachBootstrapServer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MachBootstrapServer) Autorelease() MachBootstrapServer {
	rv := objc.Send[MachBootstrapServer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMachBootstrapServer creates a new MachBootstrapServer instance.
func NewMachBootstrapServer() MachBootstrapServer {
	return getMachBootstrapServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MachBootstrapServer */
// A port name server that takes and returns Mach port objects.
//
// Port removal functionality is not supported in ; if you want to cancel a service, you have to destroy the port (invalidate the given to ).


// A port name server that takes and returns Mach port objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachBootstrapServer
type MachBootstrapServer struct {
	PortNameServer
}

// MachBootstrapServerFrom constructs a [MachBootstrapServer] from an unsafe.Pointer.
//
// A port name server that takes and returns Mach port objects.
func MachBootstrapServerFrom(ptr unsafe.Pointer) MachBootstrapServer {
	return MachBootstrapServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MachBootstrapServer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MachBootstrapServer */

// Returns the shared instance of the bootstrap server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMachBootstrapServer/sharedInstance
func (mc _MachBootstrapServerClass) SharedInstance() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MachBootstrapServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MachBootstrapServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MachBootstrapServer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMachBootstrapServer */



