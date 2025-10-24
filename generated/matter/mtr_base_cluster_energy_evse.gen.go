// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterEnergyEVSE */


/* debug [class_header]: Header for MTRBaseClusterEnergyEVSE */
// The class instance for the [MTRBaseClusterEnergyEVSE] class.
var (
	MTRBaseClusterEnergyEVSEClass     _MTRBaseClusterEnergyEVSEClass
	MTRBaseClusterEnergyEVSEClassOnce sync.Once
)

func getMTRBaseClusterEnergyEVSEClass() _MTRBaseClusterEnergyEVSEClass {
	MTRBaseClusterEnergyEVSEClassOnce.Do(func() {
		MTRBaseClusterEnergyEVSEClass = _MTRBaseClusterEnergyEVSEClass{objc.GetClass("MTRBaseClusterEnergyEVSE")}
	})
	return MTRBaseClusterEnergyEVSEClass
}

type _MTRBaseClusterEnergyEVSEClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterEnergyEVSE */
// An interface definition for the [MTRBaseClusterEnergyEVSE] class.
type IMTRBaseClusterEnergyEVSE interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterEnergyEVSE */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterEnergyEVSE */
	// methods:
	ClearTargetsWithCompletion(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterEnergyEVSE */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterEnergyEVSEClass) Alloc() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterEnergyEVSEClass) New() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterEnergyEVSE) Init() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterEnergyEVSE) Autorelease() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterEnergyEVSE creates a new MTRBaseClusterEnergyEVSE instance.
func NewMTRBaseClusterEnergyEVSE() MTRBaseClusterEnergyEVSE {
	return getMTRBaseClusterEnergyEVSEClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterEnergyEVSE */
// Cluster Energy EVSE
//
// Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.


// Cluster Energy EVSE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE
type MTRBaseClusterEnergyEVSE struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterEnergyEVSEFrom constructs a [MTRBaseClusterEnergyEVSE] from an unsafe.Pointer.
//
// Cluster Energy EVSE
func MTRBaseClusterEnergyEVSEFrom(ptr unsafe.Pointer) MTRBaseClusterEnergyEVSE {
	return MTRBaseClusterEnergyEVSE{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterEnergyEVSE *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterEnergyEVSE */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterEnergyEVSE */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterEnergyEVSE */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/clearTargets(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ClearTargetsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ClearTargetsWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterEnergyEVSE */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterEnergyEVSE */



