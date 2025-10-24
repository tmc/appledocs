// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NEFilterControlVerdict */


/* debug [class_header]: Header for NEFilterControlVerdict */
// The class instance for the [NEFilterControlVerdict] class.
var (
	NEFilterControlVerdictClass     _NEFilterControlVerdictClass
	NEFilterControlVerdictClassOnce sync.Once
)

func getNEFilterControlVerdictClass() _NEFilterControlVerdictClass {
	NEFilterControlVerdictClassOnce.Do(func() {
		NEFilterControlVerdictClass = _NEFilterControlVerdictClass{objc.GetClass("NEFilterControlVerdict")}
	})
	return NEFilterControlVerdictClass
}

type _NEFilterControlVerdictClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterControlVerdict */
// An interface definition for the [NEFilterControlVerdict] class.
type INEFilterControlVerdict interface {
	INEFilterNewFlowVerdict
	
/* debug [class_interface_properties]: Properties for NEFilterControlVerdict */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterControlVerdict */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterControlVerdict */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterControlVerdictClass) Alloc() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterControlVerdictClass) New() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterControlVerdict) Init() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterControlVerdict) Autorelease() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterControlVerdict creates a new NEFilterControlVerdict instance.
func NewNEFilterControlVerdict() NEFilterControlVerdict {
	return getNEFilterControlVerdictClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterControlVerdict */
// The result from a filter control provider.


// The result from a filter control provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict
type NEFilterControlVerdict struct {
	NEFilterNewFlowVerdict
}

// NEFilterControlVerdictFrom constructs a [NEFilterControlVerdict] from an unsafe.Pointer.
//
// The result from a filter control provider.
func NEFilterControlVerdictFrom(ptr unsafe.Pointer) NEFilterControlVerdict {
	return NEFilterControlVerdict{
		NEFilterNewFlowVerdict: NEFilterNewFlowVerdictFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterControlVerdict *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterControlVerdict */

// Create a verdict that indicates to the system that all of the flow’s data should be allowed to pass to its final destination, and that the filtering rules have been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/allow(withUpdateRules:)
func (nc _NEFilterControlVerdictClass) AllowVerdictWithUpdateRules(updateRules bool) NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("allowVerdictWithUpdateRules:"), updateRules)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllowVerdictWithUpdateRules) */


// Create a verdict that indicates to the system that all of the flow’s data should be dropped, and that the filtering rules have been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/drop(withUpdateRules:)
func (nc _NEFilterControlVerdictClass) DropVerdictWithUpdateRules(updateRules bool) NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("dropVerdictWithUpdateRules:"), updateRules)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DropVerdictWithUpdateRules) */


// Create a verdict that indicates to the system that the filtering rules have been updated, and that the Filter Data Provider needs to make a decision about the flow’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlVerdict/updateRules()
func (nc _NEFilterControlVerdictClass) UpdateRules() NEFilterControlVerdict {
	rv := objc.Send[NEFilterControlVerdict](objc.ID(nc.class), objc.Sel("updateRules"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UpdateRules) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterControlVerdict */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterControlVerdict */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterControlVerdict */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterControlVerdict */


