// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CellularPlanProvisioning] class.
type ICellularPlanProvisioning interface {
	objectivec.IObject
	SupportsEmbeddedSIM() bool
}

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

// Alloc allocates a new instance without initialization.
func (cc _CellularPlanProvisioningClass) Alloc() CellularPlanProvisioning {
	rv := objc.Send[CellularPlanProvisioning](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether the device has hardware eSIM support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning/supportsEmbeddedSIM
func (c_ CellularPlanProvisioning) SupportsEmbeddedSIM() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsEmbeddedSIM"))
	return rv
}



