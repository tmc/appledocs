// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NWUDPSession] class.
var (
	NWUDPSessionClass     _NWUDPSessionClass
	NWUDPSessionClassOnce sync.Once
)

func getNWUDPSessionClass() _NWUDPSessionClass {
	NWUDPSessionClassOnce.Do(func() {
		NWUDPSessionClass = _NWUDPSessionClass{objc.GetClass("NWUDPSession")}
	})
	return NWUDPSessionClass
}

type _NWUDPSessionClass struct {
	class objc.Class
}

// An interface definition for the [NWUDPSession] class.
type INWUDPSession interface {
	objectivec.IObject
}

// An object to manage a UDP session to a network endpoint.
//
// Since UDP does not include a handshake with the remote endpoint as part of its protocol, it is up to the client of the UDP session to provide feedback on the viability of the current endpoint. If a session is opened to a hostname, the system will resolve that hostname into potentially several IP addresses. Once the session state is , the client should try to write and read datagrams. If there is no response from the remote endpoint, the client can try the next address that was resolved using .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession
type NWUDPSession struct {
	objectivec.Object
}

// NWUDPSessionFrom constructs a [NWUDPSession] from an unsafe.Pointer.
//
// An object to manage a UDP session to a network endpoint.
func NWUDPSessionFrom(ptr unsafe.Pointer) NWUDPSession {
	return NWUDPSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWUDPSessionClass) Alloc() NWUDPSession {
	rv := objc.Send[NWUDPSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWUDPSessionClass) New() NWUDPSession {
	rv := objc.Send[NWUDPSession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWUDPSession) Init() NWUDPSession {
	rv := objc.Send[NWUDPSession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWUDPSession) Autorelease() NWUDPSession {
	rv := objc.Send[NWUDPSession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWUDPSession creates a new NWUDPSession instance.
func NewNWUDPSession() NWUDPSession {
	return getNWUDPSessionClass().New()
}




