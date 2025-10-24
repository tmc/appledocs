// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NEOnDemandRuleConnect */


/* debug [class_header]: Header for NEOnDemandRuleConnect */
// The class instance for the [NEOnDemandRuleConnect] class.
var (
	NEOnDemandRuleConnectClass     _NEOnDemandRuleConnectClass
	NEOnDemandRuleConnectClassOnce sync.Once
)

func getNEOnDemandRuleConnectClass() _NEOnDemandRuleConnectClass {
	NEOnDemandRuleConnectClassOnce.Do(func() {
		NEOnDemandRuleConnectClass = _NEOnDemandRuleConnectClass{objc.GetClass("NEOnDemandRuleConnect")}
	})
	return NEOnDemandRuleConnectClass
}

type _NEOnDemandRuleConnectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEOnDemandRuleConnect */
// An interface definition for the [NEOnDemandRuleConnect] class.
type INEOnDemandRuleConnect interface {
	INEOnDemandRule
	
/* debug [class_interface_properties]: Properties for NEOnDemandRuleConnect */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEOnDemandRuleConnect */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEOnDemandRuleConnect */
// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleConnectClass) Alloc() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEOnDemandRuleConnectClass) New() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEOnDemandRuleConnect) Init() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEOnDemandRuleConnect) Autorelease() NEOnDemandRuleConnect {
	rv := objc.Send[NEOnDemandRuleConnect](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEOnDemandRuleConnect creates a new NEOnDemandRuleConnect instance.
func NewNEOnDemandRuleConnect() NEOnDemandRuleConnect {
	return getNEOnDemandRuleConnectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEOnDemandRuleConnect */
// A VPN On Demand rule that connects the VPN.
//
// When rules of this class match, the system starts the VPN connection whenever an application running on the system opens a network connection.


// A VPN On Demand rule that connects the VPN.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEOnDemandRuleConnect
type NEOnDemandRuleConnect struct {
	NEOnDemandRule
}

// NEOnDemandRuleConnectFrom constructs a [NEOnDemandRuleConnect] from an unsafe.Pointer.
//
// A VPN On Demand rule that connects the VPN.
func NEOnDemandRuleConnectFrom(ptr unsafe.Pointer) NEOnDemandRuleConnect {
	return NEOnDemandRuleConnect{
		NEOnDemandRule: NEOnDemandRuleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEOnDemandRuleConnect *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEOnDemandRuleConnect */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEOnDemandRuleConnect */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEOnDemandRuleConnect */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEOnDemandRuleConnect */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEOnDemandRuleConnect */



