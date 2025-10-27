// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEEthernetTunnelProvider] class.
var (
	NEEthernetTunnelProviderClass     _NEEthernetTunnelProviderClass
	NEEthernetTunnelProviderClassOnce sync.Once
)

func getNEEthernetTunnelProviderClass() _NEEthernetTunnelProviderClass {
	NEEthernetTunnelProviderClassOnce.Do(func() {
		NEEthernetTunnelProviderClass = _NEEthernetTunnelProviderClass{objc.GetClass("NEEthernetTunnelProvider")}
	})
	return NEEthernetTunnelProviderClass
}

type _NEEthernetTunnelProviderClass struct {
	class objc.Class
}





// An interface definition for the [NEEthernetTunnelProvider] class.
type INEEthernetTunnelProvider interface {
	INEPacketTunnelProvider
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEEthernetTunnelProviderClass) Alloc() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEEthernetTunnelProviderClass) New() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEEthernetTunnelProvider) Init() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEEthernetTunnelProvider) Autorelease() NEEthernetTunnelProvider {
	rv := objc.Send[NEEthernetTunnelProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEEthernetTunnelProvider creates a new NEEthernetTunnelProvider instance.
func NewNEEthernetTunnelProvider() NEEthernetTunnelProvider {
	return getNEEthernetTunnelProviderClass().New()
}





// A type that implements the client side of a custom link-layer packet tunneling protocol.


// A type that implements the client side of a custom link-layer packet tunneling protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelProvider
type NEEthernetTunnelProvider struct {
	NEPacketTunnelProvider
}

// NEEthernetTunnelProviderFrom constructs a [NEEthernetTunnelProvider] from an unsafe.Pointer.
//
// A type that implements the client side of a custom link-layer packet tunneling protocol.
func NEEthernetTunnelProviderFrom(ptr unsafe.Pointer) NEEthernetTunnelProvider {
	return NEEthernetTunnelProvider{
		NEPacketTunnelProvider: NEPacketTunnelProviderFrom(ptr),
	}
}































