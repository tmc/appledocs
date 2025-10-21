// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETransparentProxyProvider] class.
var (
	NETransparentProxyProviderClass     _NETransparentProxyProviderClass
	NETransparentProxyProviderClassOnce sync.Once
)

func getNETransparentProxyProviderClass() _NETransparentProxyProviderClass {
	NETransparentProxyProviderClassOnce.Do(func() {
		NETransparentProxyProviderClass = _NETransparentProxyProviderClass{objc.GetClass("NETransparentProxyProvider")}
	})
	return NETransparentProxyProviderClass
}

type _NETransparentProxyProviderClass struct {
	class objc.Class
}

// An interface definition for the [NETransparentProxyProvider] class.
type INETransparentProxyProvider interface {
	INEAppProxyProvider
}

// An object that implements the client side of a custom transparent network proxy solution.
//
// The class has the following behavior differences from its superclass : Returning from and causes the flow to proceed to communicate directly with the flow’s ultimate destination, instead of closing the flow with a “Connection Refused” error. This provider ignores and specified within . Flows that match the within use the same DNS and proxy settings that other flows on the system currently use. Flows that are created using a “connect by name” API (such as framework or ) that match the don’t bypass DNS resolution.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyProvider
type NETransparentProxyProvider struct {
	NEAppProxyProvider
}

// NETransparentProxyProviderFrom constructs a [NETransparentProxyProvider] from an unsafe.Pointer.
//
// An object that implements the client side of a custom transparent network proxy solution.
func NETransparentProxyProviderFrom(ptr unsafe.Pointer) NETransparentProxyProvider {
	return NETransparentProxyProvider{
		NEAppProxyProvider: NEAppProxyProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NETransparentProxyProviderClass) Alloc() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETransparentProxyProviderClass) New() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETransparentProxyProvider) Init() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETransparentProxyProvider) Autorelease() NETransparentProxyProvider {
	rv := objc.Send[NETransparentProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyProvider creates a new NETransparentProxyProvider instance.
func NewNETransparentProxyProvider() NETransparentProxyProvider {
	return getNETransparentProxyProviderClass().New()
}


// An array of rules that collectively specify what traffic to route through the transparent proxy.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netransparentproxynetworksettings/includednetworkrules
func (n_ NETransparentProxyProvider) IncludedNetworkRules() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("includedNetworkRules"))
	return rv
}


// SetIncludedNetworkRules sets the value of the includedNetworkRules property.
// An array of rules that collectively specify what traffic to route through the transparent proxy.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netransparentproxynetworksettings/includednetworkrules
func (n_ NETransparentProxyProvider) SetIncludedNetworkRules(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedNetworkRules:"), value)
}



