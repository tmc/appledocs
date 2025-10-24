// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterEnergyEVSE */


/* debug [class_header]: Header for MTRClusterEnergyEVSE */
// The class instance for the [MTRClusterEnergyEVSE] class.
var (
	MTRClusterEnergyEVSEClass     _MTRClusterEnergyEVSEClass
	MTRClusterEnergyEVSEClassOnce sync.Once
)

func getMTRClusterEnergyEVSEClass() _MTRClusterEnergyEVSEClass {
	MTRClusterEnergyEVSEClassOnce.Do(func() {
		MTRClusterEnergyEVSEClass = _MTRClusterEnergyEVSEClass{objc.GetClass("MTRClusterEnergyEVSE")}
	})
	return MTRClusterEnergyEVSEClass
}

type _MTRClusterEnergyEVSEClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterEnergyEVSE */
// An interface definition for the [MTRClusterEnergyEVSE] class.
type IMTRClusterEnergyEVSE interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterEnergyEVSE */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterEnergyEVSE */
	// methods:
	DisableWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterEnergyEVSE */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterEnergyEVSEClass) Alloc() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterEnergyEVSEClass) New() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterEnergyEVSE) Init() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterEnergyEVSE) Autorelease() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterEnergyEVSE creates a new MTRClusterEnergyEVSE instance.
func NewMTRClusterEnergyEVSE() MTRClusterEnergyEVSE {
	return getMTRClusterEnergyEVSEClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterEnergyEVSE */
// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.


// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE
type MTRClusterEnergyEVSE struct {
	MTRGenericCluster
}

// MTRClusterEnergyEVSEFrom constructs a [MTRClusterEnergyEVSE] from an unsafe.Pointer.
//
// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.
func MTRClusterEnergyEVSEFrom(ptr unsafe.Pointer) MTRClusterEnergyEVSE {
	return MTRClusterEnergyEVSE{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterEnergyEVSE *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterEnergyEVSE */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterEnergyEVSE */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterEnergyEVSE */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/disable(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) DisableWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}/* debug [instance_methods/method]: DisableWithExpectedValuesExpectedValueIntervalCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterEnergyEVSE */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterEnergyEVSE */



