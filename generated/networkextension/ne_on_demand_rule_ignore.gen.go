// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEOnDemandRuleIgnore] class.
var (
	NEOnDemandRuleIgnoreClass     _NEOnDemandRuleIgnoreClass
	NEOnDemandRuleIgnoreClassOnce sync.Once
)

func getNEOnDemandRuleIgnoreClass() _NEOnDemandRuleIgnoreClass {
	NEOnDemandRuleIgnoreClassOnce.Do(func() {
		NEOnDemandRuleIgnoreClass = _NEOnDemandRuleIgnoreClass{objc.GetClass("NEOnDemandRuleIgnore")}
	})
	return NEOnDemandRuleIgnoreClass
}

type _NEOnDemandRuleIgnoreClass struct {
	class objc.Class
}

// An interface definition for the [NEOnDemandRuleIgnore] class.
type INEOnDemandRuleIgnore interface {
	INEOnDemandRule
}

// A VPN On Demand rule that doesn’t change the status of the VPN.
//
// When rules of this class match, the VPN connection is not started, and the current status of the VPN connection is left unchanged.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleIgnore
type NEOnDemandRuleIgnore struct {
	NEOnDemandRule
}

// NEOnDemandRuleIgnoreFrom constructs a [NEOnDemandRuleIgnore] from an unsafe.Pointer.
//
// A VPN On Demand rule that doesn’t change the status of the VPN.
func NEOnDemandRuleIgnoreFrom(ptr unsafe.Pointer) NEOnDemandRuleIgnore {
	return NEOnDemandRuleIgnore{
		NEOnDemandRule: NEOnDemandRuleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleIgnoreClass) Alloc() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEOnDemandRuleIgnoreClass) New() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEOnDemandRuleIgnore) Init() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEOnDemandRuleIgnore) Autorelease() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleIgnore creates a new NEOnDemandRuleIgnore instance.
func NewNEOnDemandRuleIgnore() NEOnDemandRuleIgnore {
	return getNEOnDemandRuleIgnoreClass().New()
}




