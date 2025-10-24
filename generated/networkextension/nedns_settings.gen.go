// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEDNSSettings] class.
type INEDNSSettings interface {
	objectivec.IObject
	// properties:
	DomainName() objc.IObject /* cross-framework: NSString */
	SetDomainName(value objc.IObject /* cross-framework: NSString */)
	AllowFailover() bool
	SetAllowFailover(value bool)
	DnsProtocol() unsafe.Pointer
	SetDnsProtocol(value unsafe.Pointer)
	MatchDomains() objc.IObject /* cross-framework: NSString */
	SetMatchDomains(value objc.IObject /* cross-framework: NSString */)
	MatchDomainsNoSearch() bool
	SetMatchDomainsNoSearch(value bool)
	SearchDomains() objc.IObject /* cross-framework: NSString */
	SetSearchDomains(value objc.IObject /* cross-framework: NSString */)
	Servers() objc.IObject /* cross-framework: NSString */
	SetServers(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (nc _NEDNSSettingsClass) Alloc() NEDNSSettings {
	rv := objc.Send[NEDNSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The primary domain of the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/domainName
func (n_ NEDNSSettings) DomainName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("domainName"))
	return rv
}


// The primary domain of the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettings/domainName
func (n_ NEDNSSettings) SetDomainName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDomainName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/allowfailover
func (n_ NEDNSSettings) AllowFailover() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowFailover"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/allowfailover
func (n_ NEDNSSettings) SetAllowFailover(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowFailover:"), value)
}


// The DNS protocol used by the server, such as HTTPS or TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/dnsprotocol
func (n_ NEDNSSettings) DnsProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dnsProtocol"))
	return rv
}


// The DNS protocol used by the server, such as HTTPS or TLS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/dnsprotocol
func (n_ NEDNSSettings) SetDnsProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsProtocol:"), value)
}


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSSettings) MatchDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSSettings) SetMatchDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}


// A Boolean that specifies if the domains in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomainsnosearch
func (n_ NEDNSSettings) MatchDomainsNoSearch() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("matchDomainsNoSearch"))
	return rv
}


// A Boolean that specifies if the domains in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomainsnosearch
func (n_ NEDNSSettings) SetMatchDomainsNoSearch(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomainsNoSearch:"), value)
}


// A list of domain strings used to fully qualify single-label host names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/searchdomains
func (n_ NEDNSSettings) SearchDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("searchDomains"))
	return rv
}


// A list of domain strings used to fully qualify single-label host names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/searchdomains
func (n_ NEDNSSettings) SetSearchDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchDomains:"), value)
}


// The DNS server IP addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/servers
func (n_ NEDNSSettings) Servers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("servers"))
	return rv
}


// The DNS server IP addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/servers
func (n_ NEDNSSettings) SetServers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServers:"), value)
}



