// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AppRules() INEAppRule
	SetAppRules(value INEAppRule)
	AssociatedDomains() objc.IObject /* cross-framework: NSString */
	SetAssociatedDomains(value objc.IObject /* cross-framework: NSString */)
	CalendarDomains() objc.IObject /* cross-framework: NSString */
	SetCalendarDomains(value objc.IObject /* cross-framework: NSString */)
	ContactsDomains() objc.IObject /* cross-framework: NSString */
	SetContactsDomains(value objc.IObject /* cross-framework: NSString */)
	ExcludedDomains() objc.IObject /* cross-framework: NSString */
	SetExcludedDomains(value objc.IObject /* cross-framework: NSString */)
	MailDomains() objc.IObject /* cross-framework: NSString */
	SetMailDomains(value objc.IObject /* cross-framework: NSString */)
	RoutingMethod() unsafe.Pointer
	SetRoutingMethod(value unsafe.Pointer)
	SafariDomains() objc.IObject /* cross-framework: NSString */
	SetSafariDomains(value objc.IObject /* cross-framework: NSString */)
	Connection() INEVPNConnection
	SetConnection(value INEVPNConnection)
	OnDemandRules() objc.IObject /* cross-framework: NEOnDemandRule */
	SetOnDemandRules(value objc.IObject /* cross-framework: NEOnDemandRule */)
	ProtocolConfiguration() INEVPNProtocol
	SetProtocolConfiguration(value INEVPNProtocol)
	// methods:
}

// An object to create and manage the tunnel provider’s VPN configuration.
//
// Like its superclass , you use the class to configure and control VPN connections. The difference is that is used to to configure and control VPN connections that use a custom VPN protocol. The client side of the custom protocol implementation is implemented as a Packet Tunnel Provider extension. The Packet Tunnel Provider extension’s containing app uses to create and manage VPN configurations that use the custom protocol, and to control the VPN connections specified by the configurations. The class inherits most of its functionality from the class. The key differences to be aware of when using are: The property can only be set to instances of the class The read-only property is set to an instance of the class.


// An object to create and manage the tunnel provider’s VPN configuration.
//
// [Full Topic]
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



// The rules for specific apps in a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/apprules
func (n_ NETunnelProviderManager) AppRules() INEAppRule {
	rv := objc.Send[NEAppRule](n_.ID, objc.Sel("appRules"))
	return rv
}


// The rules for specific apps in a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/apprules
func (n_ NETunnelProviderManager) SetAppRules(value INEAppRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAppRules:"), value)
}


// The domains that the system routes network traffic through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/associateddomains
func (n_ NETunnelProviderManager) AssociatedDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("associatedDomains"))
	return rv
}


// The domains that the system routes network traffic through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/associateddomains
func (n_ NETunnelProviderManager) SetAssociatedDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAssociatedDomains:"), value)
}


// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/calendardomains
func (n_ NETunnelProviderManager) CalendarDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("calendarDomains"))
	return rv
}


// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/calendardomains
func (n_ NETunnelProviderManager) SetCalendarDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCalendarDomains:"), value)
}


// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/contactsdomains
func (n_ NETunnelProviderManager) ContactsDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("contactsDomains"))
	return rv
}


// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/contactsdomains
func (n_ NETunnelProviderManager) SetContactsDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setContactsDomains:"), value)
}


// The domains that the system excludes from a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/excludeddomains
func (n_ NETunnelProviderManager) ExcludedDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("excludedDomains"))
	return rv
}


// The domains that the system excludes from a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/excludeddomains
func (n_ NETunnelProviderManager) SetExcludedDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedDomains:"), value)
}


// The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/maildomains
func (n_ NETunnelProviderManager) MailDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("mailDomains"))
	return rv
}


// The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/maildomains
func (n_ NETunnelProviderManager) SetMailDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMailDomains:"), value)
}


// The method that the system uses to route network traffic to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/routingmethod
func (n_ NETunnelProviderManager) RoutingMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("routingMethod"))
	return rv
}


// The method that the system uses to route network traffic to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/routingmethod
func (n_ NETunnelProviderManager) SetRoutingMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoutingMethod:"), value)
}


// The website domains that the system routes connections from the Safari app through a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/safaridomains
func (n_ NETunnelProviderManager) SafariDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("safariDomains"))
	return rv
}


// The website domains that the system routes connections from the Safari app through a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidermanager/safaridomains
func (n_ NETunnelProviderManager) SetSafariDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSafariDomains:"), value)
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/connection
func (n_ NETunnelProviderManager) Connection() INEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("connection"))
	return rv
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/connection
func (n_ NETunnelProviderManager) SetConnection(value INEVPNConnection) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnection:"), value)
}


// An ordered list of Connect On Demand rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/ondemandrules
func (n_ NETunnelProviderManager) OnDemandRules() objc.IObject /* cross-framework: NEOnDemandRule */ {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}


// An ordered list of Connect On Demand rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/ondemandrules
func (n_ NETunnelProviderManager) SetOnDemandRules(value objc.IObject /* cross-framework: NEOnDemandRule */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), value)
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/protocolconfiguration
func (n_ NETunnelProviderManager) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/protocolconfiguration
func (n_ NETunnelProviderManager) SetProtocolConfiguration(value INEVPNProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocolConfiguration:"), value)
}



