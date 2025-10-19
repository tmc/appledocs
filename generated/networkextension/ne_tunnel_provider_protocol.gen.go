// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETunnelProviderProtocol] class.
var nETunnelProviderProtocolClass = _NETunnelProviderProtocolClass{objc.GetClass("NETunnelProviderProtocol")}

type _NETunnelProviderProtocolClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelProviderProtocol] class.
type INETunnelProviderProtocol interface {
	INEVPNProtocol
}

// Configuration parameters for a VPN tunnel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol

type NETunnelProviderProtocol struct {
	NEVPNProtocol
}

// NETunnelProviderProtocolFrom constructs a [NETunnelProviderProtocol] from an unsafe.Pointer.
//
// Configuration parameters for a VPN tunnel.
func NETunnelProviderProtocolFrom(ptr unsafe.Pointer) NETunnelProviderProtocol {
	return NETunnelProviderProtocol{
		NEVPNProtocol: NEVPNProtocolFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderProtocolClass) Alloc() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NETunnelProviderProtocolClass) New() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProviderProtocol) Init() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProviderProtocol) Autorelease() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderProtocol creates a new NETunnelProviderProtocol instance.
func NewNETunnelProviderProtocol() NETunnelProviderProtocol {
	return nETunnelProviderProtocolClass.New()
}




