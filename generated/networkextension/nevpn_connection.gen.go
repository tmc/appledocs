// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEVPNConnection] class.
var (
	NEVPNConnectionClass     _NEVPNConnectionClass
	NEVPNConnectionClassOnce sync.Once
)

func getNEVPNConnectionClass() _NEVPNConnectionClass {
	NEVPNConnectionClassOnce.Do(func() {
		NEVPNConnectionClass = _NEVPNConnectionClass{objc.GetClass("NEVPNConnection")}
	})
	return NEVPNConnectionClass
}

type _NEVPNConnectionClass struct {
	class objc.Class
}





// An interface definition for the [NEVPNConnection] class.
type INEVPNConnection interface {
	objectivec.IObject
	

	// properties:
	ConnectedDate() foundation.foundation.INSDate
	Manager() INEVPNManager
	Status() NEVPNStatus
	NEVPNConnectionErrorDomain() foundation.foundation.INSString
	NEVPNConnectionStartOptionPassword() foundation.foundation.INSString
	NEVPNConnectionStartOptionUsername() foundation.foundation.INSString


	

	// methods:
	FetchLastDisconnectErrorWithCompletionHandler(handler unsafe.Pointer)
	StartVPNTunnelAndReturnError(error_ foundation.foundation.INSError) bool
	StartVPNTunnelWithOptionsAndReturnError(options foundation.IDictionary, error_ foundation.foundation.INSError) bool
	StopVPNTunnel()


}





// Alloc allocates a new instance without initialization.
func (nc _NEVPNConnectionClass) Alloc() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEVPNConnectionClass) New() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNConnection) Init() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNConnection) Autorelease() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNConnection creates a new NEVPNConnection instance.
func NewNEVPNConnection() NEVPNConnection {
	return getNEVPNConnectionClass().New()
}





// An object to start and stop a Personal VPN connection and get its status.
//
// objects are not instantiated directly. Instead, each object has an associated object as a read-only property. The class provides methods for starting and stopping the VPN programmatically. The other way that the VPN can be started and stopped is through VPN On Demand. See the property in and . Instances of this class are thread safe.


// An object to start and stop a Personal VPN connection and get its status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection
type NEVPNConnection struct {
	objectivec.Object
}

// NEVPNConnectionFrom constructs a [NEVPNConnection] from an unsafe.Pointer.
//
// An object to start and stop a Personal VPN connection and get its status.
func NEVPNConnectionFrom(ptr unsafe.Pointer) NEVPNConnection {
	return NEVPNConnection{objectivec.Object{objc.ID(ptr)}}
}




















// Retrives the most recent error that caused the VPN to disconnect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/fetchLastDisconnectError(completionHandler:)
func (n_ NEVPNConnection) FetchLastDisconnectErrorWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("fetchLastDisconnectErrorWithCompletionHandler:"), handler)
}


// Start the process of connecting the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/startVPNTunnel()
func (n_ NEVPNConnection) StartVPNTunnelAndReturnError(error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("startVPNTunnelAndReturnError:"), error_)
	return rv
}


// Start the process of connecting the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/startVPNTunnel(options:)
func (n_ NEVPNConnection) StartVPNTunnelWithOptionsAndReturnError(options foundation.IDictionary, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("startVPNTunnelWithOptions:andReturnError:"), options, error_)
	return rv
}


// Start the process of disconnecting the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/stopVPNTunnel()
func (n_ NEVPNConnection) StopVPNTunnel() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopVPNTunnel"))
}







// The date and time when the connection status changed to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/connectedDate
func (n_ NEVPNConnection) ConnectedDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](n_.ID, objc.Sel("connectedDate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/manager
func (n_ NEVPNConnection) Manager() INEVPNManager {
	rv := objc.Send[NEVPNManager](n_.ID, objc.Sel("manager"))
	return rv
}


// The current status of the VPN connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/status
func (n_ NEVPNConnection) Status() NEVPNStatus {
	rv := objc.Send[NEVPNStatus](n_.ID, objc.Sel("status"))
	return rv
}


// The domain for errors resulting from VPN connection calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnconnectionerrordomain
func (n_ NEVPNConnection) NEVPNConnectionErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNConnectionErrorDomain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnconnectionstartoptionpassword
func (n_ NEVPNConnection) NEVPNConnectionStartOptionPassword() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNConnectionStartOptionPassword"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnconnectionstartoptionusername
func (n_ NEVPNConnection) NEVPNConnectionStartOptionUsername() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEVPNConnectionStartOptionUsername"))
	return rv
}








