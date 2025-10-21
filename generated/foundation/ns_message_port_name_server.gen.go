// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MessagePortNameServer] class.
var (
	MessagePortNameServerClass     _MessagePortNameServerClass
	MessagePortNameServerClassOnce sync.Once
)

func getMessagePortNameServerClass() _MessagePortNameServerClass {
	MessagePortNameServerClassOnce.Do(func() {
		MessagePortNameServerClass = _MessagePortNameServerClass{objc.GetClass("NSMessagePortNameServer")}
	})
	return MessagePortNameServerClass
}

type _MessagePortNameServerClass struct {
	class objc.Class
}

// An interface definition for the [MessagePortNameServer] class.
type IMessagePortNameServer interface {
	IPortNameServer
	PortForName(name string) unsafe.Pointer
	PortForNameHost(name string, host string) unsafe.Pointer
}

// A server takes and returns message ports.
//
// This port name server takes and returns instances of . Port removal functionality is not supported in ; if you want to cancel a service, you have to destroy the port (invalidate the object given to ).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMessagePortNameServer
type MessagePortNameServer struct {
	PortNameServer
}

// MessagePortNameServerFrom constructs a [MessagePortNameServer] from an unsafe.Pointer.
//
// A server takes and returns message ports.
func MessagePortNameServerFrom(ptr unsafe.Pointer) MessagePortNameServer {
	return MessagePortNameServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MessagePortNameServerClass) Alloc() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MessagePortNameServerClass) New() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MessagePortNameServer) Init() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MessagePortNameServer) Autorelease() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMessagePortNameServer creates a new MessagePortNameServer instance.
func NewMessagePortNameServer() MessagePortNameServer {
	return getMessagePortNameServerClass().New()
}


// Returns the object registered under a given name on the local host.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMessagePortNameServer/portForName:
func (m_ MessagePortNameServer) PortForName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("portForName:"), objc.String(name))
	return rv
}

// Returns the object registered under a given name on the local host.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMessagePortNameServer/portForName:host:
func (m_ MessagePortNameServer) PortForNameHost(name string, host string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("portForName:host:"), objc.String(name), objc.String(host))
	return rv
}



