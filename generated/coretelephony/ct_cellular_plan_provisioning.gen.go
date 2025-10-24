// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCellularPlanProvisioning */


/* debug [class_header]: Header for CTCellularPlanProvisioning */
// The class instance for the [CellularPlanProvisioning] class.
var (
	CellularPlanProvisioningClass     _CellularPlanProvisioningClass
	CellularPlanProvisioningClassOnce sync.Once
)

func getCellularPlanProvisioningClass() _CellularPlanProvisioningClass {
	CellularPlanProvisioningClassOnce.Do(func() {
		CellularPlanProvisioningClass = _CellularPlanProvisioningClass{objc.GetClass("CTCellularPlanProvisioning")}
	})
	return CellularPlanProvisioningClass
}

type _CellularPlanProvisioningClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CellularPlanProvisioning */
// An interface definition for the [CellularPlanProvisioning] class.
type ICellularPlanProvisioning interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CellularPlanProvisioning */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CellularPlanProvisioning */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CellularPlanProvisioning */
// Alloc allocates a new instance without initialization.
func (cc _CellularPlanProvisioningClass) Alloc() CellularPlanProvisioning {
	rv := objc.Send[CellularPlanProvisioning](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CellularPlanProvisioningClass) New() CellularPlanProvisioning {
	rv := objc.Send[CellularPlanProvisioning](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularPlanProvisioning) Init() CellularPlanProvisioning {
	rv := objc.Send[CellularPlanProvisioning](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularPlanProvisioning) Autorelease() CellularPlanProvisioning {
	rv := objc.Send[CellularPlanProvisioning](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularPlanProvisioning creates a new CellularPlanProvisioning instance.
func NewCellularPlanProvisioning() CellularPlanProvisioning {
	return getCellularPlanProvisioningClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CellularPlanProvisioning */
// An object you use to download and install a carrier eSIM.
//
// This class is only available to carrier apps with suitable entitlements.


// An object you use to download and install a carrier eSIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning
type CellularPlanProvisioning struct {
	objectivec.Object
}

// CellularPlanProvisioningFrom constructs a [CellularPlanProvisioning] from an unsafe.Pointer.
//
// An object you use to download and install a carrier eSIM.
func CellularPlanProvisioningFrom(ptr unsafe.Pointer) CellularPlanProvisioning {
	return CellularPlanProvisioning{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CellularPlanProvisioning *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CellularPlanProvisioning */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CellularPlanProvisioning */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CellularPlanProvisioning */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CellularPlanProvisioning */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCellularPlanProvisioning */


