// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




