// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterRule */


/* debug [class_header]: Header for NEFilterRule */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterRule */
// An interface definition for the [NEFilterRule] class.
type INEFilterRule interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterRule */
	// properties:
	Action() NEFilterAction
	NetworkRule() INENetworkRule
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterRule */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterRule */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterRule */

// Creates a new filter rule from a network rule and an action to take when network traffic matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/init(networkRule:action:)
func NewNEFilterRuleWithNetworkRuleAction(networkRule INENetworkRule, action NEFilterAction) NEFilterRule {
	instance := getNEFilterRuleClass().Alloc()
	rv := objc.Send[NEFilterRule](instance.ID, objc.Sel("initWithNetworkRule:action:"), networkRule, action)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEFilterRuleWithNetworkRuleAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterRule */

// The action to take when this rule matches network traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/action
func (n_ NEFilterRule) Action() NEFilterAction {
	rv := objc.Send[NEFilterAction](n_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The network rule that defines the network traffic characteristics that this filter rule matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRule/networkRule
func (n_ NEFilterRule) NetworkRule() INENetworkRule {
	rv := objc.Send[NENetworkRule](n_.ID, objc.Sel("networkRule"))
	return rv
}/* debug [instance_properties/getter]: networkRule */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterRule */


