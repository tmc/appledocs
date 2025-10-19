// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEPacketTunnelProvider] class.
var (
	nEPacketTunnelProviderClass     _NEPacketTunnelProviderClass
	nEPacketTunnelProviderClassOnce sync.Once
)

func getNEPacketTunnelProviderClass() _NEPacketTunnelProviderClass {
	nEPacketTunnelProviderClassOnce.Do(func() {
		nEPacketTunnelProviderClass = _NEPacketTunnelProviderClass{objc.GetClass("NEPacketTunnelProvider")}
	})
	return nEPacketTunnelProviderClass
}

type _NEPacketTunnelProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEPacketTunnelProvider] class.
type INEPacketTunnelProvider interface {
	INETunnelProvider
}

// The principal class for a packet tunnel provider app extension. [Full Topic]
//
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




