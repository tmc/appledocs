// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	NERelayErrorDomain() objc.IObject /* cross-framework: NSString */
	ExcludedDomains() objc.IObject /* cross-framework: NSString */
	SetExcludedDomains(value objc.IObject /* cross-framework: NSString */)
	ExcludedFQDNs() objc.IObject /* cross-framework: NSString */
	SetExcludedFQDNs(value objc.IObject /* cross-framework: NSString */)
	IsDNSFailoverAllowed() bool
	SetIsDNSFailoverAllowed(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsUIToggleEnabled() bool
	SetIsUIToggleEnabled(value bool)
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */)
	MatchDomains() objc.IObject /* cross-framework: NSString */
	SetMatchDomains(value objc.IObject /* cross-framework: NSString */)
	MatchFQDNs() objc.IObject /* cross-framework: NSString */
	SetMatchFQDNs(value objc.IObject /* cross-framework: NSString */)
	OnDemandRules() objc.IObject /* cross-framework: NEOnDemandRule */
	SetOnDemandRules(value objc.IObject /* cross-framework: NEOnDemandRule */)
	Relays() INERelay
	SetRelays(value INERelay)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (nc _NERelayManagerClass) Alloc() NERelayManager {
	rv := objc.Send[NERelayManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The domain for errors resulting from calls to the relay manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelayerrordomain
func (n_ NERelayManager) NERelayErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NERelayErrorDomain"))
	return rv
}


// A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/excludeddomains
func (n_ NERelayManager) ExcludedDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("excludedDomains"))
	return rv
}


// A list of domain strings used to determine which connections won’t use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/excludeddomains
func (n_ NERelayManager) SetExcludedDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedDomains:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/excludedfqdns
func (n_ NERelayManager) ExcludedFQDNs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("excludedFQDNs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/excludedfqdns
func (n_ NERelayManager) SetExcludedFQDNs(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedFQDNs:"), value)
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


// A string that contains the display name of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/localizeddescription
func (n_ NERelayManager) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}


// A string that contains the display name of the relay configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/localizeddescription
func (n_ NERelayManager) SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizedDescription:"), value)
}


// A list of domain strings used to determine which connections will use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/matchdomains
func (n_ NERelayManager) MatchDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// A list of domain strings used to determine which connections will use the relay configuration contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/matchdomains
func (n_ NERelayManager) SetMatchDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/matchfqdns
func (n_ NERelayManager) MatchFQDNs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchFQDNs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/matchfqdns
func (n_ NERelayManager) SetMatchFQDNs(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchFQDNs:"), value)
}


// An array of rules you use to determine which networks the relay uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/ondemandrules
func (n_ NERelayManager) OnDemandRules() objc.IObject /* cross-framework: NEOnDemandRule */ {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// An array of rules you use to determine which networks the relay uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/ondemandrules
func (n_ NERelayManager) SetOnDemandRules(value objc.IObject /* cross-framework: NEOnDemandRule */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), value)
}


// An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/relays
func (n_ NERelayManager) Relays() INERelay {
	rv := objc.Send[NERelay](n_.ID, objc.Sel("relays"))
	return rv
}


// An array of one or two relay server configurations. If multiple relays are configured, application traffic routes through both of them in the order they appear in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelaymanager/relays
func (n_ NERelayManager) SetRelays(value INERelay) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRelays:"), value)
}



