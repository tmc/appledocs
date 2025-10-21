// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETunnelProviderSession] class.
var (
	NETunnelProviderSessionClass     _NETunnelProviderSessionClass
	NETunnelProviderSessionClassOnce sync.Once
)

func getNETunnelProviderSessionClass() _NETunnelProviderSessionClass {
	NETunnelProviderSessionClassOnce.Do(func() {
		NETunnelProviderSessionClass = _NETunnelProviderSessionClass{objc.GetClass("NETunnelProviderSession")}
	})
	return NETunnelProviderSessionClass
}

type _NETunnelProviderSessionClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelProviderSession] class.
type INETunnelProviderSession interface {
	INEVPNConnection
	SendProviderMessageReturnErrorResponseHandler(messageData unsafe.Pointer, error_ unsafe.Pointer, responseHandler unsafe.Pointer) bool
	StartTunnelWithOptionsAndReturnError(options unsafe.Pointer, error_ unsafe.Pointer) bool
	StopTunnel()
}

// An object to start and stop a tunnel connection and get its status.
//
// objects control network tunnel connections provided by Tunnel Provider extensions. objects are not instantiated directly. Instead, each object has an associated as a read-only property.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession
type NETunnelProviderSession struct {
	NEVPNConnection
}

// NETunnelProviderSessionFrom constructs a [NETunnelProviderSession] from an unsafe.Pointer.
//
// An object to start and stop a tunnel connection and get its status.
func NETunnelProviderSessionFrom(ptr unsafe.Pointer) NETunnelProviderSession {
	return NETunnelProviderSession{
		NEVPNConnection: NEVPNConnectionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderSessionClass) Alloc() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETunnelProviderSessionClass) New() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProviderSession) Init() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProviderSession) Autorelease() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderSession creates a new NETunnelProviderSession instance.
func NewNETunnelProviderSession() NETunnelProviderSession {
	return getNETunnelProviderSessionClass().New()
}


// Send a message to the Tunnel Provider extension. If the extension is not running, it should be launched to handle the message. If this method can’t start sending the message it reports an error in the parameter. If an error occurs while sending the message or returning the result, should be sent to the response handler as notification.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/sendProviderMessage(_:responseHandler:)
func (n_ NETunnelProviderSession) SendProviderMessageReturnErrorResponseHandler(messageData unsafe.Pointer, error_ unsafe.Pointer, responseHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("sendProviderMessage:returnError:responseHandler:"), messageData, error_, responseHandler)
	return rv
}

// Start the process of connecting the tunnel.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/startTunnel(options:)
func (n_ NETunnelProviderSession) StartTunnelWithOptionsAndReturnError(options unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("startTunnelWithOptions:andReturnError:"), options, error_)
	return rv
}

// Start the process of disconnecting the tunnel.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/stopTunnel()
func (n_ NETunnelProviderSession) StopTunnel() {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopTunnel"))
}



