// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NETunnelProviderManager */


/* debug [class_header]: Header for NETunnelProviderManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NETunnelProviderManager */
// An interface definition for the [NETunnelProviderManager] class.
type INETunnelProviderManager interface {
	INEVPNManager
	
/* debug [class_interface_properties]: Properties for NETunnelProviderManager */
	// properties:
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
	Connection() INEVPNConnection
	SetConnection(value INEVPNConnection)
	OnDemandRules() INEOnDemandRule
	SetOnDemandRules(value INEOnDemandRule)
	ProtocolConfiguration() INEVPNProtocol
	SetProtocolConfiguration(value INEVPNProtocol)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NETunnelProviderManager */
	// methods:
	CopyAppRules() []NEAppRule
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NETunnelProviderManager */
// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderManagerClass) Alloc() NETunnelProviderManager {
	rv := objc.Send[NETunnelProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NETunnelProviderManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NETunnelProviderManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NETunnelProviderManager */

// Returns a tunnel provider manager for managing a per-app VPN configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/forPerAppVPN()
func (nc _NETunnelProviderManagerClass) ForPerAppVPN() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("forPerAppVPN"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ForPerAppVPN) */


// Read all of the VPN configurations created by the calling app that have previously been saved to the Network Extension preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/loadAllFromPreferences(completionHandler:)
func (nc _NETunnelProviderManagerClass) LoadAllFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("loadAllFromPreferencesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadAllFromPreferencesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NETunnelProviderManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NETunnelProviderManager */

// Returns a copy of the app rules currently set in the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/copyAppRules()
func (n_ NETunnelProviderManager) CopyAppRules() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("copyAppRules"))
	return rv
}/* debug [instance_methods/method]: CopyAppRules */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NETunnelProviderManager */

// The rules for specific apps in a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/appRules
func (n_ NETunnelProviderManager) AppRules() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("appRules"))
	return rv
}/* debug [instance_properties/getter]: appRules */


// The rules for specific apps in a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/appRules
func (n_ NETunnelProviderManager) SetAppRules(value []NEAppRule) {
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
}/* debug [instance_properties/setter]: appRules */


// The domains that the system routes network traffic through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/associatedDomains
func (n_ NETunnelProviderManager) AssociatedDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("associatedDomains"))
	return rv
}/* debug [instance_properties/getter]: associatedDomains */


// The domains that the system routes network traffic through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/associatedDomains
func (n_ NETunnelProviderManager) SetAssociatedDomains(value []string) {
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
}/* debug [instance_properties/setter]: associatedDomains */


// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/calendarDomains
func (n_ NETunnelProviderManager) CalendarDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("calendarDomains"))
	return rv
}/* debug [instance_properties/getter]: calendarDomains */


// The calendar servers that the system routes connections from the Calendar app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/calendarDomains
func (n_ NETunnelProviderManager) SetCalendarDomains(value []string) {
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
}/* debug [instance_properties/setter]: calendarDomains */


// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/contactsDomains
func (n_ NETunnelProviderManager) ContactsDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("contactsDomains"))
	return rv
}/* debug [instance_properties/getter]: contactsDomains */


// The contacts servers that the system routes connections from the Contacts app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/contactsDomains
func (n_ NETunnelProviderManager) SetContactsDomains(value []string) {
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
}/* debug [instance_properties/setter]: contactsDomains */


// The domains that the system excludes from a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/excludedDomains
func (n_ NETunnelProviderManager) ExcludedDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("excludedDomains"))
	return rv
}/* debug [instance_properties/getter]: excludedDomains */


// The domains that the system excludes from a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/excludedDomains
func (n_ NETunnelProviderManager) SetExcludedDomains(value []string) {
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
}/* debug [instance_properties/setter]: excludedDomains */


// The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/mailDomains
func (n_ NETunnelProviderManager) MailDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("mailDomains"))
	return rv
}/* debug [instance_properties/getter]: mailDomains */


// The mail servers that the system routes connections from the Mail app through for a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/mailDomains
func (n_ NETunnelProviderManager) SetMailDomains(value []string) {
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
}/* debug [instance_properties/setter]: mailDomains */


// The method that the system uses to route network traffic to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/routingMethod
func (n_ NETunnelProviderManager) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](n_.ID, objc.Sel("routingMethod"))
	return rv
}/* debug [instance_properties/getter]: routingMethod */


// The website domains that the system routes connections from the Safari app through a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/safariDomains
func (n_ NETunnelProviderManager) SafariDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("safariDomains"))
	return rv
}/* debug [instance_properties/getter]: safariDomains */


// The website domains that the system routes connections from the Safari app through a per-app VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderManager/safariDomains
func (n_ NETunnelProviderManager) SetSafariDomains(value []string) {
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
}/* debug [instance_properties/setter]: safariDomains */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/connection
func (n_ NETunnelProviderManager) Connection() INEVPNConnection {
	rv := objc.Send[NEVPNConnection](n_.ID, objc.Sel("connection"))
	return rv
}/* debug [instance_properties/getter]: connection */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/connection
func (n_ NETunnelProviderManager) SetConnection(value INEVPNConnection) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnection:"), value)
}/* debug [instance_properties/setter]: connection */


// An ordered list of Connect On Demand rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/ondemandrules
func (n_ NETunnelProviderManager) OnDemandRules() INEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("onDemandRules"))
	return rv
}/* debug [instance_properties/getter]: onDemandRules */


// An ordered list of Connect On Demand rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/ondemandrules
func (n_ NETunnelProviderManager) SetOnDemandRules(value INEOnDemandRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOnDemandRules:"), value)
}/* debug [instance_properties/setter]: onDemandRules */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/protocolconfiguration
func (n_ NETunnelProviderManager) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}/* debug [instance_properties/getter]: protocolConfiguration */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnmanager/protocolconfiguration
func (n_ NETunnelProviderManager) SetProtocolConfiguration(value INEVPNProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocolConfiguration:"), value)
}/* debug [instance_properties/setter]: protocolConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NETunnelProviderManager */



