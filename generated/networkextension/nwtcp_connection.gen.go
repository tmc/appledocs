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
	

	// properties:
	ConnectedPath() INWPath
	Endpoint() INWEndpoint
	Error() foundation.foundation.INSError
	HasBetterPath() bool
	Viable() bool
	LocalAddress() INWEndpoint
	RemoteAddress() INWEndpoint
	State() NWTCPConnectionState
	TxtRecord() foundation.foundation.INSData
	IsViable() bool
	SetIsViable(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NWTCPConnectionClass) Alloc() NWTCPConnection {
	rv := objc.Send[NWTCPConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object to manage a TCP connection, with or without TLS.


// An object to manage a TCP connection, with or without TLS.
//
// [Full Topic]
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






// This convenience initializer can be used to create a new connection that will only be connected if there exists a better path (as determined by the system) to the remote endpoint of the original connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/init(upgradeFor:)
func NewNWTCPConnectionWithUpgradeForConnection(connection INWTCPConnection) NWTCPConnection {
	instance := getNWTCPConnectionClass().Alloc()
	rv := objc.Send[NWTCPConnection](instance.ID, objc.Sel("initWithUpgradeForConnection:"), connection)
	rv.Autorelease()
	return rv
}






















// The network path over which the connection was established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/connectedPath
func (n_ NWTCPConnection) ConnectedPath() INWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("connectedPath"))
	return rv
}


// The destination endpoint with which this connection was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/endpoint
func (n_ NWTCPConnection) Endpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("endpoint"))
	return rv
}


// The connection-wide error property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/error
func (n_ NWTCPConnection) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](n_.ID, objc.Sel("error"))
	return rv
}


// If a connection has a better path, new connections would use a different interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/hasBetterPath
func (n_ NWTCPConnection) HasBetterPath() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasBetterPath"))
	return rv
}


// The viability of a TCP connection indicates whether or not data can be transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/isViable
func (n_ NWTCPConnection) Viable() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("viable"))
	return rv
}


// The IP address endpoint from which the connection was established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/localAddress
func (n_ NWTCPConnection) LocalAddress() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localAddress"))
	return rv
}


// The IP address endpoint to which the connection was established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/remoteAddress
func (n_ NWTCPConnection) RemoteAddress() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteAddress"))
	return rv
}


// The status of the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/state
func (n_ NWTCPConnection) State() NWTCPConnectionState {
	rv := objc.Send[NWTCPConnectionState](n_.ID, objc.Sel("state"))
	return rv
}


// The TXT record associated with a connected Bonjour service endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWTCPConnection/txtRecord
func (n_ NWTCPConnection) TxtRecord() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("txtRecord"))
	return rv
}


// The viability of a TCP connection indicates whether or not data can be transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/isviable
func (n_ NWTCPConnection) IsViable() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isViable"))
	return rv
}


// The viability of a TCP connection indicates whether or not data can be transferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nwtcpconnection/isviable
func (n_ NWTCPConnection) SetIsViable(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsViable:"), value)
}







