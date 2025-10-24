// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEAppPushManager] class.
var (
	NEAppPushManagerClass     _NEAppPushManagerClass
	NEAppPushManagerClassOnce sync.Once
)

func getNEAppPushManagerClass() _NEAppPushManagerClass {
	NEAppPushManagerClassOnce.Do(func() {
		NEAppPushManagerClass = _NEAppPushManagerClass{objc.GetClass("NEAppPushManager")}
	})
	return NEAppPushManagerClass
}

type _NEAppPushManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEAppPushManager] class.
type INEAppPushManager interface {
	objectivec.IObject
	// properties:
	NEAppPushErrorDomain() objc.IObject /* cross-framework: NSString */
	IsActive() bool
	SetIsActive(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */)
	MatchEthernet() bool
	SetMatchEthernet(value bool)
	MatchPrivateLTENetworks() objc.IObject /* cross-framework: NEPrivateLTENetwork */
	SetMatchPrivateLTENetworks(value objc.IObject /* cross-framework: NEPrivateLTENetwork */)
	ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */
	SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */)
	ProviderConfiguration() objc.IObject /* cross-framework: NSString */
	SetProviderConfiguration(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An object that configures a push provider and manages its life cycle.
//
// Your app can create as many instances as you need. Load your managers from the persistent store and set up their delegates immediately after the app launches, so they’re ready to handle incoming calls.


// An object that configures a push provider and manages its life cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager
type NEAppPushManager struct {
	objectivec.Object
}

// NEAppPushManagerFrom constructs a [NEAppPushManager] from an unsafe.Pointer.
//
// An object that configures a push provider and manages its life cycle.
func NEAppPushManagerFrom(ptr unsafe.Pointer) NEAppPushManager {
	return NEAppPushManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppPushManagerClass) Alloc() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppPushManagerClass) New() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppPushManager) Init() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppPushManager) Autorelease() NEAppPushManager {
	rv := objc.Send[NEAppPushManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppPushManager creates a new NEAppPushManager instance.
func NewNEAppPushManager() NEAppPushManager {
	return getNEAppPushManagerClass().New()
}



// The error domain string for local push errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppusherrordomain
func (n_ NEAppPushManager) NEAppPushErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEAppPushErrorDomain"))
	return rv
}


// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isactive
func (n_ NEAppPushManager) IsActive() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isactive
func (n_ NEAppPushManager) SetIsActive(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsActive:"), value)
}


// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isenabled
func (n_ NEAppPushManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/isenabled
func (n_ NEAppPushManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}


// A string that contains the localized description of the app push manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/localizeddescription
func (n_ NEAppPushManager) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string that contains the localized description of the app push manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/localizeddescription
func (n_ NEAppPushManager) SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// A property that indicates Ethernet support for Local Push Connectivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchethernet
func (n_ NEAppPushManager) MatchEthernet() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("matchEthernet"))
	return rv
}


// A property that indicates Ethernet support for Local Push Connectivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchethernet
func (n_ NEAppPushManager) SetMatchEthernet(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchEthernet:"), value)
}


// An array of private LTE networks that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchprivateltenetworks
func (n_ NEAppPushManager) MatchPrivateLTENetworks() objc.IObject /* cross-framework: NEPrivateLTENetwork */ {
	rv := objc.Send[NEPrivateLTENetwork](n_.ID, objc.Sel("matchPrivateLTENetworks"))
	return rv
}


// An array of private LTE networks that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/matchprivateltenetworks
func (n_ NEAppPushManager) SetMatchPrivateLTENetworks(value objc.IObject /* cross-framework: NEPrivateLTENetwork */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPrivateLTENetworks:"), value)
}


// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerbundleidentifier
func (n_ NEAppPushManager) ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}


// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerbundleidentifier
func (n_ NEAppPushManager) SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderBundleIdentifier:"), value)
}


// A dictionary that contains vendor-specific key-value pairs, that you use to configure a provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerconfiguration
func (n_ NEAppPushManager) ProviderConfiguration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}


// A dictionary that contains vendor-specific key-value pairs, that you use to configure a provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapppushmanager/providerconfiguration
func (n_ NEAppPushManager) SetProviderConfiguration(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProviderConfiguration:"), value)
}


