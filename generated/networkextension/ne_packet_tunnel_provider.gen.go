// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SetPacketFlow(value INEPacketTunnelFlow)
	VirtualInterface() unsafe.Pointer
	SetVirtualInterface(value unsafe.Pointer)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (nc _NEPacketTunnelProviderClass) Alloc() NEPacketTunnelProvider {
	rv := objc.Send[NEPacketTunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelprovider/packetflow
func (n_ NEPacketTunnelProvider) PacketFlow() INEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](n_.ID, objc.Sel("packetFlow"))
	return rv
}


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelprovider/packetflow
func (n_ NEPacketTunnelProvider) SetPacketFlow(value INEPacketTunnelFlow) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPacketFlow:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelprovider/virtualinterface-7l3ol
func (n_ NEPacketTunnelProvider) VirtualInterface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("virtualInterface"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelprovider/virtualinterface-7l3ol
func (n_ NEPacketTunnelProvider) SetVirtualInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setVirtualInterface:"), value)
}



