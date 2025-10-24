//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CellularPlanProvisioningRequest


// iOS-only properties

// The address of the carrier network’s eSIM server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/address
func (c_ CellularPlanProvisioningRequest) Address() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("address"))
	return rv
}
func (c_ CellularPlanProvisioningRequest) SetAddress(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setAddress:"), value)
}

// The provisioning request’s confirmation code, provided by the network operator when initiating an eSIM download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/confirmationCode
func (c_ CellularPlanProvisioningRequest) ConfirmationCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("confirmationCode"))
	return rv
}
func (c_ CellularPlanProvisioningRequest) SetConfirmationCode(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setConfirmationCode:"), value)
}

// The provisioning request’s eUICC identifier (EID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/eid
func (c_ CellularPlanProvisioningRequest) EID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("EID"))
	return rv
}
func (c_ CellularPlanProvisioningRequest) SetEID(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setEID:"), value)
}

// The provisioning request’s Integrated Circuit Card Identifier (ICCID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/iccid
func (c_ CellularPlanProvisioningRequest) ICCID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ICCID"))
	return rv
}
func (c_ CellularPlanProvisioningRequest) SetICCID(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setICCID:"), value)
}

// The provisioning request’s matching identifier (MatchingID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/matchingID
func (c_ CellularPlanProvisioningRequest) MatchingID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("matchingID"))
	return rv
}
func (c_ CellularPlanProvisioningRequest) SetMatchingID(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setMatchingID:"), value)
}

// The provisioning request’s Object Identifier (OID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningRequest/oid
func (c_ CellularPlanProvisioningRequest) OID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("OID"))
	return rv
}
func (c_ CellularPlanProvisioningRequest) SetOID(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setOID:"), value)
}





