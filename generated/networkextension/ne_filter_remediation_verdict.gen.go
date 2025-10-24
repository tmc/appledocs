// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NEFilterRemediationVerdict */


/* debug [class_header]: Header for NEFilterRemediationVerdict */
// The class instance for the [NEFilterRemediationVerdict] class.
var (
	NEFilterRemediationVerdictClass     _NEFilterRemediationVerdictClass
	NEFilterRemediationVerdictClassOnce sync.Once
)

func getNEFilterRemediationVerdictClass() _NEFilterRemediationVerdictClass {
	NEFilterRemediationVerdictClassOnce.Do(func() {
		NEFilterRemediationVerdictClass = _NEFilterRemediationVerdictClass{objc.GetClass("NEFilterRemediationVerdict")}
	})
	return NEFilterRemediationVerdictClass
}

type _NEFilterRemediationVerdictClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterRemediationVerdict */
// An interface definition for the [NEFilterRemediationVerdict] class.
type INEFilterRemediationVerdict interface {
	INEFilterVerdict
	
/* debug [class_interface_properties]: Properties for NEFilterRemediationVerdict */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterRemediationVerdict */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterRemediationVerdict */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterRemediationVerdictClass) Alloc() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterRemediationVerdictClass) New() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterRemediationVerdict) Init() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterRemediationVerdict) Autorelease() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterRemediationVerdict creates a new NEFilterRemediationVerdict instance.
func NewNEFilterRemediationVerdict() NEFilterRemediationVerdict {
	return getNEFilterRemediationVerdictClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterRemediationVerdict */
// The result from a filter data provider after the user requests remediation for a blocked flow.


// The result from a filter data provider after the user requests remediation for a blocked flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict
type NEFilterRemediationVerdict struct {
	NEFilterVerdict
}

// NEFilterRemediationVerdictFrom constructs a [NEFilterRemediationVerdict] from an unsafe.Pointer.
//
// The result from a filter data provider after the user requests remediation for a blocked flow.
func NEFilterRemediationVerdictFrom(ptr unsafe.Pointer) NEFilterRemediationVerdict {
	return NEFilterRemediationVerdict{
		NEFilterVerdict: NEFilterVerdictFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterRemediationVerdict *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterRemediationVerdict */

// Create a verdict that indicates to the system that the Filter Data Provider will allow the flow to pass to its final destination when/if the flow is requested again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict/allow()
func (nc _NEFilterRemediationVerdictClass) AllowVerdict() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("allowVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllowVerdict) */


// Create a verdict that indicates to the system that the Filter Data Provider will continue to block the flow of network data if/when the flow is requested again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict/drop()
func (nc _NEFilterRemediationVerdictClass) DropVerdict() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("dropVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DropVerdict) */


// Create a verdict that indicates to the system that the Filter Data Provider needs the filtering rules to be updated before it can make a remediation decision about the current flow of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterRemediationVerdict/needRules()
func (nc _NEFilterRemediationVerdictClass) NeedRulesVerdict() NEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](objc.ID(nc.class), objc.Sel("needRulesVerdict"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NeedRulesVerdict) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterRemediationVerdict */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterRemediationVerdict */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterRemediationVerdict */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterRemediationVerdict */


