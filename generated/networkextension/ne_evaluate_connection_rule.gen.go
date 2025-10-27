// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NEEvaluateConnectionRule] class.
type INEEvaluateConnectionRule interface {
	objectivec.IObject
	

	// properties:
	Action() NEEvaluateConnectionRuleAction
	MatchDomains() []string
	ProbeURL() foundation.foundation.INSURL
	SetProbeURL(value foundation.foundation.INSURL)
	UseDNSServers() []string
	SetUseDNSServers(value []string)
	ConnectionRules() INEEvaluateConnectionRule
	SetConnectionRules(value INEEvaluateConnectionRule)


	

	// methods:


}





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






// Initialize an instance with a list of destination host domains and an action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/init(matchDomains:andAction:)
func NewNEEvaluateConnectionRuleWithMatchDomainsAndAction(domains []string, action NEEvaluateConnectionRuleAction) NEEvaluateConnectionRule {
	instance := getNEEvaluateConnectionRuleClass().Alloc()
	rv := objc.Send[NEEvaluateConnectionRule](instance.ID, objc.Sel("initWithMatchDomains:andAction:"), domains, action)
	rv.Autorelease()
	return rv
}






















// The action to take if the properties of the network connection being established match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/action
func (n_ NEEvaluateConnectionRule) Action() NEEvaluateConnectionRuleAction {
	rv := objc.Send[NEEvaluateConnectionRuleAction](n_.ID, objc.Sel("action"))
	return rv
}


// An array of domains used to match the destination hostname of connections. If the destination hostname of a connection matches any of the domains in the array, then the connection matches the rule. Each domain is matched against the destination hostname using suffix matching, and each label in the domain must match an entire label in the hostname. For example, the domain will match the hostname but not .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/matchDomains
func (n_ NEEvaluateConnectionRule) MatchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// An HTTP or HTTPS URL. If the rule matches the connection being established and the action is and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/probeURL
func (n_ NEEvaluateConnectionRule) ProbeURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("probeURL"))
	return rv
}


// An HTTP or HTTPS URL. If the rule matches the connection being established and the action is and a request sent to this URL results in a response with an HTTP response code other than 200, then the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/probeURL
func (n_ NEEvaluateConnectionRule) SetProbeURL(value foundation.foundation.INSURL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProbeURL:"), value)
}


// If the rule matches the connection being established and the action is , the DNS servers specified in this array are used to resolve the destination hostname of the connection while evaluating connectivity to the destination of the connection. If the resolution fails for any reason, the VPN is started.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEvaluateConnectionRule/useDNSServers
func (n_ NEEvaluateConnectionRule) UseDNSServers() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("useDNSServers"))
	return rv
}


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
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandruleevaluateconnection/connectionrules
func (n_ NEEvaluateConnectionRule) ConnectionRules() INEEvaluateConnectionRule {
	rv := objc.Send[NEEvaluateConnectionRule](n_.ID, objc.Sel("connectionRules"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neondemandruleevaluateconnection/connectionrules
func (n_ NEEvaluateConnectionRule) SetConnectionRules(value INEEvaluateConnectionRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnectionRules:"), value)
}







