// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEFilterRule] class.
var (
	NEFilterRuleClass     _NEFilterRuleClass
	NEFilterRuleClassOnce sync.Once
)

func getNEFilterRuleClass() _NEFilterRuleClass {
	NEFilterRuleClassOnce.Do(func() {
		NEFilterRuleClass = _NEFilterRuleClass{objc.GetClass("NEFilterRule")}
	})
	return NEFilterRuleClass
}

type _NEFilterRuleClass struct {
	class objc.Class
}





// An interface definition for the [NEFilterRule] class.
type INEFilterRule interface {
	objectivec.IObject
	

	// properties:
	Action() NEFilterAction
	NetworkRule() INENetworkRule


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEFilterRuleClass) Alloc() NEFilterRule {
	rv := objc.Send[NEFilterRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterRuleClass) New() NEFilterRule {
	rv := objc.Send[NEFilterRule](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterRule) Init() NEFilterRule {
	rv := objc.Send[NEFilterRule](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterRule) Autorelease() NEFilterRule {
	rv := objc.Send[NEFilterRule](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterRule creates a new NEFilterRule instance.
func NewNEFilterRule() NEFilterRule {
	return getNEFilterRuleClass().New()
}





// A rule for filters that combines a rule to match network traffic and an action to take when the rule matches.


// A rule for filters that combines a rule to match network traffic and an action to take when the rule matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule
type NEFilterRule struct {
	objectivec.Object
}

// NEFilterRuleFrom constructs a [NEFilterRule] from an unsafe.Pointer.
//
// A rule for filters that combines a rule to match network traffic and an action to take when the rule matches.
func NEFilterRuleFrom(ptr unsafe.Pointer) NEFilterRule {
	return NEFilterRule{objectivec.Object{objc.ID(ptr)}}
}






// Creates a new filter rule from a network rule and an action to take when network traffic matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/init(networkRule:action:)
func NewNEFilterRuleWithNetworkRuleAction(networkRule INENetworkRule, action NEFilterAction) NEFilterRule {
	instance := getNEFilterRuleClass().Alloc()
	rv := objc.Send[NEFilterRule](instance.ID, objc.Sel("initWithNetworkRule:action:"), networkRule, action)
	rv.Autorelease()
	return rv
}






















// The action to take when this rule matches network traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/action
func (n_ NEFilterRule) Action() NEFilterAction {
	rv := objc.Send[NEFilterAction](n_.ID, objc.Sel("action"))
	return rv
}


// The network rule that defines the network traffic characteristics that this filter rule matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/networkRule
func (n_ NEFilterRule) NetworkRule() INENetworkRule {
	rv := objc.Send[NENetworkRule](n_.ID, objc.Sel("networkRule"))
	return rv
}







