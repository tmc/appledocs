// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PortNameServer] class.
type IPortNameServer interface {
	objectivec.IObject
	PortForName(name string) Port
	PortForNameHost(name string, host string) Port
	RegisterPortName(port Port, name string) bool
	RemovePortForName(name string) bool
}

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

// Alloc allocates a new instance without initialization.
func (pc _PortNameServerClass) Alloc() PortNameServer {
	rv := objc.Send[PortNameServer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the single instance of for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer/systemDefaultPortNameServer
func (pc _PortNameServerClass) SystemDefaultPortNameServer() PortNameServer {
	rv := objc.Send[PortNameServer](objc.ID(pc.class), objc.Sel("systemDefaultPortNameServer"))
	return rv
}


// Looks up and returns the port registered under the specified name on the local host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer/portForName:
func (p_ PortNameServer) PortForName(name string) Port {
	rv := objc.Send[Port](p_.ID, objc.Sel("portForName:"), objc.String(name))
	return rv
}


// Looks up and returns the port registered under the specified name on a specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer/portForName:host:
func (p_ PortNameServer) PortForNameHost(name string, host string) Port {
	rv := objc.Send[Port](p_.ID, objc.Sel("portForName:host:"), objc.String(name), objc.String(host))
	return rv
}


// Makes a given port available on the network under a specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer/registerPort:name:
func (p_ PortNameServer) RegisterPortName(port Port, name string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("registerPort:name:"), port, objc.String(name))
	return rv
}


// Unregisters the port for a given name on the local host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortNameServer/removePortForName:
func (p_ PortNameServer) RemovePortForName(name string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("removePortForName:"), objc.String(name))
	return rv
}



