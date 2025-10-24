// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NEOnDemandRuleIgnore */


/* debug [class_header]: Header for NEOnDemandRuleIgnore */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEOnDemandRuleIgnore */
// An interface definition for the [NEOnDemandRuleIgnore] class.
type INEOnDemandRuleIgnore interface {
	INEOnDemandRule
	
/* debug [class_interface_properties]: Properties for NEOnDemandRuleIgnore */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEOnDemandRuleIgnore */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEOnDemandRuleIgnore */
// Alloc allocates a new instance without initialization.
func (nc _NEOnDemandRuleIgnoreClass) Alloc() NEOnDemandRuleIgnore {
	rv := objc.Send[NEOnDemandRuleIgnore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEOnDemandRuleIgnore */
// A VPN On Demand rule that doesn’t change the status of the VPN.
//
// When rules of this class match, the VPN connection is not started, and the current status of the VPN connection is left unchanged.


// A VPN On Demand rule that doesn’t change the status of the VPN.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEOnDemandRuleIgnore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEOnDemandRuleIgnore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEOnDemandRuleIgnore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEOnDemandRuleIgnore */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEOnDemandRuleIgnore */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEOnDemandRuleIgnore */



