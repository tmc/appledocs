// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEOnDemandRule] class.
var (
	NEOnDemandRuleClass     _NEOnDemandRuleClass
	NEOnDemandRuleClassOnce sync.Once
)

func getNEOnDemandRuleClass() _NEOnDemandRuleClass {
	NEOnDemandRuleClassOnce.Do(func() {
		NEOnDemandRuleClass = _NEOnDemandRuleClass{objc.GetClass("NEOnDemandRule")}
	})
	return NEOnDemandRuleClass
}

type _NEOnDemandRuleClass struct {
	class objc.Class
}

// An interface definition for the [NEOnDemandRule] class.
type INEOnDemandRule interface {
	objectivec.IObject
}

// A base class shared by all VPN On Demand rules.
//
// Each rule is defined by a single action and a set of optional matching conditions. The action defines how the system should trigger the VPN when the conditions are met, such as connecting automatically for all connections, connecting conditionally, or disconnecting. The optional conditions describe parameters of a network. Some common rules include disconnecting the VPN on a trusted, internal network, and triggering on all other networks. When rules are defined in an array, they are evaluated in order and the action of the first rule to match all conditions is chosen. Instances of the class should be created through one of its subclasses: , , , or .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule
type NEOnDemandRule struct {
	objectivec.Object
}

// NEOnDemandRuleFrom constructs a [NEOnDemandRule] from an unsafe.Pointer.
//
// A base class shared by all VPN On Demand rules.
func NEOnDemandRuleFrom(ptr unsafe.Pointer) NEOnDemandRule {
	return NEOnDemandRule{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleClass) Alloc() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEOnDemandRuleClass) New() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEOnDemandRule) Init() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEOnDemandRule) Autorelease() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRule creates a new NEOnDemandRule instance.
func NewNEOnDemandRule() NEOnDemandRule {
	return getNEOnDemandRuleClass().New()
}


// The action of the On Demand Rule.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/action
func (n_ NEOnDemandRule) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The action of the On Demand Rule.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/action
func (n_ NEOnDemandRule) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAction:"), value)
}

// DNS search domains that identify a network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/dnssearchdomainmatch
func (n_ NEOnDemandRule) DnsSearchDomainMatch() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("dnsSearchDomainMatch"))
	return rv
}


// SetDnsSearchDomainMatch sets the value of the dnsSearchDomainMatch property.
// DNS search domains that identify a network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/dnssearchdomainmatch
func (n_ NEOnDemandRule) SetDnsSearchDomainMatch(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSearchDomainMatch:"), value)
}

// DNS server addresses that identify a network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/dnsserveraddressmatch
func (n_ NEOnDemandRule) DnsServerAddressMatch() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("dnsServerAddressMatch"))
	return rv
}


// SetDnsServerAddressMatch sets the value of the dnsServerAddressMatch property.
// DNS server addresses that identify a network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/dnsserveraddressmatch
func (n_ NEOnDemandRule) SetDnsServerAddressMatch(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsServerAddressMatch:"), value)
}

// An interface type to identify a network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/interfacetypematch
func (n_ NEOnDemandRule) InterfaceTypeMatch() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("interfaceTypeMatch"))
	return rv
}


// SetInterfaceTypeMatch sets the value of the interfaceTypeMatch property.
// An interface type to identify a network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/interfacetypematch
func (n_ NEOnDemandRule) SetInterfaceTypeMatch(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInterfaceTypeMatch:"), value)
}

// A URL to probe when all other network identifiers match to validate that an expected resource is available.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/probeurl
func (n_ NEOnDemandRule) ProbeURL() foundation.URL {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("probeURL"))
	return rv
}


// SetProbeURL sets the value of the probeURL property.
// A URL to probe when all other network identifiers match to validate that an expected resource is available.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/probeurl
func (n_ NEOnDemandRule) SetProbeURL(value foundation.IURL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProbeURL:"), value)
}

// SSIDs that identify a network.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/ssidmatch
func (n_ NEOnDemandRule) SsidMatch() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("ssidMatch"))
	return rv
}


// SetSsidMatch sets the value of the ssidMatch property.
// SSIDs that identify a network.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandrule/ssidmatch
func (n_ NEOnDemandRule) SetSsidMatch(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSsidMatch:"), value)
}



