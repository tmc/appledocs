// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETunnelProviderManager] class.
var (
	NETunnelProviderManagerClass     _NETunnelProviderManagerClass
	NETunnelProviderManagerClassOnce sync.Once
)

func getNETunnelProviderManagerClass() _NETunnelProviderManagerClass {
	NETunnelProviderManagerClassOnce.Do(func() {
		NETunnelProviderManagerClass = _NETunnelProviderManagerClass{objc.GetClass("NETunnelProviderManager")}
	})
	return NETunnelProviderManagerClass
}

type _NETunnelProviderManagerClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelProviderManager] class.
type INETunnelProviderManager interface {
	INEVPNManager
	CopyAppRules() []NEAppRule
	AppRules() []NEAppRule
	SetAppRules(value []NEAppRule)
	AssociatedDomains() []string
	SetAssociatedDomains(value []string)
	CalendarDomains() []string
	SetCalendarDomains(value []string)
	ContactsDomains() []string
	SetContactsDomains(value []string)
	ExcludedDomains() []string
	SetExcludedDomains(value []string)
	MailDomains() []string
	SetMailDomains(value []string)
	RoutingMethod() NETunnelProviderRoutingMethod
	SafariDomains() []string
	SetSafariDomains(value []string)
	Connection() NEVPNConnection
	SetConnection(value INEVPNConnection)
	OnDemandRules() NEOnDemandRule
	SetOnDemandRules(value INEOnDemandRule)
	ProtocolConfiguration() NEVPNProtocol
	SetProtocolConfiguration(value INEVPNProtocol)
}

// An object to create and manage the tunnel provider’s VPN configuration.
//
// Like its superclass , you use the class to configure and control VPN connections. The difference is that is used to to configure and control VPN connections that use a custom VPN protocol. The client side of the custom protocol implementation is implemented as a Packet Tunnel Provider extension. The Packet Tunnel Provider extension’s containing app uses to create and manage VPN configurations that use the custom protocol, and to control the VPN connections specified by the configurations. The class inherits most of its functionality from the class. The key differences to be aware of when using are: The property can only be set to instances of the class The read-only property is set to an instance of the class.
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


// Returns a tunnel provider manager for managing a per-app VPN configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/forPerAppVPN()
func (nc _NETunnelProviderManagerClass) ForPerAppVPN() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("forPerAppVPN"))
	return rv
}

// Read all of the VPN configurations created by the calling app that have previously been saved to the Network Extension preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/loadAllFromPreferences(completionHandler:)
func (nc _NETunnelProviderManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}

// Returns a copy of the app rules currently set in the configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/copyAppRules()
func (n_ NETunnelProviderManager) CopyAppRules() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("copyAppRules"))
	return rv
}

// The rules for specific apps in a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/appRules
func (n_ NETunnelProviderManager) AppRules() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("appRules"))
	return rv
}


// SetAppRules sets the value of the appRules property.
// The rules for specific apps in a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/appRules
func (n_ NETunnelProviderManager) SetAppRules(value []NEAppRule) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setAppRules:"), nsArray)
}

// The domains that the system routes network traffic through for a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/associatedDomains
func (n_ NETunnelProviderManager) AssociatedDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("associatedDomains"))
	return rv
}


// SetAssociatedDomains sets the value of the associatedDomains property.
// The domains that the system routes network traffic through for a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/associatedDomains
func (n_ NETunnelProviderManager) SetAssociatedDomains(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setAssociatedDomains:"), nsArray)
}

// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/calendarDomains
func (n_ NETunnelProviderManager) CalendarDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("calendarDomains"))
	return rv
}


// SetCalendarDomains sets the value of the calendarDomains property.
// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/calendarDomains
func (n_ NETunnelProviderManager) SetCalendarDomains(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setCalendarDomains:"), nsArray)
}

// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/contactsDomains
func (n_ NETunnelProviderManager) ContactsDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("contactsDomains"))
	return rv
}


// SetContactsDomains sets the value of the contactsDomains property.
// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/contactsDomains
func (n_ NETunnelProviderManager) SetContactsDomains(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setContactsDomains:"), nsArray)
}

// The domains that the system excludes from a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/excludedDomains
func (n_ NETunnelProviderManager) ExcludedDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("excludedDomains"))
	return rv
}


// SetExcludedDomains sets the value of the excludedDomains property.
// The domains that the system excludes from a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/excludedDomains
func (n_ NETunnelProviderManager) SetExcludedDomains(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedDomains:"), nsArray)
}

// The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/mailDomains
func (n_ NETunnelProviderManager) MailDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("mailDomains"))
	return rv
}


// SetMailDomains sets the value of the mailDomains property.
// The mail servers that the system routes connections from the Mail app through for a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/mailDomains
func (n_ NETunnelProviderManager) SetMailDomains(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setMailDomains:"), nsArray)
}

// The method that the system uses to route network traffic to the tunnel.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/routingMethod
func (n_ NETunnelProviderManager) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](n_.ID, objc.Sel("routingMethod"))
	return rv
}

// The website domains that the system routes connections from the Safari app through a per-app VPN.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/safariDomains
func (n_ NETunnelProviderManager) SafariDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("safariDomains"))
	return rv
}


// SetSafariDomains sets the value of the safariDomains property.
// The website domains that the system routes connections from the Safari app through a per-app VPN.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/safariDomains
func (n_ NETunnelProviderManager) SetSafariDomains(value []string) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setSafariDomains:"), nsArray)
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/connection
func (n_ NETunnelProviderManager) Connection() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("connection"))
	return rv
}


// SetConnection sets the value of the connection property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/connection
func (n_ NETunnelProviderManager) SetConnection(value INEVPNConnection) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnection:"), value)
}

// An ordered list of Connect On Demand rules.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/ondemandrules
func (n_ NETunnelProviderManager) OnDemandRules() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// SetOnDemandRules sets the value of the onDemandRules property.
// An ordered list of Connect On Demand rules.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/ondemandrules
func (n_ NETunnelProviderManager) SetOnDemandRules(value INEOnDemandRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), value)
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/protocolconfiguration
func (n_ NETunnelProviderManager) ProtocolConfiguration() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}


// SetProtocolConfiguration sets the value of the protocolConfiguration property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/protocolconfiguration
func (n_ NETunnelProviderManager) SetProtocolConfiguration(value INEVPNProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocolConfiguration:"), value)
}



