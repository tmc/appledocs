// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEPacketTunnelProvider] class.
var (
	NEPacketTunnelProviderClass     _NEPacketTunnelProviderClass
	NEPacketTunnelProviderClassOnce sync.Once
)

func getNEPacketTunnelProviderClass() _NEPacketTunnelProviderClass {
	NEPacketTunnelProviderClassOnce.Do(func() {
		NEPacketTunnelProviderClass = _NEPacketTunnelProviderClass{objc.GetClass("NEPacketTunnelProvider")}
	})
	return NEPacketTunnelProviderClass
}

type _NEPacketTunnelProviderClass struct {
	class objc.Class
}





// An interface definition for the [NEPacketTunnelProvider] class.
type INEPacketTunnelProvider interface {
	INETunnelProvider
	

	// properties:
	PacketFlow() INEPacketTunnelFlow
	VirtualInterface() objectivec.IObject


	

	// methods:
	CancelTunnelWithError(error_ foundation.foundation.INSError)
	StartTunnelWithOptionsCompletionHandler(options foundation.IDictionary, completionHandler unsafe.Pointer)
	StopTunnelWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (nc _NEPacketTunnelProviderClass) Alloc() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEPacketTunnelProviderClass) New() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEPacketTunnelProvider) Init() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEPacketTunnelProvider) Autorelease() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacketTunnelProvider creates a new NEPacketTunnelProvider instance.
func NewNEPacketTunnelProvider() NEPacketTunnelProvider {
	return getNEPacketTunnelProviderClass().New()
}





// The principal class for a packet tunnel provider app extension.
//
// The class gives its subclasses access to a virtual network interface via the property. Use the method in the Packet Tunnel Provider to specify that the following network settings be associated with the virtual interface: Virtual IP address DNS resolver configuration HTTP proxy configuration IP destination networks to be routed through the tunnel IP destination networks to be routed outside the tunnel Interface MTU By specifying IP destination networks, the Packet Tunnel Provider can dictate what IP destinations will be routed to the virtual interface. IP packets with matching destination addresses will then be diverted to Packet Tunnel Provider and can be read using the property. The Packet Tunnel Provider can then encapsulate the IP packets per a custom tunneling protocol and send them to a tunnel server. When the Packet Tunnel Provider decapsulates IP packets received from the tunnel server, it can use the property to inject the packets into the networking stack.


// The principal class for a packet tunnel provider app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider
type NEPacketTunnelProvider struct {
	NETunnelProvider
}

// NEPacketTunnelProviderFrom constructs a [NEPacketTunnelProvider] from an unsafe.Pointer.
//
// The principal class for a packet tunnel provider app extension.
func NEPacketTunnelProviderFrom(ptr unsafe.Pointer) NEPacketTunnelProvider {
	return NEPacketTunnelProvider{
		NETunnelProvider: NETunnelProviderFrom(ptr),
	}
}




















// Stop the network tunnel from the Packet Tunnel Provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/cancelTunnelWithError(_:)
func (n_ NEPacketTunnelProvider) CancelTunnelWithError(error_ foundation.foundation.INSError) {
	objc.Send[objc.ID](n_.ID, objc.Sel("cancelTunnelWithError:"), error_)
}


// Start the network tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/startTunnel(options:completionHandler:)
func (n_ NEPacketTunnelProvider) StartTunnelWithOptionsCompletionHandler(options foundation.IDictionary, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("startTunnelWithOptions:completionHandler:"), options, completionHandler)
}


// Stop the network tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/stopTunnel(with:completionHandler:)
func (n_ NEPacketTunnelProvider) StopTunnelWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopTunnelWithReason:completionHandler:"), reason, completionHandler)
}







// A object which is used to receive IP packets routed to the tunnel’s virtual interface and inject IP packets into the networking stack via the tunnel’s virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/packetFlow
func (n_ NEPacketTunnelProvider) PacketFlow() INEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](n_.ID, objc.Sel("packetFlow"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelProvider/virtualInterface-9fpgd
func (n_ NEPacketTunnelProvider) VirtualInterface() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("virtualInterface"))
	return rv
}








