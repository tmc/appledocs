// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Address() string /* primitive/slice/pointer. */
	SetAddress(value string /* primitive/slice/pointer. */)
	ConfirmationCode() string /* primitive/slice/pointer. */
	SetConfirmationCode(value string /* primitive/slice/pointer. */)
	Eid() string /* primitive/slice/pointer. */
	SetEid(value string /* primitive/slice/pointer. */)
	Iccid() string /* primitive/slice/pointer. */
	SetIccid(value string /* primitive/slice/pointer. */)
	MatchingID() string /* primitive/slice/pointer. */
	SetMatchingID(value string /* primitive/slice/pointer. */)
	Oid() string /* primitive/slice/pointer. */
	SetOid(value string /* primitive/slice/pointer. */)
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



// The address of the carrier network’s eSIM server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/address
func (c_ CellularPlanProvisioningRequest) Address() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("address"))
	return rv
}


// The address of the carrier network’s eSIM server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/address
func (c_ CellularPlanProvisioningRequest) SetAddress(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAddress:"), objc.String(value))
}


// The provisioning request’s confirmation code, provided by the network operator when initiating an eSIM download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/confirmationcode
func (c_ CellularPlanProvisioningRequest) ConfirmationCode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("confirmationCode"))
	return rv
}


// The provisioning request’s confirmation code, provided by the network operator when initiating an eSIM download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/confirmationcode
func (c_ CellularPlanProvisioningRequest) SetConfirmationCode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfirmationCode:"), objc.String(value))
}


// The provisioning request’s eUICC identifier (EID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/eid
func (c_ CellularPlanProvisioningRequest) Eid() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("eid"))
	return rv
}


// The provisioning request’s eUICC identifier (EID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/eid
func (c_ CellularPlanProvisioningRequest) SetEid(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEid:"), objc.String(value))
}


// The provisioning request’s Integrated Circuit Card Identifier (ICCID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/iccid
func (c_ CellularPlanProvisioningRequest) Iccid() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("iccid"))
	return rv
}


// The provisioning request’s Integrated Circuit Card Identifier (ICCID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/iccid
func (c_ CellularPlanProvisioningRequest) SetIccid(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIccid:"), objc.String(value))
}


// The provisioning request’s matching identifier (MatchingID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/matchingid
func (c_ CellularPlanProvisioningRequest) MatchingID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("matchingID"))
	return rv
}


// The provisioning request’s matching identifier (MatchingID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/matchingid
func (c_ CellularPlanProvisioningRequest) SetMatchingID(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMatchingID:"), objc.String(value))
}


// The provisioning request’s Object Identifier (OID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/oid
func (c_ CellularPlanProvisioningRequest) Oid() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("oid"))
	return rv
}


// The provisioning request’s Object Identifier (OID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanprovisioningrequest/oid
func (c_ CellularPlanProvisioningRequest) SetOid(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOid:"), objc.String(value))
}



