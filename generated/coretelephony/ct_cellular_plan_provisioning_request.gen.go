// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ConfirmationCode() objc.IObject /* cross-framework: NSString */
	SetConfirmationCode(value objc.IObject /* cross-framework: NSString */)
	Eid() objc.IObject /* cross-framework: NSString */
	SetEid(value objc.IObject /* cross-framework: NSString */)
	Iccid() objc.IObject /* cross-framework: NSString */
	SetIccid(value objc.IObject /* cross-framework: NSString */)
	MatchingID() objc.IObject /* cross-framework: NSString */
	SetMatchingID(value objc.IObject /* cross-framework: NSString */)
	Oid() objc.IObject /* cross-framework: NSString */
	SetOid(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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



// The provisioning request’s confirmation code, provided by the network operator when initiating an eSIM download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/confirmationcode
func (c_ CellularPlanProvisioningRequest) ConfirmationCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("confirmationCode"))
	return rv
}


// The provisioning request’s confirmation code, provided by the network operator when initiating an eSIM download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/confirmationcode
func (c_ CellularPlanProvisioningRequest) SetConfirmationCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfirmationCode:"), value)
}


// The provisioning request’s eUICC identifier (EID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/eid
func (c_ CellularPlanProvisioningRequest) Eid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("eid"))
	return rv
}


// The provisioning request’s eUICC identifier (EID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/eid
func (c_ CellularPlanProvisioningRequest) SetEid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEid:"), value)
}


// The provisioning request’s Integrated Circuit Card Identifier (ICCID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/iccid
func (c_ CellularPlanProvisioningRequest) Iccid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("iccid"))
	return rv
}


// The provisioning request’s Integrated Circuit Card Identifier (ICCID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/iccid
func (c_ CellularPlanProvisioningRequest) SetIccid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIccid:"), value)
}


// The provisioning request’s matching identifier (MatchingID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/matchingid
func (c_ CellularPlanProvisioningRequest) MatchingID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("matchingID"))
	return rv
}


// The provisioning request’s matching identifier (MatchingID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/matchingid
func (c_ CellularPlanProvisioningRequest) SetMatchingID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMatchingID:"), value)
}


// The provisioning request’s Object Identifier (OID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/oid
func (c_ CellularPlanProvisioningRequest) Oid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("oid"))
	return rv
}


// The provisioning request’s Object Identifier (OID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/oid
func (c_ CellularPlanProvisioningRequest) SetOid(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOid:"), value)
}


