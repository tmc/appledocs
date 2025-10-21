// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CellularPlanProvisioningRequest] class.
type ICellularPlanProvisioningRequest interface {
	objectivec.IObject
}

// A request specifying an eSIM to download and install.
//
// You must set the property for the request to be valid. All other properties are optional. This class is only available to carrier apps with suitable entitlements.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CellularPlanProvisioningRequestClass) Alloc() CellularPlanProvisioningRequest {
	rv := objc.Send[CellularPlanProvisioningRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The address of the carrier network’s eSIM server.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/address
func (c_ CellularPlanProvisioningRequest) Address() string {
	rv := objc.Send[string](c_.ID, objc.Sel("address"))
	return rv
}


// SetAddress sets the value of the address property.
// The address of the carrier network’s eSIM server.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/address
func (c_ CellularPlanProvisioningRequest) SetAddress(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAddress:"), objc.String(value))
}



