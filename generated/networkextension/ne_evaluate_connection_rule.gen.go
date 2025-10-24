// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEEvaluateConnectionRule */


/* debug [class_header]: Header for NEEvaluateConnectionRule */
// The class instance for the [NEEvaluateConnectionRule] class.
var (
	NEEvaluateConnectionRuleClass     _NEEvaluateConnectionRuleClass
	NEEvaluateConnectionRuleClassOnce sync.Once
)

func getNEEvaluateConnectionRuleClass() _NEEvaluateConnectionRuleClass {
	NEEvaluateConnectionRuleClassOnce.Do(func() {
		NEEvaluateConnectionRuleClass = _NEEvaluateConnectionRuleClass{objc.GetClass("NEEvaluateConnectionRule")}
	})
	return NEEvaluateConnectionRuleClass
}

type _NEEvaluateConnectionRuleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEEvaluateConnectionRule */
// An interface definition for the [NEEvaluateConnectionRule] class.
type INEEvaluateConnectionRule interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEEvaluateConnectionRule */
	// properties:
	Action() NEEvaluateConnectionRuleAction
	MatchDomains() []string
	ProbeURL() objc.IObject /* cross-framework: NSURL */
	SetProbeURL(value objc.IObject /* cross-framework: NSURL */)
	UseDNSServers() []string
	SetUseDNSServers(value []string)
	ConnectionRules() INEEvaluateConnectionRule
	SetConnectionRules(value INEEvaluateConnectionRule)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEEvaluateConnectionRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEEvaluateConnectionRule */
// Alloc allocates a new instance without initialization.
func (nc _NEEvaluateConnectionRuleClass) Alloc() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEEvaluateConnectionRuleClass) New() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEEvaluateConnectionRule) Init() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEEvaluateConnectionRule) Autorelease() NEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEEvaluateConnectionRule creates a new NEEvaluateConnectionRule instance.
func NewNEEvaluateConnectionRule() NEEvaluateConnectionRule {
	return getNEEvaluateConnectionRuleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEEvaluateConnectionRule */
// associates properties of network connections with an action.


// associates properties of network connections with an action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule
type NEEvaluateConnectionRule struct {
	objectivec.Object
}

// NEEvaluateConnectionRuleFrom constructs a [NEEvaluateConnectionRule] from an unsafe.Pointer.
//
// associates properties of network connections with an action.
func NEEvaluateConnectionRuleFrom(ptr unsafe.Pointer) NEEvaluateConnectionRule {
	return NEEvaluateConnectionRule{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEEvaluateConnectionRule */

// Initialize an instance with a list of destination host domains and an action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/init(matchDomains:andAction:)
func NewNEEvaluateConnectionRuleWithMatchDomainsAndAction(domains []string, action NEEvaluateConnectionRuleAction) NEEvaluateConnectionRule {
	instance := getNEEvaluateConnectionRuleClass().Alloc()
	rv := objc.Send[NEEvaluateConnectionRule](instance.ID, objc.Sel("initWithMatchDomains:andAction:"), domains, action)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEEvaluateConnectionRuleWithMatchDomainsAndAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEEvaluateConnectionRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEEvaluateConnectionRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEEvaluateConnectionRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEEvaluateConnectionRule */

// The action to take if the properties of the network connection being established match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/action
func (n_ NEEvaluateConnectionRule) Action() NEEvaluateConnectionRuleAction {
	rv := objc.Send[NEEvaluateConnectionRuleAction](n_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// An array of domains used to match the destination hostname of connections. If the destination hostname of a connection matches any of the domains in the array, then the connection matches the rule. Each domain is matched against the destination hostname using suffix matching, and each label in the domain must match an entire label in the hostname. For example, the domain will match the hostname but not .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/matchDomains
func (n_ NEEvaluateConnectionRule) MatchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchDomains"))
	return rv
}/* debug [instance_properties/getter]: matchDomains */


// An HTTP or HTTPS URL. If the rule matches the connection being established and the action is and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/probeURL
func (n_ NEEvaluateConnectionRule) ProbeURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("probeURL"))
	return rv
}/* debug [instance_properties/getter]: probeURL */


// An HTTP or HTTPS URL. If the rule matches the connection being established and the action is and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/probeURL
func (n_ NEEvaluateConnectionRule) SetProbeURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProbeURL:"), value)
}/* debug [instance_properties/setter]: probeURL */


// If the rule matches the connection being established and the action is , the DNS servers specified in this array are used to resolve the destination hostname of the connection while evaluating connectivity to the destination of the connection. If the resolution fails for any reason, the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/useDNSServers
func (n_ NEEvaluateConnectionRule) UseDNSServers() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("useDNSServers"))
	return rv
}/* debug [instance_properties/getter]: useDNSServers */


// If the rule matches the connection being established and the action is , the DNS servers specified in this array are used to resolve the destination hostname of the connection while evaluating connectivity to the destination of the connection. If the resolution fails for any reason, the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/useDNSServers
func (n_ NEEvaluateConnectionRule) SetUseDNSServers(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setUseDNSServers:"), nsArray)
}/* debug [instance_properties/setter]: useDNSServers */


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandruleevaluateconnection/connectionrules
func (n_ NEEvaluateConnectionRule) ConnectionRules() INEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](n_.ID, objc.Sel("connectionRules"))
	return rv
}/* debug [instance_properties/getter]: connectionRules */


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandruleevaluateconnection/connectionrules
func (n_ NEEvaluateConnectionRule) SetConnectionRules(value INEEvaluateConnectionRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnectionRules:"), value)
}/* debug [instance_properties/setter]: connectionRules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEEvaluateConnectionRule */


