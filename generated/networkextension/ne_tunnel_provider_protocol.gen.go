// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETunnelProviderProtocol] class.
var (
	NETunnelProviderProtocolClass     _NETunnelProviderProtocolClass
	NETunnelProviderProtocolClassOnce sync.Once
)

func getNETunnelProviderProtocolClass() _NETunnelProviderProtocolClass {
	NETunnelProviderProtocolClassOnce.Do(func() {
		NETunnelProviderProtocolClass = _NETunnelProviderProtocolClass{objc.GetClass("NETunnelProviderProtocol")}
	})
	return NETunnelProviderProtocolClass
}

type _NETunnelProviderProtocolClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelProviderProtocol] class.
type INETunnelProviderProtocol interface {
	INEVPNProtocol
}

// Configuration parameters for a VPN tunnel.
//
// objects are used to specify configuration parameters for Tunnel Provider extensions.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getNETunnelProviderProtocolClass().New()
}


// A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (n_ NETunnelProviderProtocol) ProviderBundleIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// SetProviderBundleIdentifier sets the value of the providerBundleIdentifier property.
// A string identifying the specific Tunnel Provider extension that should be used with this configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (n_ NETunnelProviderProtocol) SetProviderBundleIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}
// A dictionary containing keys and values defined by the Tunnel Provider developer.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (n_ NETunnelProviderProtocol) ProviderConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// SetProviderConfiguration sets the value of the providerConfiguration property.
// A dictionary containing keys and values defined by the Tunnel Provider developer.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (n_ NETunnelProviderProtocol) SetProviderConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}


