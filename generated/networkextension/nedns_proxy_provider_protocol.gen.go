// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEDNSProxyProviderProtocol] class.
var (
	NEDNSProxyProviderProtocolClass     _NEDNSProxyProviderProtocolClass
	NEDNSProxyProviderProtocolClassOnce sync.Once
)

func getNEDNSProxyProviderProtocolClass() _NEDNSProxyProviderProtocolClass {
	NEDNSProxyProviderProtocolClassOnce.Do(func() {
		NEDNSProxyProviderProtocolClass = _NEDNSProxyProviderProtocolClass{objc.GetClass("NEDNSProxyProviderProtocol")}
	})
	return NEDNSProxyProviderProtocolClass
}

type _NEDNSProxyProviderProtocolClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSProxyProviderProtocol] class.
type INEDNSProxyProviderProtocol interface {
	INEVPNProtocol
}

// Configuration parameters for a DNS proxy.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol
type NEDNSProxyProviderProtocol struct {
	NEVPNProtocol
}

// NEDNSProxyProviderProtocolFrom constructs a [NEDNSProxyProviderProtocol] from an unsafe.Pointer.
//
// Configuration parameters for a DNS proxy.
func NEDNSProxyProviderProtocolFrom(ptr unsafe.Pointer) NEDNSProxyProviderProtocol {
	return NEDNSProxyProviderProtocol{
		NEVPNProtocol: NEVPNProtocolFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyProviderProtocolClass) Alloc() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSProxyProviderProtocolClass) New() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSProxyProviderProtocol) Init() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSProxyProviderProtocol) Autorelease() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyProviderProtocol creates a new NEDNSProxyProviderProtocol instance.
func NewNEDNSProxyProviderProtocol() NEDNSProxyProviderProtocol {
	return getNEDNSProxyProviderProtocolClass().New()
}


// A string containing the bundle identifier of the proxy provider to be used by this configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyproviderprotocol/providerbundleidentifier
func (n_ NEDNSProxyProviderProtocol) ProviderBundleIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// SetProviderBundleIdentifier sets the value of the providerBundleIdentifier property.
// A string containing the bundle identifier of the proxy provider to be used by this configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyproviderprotocol/providerbundleidentifier
func (n_ NEDNSProxyProviderProtocol) SetProviderBundleIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), objc.String(value))
}

// A dictionary containing vendor-specific configuration parameters for a proxy provider.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyproviderprotocol/providerconfiguration
func (n_ NEDNSProxyProviderProtocol) ProviderConfiguration() string {
	rv := objc.Send[string](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// SetProviderConfiguration sets the value of the providerConfiguration property.
// A dictionary containing vendor-specific configuration parameters for a proxy provider.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyproviderprotocol/providerconfiguration
func (n_ NEDNSProxyProviderProtocol) SetProviderConfiguration(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), objc.String(value))
}



