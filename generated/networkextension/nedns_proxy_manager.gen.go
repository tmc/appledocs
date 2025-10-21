// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEDNSProxyManager] class.
var (
	NEDNSProxyManagerClass     _NEDNSProxyManagerClass
	NEDNSProxyManagerClassOnce sync.Once
)

func getNEDNSProxyManagerClass() _NEDNSProxyManagerClass {
	NEDNSProxyManagerClassOnce.Do(func() {
		NEDNSProxyManagerClass = _NEDNSProxyManagerClass{objc.GetClass("NEDNSProxyManager")}
	})
	return NEDNSProxyManagerClass
}

type _NEDNSProxyManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSProxyManager] class.
type INEDNSProxyManager interface {
	objectivec.IObject
}

// An object to create and manage an DNS proxy provider’s configuration.
//
// A DNS proxy allows your app to intercept all DNS traffic generated on a device. You can use this capability to provide services like DNS traffic encryption, typically by redirecting DNS traffic to your own server. You usually do this in the context of managed devices, such as those owned by a school or an enterprise. You create a DNS proxy as an app extension based on a custom subclass of the class. You enable and configure this proxy from within your app using the singleton proxy manager instance provided by the type method of the class. For example, for a proxy that performs a simple redirect, you can use the proxy manager to define and dynamically configure the destination IP address of the redirected traffic. Instances of the proxy manager are thread safe.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager
type NEDNSProxyManager struct {
	objectivec.Object
}

// NEDNSProxyManagerFrom constructs a [NEDNSProxyManager] from an unsafe.Pointer.
//
// An object to create and manage an DNS proxy provider’s configuration.
func NEDNSProxyManagerFrom(ptr unsafe.Pointer) NEDNSProxyManager {
	return NEDNSProxyManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyManagerClass) Alloc() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSProxyManagerClass) New() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSProxyManager) Init() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSProxyManager) Autorelease() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyManager creates a new NEDNSProxyManager instance.
func NewNEDNSProxyManager() NEDNSProxyManager {
	return getNEDNSProxyManagerClass().New()
}




