// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NERelayManager] class.
var (
	NERelayManagerClass     _NERelayManagerClass
	NERelayManagerClassOnce sync.Once
)

func getNERelayManagerClass() _NERelayManagerClass {
	NERelayManagerClassOnce.Do(func() {
		NERelayManagerClass = _NERelayManagerClass{objc.GetClass("NERelayManager")}
	})
	return NERelayManagerClass
}

type _NERelayManagerClass struct {
	class objc.Class
}





// An interface definition for the [NERelayManager] class.
type INERelayManager interface {
	objectivec.IObject
	

	// properties:
	ExcludedDomains() []string
	SetExcludedDomains(value []string)
	ExcludedFQDNs() []string
	SetExcludedFQDNs(value []string)
	AllowDNSFailover() bool
	SetAllowDNSFailover(value bool)
	Enabled() bool
	SetEnabled(value bool)
	UIToggleEnabled() bool
	SetUIToggleEnabled(value bool)
	LocalizedDescription() foundation.foundation.INSString
	SetLocalizedDescription(value foundation.foundation.INSString)
	MatchDomains() []string
	SetMatchDomains(value []string)
	MatchFQDNs() []string
	SetMatchFQDNs(value []string)
	OnDemandRules() []NEOnDemandRule
	SetOnDemandRules(value []NEOnDemandRule)
	Relays() []NERelay
	SetRelays(value []NERelay)
	NERelayErrorDomain() foundation.foundation.INSString
	IsDNSFailoverAllowed() bool
	SetIsDNSFailoverAllowed(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsUIToggleEnabled() bool
	SetIsUIToggleEnabled(value bool)


	

	// methods:
	GetLastClientErrorsCompletionHandler(seconds float64, completionHandler unsafe.Pointer)
	LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (nc _NERelayManagerClass) Alloc() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NERelayManagerClass) New() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NERelayManager) Init() NERelayManager {
	rv := objc.Send[NERelayManager](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NERelayManager) Autorelease() NERelayManager {
	rv := objc.Send[NERelayManager](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNERelayManager creates a new NERelayManager instance.
func NewNERelayManager() NERelayManager {
	return getNERelayManagerClass().New()
}





// An object you use to create and manage a network relay configuration.
//
// When your app starts up, access the shared instance of the relay manager, and load existing settings from the preferences using . You can define your relay server configuration, and persist it by calling .


// An object you use to create and manage a network relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager
type NERelayManager struct {
	objectivec.Object
}

// NERelayManagerFrom constructs a [NERelayManager] from an unsafe.Pointer.
//
// An object you use to create and manage a network relay configuration.
func NERelayManagerFrom(ptr unsafe.Pointer) NERelayManager {
	return NERelayManager{objectivec.Object{objc.ID(ptr)}}
}










// Asynchronously reads all the relay configurations previously created and saved by the calling app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/loadAllManagersFromPreferences(completionHandler:)
func (nc _NERelayManagerClass) LoadAllManagersFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllManagersFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Access the single instance of a network relay manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/shared()
func (nc _NERelayManagerClass) SharedManager() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("sharedManager"))
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/getLastClientErrors(_:completionHandler:)
func (n_ NERelayManager) GetLastClientErrorsCompletionHandler(seconds float64, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("getLastClientErrors:completionHandler:"), seconds, completionHandler)
}


// Load your relay configuration from the system networking preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/loadFromPreferences(completionHandler:)
func (n_ NERelayManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Remove your relay configuration from the system networking preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/removeFromPreferences(completionHandler:)
func (n_ NERelayManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}


// Save your relay configuration to the system networking preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/saveToPreferences(completionHandler:)
func (n_ NERelayManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}







// A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/excludedDomains
func (n_ NERelayManager) ExcludedDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("excludedDomains"))
	return rv
}


// A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/excludedDomains
func (n_ NERelayManager) SetExcludedDomains(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedDomains:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/excludedFQDNs
func (n_ NERelayManager) ExcludedFQDNs() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("excludedFQDNs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/excludedFQDNs
func (n_ NERelayManager) SetExcludedFQDNs(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedFQDNs:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isDNSFailoverAllowed
func (n_ NERelayManager) AllowDNSFailover() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowDNSFailover"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isDNSFailoverAllowed
func (n_ NERelayManager) SetAllowDNSFailover(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowDNSFailover:"), value)
}


// A Boolean used to toggle the enabled state of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isEnabled
func (n_ NERelayManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean used to toggle the enabled state of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isEnabled
func (n_ NERelayManager) SetEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isUIToggleEnabled
func (n_ NERelayManager) UIToggleEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("UIToggleEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/isUIToggleEnabled
func (n_ NERelayManager) SetUIToggleEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUIToggleEnabled:"), value)
}


// A string that contains the display name of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/localizedDescription
func (n_ NERelayManager) LocalizedDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string that contains the display name of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/localizedDescription
func (n_ NERelayManager) SetLocalizedDescription(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// A list of domain strings used to determine which connections will use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/matchDomains
func (n_ NERelayManager) MatchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// A list of domain strings used to determine which connections will use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/matchDomains
func (n_ NERelayManager) SetMatchDomains(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/matchFQDNs
func (n_ NERelayManager) MatchFQDNs() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchFQDNs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/matchFQDNs
func (n_ NERelayManager) SetMatchFQDNs(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchFQDNs:"), nsArray)
}


// An array of rules you use to determine which networks the relay uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/onDemandRules
func (n_ NERelayManager) OnDemandRules() []NEOnDemandRule {
	rv := objc.Send[[]NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// An array of rules you use to determine which networks the relay uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/onDemandRules
func (n_ NERelayManager) SetOnDemandRules(value []NEOnDemandRule) {
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


// An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/relays
func (n_ NERelayManager) Relays() []NERelay {
	rv := objc.Send[[]NERelay](n_.ID, objc.Sel("relays"))
	return rv
}


// An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelayManager/relays
func (n_ NERelayManager) SetRelays(value []NERelay) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setRelays:"), nsArray)
}


// The domain for errors resulting from calls to the relay manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelayerrordomain
func (n_ NERelayManager) NERelayErrorDomain() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NERelayErrorDomain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/isdnsfailoverallowed
func (n_ NERelayManager) IsDNSFailoverAllowed() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isDNSFailoverAllowed"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/isdnsfailoverallowed
func (n_ NERelayManager) SetIsDNSFailoverAllowed(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsDNSFailoverAllowed:"), value)
}


// A Boolean used to toggle the enabled state of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/isenabled
func (n_ NERelayManager) IsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean used to toggle the enabled state of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/isenabled
func (n_ NERelayManager) SetIsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/isuitoggleenabled
func (n_ NERelayManager) IsUIToggleEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isUIToggleEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/isuitoggleenabled
func (n_ NERelayManager) SetIsUIToggleEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsUIToggleEnabled:"), value)
}








