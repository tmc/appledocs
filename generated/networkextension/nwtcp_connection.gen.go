// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NWTCPConnection] class.
var (
	NWTCPConnectionClass     _NWTCPConnectionClass
	NWTCPConnectionClassOnce sync.Once
)

func getNWTCPConnectionClass() _NWTCPConnectionClass {
	NWTCPConnectionClassOnce.Do(func() {
		NWTCPConnectionClass = _NWTCPConnectionClass{objc.GetClass("NWTCPConnection")}
	})
	return NWTCPConnectionClass
}

type _NWTCPConnectionClass struct {
	class objc.Class
}

// An interface definition for the [NWTCPConnection] class.
type INWTCPConnection interface {
	objectivec.IObject
}

// An object to manage a TCP connection, with or without TLS.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection
type NWTCPConnection struct {
	objectivec.Object
}

// NWTCPConnectionFrom constructs a [NWTCPConnection] from an unsafe.Pointer.
//
// An object to manage a TCP connection, with or without TLS.
func NWTCPConnectionFrom(ptr unsafe.Pointer) NWTCPConnection {
	return NWTCPConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWTCPConnectionClass) Alloc() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWTCPConnectionClass) New() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWTCPConnection) Init() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWTCPConnection) Autorelease() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWTCPConnection creates a new NWTCPConnection instance.
func NewNWTCPConnection() NWTCPConnection {
	return getNWTCPConnectionClass().New()
}


// The status of the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/state
func (n_ NWTCPConnection) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The status of the connection.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/state
func (n_ NWTCPConnection) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setState:"), value)
}

// The viability of a TCP connection indicates whether or not data can be transferred.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/isviable
func (n_ NWTCPConnection) IsViable() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isViable"))
	return rv
}


// SetIsViable sets the value of the isViable property.
// The viability of a TCP connection indicates whether or not data can be transferred.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/isviable
func (n_ NWTCPConnection) SetIsViable(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsViable:"), value)
}

// The IP address endpoint to which the connection was established.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/remoteaddress
func (n_ NWTCPConnection) RemoteAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("remoteAddress"))
	return rv
}


// SetRemoteAddress sets the value of the remoteAddress property.
// The IP address endpoint to which the connection was established.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/remoteaddress
func (n_ NWTCPConnection) SetRemoteAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRemoteAddress:"), value)
}

// The network path over which the connection was established.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/connectedpath
func (n_ NWTCPConnection) ConnectedPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("connectedPath"))
	return rv
}


// SetConnectedPath sets the value of the connectedPath property.
// The network path over which the connection was established.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/connectedpath
func (n_ NWTCPConnection) SetConnectedPath(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnectedPath:"), value)
}

// The destination endpoint with which this connection was created.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/endpoint
func (n_ NWTCPConnection) Endpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
// The destination endpoint with which this connection was created.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/endpoint
func (n_ NWTCPConnection) SetEndpoint(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEndpoint:"), value)
}

// The IP address endpoint from which the connection was established.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/localaddress
func (n_ NWTCPConnection) LocalAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("localAddress"))
	return rv
}


// SetLocalAddress sets the value of the localAddress property.
// The IP address endpoint from which the connection was established.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/localaddress
func (n_ NWTCPConnection) SetLocalAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalAddress:"), value)
}

// The connection-wide error property.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/error
func (n_ NWTCPConnection) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
// The connection-wide error property.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/error
func (n_ NWTCPConnection) SetError(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setError:"), value)
}

// The TXT record associated with a connected Bonjour service endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/txtrecord
func (n_ NWTCPConnection) TxtRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("txtRecord"))
	return rv
}


// SetTxtRecord sets the value of the txtRecord property.
// The TXT record associated with a connected Bonjour service endpoint.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/txtrecord
func (n_ NWTCPConnection) SetTxtRecord(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTxtRecord:"), value)
}

// If a connection has a better path, new connections would use a different interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/hasbetterpath
func (n_ NWTCPConnection) HasBetterPath() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasBetterPath"))
	return rv
}


// SetHasBetterPath sets the value of the hasBetterPath property.
// If a connection has a better path, new connections would use a different interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/hasbetterpath
func (n_ NWTCPConnection) SetHasBetterPath(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasBetterPath:"), value)
}



