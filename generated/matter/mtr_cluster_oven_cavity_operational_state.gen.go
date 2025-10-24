// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterOvenCavityOperationalState */


/* debug [class_header]: Header for MTRClusterOvenCavityOperationalState */
// The class instance for the [MTRClusterOvenCavityOperationalState] class.
var (
	MTRClusterOvenCavityOperationalStateClass     _MTRClusterOvenCavityOperationalStateClass
	MTRClusterOvenCavityOperationalStateClassOnce sync.Once
)

func getMTRClusterOvenCavityOperationalStateClass() _MTRClusterOvenCavityOperationalStateClass {
	MTRClusterOvenCavityOperationalStateClassOnce.Do(func() {
		MTRClusterOvenCavityOperationalStateClass = _MTRClusterOvenCavityOperationalStateClass{objc.GetClass("MTRClusterOvenCavityOperationalState")}
	})
	return MTRClusterOvenCavityOperationalStateClass
}

type _MTRClusterOvenCavityOperationalStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterOvenCavityOperationalState */
// An interface definition for the [MTRClusterOvenCavityOperationalState] class.
type IMTRClusterOvenCavityOperationalState interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterOvenCavityOperationalState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterOvenCavityOperationalState */
	// methods:
	ReadAttributeOperationalStateListWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterOvenCavityOperationalState */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOvenCavityOperationalStateClass) Alloc() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterOvenCavityOperationalStateClass) New() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOvenCavityOperationalState) Init() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOvenCavityOperationalState) Autorelease() MTRClusterOvenCavityOperationalState {
	rv := objc.Send[MTRClusterOvenCavityOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOvenCavityOperationalState creates a new MTRClusterOvenCavityOperationalState instance.
func NewMTRClusterOvenCavityOperationalState() MTRClusterOvenCavityOperationalState {
	return getMTRClusterOvenCavityOperationalStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterOvenCavityOperationalState */
// Cluster Oven Cavity Operational State This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.


// Cluster Oven Cavity Operational State This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState
type MTRClusterOvenCavityOperationalState struct {
	MTRGenericCluster
}

// MTRClusterOvenCavityOperationalStateFrom constructs a [MTRClusterOvenCavityOperationalState] from an unsafe.Pointer.
//
// Cluster Oven Cavity Operational State This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.
func MTRClusterOvenCavityOperationalStateFrom(ptr unsafe.Pointer) MTRClusterOvenCavityOperationalState {
	return MTRClusterOvenCavityOperationalState{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterOvenCavityOperationalState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterOvenCavityOperationalState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterOvenCavityOperationalState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterOvenCavityOperationalState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOvenCavityOperationalState/readAttributeOperationalStateList(with:)
func (m_ MTRClusterOvenCavityOperationalState) ReadAttributeOperationalStateListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeOperationalStateListWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeOperationalStateListWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterOvenCavityOperationalState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterOvenCavityOperationalState */



