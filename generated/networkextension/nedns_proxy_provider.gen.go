// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEDNSProxyProvider] class.
var (
	NEDNSProxyProviderClass     _NEDNSProxyProviderClass
	NEDNSProxyProviderClassOnce sync.Once
)

func getNEDNSProxyProviderClass() _NEDNSProxyProviderClass {
	NEDNSProxyProviderClassOnce.Do(func() {
		NEDNSProxyProviderClass = _NEDNSProxyProviderClass{objc.GetClass("NEDNSProxyProvider")}
	})
	return NEDNSProxyProviderClass
}

type _NEDNSProxyProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSProxyProvider] class.
type INEDNSProxyProvider interface {
	INEProvider
	// properties:
	SystemDNSSettings() INEDNSSettings
	SetSystemDNSSettings(value INEDNSSettings)
	// methods:
	HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint unsafe.Pointer) bool
}

// The principal class for a DNS proxy provider app extension.
//
// A DNS proxy allows your app to intercept all DNS traffic generated on a device. You can use this capability to provide services like DNS traffic encryption, typically by redirecting DNS traffic to your own server. You usually do this in the context of managed devices, such as those owned by a school or an enterprise. You create a DNS proxy as an app extension based on a custom subclass of the class. Once active, the proxy receives access to flows of DNS traffic in the form of instances. Each flow corresponds to a socket opened by an app to UDP port 53 or TCP port 53. Your DNS proxy provider acts as a transparent DNS proxy for the flows of network data that it receives. When you subclass , you must provide implementations for the following methods:


// The principal class for a DNS proxy provider app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider
type NEDNSProxyProvider struct {
	NEProvider
}

// NEDNSProxyProviderFrom constructs a [NEDNSProxyProvider] from an unsafe.Pointer.
//
// The principal class for a DNS proxy provider app extension.
func NEDNSProxyProviderFrom(ptr unsafe.Pointer) NEDNSProxyProvider {
	return NEDNSProxyProvider{
		NEProvider: NEProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyProviderClass) Alloc() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSProxyProviderClass) New() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSProxyProvider) Init() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSProxyProvider) Autorelease() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyProvider creates a new NEDNSProxyProvider instance.
func NewNEDNSProxyProvider() NEDNSProxyProvider {
	return getNEDNSProxyProviderClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/handleNewUDPFlow:initialRemoteFlowEndpoint:
func (n_ NEDNSProxyProvider) HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("handleNewUDPFlow:initialRemoteFlowEndpoint:"), flow, remoteEndpoint)
	return rv
}


// The current system DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyprovider/systemdnssettings
func (n_ NEDNSProxyProvider) SystemDNSSettings() INEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("systemDNSSettings"))
	return rv
}


// The current system DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyprovider/systemdnssettings
func (n_ NEDNSProxyProvider) SetSystemDNSSettings(value INEDNSSettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSystemDNSSettings:"), value)
}



