// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEDNSSettings */


/* debug [class_header]: Header for NEDNSSettings */
// The class instance for the [NEDNSSettings] class.
var (
	NEDNSSettingsClass     _NEDNSSettingsClass
	NEDNSSettingsClassOnce sync.Once
)

func getNEDNSSettingsClass() _NEDNSSettingsClass {
	NEDNSSettingsClassOnce.Do(func() {
		NEDNSSettingsClass = _NEDNSSettingsClass{objc.GetClass("NEDNSSettings")}
	})
	return NEDNSSettingsClass
}

type _NEDNSSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEDNSSettings */
// An interface definition for the [NEDNSSettings] class.
type INEDNSSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEDNSSettings */
	// properties:
	AllowFailover() bool
	SetAllowFailover(value bool)
	DnsProtocol() NEDNSProtocol
	DomainName() objc.IObject /* cross-framework: NSString */
	SetDomainName(value objc.IObject /* cross-framework: NSString */)
	MatchDomains() []string
	SetMatchDomains(value []string)
	MatchDomainsNoSearch() bool
	SetMatchDomainsNoSearch(value bool)
	SearchDomains() []string
	SetSearchDomains(value []string)
	Servers() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEDNSSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEDNSSettings */
// Alloc allocates a new instance without initialization.
func (nc _NEDNSSettingsClass) Alloc() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEDNSSettingsClass) New() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSSettings) Init() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSSettings) Autorelease() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSSettings creates a new NEDNSSettings instance.
func NewNEDNSSettings() NEDNSSettings {
	return getNEDNSSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEDNSSettings */
// The DNS resolver settings of a network tunnel or a system-wide configuration.


// The DNS resolver settings of a network tunnel or a system-wide configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings
type NEDNSSettings struct {
	objectivec.Object
}

// NEDNSSettingsFrom constructs a [NEDNSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings of a network tunnel or a system-wide configuration.
func NEDNSSettingsFrom(ptr unsafe.Pointer) NEDNSSettings {
	return NEDNSSettings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEDNSSettings */

// Initialize the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/init(servers:)
func NewNEDNSSettingsWithServers(servers []string) NEDNSSettings {
	instance := getNEDNSSettingsClass().Alloc()
	rv := objc.Send[NEDNSSettings](instance.ID, objc.Sel("initWithServers:"), servers)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEDNSSettingsWithServers */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEDNSSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEDNSSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEDNSSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEDNSSettings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/allowFailover
func (n_ NEDNSSettings) AllowFailover() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowFailover"))
	return rv
}/* debug [instance_properties/getter]: allowFailover */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/allowFailover
func (n_ NEDNSSettings) SetAllowFailover(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowFailover:"), value)
}/* debug [instance_properties/setter]: allowFailover */


// The DNS protocol used by the server, such as HTTPS or TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/dnsProtocol
func (n_ NEDNSSettings) DnsProtocol() NEDNSProtocol {
	rv := objc.Send[NEDNSProtocol](n_.ID, objc.Sel("dnsProtocol"))
	return rv
}/* debug [instance_properties/getter]: dnsProtocol */


// The primary domain of the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/domainName
func (n_ NEDNSSettings) DomainName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("domainName"))
	return rv
}/* debug [instance_properties/getter]: domainName */


// The primary domain of the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/domainName
func (n_ NEDNSSettings) SetDomainName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDomainName:"), value)
}/* debug [instance_properties/setter]: domainName */


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/matchDomains
func (n_ NEDNSSettings) MatchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchDomains"))
	return rv
}/* debug [instance_properties/getter]: matchDomains */


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/matchDomains
func (n_ NEDNSSettings) SetMatchDomains(value []string) {
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
}/* debug [instance_properties/setter]: matchDomains */


// A Boolean that specifies if the domains in the list should not be appended to the resolver’s list of search domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/matchDomainsNoSearch
func (n_ NEDNSSettings) MatchDomainsNoSearch() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("matchDomainsNoSearch"))
	return rv
}/* debug [instance_properties/getter]: matchDomainsNoSearch */


// A Boolean that specifies if the domains in the list should not be appended to the resolver’s list of search domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/matchDomainsNoSearch
func (n_ NEDNSSettings) SetMatchDomainsNoSearch(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomainsNoSearch:"), value)
}/* debug [instance_properties/setter]: matchDomainsNoSearch */


// A list of domain strings used to fully qualify single-label host names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/searchDomains
func (n_ NEDNSSettings) SearchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("searchDomains"))
	return rv
}/* debug [instance_properties/getter]: searchDomains */


// A list of domain strings used to fully qualify single-label host names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/searchDomains
func (n_ NEDNSSettings) SetSearchDomains(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchDomains:"), nsArray)
}/* debug [instance_properties/setter]: searchDomains */


// The DNS server IP addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/servers
func (n_ NEDNSSettings) Servers() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("servers"))
	return rv
}/* debug [instance_properties/getter]: servers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEDNSSettings */


