// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCellularPlanProvisioningRequest */


/* debug [class_header]: Header for CTCellularPlanProvisioningRequest */
// The class instance for the [CellularPlanProvisioningRequest] class.
var (
	CellularPlanProvisioningRequestClass     _CellularPlanProvisioningRequestClass
	CellularPlanProvisioningRequestClassOnce sync.Once
)

func getCellularPlanProvisioningRequestClass() _CellularPlanProvisioningRequestClass {
	CellularPlanProvisioningRequestClassOnce.Do(func() {
		CellularPlanProvisioningRequestClass = _CellularPlanProvisioningRequestClass{objc.GetClass("CTCellularPlanProvisioningRequest")}
	})
	return CellularPlanProvisioningRequestClass
}

type _CellularPlanProvisioningRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CellularPlanProvisioningRequest */
// An interface definition for the [CellularPlanProvisioningRequest] class.
type ICellularPlanProvisioningRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CellularPlanProvisioningRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CellularPlanProvisioningRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CellularPlanProvisioningRequest */
// Alloc allocates a new instance without initialization.
func (cc _CellularPlanProvisioningRequestClass) Alloc() CellularPlanProvisioningRequest {
	rv := objc.Send[CellularPlanProvisioningRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CellularPlanProvisioningRequestClass) New() CellularPlanProvisioningRequest {
	rv := objc.Send[CellularPlanProvisioningRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularPlanProvisioningRequest) Init() CellularPlanProvisioningRequest {
	rv := objc.Send[CellularPlanProvisioningRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularPlanProvisioningRequest) Autorelease() CellularPlanProvisioningRequest {
	rv := objc.Send[CellularPlanProvisioningRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularPlanProvisioningRequest creates a new CellularPlanProvisioningRequest instance.
func NewCellularPlanProvisioningRequest() CellularPlanProvisioningRequest {
	return getCellularPlanProvisioningRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CellularPlanProvisioningRequest */
// A request specifying an eSIM to download and install.
//
// You must set the property for the request to be valid. All other properties are optional. This class is only available to carrier apps with suitable entitlements.


// A request specifying an eSIM to download and install.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest
type CellularPlanProvisioningRequest struct {
	objectivec.Object
}

// CellularPlanProvisioningRequestFrom constructs a [CellularPlanProvisioningRequest] from an unsafe.Pointer.
//
// A request specifying an eSIM to download and install.
func CellularPlanProvisioningRequestFrom(ptr unsafe.Pointer) CellularPlanProvisioningRequest {
	return CellularPlanProvisioningRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CellularPlanProvisioningRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CellularPlanProvisioningRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CellularPlanProvisioningRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CellularPlanProvisioningRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CellularPlanProvisioningRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCellularPlanProvisioningRequest */


