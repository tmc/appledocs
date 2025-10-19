// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEDNSProxyProviderProtocol] class.
var nEDNSProxyProviderProtocolClass = _NEDNSProxyProviderProtocolClass{objc.GetClass("NEDNSProxyProviderProtocol")}

type _NEDNSProxyProviderProtocolClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSProxyProviderProtocol] class.
type INEDNSProxyProviderProtocol interface {
	INEVPNProtocol
}

// Configuration parameters for a DNS proxy. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return nEDNSProxyProviderProtocolClass.New()
}




