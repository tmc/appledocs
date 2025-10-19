// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETunnelProviderManager] class.
var (
	nETunnelProviderManagerClass     _NETunnelProviderManagerClass
	nETunnelProviderManagerClassOnce sync.Once
)

func getNETunnelProviderManagerClass() _NETunnelProviderManagerClass {
	nETunnelProviderManagerClassOnce.Do(func() {
		nETunnelProviderManagerClass = _NETunnelProviderManagerClass{objc.GetClass("NETunnelProviderManager")}
	})
	return nETunnelProviderManagerClass
}

type _NETunnelProviderManagerClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelProviderManager] class.
type INETunnelProviderManager interface {
	INEVPNManager
}

// An object to create and manage the tunnel provider’s VPN configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager
type NETunnelProviderManager struct {
	NEVPNManager
}

// NETunnelProviderManagerFrom constructs a [NETunnelProviderManager] from an unsafe.Pointer.
//
// An object to create and manage the tunnel provider’s VPN configuration.
func NETunnelProviderManagerFrom(ptr unsafe.Pointer) NETunnelProviderManager {
	return NETunnelProviderManager{
		NEVPNManager: NEVPNManagerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderManagerClass) Alloc() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETunnelProviderManagerClass) New() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProviderManager) Init() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProviderManager) Autorelease() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderManager creates a new NETunnelProviderManager instance.
func NewNETunnelProviderManager() NETunnelProviderManager {
	return getNETunnelProviderManagerClass().New()
}




