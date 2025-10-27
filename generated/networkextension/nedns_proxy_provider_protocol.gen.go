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
	

	// properties:
	ProviderBundleIdentifier() foundation.foundation.INSString
	SetProviderBundleIdentifier(value foundation.foundation.INSString)
	ProviderConfiguration() foundation.IDictionary
	SetProviderConfiguration(value foundation.IDictionary)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyProviderProtocolClass) Alloc() NEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// Configuration parameters for a DNS proxy.


// Configuration parameters for a DNS proxy.
//
// [Full Topic]
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

























// A string containing the bundle identifier of the proxy provider to be used by this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerBundleIdentifier
func (n_ NEDNSProxyProviderProtocol) ProviderBundleIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// A string containing the bundle identifier of the proxy provider to be used by this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerBundleIdentifier
func (n_ NEDNSProxyProviderProtocol) SetProviderBundleIdentifier(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}


// A dictionary containing vendor-specific configuration parameters for a proxy provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerConfiguration
func (n_ NEDNSProxyProviderProtocol) ProviderConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// A dictionary containing vendor-specific configuration parameters for a proxy provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProviderProtocol/providerConfiguration
func (n_ NEDNSProxyProviderProtocol) SetProviderConfiguration(value foundation.IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}








