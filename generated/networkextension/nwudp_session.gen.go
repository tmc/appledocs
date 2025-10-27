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
	

	// properties:
	CurrentPath() INWPath
	Endpoint() INWEndpoint
	HasBetterPath() bool
	Viable() bool
	MaximumDatagramLength() uint
	ResolvedEndpoint() INWEndpoint
	State() NWUDPSessionState
	IsViable() bool
	SetIsViable(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NWUDPSessionClass) Alloc() NWUDPSession {
	rv := objc.Send[NWUDPSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object to manage a UDP session to a network endpoint.
//
// Since UDP does not include a handshake with the remote endpoint as part of its protocol, it is up to the client of the UDP session to provide feedback on the viability of the current endpoint. If a session is opened to a hostname, the system will resolve that hostname into potentially several IP addresses. Once the session state is , the client should try to write and read datagrams. If there is no response from the remote endpoint, the client can try the next address that was resolved using .


// An object to manage a UDP session to a network endpoint.
//
// [Full Topic]
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






// This convenience initializer can be used to create a new session based on the original session’s endpoint and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/init(upgradeFor:)
func NewNWUDPSessionWithUpgradeForSession(session INWUDPSession) NWUDPSession {
	instance := getNWUDPSessionClass().Alloc()
	rv := objc.Send[NWUDPSession](instance.ID, objc.Sel("initWithUpgradeForSession:"), session)
	rv.Autorelease()
	return rv
}






















// The current evaluated path for the session’s property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/currentPath
func (n_ NWUDPSession) CurrentPath() INWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("currentPath"))
	return rv
}


// The destination endpoint with which this session was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/endpoint
func (n_ NWUDPSession) Endpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("endpoint"))
	return rv
}


// If a session has a better path, new session would use a different interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/hasBetterPath
func (n_ NWUDPSession) HasBetterPath() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasBetterPath"))
	return rv
}


// The viability of a UDP session represents whether or not data can be transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/isViable
func (n_ NWUDPSession) Viable() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("viable"))
	return rv
}


// The maximum size of a datagram to be written currently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/maximumDatagramLength
func (n_ NWUDPSession) MaximumDatagramLength() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumDatagramLength"))
	return rv
}


// The currently targeted remote endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/resolvedEndpoint
func (n_ NWUDPSession) ResolvedEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("resolvedEndpoint"))
	return rv
}


// The current state of the UDP session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWUDPSession/state
func (n_ NWUDPSession) State() NWUDPSessionState {
	rv := objc.Send[NWUDPSessionState](n_.ID, objc.Sel("state"))
	return rv
}


// The viability of a UDP session represents whether or not data can be transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/isviable
func (n_ NWUDPSession) IsViable() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isViable"))
	return rv
}


// The viability of a UDP session represents whether or not data can be transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/isviable
func (n_ NWUDPSession) SetIsViable(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsViable:"), value)
}







