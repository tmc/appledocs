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
	

	// properties:
	ProviderBundleIdentifier() foundation.foundation.INSString
	SetProviderBundleIdentifier(value foundation.foundation.INSString)
	ProviderConfiguration() foundation.IDictionary
	SetProviderConfiguration(value foundation.IDictionary)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderProtocolClass) Alloc() NETunnelProviderProtocol {
	rv := objc.Send[NETunnelProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// Configuration parameters for a VPN tunnel.
//
// objects are used to specify configuration parameters for Tunnel Provider extensions.


// Configuration parameters for a VPN tunnel.
//
// [Full Topic]
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

























// A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (n_ NETunnelProviderProtocol) ProviderBundleIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// A string identifying the specific Tunnel Provider extension that should be used with this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerBundleIdentifier
func (n_ NETunnelProviderProtocol) SetProviderBundleIdentifier(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}


// A dictionary containing keys and values defined by the Tunnel Provider developer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (n_ NETunnelProviderProtocol) ProviderConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// A dictionary containing keys and values defined by the Tunnel Provider developer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderProtocol/providerConfiguration
func (n_ NETunnelProviderProtocol) SetProviderConfiguration(value foundation.IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}








