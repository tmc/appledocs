// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [NEOnDemandRuleEvaluateConnection] class.
var (
	NEOnDemandRuleEvaluateConnectionClass     _NEOnDemandRuleEvaluateConnectionClass
	NEOnDemandRuleEvaluateConnectionClassOnce sync.Once
)

func getNEOnDemandRuleEvaluateConnectionClass() _NEOnDemandRuleEvaluateConnectionClass {
	NEOnDemandRuleEvaluateConnectionClassOnce.Do(func() {
		NEOnDemandRuleEvaluateConnectionClass = _NEOnDemandRuleEvaluateConnectionClass{objc.GetClass("NEOnDemandRuleEvaluateConnection")}
	})
	return NEOnDemandRuleEvaluateConnectionClass
}

type _NEOnDemandRuleEvaluateConnectionClass struct {
	class objc.Class
}





// An interface definition for the [NEOnDemandRuleEvaluateConnection] class.
type INEOnDemandRuleEvaluateConnection interface {
	INEOnDemandRule
	

	// properties:
	ConnectionRules() []NEEvaluateConnectionRule
	SetConnectionRules(value []NEEvaluateConnectionRule)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleEvaluateConnectionClass) Alloc() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEOnDemandRuleEvaluateConnectionClass) New() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEOnDemandRuleEvaluateConnection) Init() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEOnDemandRuleEvaluateConnection) Autorelease() NEOnDemandRuleEvaluateConnection {
	rv := objc.Send[NEOnDemandRuleEvaluateConnection](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleEvaluateConnection creates a new NEOnDemandRuleEvaluateConnection instance.
func NewNEOnDemandRuleEvaluateConnection() NEOnDemandRuleEvaluateConnection {
	return getNEOnDemandRuleEvaluateConnectionClass().New()
}





// A VPN On Demand rule that evaluate the app’s connection to determine whether to run its action.
//
// When rules of this class match, the properties of the network connection being established are matched against a set of connection rules. The action of the matched rule (if any) is used to determine whether or not the VPN will be started.


// A VPN On Demand rule that evaluate the app’s connection to determine whether to run its action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection
type NEOnDemandRuleEvaluateConnection struct {
	NEOnDemandRule
}

// NEOnDemandRuleEvaluateConnectionFrom constructs a [NEOnDemandRuleEvaluateConnection] from an unsafe.Pointer.
//
// A VPN On Demand rule that evaluate the app’s connection to determine whether to run its action.
func NEOnDemandRuleEvaluateConnectionFrom(ptr unsafe.Pointer) NEOnDemandRuleEvaluateConnection {
	return NEOnDemandRuleEvaluateConnection{
		NEOnDemandRule: NEOnDemandRuleFrom(ptr),
	}
}

























// An array of objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection/connectionRules
func (n_ NEOnDemandRuleEvaluateConnection) ConnectionRules() []NEEvaluateConnectionRule {
	rv := objc.Send[[]NEEvaluateConnectionRule](n_.ID, objc.Sel("connectionRules"))
	return rv
}


// An array of objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection/connectionRules
func (n_ NEOnDemandRuleEvaluateConnection) SetConnectionRules(value []NEEvaluateConnectionRule) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setConnectionRules:"), nsArray)
}








