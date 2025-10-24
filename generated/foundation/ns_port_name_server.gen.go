// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPortNameServer */


/* debug [class_header]: Header for NSPortNameServer */
// The class instance for the [PortNameServer] class.
var (
	PortNameServerClass     _PortNameServerClass
	PortNameServerClassOnce sync.Once
)

func getPortNameServerClass() _PortNameServerClass {
	PortNameServerClassOnce.Do(func() {
		PortNameServerClass = _PortNameServerClass{objc.GetClass("NSPortNameServer")}
	})
	return PortNameServerClass
}

type _PortNameServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PortNameServer */
// An interface definition for the [PortNameServer] class.
type IPortNameServer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PortNameServer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PortNameServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PortNameServer */
// Alloc allocates a new instance without initialization.
func (pc _PortNameServerClass) Alloc() PortNameServer {
	rv := objc.Send[PortNameServer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PortNameServerClass) New() PortNameServer {
	rv := objc.Send[PortNameServer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PortNameServer) Init() PortNameServer {
	rv := objc.Send[PortNameServer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PortNameServer) Autorelease() PortNameServer {
	rv := objc.Send[PortNameServer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPortNameServer creates a new PortNameServer instance.
func NewPortNameServer() PortNameServer {
	return getPortNameServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PortNameServer */
// An object-oriented interface to the port registration service used by the distributed objects system.
//
// objects use this interface to contact each other and to distribute objects over the network; you should rarely need to interact directly with an . You get an object by using the class method—never allocate and initialize an instance directly. With the default server object you can register an object under a given name, making it available on the network, and also unregister it so that it can’t be looked up (although other applications that have already looked up the object can still use it until it becomes invalid). See the class specification for more information.


// An object-oriented interface to the port registration service used by the distributed objects system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer
type PortNameServer struct {
	objectivec.Object
}

// PortNameServerFrom constructs a [PortNameServer] from an unsafe.Pointer.
//
// An object-oriented interface to the port registration service used by the distributed objects system.
func PortNameServerFrom(ptr unsafe.Pointer) PortNameServer {
	return PortNameServer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PortNameServer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PortNameServer */

// Returns the single instance of for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer/systemDefaultPortNameServer
func (pc _PortNameServerClass) SystemDefaultPortNameServer() IPortNameServer {
	rv := objc.Send[PortNameServer](objc.ID(pc.class), objc.Sel("systemDefaultPortNameServer"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SystemDefaultPortNameServer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PortNameServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PortNameServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PortNameServer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPortNameServer */



