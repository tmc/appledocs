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
	CurrentPath() NWPath
	SetCurrentPath(value INWPath)
	Endpoint() NWEndpoint
	SetEndpoint(value INWEndpoint)
	HasBetterPath() bool
	SetHasBetterPath(value bool)
	IsViable() bool
	SetIsViable(value bool)
	MaximumDatagramLength() int
	SetMaximumDatagramLength(value int)
	ResolvedEndpoint() NWEndpoint
	SetResolvedEndpoint(value INWEndpoint)
	State() unsafe.Pointer
	SetState(value unsafe.Pointer)
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


// The current evaluated path for the session’s
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/currentpath
func (n_ NWUDPSession) CurrentPath() NWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("currentPath"))
	return rv
}


// SetCurrentPath sets the value of the currentPath property.
// The current evaluated path for the session’s

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/currentpath
func (n_ NWUDPSession) SetCurrentPath(value INWPath) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrentPath:"), value)
}

// The destination endpoint with which this session was created.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/endpoint
func (n_ NWUDPSession) Endpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
// The destination endpoint with which this session was created.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/endpoint
func (n_ NWUDPSession) SetEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEndpoint:"), value)
}

// If a session has a better path, new session would use a different interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/hasbetterpath
func (n_ NWUDPSession) HasBetterPath() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasBetterPath"))
	return rv
}


// SetHasBetterPath sets the value of the hasBetterPath property.
// If a session has a better path, new session would use a different interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/hasbetterpath
func (n_ NWUDPSession) SetHasBetterPath(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasBetterPath:"), value)
}

// The viability of a UDP session represents whether or not data can be transferred.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/isviable
func (n_ NWUDPSession) IsViable() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isViable"))
	return rv
}


// SetIsViable sets the value of the isViable property.
// The viability of a UDP session represents whether or not data can be transferred.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/isviable
func (n_ NWUDPSession) SetIsViable(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsViable:"), value)
}

// The maximum size of a datagram to be written currently.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/maximumdatagramlength
func (n_ NWUDPSession) MaximumDatagramLength() int {
	rv := objc.Send[int](n_.ID, objc.Sel("maximumDatagramLength"))
	return rv
}


// SetMaximumDatagramLength sets the value of the maximumDatagramLength property.
// The maximum size of a datagram to be written currently.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/maximumdatagramlength
func (n_ NWUDPSession) SetMaximumDatagramLength(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumDatagramLength:"), value)
}

// The currently targeted remote endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/resolvedendpoint
func (n_ NWUDPSession) ResolvedEndpoint() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("resolvedEndpoint"))
	return rv
}


// SetResolvedEndpoint sets the value of the resolvedEndpoint property.
// The currently targeted remote endpoint.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/resolvedendpoint
func (n_ NWUDPSession) SetResolvedEndpoint(value INWEndpoint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setResolvedEndpoint:"), value)
}

// The current state of the UDP session.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/state
func (n_ NWUDPSession) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The current state of the UDP session.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwudpsession/state
func (n_ NWUDPSession) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setState:"), value)
}




