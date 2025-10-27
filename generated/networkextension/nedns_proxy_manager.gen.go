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
	

	// properties:
	Enabled() bool
	SetEnabled(value bool)
	LocalizedDescription() foundation.foundation.INSString
	SetLocalizedDescription(value foundation.foundation.INSString)
	ProviderProtocol() INEDNSProxyProviderProtocol
	SetProviderProtocol(value INEDNSProxyProviderProtocol)
	NEDNSProxyErrorDomain() foundation.foundation.INSString
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyManagerClass) Alloc() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object to create and manage an DNS proxy provider’s configuration.
//
// A DNS proxy allows your app to intercept all DNS traffic generated on a device. You can use this capability to provide services like DNS traffic encryption, typically by redirecting DNS traffic to your own server. You usually do this in the context of managed devices, such as those owned by a school or an enterprise. You create a DNS proxy as an app extension based on a custom subclass of the class. You enable and configure this proxy from within your app using the singleton proxy manager instance provided by the type method of the class. For example, for a proxy that performs a simple redirect, you can use the proxy manager to define and dynamically configure the destination IP address of the redirected traffic. Instances of the proxy manager are thread safe.


// An object to create and manage an DNS proxy provider’s configuration.
//
// [Full Topic]
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










// Returns a singleton DNS proxy manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/shared()
func (nc _NEDNSProxyManagerClass) SharedManager() NEDNSProxyManager {
	rv := objc.Send[NEDNSProxyManager](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}












// Loads the current DNS proxy configuration from the caller’s DNS proxy preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/loadFromPreferences(completionHandler:)
func (n_ NEDNSProxyManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Removes the DNS proxy configuration from the caller’s DNS proxy preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/removeFromPreferences(completionHandler:)
func (n_ NEDNSProxyManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Saves the DNS proxy configuration in the caller’s DNS proxy preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/saveToPreferences(completionHandler:)
func (n_ NEDNSProxyManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}







// The status of a DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/isEnabled
func (n_ NEDNSProxyManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}


// The status of a DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/isEnabled
func (n_ NEDNSProxyManager) SetEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnabled:"), value)
}


// A description of the DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/localizedDescription
func (n_ NEDNSProxyManager) LocalizedDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A description of the DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/localizedDescription
func (n_ NEDNSProxyManager) SetLocalizedDescription(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// The provider-specific portion of the DNS proxy configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/providerProtocol
func (n_ NEDNSProxyManager) ProviderProtocol() INEDNSProxyProviderProtocol {
	rv := objc.Send[NEDNSProxyProviderProtocol](n_.ID, objc.Sel("providerProtocol"))
	return rv
}


// The provider-specific portion of the DNS proxy configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyManager/providerProtocol
func (n_ NEDNSProxyManager) SetProviderProtocol(value INEDNSProxyProviderProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderProtocol:"), value)
}


// The DNS proxy error domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxyerrordomain
func (n_ NEDNSProxyManager) NEDNSProxyErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEDNSProxyErrorDomain"))
	return rv
}


// The status of a DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxymanager/isenabled
func (n_ NEDNSProxyManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// The status of a DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednsproxymanager/isenabled
func (n_ NEDNSProxyManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}








