// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEDNSSettingsManager] class.
var (
	NEDNSSettingsManagerClass     _NEDNSSettingsManagerClass
	NEDNSSettingsManagerClassOnce sync.Once
)

func getNEDNSSettingsManagerClass() _NEDNSSettingsManagerClass {
	NEDNSSettingsManagerClassOnce.Do(func() {
		NEDNSSettingsManagerClass = _NEDNSSettingsManagerClass{objc.GetClass("NEDNSSettingsManager")}
	})
	return NEDNSSettingsManagerClass
}

type _NEDNSSettingsManagerClass struct {
	class objc.Class
}

// An interface definition for the [NEDNSSettingsManager] class.
type INEDNSSettingsManager interface {
	objectivec.IObject
	// properties:
	NEDNSSettingsErrorDomain() objc.IObject /* cross-framework: NSString */
	DnsSettings() INEDNSSettings
	SetDnsSettings(value INEDNSSettings)
	IsEnabled() bool
	SetIsEnabled(value bool)
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */)
	OnDemandRules() objc.IObject /* cross-framework: NEOnDemandRule */
	SetOnDemandRules(value objc.IObject /* cross-framework: NEOnDemandRule */)
	// methods:
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object you use to create and manage a DNS settings configuration.
//
// When your app starts up, access the shared instance of the DNS settings manager, and load existing settings from the preferences using . You can define your DNS server configuration, and persist it by calling . In order to use your DNS settings, the user needs to enable it in the Settings app on iOS or in System Preferences on macOS.


// An object you use to create and manage a DNS settings configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager
type NEDNSSettingsManager struct {
	objectivec.Object
}

// NEDNSSettingsManagerFrom constructs a [NEDNSSettingsManager] from an unsafe.Pointer.
//
// An object you use to create and manage a DNS settings configuration.
func NEDNSSettingsManagerFrom(ptr unsafe.Pointer) NEDNSSettingsManager {
	return NEDNSSettingsManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEDNSSettingsManagerClass) Alloc() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEDNSSettingsManagerClass) New() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSSettingsManager) Init() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSSettingsManager) Autorelease() NEDNSSettingsManager {
	rv := objc.Send[NEDNSSettingsManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSSettingsManager creates a new NEDNSSettingsManager instance.
func NewNEDNSSettingsManager() NEDNSSettingsManager {
	return getNEDNSSettingsManagerClass().New()
}



// Load your DNS settings configuration from the system networking preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/loadFromPreferences(completionHandler:)
func (n_ NEDNSSettingsManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}


// The domain for errors resulting from calls to the DNS settings manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingserrordomain
func (n_ NEDNSSettingsManager) NEDNSSettingsErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEDNSSettingsErrorDomain"))
	return rv
}


// An object that contains the configuration settings for a DNS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/dnssettings
func (n_ NEDNSSettingsManager) DnsSettings() INEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("dnsSettings"))
	return rv
}


// An object that contains the configuration settings for a DNS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/dnssettings
func (n_ NEDNSSettingsManager) SetDnsSettings(value INEDNSSettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSettings:"), value)
}


// A Boolean you use to query the enabled state of the DNS settings configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/isenabled
func (n_ NEDNSSettingsManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean you use to query the enabled state of the DNS settings configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/isenabled
func (n_ NEDNSSettingsManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}


// A string that contains the display name of the DNS settings configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/localizeddescription
func (n_ NEDNSSettingsManager) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string that contains the display name of the DNS settings configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/localizeddescription
func (n_ NEDNSSettingsManager) SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// A list of ordered rules that defines the networks on which the DNS settings will apply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/ondemandrules
func (n_ NEDNSSettingsManager) OnDemandRules() objc.IObject /* cross-framework: NEOnDemandRule */ {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// A list of ordered rules that defines the networks on which the DNS settings will apply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettingsmanager/ondemandrules
func (n_ NEDNSSettingsManager) SetOnDemandRules(value objc.IObject /* cross-framework: NEOnDemandRule */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), value)
}



