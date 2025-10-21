// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEAppProxyProviderManager] class.
var (
	NEAppProxyProviderManagerClass     _NEAppProxyProviderManagerClass
	NEAppProxyProviderManagerClassOnce sync.Once
)

func getNEAppProxyProviderManagerClass() _NEAppProxyProviderManagerClass {
	NEAppProxyProviderManagerClassOnce.Do(func() {
		NEAppProxyProviderManagerClass = _NEAppProxyProviderManagerClass{objc.GetClass("NEAppProxyProviderManager")}
	})
	return NEAppProxyProviderManagerClass
}

type _NEAppProxyProviderManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEAppProxyProviderManager] class.
type INEAppProxyProviderManager interface {
	INETunnelProviderManager
}

// An object to create and manage the app proxy provider’s VPN configuration.
//
// Objects cannot be directly instantiated. Instead, App Proxy configurations are created exclusively from payloads in configuration profiles. App Proxy configurations can only be used with Per-App VPN routing rules. For more details about how to create App Proxy configurations and configure Per-App VPN, see .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProviderManager
type NEAppProxyProviderManager struct {
	NETunnelProviderManager
}

// NEAppProxyProviderManagerFrom constructs a [NEAppProxyProviderManager] from an unsafe.Pointer.
//
// An object to create and manage the app proxy provider’s VPN configuration.
func NEAppProxyProviderManagerFrom(ptr unsafe.Pointer) NEAppProxyProviderManager {
	return NEAppProxyProviderManager{
		NETunnelProviderManager: NETunnelProviderManagerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyProviderManagerClass) Alloc() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppProxyProviderManagerClass) New() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyProviderManager) Init() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyProviderManager) Autorelease() NEAppProxyProviderManager {
	rv := objc.Send[NEAppProxyProviderManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyProviderManager creates a new NEAppProxyProviderManager instance.
func NewNEAppProxyProviderManager() NEAppProxyProviderManager {
	return getNEAppProxyProviderManagerClass().New()
}


// Load all of the App Proxy configurations associated with the calling app that have previously been saved to the Network Extension preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProviderManager/loadAllFromPreferences(completionHandler:)
func (nc _NEAppProxyProviderManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}



