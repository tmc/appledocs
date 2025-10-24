// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEOnDemandRule */


/* debug [class_header]: Header for NEOnDemandRule */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEOnDemandRule */
// An interface definition for the [NEOnDemandRule] class.
type INEOnDemandRule interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEOnDemandRule */
	// properties:
	Action() NEOnDemandRuleAction
	DNSSearchDomainMatch() []string
	SetDNSSearchDomainMatch(value []string)
	DNSServerAddressMatch() []string
	SetDNSServerAddressMatch(value []string)
	InterfaceTypeMatch() NEOnDemandRuleInterfaceType
	SetInterfaceTypeMatch(value NEOnDemandRuleInterfaceType)
	ProbeURL() objc.IObject /* cross-framework: NSURL */
	SetProbeURL(value objc.IObject /* cross-framework: NSURL */)
	SSIDMatch() []string
	SetSSIDMatch(value []string)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEOnDemandRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEOnDemandRule */
// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleClass) Alloc() NEOnDemandRule {
	rv := objc.Send[NEOnDemandRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEOnDemandRule */
// A base class shared by all VPN On Demand rules.
//
// Each rule is defined by a single action and a set of optional matching conditions. The action defines how the system should trigger the VPN when the conditions are met, such as connecting automatically for all connections, connecting conditionally, or disconnecting. The optional conditions describe parameters of a network. Some common rules include disconnecting the VPN on a trusted, internal network, and triggering on all other networks. When rules are defined in an array, they are evaluated in order and the action of the first rule to match all conditions is chosen. Instances of the class should be created through one of its subclasses: , , , or .


// A base class shared by all VPN On Demand rules.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEOnDemandRule *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEOnDemandRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEOnDemandRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEOnDemandRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEOnDemandRule */

// The action of the On Demand Rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/action
func (n_ NEOnDemandRule) Action() NEOnDemandRuleAction {
	rv := objc.Send[NEOnDemandRuleAction](n_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// DNS search domains that identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/dnsSearchDomainMatch
func (n_ NEOnDemandRule) DNSSearchDomainMatch() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("DNSSearchDomainMatch"))
	return rv
}/* debug [instance_properties/getter]: DNSSearchDomainMatch */


// DNS search domains that identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/dnsSearchDomainMatch
func (n_ NEOnDemandRule) SetDNSSearchDomainMatch(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setDNSSearchDomainMatch:"), nsArray)
}/* debug [instance_properties/setter]: DNSSearchDomainMatch */


// DNS server addresses that identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/dnsServerAddressMatch
func (n_ NEOnDemandRule) DNSServerAddressMatch() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("DNSServerAddressMatch"))
	return rv
}/* debug [instance_properties/getter]: DNSServerAddressMatch */


// DNS server addresses that identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/dnsServerAddressMatch
func (n_ NEOnDemandRule) SetDNSServerAddressMatch(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setDNSServerAddressMatch:"), nsArray)
}/* debug [instance_properties/setter]: DNSServerAddressMatch */


// An interface type to identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/interfaceTypeMatch
func (n_ NEOnDemandRule) InterfaceTypeMatch() NEOnDemandRuleInterfaceType {
	rv := objc.Send[NEOnDemandRuleInterfaceType](n_.ID, objc.Sel("interfaceTypeMatch"))
	return rv
}/* debug [instance_properties/getter]: interfaceTypeMatch */


// An interface type to identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/interfaceTypeMatch
func (n_ NEOnDemandRule) SetInterfaceTypeMatch(value NEOnDemandRuleInterfaceType) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInterfaceTypeMatch:"), value)
}/* debug [instance_properties/setter]: interfaceTypeMatch */


// A URL to probe when all other network identifiers match to validate that an expected resource is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/probeURL
func (n_ NEOnDemandRule) ProbeURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("probeURL"))
	return rv
}/* debug [instance_properties/getter]: probeURL */


// A URL to probe when all other network identifiers match to validate that an expected resource is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/probeURL
func (n_ NEOnDemandRule) SetProbeURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProbeURL:"), value)
}/* debug [instance_properties/setter]: probeURL */


// SSIDs that identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/ssidMatch
func (n_ NEOnDemandRule) SSIDMatch() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("SSIDMatch"))
	return rv
}/* debug [instance_properties/getter]: SSIDMatch */


// SSIDs that identify a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRule/ssidMatch
func (n_ NEOnDemandRule) SetSSIDMatch(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setSSIDMatch:"), nsArray)
}/* debug [instance_properties/setter]: SSIDMatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEOnDemandRule */



