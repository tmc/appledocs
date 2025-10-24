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





