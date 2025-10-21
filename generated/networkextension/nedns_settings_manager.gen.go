// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object you use to create and manage a DNS settings configuration.
//
// When your app starts up, access the shared instance of the DNS settings manager, and load existing settings from the preferences using . You can define your DNS server configuration, and persist it by calling . In order to use your DNS settings, the user needs to enable it in the Settings app on iOS or in System Preferences on macOS.
//
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


// Access the single instance of a DNS settings manager.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/shared()
func (nc _NEDNSSettingsManagerClass) SharedManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}

// Load your DNS settings configuration from the system networking preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/loadFromPreferences(completionHandler:)
func (n_ NEDNSSettingsManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}

// Remove your DNS settings configuration from the system networking preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/removeFromPreferences(completionHandler:)
func (n_ NEDNSSettingsManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}

// Save your DNS settings configuration to the system networking preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/saveToPreferences(completionHandler:)
func (n_ NEDNSSettingsManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}

// An object that contains the configuration settings for a DNS server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/dnsSettings
func (n_ NEDNSSettingsManager) DnsSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dnsSettings"))
	return rv
}


// SetDnsSettings sets the value of the dnsSettings property.
// An object that contains the configuration settings for a DNS server.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/dnsSettings
func (n_ NEDNSSettingsManager) SetDnsSettings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSettings:"), value)
}
// A Boolean you use to query the enabled state of the DNS settings configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/isEnabled
func (n_ NEDNSSettingsManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}

// A string that contains the display name of the DNS settings configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/localizedDescription
func (n_ NEDNSSettingsManager) LocalizedDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// SetLocalizedDescription sets the value of the localizedDescription property.
// A string that contains the display name of the DNS settings configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/localizedDescription
func (n_ NEDNSSettingsManager) SetLocalizedDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}
// A list of ordered rules that defines the networks on which the DNS settings will apply.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/onDemandRules
func (n_ NEDNSSettingsManager) OnDemandRules() []NEOnDemandRule {
	rv := objc.Send[[]NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// SetOnDemandRules sets the value of the onDemandRules property.
// A list of ordered rules that defines the networks on which the DNS settings will apply.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsManager/onDemandRules
func (n_ NEDNSSettingsManager) SetOnDemandRules(value []NEOnDemandRule) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), nsArray)
}


