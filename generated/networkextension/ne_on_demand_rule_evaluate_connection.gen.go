// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NEOnDemandRuleEvaluateConnection */


/* debug [class_header]: Header for NEOnDemandRuleEvaluateConnection */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEOnDemandRuleEvaluateConnection */
// An interface definition for the [NEOnDemandRuleEvaluateConnection] class.
type INEOnDemandRuleEvaluateConnection interface {
	INEOnDemandRule
	
/* debug [class_interface_properties]: Properties for NEOnDemandRuleEvaluateConnection */
	// properties:
	ConnectionRules() []NEEvaluateConnectionRule
	SetConnectionRules(value []NEEvaluateConnectionRule)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEOnDemandRuleEvaluateConnection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEOnDemandRuleEvaluateConnection */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEOnDemandRuleEvaluateConnection */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEOnDemandRuleEvaluateConnection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEOnDemandRuleEvaluateConnection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEOnDemandRuleEvaluateConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEOnDemandRuleEvaluateConnection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEOnDemandRuleEvaluateConnection */

// An array of objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleEvaluateConnection/connectionRules
func (n_ NEOnDemandRuleEvaluateConnection) ConnectionRules() []NEEvaluateConnectionRule {
	rv := objc.Send[[]NEEvaluateConnectionRule](n_.ID, objc.Sel("connectionRules"))
	return rv
}/* debug [instance_properties/getter]: connectionRules */


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
}/* debug [instance_properties/setter]: connectionRules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEOnDemandRuleEvaluateConnection */



