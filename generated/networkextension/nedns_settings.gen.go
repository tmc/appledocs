// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AllowFailover() bool
	SetAllowFailover(value bool)
	DnsProtocol() unsafe.Pointer
	SetDnsProtocol(value unsafe.Pointer)
	DomainName() string
	SetDomainName(value string)
	MatchDomains() string
	SetMatchDomains(value string)
	MatchDomainsNoSearch() bool
	SetMatchDomainsNoSearch(value bool)
	SearchDomains() string
	SetSearchDomains(value string)
	Servers() string
	SetServers(value string)
}

// The DNS resolver settings of a network tunnel or a system-wide configuration.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/allowfailover
func (n_ NEDNSSettings) AllowFailover() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowFailover"))
	return rv
}


// SetAllowFailover sets the value of the allowFailover property.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/allowfailover
func (n_ NEDNSSettings) SetAllowFailover(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowFailover:"), value)
}

// The DNS protocol used by the server, such as HTTPS or TLS.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/dnsprotocol
func (n_ NEDNSSettings) DnsProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dnsProtocol"))
	return rv
}


// SetDnsProtocol sets the value of the dnsProtocol property.
// The DNS protocol used by the server, such as HTTPS or TLS.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/dnsprotocol
func (n_ NEDNSSettings) SetDnsProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsProtocol:"), value)
}

// The primary domain of the tunnel.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/domainname
func (n_ NEDNSSettings) DomainName() string {
	rv := objc.Send[string](n_.ID, objc.Sel("domainName"))
	return rv
}


// SetDomainName sets the value of the domainName property.
// The primary domain of the tunnel.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/domainname
func (n_ NEDNSSettings) SetDomainName(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDomainName:"), objc.String(value))
}

// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSSettings) MatchDomains() string {
	rv := objc.Send[string](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// SetMatchDomains sets the value of the matchDomains property.
// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSSettings) SetMatchDomains(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), objc.String(value))
}

// A Boolean that specifies if the domains in the
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomainsnosearch
func (n_ NEDNSSettings) MatchDomainsNoSearch() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("matchDomainsNoSearch"))
	return rv
}


// SetMatchDomainsNoSearch sets the value of the matchDomainsNoSearch property.
// A Boolean that specifies if the domains in the

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomainsnosearch
func (n_ NEDNSSettings) SetMatchDomainsNoSearch(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomainsNoSearch:"), value)
}

// A list of domain strings used to fully qualify single-label host names.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/searchdomains
func (n_ NEDNSSettings) SearchDomains() string {
	rv := objc.Send[string](n_.ID, objc.Sel("searchDomains"))
	return rv
}


// SetSearchDomains sets the value of the searchDomains property.
// A list of domain strings used to fully qualify single-label host names.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/searchdomains
func (n_ NEDNSSettings) SetSearchDomains(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSearchDomains:"), objc.String(value))
}

// The DNS server IP addresses.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/servers
func (n_ NEDNSSettings) Servers() string {
	rv := objc.Send[string](n_.ID, objc.Sel("servers"))
	return rv
}


// SetServers sets the value of the servers property.
// The DNS server IP addresses.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/servers
func (n_ NEDNSSettings) SetServers(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServers:"), objc.String(value))
}



