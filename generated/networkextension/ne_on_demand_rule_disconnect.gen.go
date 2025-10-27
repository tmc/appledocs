// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEOnDemandRuleDisconnect] class.
var (
	NEOnDemandRuleDisconnectClass     _NEOnDemandRuleDisconnectClass
	NEOnDemandRuleDisconnectClassOnce sync.Once
)

func getNEOnDemandRuleDisconnectClass() _NEOnDemandRuleDisconnectClass {
	NEOnDemandRuleDisconnectClassOnce.Do(func() {
		NEOnDemandRuleDisconnectClass = _NEOnDemandRuleDisconnectClass{objc.GetClass("NEOnDemandRuleDisconnect")}
	})
	return NEOnDemandRuleDisconnectClass
}

type _NEOnDemandRuleDisconnectClass struct {
	class objc.Class
}





// An interface definition for the [NEOnDemandRuleDisconnect] class.
type INEOnDemandRuleDisconnect interface {
	INEOnDemandRule
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleDisconnectClass) Alloc() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEOnDemandRuleDisconnectClass) New() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEOnDemandRuleDisconnect) Init() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEOnDemandRuleDisconnect) Autorelease() NEOnDemandRuleDisconnect {
	rv := objc.Send[NEOnDemandRuleDisconnect](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleDisconnect creates a new NEOnDemandRuleDisconnect instance.
func NewNEOnDemandRuleDisconnect() NEOnDemandRuleDisconnect {
	return getNEOnDemandRuleDisconnectClass().New()
}





// A VPN On Demand rule that disconnects the VPN.
//
// When rules of this class match, the VPN connection is not started, and the VPN connection is disconnected if it is not already disconnected.


// A VPN On Demand rule that disconnects the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleDisconnect
type NEOnDemandRuleDisconnect struct {
	NEOnDemandRule
}

// NEOnDemandRuleDisconnectFrom constructs a [NEOnDemandRuleDisconnect] from an unsafe.Pointer.
//
// A VPN On Demand rule that disconnects the VPN.
func NEOnDemandRuleDisconnectFrom(ptr unsafe.Pointer) NEOnDemandRuleDisconnect {
	return NEOnDemandRuleDisconnect{
		NEOnDemandRule: NEOnDemandRuleFrom(ptr),
	}
}































